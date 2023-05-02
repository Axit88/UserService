package incoming

import (
	pb "github.com/Axit88/UserGrpc/storage-proto"
	"github.com/Axit88/UserService/src/domain/userService/core/model"
)

type UserService interface {
	AddUserClient(grpc pb.TestApiClient, input *model.User) error
	GetUserClient(grpc pb.TestApiClient, userId string) (*model.User, error)
	DeleteUserClient(grpc pb.TestApiClient, userId string) error
	UpdateUserClient(grpc pb.TestApiClient, userId string, userName string) error
}

