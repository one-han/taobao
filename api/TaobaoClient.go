package api

import (
	"crypto/md5"
	"crypto/tls"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
	"time"
)

const (
	APP_KEY     = "app_key"
	FORMAT      = "format"
	METHOD      = "method"
	TIMESTAMP   = "timestamp"
	VERSION     = "v"
	SIGN        = "sign"
	SIGN_METHOD = "sign_method"
	PARTNER_ID  = "partner_id"
	SESSION     = "session"
	SIMPLIFY    = "simplify"
)

type TaobaoClient interface {
	Excute(request TaobaoRequest, response interface{}, sessionKey string) ([]byte, error)
}

type DefaultClient struct {
	ServerUrl string
	Format    string
	Version   string
}
type DefaultTaobaoClient struct {
	DefaultClient
	SignMethod string
	AppKey     string
	AppSecret  string
}
type HttpsTaobaoClient struct { //https://eco.taobao.com/router/rest
	DefaultClient
}

func (c *DefaultTaobaoClient) Excute(request TaobaoRequest, response interface{}, sessionKey string) ([]byte, error) {
	request.Set(APP_KEY, c.AppKey)
	request.Set(METHOD, request.GetApiMethodName())
	request.Set(FORMAT, c.Format)
	request.Set(VERSION, c.Version)
	//request.Set(SESSION, sessionKey)
	request.Set(SIGN_METHOD, c.SignMethod)
	request.Set(TIMESTAMP, time.Now().Format("2006-01-02 15:04:05"))
	request.Set(SIGN, md5Signature(c.AppSecret, request))
	body := strings.NewReader(request.GetValues().Encode())
	req, err := http.NewRequest("POST", c.ServerUrl, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded") //必须

	tr := &http.Transport{
		TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
		DisableCompression: false,
		DisableKeepAlives:  false,
	}
	client := &http.Client{Transport: tr}
	return clientDo(client, req, response)
}

func (c *HttpsTaobaoClient) Excute(request TaobaoRequest, response interface{}, accessToken string) ([]byte, error) {
	request.Set(METHOD, request.GetApiMethodName())
	request.Set(FORMAT, c.Format)
	request.Set(VERSION, c.Version)
	request.Set("access_token", accessToken)

	body := strings.NewReader(request.GetValues().Encode())
	req, err := http.NewRequest("POST", c.ServerUrl, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded") //必须

	transport := new(http.Transport)
	transport.TLSClientConfig = &tls.Config{InsecureSkipVerify: false}

	client := new(http.Client)
	client.Transport = transport

	return clientDo(client, req, response)
}

func clientDo(client *http.Client, req *http.Request, response interface{}) ([]byte, error) {
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if response != nil {
		return data, json.Unmarshal(data, &response)
	}
	return data, nil
}

func md5Signature(secret string, request TaobaoRequest) string {
	keys := make([]string, 0)
	for k := range request.GetValues() {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	str := secret
	for _, k := range keys {
		v := request.GetValues().Get(k)
		if k != "" && v != "" {
			str += k + v
		}
	}
	str += secret
	h := md5.New()
	h.Write([]byte(str))
	return strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
}

func GetDefaultTaobaoClient(appkey, appSecret, serverUrl string) *DefaultTaobaoClient {
	if serverUrl == "" {
		serverUrl = "http://gw.api.taobao.com/router/rest"
	}
	return &DefaultTaobaoClient{
		DefaultClient: DefaultClient{
			serverUrl,
			"json",
			"2.0",
		},
		SignMethod: "md5",
		AppKey:     appkey,
		AppSecret:  appSecret,
	}
}
func GetHttpsTaobaoClient(serverUrl string) *HttpsTaobaoClient {
	if serverUrl == "" {
		serverUrl = "https://eco.taobao.com/router/rest"
	}
	return &HttpsTaobaoClient{
		DefaultClient: DefaultClient{
			serverUrl,
			"json",
			"2.0",
		},
	}
}
