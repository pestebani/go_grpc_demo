package database

import (
	"context"
	"errors"
	"go_grpc_demo/pkg/model"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var IdNotExistsError = errors.New("database: ID does not exist")
var UnimplementedError = errors.New("database: Unimplemented")
var AlreadyExistsError = errors.New("database: Already exists")

type Database interface {
	Initiate() error
	RetrieveFromDatabase(context.Context, int) (model.Agenda, error)
	RetrieveListFromDatabase(ctx context.Context, page int, elementsPage int) (agendas []model.Agenda, nextPage int, TotalAgendas int, err error)
	StoreInDatabase(context.Context, model.Agenda) (model.Agenda, error)
	UpdateInDatabase(context.Context, int, model.Agenda) (model.Agenda, error)
	DeleteFromDatabase(context.Context, int) error
	Close() error
}

func ConvertErrorToGRPCStatus(err error) error {
	switch {
	case errors.Is(err, IdNotExistsError):
		return status.Error(codes.NotFound, err.Error())
	case errors.Is(err, UnimplementedError):
		return status.Error(codes.Unimplemented, err.Error())
	case errors.Is(err, AlreadyExistsError):
		return status.Error(codes.AlreadyExists, err.Error())
	default:
		return status.Error(codes.Unknown, err.Error())
	}
}
