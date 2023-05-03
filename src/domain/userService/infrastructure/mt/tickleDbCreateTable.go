package main

import (
	"context"
	"fmt"

	"github.com/Axit88/UserService/src/constants"
	"github.com/Axit88/UserService/src/domain/userService/core/model"
	outgoing "github.com/Axit88/UserService/src/domain/userService/core/ports/outgoing/outgoingMt"
	pb "github.com/MindTickle/storageprotos/pb/tickleDb"
	"google.golang.org/grpc"
)

type TickleDbStoreImplImpl struct {
	TickleDbStoreImplService pb.StoreManagerClient
}

func NewTickleDbStoreImplClient() outgoing.TickleDbCreateTableClient {
	connection := fmt.Sprintf("%v:%v", constants.GRPC_HOST, constants.GRPC_PORT)
	conn, err := grpc.Dial(connection, grpc.WithInsecure())
	if err != nil {
		return nil
	}
	res := TickleDbStoreImplImpl{}
	res.TickleDbStoreImplService = pb.NewStoreManagerClient(conn)
	return res
}

func (client TickleDbStoreImplImpl) CreateTable(dbDetail model.TickleDbEnvDetail) error {

	myTable := &pb.Table{
		TableName: dbDetail.TableName,
		Ttl:       0,
		Version:   0,
		Namespace: dbDetail.Namespace,
		Env:       dbDetail.Env,
		Columns: []*pb.Field{
			&pb.Field{
				FieldName:          "uid",
				DataType:           1,
				EnumExpectedValues: nil,
				NestedFields:       nil,
				Required:           true,
				Size:               10,
				DefaultValue:       "1",
			},
			&pb.Field{
				FieldName:          "uname",
				DataType:           1,
				EnumExpectedValues: nil,
				NestedFields:       nil,
				Required:           true,
				Size:               255,
				DefaultValue:       "",
			},
		},
		PrimaryKey: &pb.PrimaryKey{Columns: []string{"uid"}},
		//IndexColumns:       []*pb.IndexField{{FieldPath: []string{"uname"}},},
		PartitionStrategy: pb.PartitionStrategy_HASH_BASED,
		PartitionKey:      "",
	}

	res, err := client.TickleDbStoreImplService.CreateTable(context.Background(), &pb.CreateTableRequest{Table: myTable})
	if err != nil {
		fmt.Println("Create Table error", err)
		return err
	}
	fmt.Println(res)
	return nil
}
