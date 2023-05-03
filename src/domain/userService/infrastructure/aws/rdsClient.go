package infrastructure

import (
	"database/sql"
	"fmt"

	outgoing "github.com/Axit88/UserService/src/domain/userService/core/ports/outgoing/outgoingAws"
	_ "github.com/go-sql-driver/mysql"
)

type OutgoingRds struct {
}

func NewOutgoingRdsClient() outgoing.RdsClient {
	return &OutgoingRds{}
}

func (client OutgoingRds) RdsMySQL(dbName string, tableName string, rdsEndpoint string) error {

	connection := "admin" + ":" + "admin" + "@tcp(" + rdsEndpoint + ")/"
	db, err := sql.Open("mysql", connection)
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + dbName)
	if err != nil {
		return err
	}

	db, err = sql.Open("mysql", connection+dbName)
	if err != nil {
		return err
	}
	defer db.Close()

	query := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
        id INT NOT NULL AUTO_INCREMENT,
        name VARCHAR(255) NOT NULL,
        age INT NOT NULL,
        PRIMARY KEY (id)
    )`, tableName)

	_, err = db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}
