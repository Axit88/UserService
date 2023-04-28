package infrastructure

import (
	"fmt"

	"github.com/Axit88/UserService/src/constants"
	pb "github.com/Axit88/UserService/src/domain/userService/infrastructure/Grpc/storage-proto"
	"google.golang.org/grpc"
)

func CreateGrpcConnection() (pb.TestApiClient, error) {
	connection := fmt.Sprintf("%v:%v", constants.GRPC_HOST, constants.GRPC_PORT)
	conn, err := grpc.Dial(connection, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	client := pb.NewTestApiClient(conn)
	return client, nil
}
