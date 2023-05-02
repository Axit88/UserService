package outgoing

type KinesisClient interface {
	PushRecordToKinesis(kinesisStreamName string, data string, region string, partitionKey string) error
}
