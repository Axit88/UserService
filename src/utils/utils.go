package utils

import (
	"log"
	"os"

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
