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
	"google.golang.org/grpc"
)

type UserServiceClientImpl struct {
	userService pb.TestApiClient
}

func NewUserServiceClient() outgoing.UserServiceClient {
	var cfn, _ = config.NewConfig()
	connection := cfn.UserServiceUrl.GrpcUrl
	conn, err := grpc.Dial(connection, grpc.WithInsecure())
	if err != nil {
		return nil
	}

	res := UserServiceClientImpl{}
	res.userService = pb.NewTestApiClient(conn)
	return res
}

func (client UserServiceClientImpl) AddUser(input *model.User) error {
	in := &pb.AddUserInput{
		UserName: input.Username,
		UserId:   input.UserId,
	}

	_, err := client.userService.AddUser(context.Background(), in)
	return err
}

func (client UserServiceClientImpl) GetUser(userId string) (*model.User, error) {
	in := &pb.GetUserInput{
		UserId: userId,
	}

	res, err := client.userService.GetUser(context.Background(), in)
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
	return err
}

func (client UserServiceClientImpl) UpdateUser(userId string, userName string) error {
	input := pb.UpdateUserInput{
		UserId:   userId,
		UserName: userName,
	}

	_, err := client.userService.UpdateUser(context.Background(), &input)
	return err
}

func AddUserRest(input *model.User) error {
	inputString := fmt.Sprintf(`{"id": "%s", "name": "%s"}`, input.UserId, input.Username)
	requestBody := strings.NewReader(inputString)
	var cfn, _ = config.NewConfig()
	url := cfn.UserServiceUrl.RestUrl + "/User"

	resp, err := http.Post(url, "application/json", requestBody)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

func GetUserRest(userId string) (*model.User, error) {
	var cfn, _ = config.NewConfig()
	url := cfn.UserServiceUrl.RestUrl + "/User/" + string(userId)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	res := model.User{}
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
