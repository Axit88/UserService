package main

import (
	"fmt"

	"github.com/Axit88/UserService/src/constants"
	"github.com/Axit88/UserService/src/domain/userService/core/model"
	outgoing "github.com/Axit88/UserService/src/domain/userService/core/ports/outgoing/outgoingMt"
	pb "github.com/MindTickle/platform-protos/pb/event_channel"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type EventChannelImpl struct {
	EventChannelService pb.EventChannelServiceClient
}

func NewEventChannelClient() outgoing.EventChannelClient {
	connection := fmt.Sprintf("%v:%v", constants.GRPC_HOST, constants.GRPC_PORT)
	conn, err := grpc.Dial(connection, grpc.WithInsecure())
	if err != nil {
		return nil
	}
	res := EventChannelImpl{}
	res.EventChannelService = pb.NewEventChannelServiceClient(conn)
	return res
}

func (client EventChannelImpl) CreateEventChannel(ctx context.Context, url string, channelData model.EventChannelField) error {

	newChannel := &pb.EventChannel{
		Name:        channelData.Name,
		Project:     channelData.Project,
		Parallelism: channelData.Parallelism,
		State:       pb.EventChannel_ENABLED,
	}
	eventChannel, err := client.EventChannelService.CreateEventChannel(ctx, &pb.CreateEventChannelRequest{EventChannel: newChannel})
	if err != nil {
		return err
	}

	fmt.Println(eventChannel)
	return nil
}
