package infrastructure

import (
	"context"

	"github.com/Axit88/UserService/src/domain/userService/core/model"
	"github.com/Axit88/UserService/src/domain/userService/core/ports/incoming"
	"github.com/Axit88/UserService/src/domain/userService/core/ports/outgoing"
	"github.com/MindTickle/mt-go-logger/logger"
)

type UserServiceFacade struct {
	userServiceClient outgoing.UserServiceClient
	logger            *logger.LoggerImpl
}

func NewUserServiceFacade(newUserServiceClient outgoing.UserServiceClient, l *logger.LoggerImpl) incoming.UserService {
	return &UserServiceFacade{
		userServiceClient: newUserServiceClient,
		logger:            l,
	}
}

func (facade *UserServiceFacade) AddUser(input *model.User) error {
	err := facade.userServiceClient.AddUser(input)
	if err != nil {
		facade.logger.Errorf(context.Background(), "Failed To Process Outgoing AddUser Request")
	}
	return err
}

func (facade *UserServiceFacade) GetUser(userId string) (*model.User, error) {
	res, err := facade.userServiceClient.GetUser(userId)
	if err != nil {
		facade.logger.Errorf(context.Background(), "Failed To Process Outgoing GetUser Request")
	}
	return res, err
}

func (facade *UserServiceFacade) UpdateUser(userId string, userName string) error {
	err := facade.userServiceClient.UpdateUser(userId, userName)
	if err != nil {
		facade.logger.Errorf(context.Background(), "Failed To Process Outgoing UpdateUser Request")
	}
	return err
}

func (facade *UserServiceFacade) DeleteUser(userId string) error {
	err := facade.userServiceClient.DeleteUser(userId)
	if err != nil {
		facade.logger.Errorf(context.Background(), "Failed To Process Outgoing DeleteUser Request")
	}
	return err
}
