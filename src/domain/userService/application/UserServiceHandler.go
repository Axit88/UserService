package application

import (
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
	return worker.facade.AddUser(&input)
}

func (worker *UserServiceApplication) GetUserApplication(userId string) (*model.User, error) {
	return worker.facade.GetUser(userId)
	//return infrastructure.GetUserRest(userId)
}
