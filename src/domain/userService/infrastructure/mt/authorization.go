package main

import (
	"context"
	"fmt"

	"github.com/Axit88/UserService/src/constants"
	outgoing "github.com/Axit88/UserService/src/domain/userService/core/ports/outgoing/outgoingMt"
	pb "github.com/MindTickle/content-protos/pb/authorisation"
	"github.com/MindTickle/content-protos/pb/common"
	"google.golang.org/grpc"
)

type AuthorizationImpl struct {
	AuthorizationService pb.RolePermissionServiceClient
}

func NewAuthorizationClient() outgoing.AuthorizationClient {
	connection := fmt.Sprintf("%v:%v", constants.GRPC_HOST, constants.GRPC_PORT)
	conn, err := grpc.Dial(connection, grpc.WithInsecure())
	if err != nil {
		return nil
	}
	res := AuthorizationImpl{}
	res.AuthorizationService = pb.NewRolePermissionServiceClient(conn)
	return res
}

func (client AuthorizationImpl) GetCompnanyRolePermission(url string, companyId string, reqMeta common.RequestMeta, recMeta common.RecordMeta) error {

	data := pb.GetRoleAndPermissionsRequest{RequestMeta: &reqMeta, RecordMeta: &recMeta}
	data.CompanyIds = append(data.CompanyIds, companyId)

	roles, err := client.AuthorizationService.GetRolesAndPermissions(context.Background(), &data)
	if err != nil {
		return err
	}
	fmt.Println(roles)
	return nil
}

func (client AuthorizationImpl) GetUserRolePermission(url string, userId string, reqMeta common.RequestMeta, recMeta common.RecordMeta) error {

	data := pb.GetUserRolesAndPermissionsRequest{UserId: userId, RequestMeta: &reqMeta, RecordMeta: &recMeta}

	roles, err := client.AuthorizationService.GetUserRolesAndPermissions(context.Background(), &data)
	fmt.Println(roles)
	if err != nil {
		return err
	}
	return nil
}
