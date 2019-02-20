package main

import (
	"github.com/one-han/taobao/api"
	"github.com/one-han/taobao/api/request"
	"github.com/one-han/taobao/api/response"
	"log"
)

var (
	serverUrl   = "" //为空则默认使用正式环境
	appKey      = "xxx"
	appSecret   = "xxx"
	sessionKey  = "xxx"
	accessToken = "xxx"
)

func main() {
	client := api.GetDefaultTaobaoClient(appKey, appSecret, serverUrl)
	//免签名方式(https)
	// client := api.GetHttpsTaobaoClient(serverUrl)

	var resp response.ItemGetResponse
	var req = new(request.ItemGetRequest)
	req.SetFields("desc")
	req.SetNumIid("21301327956")

	//免签名方式(https)时使用accessToken
	data, err := client.Excute(req, &resp, sessionKey) //"sessionKey or accessToken"
	log.Println("data:", string(data))                 //data为淘宝平台返回的字符串
	if err != nil {
		panic(err)
	}
	if resp.ErrorResponse != nil {
		panic(resp.Msg)
	}
	log.Println("desc:", resp.Item.Desc)
}
