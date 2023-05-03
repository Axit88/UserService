package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Axit88/UserService/src/domain/userService/core/model"
)

func VerifySession(url string, sessionId string) (*model.AuthenticatioResponse, error) {
	client := &http.Client{}
	//url := fmt.Sprintf("%v/auth-token/verifySession", baseUrl)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return &model.AuthenticatioResponse{}, nil
	}

	req.Header.Set("x-token", sessionId)
	resp, _ := client.Do(req)
	var data model.AuthenticatioResponse

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		fmt.Println(err)
		return &model.AuthenticatioResponse{}, nil
	}

	return &data, nil
}
