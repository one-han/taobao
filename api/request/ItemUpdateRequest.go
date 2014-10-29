package request

import (
	"net/url"
)

type ItemUpdateRequest struct {
	values url.Values
}

func (r *ItemUpdateRequest) GetApiMethodName() string {
	return "taobao.item.update"
}
func (r *ItemUpdateRequest) SetValue(key, value string) {
	if r.values == nil {
		r.values = url.Values{}
	}
	r.values.Set(key, value)
}
func (r *ItemUpdateRequest) GetValues() url.Values {
	return r.values
}

/* 售后服务说明模板id */
func (r *ItemUpdateRequest) SetAfterSaleId(value string) {
	r.SetValue("after_sale_id", value)
}

/* 商品上传后的状态。可选值:onsale（出售中）,instock（库中），如果同时更新商品状态为出售中及list_time为将来的时间，则商品还是处于定时上架的状态, 此时item.is_timing为true */
func (r *ItemUpdateRequest) SetApproveStatus(value string) {
	r.SetValue("approve_status", value)
}

/* 商品的积分返点比例。如：5 表示返点比例0.5%. 注意：返点比例必须是>0的整数，而且最大是90,即为9%.B商家在发布非虚拟商品时，返点必须是 5的倍数，即0.5%的倍数。其它是1的倍数，即0.1%的倍数。无名良品商家发布商品时，复用该字段记录积分宝返点比例，返点必须是对应类目的返点步长的整数倍，默认是5，即0.5%。注意此时该字段值依旧必须是>0的整数，注意此时该字段值依旧必须是>0的整数，最高值不超过500，即50% */
func (r *ItemUpdateRequest) SetAuctionPoint(value string) {
	r.SetValue("auction_point", value)
}

/* 代充商品类型。只有少数类目下的商品可以标记上此字段，具体哪些类目可以上传可以通过taobao.itemcat.features.get获得。在代充商品的类目下，不传表示不标记商品类型（交易搜索中就不能通过标记搜到相关的交易了）。可选类型：
no_mark(不做类型标记)
time_card(点卡软件代充)
fee_card(话费软件代充) */
func (r *ItemUpdateRequest) SetAutoFill(value string) {
	r.SetValue("auto_fill", value)
}

/* 商品条形码 */
func (r *ItemUpdateRequest) SetBarcode(value string) {
	r.SetValue("barcode", value)
}

/* 商品基础色，数据格式为：pid:vid:rvid1,rvid2,rvid3;pid:vid:rvid1 */
func (r *ItemUpdateRequest) SetChangeProp(value string) {
	r.SetValue("change_prop", value)
}

/* 天猫超市扩展字段，天猫超市专用 */
func (r *ItemUpdateRequest) SetChaoshiExtendsInfo(value string) {
	r.SetValue("chaoshi_extends_info", value)
}

/* 叶子类目id<br /> 支持最小值为：0 */
func (r *ItemUpdateRequest) SetCid(value string) {
	r.SetValue("cid", value)
}

/* 货到付款运费模板ID
该字段已经废弃，货到付款模板已经集成到运费模板中。 */
func (r *ItemUpdateRequest) SetCodPostageId(value string) {
	r.SetValue("cod_postage_id", value)
}

/* 商品描述. 字数要大于5个字符，小于25000个字符 ，受违禁词控制<br /> 支持最大长度为：200000<br /> 支持的最大列表长度为：200000 */
func (r *ItemUpdateRequest) SetDesc(value string) {
	r.SetValue("desc", value)
}

/* 商品描述模块化，模块列表；数据结构可参考Item_Desc_Module 。详细使用说明：http://open.taobao.com/support/question_detail.htm?spm=0.0.0.0.cRcj3S&id=147498 ； */
func (r *ItemUpdateRequest) SetDescModules(value string) {
	r.SetValue("desc_modules", value)
}

/* 支持宝贝信息的删除,如需删除对应的食品安全信息中的储藏方法、保质期， 则应该设置此参数的值为：food_security.plan_storage,food_security.period; 各个参数的名称之间用【,】分割, 如果对应的参数有设置过值，即使在这个列表中，也不会被删除; 目前支持此功能的宝贝信息如下：食品安全信息所有字段、电子交易凭证字段（locality_life，locality_life.verification，locality_life.refund_ratio，locality_life.network_id ，locality_life.onsale_auto_refund_ratio）。支持对全球购宝贝信息的清除（字符串中包含global_stock） */
func (r *ItemUpdateRequest) SetEmptyFields(value string) {
	r.SetValue("empty_fields", value)
}

/* ems费用。取值范围:0.01-999.00;精确到2位小数;单位:元。如:25.07，表示:25元7分 */
func (r *ItemUpdateRequest) SetEmsFee(value string) {
	r.SetValue("ems_fee", value)
}

/* 快递费用。取值范围:0.01-999.00;精确到2位小数;单位:元。如:15.07，表示:15元7分 */
func (r *ItemUpdateRequest) SetExpressFee(value string) {
	r.SetValue("express_fee", value)
}

/* 厂家联系方式 */
func (r *ItemUpdateRequest) SetFoodSecurityContact(value string) {
	r.SetValue("food_security.contact", value)
}

/* 产品标准号 */
func (r *ItemUpdateRequest) SetFoodSecurityDesignCode(value string) {
	r.SetValue("food_security.design_code", value)
}

/* 厂名 */
func (r *ItemUpdateRequest) SetFoodSecurityFactory(value string) {
	r.SetValue("food_security.factory", value)
}

/* 厂址 */
func (r *ItemUpdateRequest) SetFoodSecurityFactorySite(value string) {
	r.SetValue("food_security.factory_site", value)
}

/* 食品添加剂 */
func (r *ItemUpdateRequest) SetFoodSecurityFoodAdditive(value string) {
	r.SetValue("food_security.food_additive", value)
}

/* 健字号，保健品/膳食营养补充剂 这个类目下特有的信息，此类目下无需填写生产许可证编号（QS），如果填写了生产许可证编号（QS）将被忽略不保存；保存宝贝时，标题前会自动加上健字号产品名称一起作为宝贝标题； */
func (r *ItemUpdateRequest) SetFoodSecurityHealthProductNo(value string) {
	r.SetValue("food_security.health_product_no", value)
}

/* 配料表 */
func (r *ItemUpdateRequest) SetFoodSecurityMix(value string) {
	r.SetValue("food_security.mix", value)
}

/* 保质期 */
func (r *ItemUpdateRequest) SetFoodSecurityPeriod(value string) {
	r.SetValue("food_security.period", value)
}

/* 储藏方法 */
func (r *ItemUpdateRequest) SetFoodSecurityPlanStorage(value string) {
	r.SetValue("food_security.plan_storage", value)
}

/* 生产许可证号 */
func (r *ItemUpdateRequest) SetFoodSecurityPrdLicenseNo(value string) {
	r.SetValue("food_security.prd_license_no", value)
}

/* 生产结束日期,格式必须为yyyy-MM-dd */
func (r *ItemUpdateRequest) SetFoodSecurityProductDateEnd(value string) {
	r.SetValue("food_security.product_date_end", value)
}

/* 生产开始日期，格式必须为yyyy-MM-dd */
func (r *ItemUpdateRequest) SetFoodSecurityProductDateStart(value string) {
	r.SetValue("food_security.product_date_start", value)
}

/* 进货结束日期，要在生产日期之后，格式必须为yyyy-MM-dd */
func (r *ItemUpdateRequest) SetFoodSecurityStockDateEnd(value string) {
	r.SetValue("food_security.stock_date_end", value)
}

/* 进货开始日期，要在生产日期之后，格式必须为yyyy-MM-dd */
func (r *ItemUpdateRequest) SetFoodSecurityStockDateStart(value string) {
	r.SetValue("food_security.stock_date_start", value)
}

/* 供货商 */
func (r *ItemUpdateRequest) SetFoodSecuritySupplier(value string) {
	r.SetValue("food_security.supplier", value)
}

/* 运费承担方式。运费承担方式。可选值:seller（卖家承担）,buyer(买家承担); */
func (r *ItemUpdateRequest) SetFreightPayer(value string) {
	r.SetValue("freight_payer", value)
}

/* 全球购商品采购地（地区/国家）,默认值只在全球购商品采购地（库存类型选择情况生效），地区国家值为（美国, 香港, 日本, 英国, 新西兰, 德国, 韩国, 荷兰, 澳洲, 法国, 意大利, 台湾, 澳门, 加拿大, 瑞士, 西班牙, 泰国, 新加坡, 马来西亚, 菲律宾, 其他）<br /> 支持最大长度为：30<br /> 支持的最大列表长度为：30 */
func (r *ItemUpdateRequest) SetGlobalStockCountry(value string) {
	r.SetValue("global_stock_country", value)
}

/* 全球购商品采购地（库存类型）
全球购商品有两种库存类型：现货和代购 参数值为1时代表现货，值为2时代表代购。注意：使用时请与 全球购商品采购地（地区/国家）配合使用 */
func (r *ItemUpdateRequest) SetGlobalStockType(value string) {
	r.SetValue("global_stock_type", value)
}

/* 支持会员打折。可选值:true,false; */
func (r *ItemUpdateRequest) SetHasDiscount(value string) {
	r.SetValue("has_discount", value)
}

/* 是否有发票。可选值:true,false (商城卖家此字段必须为true) */
func (r *ItemUpdateRequest) SetHasInvoice(value string) {
	r.SetValue("has_invoice", value)
}

/* 橱窗推荐。可选值:true,false; */
func (r *ItemUpdateRequest) SetHasShowcase(value string) {
	r.SetValue("has_showcase", value)
}

/* 是否有保修。可选值:true,false; */
func (r *ItemUpdateRequest) SetHasWarranty(value string) {
	r.SetValue("has_warranty", value)
}

/* 商品图片。类型:JPG,GIF;最大长度:500k<br /> 支持的文件类型为：gif,jpg,jpeg,png<br /> 支持的最大列表长度为：524288 */
func (r *ItemUpdateRequest) SetImage(value string) {
	r.SetValue("image", value)
}

/* 加价(降价)幅度。如果为0，代表系统代理幅度。对于增价拍和荷兰拍来说是加价幅度，对于降价拍来说是降价幅度。 */
func (r *ItemUpdateRequest) SetIncrement(value string) {
	r.SetValue("increment", value)
}

/* 用户自行输入的类目属性ID串，结构："pid1,pid2,pid3"，如："20000"（表示品牌） 注：通常一个类目下用户可输入的关键属性不超过1个。 */
func (r *ItemUpdateRequest) SetInputPids(value string) {
	r.SetValue("input_pids", value)
}

/* 用户自行输入的子属性名和属性值，结构:"父属性值;一级子属性名;一级子属性值;二级子属性名;自定义输入值,....",如：“耐克;耐克系列;科比系列;科比系列;2K5,Nike乔丹鞋;乔丹系列;乔丹鞋系列;乔丹鞋系列;json5”，多个自定义属性用','分割，input_str需要与input_pids一一对应，注：通常一个类目下用户可输入的关键属性不超过1个。所有属性别名加起来不能超过3999字节。此处不可以使用“其他”、“其它”和“其她”这三个词。 */
func (r *ItemUpdateRequest) SetInputStr(value string) {
	r.SetValue("input_str", value)
}

/* 是否是3D */
func (r *ItemUpdateRequest) SetIs3D(value string) {
	r.SetValue("is_3D", value)
}

/* 是否在外店显示 */
func (r *ItemUpdateRequest) SetIsEx(value string) {
	r.SetValue("is_ex", value)
}

/* 实物闪电发货。注意：在售的闪电发货产品不允许取消闪电发货，需要先下架商品才能取消闪电发货标记 */
func (r *ItemUpdateRequest) SetIsLightningConsignment(value string) {
	r.SetValue("is_lightning_consignment", value)
}

/* 是否是线下商品。 1：线上商品（默认值）； 2：线上或线下商品； 3：线下商品。 */
func (r *ItemUpdateRequest) SetIsOffline(value string) {
	r.SetValue("is_offline", value)
}

/* 是否替换sku */
func (r *ItemUpdateRequest) SetIsReplaceSku(value string) {
	r.SetValue("is_replace_sku", value)
}

/* 是否在淘宝上显示（如果传FALSE，则在淘宝主站无法显示该商品） */
func (r *ItemUpdateRequest) SetIsTaobao(value string) {
	r.SetValue("is_taobao", value)
}

/* 商品是否为新品。只有在当前类目开通新品,并且当前用户拥有该类目下发布新品权限时才能设置is_xinpin为true，否则设置true后会返回错误码:isv.invalid-permission:xinpin。同时只有一口价全新的宝贝才能设置为新品，否则会返回错误码：isv.invalid-parameter:xinpin。不设置参数就保持原有值。 */
func (r *ItemUpdateRequest) SetIsXinpin(value string) {
	r.SetValue("is_xinpin", value)
}

/* 表示商品的体积，如果需要使用按体积计费的运费模板，一定要设置这个值。该值的单位为立方米（m3），如果是其他单位，请转换成成立方米。
该值支持两种格式的设置：格式1：bulk:3,单位为立方米(m3),表示直接设置为商品的体积。格式2：length:10;breadth:10;height:10，单位为米（m）。体积和长宽高都支持小数类型。
在传入体积或长宽高时候，不能带单位。体积的单位默认为立方米（m3），长宽高的单位默认为米(m)
在编辑的时候，如果需要删除体积属性，请设置该值为0，如bulk:0 */
func (r *ItemUpdateRequest) SetItemSize(value string) {
	r.SetValue("item_size", value)
}

/* 商品的重量，用于按重量计费的运费模板。注意：单位为kg。 只能传入数值类型（包含小数），不能带单位，单位默认为kg。 在编辑时候，如果需要在商品里删除重量的信息，就需要将值设置为0 */
func (r *ItemUpdateRequest) SetItemWeight(value string) {
	r.SetValue("item_weight", value)
}

/* 商品文字的版本，繁体传入”zh_HK”，简体传入”zh_CN” */
func (r *ItemUpdateRequest) SetLang(value string) {
	r.SetValue("lang", value)
}

/* 上架时间。大于当前时间则宝贝会下架进入定时上架的宝贝中。 */
func (r *ItemUpdateRequest) SetListTime(value string) {
	r.SetValue("list_time", value)
}

/* 编辑电子凭证宝贝时候表示是否使用邮寄
0: 代表不使用邮寄；
1：代表使用邮寄；
如果不设置这个值，代表不使用邮寄 */
func (r *ItemUpdateRequest) SetLocalityLifeChooseLogis(value string) {
	r.SetValue("locality_life.choose_logis", value)
}

/* 本地生活电子交易凭证业务，目前此字段只涉及到的信息为有效期;
如果有效期为起止日期类型，此值为2012-08-06,2012-08-16
如果有效期为【购买成功日 至】类型则格式为2012-08-16
如果有效期为天数类型则格式为15 */
func (r *ItemUpdateRequest) SetLocalityLifeExpirydate(value string) {
	r.SetValue("locality_life.expirydate", value)
}

/* 码商信息，格式为 码商id:nick */
func (r *ItemUpdateRequest) SetLocalityLifeMerchant(value string) {
	r.SetValue("locality_life.merchant", value)
}

/* 网点ID,在参数empty_fields里设置locality_life.network_id可删除网点ID */
func (r *ItemUpdateRequest) SetLocalityLifeNetworkId(value string) {
	r.SetValue("locality_life.network_id", value)
}

/* 电子凭证售中自动退款比例，百分比%前的数字，介于1-100之间的整数 */
func (r *ItemUpdateRequest) SetLocalityLifeOnsaleAutoRefundRatio(value string) {
	r.SetValue("locality_life.onsale_auto_refund_ratio", value)
}

/* 退款比例，百分比%前的数字,1-100的正整数值; 在参数empty_fields里设置locality_life.refund_ratio可删除退款比例 */
func (r *ItemUpdateRequest) SetLocalityLifeRefundRatio(value string) {
	r.SetValue("locality_life.refund_ratio", value)
}

/* 退款码费承担方。发布电子凭证宝贝的时候会增加“退款码费承担方”配置项，可选填：(1)s（卖家承担） (2)b(买家承担) */
func (r *ItemUpdateRequest) SetLocalityLifeRefundmafee(value string) {
	r.SetValue("locality_life.refundmafee", value)
}

/* 核销打款,1代表核销打款 0代表非核销打款; 在参数empty_fields里设置locality_life.verification可删除核销打款 */
func (r *ItemUpdateRequest) SetLocalityLifeVerification(value string) {
	r.SetValue("locality_life.verification", value)
}

/* 所在地城市。如杭州 */
func (r *ItemUpdateRequest) SetLocationCity(value string) {
	r.SetValue("location.city", value)
}

/* 所在地省份。如浙江 */
func (r *ItemUpdateRequest) SetLocationState(value string) {
	r.SetValue("location.state", value)
}

/* 该宝贝是否支持【7天无理由退货】，卖家选择的值只是一个因素，最终以类目和选择的属性条件来确定是否支持7天。填入字符0，表示不支持；未填写或填人字符1，表示支持7天无理由退货；<br>注意：使用该API修改商品其它属性如标题title时，如需保持商品不支持7天无理由退货状态，该字段需传入0 。 */
func (r *ItemUpdateRequest) SetNewprepay(value string) {
	r.SetValue("newprepay", value)
}

/* 商品数量，取值范围:0-900000000的整数。且需要等于Sku所有数量的和 拍卖商品中增加拍只能为1，荷兰拍要在[2,500)范围内。<br /> 支持最大值为：900000000<br /> 支持最小值为：0 */
func (r *ItemUpdateRequest) SetNum(value string) {
	r.SetValue("num", value)
}

/* 商品数字ID，该参数必须<br /> 支持最小值为：1 */
func (r *ItemUpdateRequest) SetNumIid(value string) {
	r.SetValue("num_iid", value)
}

/* 商家编码 */
func (r *ItemUpdateRequest) SetOuterId(value string) {
	r.SetValue("outer_id", value)
}

/* 拍卖宝贝的保证金。对于增价拍和荷兰拍来说保证金有两种模式：淘宝默认模式（首次出价金额的10%），自定义固定保证金（固定冻结金额只能输入不超过30万的正整数），并且保证金只冻结1次。对于降价拍来说保证金只有淘宝默认的（竞拍价格的10% * 竞拍数量），并且每次出价都需要冻结保证金。
对于拍卖宝贝来说，保证金是必须的，但是默认使用淘宝默认保证金模式，只有用户需要使用自定义固定保证金的时候才需要使用到这个参数。如果该参数不传或传入0则代表使用默认。 */
func (r *ItemUpdateRequest) SetPaimaiInfoDeposit(value string) {
	r.SetValue("paimai_info.deposit", value)
}

/* 降价拍宝贝的降价周期(分钟)。降价拍宝贝的价格每隔paimai_info.interval时间会下降一次increment。<br /> 支持最大值为：60<br /> 支持最小值为：1 */
func (r *ItemUpdateRequest) SetPaimaiInfoInterval(value string) {
	r.SetValue("paimai_info.interval", value)
}

/* 拍卖商品选择的拍卖类型，拍卖类型包括三种：增价拍(1)，荷兰拍(2)和降价拍(3)。<br /> 支持最大值为：3<br /> 支持最小值为：1 */
func (r *ItemUpdateRequest) SetPaimaiInfoMode(value string) {
	r.SetValue("paimai_info.mode", value)
}

/* 降价拍宝贝的保留价。对于降价拍来说，paimai_info.reserve必须大于0，且小于price-increment，而且（price-paimai_info.reserve）/increment的计算结果必须为整数 */
func (r *ItemUpdateRequest) SetPaimaiInfoReserve(value string) {
	r.SetValue("paimai_info.reserve", value)
}

/* 自定义销售周期的小时数。拍卖宝贝可以自定义销售周期，这里指定销售周期的小时数。自定义销售周期的小时数。拍卖宝贝可以自定义销售周期，这里指定销售周期的小时数。注意，该参数只作为输入参数，不能通过taobao.item.get接口获取。<br /> 支持最大值为：48<br /> 支持最小值为：1 */
func (r *ItemUpdateRequest) SetPaimaiInfoValidHour(value string) {
	r.SetValue("paimai_info.valid_hour", value)
}

/* 自定义销售周期的分钟数。拍卖宝贝可以自定义销售周期，这里是指定销售周期的分钟数。自定义销售周期的小时数。拍卖宝贝可以自定义销售周期，这里指定销售周期的小时数。注意，该参数只作为输入参数，不能通过taobao.item.get接口获取。<br /> 支持最大值为：59<br /> 支持最小值为：0 */
func (r *ItemUpdateRequest) SetPaimaiInfoValidMinute(value string) {
	r.SetValue("paimai_info.valid_minute", value)
}

/* 商品主图需要关联的图片空间的相对url。这个url所对应的图片必须要属于当前用户。pic_path和image只需要传入一个,如果两个都传，默认选择pic_path */
func (r *ItemUpdateRequest) SetPicPath(value string) {
	r.SetValue("pic_path", value)
}

/* 平邮费用。取值范围:0.01-999.00;精确到2位小数;单位:元。如:5.07，表示:5元7分, 注:post_fee,express_fee,ems_fee需一起填写 */
func (r *ItemUpdateRequest) SetPostFee(value string) {
	r.SetValue("post_fee", value)
}

/* 宝贝所属的运费模板ID。取值范围：整数且必须是该卖家的运费模板的ID（可通过taobao.delivery.templates.get获得当前会话用户的所有邮费模板） */
func (r *ItemUpdateRequest) SetPostageId(value string) {
	r.SetValue("postage_id", value)
}

/* 商品价格。取值范围:0-100000000;精确到2位小数;单位:元。如:200.07，表示:200元7分。需要在正确的价格区间内。 拍卖商品对应的起拍价。 */
func (r *ItemUpdateRequest) SetPrice(value string) {
	r.SetValue("price", value)
}

/* 商品所属的产品ID(B商家发布商品需要用) */
func (r *ItemUpdateRequest) SetProductId(value string) {
	r.SetValue("product_id", value)
}

/* 属性值别名。如pid:vid:别名;pid1:vid1:别名1， pid:属性id vid:值id。总长度不超过800字节<br /> 支持最大长度为：800<br /> 支持的最大列表长度为：800 */
func (r *ItemUpdateRequest) SetPropertyAlias(value string) {
	r.SetValue("property_alias", value)
}

/* 商品属性列表。格式:pid:vid;pid:vid。属性的pid调用taobao.itemprops.get取得，属性值的vid用taobao.itempropvalues.get取得vid。 如果该类目下面没有属性，可以不用填写。如果有属性，必选属性必填，其他非必选属性可以选择不填写.属性不能超过35对。所有属性加起来包括分割符不能超过549字节，单个属性没有限制。 如果有属性是可输入的话，则用字段input_str填入属性的值。 */
func (r *ItemUpdateRequest) SetProps(value string) {
	r.SetValue("props", value)
}

/* 景区门票在选择订金支付时候，需要交的预订费。传入的值是1到20之间的数值，小数点后最多可以保留两位（多余的部分将做四舍五入的处理）。这个数值表示的是预订费的比例，最终的预订费为 scenic_ticket_book_cost乘一口价除以100 */
func (r *ItemUpdateRequest) SetScenicTicketBookCost(value string) {
	r.SetValue("scenic_ticket_book_cost", value)
}

/* 景区门票类宝贝编辑时候，当卖家签订了支付宝代扣协议时候，需要选择支付方式：全额支付和订金支付。当scenic_ticket_pay_way为1时表示全额支付，为2时表示订金支付 */
func (r *ItemUpdateRequest) SetScenicTicketPayWay(value string) {
	r.SetValue("scenic_ticket_pay_way", value)
}

/* 商品卖点信息，最长150个字符。天猫和集市都可用<br /> 支持最大长度为：150<br /> 支持的最大列表长度为：150 */
func (r *ItemUpdateRequest) SetSellPoint(value string) {
	r.SetValue("sell_point", value)
}

/* 是否承诺退换货服务!虚拟商品无须设置此项! */
func (r *ItemUpdateRequest) SetSellPromise(value string) {
	r.SetValue("sell_promise", value)
}

/* 重新关联商品与店铺类目，结构:",cid1,cid2,...,"，如果店铺类目存在二级类目，必须传入子类目cids。 */
func (r *ItemUpdateRequest) SetSellerCids(value string) {
	r.SetValue("seller_cids", value)
}

/* sku层面的条形码，多个SKU情况，与SKU价格库存格式类似，用逗号分隔 */
func (r *ItemUpdateRequest) SetSkuBarcode(value string) {
	r.SetValue("sku_barcode", value)
}

/* Sku的外部id串，结构如：1234,1342,… sku_properties, sku_quantities, sku_prices, sku_outer_ids在输入数据时要一一对应，如果没有sku_outer_ids也要写上这个参数，入参是","(这个是两个sku的示列，逗号数应该是sku个数减1)；该参数最大长度是512个字节 */
func (r *ItemUpdateRequest) SetSkuOuterIds(value string) {
	r.SetValue("sku_outer_ids", value)
}

/* 更新的Sku的价格串，结构如：10.00,5.00,… 精确到2位小数;单位:元。如:200.07，表示:200元7分 */
func (r *ItemUpdateRequest) SetSkuPrices(value string) {
	r.SetValue("sku_prices", value)
}

/* 更新的Sku的属性串，调用taobao.itemprops.get获取类目属性，如果属性是销售属性，再用taobao.itempropvalues.get取得vid。格式:pid:vid;pid:vid,多个sku之间用逗号分隔。该字段内的销售属性(自定义的除外)也需要在props字段填写 . 规则：如果该SKU存在旧商品，则修改；否则新增Sku。如果更新时有对Sku进行操作，则Sku的properties一定要传入。如果存在自定义销售属性，则格式为pid:vid;pid2:vid2;$pText:vText，其中$pText:vText为自定义属性。限制：其中$pText的’$’前缀不能少，且pText和vText文本中不可以存在 冒号:和分号;以及逗号 */
func (r *ItemUpdateRequest) SetSkuProperties(value string) {
	r.SetValue("sku_properties", value)
}

/* 更新的Sku的数量串，结构如：num1,num2,num3 如:2,3,4 */
func (r *ItemUpdateRequest) SetSkuQuantities(value string) {
	r.SetValue("sku_quantities", value)
}

/* 暂时不可用 */
func (r *ItemUpdateRequest) SetSkuSpecIds(value string) {
	r.SetValue("sku_spec_ids", value)
}

/* 商品新旧程度。可选值:new（全新）,unused（闲置）,second（二手）。 */
func (r *ItemUpdateRequest) SetStuffStatus(value string) {
	r.SetValue("stuff_status", value)
}

/* 商品是否支持拍下减库存:1支持;2取消支持(付款减库存);0(默认)不更改
集市卖家默认拍下减库存;
商城卖家默认付款减库存 */
func (r *ItemUpdateRequest) SetSubStock(value string) {
	r.SetValue("sub_stock", value)
}

/* 宝贝标题. 不能超过30字符,受违禁词控制<br /> 支持最大长度为：120<br /> 支持的最大列表长度为：120 */
func (r *ItemUpdateRequest) SetTitle(value string) {
	r.SetValue("title", value)
}

/* 有效期。可选值:7,14;单位:天; */
func (r *ItemUpdateRequest) SetValidThru(value string) {
	r.SetValue("valid_thru", value)
}

/* 商品的重量(商超卖家专用字段) */
func (r *ItemUpdateRequest) SetWeight(value string) {
	r.SetValue("weight", value)
}
