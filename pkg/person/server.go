package person

import (
	"context"

	"github.com/habuvo/testserverone/proto"
)

type personServer struct {
	proto.UnimplementedPersonServiceServer
}

func New() personServer {
	return personServer{}
}

func (p personServer) GetPerson(ctx context.Context, in *proto.GetRequest) (*proto.Person, error) {
	return &proto.Person{
		Name: 0,
	}, nil
}

func (p personServer) SetPerson(ctx context.Context, person *proto.Person) (*proto.SetResponse, error) {
	return &proto.SetResponse{
		Id: 0,
	}, nil
}
