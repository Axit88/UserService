package incoming

import (
	"github.com/Axit88/UserService/src/domain/userService/core/model"
)

type ContentSync interface {
	AddUser(request *model.User) error
	GetUser(userId string) (*model.User,error)
}

