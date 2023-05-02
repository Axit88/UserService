package outgoing

type S3Client interface {
	PutObjectInS3(bucketname string, region string) error 
}
