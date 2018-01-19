package main

import (
	"net/http"
	"bytes"
	"encoding/json"
	"yhhttp"
)

func main() {
	yhhttp.HttpRequest(SendMsgRequest())
}

func SendMsgRequest() func() (*http.Request, error) {
	return func() (*http.Request, error) {
		url := "https://api.bmob.cn/1/requestSms"
		msg := make(map[string]interface{})

		msg["mobilePhoneNumber"]="15680571195"
		msg["content"]="content"
		msg["sendTime"]="15680571195"
		data,_ := json.Marshal(msg)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
		if err != nil {
			return nil,err
		}
		req.Header.Set("Content-type", "application/json;charset=UTF-8")
		req.Header.Set("X-Bmob-Application-Id", "189217726a2cc9653bdfa9f5e6453eea")
		req.Header.Set("X-Bmob-REST-API-Key", "0fd330f331db278bc594ec06a631633d")
		return req, nil
	}
}