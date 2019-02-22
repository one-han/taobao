package request

import (
	"net/url"
)

type TbkTpwdCreateRequest struct {
	values url.Values
}

func (this *TbkTpwdCreateRequest) GetApiMethodName() string {
	return "taobao.tbk.tpwd.create"
}
func (this *TbkTpwdCreateRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *TbkTpwdCreateRequest) GetValues() url.Values {
	return this.values
}

/* 扩展字段JSON格式 */
func (this *TbkTpwdCreateRequest) SetExt(value string) {
	this.Set("ext", value)
}

/* 口令弹框logoURL */
func (this *TbkTpwdCreateRequest) SetLogo(value string) {
	this.Set("logo", value)
}

/* 口令弹框内容 */
func (this *TbkTpwdCreateRequest) SetText(value string) {
	this.Set("text", value)
}

/* 口令跳转目标页 */
func (this *TbkTpwdCreateRequest) SetUrl(value string) {
	this.Set("url", value)
}

/* 生成口令的淘宝用户ID */
func (this *TbkTpwdCreateRequest) SetUserId(value string) {
	this.Set("user_id", value)
}
