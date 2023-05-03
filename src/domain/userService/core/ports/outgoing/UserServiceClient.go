package outgoing

import (
	"github.com/Axit88/UserService/src/domain/userService/core/model"
)

type UserServiceClient interface {
	AddUser(input *model.User) error
	GetUser(userId string) (*model.User, error)
	DeleteUser(userId string) error
	UpdateUser(userId string, userName string) error
}
