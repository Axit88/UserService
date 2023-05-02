package lambdaWorkers

import (
	"fmt"

	"github.com/Axit88/UserService/src/domain/userService/application"
)

func StartWorker(userServiceApplication *application.UserServiceApplication) {
	err := userServiceApplication.AddUserApplication("11", "pk")
	if err != nil {
		fmt.Println(err)
	}

	res, err := userServiceApplication.GetUserApplication("11")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
