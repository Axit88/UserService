package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Axit88/UserService/src/domain/userService/core/model"
)

func SendNotification(url string, input model.NotificationField) (*model.EmailResponse, error) {
	res := model.EmailResponse{
		JobId: "",
	}
	c := http.Client{Timeout: time.Duration(1) * time.Second}
	data := map[string]interface{}{
		"params": map[string]interface{}{
			"from_name":   input.From_Name,
			"domain_base": input.DomainBase,
			"cname":       input.Cname,
			"series":      input.Series,
			"entity":      input.Entity,
			"user_type":   input.UserType,
		},
		"reply_to":            input.ReplyTo,
		"from":                input.From,
		"mailer":              "ses",
		"category":            input.Category,
		"template_evaluation": "strict",
		"html":                true,
		"notificationChannel": input.NotificationChannel,
		"to":                  input.To,
		"template":            input.Template,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("Error %s", err)
		return nil, err
	}

	url = fmt.Sprintf("%v/immediate", url)

	resp, err := c.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("Error %s", err)
		return nil, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return nil, err
	}
	return &res, err
}
