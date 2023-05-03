package main

import (
	"context"
	"fmt"

	"github.com/Axit88/UserService/src/constants"
	"github.com/Axit88/UserService/src/domain/userService/core/model"
	outgoing "github.com/Axit88/UserService/src/domain/userService/core/ports/outgoing/outgoingMt"
	pb "github.com/MindTickle/storageprotos/pb/tickleDbSqlStore"
	"google.golang.org/grpc"
)

type TickleDbSqlImpl struct {
	TickleDbSqlService pb.SqlStoreClient
}

func NewTickleDbSqlClient() outgoing.TickleDbInsertClient {
	connection := fmt.Sprintf("%v:%v", constants.GRPC_HOST, constants.GRPC_PORT)
	conn, err := grpc.Dial(connection, grpc.WithInsecure())
	if err != nil {
		return nil
	}
	res := TickleDbSqlImpl{}
	res.TickleDbSqlService = pb.NewSqlStoreClient(conn)
	return res
}

func (client TickleDbSqlImpl) InsertRow(id string, field model.User, url string, tableName string, reqContext pb.RequestContext, authMeta pb.AuthMeta) error {

	rowValue := pb.RowValue{
		RowInBytes: []byte(fmt.Sprintf(`{"uid": "%s", "uname":"%s"}`, field.UserId, field.Username)),
		AuthMeta:   &authMeta,
	}

	row := pb.Row{
		Id:       id,
		RowValue: &rowValue,
	}
	data := pb.CreateRowsRequest{TableName: tableName, RequestContext: &reqContext}
	data.Rows = append(data.Rows, &row)
	res, err := client.TickleDbSqlService.CreateRows(context.Background(), &data)

	if err != nil {
		return err
	}
	fmt.Println(res)
	return nil
}
