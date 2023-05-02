package outgoing

import (
	pb "github.com/Axit88/UserGrpc/storage-proto"
)

type GrpcClient interface {
	CreateGrpcConnection() (pb.TestApiClient,error)
}