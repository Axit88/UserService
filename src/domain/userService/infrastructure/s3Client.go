package infrastructure

import (
	"bytes"
	"fmt"

	"github.com/Axit88/UserService/src/domain/userService/core/ports/outgoing"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type OutgoingS3 struct {
}

func NewOutgoingS3Client() outgoing.S3Client {
	return &OutgoingS3{}
}

func (client OutgoingS3) PutObjectInS3(bucketname string, region string) error {
	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String(region)},
	)

	svc := s3.New(sess)
	fileContent := []byte("This is a test file")
	res, err := svc.PutObject(&s3.PutObjectInput{
		Body:   bytes.NewReader(fileContent),
		Bucket: aws.String(bucketname),
		Key:    aws.String("test.txt"),
	})
	if err != nil {
		return err
	}
	fmt.Println(res)
	return nil
}
