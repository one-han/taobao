package request

import (
	"net/url"
)

type TbkDgMaterialOptionalRequest struct {
	values url.Values
}

func (this *TbkDgMaterialOptionalRequest) GetApiMethodName() string {
	return "taobao.tbk.dg.material.optional"
}
func (this *TbkDgMaterialOptionalRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *TbkDgMaterialOptionalRequest) GetValues() url.Values {
	return this.values
}

/* mm_xxx_xxx_xxx的第三位 */
func (this *TbkDgMaterialOptionalRequest) SetAdzoneId(value string) {
	this.Set("adzone_id", value)
}

/* 后台类目ID，用,分割，最大10个，该ID可以通过taobao.itemcats.get接口获取到 */
func (this *TbkDgMaterialOptionalRequest) SetCat(value string) {
	this.Set("cat", value)
}

/* 设备号加密类型：MD5 */
func (this *TbkDgMaterialOptionalRequest) SetDeviceEncrypt(value string) {
	this.Set("device_encrypt", value)
}

/* 设备号类型：IMEI，或者IDFA，或者UTDID */
func (this *TbkDgMaterialOptionalRequest) SetDeviceType(value string) {
	this.Set("device_type", value)
}

/* 设备号加密后的值 */
func (this *TbkDgMaterialOptionalRequest) SetDeviceValue(value string) {
	this.Set("device_value", value)
}

/* KA媒体淘客佣金比率上限，如：1234表示12.34% */
func (this *TbkDgMaterialOptionalRequest) SetEndKaTkRate(value string) {
	this.Set("end_ka_tk_rate", value)
}

/* 折扣价范围上限，单位：元 */
func (this *TbkDgMaterialOptionalRequest) SetEndPrice(value string) {
	this.Set("end_price", value)
}

/* 淘客佣金比率上限，如：1234表示12.34% */
func (this *TbkDgMaterialOptionalRequest) SetEndTkRate(value string) {
	this.Set("end_tk_rate", value)
}

/* 是否有优惠券，设置为true表示该商品有优惠券，设置为false或不设置表示不判断这个属性 */
func (this *TbkDgMaterialOptionalRequest) SetHasCoupon(value string) {
	this.Set("has_coupon", value)
}

/* 好评率是否高于行业均值 */
func (this *TbkDgMaterialOptionalRequest) SetIncludeGoodRate(value string) {
	this.Set("include_good_rate", value)
}

/* 成交转化是否高于行业均值 */
func (this *TbkDgMaterialOptionalRequest) SetIncludePayRate30(value string) {
	this.Set("include_pay_rate_30", value)
}

/* 退款率是否低于行业均值 */
func (this *TbkDgMaterialOptionalRequest) SetIncludeRfdRate(value string) {
	this.Set("include_rfd_rate", value)
}

/* ip参数影响邮费获取，如果不传或者传入不准确，邮费无法精准提供 */
func (this *TbkDgMaterialOptionalRequest) SetIp(value string) {
	this.Set("ip", value)
}

/* 是否海外商品，设置为true表示该商品是属于海外商品，设置为false或不设置表示不判断这个属性 */
func (this *TbkDgMaterialOptionalRequest) SetIsOverseas(value string) {
	this.Set("is_overseas", value)
}

/* 是否商城商品，设置为true表示该商品是属于淘宝商城商品，设置为false或不设置表示不判断这个属性 */
func (this *TbkDgMaterialOptionalRequest) SetIsTmall(value string) {
	this.Set("is_tmall", value)
}

/* 所在地 */
func (this *TbkDgMaterialOptionalRequest) SetItemloc(value string) {
	this.Set("itemloc", value)
}

/* 官方的物料Id(详细物料id见：https://tbk.bbs.taobao.com/detail.html?appId=45301&postId=8576096)，不传时默认为2836 */
func (this *TbkDgMaterialOptionalRequest) SetMaterialId(value string) {
	this.Set("material_id", value)
}

/* 是否包邮，true表示包邮，空或false表示不限 */
func (this *TbkDgMaterialOptionalRequest) SetNeedFreeShipment(value string) {
	this.Set("need_free_shipment", value)
}

/* 是否加入消费者保障，true表示加入，空或false表示不限 */
func (this *TbkDgMaterialOptionalRequest) SetNeedPrepay(value string) {
	this.Set("need_prepay", value)
}

/* 牛皮癣程度，取值：1:不限，2:无，3:轻微 */
func (this *TbkDgMaterialOptionalRequest) SetNpxLevel(value string) {
	this.Set("npx_level", value)
}

/* 第几页，默认：１ */
func (this *TbkDgMaterialOptionalRequest) SetPageNo(value string) {
	this.Set("page_no", value)
}

/* 页大小，默认20，1~100 */
func (this *TbkDgMaterialOptionalRequest) SetPageSize(value string) {
	this.Set("page_size", value)
}

/* 链接形式：1：PC，2：无线，默认：１ */
func (this *TbkDgMaterialOptionalRequest) SetPlatform(value string) {
	this.Set("platform", value)
}

/* 查询词 */
func (this *TbkDgMaterialOptionalRequest) SetQ(value string) {
	this.Set("q", value)
}

/* 排序_des（降序），排序_asc（升序），销量（total_sales），淘客佣金比率（tk_rate）， 累计推广量（tk_total_sales），总支出佣金（tk_total_commi），价格（price） */
func (this *TbkDgMaterialOptionalRequest) SetSort(value string) {
	this.Set("sort", value)
}

/* 店铺dsr评分，筛选高于等于当前设置的店铺dsr评分的商品0-50000之间 */
func (this *TbkDgMaterialOptionalRequest) SetStartDsr(value string) {
	this.Set("start_dsr", value)
}

/* KA媒体淘客佣金比率下限，如：1234表示12.34% */
func (this *TbkDgMaterialOptionalRequest) SetStartKaTkRate(value string) {
	this.Set("start_ka_tk_rate", value)
}

/* 折扣价范围下限，单位：元 */
func (this *TbkDgMaterialOptionalRequest) SetStartPrice(value string) {
	this.Set("start_price", value)
}

/* 淘客佣金比率下限，如：1234表示12.34% */
func (this *TbkDgMaterialOptionalRequest) SetStartTkRate(value string) {
	this.Set("start_tk_rate", value)
}
