package api

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
	"time"
)

type TaobaoClient interface {
	Excute(request TaobaoRequest, response interface{}, sessionKey string) ([]byte, error)
}

type DefaultTaobaoClient struct {
	ServerUrl string
	AppKey    string
	AppSecret string
}

func (c *DefaultTaobaoClient) Excute(request TaobaoRequest, response interface{}, sessionKey string) ([]byte, error) {
	request.SetValue("app_key", c.AppKey)
	request.SetValue("format", "json")
	request.SetValue("method", request.GetApiMethodName())
	// request.SetValue("sign_method", "md5")
	request.SetValue("v", "2.0")
	request.SetValue("session", sessionKey)
	request.SetValue("timestamp", time.Now().Format("2006-01-02 15:04:05"))
	request.SetValue("sign", md5Signature(c.AppSecret, request))

	body := strings.NewReader(request.GetValues().Encode())
	req, err := http.NewRequest("POST", c.ServerUrl, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded") //必须

	client := &http.Client{}
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
	h := md5.New()
	h.Write([]byte(str))
	return strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
}
