package application

import (
	"github.com/Axit88/UserService/src/domain/userService/core/model"
	"github.com/Axit88/UserService/src/domain/userService/core/ports/incoming"
	"github.com/Axit88/UserService/src/domain/userService/core/ports/outgoing"
)

type UserServiceApplication struct {
	facade incoming.UserService
	grpc   outgoing.GrpcClient
}

func NewUserServiceApplication(grpc outgoing.GrpcClient, facade incoming.UserService) *UserServiceApplication {
	res := new(UserServiceApplication)
	res.grpc = grpc
	res.facade = facade
	return res
}

func (worker *UserServiceApplication) AddUserApplication(userId string, userName string) error {
	input := model.User{
		UserId:   userId,
		Username: userName,
	}
	grpc_client, err := worker.grpc.CreateGrpcConnection()
	if err != nil {
		return err
	}
	return worker.facade.AddUserClient(grpc_client, &input)
}

func (worker *UserServiceApplication) GetUserApplication(userId string) (*model.User, error) {
	grpc_client, err := worker.grpc.CreateGrpcConnection()
	if err != nil {
		return nil, err
	}
	return worker.facade.GetUserClient(grpc_client, userId)
}
