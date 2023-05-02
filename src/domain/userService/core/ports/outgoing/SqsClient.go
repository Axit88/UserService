package outgoing

type SqsClient interface {
	SendMessageToSqsQueue(queueUrl string, messageBody string, region string) error
}
