package main

import (
	"fmt"

	"github.com/Axit88/UserService/src/constants"
	"github.com/Axit88/UserService/src/domain/userService/core/model"
	outgoing "github.com/Axit88/UserService/src/domain/userService/core/ports/outgoing/outgoingMt"
	pb "github.com/MindTickle/platform-protos/pb/auditlogsservice"
	"github.com/MindTickle/platform-protos/pb/common"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type AuditImpl struct {
	AuditService pb.AuditLogsServiceClient
}

func NewAuditClient() outgoing.AuditLogClient {
	connection := fmt.Sprintf("%v:%v", constants.GRPC_HOST, constants.GRPC_PORT)
	conn, err := grpc.Dial(connection, grpc.WithInsecure())
	if err != nil {
		return nil
	}
	res := AuditImpl{}
	res.AuditService = pb.NewAuditLogsServiceClient(conn)
	return res
}

func (client AuditImpl) AddLog(myClient pb.AuditLogsServiceClient, url string, reqM common.RequestMeta, schemaField []pb.IngestField, field model.AuditField) error {

	data := pb.AddAuditLogRequest{
		RequestMeta: &reqM,
		Type:        pb.AuditLogType(field.AuditType),
		Timestamp:   field.TimeStamp}

	for i, _ := range schemaField {
		data.Fields = append(data.Fields, &schemaField[i])
	}

	res, err := client.AuditService.AddAuditLog(context.Background(), &data)
	if err != nil {
		return err
	}
	fmt.Println(res)
	return nil
}
