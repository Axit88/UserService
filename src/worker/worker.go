package lambdaWorkers

import (
	"fmt"

	"github.com/Axit88/UserService/src/domain/userService/application"
)

func StartWorker(userServiceApplication *application.UserServiceApplication) {
	err := userServiceApplication.AddUserApplication("12", "ajay")
	if err != nil {
		fmt.Println(err)
	}

	res, err := userServiceApplication.GetUserApplication("12")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
