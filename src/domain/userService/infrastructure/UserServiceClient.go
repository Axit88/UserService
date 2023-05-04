package infrastructure

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	pb "github.com/Axit88/UserGrpc/storage-proto"
	"github.com/Axit88/UserService/src/config"
	"github.com/Axit88/UserService/src/domain/userService/core/model"
	"github.com/Axit88/UserService/src/domain/userService/core/ports/outgoing"
	"github.com/MindTickle/mt-go-logger/logger"
	"google.golang.org/grpc"
)

type UserServiceClientImpl struct {
	userService pb.TestApiClient
	logger      *logger.LoggerImpl
}

func NewUserServiceClient(l *logger.LoggerImpl) outgoing.UserServiceClient {
	var cfn, _ = config.NewConfig()
	url := cfn.UserServiceUrl.GrpcUrl
	conn, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		return nil
	}

	res := UserServiceClientImpl{}
	res.userService = pb.NewTestApiClient(conn)
	res.logger = l
	return res
}

func (client UserServiceClientImpl) AddUser(input *model.User) error {
	in := &pb.AddUserInput{
		UserName: input.Username,
		UserId:   input.UserId,
	}

	_, err := client.userService.AddUser(context.Background(), in)
	if err != nil {
		client.logger.Errorf(context.Background(), "Failed To Process GRPC AddUser Request")
	}
	return err
}

func (client UserServiceClientImpl) GetUser(userId string) (*model.User, error) {
	in := &pb.GetUserInput{
		UserId: userId,
	}

	res, err := client.userService.GetUser(context.Background(), in)
	if err != nil {
		client.logger.Errorf(context.Background(), "Failed To Process GRPC GetUser Request")
	}

	output := model.User{
		UserId:   res.UserId,
		Username: res.UserName,
	}
	return &output, err
}

func (client UserServiceClientImpl) DeleteUser(userId string) error {
	in := &pb.DeleteUserInput{
		UserId: userId,
	}

	_, err := client.userService.DeleteUser(context.Background(), in)
	if err != nil {
		client.logger.Errorf(context.Background(), "Failed To Process GRPC DeleteUser Request")
	}
	return err
}

func (client UserServiceClientImpl) UpdateUser(userId string, userName string) error {
	input := pb.UpdateUserInput{
		UserId:   userId,
		UserName: userName,
	}

	_, err := client.userService.UpdateUser(context.Background(), &input)
	if err != nil {
		client.logger.Errorf(context.Background(), "Failed To Process GRPC UpdateUser Request")
	}
	return err
}

func AddUserRest(input *model.User) error {
	inputString := fmt.Sprintf(`{"id": "%s", "name": "%s"}`, input.UserId, input.Username)
	requestBody := strings.NewReader(inputString)
	var cfn, _ = config.NewConfig()
	url := "http://" + cfn.UserServiceUrl.RestUrl + "/User"

	resp, err := http.Post(url, "application/json", requestBody)
	if err != nil {
		logger.Logger.Errorf(context.Background(), "Failed To Process REST AddUser Request")
		return err
	}
	defer resp.Body.Close()

	return nil
}

func GetUserRest(userId string) (*model.User, error) {
	var cfn, _ = config.NewConfig()
	url := "http://" + cfn.UserServiceUrl.RestUrl + "/User/" + string(userId)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	res := model.User{}
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		logger.Logger.Errorf(context.Background(), "Failed To Process REST GetUser Request")
		return nil, err
	}

	return &res, nil
}
