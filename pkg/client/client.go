package client

import (
	"context"
	pb "go_grpc_demo/pkg/agenda_server/v1"
	"go_grpc_demo/pkg/model"
	"google.golang.org/grpc"
)

type Client struct {
	agendaClient pb.AgendaServiceClient
}

func NewClient(target string, opts ...grpc.DialOption) (Client, error) {
	cl, err := grpc.NewClient(target, opts...)
	clAgenda := pb.NewAgendaServiceClient(cl)
	return Client{
		agendaClient: clAgenda,
	}, err
}

func (c *Client) Ping(ctx context.Context) error {
	_, err := c.agendaClient.Ping(ctx, new(pb.PingRequest))
	return err
}

func (c *Client) CreateAgenda(ctx context.Context, ag model.Agenda) (model.Agenda, error) {
	agInProto, err := ag.Encode()

	if err != nil {
		return model.Agenda{}, err
	}

	agOutProto, err := c.agendaClient.CreateAgenda(ctx, &pb.CreateAgendaRequest{Agenda: agInProto})

	if err != nil {
		return model.Agenda{}, err
	}

	agOut := new(model.Agenda)

	err = agOut.Decode(agOutProto.GetAgenda())

	return *agOut, err
}

func (c *Client) GetAgenda(ctx context.Context, id int) (model.Agenda, error) {
	agOutProto, err := c.agendaClient.GetAgenda(ctx, &pb.GetAgendaRequest{Id: int64(id)})
	if err != nil {
		return model.Agenda{}, err
	}

	agOut := new(model.Agenda)

	err = agOut.Decode(agOutProto.GetAgenda())

	return *agOut, err
}

func (c *Client) GetAgendas(ctx context.Context, page int, elementsPage int) ([]model.Agenda, int, int, error) {
	agOutProto, err := c.agendaClient.GetAgendas(ctx, &pb.GetAgendasRequest{Page: int64(page), Items: int64(elementsPage)})

	if err != nil {
		return nil, 0, 0, err
	}

	agOut := make([]model.Agenda, len(agOutProto.Agendas))

	for i, ag := range agOutProto.Agendas {
		agOut[i].Decode(ag)
	}

	return agOut, int(agOutProto.GetNextPage()), int(agOutProto.GetTotal()), nil
}

func (c *Client) UpdateAgenda(ctx context.Context, id int, ag model.Agenda) (model.Agenda, error) {
	agInProto, err := ag.Encode()

	if err != nil {
		return model.Agenda{}, err
	}

	agOutProto, err := c.agendaClient.UpdateAgenda(ctx, &pb.UpdateAgendaRequest{Id: int64(id), Agenda: agInProto})

	if err != nil {
		return model.Agenda{}, err
	}

	agOut := new(model.Agenda)

	err = agOut.Decode(agOutProto.GetAgenda())

	return *agOut, err
}

func (c *Client) DeleteAgenda(ctx context.Context, id int) error {
	_, err := c.agendaClient.DeleteAgenda(ctx, &pb.DeleteAgendaRequest{Id: int64(id)})
	return err
}
