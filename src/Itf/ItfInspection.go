package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"verifyItf"
	"time"
	"yhhttp"
)

func main() {
	for {

		currentPath := verifyitf.GetCurrentDirectory()
		fmt.Println("path = ", currentPath)
		rbc, err := verifyitf.ReadBaseConfig(currentPath + verifyitf.BaseConfigName)
		if err != nil {
			fmt.Println("sorroy ,need config,Please contact yh ", err)
			return
		}

		for _, itfName := range rbc.Order {
			fmt.Println("pa = ", itfName)
			//获取接口的参数
			itf, err := verifyitf.ReadItfParam(rbc.Path + itfName)

			if err != nil {
				fmt.Println("read ReadItfParam config failed = ", err)
			} else {
				re, _ := yhhttp.HttpRequest(verifyitf.GenerateRequest(rbc, itf))
				if rbc.RefreshToken == itfName { //刷新token的接口和读取的参数文件名一致时，刷新token
					//当时登录接口时，将token值赋给bc的tokenValue属性上。
					str := fmt.Sprintf("%v", re.Data["token"])
					rbc.Headers[rbc.TokenName] = string(str)
					fmt.Println("token", string(str))
				} else {
					if re.Code != 200 {
						if re.Code == 401 {
							yhhttp.HttpRequest(Notify("Token异常了"))
						} else if re.Code == 4014 {
							yhhttp.HttpRequest(Notify("刷新接口也异常了。"))
						} else {
							yhhttp.HttpRequest(Notify(itfName+"接口出现了问题。"))
						}
					}
				}
			}
		}

		time.Sleep(rbc.Time * time.Minute)
	}
}

func Notify(message string) func() (*http.Request, error) {
	return func() (*http.Request, error) {
		url := "https://oapi.dingtalk.com/robot/send?access_token=454a9c24fa1f6f871e7e3d0bb8a9748a822df3eb7da033373a514df3dc16c2df"
		msg := make(map[string]interface{})

		msg["msgtype"] = "text"

		content := make(map[string]string)
		content["content"] = message //"d8zone生成微信预支付订单接口出问题了"
		msg["text"] = content

		at := make(map[string]interface{})
		at["atMobiles"] = "[15680571195]"
		at["isAtAll"] = false
		msg["at"] = &at

		data, err := json.Marshal(msg)
		if err != nil {
			return nil, err
		}
		fmt.Println("data = ", string(data))

		req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
		req.Header.Set("Content-type", "application/json;charset=UTF-8")
		return req, nil
	}
}
