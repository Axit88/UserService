package application

import (
	"context"

	"github.com/Axit88/UserService/src/domain/userService/core/model"
	"github.com/Axit88/UserService/src/domain/userService/core/ports/incoming"
	"github.com/MindTickle/mt-go-logger/logger"
)

type UserServiceApplication struct {
	facade incoming.UserService
	Logger *logger.LoggerImpl
}

func NewUserServiceApplication(facade incoming.UserService, l *logger.LoggerImpl) *UserServiceApplication {
	res := new(UserServiceApplication)
	res.facade = facade
	res.Logger = l
	return res
}

func (worker *UserServiceApplication) AddUserApplication(userId string, userName string) error {
	input := model.User{
		UserId:   userId,
		Username: userName,
	}

	ctx := context.Background()
	err := worker.facade.AddUser(&input)
	if err != nil {
		worker.Logger.Errorf(ctx, "Failed To Process Add User Application Request", err)
	}
	return err
}

func (worker *UserServiceApplication) GetUserApplication(userId string) (*model.User, error) {
	res, err := worker.facade.GetUser(userId)
	ctx := context.Background()
	if err != nil {
		worker.Logger.Errorf(ctx, "Failed To Process Get User Application Request", err)
	}
	return res, err
	// return infrastructure.GetUserRest(userId)
}
