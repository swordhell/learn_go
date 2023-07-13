package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
)

func currentTimeMillis() int64 {
	return time.Now().UnixNano() / 1000000
}
func main() {
	logrus.Info("start")
	paramters := make(map[string]string)
	paramters["chainId"] = "XRP"
	paramters["hash"] = "0xssfdafdasfdsa"
	paramters["appId"] = ""
	paramters["timestamp"] = strconv.Itoa(int(currentTimeMillis()))

	const NotifyUrl = "http://106.14.252.100:6012/notify_deposit_url_test"

	signature := "0xsfdajfljkleejql"
	paramters["signature"] = signature
	jsonData, err := json.Marshal(paramters)
	if err != nil {
		logrus.Errorln("Error encoding JSON:", err)
		return
	}
	// 创建一个 HTTP 客户端
	client := &http.Client{}

	// 创建一个 HTTP POST 请求
	request, err := http.NewRequest("POST", NotifyUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		logrus.Errorln("Error:", err)
		return
	}

	nonce := fmt.Sprintf("%d", (time.Now().UnixNano() /1000))
	// 设置请求头
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Accept", "application/json")
	request.Header.Add("Biz-Api-Nonce", nonce)

	// 发送请求并获取响应
	response, err := client.Do(request)
	if err != nil {
		logrus.Errorln("Error:", err)
		return
	}
	defer response.Body.Close()

	byteValue, _ := ioutil.ReadAll(response.Body)  
    jsonstr := string(byteValue)  

	logrus.Info(jsonstr)
}
