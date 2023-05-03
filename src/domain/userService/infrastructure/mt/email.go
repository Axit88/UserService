package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Axit88/UserService/src/domain/userService/core/model"
)

func SendEmail(url string, input model.EmailField) (*model.EmailResponse, error) {
	res := model.EmailResponse{
		JobId: "",
	}

	data := map[string]interface{}{
		"mailer":              "ses",
		"from":                input.From,
		"to":                  input.To,
		"reply_to":            input.ReplyTo,
		"template":            input.Template,
		"template_evaluation": "strict",
		"params": map[string]interface{}{
			"cname":           input.Cname,
			"from_name":       input.FromName,
			"fromDate":        "24 Dec",
			"toDate":          "31 Dec",
			"digestFrequency": "Weekly",
			"digest": map[string]interface{}{
				"sampleCategory": map[string]interface{}{
					"totalCount": 5,
					"digestUrl":  "s3:url-for-user1",
				},
			},
		},
		"html": "true",
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("Error %s", err)
		return nil, err
	}

	c := http.Client{Timeout: time.Duration(1) * time.Second}
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

	return &res, nil
}
