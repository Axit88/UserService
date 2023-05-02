package infrastructure

import (
	"fmt"

	"github.com/Axit88/UserService/src/domain/userService/core/ports/outgoing"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type OutgoingDynamoDb struct {
}

func NewOutgoingDynamoDbClient() outgoing.DynamoDbClient {
	return &OutgoingDynamoDb{}
}

func (client OutgoingDynamoDb) PushItemToDynamoDb(tableName string, region string, id string, name string) error {
	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String(region)},
	)

	svc := dynamodb.New(sess)
	input := dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item: map[string]*dynamodb.AttributeValue{
			"id":   {S: aws.String(id)},
			"name": {S: aws.String(name)},
		},
	}
	res, err := svc.PutItem(&input)
	if err != nil {
		return err
	}
	fmt.Println(res)
	return nil
}