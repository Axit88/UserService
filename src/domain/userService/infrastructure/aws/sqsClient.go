package infrastructure

import (
	"fmt"

	outgoing "github.com/Axit88/UserService/src/domain/userService/core/ports/outgoing/outgoingAws"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type OutgoingSqs struct {
}

func NewOutgoingSqsClient() outgoing.SqsClient {
	return &OutgoingSqs{}
}

func (client OutgoingSqs) SendMessageToSqsQueue(queueUrl string, messageBody string, region string) error {

	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(region),
	}))

	svc := sqs.New(sess)

	input := &sqs.SendMessageInput{
		MessageBody: aws.String(messageBody),
		QueueUrl:    aws.String(queueUrl),
	}

	result, err := svc.SendMessage(input)
	if err != nil {
		return err
	}

	fmt.Println(result)
	return nil
}
