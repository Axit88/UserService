package infrastructure

import (
	"context"

	pb "github.com/Axit88/UserGrpc/storage-proto"
	"github.com/Axit88/UserService/src/domain/userService/core/model"
	"github.com/Axit88/UserService/src/domain/userService/core/ports/incoming"
	"github.com/Axit88/UserService/src/domain/userService/core/ports/outgoing"
)

type UserServiceFacade struct {
	grpc outgoing.GrpcClient
}

func NewUserServiceFacade(grpcCliet outgoing.GrpcClient) incoming.UserService {
	return &UserServiceFacade{
		grpc: grpcCliet,
	}
}

func (facade *UserServiceFacade) AddUserClient(client pb.TestApiClient, input *model.User) error {
	in := &pb.AddUserInput{
		UserName: input.Username,
		UserId:   input.UserId,
	}
	_, err := client.AddUser(context.Background(), in)
	return err
}

func (facade *UserServiceFacade) GetUserClient(client pb.TestApiClient, userId string) (*model.User, error) {
	in := &pb.GetUserInput{
		UserId: userId,
	}
	res, err := client.GetUser(context.Background(), in)
	output := model.User{
		UserId:   res.UserId,
		Username: res.UserName,
	}
	return &output, err
}

func (facade *UserServiceFacade) UpdateUserClient(client pb.TestApiClient, userId string, userName string) error {
	input := pb.UpdateUserInput{
		UserId:   userId,
		UserName: userName,
	}
	_, err := client.UpdateUser(context.Background(), &input)
	return err
}

func (facade *UserServiceFacade) DeleteUserClient(client pb.TestApiClient, userId string) error {
	in := &pb.DeleteUserInput{
		UserId: userId,
	}
	_, err := client.DeleteUser(context.Background(), in)
	return err
}
