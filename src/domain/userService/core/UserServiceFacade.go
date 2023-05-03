package infrastructure

import (
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
	return err
}

func (facade *UserServiceFacade) GetUser(userId string) (*model.User, error) {
	res, err := facade.userServiceClient.GetUser(userId)
	return res, err
}

func (facade *UserServiceFacade) UpdateUser(userId string, userName string) error {
	err := facade.userServiceClient.UpdateUser(userId, userName)
	return err
}

func (facade *UserServiceFacade) DeleteUser(userId string) error {
	err := facade.userServiceClient.DeleteUser(userId)
	return err
}
