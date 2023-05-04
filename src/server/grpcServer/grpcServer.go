package main

import (
	"context"
	"fmt"
	"net"

	pb "github.com/Axit88/UserGrpc/storage-proto"
	"github.com/Axit88/UserService/src/config"
	"github.com/Axit88/UserService/src/utils"
	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
)

type testApiServer struct {
	pb.UnimplementedTestApiServer
}

func (s *testApiServer) AddUser(ctx context.Context, input *pb.AddUserInput) (*pb.AddUserOutput, error) {
	db, err := utils.CreateDbClient()
	if err != nil {
		return nil, err
	}

	_, err = db.Exec("INSERT INTO USER (UserId, UserName) VALUES (?, ?)", input.UserId, input.UserName)
	if err != nil {
		return nil, err
	}

	output := pb.AddUserOutput{
		Message: fmt.Sprintf("User %v Added Successfully", input.UserName),
	}

	return &output, nil
}

func (s *testApiServer) GetUser(ctx context.Context, input *pb.GetUserInput) (*pb.GetUserOutput, error) {
	db, err := utils.CreateDbClient()
	if err != nil {
		return nil, err
	}

	queryResult, err := db.Query("SELECT * FROM USER WHERE UserId = ?", input.UserId)
	if err != nil {
		return nil, err
	}

	output := pb.GetUserOutput{}
	for queryResult.Next() {
		err = queryResult.Scan(&output.UserId, &output.UserName)
		if err != nil {
			return nil, err
		}
	}
	return &output, nil
}

func (s *testApiServer) DeleteUser(ctx context.Context, input *pb.DeleteUserInput) (*pb.DeleteUserOutput, error) {
	db, err := utils.CreateDbClient()
	if err != nil {
		return nil, err
	}

	_, err = db.Query("DELETE FROM USER WHERE UserId = ?", input.UserId)
	if err != nil {
		return nil, err
	}

	output := pb.DeleteUserOutput{}
	output.Message = fmt.Sprintf("User with id %v Deleted", input.UserId)
	return &output, nil
}

func (s *testApiServer) UpdateUser(ctx context.Context, input *pb.UpdateUserInput) (*pb.UpdateUserOutput, error) {
	db, err := utils.CreateDbClient()
	if err != nil {
		return nil, err
	}

	_, err = db.Exec("UPDATE USER SET UserName = ? WHERE UserId = ?", input.UserName, input.UserId)
	if err != nil {
		return nil, err
	}

	output := pb.UpdateUserOutput{
		Message: fmt.Sprintf("User %v Updated Name %v", input.UserId, input.UserName),
	}

	return &output, nil
}

func RunGrpcServer() error {
	var cfn, _ = config.NewConfig()
	url := cfn.UserServiceUrl.GrpcUrl
	listner, err := net.Listen("tcp", url)
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()
	pb.RegisterTestApiServer(grpcServer, &testApiServer{})
	err = grpcServer.Serve(listner)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	utils.SetEnv()
	err := RunGrpcServer()
	if err != nil {
		fmt.Println(err)
	}
}
