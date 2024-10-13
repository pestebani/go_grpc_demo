package service

import (
	"context"
	"go.opentelemetry.io/contrib/bridges/otelslog"
	"go.opentelemetry.io/otel"
	v1agenda "go_grpc_demo/pkg/agenda_server/v1"
	"go_grpc_demo/pkg/database"
	"go_grpc_demo/pkg/dblayer"
	"go_grpc_demo/pkg/model"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const name = "go_grpc_server.service"

var (
	tracer = otel.Tracer(name)
	logger = otelslog.NewLogger(name)
)

type Service struct {
	v1agenda.UnimplementedAgendaServiceServer
	database database.Database
}

func NewService() (*Service, error) {
	db, err := dblayer.NewDBLayer()
	if err != nil {
		return nil, database.ConvertErrorToGRPCStatus(err)
	}

	err = db.Initiate()
	return &Service{
		database: db,
	}, err
}

func (s *Service) Close() error {
	return s.database.Close()
}

func (s *Service) Ping(ctx context.Context, _ *v1agenda.PingRequest) (*v1agenda.PingResponse, error) {
	ctx, span := tracer.Start(ctx, "Ping")
	defer span.End()

	logger.InfoContext(ctx, "Ping function", "result", "Pong")

	return &v1agenda.PingResponse{
		Response: "Pong",
	}, status.Error(codes.OK, "")
}

func (s *Service) CreateAgenda(ctx context.Context, ag *v1agenda.CreateAgendaRequest) (*v1agenda.CreateAgendaResponse, error) {
	ctx, span := tracer.Start(ctx, "CreateAgenda")
	defer span.End()

	logger.InfoContext(ctx, "going to create an agenda")

	var agm model.Agenda

	err := agm.Decode(ag.GetAgenda())

	if err != nil {
		logger.ErrorContext(ctx, "Error decoding agenda", "error", err)
		span.RecordError(err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	mod, err := s.database.StoreInDatabase(ctx, agm)

	if err != nil {
		logger.ErrorContext(ctx, "Error storing in database", "error", err)
		span.RecordError(err)
		return nil, database.ConvertErrorToGRPCStatus(err)
	}

	agEnc, err := mod.Encode()

	if err != nil {
		logger.ErrorContext(ctx, "Error encoding agenda", "error", err)
		span.RecordError(err)
		return nil, status.Error(codes.Unknown, err.Error())
	}

	return &v1agenda.CreateAgendaResponse{
		Agenda: agEnc,
	}, status.Error(codes.OK, "")
}

func (s *Service) GetAgenda(ctx context.Context, ag *v1agenda.GetAgendaRequest) (*v1agenda.GetAgendaResponse, error) {
	ctx, span := tracer.Start(ctx, "GetAgenda")
	defer span.End()

	logger.InfoContext(ctx, "retrieving agenda with id", "id", ag.GetId())

	id := int(ag.GetId())

	mod, err := s.database.RetrieveFromDatabase(ctx, id)

	if err != nil {
		logger.ErrorContext(ctx, "Error retrieving from database", "error", err)
		span.RecordError(err)
		return nil, database.ConvertErrorToGRPCStatus(err)
	}

	agEnc, err := mod.Encode()

	if err != nil {
		logger.ErrorContext(ctx, "Error encoding agenda", "error", err)
		span.RecordError(err)
		return nil, status.Error(codes.Unknown, err.Error())
	}

	return &v1agenda.GetAgendaResponse{
		Agenda: agEnc,
	}, status.Error(codes.OK, "")
}

func (s *Service) GetAgendas(ctx context.Context, ag *v1agenda.GetAgendasRequest) (*v1agenda.GetAgendasResponse, error) {
	ctx, span := tracer.Start(ctx, "GetAgendas")
	defer span.End()

	logger.InfoContext(ctx, "retrieving agendas")

	page := int(ag.GetPage())
	elementsPage := int(ag.GetItems())

	mod, nextPage, total, err := s.database.RetrieveListFromDatabase(ctx, page, elementsPage)

	if err != nil {
		logger.ErrorContext(ctx, "Error retrieving list from database", "error", err)
		span.RecordError(err)
		return nil, database.ConvertErrorToGRPCStatus(err)
	}

	var ags []*v1agenda.Agenda

	for _, ag := range mod {
		agEnc, err := ag.Encode()

		if err != nil {
			logger.ErrorContext(ctx, "Error encoding agenda", "error", err)
			span.RecordError(err)
			return nil, status.Error(codes.Unknown, err.Error())
		}

		ags = append(ags, agEnc)
	}

	return &v1agenda.GetAgendasResponse{
		Agendas:  ags,
		Total:    int64(total),
		NextPage: int64(nextPage),
	}, status.Error(codes.OK, "")
}

func (s *Service) UpdateAgenda(ctx context.Context, ag *v1agenda.UpdateAgendaRequest) (*v1agenda.UpdateAgendaResponse, error) {
	ctx, span := tracer.Start(ctx, "UpdateAgenda")
	defer span.End()

	logger.InfoContext(ctx, "updating agenda with id", "id", ag.GetId())

	id := int(ag.GetId())

	var agm model.Agenda

	err := agm.Decode(ag.GetAgenda())

	if err != nil {
		logger.ErrorContext(ctx, "Error decoding agenda", "error", err)
		span.RecordError(err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	mod, err := s.database.UpdateInDatabase(ctx, id, agm)

	if err != nil {
		logger.ErrorContext(ctx, "Error updating in database", "error", err)
		span.RecordError(err)
		return nil, database.ConvertErrorToGRPCStatus(err)
	}

	agEnc, err := mod.Encode()

	if err != nil {
		logger.ErrorContext(ctx, "Error encoding agenda", "error", err)
		span.RecordError(err)
		return nil, status.Error(codes.Unknown, err.Error())
	}

	return &v1agenda.UpdateAgendaResponse{
		Agenda: agEnc,
	}, status.Error(codes.OK, "")
}

func (s *Service) DeleteAgenda(ctx context.Context, ag *v1agenda.DeleteAgendaRequest) (*v1agenda.DeleteAgendaResponse, error) {
	ctx, span := tracer.Start(ctx, "DeleteAgenda")
	defer span.End()

	logger.InfoContext(ctx, "deleting agenda with id", "id", ag.GetId())

	id := int(ag.GetId())

	err := s.database.DeleteFromDatabase(ctx, id)

	if err != nil {
		logger.ErrorContext(ctx, "Error deleting from database", "error", err)
		span.RecordError(err)
		return nil, database.ConvertErrorToGRPCStatus(err)
	}

	return &v1agenda.DeleteAgendaResponse{}, status.Error(codes.OK, "")
}
