package outgoing

type DynamoDbClient interface {
	PushItemToDynamoDb(tableName string, region string, id string, name string) error 
}
