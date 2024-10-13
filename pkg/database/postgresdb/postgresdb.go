package postgresdb

import (
	"context"
	"database/sql"
	"errors"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
	"go.opentelemetry.io/contrib/bridges/otelslog"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go_grpc_demo/pkg/database"
	"go_grpc_demo/pkg/model"
	"os"
)

const name = "go_grpc_server.agenda-database"

var (
	tracer = otel.Tracer(name)
	logger = otelslog.NewLogger(name)
)

type PostgresDB struct {
	database *sql.DB
}

// NewPostgresDB creates a new connection to database
func NewPostgresDB() (*PostgresDB, error) {
	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		connStr = "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
	}
	// Connect to the database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		logger.Error("Failed to connect to the database", "error", err)
		return nil, err
	}

	// Check if the connection is successful
	err = db.Ping()
	if err != nil {
		logger.Error("Failed to connect to the database", "error", err)
		return nil, err
	}

	logger.Info("Connected to the database")

	return &PostgresDB{
		database: db,
	}, nil
}

// Close closes the database connection
func (p *PostgresDB) Close() error {
	return p.database.Close()
}

// Initiate the database connection
func (p *PostgresDB) Initiate() error {
	// Create the table called agenda if it does not exist. It has an autoincrement field called id, a name field, an email field, and a phone field.
	// name field is of type TEXT and unique, email field is of type TEXT, and phone field is of type TEXT.
	query := `CREATE TABLE IF NOT EXISTS agenda (
			id SERIAL PRIMARY KEY,
			name TEXT UNIQUE NOT NULL,
			email TEXT NOT NULL,
			phone TEXT NOT NULL);`

	_, err := p.database.Exec(query)

	if err != nil {
		logger.Error("Failed to create the table", "error", err)
		return err
	}

	logger.Info("Table created")

	return nil
}

// RetrieveFromDatabase retrieves an agenda from the database
func (p *PostgresDB) RetrieveFromDatabase(ctx context.Context, id int) (model.Agenda, error) {
	ctx, span := tracer.Start(context.Background(), "RetrieveFromDatabase")
	defer span.End()

	span.SetAttributes(attribute.Int("id", id))

	query := `SELECT * FROM agenda WHERE id = $1;`
	row := p.database.QueryRow(query, id)

	var ag model.Agenda

	err := row.Scan(&ag.ID, &ag.Name, &ag.Email, &ag.Phone)

	if err != nil {
		logger.ErrorContext(ctx, "Failed to retrieve from the database", "error", err)
		span.RecordError(err)
		return ag, convertPostgresErrorToDBError(err)
	}

	logger.InfoContext(ctx, "Retrieved from the database", "model", ag)

	return ag, nil
}

// RetrieveListFromDatabase retrieves a list of agendas from the database
func (p *PostgresDB) RetrieveListFromDatabase(ctx context.Context, page int, elementsPage int) ([]model.Agenda, int, int, error) {
	ctx, span := tracer.Start(context.Background(), "RetrieveListFromDatabase")
	defer span.End()

	span.SetAttributes(attribute.Int("page", page))
	span.SetAttributes(attribute.Int("elementsPage", elementsPage))

	query := `SELECT (SELECT COUNT(*) FROM agenda) as total, id, name, email, phone FROM agenda LIMIT $1 OFFSET $2;`
	rows, err := p.database.Query(query, elementsPage, (page-1)*elementsPage)

	if err != nil {
		logger.ErrorContext(ctx, "Failed to retrieve from the database", "error", err)
		span.RecordError(err)
		return nil, 0, 0, err
	}

	var ags []model.Agenda
	var totalAgendas int

	for rows.Next() {
		var ag model.Agenda
		err := rows.Scan(&totalAgendas, &ag.ID, &ag.Name, &ag.Email, &ag.Phone)

		if err != nil {
			logger.ErrorContext(ctx, "Failed to retrieve from the database", "error", err)
			span.RecordError(err)
			return nil, 0, 0, err
		}

		ags = append(ags, ag)
	}

	logger.InfoContext(ctx, "Retrieved from the database", "models", ags)

	if page*elementsPage >= totalAgendas {
		page = -1
	}

	return ags, page + 1, totalAgendas, nil
}

// StoreInDatabase stores an agenda in the database
func (p *PostgresDB) StoreInDatabase(ctx context.Context, ag model.Agenda) (model.Agenda, error) {
	ctx, span := tracer.Start(context.Background(), "StoreInDatabase")
	defer span.End()

	span.SetAttributes(attribute.String("name", ag.Name))
	span.SetAttributes(attribute.String("email", ag.Email))
	span.SetAttributes(attribute.String("phone", ag.Phone))

	query := `INSERT INTO agenda (name, email, phone) VALUES ($1, $2, $3) RETURNING id, name, email, phone;`
	row := p.database.QueryRow(query, ag.Name, ag.Email, ag.Phone)

	var agRet model.Agenda

	err := row.Scan(&agRet.ID, &agRet.Name, &agRet.Email, &agRet.Phone)

	if err != nil {
		logger.ErrorContext(ctx, "Failed to store in the database", "error", err)
		span.RecordError(err)
		return agRet, convertPostgresErrorToDBError(err)
	}

	logger.InfoContext(ctx, "Stored in the database", "model", agRet)

	return agRet, nil
}

// UpdateInDatabase updates an agenda in the database
func (p *PostgresDB) UpdateInDatabase(ctx context.Context, id int, ag model.Agenda) (model.Agenda, error) {
	ctx, span := tracer.Start(context.Background(), "UpdateInDatabase")
	defer span.End()

	span.SetAttributes(attribute.Int("id", id))
	span.SetAttributes(attribute.String("name", ag.Name))
	span.SetAttributes(attribute.String("email", ag.Email))
	span.SetAttributes(attribute.String("phone", ag.Phone))

	query := `UPDATE agenda SET name = $1, email = $2, phone = $3 WHERE id = $4 RETURNING id, name, email, phone;`
	row := p.database.QueryRow(query, ag.Name, ag.Email, ag.Phone, id)

	var agRet model.Agenda

	err := row.Scan(&agRet.ID, &agRet.Name, &agRet.Email, &agRet.Phone)

	if err != nil {
		logger.ErrorContext(ctx, "Failed to update in the database", "error", err)
		span.RecordError(err)
		return agRet, convertPostgresErrorToDBError(err)
	}

	logger.InfoContext(ctx, "Updated in the database", "model", agRet)

	return agRet, nil
}

// DeleteFromDatabase deletes an agenda from the database
func (p *PostgresDB) DeleteFromDatabase(ctx context.Context, id int) error {
	ctx, span := tracer.Start(context.Background(), "DeleteFromDatabase")
	defer span.End()

	span.SetAttributes(attribute.Int("id", id))

	query := `DELETE FROM agenda WHERE id = $1;`
	res, err := p.database.Exec(query, id)

	if err != nil {
		logger.ErrorContext(ctx, "Failed to delete from the database", "error", err)
		span.RecordError(err)
		return err
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		logger.ErrorContext(ctx, "Failed to delete from the database", "error", err)
		span.RecordError(err)
		return err
	}

	if rowsAffected == 0 {
		logger.ErrorContext(ctx, "Failed to delete from the database", "error", "no rows affected")
		span.RecordError(err)
		return database.IdNotExistsError
	}

	logger.InfoContext(ctx, "Deleted from the database", "id", id)

	return nil
}

func convertPostgresErrorToDBError(err error) error {
	var pqErr *pq.Error
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return database.IdNotExistsError
	case errors.As(err, &pqErr):
		if pqErr.Code == "23505" {
			return database.AlreadyExistsError
		}
		return err
	default:
		return err
	}
}
