package infrastructure

import (
	"fmt"

	pb "github.com/Axit88/UserGrpc/storage-proto"
	"github.com/Axit88/UserService/src/constants"
	"github.com/Axit88/UserService/src/domain/userService/core/ports/outgoing"
	"google.golang.org/grpc"
)

type OutgoingGrpc struct {
}

func NewOutgoingGrpcClient() outgoing.GrpcClient {
	return &OutgoingGrpc{}
}

func (client OutgoingGrpc) CreateGrpcConnection() (pb.TestApiClient, error) {
	connection := fmt.Sprintf("%v:%v", constants.GRPC_HOST, constants.GRPC_PORT)
	conn, err := grpc.Dial(connection, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	grpcclient := pb.NewTestApiClient(conn)
	return grpcclient, nil
}
