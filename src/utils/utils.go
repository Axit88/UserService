package utils

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/Axit88/UserService/src/config"
	"github.com/Axit88/UserService/src/constants"
	"github.com/joho/godotenv"
)

func SetEnv() {
	if os.Getenv("LAMBDA_TASK_ROOT") == "" && os.Getenv("AWS_EXECUTION_ENV") == "" {
		//pwd, _ := os.Getwd()
		// err := godotenv.Load(pwd + "/env/local.env")
		err := godotenv.Load("/Users/axit/Desktop/UserApi/src/env/local.env")
		if err != nil {
			log.Fatalf("Some error occured. Err: %s", err)
		}
	}
}

func CreateDbClient() (sql.DB, error) {
	var cfn, _ = config.NewConfig()
	connection := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", cfn.DbConfig.UserName, cfn.DbConfig.Password, cfn.DbConfig.Host, cfn.DbConfig.Port, constants.DB_NAME)
	db, err := sql.Open("mysql", connection)
	if err != nil {
		return sql.DB{}, err
	}
	defer db.Close()
	return *db, nil
}
