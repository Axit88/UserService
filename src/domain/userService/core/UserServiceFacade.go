package infrastructure

import (
	"context"

	"github.com/Axit88/UserService/src/domain/userService/core/model"
	"github.com/Axit88/UserService/src/domain/userService/core/ports/incoming"
	"github.com/Axit88/UserService/src/domain/userService/infrastructure"
	pb "github.com/Axit88/UserService/src/domain/userService/infrastructure/Grpc/storage-proto"
)

type UserServiceFacade struct {
}

func NewUserServiceFacade() incoming.ContentSync {
	return &UserServiceFacade{}
}

func (facade *UserServiceFacade) AddUser(input *model.User) error {
	client, err := infrastructure.CreateGrpcConnection()
	if err != nil {
		return err
	}

	in := &pb.AddUserInput{
		UserName: input.Username,
		UserId:   input.UserId,
	}
	_, err = client.AddUser(context.Background(), in)
	return err
}

func (facade *UserServiceFacade) GetUser(InputUserId string) (*model.User, error) {
	client, err := infrastructure.CreateGrpcConnection()
	if err != nil {
		return nil,err
	}

	in := &pb.GetUserInput{
		UserId: InputUserId,
	}
	res, err := client.GetUser(context.Background(), in)
	output := model.User{
		UserId:   res.UserId,
		Username: res.UserName,
	}
	return &output, err
}
