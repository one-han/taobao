package request

import (
	"net/url"
)

type TbkDgOptimusMaterialRequest struct {
	values url.Values
}

func (this *TbkDgOptimusMaterialRequest) GetApiMethodName() string {
	return "taobao.tbk.dg.optimus.material"
}
func (this *TbkDgOptimusMaterialRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *TbkDgOptimusMaterialRequest) GetValues() url.Values {
	return this.values
}

/* mm_xxx_xxx_xxx的第三位 */
func (this *TbkDgOptimusMaterialRequest) SetAdzoneId(value string) {
	this.Set("adzone_id", value)
}

/* 内容详情ID */
func (this *TbkDgOptimusMaterialRequest) SetContentId(value string) {
	this.Set("content_id", value)
}

/* 内容渠道信息 */
func (this *TbkDgOptimusMaterialRequest) SetContentSource(value string) {
	this.Set("content_source", value)
}

/* 设备号加密类型：MD5 */
func (this *TbkDgOptimusMaterialRequest) SetDeviceEncrypt(value string) {
	this.Set("device_encrypt", value)
}

/* 设备号类型：IMEI，或者IDFA，或者UTDID */
func (this *TbkDgOptimusMaterialRequest) SetDeviceType(value string) {
	this.Set("device_type", value)
}

/* 设备号加密后的值 */
func (this *TbkDgOptimusMaterialRequest) SetDeviceValue(value string) {
	this.Set("device_value", value)
}

/* 官方的物料Id(详细物料id见：https://tbk.bbs.taobao.com/detail.html?appId=45301&postId=8576096) */
func (this *TbkDgOptimusMaterialRequest) SetMaterialId(value string) {
	this.Set("material_id", value)
}

/* 第几页，默认：1 */
func (this *TbkDgOptimusMaterialRequest) SetPageNo(value string) {
	this.Set("page_no", value)
}

/* 页大小，默认20，1~100 */
func (this *TbkDgOptimusMaterialRequest) SetPageSize(value string) {
	this.Set("page_size", value)
}
