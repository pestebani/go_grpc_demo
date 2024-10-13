package dblayer

import (
	"errors"
	"go_grpc_demo/pkg/database"
	"go_grpc_demo/pkg/database/postgresdb"
	"os"
)

func NewDBLayer() (database.Database, error) {
	switch os.Getenv("DATABASE_TYPE") {
	case "postgres":
		return postgresdb.NewPostgresDB()
	default:
		return nil, errors.New("database not supported")
	}
}
