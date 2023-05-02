package outgoing

type RdsClient interface {
	RdsMySQL(databaseName string, tableName string, rdsEndpoint string) error
}
