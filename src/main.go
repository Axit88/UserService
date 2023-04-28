package main

import (
	"fmt"

	core "github.com/Axit88/UserService/src/domain/userService/core"
	"github.com/Axit88/UserService/src/domain/userService/core/model"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	client := core.NewUserServiceFacade()
	user := model.User{
		Username: "Jay",
		UserId:   "3",
	}
	err := client.AddUser(&user)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("UserAdded Success")
	}

	res, err := client.GetUser("2")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res.Username)
	}
}
