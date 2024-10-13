package model

import v1agenda "go_grpc_demo/pkg/agenda_server/v1"

type Agenda struct {
	ID    int
	Name  string
	Email string
	Phone string
}

func (ag *Agenda) Encode() (*v1agenda.Agenda, error) {
	return &v1agenda.Agenda{
		Id:    int64(ag.ID),
		Name:  ag.Name,
		Email: ag.Email,
		Phone: ag.Phone,
	}, nil
}

func (ag *Agenda) Decode(a *v1agenda.Agenda) error {
	ag.ID = int(a.Id)
	ag.Name = a.Name
	ag.Email = a.Email
	ag.Phone = a.Phone
	return nil
}
