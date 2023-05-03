package main

import (
	"github.com/Axit88/UserService/src/container"
	"github.com/Axit88/UserService/src/utils"
	worker "github.com/Axit88/UserService/src/worker"
)

func main() {
	utils.SetEnv()
	container, err := container.UserServiceContainer()
	if err != nil {
		panic(err)
	}

	err = container.Invoke(worker.StartWorker)
	if err != nil {
		panic(err)
	}
}
