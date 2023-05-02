package main

import (
	"github.com/Axit88/UserService/src/container"
	worker "github.com/Axit88/UserService/src/worker"
)

func main() {
	container, err := container.UserServiceContainer()
	if err != nil {
		panic(err)
	}

	err = container.Invoke(worker.StartWorker)
	if err != nil {
		panic(err)
	}
}
