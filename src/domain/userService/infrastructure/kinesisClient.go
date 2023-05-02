package infrastructure

import (
	"fmt"

	"github.com/Axit88/UserService/src/domain/userService/core/ports/outgoing"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kinesis"
)

type OutgoingKinesis struct {
}

func NewOutgoingKinesisClient() outgoing.KinesisClient {
	return &OutgoingKinesis{}
}

func (client OutgoingKinesis) PushRecordToKinesis(kinesisStreamName string, data string, region string, partitionKey string) error {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(region),
	}))

	svc := kinesis.New(sess)
	res, err := svc.PutRecord(&kinesis.PutRecordInput{
		Data:         []byte(data),
		StreamName:   aws.String(kinesisStreamName),
		PartitionKey: aws.String(partitionKey),
	})

	if err != nil {
		return err
	}
	fmt.Println(res)

	return nil
}
