package api

import "net/url"

/* 用户信用 */
type UserCredit struct {
	/* 收到的好评总条数。取值范围:大于零的整数 */
	GoodNum int64 `json:"good_num"`

	/* 信用等级（是根据score生成的），信用等级：淘宝会员在淘宝网上的信用度，分为20个级别，级别如：level = 1 时，表示一心；level = 2 时，表示二心 */
	Level int64 `json:"level"`

	/* 信用总分（“好评”加一分，“中评”不加分，“差评”扣一分。分越高，等级越高） */
	Score int64 `json:"score"`

	/* 收到的评价总条数。取值范围:大于零的整数 */
	TotalNum int64 `json:"total_num"`
}

/* 批次库存查询结果记录 */
type WlbItemBatch struct {
	/* 批次编号 */
	BatchCode string `json:"batch_code"`

	/* 创建者 */
	Creator string `json:"creator"`

	/* 残次数量 */
	DefectQuantity int64 `json:"defect_quantity"`

	/* 到期时间 */
	DueDate string `json:"due_date"`

	/* 创建时间 */
	GmtCreate string `json:"gmt_create"`

	/* 最后修改时间 */
	GmtModified string `json:"gmt_modified"`

	/* 保质期 */
	GuaranteePeriod string `json:"guarantee_period"`

	/* 天（单位） */
	GuaranteeUnit int64 `json:"guarantee_unit"`

	/* 商品批次记录id */
	Id int64 `json:"id"`

	/* 是否删除。0：正常 1：删除 */
	IsDeleted bool `json:"is_deleted"`

	/* 商品id */
	ItemId int64 `json:"item_id"`

	/* 最后修改者 */
	LastModifier string `json:"last_modifier"`

	/* 产地 */
	ProduceArea string `json:"produce_area"`

	/* 生产编号 */
	ProduceCode string `json:"produce_code"`

	/* 生产日期 */
	ProduceDate string `json:"produce_date"`

	/* 商品数量 */
	Quantity int64 `json:"quantity"`

	/* 接受日期 */
	ReceiveDate string `json:"receive_date"`

	/* 描述 */
	Remarks string `json:"remarks"`

	/* 状态。item_batch_status_open:开放 item_batch_status_lock:冻结 item_batch_status_invalid:无效 */
	Status string `json:"status"`

	/* 存储类型 */
	StoreCode string `json:"store_code"`

	/* 用户id */
	UserId int64 `json:"user_id"`

	/* 版本 */
	Version int64 `json:"version"`
}

/* 用户地址 */
type Location struct {
	/* 详细地址，最大256个字节（128个中文） */
	Address string `json:"address"`

	/* 所在城市（中文名称） */
	City string `json:"city"`

	/* 国家名称 */
	Country string `json:"country"`

	/* 区/县（只适用于物流API） */
	District string `json:"district"`

	/* 所在省份（中文名称） */
	State string `json:"state"`

	/* 邮政编码 */
	Zip string `json:"zip"`
}

/* 用户 */
type User struct {
	/* 有无绑定。可选值:bind(绑定),notbind(未绑定) */
	AlipayBind string `json:"alipay_bind"`

	/* 是否受限制。可选值:limited(受限制),unlimited(不受限) */
	AutoRepost string `json:"auto_repost"`

	/* 用户头像地址 */
	Avatar string `json:"avatar"`

	/* 生日 */
	Birthday string `json:"birthday"`

	/* 买家信用 */
	BuyerCredit *UserCredit `json:"buyer_credit"`

	/* 是否参加消保 */
	ConsumerProtection bool `json:"consumer_protection"`

	/* 用户注册时间。格式:yyyy-MM-dd HH:mm:ss */
	Created string `json:"created"`

	/* 联系人email */
	Email string `json:"email"`

	/* 是否购买多图服务。可选值:true(是),false(否) */
	HasMorePic bool `json:"has_more_pic"`

	/* 用户作为卖家是否开过店 */
	HasShop bool `json:"has_shop"`

	/* 表示用户是否具备修改商品减库存逻辑的权限（一共有拍下减库存和付款减库存两种逻辑） 值含义： 1）true：是 2）false：否。 */
	HasSubStock bool `json:"has_sub_stock"`

	/* 用户是否是金牌卖家 */
	IsGoldenSeller bool `json:"is_golden_seller"`

	/* 是否24小时闪电发货(实物类) */
	IsLightningConsignment bool `json:"is_lightning_consignment"`

	/* 可上传商品图片数量 */
	ItemImgNum int64 `json:"item_img_num"`

	/* 单张商品图片最大容量(商品主图大小)。单位:k */
	ItemImgSize int64 `json:"item_img_size"`

	/* 最近登陆时间。格式:yyyy-MM-dd HH:mm:ss */
	LastVisit string `json:"last_visit"`

	/* 是否是无名良品用户，true or false */
	Liangpin bool `json:"liangpin"`

	/* 用户当前居住地公开信息。如：location.city获取其中的city数据 */
	Location *Location `json:"location"`

	/* 是否订阅了淘宝天下杂志 */
	MagazineSubscribe bool `json:"magazine_subscribe"`

	/* 用户昵称 */
	Nick string `json:"nick"`

	/* 用户是否为网游用户，属于隐私信息，需要登陆才能查看自己的。 目前仅供taobao.user.get使用 */
	OnlineGaming bool `json:"online_gaming"`

	/* 有无实名认证。可选值:authentication(实名认证),not authentication(没有认证) */
	PromotedType string `json:"promoted_type"`

	/* 可上传属性图片数量 */
	PropImgNum int64 `json:"prop_img_num"`

	/* 单张销售属性图片最大容量（非主图的商品图片和商品属性图片）。单位:k */
	PropImgSize int64 `json:"prop_img_size"`

	/* 卖家信用 */
	SellerCredit *UserCredit `json:"seller_credit"`

	/* 性别。可选值:m(男),f(女) */
	Sex string `json:"sex"`

	/* 卖家是否签署食品卖家承诺协议 */
	SignFoodSellerPromise bool `json:"sign_food_seller_promise"`

	/* 状态。可选值:normal(正常),inactive(未激活),delete(删除),reeze(冻结),supervise(监管) */
	Status string `json:"status"`

	/* 用户类型。可选值:B(B商家),C(C商家) */
	Type string `json:"type"`

	/* 用户字符串ID */
	Uid string `json:"uid"`

	/* 用户数字ID */
	UserId int64 `json:"user_id"`

	/* 用户参与垂直市场类型。shoes表示鞋城垂直市场用户，3C表示3C垂直市场用户。多个类型之间用","分隔。如：一个用户既是3C用户又是鞋城用户，那么这个字段返回就是"3C,shoes"。如果用户不是垂直市场用户，此字段返回为""。 */
	VerticalMarket string `json:"vertical_market"`

	/* 用户的全站vip信息，可取值如下：c(普通会员),asso_vip(荣誉会员)，vip1,vip2,vip3,vip4,vip5,vip6(六个等级的正式vip会员)，共8种取值，其中asso_vip是由vip会员衰退而成，与主站上的vip0对应。 */
	VipInfo string `json:"vip_info"`
}

/* 产品图片 */
type ProductImg struct {
	/* 添加时间.格式:yyyy-mm-dd hh:mm:ss */
	Created string `json:"created"`

	/* 产品图片ID */
	Id int64 `json:"id"`

	/* 修改时间.格式:yyyy-mm-dd hh:mm:ss */
	Modified string `json:"modified"`

	/* 图片序号。产品里的图片展示顺序，数据越小越靠前。要求是正整数。 */
	Position int64 `json:"position"`

	/* 图片所属产品的ID */
	ProductId int64 `json:"product_id"`

	/* 图片地址.(绝对地址,格式:http://host/image_path) */
	Url string `json:"url"`
}

/* 图书类目导入返回结果 */
type ProductBooks struct {
	/* 作者/著者 */
	Author string `json:"author"`

	/* 条形码，13位，9787开头 */
	BarCode string `json:"bar_code"`

	/* 完整的图书名称 */
	BookName string `json:"book_name"`

	/* 类目id */
	CategoryId int64 `json:"category_id"`

	/* ISBN号 */
	Isbn string `json:"isbn"`

	/* 图书价格，若有多个，则以';'号分隔 */
	Price string `json:"price"`
}

/* 图片空间的用户信息获取，包括订购容量等 */
type UserInfo struct {
	/* 用户的可用容量，即订购量与免费量之和 */
	AvailableSpace string `json:"available_space"`

	/* 图片空间的免费容量 */
	FreeSpace string `json:"free_space"`

	/* 图片空间的订购有效期 */
	OrderExpiryDate string `json:"order_expiry_date"`

	/* 用户订购的图片空间容量 */
	OrderSpace string `json:"order_space"`

	/* 剩余的图片空间容量 */
	RemainingSpace string `json:"remaining_space"`

	/* 已使用的图片空间容量 */
	UsedSpace string `json:"used_space"`

	/* 用户自定义的水印参数，通过"|"分割开，如果用户没有定义则为""
	具体水印参数组合方法，用"|"分开，顺序按"是否全局设置|水印文字|是否文字水印优先|透明度|字体|字体大小|字体是否加粗|字体是否斜体|字体是否加下划线|字体颜色|旋转角度|是否带阴影|水印位置|图片水印URL|reference水印相对位置" reference取值有左上（1）/中间（3）/右下（2）,其中的null代表为空 */
	WaterMark string `json:"water_mark"`
}

/* 产品结构 */
type Product struct {
	/* 产品条码信息，仅在taobao.products.search接口且商城可用 */
	BarcodeStr string `json:"barcode_str"`

	/* 产品的非关键属性列表.格式:pid:vid;pid:vid. */
	Binds string `json:"binds"`

	/* 产品的非关键属性字符串列表.格式同props_str(<strong>注：</strong><font color="red">属性名称中的冒号":"被转换为："#cln#";
	分号";"被转换为："#scln#"
	</font>) */
	BindsStr string `json:"binds_str"`

	/* 商品类目名称 */
	CatName string `json:"cat_name"`

	/* 商品类目ID.必须是叶子类目ID */
	Cid int64 `json:"cid"`

	/* 产品的collect次数（不提供数据，返回0) */
	CollectNum int64 `json:"collect_num"`

	/* 品类ID */
	CommodityId int64 `json:"commodity_id"`

	/* 创建时间.格式:yyyy-mm-dd hh:mm:ss */
	Created string `json:"created"`

	/* 标识是否为达尔文体系下的产品 。
	如果为空表示是非达尔文体系下的产品
	如果cspu:0 表示是达尔文体系下的产品，有cspu正在待小二审核，但不能发布商品。
	如果cspu:1 表示是达尔文体系下的产品，且有小二确认的cspu，能发布商品 */
	CspuFeature string `json:"cspu_feature"`

	/* 用户自定义属性,结构：pid1:value1;pid2:value2 例如：“20000:优衣库”，表示“品牌:优衣库” */
	CustomerProps string `json:"customer_props"`

	/* 产品的描述.最大25000个字节 */
	Desc string `json:"desc"`

	/* 标识套装产品是否有效，无效的套装产品需要重新发布 */
	IsSuiteEffective bool `json:"is_suite_effective"`

	/* 产品的级别level */
	Level int64 `json:"level"`

	/* 修改时间.格式:yyyy-mm-dd hh:mm:ss */
	Modified string `json:"modified"`

	/* 产品名称 */
	Name string `json:"name"`

	/* 外部产品ID */
	OuterId string `json:"outer_id"`

	/* 产品对应的图片路径 */
	PicPath string `json:"pic_path"`

	/* 产品的主图片地址.(绝对地址,格式:http://host/image_path) */
	PicUrl string `json:"pic_url"`

	/* 产品的市场价.单位为元.精确到2位小数;如:200.07 */
	Price float64 `json:"price,string"`

	/* 产品扩展信息 */
	ProductExtraInfos []*ProductExtraInfo `json:"product_extra_infos>product_extra_info"`

	/* 产品ID */
	ProductId int64 `json:"product_id"`

	/* 产品的子图片.目前最多支持4张。fields中设置为product_imgs.id、product_imgs.url、product_imgs.position 等形式就会返回相应的字段 */
	ProductImgs []*ProductImg `json:"product_imgs>product_img"`

	/* 产品的属性图片.比如说黄色对应的产品图片,绿色对应的产品图片。fields中设置为product_prop_imgs.id、
	product_prop_imgs.props、product_prop_imgs.url、product_prop_imgs.position等形式就会返回相应的字段 */
	ProductPropImgs []*ProductPropImg `json:"product_prop_imgs>product_prop_img"`

	/* 销售属性值别名。格式为pid1:vid1:alias1;pid1:vid2:alia2。 */
	PropertyAlias string `json:"property_alias"`

	/* 产品的关键属性列表.格式：pid:vid;pid:vid */
	Props string `json:"props"`

	/* 产品的关键属性字符串列表.比如:品牌:诺基亚;型号:N73(<strong>注：</strong><font color="red">属性名称中的冒号":"被转换为："#cln#";
	分号";"被转换为："#scln#"
	</font>) */
	PropsStr string `json:"props_str"`

	/* 产品的评分次数 */
	RateNum int64 `json:"rate_num"`

	/* 产品的销售量 */
	SaleNum int64 `json:"sale_num"`

	/* 产品的销售属性列表.格式:pid:vid;pid:vid */
	SaleProps string `json:"sale_props"`

	/* 产品的销售属性字符串列表.格式同props_str(<strong>注：</strong><font color="red">属性名称中的冒号":"被转换为："#cln#";
	分号";"被转换为："#scln#"
	</font>) */
	SalePropsStr string `json:"sale_props_str"`

	/* 产品卖点描述，长度限制20个汉字 */
	SellPt string `json:"sell_pt"`

	/* 产品的店内价格 */
	ShopPrice string `json:"shop_price"`

	/* 产品的标准价格 */
	StandardPrice string `json:"standard_price"`

	/* 当前状态(0 商家确认 1 屏蔽 3 小二确认 2 未确认 -1 删除) */
	Status int64 `json:"status"`

	/* 套装产品关联的子规格,同时该字段不为空标识该产品是套装产品 */
	SuiteItemsStr string `json:"suite_items_str"`

	/* 模板ID */
	TemplateId int64 `json:"template_id"`

	/* 淘宝标准产品编码 */
	Tsc string `json:"tsc"`

	/* 垂直市场,如：3（3C），4（鞋城） */
	VerticalMarket int64 `json:"vertical_market"`
}

/* 退款超时 */
type RefundRemindTimeout struct {
	/* 是否存在超时。可选值:true(是),false(否) */
	ExistTimeout bool `json:"exist_timeout"`

	/* 提醒的类型（退款详情中提示信息的类型） */
	RemindType int64 `json:"remind_type"`

	/* 超时时间。格式:yyyy-MM-dd HH:mm:ss */
	Timeout string `json:"timeout"`
}

/* 退款留言 */
type RefundMessage struct {
	/* 留言内容。最大长度: 400个字节 */
	Content string `json:"content"`

	/* 留言创建时间。格式:yyyy-MM-dd HH:mm:ss */
	Created string `json:"created"`

	/* 留言编号 */
	Id int64 `json:"id"`

	/* 退款类型：NORMAL（普通留言），RETURN_GOODS_APPROVED（卖家留退货地址时留言）；如果为RETURN_GOODS_APPROVED，则退款留言中有卖家收货地址 */
	MessageType string `json:"message_type"`

	/* 留言者编号 */
	OwnerId int64 `json:"owner_id"`

	/* 留言者昵称 */
	OwnerNick string `json:"owner_nick"`

	/* 留言者身份1代表买家，2代表卖家，3代表小二 */
	OwnerRole string `json:"owner_role"`

	/* 凭证附件地址（图片） */
	PicUrls []*PicUrl `json:"pic_urls>pic_url"`

	/* 退款编号。 */
	RefundId int64 `json:"refund_id"`

	/* 退款阶段，可选值：onsale（售中）, aftersale（售后） */
	RefundPhase string `json:"refund_phase"`
}

/* 订单结构 */
type Order struct {
	/* 手工调整金额.格式为:1.01;单位:元;精确到小数点后两位. */
	AdjustFee string `json:"adjust_fee"`

	/* 捆绑的子订单号，表示该子订单要和捆绑的子订单一起发货，用于卖家子订单捆绑发货 */
	BindOid int64 `json:"bind_oid"`

	/* 买家昵称 */
	BuyerNick string `json:"buyer_nick"`

	/* 买家是否已评价。可选值：true(已评价)，false(未评价) */
	BuyerRate bool `json:"buyer_rate"`

	/* 交易商品对应的类目ID */
	Cid int64 `json:"cid"`

	/* 子订单发货时间，当卖家对订单进行了多次发货，子订单的发货时间和主订单的发货时间可能不一样了，那么就需要以子订单的时间为准。（没有进行多次发货的订单，主订单的发货时间和子订单的发货时间都一样） */
	ConsignTime string `json:"consign_time"`

	/* 子订单级订单优惠金额。精确到2位小数;单位:元。如:200.07，表示:200元7分 */
	DiscountFee string `json:"discount_fee"`

	/* 分摊之后的实付金额 */
	DivideOrderFee string `json:"divide_order_fee"`

	/* 子订单的交易结束时间
	说明：子订单有单独的结束时间，与主订单的结束时间可能有所不同，在有退款发起的时候或者是主订单分阶段付款的时候，子订单的结束时间会早于主订单的结束时间，所以开放这个字段便于订单结束状态的判断 */
	EndTime string `json:"end_time"`

	/* 商品的字符串编号(注意：iid近期即将废弃，请用num_iid参数) */
	Iid string `json:"iid"`

	/* 子订单所在包裹的运单号 */
	InvoiceNo string `json:"invoice_no"`

	/* 表示订单交易是否含有对应的代销采购单。
	如果该订单中存在一个对应的代销采购单，那么该值为true；反之，该值为false。 */
	IsDaixiao bool `json:"is_daixiao"`

	/* 是否超卖 */
	IsOversold bool `json:"is_oversold"`

	/* 是否是服务订单，是返回true，否返回false。 */
	IsServiceOrder bool `json:"is_service_order"`

	/* 子订单是否是www订单 */
	IsWww bool `json:"is_www"`

	/* 套餐ID */
	ItemMealId int64 `json:"item_meal_id"`

	/* 套餐的值。如：M8原装电池:便携支架:M8专用座充:莫凡保护袋 */
	ItemMealName string `json:"item_meal_name"`

	/* 子订单发货的快递公司名称 */
	LogisticsCompany string `json:"logistics_company"`

	/* 订单修改时间，目前只有taobao.trade.ordersku.update会返回此字段。 */
	Modified string `json:"modified"`

	/* 购买数量。取值范围:大于零的整数 */
	Num int64 `json:"num"`

	/* 商品数字ID */
	NumIid int64 `json:"num_iid"`

	/* 子订单编号 */
	Oid int64 `json:"oid"`

	/* 子订单来源,如jhs(聚划算)、taobao(淘宝)、wap(无线) */
	OrderFrom string `json:"order_from"`

	/* 商家外部编码(可与商家外部系统对接)。外部商家自己定义的商品Item的id，可以通过taobao.items.custom.get获取商品的Item的信息 */
	OuterIid string `json:"outer_iid"`

	/* 外部网店自己定义的Sku编号 */
	OuterSkuId string `json:"outer_sku_id"`

	/* 优惠分摊 */
	PartMjzDiscount string `json:"part_mjz_discount"`

	/* 子订单实付金额。精确到2位小数，单位:元。如:200.07，表示:200元7分。对于多子订单的交易，计算公式如下：payment = price * num + adjust_fee - discount_fee ；单子订单交易，payment与主订单的payment一致，对于退款成功的子订单，由于主订单的优惠分摊金额，会造成该字段可能不为0.00元。建议使用退款前的实付金额减去退款单中的实际退款金额计算。 */
	Payment string `json:"payment"`

	/* 商品图片的绝对路径 */
	PicPath string `json:"pic_path"`

	/* 商品价格。精确到2位小数;单位:元。如:200.07，表示:200元7分 */
	Price string `json:"price"`

	/* 最近退款ID */
	RefundId int64 `json:"refund_id"`

	/* 退款状态。退款状态。可选值 WAIT_SELLER_AGREE(买家已经申请退款，等待卖家同意) WAIT_BUYER_RETURN_GOODS(卖家已经同意退款，等待买家退货) WAIT_SELLER_CONFIRM_GOODS(买家已经退货，等待卖家确认收货) SELLER_REFUSE_BUYER(卖家拒绝退款) CLOSED(退款关闭) SUCCESS(退款成功) */
	RefundStatus string `json:"refund_status"`

	/* 卖家昵称 */
	SellerNick string `json:"seller_nick"`

	/* 卖家是否已评价。可选值：true(已评价)，false(未评价) */
	SellerRate bool `json:"seller_rate"`

	/* 卖家类型，可选值为：B（商城商家），C（普通卖家） */
	SellerType string `json:"seller_type"`

	/* 子订单的运送方式（卖家对订单进行多次发货之后，一个主订单下的子订单的运送方式可能不同，用order.shipping_type来区分子订单的运送方式） */
	ShippingType string `json:"shipping_type"`

	/* 商品的最小库存单位Sku的id.可以通过taobao.item.sku.get获取详细的Sku信息 */
	SkuId string `json:"sku_id"`

	/* SKU的值。如：机身颜色:黑色;手机套餐:官方标配 */
	SkuPropertiesName string `json:"sku_properties_name"`

	/* 订单快照详细信息 */
	Snapshot string `json:"snapshot"`

	/* 订单快照URL */
	SnapshotUrl string `json:"snapshot_url"`

	/* 订单状态（请关注此状态，如果为TRADE_CLOSED_BY_TAOBAO状态，则不要对此订单进行发货，切记啊！）。可选值:
	<ul>
	<li>TRADE_NO_CREATE_PAY(没有创建支付宝交易)
	<li>WAIT_BUYER_PAY(等待买家付款)
	<li>WAIT_SELLER_SEND_GOODS(等待卖家发货,即:买家已付款)
	<li>WAIT_BUYER_CONFIRM_GOODS(等待买家确认收货,即:卖家已发货)
	<li>TRADE_BUYER_SIGNED(买家已签收,货到付款专用)
	<li>TRADE_FINISHED(交易成功)
	<li>TRADE_CLOSED(付款以后用户退款成功，交易自动关闭)
	<li>TRADE_CLOSED_BY_TAOBAO(付款以前，卖家或买家主动关闭交易)
	<li>PAY_PENDING(国际信用卡支付付款确认中) */
	Status string `json:"status"`

	/* 发货的仓库编码 */
	StoreCode string `json:"store_code"`

	/* 门票有效期的key */
	TicketExpdateKey string `json:"ticket_expdate_key"`

	/* 对应门票有效期的外部id */
	TicketOuterId string `json:"ticket_outer_id"`

	/* 订单超时到期时间。格式:yyyy-MM-dd HH:mm:ss */
	TimeoutActionTime string `json:"timeout_action_time"`

	/* 商品标题 */
	Title string `json:"title"`

	/* 应付金额（商品价格 * 商品数量 + 手工调整金额 - 子订单级订单优惠金额）。精确到2位小数;单位:元。如:200.07，表示:200元7分 */
	TotalFee string `json:"total_fee"`
}

/* 退款详情 */
type Refund struct {
	/* 卖家收货地址 */
	Address string `json:"address"`

	/* 退款先行垫付默认的未申请状态 0;退款先行垫付申请中  1;退款先行垫付，垫付完成 2;退款先行垫付，卖家拒绝收货 3;退款先行垫付，垫付关闭 4;退款先行垫付，垫付分账成功 5; */
	AdvanceStatus int64 `json:"advance_status"`

	/* 支付宝交易号 */
	AlipayNo string `json:"alipay_no"`

	/* 退款扩展属性 */
	Attribute string `json:"attribute"`

	/* 买家昵称 */
	BuyerNick string `json:"buyer_nick"`

	/* 物流公司名称 */
	CompanyName string `json:"company_name"`

	/* 退款申请时间。格式:yyyy-MM-dd HH:mm:ss */
	Created string `json:"created"`

	/* 不需客服介入1;需要客服介入2;客服已经介入3;客服初审完成 4;客服主管复审失败5;客服处理完成6; */
	CsStatus int64 `json:"cs_status"`

	/* 退款说明 */
	Desc string `json:"desc"`

	/* 退货时间。格式:yyyy-MM-dd HH:mm:ss */
	GoodReturnTime string `json:"good_return_time"`

	/* 货物状态。可选值BUYER_NOT_RECEIVED (买家未收到货) BUYER_RECEIVED (买家已收到货) BUYER_RETURNED_GOODS (买家已退货) */
	GoodStatus string `json:"good_status"`

	/* 买家是否需要退货。可选值:true(是),false(否) */
	HasGoodReturn bool `json:"has_good_return"`

	/* 申请退款的商品字符串编号(注意：iid近期即将废弃，请用num_iid参数) */
	Iid string `json:"iid"`

	/* 更新时间。格式:yyyy-MM-dd HH:mm:ss */
	Modified string `json:"modified"`

	/* 商品购买数量 */
	Num int64 `json:"num"`

	/* 申请退款的商品数字编号 */
	NumIid int64 `json:"num_iid"`

	/* 子订单号。如果是单笔交易oid会等于tid */
	Oid int64 `json:"oid"`

	/* 退款约束，可选值：cannot_refuse（不允许操作），refund_onweb（需要到网页版操作） */
	OperationContraint string `json:"operation_contraint"`

	/* 退款对应的订单交易状态。可选值TRADE_NO_CREATE_PAY(没有创建支付宝交易) WAIT_BUYER_PAY(等待买家付款) WAIT_SELLER_SEND_GOODS(等待卖家发货,即:买家已付款) WAIT_BUYER_CONFIRM_GOODS(等待买家确认收货,即:卖家已发货) TRADE_BUYER_SIGNED(买家已签收,货到付款专用) TRADE_FINISHED(交易成功) TRADE_CLOSED(交易关闭) TRADE_CLOSED_BY_TAOBAO(交易被淘宝关闭) ALL_WAIT_PAY(包含：WAIT_BUYER_PAY、TRADE_NO_CREATE_PAY) ALL_CLOSED(包含：TRADE_CLOSED、TRADE_CLOSED_BY_TAOBAO) 取自"http://open.taobao.com/dev/index.php/%E4%BA%A4%E6%98%93%E7%8A%B6%E6%80%81" */
	OrderStatus string `json:"order_status"`

	/* 商品外部商家编码 */
	OuterId string `json:"outer_id"`

	/* 支付给卖家的金额(交易总金额-退还给买家的金额)。精确到2位小数;单位:元。如:200.07，表示:200元7分 */
	Payment string `json:"payment"`

	/* 商品价格。精确到2位小数;单位:元。如:200.07，表示:200元7分 */
	Price string `json:"price"`

	/* 退款原因 */
	Reason string `json:"reason"`

	/* 退还金额(退还给买家的金额)。精确到2位小数;单位:元。如:200.07，表示:200元7分 */
	RefundFee string `json:"refund_fee"`

	/* 退款单号 */
	RefundId int64 `json:"refund_id"`

	/* 退款阶段，可选值：onsale/aftersale */
	RefundPhase string `json:"refund_phase"`

	/* 退款超时结构RefundRemindTimeout */
	RefundRemindTimeout *RefundRemindTimeout `json:"refund_remind_timeout"`

	/* 退款版本号（时间戳） */
	RefundVersion int64 `json:"refund_version"`

	/* 卖家昵称 */
	SellerNick string `json:"seller_nick"`

	/* 物流方式.可选值:free(卖家包邮),post(平邮),express(快递),ems(EMS). */
	ShippingType string `json:"shipping_type"`

	/* 退货运单号 */
	Sid string `json:"sid"`

	/* 商品SKU信息 */
	Sku string `json:"sku"`

	/* 分账给卖家的钱 */
	SplitSellerFee string `json:"split_seller_fee"`

	/* 分账给淘宝的钱 */
	SplitTaobaoFee string `json:"split_taobao_fee"`

	/* 退款状态。可选值WAIT_SELLER_AGREE(买家已经申请退款，等待卖家同意) WAIT_BUYER_RETURN_GOODS(卖家已经同意退款，等待买家退货) WAIT_SELLER_CONFIRM_GOODS(买家已经退货，等待卖家确认收货) SELLER_REFUSE_BUYER(卖家拒绝退款) CLOSED(退款关闭) SUCCESS(退款成功) */
	Status string `json:"status"`

	/* 淘宝交易单号 */
	Tid int64 `json:"tid"`

	/* 商品标题 */
	Title string `json:"title"`

	/* 交易总金额。精确到2位小数;单位:元。如:200.07，表示:200元7分 */
	TotalFee string `json:"total_fee"`
}

/* 运费模板对象 */
type DeliveryTemplate struct {
	/* 运费模板上设置的发货地址 */
	Address string `json:"address"`

	/* 可选值：0,1,2,3<br>,说明如下<br>
	1)买家承担运费的模版<br>
	0：买家承担服务费<br>
	1: 卖家承担服务费<br>
	2)卖家承担运费的模版<br>
	2:卖家承担运费的模版（集市），卖家承担服务费<br>
	3:卖家承担运费的模版（天猫），卖家承担服务费<br> */
	Assumer int64 `json:"assumer"`

	/* 该模板上设置的卖家发货地址区域ID，如：address为浙江省杭州市西湖去文三路XX号那么这个consign_area_id的值就是西湖区的ID */
	ConsignAreaId int64 `json:"consign_area_id"`

	/* 模板创建时间 */
	Created string `json:"created"`

	/* 运费模板中运费详细信息对象，包含默认运费和指定地区运费 */
	FeeList []*TopFee `json:"fee_list>top_fee"`

	/* 模板修改时间 */
	Modified string `json:"modified"`

	/* 模板名称，长度不能超过25个字节 */
	Name string `json:"name"`

	/* 物流服务模板支持增值服务列表，多个用分号隔开。cod:货到付款 mops：刷卡付款 */
	Supports string `json:"supports"`

	/* 模块ID */
	TemplateId int64 `json:"template_id"`

	/* 可选值：0
	说明：
	0:表示按宝贝件数计算运费
	<br/><br/>
	1:表示按宝贝重量计算运费
	<br/><br/>
	3:表示按宝贝体积计算运费 */
	Valuation int64 `json:"valuation"`
}

/* 交易结构 */
type Trade struct {
	/* 卖家手工调整金额，精确到2位小数，单位：元。如：200.07，表示：200元7分。来源于订单价格修改，如果有多笔子订单的时候，这个为0，单笔的话则跟[order].adjust_fee一样 */
	AdjustFee string `json:"adjust_fee"`

	/* 买家的支付宝id号，在UIC中有记录，买家支付宝的唯一标示，不因为买家更换Email账号而改变。 */
	AlipayId int64 `json:"alipay_id"`

	/* 支付宝交易号，如：2009112081173831 */
	AlipayNo string `json:"alipay_no"`

	/* 付款时使用的支付宝积分的额度,单位分，比如返回1，则为1分钱 */
	AlipayPoint int64 `json:"alipay_point"`

	/* 创建交易接口成功后，返回的支付url */
	AlipayUrl string `json:"alipay_url"`

	/* 淘宝下单成功了,但由于某种错误支付宝订单没有创建时返回的信息。taobao.trade.add接口专用 */
	AlipayWarnMsg string `json:"alipay_warn_msg"`

	/* 区域id，代表订单下单的区位码，区位码是通过省市区转换而来，通过区位码能精确到区内的划分，比如310012是杭州市西湖区华星路 */
	AreaId string `json:"area_id"`

	/* 物流到货时效截单时间，格式 HH:mm */
	ArriveCutTime string `json:"arrive_cut_time"`

	/* 物流到货时效，单位天 */
	ArriveInterval int64 `json:"arrive_interval"`

	/* 同步到卖家库的时间，taobao.trades.sold.incrementv.get接口返回此字段 */
	AsyncModified string `json:"async_modified"`

	/* 交易中剩余的确认收货金额（这个金额会随着子订单确认收货而不断减少，交易成功后会变为零）。精确到2位小数;单位:元。如:200.07，表示:200 元7分 */
	AvailableConfirmFee string `json:"available_confirm_fee"`

	/* 买家支付宝账号 */
	BuyerAlipayNo string `json:"buyer_alipay_no"`

	/* 买家下单的地区 */
	BuyerArea string `json:"buyer_area"`

	/* 买家货到付款服务费。精确到2位小数;单位:元。如:12.07，表示:12元7分 */
	BuyerCodFee string `json:"buyer_cod_fee"`

	/* 买家邮件地址 */
	BuyerEmail string `json:"buyer_email"`

	/* 买家备注旗帜（与淘宝网上订单的买家备注旗帜对应，只有买家才能查看该字段）
	红、黄、绿、蓝、紫 分别对应 1、2、3、4、5 */
	BuyerFlag int64 `json:"buyer_flag"`

	/* 买家备注（与淘宝网上订单的买家备注对应，只有买家才能查看该字段） */
	BuyerMemo string `json:"buyer_memo"`

	/* 买家留言 */
	BuyerMessage string `json:"buyer_message"`

	/* 买家昵称 */
	BuyerNick string `json:"buyer_nick"`

	/* 买家获得积分,返点的积分。格式:100;单位:个。返点的积分要交易成功之后才能获得。 */
	BuyerObtainPointFee int64 `json:"buyer_obtain_point_fee"`

	/* 买家是否已评价。可选值:true(已评价),false(未评价)。如买家只评价未打分，此字段仍返回false */
	BuyerRate bool `json:"buyer_rate"`

	/* 买家可以通过此字段查询是否当前交易可以评论，can_rate=true可以评价，false则不能评价。 */
	CanRate bool `json:"can_rate"`

	/* 货到付款服务费。精确到2位小数;单位:元。如:12.07，表示:12元7分。 */
	CodFee string `json:"cod_fee"`

	/* 货到付款物流状态。
	初始状态 NEW_CREATED,
	接单成功 ACCEPTED_BY_COMPANY,
	接单失败 REJECTED_BY_COMPANY,
	接单超时 RECIEVE_TIMEOUT,
	揽收成功 TAKEN_IN_SUCCESS,
	揽收失败 TAKEN_IN_FAILED,
	揽收超时 TAKEN_TIMEOUT,
	签收成功 SIGN_IN,
	签收失败 REJECTED_BY_OTHER_SIDE,
	订单等待发送给物流公司 WAITING_TO_BE_SENT,
	用户取消物流订单 CANCELED */
	CodStatus string `json:"cod_status"`

	/* 交易佣金。精确到2位小数;单位:元。如:200.07，表示:200元7分 */
	CommissionFee string `json:"commission_fee"`

	/* 物流发货时效，单位小时 */
	ConsignInterval int64 `json:"consign_interval"`

	/* 卖家发货时间。格式:yyyy-MM-dd HH:mm:ss */
	ConsignTime string `json:"consign_time"`

	/* 交易创建时间。格式:yyyy-MM-dd HH:mm:ss */
	Created string `json:"created"`

	/* 使用信用卡支付金额数 */
	CreditCardFee string `json:"credit_card_fee"`

	/* 建议使用trade.promotion_details查询系统优惠
	系统优惠金额（如打折，VIP，满就送等），精确到2位小数，单位：元。如：200.07，表示：200元7分 */
	DiscountFee string `json:"discount_fee"`

	/* 交易结束时间。交易成功时间(更新交易状态为成功的同时更新)/确认收货时间或者交易关闭时间 。格式:yyyy-MM-dd HH:mm:ss */
	EndTime string `json:"end_time"`

	/* 电子凭证的垂直信息 */
	EticketExt string `json:"eticket_ext"`

	/* 快递代收款。精确到2位小数;单位:元。如:212.07，表示:212元7分 */
	ExpressAgencyFee string `json:"express_agency_fee"`

	/* 判断订单是否有买家留言，有买家留言返回true，否则返回false */
	HasBuyerMessage bool `json:"has_buyer_message"`

	/* 是否包含邮费。与available_confirm_fee同时使用。可选值:true(包含),false(不包含) */
	HasPostFee bool `json:"has_post_fee"`

	/* 订单中是否包含运费险订单，如果包含运费险订单返回true，不包含运费险订单，返回false */
	HasYfx bool `json:"has_yfx"`

	/* 商品字符串编号(注意：iid近期即将废弃，请用num_iid参数) */
	Iid string `json:"iid"`

	/* 发票抬头 */
	InvoiceName string `json:"invoice_name"`

	/* 发票类型 */
	InvoiceType string `json:"invoice_type"`

	/* 是否是3D淘宝交易 */
	Is3D bool `json:"is_3D"`

	/* 表示是否是品牌特卖（常规特卖，不包括特卖惠和特实惠）订单，如果是返回true，如果不是返回false。当此字段与is_force_wlb均为true时，订单强制物流宝发货。 */
	IsBrandSale bool `json:"is_brand_sale"`

	/* 表示订单交易是否含有对应的代销采购单。
	如果该订单中存在一个对应的代销采购单，那么该值为true；反之，该值为false。 */
	IsDaixiao bool `json:"is_daixiao"`

	/* 订单是否强制使用物流宝发货。当此字段与is_brand_sale均为true时，订单强制物流宝发货。此字段为false时，该订单根据流转规则设置可以使用物流宝或者常规方式发货 */
	IsForceWlb bool `json:"is_force_wlb"`

	/* 是否保障速递，如果为true，则为保障速递订单，使用线下联系发货接口发货，如果未false，则该订单非保障速递，根据卖家设置的订单流转规则可使用物流宝或者常规物流发货。 */
	IsLgtype bool `json:"is_lgtype"`

	/* 是否是多次发货的订单
	如果卖家对订单进行多次发货，则为true
	否则为false */
	IsPartConsign bool `json:"is_part_consign"`

	/* 表示订单交易是否网厅订单。 如果该订单是网厅订单，那么该值为true；反之，该值为false。 */
	IsWt bool `json:"is_wt"`

	/* 次日达订单送达时间 */
	LgAging string `json:"lg_aging"`

	/* 次日达，三日达等送达类型 */
	LgAgingType string `json:"lg_aging_type"`

	/* 订单出现异常问题的时候，给予用户的描述,没有异常的时候，此值为空 */
	MarkDesc string `json:"mark_desc"`

	/* 交易修改时间(用户对订单的任何修改都会更新此字段)。格式:yyyy-MM-dd HH:mm:ss */
	Modified string `json:"modified"`

	/* 商品购买数量。取值范围：大于零的整数,对于一个trade对应多个order的时候（一笔主订单，对应多笔子订单），num=0，num是一个跟商品关联的属性，一笔订单对应多比子订单的时候，主订单上的num无意义。 */
	Num int64 `json:"num"`

	/* 商品数字编号 */
	NumIid int64 `json:"num_iid"`

	/* 卡易售垂直表信息，去除下单ip之后的结果 */
	NutFeature string `json:"nut_feature"`

	/* 导购宝=crm */
	O2o string `json:"o2o"`

	/* 导购宝提货方式，inshop：店内提货，online：线上发货 */
	O2oDelivery string `json:"o2o_delivery"`

	/* 导购员id */
	O2oGuideId string `json:"o2o_guide_id"`

	/* 导购员名称 */
	O2oGuideName string `json:"o2o_guide_name"`

	/* 外部订单号 */
	O2oOutTradeId string `json:"o2o_out_trade_id"`

	/* 导购员门店id */
	O2oShopId string `json:"o2o_shop_id"`

	/* 导购门店名称 */
	O2oShopName string `json:"o2o_shop_name"`

	/* 订单列表 */
	Orders []*Order `json:"orders>order"`

	/* 付款时间。格式:yyyy-MM-dd HH:mm:ss。订单的付款时间即为物流订单的创建时间。 */
	PayTime string `json:"pay_time"`

	/* 实付金额。精确到2位小数;单位:元。如:200.07，表示:200元7分 */
	Payment string `json:"payment"`

	/* 天猫点券卡实付款金额,单位分 */
	PccAf int64 `json:"pcc_af"`

	/* 商品图片绝对途径 */
	PicPath string `json:"pic_path"`

	/* 买家使用积分,下单时生成，且一直不变。格式:100;单位:个. */
	PointFee int64 `json:"point_fee"`

	/* 邮费。精确到2位小数;单位:元。如:200.07，表示:200元7分 */
	PostFee string `json:"post_fee"`

	/* 商品价格。精确到2位小数；单位：元。如：200.07，表示：200元7分 */
	Price string `json:"price"`

	/* 交易促销详细信息 */
	Promotion string `json:"promotion"`

	/* 优惠详情 */
	PromotionDetails []*PromotionDetail `json:"promotion_details>promotion_detail"`

	/* 买家实际使用积分（扣除部分退款使用的积分），交易完成后生成（交易成功或关闭），交易未完成时该字段值为0。格式:100;单位:个 */
	RealPointFee int64 `json:"real_point_fee"`

	/* 卖家实际收到的支付宝打款金额（由于子订单可以部分确认收货，这个金额会随着子订单的确认收货而不断增加，交易成功后等于买家实付款减去退款金额）。精确到2位小数;单位:元。如:200.07，表示:200元7分 */
	ReceivedPayment string `json:"received_payment"`

	/* 收货人的详细地址 */
	ReceiverAddress string `json:"receiver_address"`

	/* 收货人的所在城市<br>
	注：因为国家对于城市和地区的划分的有：省直辖市和省直辖县级行政区（区级别的）划分的，淘宝这边根据这个差异保存在不同字段里面
	比如：广东广州：广州属于一个直辖市是放在的receiver_city的字段里面；而河南济源：济源属于省直辖县级行政区划分，是区级别的，放在了receiver_district里面
	<br>建议：程序依赖于城市字段做物流等判断的操作，最好加一个判断逻辑：如果返回值里面只有receiver_district参数，该参数作为城市 */
	ReceiverCity string `json:"receiver_city"`

	/* 收货人的所在地区<br>
	注：因为国家对于城市和地区的划分的有：省直辖市和省直辖县级行政区（区级别的）划分的，淘宝这边根据这个差异保存在不同字段里面
	比如：广东广州：广州属于一个直辖市是放在的receiver_city的字段里面；而河南济源：济源属于省直辖县级行政区划分，是区级别的，放在了receiver_district里面
	<br>建议：程序依赖于城市字段做物流等判断的操作，最好加一个判断逻辑：如果返回值里面只有receiver_district参数，该参数作为城市 */
	ReceiverDistrict string `json:"receiver_district"`

	/* 收货人的手机号码 */
	ReceiverMobile string `json:"receiver_mobile"`

	/* 收货人的姓名 */
	ReceiverName string `json:"receiver_name"`

	/* 收货人的电话号码 */
	ReceiverPhone string `json:"receiver_phone"`

	/* 收货人的所在省份 */
	ReceiverState string `json:"receiver_state"`

	/* 收货人的邮编 */
	ReceiverZip string `json:"receiver_zip"`

	/* 卖家支付宝账号 */
	SellerAlipayNo string `json:"seller_alipay_no"`

	/* 卖家是否可以对订单进行评价 */
	SellerCanRate bool `json:"seller_can_rate"`

	/* 卖家货到付款服务费。精确到2位小数;单位:元。如:12.07，表示:12元7分。卖家不承担服务费的订单：未发货的订单获取服务费为0，发货后就能获取到正确值。 */
	SellerCodFee string `json:"seller_cod_fee"`

	/* 卖家邮件地址 */
	SellerEmail string `json:"seller_email"`

	/* 卖家备注旗帜（与淘宝网上订单的卖家备注旗帜对应，只有卖家才能查看该字段）
	红、黄、绿、蓝、紫 分别对应 1、2、3、4、5 */
	SellerFlag int64 `json:"seller_flag"`

	/* 卖家备注（与淘宝网上订单的卖家备注对应，只有卖家才能查看该字段） */
	SellerMemo string `json:"seller_memo"`

	/* 卖家手机 */
	SellerMobile string `json:"seller_mobile"`

	/* 卖家姓名 */
	SellerName string `json:"seller_name"`

	/* 卖家昵称 */
	SellerNick string `json:"seller_nick"`

	/* 卖家电话 */
	SellerPhone string `json:"seller_phone"`

	/* 卖家是否已评价。可选值:true(已评价),false(未评价) */
	SellerRate bool `json:"seller_rate"`

	/* 订单将在此时间前发出，主要用于预售订单 */
	SendTime string `json:"send_time"`

	/* 服务子订单列表 */
	ServiceOrders []*ServiceOrder `json:"service_orders>service_order"`

	/* 物流标签 */
	ServiceTags []*LogisticsTag `json:"service_tags>logistics_tag"`

	/* 创建交易时的物流方式（交易完成前，物流方式有可能改变，但系统里的这个字段一直不变）。可选值：free(卖家包邮),post(平邮),express(快递),ems(EMS),virtual(虚拟发货)，25(次日必达)，26(预约配送)。 */
	ShippingType string `json:"shipping_type"`

	/* 交易快照详细信息 */
	Snapshot string `json:"snapshot"`

	/* 交易快照地址 */
	SnapshotUrl string `json:"snapshot_url"`

	/* 交易状态。可选值:
	 * TRADE_NO_CREATE_PAY(没有创建支付宝交易)
	 * WAIT_BUYER_PAY(等待买家付款)
	 * SELLER_CONSIGNED_PART(卖家部分发货)
	 * WAIT_SELLER_SEND_GOODS(等待卖家发货,即:买家已付款)
	 * WAIT_BUYER_CONFIRM_GOODS(等待买家确认收货,即:卖家已发货)
	 * TRADE_BUYER_SIGNED(买家已签收,货到付款专用)
	 * TRADE_FINISHED(交易成功)
	 * TRADE_CLOSED(付款以后用户退款成功，交易自动关闭)
	 * TRADE_CLOSED_BY_TAOBAO(付款以前，卖家或买家主动关闭交易)
	 * PAY_PENDING(国际信用卡支付付款确认中)
	 * WAIT_PRE_AUTH_CONFIRM(0元购合约中) */
	Status string `json:"status"`

	/* 分阶段付款的已付金额（万人团订单已付金额） */
	StepPaidFee string `json:"step_paid_fee"`

	/* 分阶段付款的订单状态（例如万人团订单等），目前有三返回状态
	FRONT_NOPAID_FINAL_NOPAID(定金未付尾款未付)，FRONT_PAID_FINAL_NOPAID(定金已付尾款未付)，FRONT_PAID_FINAL_PAID(定金和尾款都付) */
	StepTradeStatus string `json:"step_trade_status"`

	/* 交易编号 (父订单的交易编号) */
	Tid int64 `json:"tid"`

	/* 超时到期时间。格式:yyyy-MM-dd HH:mm:ss。业务规则：
	前提条件：只有在买家已付款，卖家已发货的情况下才有效
	如果申请了退款，那么超时会落在子订单上；比如说3笔ABC，A申请了，那么返回的是BC的列表, 主定单不存在
	如果没有申请过退款，那么超时会挂在主定单上；比如ABC，返回主定单，ABC的超时和主定单相同 */
	TimeoutActionTime string `json:"timeout_action_time"`

	/* 交易标题，以店铺名作为此标题的值。注:taobao.trades.get接口返回的Trade中的title是商品名称 */
	Title string `json:"title"`

	/* 商品金额（商品价格乘以数量的总金额）。精确到2位小数;单位:元。如:200.07，表示:200元7分 */
	TotalFee string `json:"total_fee"`

	/* 交易内部来源。
	WAP(手机);HITAO(嗨淘);TOP(TOP平台);TAOBAO(普通淘宝);JHS(聚划算)
	一笔订单可能同时有以上多个标记，则以逗号分隔 */
	TradeFrom string `json:"trade_from"`

	/* 交易备注。 */
	TradeMemo string `json:"trade_memo"`

	/* 交易外部来源：ownshop(商家官网) */
	TradeSource string `json:"trade_source"`

	/* 交易类型列表，同时查询多种交易类型可用逗号分隔。默认同时查询guarantee_trade, auto_delivery, ec, cod的4种交易类型的数据
	可选值
	fixed(一口价)
	auction(拍卖)
	guarantee_trade(一口价、拍卖)
	auto_delivery(自动发货)
	independent_simple_trade(旺店入门版交易)
	independent_shop_trade(旺店标准版交易)
	ec(直冲)
	cod(货到付款)
	fenxiao(分销)
	game_equipment(游戏装备)
	shopex_trade(ShopEX交易)
	netcn_trade(万网交易)
	external_trade(统一外部交易)
	o2o_offlinetrade（O2O交易）
	step (万人团)
	nopaid(无付款订单)
	pre_auth_type(预授权0元购机交易) */
	Type string `json:"type"`

	/* 订单的运费险，单位为元 */
	YfxFee string `json:"yfx_fee"`

	/* 运费险支付号 */
	YfxId string `json:"yfx_id"`

	/* 运费险类型 */
	YfxType string `json:"yfx_type"`

	/* 在返回的trade对象上专门增加一个字段zero_purchase来区分，这个为true的就是0元购机预授权交易 */
	ZeroPurchase bool `json:"zero_purchase"`
}

/* 运费模板中运费信息对象 */
type TopFee struct {
	/* 增费：输入0.00-999.99（最多包含两位小数） 不能为空不输入运费时必须输入0 */
	AddFee string `json:"add_fee"`

	/* 增费标准：当valuation(记价方式)为0时输入1-9999范围内的整数 */
	AddStandard string `json:"add_standard"`

	/* 邮费子项涉及的地区,多个地区用逗号连接数量串;可以用taobao.areas.get接口获取.或者参考:http://www.stats.gov.cn/tjbz/xzqhdm/t20080215_402462675.htm 例 (110000,310000,320000,330000).就代表邮费子项涉 及(北京,上海,江苏,浙江)四个地区.填写时要注意对照地区代码值,如果填入错误地区代码,将会出现错误信息. */
	Destination string `json:"destination"`

	/* 可选值：post:平邮; cod:货到付款; ems:EMS; express:快递公司 */
	ServiceType string `json:"service_type"`

	/* 首费：输入0.01-999.99（最多包含两位小数） 不能为空也不能为0 */
	StartFee string `json:"start_fee"`

	/* 首费标准：当valuation(记价方式)为0时输入1-9999范围内的整数 */
	StartStandard string `json:"start_standard"`
}

/* Sku结构 */
type Sku struct {
	/* 商品级别的条形码 */
	Barcode string `json:"barcode"`

	/* 基础色数据 */
	ChangeProp string `json:"change_prop"`

	/* sku创建日期 时间格式：yyyy-MM-dd HH:mm:ss */
	Created string `json:"created"`

	/* sku所属商品id(注意：iid近期即将废弃，请用num_iid参数) */
	Iid string `json:"iid"`

	/* sku最后修改日期 时间格式：yyyy-MM-dd HH:mm:ss */
	Modified string `json:"modified"`

	/* sku所属商品数字id */
	NumIid int64 `json:"num_iid"`

	/* 商家设置的外部id。天猫和集市的卖家，需要登录才能获取到自己的商家编码，不能获取到他人的商家编码。 */
	OuterId string `json:"outer_id"`

	/* 属于这个sku的商品的价格 取值范围:0-100000000;精确到2位小数;单位:元。如:200.07，表示:200元7分。 */
	Price string `json:"price"`

	/* sku的销售属性组合字符串（颜色，大小，等等，可通过类目API获取某类目下的销售属性）,格式是p1:v1;p2:v2 */
	Properties string `json:"properties"`

	/* sku所对应的销售属性的中文名字串，格式如：pid1:vid1:pid_name1:vid_name1;pid2:vid2:pid_name2:vid_name2…… */
	PropertiesName string `json:"properties_name"`

	/* 属于这个sku的商品的数量， */
	Quantity int64 `json:"quantity"`

	/* sku级别发货时间 */
	SkuDeliveryTime string `json:"sku_delivery_time"`

	/* sku的id */
	SkuId int64 `json:"sku_id"`

	/* 表示SKu上的产品规格信息 */
	SkuSpecId int64 `json:"sku_spec_id"`

	/* sku状态。 normal:正常 ；delete:删除 */
	Status string `json:"status"`

	/* 商品在付款减库存的状态下，该sku上未付款的订单数量 */
	WithHoldQuantity int64 `json:"with_hold_quantity"`
}

/* 商品视频关联记录 */
type Video struct {
	/* 视频关联记录创建时间（格式：yyyy-MM-dd HH:mm:ss） */
	Created string `json:"created"`

	/* 视频关联记录的id，和商品相对应 */
	Id int64 `json:"id"`

	/* 视频记录关联的商品的数字id(注意：iid近期即将废弃，请用num_iid参数) */
	Iid string `json:"iid"`

	/* 视频关联记录修改时间（格式：yyyy-MM-dd HH:mm:ss） */
	Modified string `json:"modified"`

	/* 视频记录所关联的商品的数字id */
	NumIid int64 `json:"num_iid"`

	/* video的url连接地址。淘秀里视频记录里面存储的url地址 */
	Url string `json:"url"`

	/* video的id，对应于视频在淘秀的存储记录 */
	VideoId int64 `json:"video_id"`
}

/* 店铺动态评分信息 */
type ShopScore struct {
	/* 发货速度评分 */
	DeliveryScore string `json:"delivery_score"`

	/* 商品描述评分 */
	ItemScore string `json:"item_score"`

	/* 服务态度评分 */
	ServiceScore string `json:"service_score"`
}

/* 店铺内卖家自定义类目 */
type SellerCat struct {
	/* 卖家自定义类目编号 */
	Cid int64 `json:"cid"`

	/* 创建时间。格式：yyyy-MM-dd HH:mm:ss */
	Created string `json:"created"`

	/* 修改时间。格式：yyyy-MM-dd HH:mm:ss */
	Modified string `json:"modified"`

	/* 卖家自定义类目名称 */
	Name string `json:"name"`

	/* 父类目编号，值等于0：表示此类目为店铺下的一级类目，值不等于0：表示此类目有父类目 */
	ParentCid int64 `json:"parent_cid"`

	/* 链接图片地址 */
	PicUrl string `json:"pic_url"`

	/* 该类目在页面上的排序位置 */
	SortOrder int64 `json:"sort_order"`

	/* 店铺类目类型：可选值：manual_type：手动分类，new_type：新品上价， tree_type：二三级类目树 ，property_type：属性叶子类目树， brand_type：品牌推广 */
	Type string `json:"type"`
}

/* 采购单子单退款详情 */
type RefundDetail struct {
	/* 下游买家的退款信息 */
	BuyerRefund *BuyerRefund `json:"buyer_refund"`

	/* 分销商nick */
	DistributorNick string `json:"distributor_nick"`

	/* 是否退货 */
	IsReturnGoods bool `json:"is_return_goods"`

	/* 退款修改时间。格式:yyyy-MM-dd HH:mm:ss */
	Modified string `json:"modified"`

	/* 支付给供应商的金额 */
	PaySupFee float64 `json:"pay_sup_fee,string"`

	/* 主采购单id */
	PurchaseOrderId int64 `json:"purchase_order_id"`

	/* 退款创建时间 */
	RefundCreateTime string `json:"refund_create_time"`

	/* 退款说明 */
	RefundDesc string `json:"refund_desc"`

	/* 退款的金额 */
	RefundFee float64 `json:"refund_fee,string"`

	/* 退款流程类型：
	4：发货前退款；
	1：发货后退款不退货；
	2：发货后退款退货 */
	RefundFlowType int64 `json:"refund_flow_type"`

	/* 退款原因 */
	RefundReason string `json:"refund_reason"`

	/* 退款状态
	1：买家已经申请退款，等待卖家同意
	2：卖家已经同意退款，等待买家退货
	3：买家已经退货，等待卖家确认收货
	4：退款关闭
	5：退款成功
	6：卖家拒绝退款
	12：同意退款，待打款
	9：没有申请退款
	10：卖家拒绝确认收货 */
	RefundStatus int64 `json:"refund_status"`

	/* 子单id */
	SubOrderId int64 `json:"sub_order_id"`

	/* 供应商nick */
	SupplierNick string `json:"supplier_nick"`

	/* 超时时间 */
	Timeout string `json:"timeout"`

	/* 超时类型：
	1：供应商同意退款/同意退货超时；
	2：供应商确认收货超时 */
	ToType int64 `json:"to_type"`
}

/* 登录日志 */
type LoginLog struct {
	/* 用户登录或退出客户端的时间 */
	Time string `json:"time"`

	/* 标志用户登录或退出。
	0表示登陆，1表示退出。 */
	Type string `json:"type"`
}

/* 产品的各种图片信息 */
type FenxiaoImage struct {
	/* 图片id */
	ImageId int64 `json:"image_id"`

	/* 图片顺序 */
	ImagePosition int64 `json:"image_position"`

	/* 图片对应的url */
	ImageUrl string `json:"image_url"`

	/* 当图片类型为属性图片时，表示图片对应的属性pv对。 */
	Properties string `json:"properties"`

	/* 图片类型（PRODUCT：产品图片  EXTPRODUCT：产品辅图  PROPERTIES：产品属性图片 ） */
	Type string `json:"type"`
}

/* 评价详细 */
type EvalDetail struct {
	/* 评价详细：
	1 非常满意
	2 满意
	3 一般
	4 不满意 */
	EvalCode int64 `json:"eval_code"`

	/* 评价接收者 */
	EvalRecer string `json:"eval_recer"`

	/* 评价发送者 */
	EvalSender string `json:"eval_sender"`

	/* 评价时间 */
	EvalTime string `json:"eval_time"`

	/* 评价发送时间 */
	SendTime string `json:"send_time"`
}

/* 分流权重 */
type StreamWeight struct {
	/* 账号 */
	User string `json:"user"`

	/* 账号对应的权重 */
	Weight int64 `json:"weight"`
}

/* 评价大印象返回的印象词接口 */
type ImprItemDO struct {
	/* 印象标签的id */
	AttributeId int64 `json:"attribute_id"`

	/* 被买家评价提及到的印象标签次数 */
	Count int64 `json:"count"`

	/* 印象标签情感倾向 1表示正面 -1表示负面 0表示中性 */
	Status int64 `json:"status"`

	/* 印象标签的名称 */
	Title string `json:"title"`
}

/* 采购单留言列表 */
type OrderMessage struct {
	/* 留言内容 */
	MessageContent string `json:"message_content"`

	/* 留言时间 */
	MessageTime string `json:"message_time"`

	/* 留言标题，例如：分销商留言，供应商留言，买家留言 */
	MessageTitle string `json:"message_title"`

	/* 留言时的图片地址 */
	PicUrl string `json:"pic_url"`
}

/* 下游买家退款信息 */
type BuyerRefund struct {
	/* 订单id */
	BizOrderId int64 `json:"biz_order_id"`

	/* 下游买家nick */
	BuyerNick string `json:"buyer_nick"`

	/* 货物的状态：
	买家已收到货
	买家已退货
	买家未收到货 */
	GoodsStatusDesc string `json:"goods_status_desc"`

	/* 退款修改时间。格式:yyyy-MM-dd HH:mm:ss */
	Modified string `json:"modified"`

	/* 买家是否退货 */
	NeedReturnGoods bool `json:"need_return_goods"`

	/* 退款创建时间 */
	RefundCreateTime string `json:"refund_create_time"`

	/* 退款说明 */
	RefundDesc string `json:"refund_desc"`

	/* 交易退款id */
	RefundId int64 `json:"refund_id"`

	/* 退款原因 */
	RefundReason string `json:"refund_reason"`

	/* 退款状态 */
	RefundStatus int64 `json:"refund_status"`

	/* 退还买家的金额 */
	ReturnFee int64 `json:"return_fee"`

	/* 采购单子单id */
	SubOrderId int64 `json:"sub_order_id"`

	/* 支付分销商的金额 */
	ToSellerFee int64 `json:"to_seller_fee"`
}

/* 确认收货费用结构 */
type TradeConfirmFee struct {
	/* 确认收货的金额(不包含邮费)。精确到2位小数;单位:元。如:200.07，表示:200元7分 */
	ConfirmFee float64 `json:"confirm_fee,string"`

	/* 需确认收货的邮费，当不是最后一笔收货或者没有邮费时是0.00。精确到2位小数;单位:元。如:200.07，表示:200元7分 */
	ConfirmPostFee float64 `json:"confirm_post_fee,string"`

	/* 是否是最后一笔订单（只对子订单有效，当其他子订单都是交易完成时，返回true，否则false） */
	IsLastOrder bool `json:"is_last_order"`
}

/* ItemImg结构 */
type ItemImg struct {
	/* 图片创建时间 时间格式：yyyy-MM-dd HH:mm:ss */
	Created string `json:"created"`

	/* 商品图片的id，和商品相对应（主图id默认为0） */
	Id int64 `json:"id"`

	/* 图片放在第几张（多图时可设置） */
	Position int64 `json:"position"`

	/* 图片链接地址 */
	Url string `json:"url"`
}

/* 店铺类目 */
type ShopCat struct {
	/* 类目编号 */
	Cid int64 `json:"cid"`

	/* 该类目是否为父类目。即：该类目是否还有子类目 */
	IsParent bool `json:"is_parent"`

	/* 类目名称 */
	Name string `json:"name"`

	/* 父类目编号，注：此类目指前台类目，值等于0：表示此类目为一级类目，值不等于0：表示此类目有父类目 */
	ParentCid int64 `json:"parent_cid"`
}

/* 物流公司资费相关信息 */
type CarriageDetail struct {
	/* 续费（单位：元） */
	AddFee int64 `json:"add_fee"`

	/* 续重（单位：千克） */
	AddWeight int64 `json:"add_weight"`

	/* 破损赔付 */
	DamagePayment string `json:"damage_payment"`

	/* 物流公司揽收时间段 */
	GotTime string `json:"got_time"`

	/* 首费（单位：元） */
	InitialFee int64 `json:"initial_fee"`

	/* 首重（单位：千克） */
	InitialWeight int64 `json:"initial_weight"`

	/* 丢单赔付 */
	LostPayment string `json:"lost_payment"`

	/* 快件送达所需的时间(单位：天) */
	WayDay string `json:"way_day"`
}

/* 产品属性图片 */
type ProductPropImg struct {
	/* 添加时间.格式:yyyy-mm-dd hh:mm:ss */
	Created string `json:"created"`

	/* 产品属性图片ID */
	Id int64 `json:"id"`

	/* 修改时间.格式:yyyy-mm-dd hh:mm:ss */
	Modified string `json:"modified"`

	/* 图片序号。产品里的图片展示顺序，数据越小越靠前。要求是正整数。 */
	Position int64 `json:"position"`

	/* 图片所属产品的ID */
	ProductId int64 `json:"product_id"`

	/* 属性串(pid:vid),目前只有颜色属性.如:颜色:红色表示为　1627207:28326 */
	Props string `json:"props"`

	/* 图片地址.(绝对地址,格式:http://host/image_path) */
	Url string `json:"url"`
}

/* 店铺信息 */
type Shop struct {
	/* 总橱窗数量，对于C卖家返回总橱窗数，对于B卖家返回0（只有taobao.shop.remainshowcase.get可以返回） */
	AllCount int64 `json:"all_count"`

	/* 店铺公告 */
	Bulletin string `json:"bulletin"`

	/* 店铺所属的类目编号 */
	Cid int64 `json:"cid"`

	/* 开店时间。格式：yyyy-MM-dd HH:mm:ss */
	Created string `json:"created"`

	/* 店铺描述 */
	Desc string `json:"desc"`

	/* 最后修改时间。格式：yyyy-MM-dd HH:mm:ss */
	Modified string `json:"modified"`

	/* 卖家昵称 */
	Nick string `json:"nick"`

	/* 店标地址。返回相对路径，可以用"http://logo.taobao.com/shop-logo"来拼接成绝对路径 */
	PicPath string `json:"pic_path"`

	/* 剩余橱窗数量，对于C卖家返回剩余橱窗数，对于B卖家返回-1（只有taobao.shop.remainshowcase.get可以返回） */
	RemainCount int64 `json:"remain_count"`

	/* 店铺动态评分信息 */
	ShopScore *ShopScore `json:"shop_score"`

	/* 店铺编号。shop+sid.taobao.com即店铺地址，如shop123456.taobao.com */
	Sid int64 `json:"sid"`

	/* 店铺标题 */
	Title string `json:"title"`

	/* 已用的橱窗数量，对于C卖家返回已使用的橱窗数，对于B卖家返回-1（只有taobao.shop.remainshowcase.get可以返回） */
	UsedCount int64 `json:"used_count"`
}

/* E客服账号操作信息 */
type AccountInfo struct {
	/* 账号 */
	Account string `json:"account"`

	/* 操作列表 */
	AccountStats []*AccountStat `json:"account_stats>account_stat"`

	/* account_stats列表的长度 */
	Count int64 `json:"count"`
}

/* 商城虚拟服务子订单数据结构 */
type ServiceOrder struct {
	/* 卖家昵称 */
	BuyerNick string `json:"buyer_nick"`

	/* 服务所属的交易订单号。如果服务为一年包换，则item_oid这笔订单享受改服务的保护 */
	ItemOid int64 `json:"item_oid"`

	/* 购买数量，取值范围为大于0的整数 */
	Num int64 `json:"num"`

	/* 虚拟服务子订单订单号 */
	Oid int64 `json:"oid"`

	/* 子订单实付金额。精确到2位小数，单位:元。如:200.07，表示:200元7分。 */
	Payment string `json:"payment"`

	/* 服务图片地址 */
	PicPath string `json:"pic_path"`

	/* 服务价格，精确到小数点后两位：单位:元 */
	Price string `json:"price"`

	/* 最近退款的id */
	RefundId int64 `json:"refund_id"`

	/* 卖家昵称 */
	SellerNick string `json:"seller_nick"`

	/* 服务详情的URL地址 */
	ServiceDetailUrl string `json:"service_detail_url"`

	/* 服务数字id */
	ServiceId int64 `json:"service_id"`

	/* 商品名称 */
	Title string `json:"title"`

	/* 服务子订单总费用 */
	TotalFee string `json:"total_fee"`
}

/* 商品属性图片结构 */
type PropImg struct {
	/* 图片创建时间 时间格式：yyyy-MM-dd HH:mm:ss */
	Created string `json:"created"`

	/* 属性图片的id，和商品相对应 */
	Id int64 `json:"id"`

	/* 图片放在第几张（多图时可设置） */
	Position int64 `json:"position"`

	/* 图片所对应的属性组合的字符串 */
	Properties string `json:"properties"`

	/* 图片链接地址 */
	Url string `json:"url"`
}

/* 用户订购信息 */
type UserSubscribe struct {
	/* 订购结束时间。格式:yyyy-MM-dd HH:mm:ss */
	EndDate string `json:"end_date"`

	/* 订购开始时间。格式:yyyy-MM-dd HH:mm:ss */
	StartDate string `json:"start_date"`

	/* 订购状况。应用订购者：subscribeUser;尚未订购：unsubscribeUser；非法用户：invalidateUser */
	Status string `json:"status"`

	/* 0-无版本信息；1-初级版；2-中级版；3-高级版 */
	VersionNo int64 `json:"version_no"`
}

/* 品牌 */
type Brand struct {
	/* vid的值 */
	Name string `json:"name"`

	/* 品牌的属性id */
	Pid int64 `json:"pid"`

	/* 品牌的属性名 */
	PropName string `json:"prop_name"`

	/* 对应属性的 pid:vid 串中的vid */
	Vid int64 `json:"vid"`
}

/* 商品类目结构 */
type ItemCat struct {
	/* 商品所属类目ID */
	Cid int64 `json:"cid"`

	/* Feature对象列表
	目前已有的属性：
	若Attr_key为 udsaleprop，attr_value为1 则允许卖家在改类目新增自定义销售属性,不然为不允许 */
	Features []*Feature `json:"features>feature"`

	/* 该类目是否为父类目(即：该类目是否还有子类目) */
	IsParent bool `json:"is_parent"`

	/* 类目名称 */
	Name string `json:"name"`

	/* 父类目ID=0时，代表的是一级的类目 */
	ParentCid int64 `json:"parent_cid"`

	/* 排列序号，表示同级类目的展现次序，如数值相等则按名称次序排列。取值范围:大于零的整数 */
	SortOrder int64 `json:"sort_order"`

	/* 状态。可选值:normal(正常),deleted(删除) */
	Status string `json:"status"`
}

/* 属性值 */
type PropValue struct {
	/* 类目ID */
	Cid int64 `json:"cid"`

	/* 属性值feature */
	Features []*Feature `json:"features>feature"`

	/* 是否为父类目属性 */
	IsParent bool `json:"is_parent"`

	/* 修改时间（类目增量专用） */
	ModifiedTime string `json:"modified_time"`

	/* 三种枚举类型：modify，add，delete (增量类目专用) */
	ModifiedType string `json:"modified_type"`

	/* 属性值 */
	Name string `json:"name"`

	/* 属性值别名 */
	NameAlias string `json:"name_alias"`

	/* 属性 ID */
	Pid int64 `json:"pid"`

	/* 属性名 */
	PropName string `json:"prop_name"`

	/* 排列序号。取值范围:大于零的整数 */
	SortOrder int64 `json:"sort_order"`

	/* 状态。可选值:normal(正常),deleted(删除) */
	Status string `json:"status"`

	/* 属性值ID */
	Vid int64 `json:"vid"`
}

/* E客服账号操作记录 */
type AccountStat struct {
	/* 登录状态。0：离线；1：在线 */
	LoginStatus int64 `json:"login_status"`

	/* 服务状态。0：挂起；1：服务 */
	OnserviceStatus int64 `json:"onservice_status"`

	/* 操作时间。格式：YYYY-mm-dd HH:MM:SS */
	Time string `json:"time"`
}

/* Item(商品)结构 */
type Item struct {
	/* 售后服务ID,该字段仅在taobao.item.get接口中返回 */
	AfterSaleId int64 `json:"after_sale_id"`

	/* 商品上传后的状态。onsale出售中，instock库中 */
	ApproveStatus string `json:"approve_status"`

	/* 天猫订单抽佣比例，为5的倍数，最低0.5%。跟淘客佣金没有关系。 */
	AuctionPoint int64 `json:"auction_point"`

	/* 代充商品类型。在代充商品的类目下，不传表示不标记商品类型（交易搜索中就不能通过标记搜到相关的交易了）。可选类型： no_mark(不做类型标记) time_card(点卡软件代充) fee_card(话费软件代充) */
	AutoFill string `json:"auto_fill"`

	/* 商品级别的条形码 */
	Barcode string `json:"barcode"`

	/* 基础色数据 */
	ChangeProp string `json:"change_prop"`

	/* 商品所属的叶子类目 id */
	Cid int64 `json:"cid"`

	/* 货到付款运费模板ID */
	CodPostageId int64 `json:"cod_postage_id"`

	/* Item的发布时间，目前仅供taobao.item.add和taobao.item.get可用 */
	Created string `json:"created"`

	/* 定制工具Id */
	CustomMadeTypeId string `json:"custom_made_type_id"`

	/* 下架时间（格式：yyyy-MM-dd HH:mm:ss） */
	DelistTime string `json:"delist_time"`

	/* 发货时间信息 */
	DeliveryTime *DeliveryTime `json:"delivery_time"`

	/* 商品描述, 字数要大于5个字符，小于25000个字符 */
	Desc string `json:"desc"`

	/* 宝贝描述规范化模块锚点信息 */
	DescModuleInfo *DescModuleInfo `json:"desc_module_info"`

	/* 商品描述模块化，模块列表，由List转化成jsonArray存入，后端逻辑验证通过，拼装成模块内容+锚点导航后存入desc中。数据结构具体参见Item_Desc_Module */
	DescModules string `json:"desc_modules"`

	/* 商品url */
	DetailUrl string `json:"detail_url"`

	/* ems费用,格式：5.00；单位：元；精确到：分 */
	EmsFee float64 `json:"ems_fee,string"`

	/* 快递费用,格式：5.00；单位：元；精确到：分 */
	ExpressFee float64 `json:"express_fee,string"`

	/* 宝贝特征值，
	只有在Top支持的特征值才能保存到宝贝上 */
	Features string `json:"features"`

	/* 食品安全信息，包括：生产许可证编号、产品标准号、厂名、厂址等 */
	FoodSecurity *FoodSecurity `json:"food_security"`

	/* 运费承担方式,seller（卖家承担），buyer(买家承担） */
	FreightPayer string `json:"freight_payer"`

	/* 全球购商品采购地信息（地区/国家），代表全球购商品的产地信息。 */
	GlobalStockCountry string `json:"global_stock_country"`

	/* 全球购商品采购地信息（库存类型），有两种库存类型：现货和代购;
	参数值为1时代表现货，值为2时代表代购 */
	GlobalStockType string `json:"global_stock_type"`

	/* 支持会员打折,true/false */
	HasDiscount bool `json:"has_discount"`

	/* 是否有发票,true/false */
	HasInvoice bool `json:"has_invoice"`

	/* 橱窗推荐,true/false */
	HasShowcase bool `json:"has_showcase"`

	/* 是否有保修,true/false */
	HasWarranty bool `json:"has_warranty"`

	/* 加价幅度。如果为0，代表系统代理幅度。
	在竞拍中，为了超越上一个出价，会员需要在当前出价上增加金额，这个金额就是加价幅度。卖家在发布宝贝的时候可以自定义加价幅度，也可以让系统自动代理加价。系统自动代理加价的加价幅度随着当前出价金额的增加而增加，我们建议会员使用系统自动代理加价，并请买家在出价前看清楚加价幅度的具体金额。另外需要注意是，此功能只适用于拍卖的商品。
	以下是系统自动代理加价幅度表：
	当前价（加价幅度 ）
	1-40（ 1 ）、41-100（ 2 ）、101-200（5 ）、201-500 （10）、501-1001（15）、001-2000（25）、2001-5000（50）、5001-10000（100）
	10001以上         200 */
	Increment string `json:"increment"`

	/* 用户内店宝贝装修模板id */
	InnerShopAuctionTemplateId int64 `json:"inner_shop_auction_template_id"`

	/* 用户自行输入的类目属性ID串。结构："pid1,pid2,pid3"，如："20000"（表示品牌） 注：通常一个类目下用户可输入的关键属性不超过1个。 */
	InputPids string `json:"input_pids"`

	/* 用户自行输入的子属性名和属性值，结构:"父属性值;一级子属性名;一级子属性值;二级子属性名;自定义输入值,....",如：“耐克;耐克系列;科比系列;科比系列;2K5”，input_str需要与input_pids一一对应，注：通常一个类目下用户可输入的关键属性不超过1个。所有属性别名加起来不能超过 3999字节。 */
	InputStr string `json:"input_str"`

	/* 是否是3D淘宝的商品 */
	Is3D bool `json:"is_3D"`

	/* 是否为达尔文挂接成功了的商品 */
	IsCspu bool `json:"is_cspu"`

	/* 是否在外部网店显示 */
	IsEx bool `json:"is_ex"`

	/* 非分销商品：0，代销：1，经销：2 */
	IsFenxiao int64 `json:"is_fenxiao"`

	/* 是否24小时闪电发货 */
	IsLightningConsignment bool `json:"is_lightning_consignment"`

	/* 是否是线下商品。
	1：线上商品（默认值）；
	2：线上或线下商品；
	3：线下商品。 */
	IsOffline string `json:"is_offline"`

	/* 商品是否为先行赔付
	taobao.items.search和taobao.items.vip.search专用 */
	IsPrepay bool `json:"is_prepay"`

	/* 是否在淘宝显示 */
	IsTaobao bool `json:"is_taobao"`

	/* 是否定时上架商品 */
	IsTiming bool `json:"is_timing"`

	/* 虚拟商品的状态字段 */
	IsVirtual bool `json:"is_virtual"`

	/* 标示商品是否为新品。
	值含义：true-是，false-否。 */
	IsXinpin bool `json:"is_xinpin"`

	/* 商品图片列表(包括主图)。fields中只设置item_img可以返回ItemImg结构体中所有字段，如果设置为item_img.id、item_img.url、item_img.position等形式就只会返回相应的字段 */
	ItemImgs []*ItemImg `json:"item_imgs>item_img"`

	/* 表示商品的体积，用于按体积计费的运费模板。该值的单位为立方米（m3）。
	该值支持两种格式的设置：格式1：bulk:3,单位为立方米(m3),表示直接设置为商品的体积。格式2：weight:10;breadth:10;height:10，单位为米（m） */
	ItemSize string `json:"item_size"`

	/* 商品的重量，用于按重量计费的运费模板。注意：单位为kg */
	ItemWeight string `json:"item_weight"`

	/* 上架时间（格式：yyyy-MM-dd HH:mm:ss） */
	ListTime string `json:"list_time"`

	/* 本地生活电子交易凭证业务，目前此字段只涉及到的信息为有效期:
	如果有效期为起止日期类型，此值为2012-08-06,2012-08-16
	如果有效期为【购买成功日 至】类型则格式为2012-08-16
	如果有效期为天数类型则格式为3 */
	LocalityLife *LocalityLife `json:"locality_life"`

	/* 商品所在地 */
	Location *Location `json:"location"`

	/* 商品修改时间（格式：yyyy-MM-dd HH:mm:ss） */
	Modified string `json:"modified"`

	/* 宝贝主图视频的数据信息，包括：视频ID，视频缩略图URL，视频时长，视频状态等信息。 */
	MpicVideo *MpicVideo `json:"mpic_video"`

	/* 是否为新消保法中的7天无理由退货 */
	Newprepay string `json:"newprepay"`

	/* 卖家昵称 */
	Nick string `json:"nick"`

	/* 商品数量 */
	Num int64 `json:"num"`

	/* 商品数字id */
	NumIid int64 `json:"num_iid"`

	/* 是否淘1站商品 */
	OneStation bool `json:"one_station"`

	/* open_iid */
	OpenIid string `json:"open_iid"`

	/* 商家外部编码(可与商家外部系统对接)。需要授权才能获取。 */
	OuterId string `json:"outer_id"`

	/* 用户外店装修模板id */
	OuterShopAuctionTemplateId int64 `json:"outer_shop_auction_template_id"`

	/* 用于保存拍卖有关的信息 */
	PaimaiInfo *PaimaiInfo `json:"paimai_info"`

	/* 商品主图片地址 */
	PicUrl string `json:"pic_url"`

	/* 平邮费用,格式：5.00；单位：元；精确到：分 */
	PostFee float64 `json:"post_fee,string"`

	/* 宝贝所属的运费模板ID，如果没有返回则说明没有使用运费模板 */
	PostageId int64 `json:"postage_id"`

	/* 商品价格，格式：5.00；单位：元；精确到：分 */
	Price float64 `json:"price,string"`

	/* 宝贝所属产品的id(可能为空). 该字段可以通过taobao.products.search 得到 */
	ProductId int64 `json:"product_id"`

	/* 消保类型，多个类型以,分割。可取以下值：
	2：假一赔三；4：7天无理由退换货；taobao.items.search和taobao.items.vip.search专用 */
	PromotedService string `json:"promoted_service"`

	/* 商品属性图片列表。fields中只设置prop_img可以返回PropImg结构体中所有字段，如果设置为prop_img.id、prop_img.url、prop_img.properties、prop_img.position等形式就只会返回相应的字段 */
	PropImgs []*PropImg `json:"prop_imgs>prop_img"`

	/* 属性值别名,比如颜色的自定义名称 */
	PropertyAlias string `json:"property_alias"`

	/* 商品属性 格式：pid:vid;pid:vid */
	Props string `json:"props"`

	/* 商品属性名称。标识着props内容里面的pid和vid所对应的名称。格式为：pid1:vid1:pid_name1:vid_name1;pid2:vid2:pid_name2:vid_name2……(<strong>注：</strong><font color="red">属性名称中的冒号":"被转换为："#cln#";
	分号";"被转换为："#scln#"
	</font>) */
	PropsName string `json:"props_name"`

	/* 商品资质的信息，用URLEncoder做过转换，使用时，需要URLDecoder转换回来，默认字符集为：UTF-8 */
	Qualification string `json:"qualification"`

	/* 商品所属卖家的信用等级数，1表示1心，2表示2心……，只有调用商品搜索:taobao.items.get和taobao.items.search的时候才能返回 */
	Score int64 `json:"score"`

	/* 秒杀商品类型。打上秒杀标记的商品，用户只能下架并不能再上架，其他任何编辑或删除操作都不能进行。如果用户想取消秒杀标记，需要联系小二进行操作。如果秒杀结束需要自由编辑请联系活动负责人（小二）去掉秒杀标记。可选类型
	web_only(只能通过web网络秒杀)
	wap_only(只能通过wap网络秒杀)
	web_and_wap(既能通过web秒杀也能通过wap秒杀) */
	SecondKill string `json:"second_kill"`

	/* 达尔文数据挂接，第二步保存结果 */
	SecondResult bool `json:"second_result"`

	/* 商品卖点信息，天猫商家使用字段，最长150个字符。 */
	SellPoint string `json:"sell_point"`

	/* 是否承诺退换货服务! */
	SellPromise bool `json:"sell_promise"`

	/* 商品所属的店铺内卖家自定义类目列表 */
	SellerCids string `json:"seller_cids"`

	/* 店铺类型，B：天猫店铺，C：淘宝店铺 */
	ShopType string `json:"shop_type"`

	/* <a href="http://open.taobao.com/dev/index.php/Sku">Sku</a>列表。fields中只设置sku可以返回Sku结构体中所有字段，如果设置为sku.sku_id、sku.properties、sku.quantity等形式就只会返回相应的字段 */
	Skus []*Sku `json:"skus>sku"`

	/* 商品销量 */
	SoldQuantity int64 `json:"sold_quantity"`

	/* 商品新旧程度(全新:new，闲置:unused，二手：second) */
	StuffStatus string `json:"stuff_status"`

	/* 商品是否支持拍下减库存:1支持;2取消支持(付款减库存);0(默认)不更改 集市卖家默认拍下减库存; 商城卖家默认付款减库存 */
	SubStock int64 `json:"sub_stock"`

	/* 商品的子标题，给商品增加卖点等描述 */
	SubTitle string `json:"sub_title"`

	/* 页面模板id */
	TemplateId string `json:"template_id"`

	/* 商品标题,不能超过60字节 */
	Title string `json:"title"`

	/* 商品类型(fixed:一口价;auction:拍卖)注：取消团购 */
	Type string `json:"type"`

	/* 有效期,7或者14（默认是7天） */
	ValidThru int64 `json:"valid_thru"`

	/* 该字段废弃，请勿使用。 */
	VideoId int64 `json:"video_id"`

	/* 商品视频列表(目前只支持单个视频关联)。fields中只设置video可以返回Video结构体中所有字段，如果设置为video.id、video.video_id、
	video.url等形式就只会返回相应的字段 */
	Videos []*Video `json:"videos>video"`

	/* 商品是否违规，违规：true , 不违规：false */
	Violation bool `json:"violation"`

	/* 不带html标签的desc文本信息，该字段只在taobao.item.get接口中返回 */
	WapDesc string `json:"wap_desc"`

	/* 适合wap应用的商品详情url ，该字段只在taobao.item.get接口中返回 */
	WapDetailUrl string `json:"wap_detail_url"`

	/* 无线的宝贝描述 */
	WirelessDesc string `json:"wireless_desc"`

	/* 预扣库存，即付款减库存的商品现在有多少处于未付款状态的订单 */
	WithHoldQuantity int64 `json:"with_hold_quantity"`

	/* 商品所属的商家的旺旺在线状况，
	taobao.items.search和taobao.items.vip.search专用 */
	WwStatus bool `json:"ww_status"`
}

/* 商品的库存信息和批次信息 */
type WlbItemBatchInventory struct {
	/* 批次库存查询结果 */
	ItemBatch []*WlbItemBatch `json:"item_batch>wlb_item_batch"`

	/* 商品ID */
	ItemId int64 `json:"item_id"`

	/* 商品库存查询结果 */
	ItemInventorys []*WlbItemInventory `json:"item_inventorys>wlb_item_inventory"`

	/* 商品在所有仓库的可销库存总数 */
	TotalQuantity int64 `json:"total_quantity"`
}

/* 物流宝商品库存 */
type WlbItemInventory struct {
	/* 商品id */
	ItemId int64 `json:"item_id"`

	/* 锁定库存数量 */
	LockQuantity int64 `json:"lock_quantity"`

	/* 库存数量 */
	Quantity int64 `json:"quantity"`

	/* 仓库编码 */
	StoreCode string `json:"store_code"`

	/* SELLALBE 可销售库存
	DEFECTIVE 残次
	JISHUN 机损
	XIANGSHUN 箱损
	FREEZE 冻结库存
	ONROAD 在途库存 */
	Type string `json:"type"`
}

/* 视频 */
type VideoItem struct {
	/* 视频封面url */
	CoverUrl string `json:"cover_url"`

	/* 视频描述 */
	Description string `json:"description"`

	/* 视频时长 */
	Duration int64 `json:"duration"`

	/* 是否允许他人观看 */
	IsOpenToOther bool `json:"is_open_to_other"`

	/* 视频状态：等待转码（1），转码中（2），转码失败（3），等待审核（4），未通过审核（5），通过审核（6） */
	State int64 `json:"state"`

	/* 视频标签 */
	Tags []string `json:"tags"`

	/* 视频标题 */
	Title string `json:"title"`

	/* 视频上传时间 */
	UploadTime string `json:"upload_time"`

	/* 视频上传者数字id */
	UploaderId int64 `json:"uploader_id"`

	/* 视频id */
	VideoId int64 `json:"video_id"`

	/* 视频播放地址 */
	VideoPlayInfo *VideoPlayInfo `json:"video_play_info"`
}

/* 视频播放信息 */
type VideoPlayInfo struct {
	/* android  pad兼播放的m3u8列表文件（包含多码率）。适用大于等于3.0版本Android。 */
	AndroidpadUrl string `json:"androidpad_url"`

	/* android pad播放的mp4文件列表。适用2.3版本的Android */
	AndroidpadV23Url *AndroidVlowUrl `json:"androidpad_v23_url"`

	/* android  phone播放的m3u8列表文件（包含多码率，）。适用大于等于3.0版本Android。 */
	AndroidphoneUrl string `json:"androidphone_url"`

	/* android  phone播放的mp4文件列表。适用2.3版本的Android。 */
	AndroidphoneV23Url *AndroidVlowUrl `json:"androidphone_v23_url"`

	/* Flash播放器地址，可直接通过PC浏览器播放 */
	FlashUrl string `json:"flash_url"`

	/* ipad播放的m3u8列表文件（包含多码率） */
	IpadUrl string `json:"ipad_url"`

	/* iphone播放的m3u8列表文件（包含多码率） */
	IphoneUrl string `json:"iphone_url"`

	/* Web嵌入html代码，可直接嵌入页面中，支持html5的video标签，支持HLS播放协议最终返回m3u8资源，否则返回mp4资源 */
	WebUrl string `json:"web_url"`
}

/* 搜索返回的结果类 */
type TOPSearchResult struct {
	/* 分页信息 */
	Paginator *TOPPaginator `json:"paginator"`

	/* 视频列表 */
	VideoItems []*VideoItem `json:"video_items>video_item"`
}

/* 授权 */
type SellerAuthorize struct {
	/* 品牌列表 */
	Brands []*Brand `json:"brands>brand"`

	/* 类目列表 */
	ItemCats []*ItemCat `json:"item_cats>item_cat"`

	/* 被授权的新品类目列表 */
	XinpinItemCats []*ItemCat `json:"xinpin_item_cats>item_cat"`
}

/* 分页信息 */
type TOPPaginator struct {
	/* 当前页码。取值范围:大于零的整数。默认值为1,即默认返回第一页数据。 */
	CurrentPage int64 `json:"current_page"`

	/* 是否最后一页 */
	IsLastPage bool `json:"is_last_page"`

	/* 每页条数。默认值：12 */
	PageSize int64 `json:"page_size"`

	/* 搜索到符合条件的结果总数 */
	TotalResults int64 `json:"total_results"`
}

/* 大家印象标签对应评价信息对象 */
type ImprFeedInfoDO struct {
	/* 1:主评  2:追评 */
	BizType int64 `json:"biz_type"`

	/* 评价内容 */
	Feedback string `json:"feedback"`

	/* 评价创建时间 */
	GmtCreate string `json:"gmt_create"`

	/* 评价的修改时间 */
	GmtModified string `json:"gmt_modified"`

	/* 从评价内容里提取的语义标签 */
	ImprWords []string `json:"impr_words"`
}

/* 商品属性 */
type ItemProp struct {
	/* 子属性的模板（卖家自行输入属性时需要用到） */
	ChildTemplate string `json:"child_template"`

	/* 类目ID */
	Cid int64 `json:"cid"`

	/* 属性的feature列表 */
	Features []*Feature `json:"features>feature"`

	/* 是否允许别名。可选值：true（是），false（否） */
	IsAllowAlias bool `json:"is_allow_alias"`

	/* 是否颜色属性。可选值:true(是),false(否) */
	IsColorProp bool `json:"is_color_prop"`

	/* 是否是可枚举属性。可选值:true(是),false(否) */
	IsEnumProp bool `json:"is_enum_prop"`

	/* 在is_enum_prop是true的前提下，是否是卖家可以自行输入的属性（注：如果is_enum_prop返回false，该参数统一返回false）。可选值:true(是),false(否)。<b>对于品牌和型号属性（包括子属性）：如果用户是C卖家，则可自定义属性；如果是B卖家，则不可自定义属性，而必须要授权的属性。</b> */
	IsInputProp bool `json:"is_input_prop"`

	/* 是否商品属性。可选值:true(是),false(否) */
	IsItemProp bool `json:"is_item_prop"`

	/* 是否关键属性。可选值:true(是),false(否) */
	IsKeyProp bool `json:"is_key_prop"`

	/* 是否销售属性。可选值:true(是),false(否) */
	IsSaleProp bool `json:"is_sale_prop"`

	/* 属性修改时间（增量类目专用） */
	ModifiedTime string `json:"modified_time"`

	/* 三种枚举类型：modify，add，delete（增量类目专用） */
	ModifiedType string `json:"modified_type"`

	/* 发布产品或商品时是否可以多选。可选值:true(是),false(否) */
	Multi bool `json:"multi"`

	/* 发布产品或商品时是否为必选属性。可选值:true(是),false(否) */
	Must bool `json:"must"`

	/* 属性名 */
	Name string `json:"name"`

	/* 上级属性ID */
	ParentPid int64 `json:"parent_pid"`

	/* 上级属性值ID */
	ParentVid int64 `json:"parent_vid"`

	/* 属性 ID 例：品牌的PID=20000 */
	Pid int64 `json:"pid"`

	/*  */
	PropValues []*PropValue `json:"prop_values>prop_value"`

	/* 发布产品或商品时是否为必选属性(与must相同)。可选值:true(是),false(否) */
	Required bool `json:"required"`

	/* 排列序号。取值范围:大于零的整排列序号。取值范围:大于零的整数 */
	SortOrder int64 `json:"sort_order"`

	/* 状态。可选值:normal(正常),deleted(删除) */
	Status string `json:"status"`

	/* 属性值类型。可选值：
	multiCheck(枚举多选)
	optional(枚举单选)
	multiCheckText(枚举可输入多选)
	optionalText(枚举可输入单选)
	text(非枚举可输入) */
	Type string `json:"type"`
}

/* 图片 */
type Picture struct {
	/* 图片上传的来源，有电脑版本宝贝发布，手机版本宝贝发布 */
	ClientType string `json:"client_type"`

	/* 图片的创建时间 */
	Created string `json:"created"`

	/* 图片是否删除的标记 */
	Deleted string `json:"deleted"`

	/* 图片在后台处理之后的md5值
	当md5值为32位长度的字符串时为图片搬家后的文件md5验证码
	md5值为长整数时为图片替换后的时间戳 */
	Md5 string `json:"md5"`

	/* 图片的修改时间 */
	Modified string `json:"modified"`

	/* 图片分类ID */
	PictureCategoryId int64 `json:"picture_category_id"`

	/* 图片ID */
	PictureId int64 `json:"picture_id"`

	/* 返回的是绝对路径如：http://img07.taobaocdn.com/imgextra/i7/22670458/T2dD0kXb4cXXXXXXXX_!!22670458.jpg */
	PicturePath string `json:"picture_path"`

	/* 图片相素,格式:长x宽，如450x150 */
	Pixel string `json:"pixel"`

	/* 图片是否被引用 */
	Referenced bool `json:"referenced"`

	/* 图片大小,bite单位 */
	Sizes int64 `json:"sizes"`

	/* 图片状态,unfroze代表没有被冻结，froze代表被冻结,pass代表排查通过 */
	Status string `json:"status"`

	/* 图片标题 */
	Title string `json:"title"`

	/* 卖家ID */
	Uid int64 `json:"uid"`
}

/* 地址区域结构 */
type Area struct {
	/* 标准行政区域代码.参考:http://www.stats.gov.cn/tjbz/xzqhdm/t20120105_402777427.htm */
	Id int64 `json:"id"`

	/* 地域名称.如北京市,杭州市,西湖区,每一个area_id 都代表了一个具体的地区. */
	Name string `json:"name"`

	/* 父节点区域标识.如北京市的area_id是110100,朝阳区是北京市的一个区,所以朝阳区的parent_id就是北京市的area_id. */
	ParentId int64 `json:"parent_id"`

	/* 区域类型.area区域 1:country/国家;2:province/省/自治区/直辖市;3:city/地区(省下面的地级市);4:district/县/市(县级市)/区;abroad:海外. 比如北京市的area_type = 2,朝阳区是北京市的一个区,所以朝阳区的area_type = 4. */
	Type int64 `json:"type"`

	/* 具体一个地区的邮编 */
	Zip string `json:"zip"`
}

/* 物流公司基础数据结构 */
type LogisticsCompany struct {
	/* 物流公司代码 */
	Code string `json:"code"`

	/* 物流公司标识 */
	Id int64 `json:"id"`

	/* 物流公司简称 */
	Name string `json:"name"`

	/* 运单号验证正则表达式 */
	RegMailNo string `json:"reg_mail_no"`
}

/* 获取的物流订单详情列表 返回的Shipping包含的具体信息为入参fields请求的字段信息 */
type Shipping struct {
	/* 买家昵称 */
	BuyerNick string `json:"buyer_nick"`

	/* 物流公司名称 */
	CompanyName string `json:"company_name"`

	/* 运单创建时间 */
	Created string `json:"created"`

	/* 预约取货结束时间 */
	DeliveryEnd string `json:"delivery_end"`

	/* 预约取货开始时间 */
	DeliveryStart string `json:"delivery_start"`

	/* 谁承担运费.可选值:buyer(买家承担),seller(卖家承担运费). */
	FreightPayer string `json:"freight_payer"`

	/* 标示为是否快捷COD订单 */
	IsQuickCodOrder bool `json:"is_quick_cod_order"`

	/* 表明是否是拆单，默认值0，1表示拆单 */
	IsSpilt int64 `json:"is_spilt"`

	/* 返回发货是否成功。 */
	IsSuccess bool `json:"is_success"`

	/* 货物名称 */
	ItemTitle string `json:"item_title"`

	/* 收件人地址信息(在传输请求参数Fields字段时，必须使用“receiver_location”才能返回此字段) */
	Location *Location `json:"location"`

	/* 运单修改时间 */
	Modified string `json:"modified"`

	/* 物流订单编号 */
	OrderCode string `json:"order_code"`

	/* 运单号.具体一个物流公司的运单号码. */
	OutSid string `json:"out_sid"`

	/* 收件人手机号码 */
	ReceiverMobile string `json:"receiver_mobile"`

	/* 收件人姓名 */
	ReceiverName string `json:"receiver_name"`

	/* 收件人电话 */
	ReceiverPhone string `json:"receiver_phone"`

	/* 卖家是否确认发货.可选值:yes(是),no(否). */
	SellerConfirm string `json:"seller_confirm"`

	/* 卖家昵称 */
	SellerNick string `json:"seller_nick"`

	/* 物流订单状态,可选值:
	CREATED(订单已创建)
	RECREATED(订单重新创建)
	CANCELLED(订单已取消)
	CLOSED(订单关闭)
	SENDING(等候发送给物流公司)
	ACCEPTING(已发送给物流公司,等待接单)
	ACCEPTED(物流公司已接单)
	REJECTED(物流公司不接单)
	PICK_UP(物流公司揽收成功)
	PICK_UP_FAILED(物流公司揽收失败)
	LOST(物流公司丢单)
	REJECTED_BY_RECEIVER(对方拒签)
	ACCEPTED_BY_RECEIVER(发货方式在线下单：对方已签收；自己联系：卖家已发货) */
	Status string `json:"status"`

	/* 拆单子订单列表，对应的数据是：该物流订单下的全部子订单 */
	SubTids []int64 `json:"sub_tids"`

	/* 交易ID */
	Tid int64 `json:"tid"`

	/* 物流方式.可选值:free(卖家包邮),post(平邮),express(快递),ems(EMS). */
	Type string `json:"type"`
}

/* android  phone和pad播放的mp4文件类。适用2.3版本的Android。 */
type AndroidVlowUrl struct {
	/* android phone和pad播放的高清mp4文件。适用2.3版本的Android */
	Hd string `json:"hd"`

	/* android phone和pad播放的流畅mp4文件。适用2.3版本的Android */
	Ld string `json:"ld"`

	/* android phone和pad播放的标清mp4文件。适用2.3版本的Android */
	Sd string `json:"sd"`

	/* android phone和pad播放的超清mp4文件。适用2.3版本的Android */
	Ud string `json:"ud"`
}

/* 评价列表 */
type TradeRate struct {
	/* 评价内容,最大长度:500个汉字 */
	Content string `json:"content"`

	/* 评价创建时间,格式:yyyy-MM-dd HH:mm:ss */
	Created string `json:"created"`

	/* 商品价格,精确到2位小数;单位:元.如:200.07，表示:200元7分 */
	ItemPrice float64 `json:"item_price,string"`

	/* 商品标题 */
	ItemTitle string `json:"item_title"`

	/* 评价者昵称 */
	Nick string `json:"nick"`

	/* 商品的数字ID */
	NumIid int64 `json:"num_iid"`

	/* 子订单ID */
	Oid int64 `json:"oid"`

	/* 被评价者昵称 */
	RatedNick string `json:"rated_nick"`

	/* 评价解释,最大长度:500个汉字 */
	Reply string `json:"reply"`

	/* 评价结果,可选值:good(好评),neutral(中评),bad(差评) */
	Result string `json:"result"`

	/* 评价者角色.可选值:seller(卖家),buyer(买家) */
	Role string `json:"role"`

	/* 交易ID */
	Tid int64 `json:"tid"`

	/* 评价信息是否用于记分，
	可取值：true(参与记分)和false(不参与记分) */
	ValidScore bool `json:"valid_score"`
}

/* 库存明细 */
type IpcInventoryDetailDo struct {
	/* 主订单号 */
	BizOrderId int64 `json:"biz_order_id"`

	/* 子订单号 */
	BizSubOrderId int64 `json:"biz_sub_order_id"`

	/* 1拍减 2付减 */
	Flag int64 `json:"flag"`

	/* 占用数量 */
	OccupyQuantity int64 `json:"occupy_quantity"`

	/* 货主昵称 */
	OwnerNick string `json:"owner_nick"`

	/* 预扣库存数 */
	ReserveQuantity int64 `json:"reserve_quantity"`

	/* 仓储商品id */
	ScItemId int64 `json:"sc_item_id"`

	/* 仓库编码 */
	StoreCode string `json:"store_code"`
}

/* 分销产品SKU */
type FenxiaoSku struct {
	/* 代销采购价，单位：元 */
	CostPrice string `json:"cost_price"`

	/* 经销采购价 */
	DealerCostPrice string `json:"dealer_cost_price"`

	/* SkuID */
	Id int64 `json:"id"`

	/* 名称 */
	Name string `json:"name"`

	/* 商家编码 */
	OuterId string `json:"outer_id"`

	/* sku的销售属性组合字符串。格式:pid:vid;pid:vid,如:1627207:3232483;1630696:3284570,表示:机身颜色:军绿色;手机套餐:一电一充。 */
	Properties string `json:"properties"`

	/* 库存 */
	Quantity int64 `json:"quantity"`

	/* 配额可用库存 */
	QuotaQuantity int64 `json:"quota_quantity"`

	/* 预扣库存 */
	ReservedQuantity int64 `json:"reserved_quantity"`

	/* 关联的后端商品id */
	ScitemId int64 `json:"scitem_id"`

	/* 市场价 */
	StandardPrice string `json:"standard_price"`
}

/* 营销工具满就送 */
type PromotionMjs struct {
	/* 不免邮省份的area_id列表。用';'间隔。最多10个。[110000:"北京";
	120000:"天津";
	130000:"河北";
	140000:"山西";
	150000:"内蒙古";
	210000:"辽宁";
	220000:"吉林";
	230000:"黑龙江";
	310000:"上海";
	320000:"江苏";
	330000:"浙江";
	340000:"安徽";
	350000:"福建";
	360000:"江西";
	370000:"山东";
	410000:"河南";
	420000:"湖北";
	430000:"湖南";
	440000:"广东";
	450000:"广西";
	460000:"海南";
	500000:"重庆";
	510000:"四川";
	520000:"贵州";
	530000:"云南";
	540000:"西藏";
	610000:"陕西";
	620000:"甘肃";
	630000:"青海";
	640000:"宁夏";
	650000:"新疆";
	710000:"台湾";
	810000:"香港";
	820000:"澳门";
	990000:"海外"] */
	AreaIds string `json:"area_ids"`

	/* 送优惠卷。优惠卷的有效期时间。 */
	CouponTime string `json:"coupon_time"`

	/* 送优惠卷面额。分为[3,5,10,20,50,100]面值。 */
	CouponValue int64 `json:"coupon_value"`

	/* 是否免邮！ */
	FreePost bool `json:"free_post"`

	/* 满就送的满多少钱。正整数。 */
	FullMoney int64 `json:"full_money"`

	/* 最长15个汉字。 */
	GiftName string `json:"gift_name"`

	/* 礼品的url地址。礼品链接必须是以http://或https://开头的淘宝站内链接。正则表达式:"^http[s]{0,1}://([a-z0-9\\-_]+\\.)*(mall\\.)?(taobao|taobaocdn)\\.(com|net|cn|com\\.cn)(/.*)*$" */
	GiftUrl string `json:"gift_url"`

	/* 满就送积分。 */
	Points int64 `json:"points"`

	/* true时，减后满full_money才免邮费。false时，减前满full_money就免邮。 */
	ReduceAfter bool `json:"reduce_after"`

	/* 满就送减掉的钱数。当满足full_money时，就减掉的钱数。这个钱数是小于full_money的正整数。 */
	ReduceMoney int64 `json:"reduce_money"`
}

/* 收货人详细信息 */
type Receiver struct {
	/* 收货人的详细地址信息 */
	Address string `json:"address"`

	/* 证件号 */
	CardId string `json:"card_id"`

	/* 收货人的城市 */
	City string `json:"city"`

	/* 区/县 */
	District string `json:"district"`

	/* 移动电话。 */
	MobilePhone string `json:"mobile_phone"`

	/* 收货人全名。 */
	Name string `json:"name"`

	/* 固定电话。 */
	Phone string `json:"phone"`

	/* 收货人所在省份 */
	State string `json:"state"`

	/* 邮政编码 */
	Zip string `json:"zip"`
}

/* 权限信息 */
type Permission struct {
	/* 1 :允许授权  2：不允许授权 6：不允许授权但默认已有权限 */
	IsAuthorize int64 `json:"is_authorize"`

	/* 父权限code */
	ParentCode string `json:"parent_code"`

	/* 注册到权限中心的code值 */
	PermissionCode string `json:"permission_code"`

	/* 权限名称 */
	PermissionName string `json:"permission_name"`
}

/* 登录分销用户信息 */
type LoginUser struct {
	/* 入驻时间 */
	CreateTime string `json:"create_time"`

	/* 会员NICK */
	Nick string `json:"nick"`

	/* 分销用户ID */
	UserId int64 `json:"user_id"`

	/* 分销用户类型(1:分销商，2:供应商，3:品牌商；2、3都表示有供货能力，只是身份不同) */
	UserType string `json:"user_type"`
}

/* 商品详情模块 */
type ItemDescModule struct {
	/* 描述具体内容 */
	Content string `json:"content"`

	/* 一个List<String>的Json串，里面是模块引导提示，不超过3条（isv提交时可忽略不传） */
	Intros string `json:"intros"`

	/* 模块id ,(注意:用户自定义模块不用填此项。) */
	ModuleId int64 `json:"module_id"`

	/* 模块名称 */
	ModuleName string `json:"module_name"`

	/* 是否必填 （isv提交时可忽略不传） */
	Required bool `json:"required"`

	/* 淘宝图片空间的地址链接，示例模板，最多不超过三个（isv提交时可忽略不传） */
	TplUrls string `json:"tpl_urls"`

	/* cat_mod:表示是运营设置的类目维度模块，usr_mod表示的是商家自定义模块。 */
	Type string `json:"type"`
}

/* 商品描述模块化信息，包含类目级别的模块信息以及用户自定义模块约束等其他信息 */
type ModularDescInfo struct {
	/* 旧描述与新描述共存时间（切换时间）。 */
	DeadLine string `json:"dead_line"`

	/* 运营定义的该商品所属类目的类目级别模块信息列表，列表顺序即为模块顺序。 */
	ItmDscModules []*ItemDescModule `json:"itm_dsc_modules>item_desc_module"`

	/* 默认值为false，如果此字段为true，则卖家上传的模块化的描述信息可以自定义排序。 */
	UserOrder bool `json:"user_order"`

	/* 用户自定义模块数量最大值限制。类目级别模块+用户级别模块必须小于<20 */
	UsrDefMaxNum int64 `json:"usr_def_max_num"`
}

/* 采购单及子采购单信息 */
type PurchaseOrder struct {
	/* 支付宝交易号。 */
	AlipayNo string `json:"alipay_no"`

	/* 买家nick，供应商查询不会返回买家昵称，分销商查询才会返回。 */
	BuyerNick string `json:"buyer_nick"`

	/* 买家支付给分销商的总金额。注意买家购买的商品可能不是全部来自同一供货商，请同时参考子单上的相关金额。（精确到2位小数;单位:元。如:200.07，表示:200元7分） */
	BuyerPayment float64 `json:"buyer_payment,string"`

	/* 加密后的买家淘宝ID，长度为32位 */
	BuyerTaobaoId string `json:"buyer_taobao_id"`

	/* 物流发货时间。格式:yyyy-MM-dd HH:mm:ss */
	ConsignTime string `json:"consign_time"`

	/* 采购单创建时间。格式:yyyy-MM-dd HH:mm:ss */
	Created string `json:"created"`

	/* 分销商来源网站（taobao）。 */
	DistributorFrom string `json:"distributor_from"`

	/* 分销商实付金额。(精确到2位小数;单位:元。如:200.07，表示:200元7分 ) */
	DistributorPayment float64 `json:"distributor_payment,string"`

	/* 分销商在来源网站的帐号名。 */
	DistributorUsername string `json:"distributor_username"`

	/* 交易结束时间 */
	EndTime string `json:"end_time"`

	/* 主订单属性信息，key-value形式：
	orderNovice ：订单发票抬头；
	orderNoviceContent ：代表发票明细 */
	Features []*Feature `json:"features>feature"`

	/* 分销流水号，分销平台产生的主键 */
	FenxiaoId int64 `json:"fenxiao_id"`

	/* 供应商交易ID 非采购单ID，如果改发货状态 是需要该ID，ID在用户未付款前为0，付款后有具体值（发货时使用该ID） */
	Id int64 `json:"id"`

	/* 自定义key */
	IsvCustomKey []string `json:"isv_custom_key"`

	/* 自定义值 */
	IsvCustomValue []string `json:"isv_custom_value"`

	/* 物流公司 */
	LogisticsCompanyName string `json:"logistics_company_name"`

	/* 运单号 */
	LogisticsId string `json:"logistics_id"`

	/* 采购单留言。（代销模式下信息包括买家留言和分销商留言） */
	Memo string `json:"memo"`

	/* 交易修改时间。格式:yyyy-MM-dd HH:mm:ss */
	Modified string `json:"modified"`

	/* 采购单留言列表 */
	OrderMessages []*OrderMessage `json:"order_messages>order_message"`

	/* 付款时间。格式:yyyy-MM-dd HH:mm:ss */
	PayTime string `json:"pay_time"`

	/* 支付方式：ALIPAY_SURETY（支付宝担保交易）、ALIPAY_CHAIN（分账交易）、TRANSFER（线下转账）、PREPAY（预存款）、IMMEDIATELY（即时到账）、CASHGOODS（先款后货） */
	PayType string `json:"pay_type"`

	/* 采购单邮费。(精确到2位小数;单位:元。如:200.07，表示:200元7分 ) */
	PostFee float64 `json:"post_fee,string"`

	/* 买家详细的信息。 */
	Receiver *Receiver `json:"receiver"`

	/* 配送方式，FAST(快速)、EMS、ORDINARY(平邮)、SELLER(卖家包邮) */
	Shipping string `json:"shipping"`

	/* 订单快照URL */
	SnapshotUrl string `json:"snapshot_url"`

	/* 采购单交易状态。可选值：<br>
	WAIT_BUYER_PAY(等待付款)<br>
	WAIT_SELLER_SEND_GOODS(已付款，待发货）<br>
	WAIT_BUYER_CONFIRM_GOODS(已付款，已发货)<br>
	TRADE_FINISHED(交易成功)<br>
	TRADE_CLOSED(交易关闭)<br>
	WAIT_BUYER_CONFIRM_GOODS_ACOUNTED(已付款（已分账），已发货。只对代销分账支持)<br>
	PAY_ACOUNTED_GOODS_CONFIRM （已分账发货成功）<br>
	PAY_WAIT_ACOUNT_GOODS_CONFIRM（已付款，确认收货） */
	Status string `json:"status"`

	/* 子订单的详细信息列表。 */
	SubPurchaseOrders []*SubPurchaseOrder `json:"sub_purchase_orders>sub_purchase_order"`

	/* 返回供应商备注旗帜
	vlaue在1-5之间。非1-5之间，都采用1作为默认。 1:红色 2:黄色 3:绿色 4:蓝色 5:粉红色 */
	SupplierFlag int64 `json:"supplier_flag"`

	/* 供应商来源网站, values: taobao, alibaba */
	SupplierFrom string `json:"supplier_from"`

	/* 供应商备注 */
	SupplierMemo string `json:"supplier_memo"`

	/* 供应商在来源网站的帐号名。 */
	SupplierUsername string `json:"supplier_username"`

	/* 主订单ID （经销不显示） */
	TcOrderId int64 `json:"tc_order_id"`

	/* 采购单总额（不含邮费,精确到2位小数;单位:元。如:200.07，表示:200元7分 ) */
	TotalFee float64 `json:"total_fee,string"`

	/* 分销方式：AGENT（代销）、DEALER（经销） */
	TradeType string `json:"trade_type"`
}

/* 分销产品 */
type FenxiaoProduct struct {
	/* 警戒库存 */
	AlarmNumber int64 `json:"alarm_number"`

	/* 类目id */
	CategoryId string `json:"category_id"`

	/* 所在地：市 */
	City string `json:"city"`

	/* 采购价格，单位：元。 */
	CostPrice float64 `json:"cost_price,string"`

	/* 创建时间 */
	Created string `json:"created"`

	/* 经销采购价 */
	DealerCostPrice float64 `json:"dealer_cost_price,string"`

	/* 产品描述路径，通过http请求获取 */
	DescPath string `json:"desc_path"`

	/* 产品描述的内容 */
	Description string `json:"description"`

	/* 折扣ID（新增参数） */
	DiscountId int64 `json:"discount_id"`

	/* 是否有保修，可选值：false（否）、true（是） */
	HaveGuarantee bool `json:"have_guarantee"`

	/* 是否有发票，可选值：false（否）、true（是） */
	HaveInvoice bool `json:"have_invoice"`

	/* 产品图片 */
	Images []*FenxiaoImage `json:"images>fenxiao_image"`

	/* 自定义属性，格式为pid:value;pid:value */
	InputProperties string `json:"input_properties"`

	/* 查询产品列表时，查询入参增加is_authz:yes|no
	yes:需要授权
	no:不需要授权 */
	IsAuthz string `json:"is_authz"`

	/* 导入的商品ID */
	ItemId int64 `json:"item_id"`

	/* 下载人数 */
	ItemsCount int64 `json:"items_count"`

	/* 更新时间 */
	Modified string `json:"modified"`

	/* 产品名称 */
	Name string `json:"name"`

	/* 累计采购次数 */
	OrdersCount int64 `json:"orders_count"`

	/* 商家编码 */
	OuterId string `json:"outer_id"`

	/* 产品分销商信息 */
	Pdus []*FenxiaoPdu `json:"pdus>fenxiao_pdu"`

	/* 产品图片路径 */
	Pictures string `json:"pictures"`

	/* 产品ID */
	Pid int64 `json:"pid"`

	/* ems费用，单位：元 */
	PostageEms float64 `json:"postage_ems,string"`

	/* 快递费用，单位：元 */
	PostageFast float64 `json:"postage_fast,string"`

	/* 运费模板ID */
	PostageId int64 `json:"postage_id"`

	/* 平邮费用，单位：元 */
	PostageOrdinary float64 `json:"postage_ordinary,string"`

	/* 运费类型，可选值：seller（供应商承担运费）、buyer（分销商承担运费） */
	PostageType string `json:"postage_type"`

	/* 产品线ID */
	ProductcatId int64 `json:"productcat_id"`

	/* 产品属性，格式为pid:vid;pid:vid */
	Properties string `json:"properties"`

	/* 属性别名，格式为pid:vid:alias;pid:vid:alias */
	PropertyAlias string `json:"property_alias"`

	/* 所在地：省 */
	Prov string `json:"prov"`

	/* 产品库存 */
	Quantity int64 `json:"quantity"`

	/* 根据商品ID查询时，返回这个产品对应的商品ID，只对分销商查询接口有效 */
	QueryItemId int64 `json:"query_item_id"`

	/* 配额可用库存 */
	QuotaQuantity int64 `json:"quota_quantity"`

	/* 预扣库存 */
	ReservedQuantity int64 `json:"reserved_quantity"`

	/* 最高零售价，单位：分。 */
	RetailPriceHigh float64 `json:"retail_price_high,string"`

	/* 最低零售价，单位：分。 */
	RetailPriceLow float64 `json:"retail_price_low,string"`

	/* 关联的后端商品id */
	ScitemId int64 `json:"scitem_id"`

	/* sku列表 */
	Skus []*FenxiaoSku `json:"skus>fenxiao_sku"`

	/* 产品spu id */
	SpuId int64 `json:"spu_id"`

	/* 采购基准价，单位：元。 */
	StandardPrice float64 `json:"standard_price,string"`

	/* 零售基准价，单位：元 */
	StandardRetailPrice float64 `json:"standard_retail_price,string"`

	/* 发布状态，可选值：up（上架）、down（下架） */
	Status string `json:"status"`

	/* 分销方式：AGENT（只做代销，默认值）、DEALER（只做经销）、ALL（代销和经销都做） */
	TradeType string `json:"trade_type"`

	/* 铺货时间 */
	UpshelfTime string `json:"upshelf_time"`
}

/* 产品扩展信息 */
type ProductExtraInfo struct {
	/* 产品扩展信息键 */
	FieldKey string `json:"field_key"`

	/* 产品扩展信息名称 */
	FieldName string `json:"field_name"`

	/* 产品扩展信息简介 */
	FieldValue string `json:"field_value"`

	/* 产品ID */
	ProductId int64 `json:"product_id"`
}

/* 子账号基本信息 */
type SubUserInfo struct {
	/* 子账号姓名 */
	FullName string `json:"full_name"`

	/* 是否参与分流 1不参与 2参与 */
	IsOnline int64 `json:"is_online"`

	/* 子账号用户名 */
	Nick string `json:"nick"`

	/* 子账号所属的主账号的唯一标识 */
	SellerId int64 `json:"seller_id"`

	/* 主账号昵称 */
	SellerNick string `json:"seller_nick"`

	/* 子账号当前状态 1正常 -1删除  2冻结 */
	Status int64 `json:"status"`

	/* 子账号Id */
	SubId int64 `json:"sub_id"`
}

/* 使用数据推送的用户信息 */
type JdpUser struct {
	/* rds数据库的实例名 */
	RdsName string `json:"rds_name"`

	/* 0:暂停
	1：正常
	2：sessoin失效，停止
	3：已删除 */
	Status int64 `json:"status"`

	/* 用户昵称 */
	UserNick string `json:"user_nick"`
}

/* 子采购单详细信息 */
type SubPurchaseOrder struct {
	/* 分销商店铺中宝贝一口价 */
	AuctionPrice float64 `json:"auction_price,string"`

	/* 发票应开金额。根据买家实际付款去除邮费后，按各个子单(商品)金额比例进行分摊后的金额，仅供开发票时做票面金额参考。 */
	BillFee float64 `json:"bill_fee,string"`

	/* 买家订单上对应的子单零售金额，除以num（数量）后等于最终宝贝的零售价格（精确到2位小数;单位:元。如:200.07，表示:200元7分） */
	BuyerPayment float64 `json:"buyer_payment,string"`

	/* 创建时间。格式 yyyy-MM-dd HH:mm:ss 。 */
	Created string `json:"created"`

	/* 优惠活动的折扣金额 */
	DiscountFee float64 `json:"discount_fee,string"`

	/* 分销商实付金额=total_fee（分销商应付金额）+改价-优惠。（精确到2位小数;单位:元。如:200.07，表示:200元7分） */
	DistributorPayment float64 `json:"distributor_payment,string"`

	/* Feature对象列表
	目前已有的属性：
	attr_key为 www，attr_value为1 表示是www子订单
	attr_key为 wwwStoreCode，attr_value 是www子订单发货的仓库编码
	attr_key为isWt,attr_value为1 表示是网厅子订单
	attr_key为wtInfo,attr_value为网厅相关合约信息 */
	Features []*Feature `json:"features>feature"`

	/* 分销平台的子采购单主键 */
	FenxiaoId int64 `json:"fenxiao_id"`

	/* 子采购单id，淘宝交易主键，采购单未付款时为0.（只有支付宝 付款才有这个id，其余付款形式该字段为0） */
	Id int64 `json:"id"`

	/* 分销平台上的产品id，同FenxiaoProduct 的pid */
	ItemId int64 `json:"item_id"`

	/* 分销平台上商品商家编码。 */
	ItemOuterId string `json:"item_outer_id"`

	/* 产品的采购数量。取值范围:大于零的整数 */
	Num int64 `json:"num"`

	/* 老的SKU属性值。如: 颜色:红色:别名;尺码:L:别名 */
	OldSkuProperties string `json:"old_sku_properties"`

	/* 代销采购单对应下游200订单状态。
	可选值：
	WAIT_SELLER_SEND_GOODS(已付款，待发货)
	WAIT_BUYER_CONFIRM_GOODS(已付款，已发货)
	TRADE_CLOSED(已退款成功)
	TRADE_REFUNDING(退款中)
	TRADE_FINISHED(交易成功)
	TRADE_CLOSED_BY_TAOBAO(交易关闭) */
	Order200Status string `json:"order_200_status"`

	/* 产品的采购价格。（精确到2位小数;单位:元。如:200.07，表示:200元7分） */
	Price float64 `json:"price,string"`

	/* 优惠活动类型
	0=无优惠
	1=限时折 */
	PromotionType string `json:"promotion_type"`

	/* 退款金额 */
	RefundFee float64 `json:"refund_fee,string"`

	/* 后端商品id */
	ScItemId int64 `json:"sc_item_id"`

	/* 商品的SKU id。当存在时才会有值，建议使用sku_outer_id，sku_properties这两个值 */
	SkuId int64 `json:"sku_id"`

	/* SKU商家编码。 */
	SkuOuterId string `json:"sku_outer_id"`

	/* SKU属性值。如: 颜色:红色:别名;尺码:L:别名 */
	SkuProperties string `json:"sku_properties"`

	/* 快照地址。 */
	SnapshotUrl string `json:"snapshot_url"`

	/* 交易状态。可选值：<br>
	WAIT_BUYER_PAY(等待付款)<br>
	WAIT_CONFIRM(付款信息待确认)<br>
	WAIT_CONFIRM_WAIT_SEND_GOODS(付款信息待确认，待发货)<br>
	WAIT_CONFIRM_SEND_GOODS(付款信息待确认，已发货)<br>
	WAIT_CONFIRM_GOODS_CONFIRM(付款信息待确认，已收货)<br>
	WAIT_SELLER_SEND_GOODS(已付款，待发货)<br>
	WAIT_BUYER_CONFIRM_GOODS(已付款，已发货)<br>
	WAIT_BUYER_CONFIRM_GOODS_ACOUNTED(已付款（已分账），已发货。只对代销分账支持)<br>
	CONFIRM_WAIT_SEND_GOODS(付款信息已确认，待发货)<br>
	CONFIRM_SEND_GOODS(付款信息已确认，已发货)<br>
	TRADE_REFUNDED(已退款)<br>
	TRADE_REFUNDING(退款中)<br>
	TRADE_FINISHED(交易成功)<br>
	TRADE_CLOSED(交易关闭)<br>
	PAY_ACOUNTED_GOODS_CONFIRM （已分账发货成功）<br>
	PAY_WAIT_ACOUNT_GOODS_CONFIRM（已付款，确认收货） */
	Status string `json:"status"`

	/* 商品的卖出金额调整，金额增加时为正数，金额减少时为负数，单位是分，不带小数 */
	TcAdjustFee int64 `json:"tc_adjust_fee"`

	/* 优惠金额，始终为正数，单位是分，不带小数 */
	TcDiscountFee int64 `json:"tc_discount_fee"`

	/* TC子订单ID（经销不显示） */
	TcOrderId int64 `json:"tc_order_id"`

	/* 商品优惠类型：聚划算、秒杀或其他 */
	TcPreferentialType string `json:"tc_preferential_type"`

	/* 采购的产品标题。 */
	Title string `json:"title"`

	/* 分销商应付金额=num(采购数量)*price(采购价)。（精确到2位小数;单位:元。如:200.07，表示:200元7分） */
	TotalFee float64 `json:"total_fee,string"`
}

/* rds的配置信息 */
type RdsConfig struct {
	/* appkey */
	Appkey string `json:"appkey"`

	/* 数据库类型，mysql/sqlserver */
	DbType string `json:"db_type"`

	/* 数据库用户名 */
	DbUser string `json:"db_user"`

	/* 主机地址 */
	Host string `json:"host"`

	/* 数据推送rds表里的id */
	Id int64 `json:"id"`

	/* rds数据库实例名 */
	Name string `json:"name"`

	/* rds数据库端口 */
	Port string `json:"port"`

	/* 数据库状态 */
	Status string `json:"status"`
}

/* 产品线 */
type ProductCat struct {
	/* 代销采购价百分比 */
	CostPercentAgent string `json:"cost_percent_agent"`

	/* 经销采购价百分比 */
	CostPercentDealer string `json:"cost_percent_dealer"`

	/* 产品线ID */
	Id int64 `json:"id"`

	/* 产品线名称 */
	Name string `json:"name"`

	/* 产品数量 */
	ProductNum int64 `json:"product_num"`

	/* 最高零食价百分比 */
	RetailHighPercent string `json:"retail_high_percent"`

	/* 最低零售价百分比 */
	RetailLowPercent string `json:"retail_low_percent"`
}

/* 类目属性 */
type Feature struct {
	/* 属性键 */
	AttrKey string `json:"attr_key"`

	/* 属性值 */
	AttrValue string `json:"attr_value"`
}

/* 消息通道的用户结构 */
type TmcUser struct {
	/* 用户首次开通时间 */
	Created string `json:"created"`

	/* 接收用户消息的组名 */
	GroupName string `json:"group_name"`

	/* 用户授权是否有效，true表示授权有效，false表示授权过期 */
	IsValid bool `json:"is_valid"`

	/* 用户最后开通时间 */
	Modified string `json:"modified"`

	/* 用户开通的消息类型列表。如果为空表示应用开通的所有类型 */
	Topics []string `json:"topics"`

	/* 用户ID */
	UserId int64 `json:"user_id"`

	/* 用户昵称 */
	UserNick string `json:"user_nick"`
}

/* 消息通道的通用消息结构 */
type TmcMessage struct {
	/* 消息详细内容，格式为JSON/XML */
	Content string `json:"content"`

	/* 消息ID */
	Id int64 `json:"id"`

	/* 消息发布者的AppKey */
	PubAppKey string `json:"pub_app_key"`

	/* 消息发布时间 */
	PubTime string `json:"pub_time"`

	/* 消息所属主题 */
	Topic string `json:"topic"`

	/* 消息所属的用户编号 */
	UserId int64 `json:"user_id"`

	/* 用户的昵称 */
	UserNick string `json:"user_nick"`
}

/* 消息通道的分组结构 */
type TmcGroup struct {
	/* 分组名称 */
	Name string `json:"name"`

	/* 用户列表 */
	Users []string `json:"users"`
}

/* 采购申请和经销采购单 */
type DealerOrder struct {
	/* 支付宝交易号 */
	AlipayNo string `json:"alipay_no"`

	/* 申请时间 */
	AppliedTime string `json:"applied_time"`

	/* 分销商nick */
	ApplierNick string `json:"applier_nick"`

	/* 分销商最后一次确认/申请/拒绝的时间 */
	AuditTimeApplier string `json:"audit_time_applier"`

	/* 供应商最后一次审核通过/修改/驳回的时间 */
	AuditTimeSupplier string `json:"audit_time_supplier"`

	/* 关闭原因 */
	CloseReason string `json:"close_reason"`

	/* 产品明细 */
	DealerOrderDetails []*DealerOrderDetail `json:"dealer_order_details>dealer_order_detail"`

	/* 经销采购单编号，API发货使用此字段 */
	DealerOrderId int64 `json:"dealer_order_id"`

	/* 已发货数量 */
	DeliveredQuantityCount int64 `json:"delivered_quantity_count"`

	/* 物流费用(精确到2位小数;单位:元。如:200.07，表示:200元7分 ) */
	LogisticsFee float64 `json:"logistics_fee,string"`

	/* 物流方式：
	SELF_PICKUP（自提）、LOGISTICS（物流发货) */
	LogisticsType string `json:"logistics_type"`

	/* 修改时间 */
	ModifiedTime string `json:"modified_time"`

	/* WAIT_FOR_SUPPLIER_AUDIT1：分销商提交申请，待供应商审核；
	SUPPLIER_REFUSE：供应商驳回申请，待分销商确认；
	WAIT_FOR_APPLIER_AUDIT：供应商修改后，待分销商确认；
	WAIT_FOR_SUPPLIER_AUDIT2：分销商拒绝修改，待供应商再审核；
	BOTH_AGREE_WAIT_PAY：审核通过下单成功，待分销商付款
	WAIT_FOR_SUPPLIER_DELIVER：付款成功，待供应商发货；
	WAIT_FOR_APPLIER_STORAGE：供应商发货，待分销商收货；
	TRADE_FINISHED：分销商收货，交易成功；
	TRADE_CLOSED：经销采购单关闭。 */
	OrderStatus string `json:"order_status"`

	/* 付款时间 */
	PayTime string `json:"pay_time"`

	/* 支付方式：
	ALIPAY_SURETY（支付宝担保交易）
	TRANSFER（线下转账）
	PREPAY（预存款）
	IMMEDIATELY（即时到账） */
	PayType string `json:"pay_type"`

	/* 总采购数量 */
	QuantityCount int64 `json:"quantity_count"`

	/* 收货人信息 */
	Receiver *Receiver `json:"receiver"`

	/* 分销商拒绝供应商修改的原因 */
	RefuseReasonApplier string `json:"refuse_reason_applier"`

	/* 供应商驳回申请的原因 */
	RefuseReasonSupplier string `json:"refuse_reason_supplier"`

	/* 供应商备注。
	仅供应商可见。 */
	SupplierMemo string `json:"supplier_memo"`

	/* 供应商备注旗帜。
	1:红色 2:黄色 3:绿色 4:蓝色 5:粉红色。
	仅供应商可见。 */
	SupplierMemoFlag int64 `json:"supplier_memo_flag"`

	/* 供应商nick */
	SupplierNick string `json:"supplier_nick"`

	/* 采购总价(精确到2位小数;单位:元。如:200.07，表示:200元7分 ) */
	TotalPrice float64 `json:"total_price,string"`
}

/* 采购申请/经销采购单中的商品明细 */
type DealerOrderDetail struct {
	/* 经销采购单明细id */
	DealerDetailId int64 `json:"dealer_detail_id"`

	/* 经销采购单编号 */
	DealerOrderId int64 `json:"dealer_order_id"`

	/* 最终价格(精确到2位小数;单位:元。如:200.07，表示:200元7分 ) */
	FinalPrice float64 `json:"final_price,string"`

	/* 该条明细是否已删除。true：已删除；false：未删除。 */
	IsDeleted bool `json:"is_deleted"`

	/* 原始价格(精确到2位小数;单位:元。如:200.07，表示:200元7分 ) */
	OriginalPrice float64 `json:"original_price,string"`

	/* 金额小计(采购数量乘以最终价格。精确到2位小数;单位:元。如:200.07，表示:200元7分 ) */
	PriceCount float64 `json:"price_count,string"`

	/* 产品id */
	ProductId int64 `json:"product_id"`

	/* 产品标题 */
	ProductTitle string `json:"product_title"`

	/* 采购数量 */
	Quantity int64 `json:"quantity"`

	/* sku id */
	SkuId int64 `json:"sku_id"`

	/* 商家编码，是sku则为sku的商家编码，否则是产品的商家编码 */
	SkuNumber string `json:"sku_number"`

	/* 产品规格 */
	SkuSpec string `json:"sku_spec"`

	/* 产品快照url */
	SnapshotUrl string `json:"snapshot_url"`
}

/* 产品资质认证文本信息，包括认证类型以及文本信息 */
type CertTxtInfo struct {
	/* 认证类型的数值id */
	CertType int64 `json:"cert_type"`

	/* 认证文本信息 */
	Text string `json:"text"`
}

/* tmall权益 */
type GradeEquity struct {
	/* 买家会员级别  0:店铺客户  1：普通会员 2：高级会员 3：VIP会员 4：至尊VIP */
	CurGrade string `json:"cur_grade"`

	/* 店铺客户、普通会员 、高级会员、VIP会员、至尊VIP */
	CurGradeName string `json:"cur_grade_name"`

	/* 在包邮的前提下的不免邮区域 */
	ExcludeArea string `json:"exclude_area"`

	/* 返回几倍tmall积分 */
	Point int64 `json:"point"`

	/* 是否该等级开启包邮 */
	Postage bool `json:"postage"`
}

/* 图片分类 */
type PictureCategory struct {
	/* 图片类目的创建时间 */
	Created string `json:"created"`

	/* 图片分类的修改时间 */
	Modified string `json:"modified"`

	/* 一级分类的parent_id为0
	二级分类的则为其父分类的picture_category_id */
	ParentId int64 `json:"parent_id"`

	/* 图片分类ID */
	PictureCategoryId int64 `json:"picture_category_id"`

	/* 图片分类名 */
	PictureCategoryName string `json:"picture_category_name"`

	/* 图片分类排序 */
	Position int64 `json:"position"`

	/* 图片分类型别，sys-fixture代表店铺装修分类(系统分类)，sys-auction代表宝贝图片分类(系统分类)，user-define代表用户自定义分类 */
	Type string `json:"type"`
}

/* 分销API返回数据结构 */
type Distributor struct {
	/* 分销商的支付宝帐户 */
	AlipayAccount string `json:"alipay_account"`

	/* 分销商的淘宝卖家评价 */
	Appraise int64 `json:"appraise"`

	/* 分销商店铺主营类目 */
	CategoryId int64 `json:"category_id"`

	/* 联系人 */
	ContactPerson string `json:"contact_person"`

	/* 分销商创建时间 时间格式：yyyy-MM-dd HH:mm:ss */
	Created string `json:"created"`

	/* 分销商Id */
	DistributorId int64 `json:"distributor_id"`

	/* 分销商姓名 */
	DistributorName string `json:"distributor_name"`

	/* 分销商的email */
	Email string `json:"email"`

	/* 分销商的真实姓名，认证姓名 */
	FullName string `json:"full_name"`

	/* 店铺等级 */
	Level int64 `json:"level"`

	/* 分销商的手机号 */
	MobilePhone string `json:"mobile_phone"`

	/* 分销商的电话 */
	Phone string `json:"phone"`

	/* 分销商的网店链接 */
	ShopWebLink string `json:"shop_web_link"`

	/* 分销商卖家的开店时间 */
	Starts string `json:"starts"`

	/* 分销商ID */
	UserId int64 `json:"user_id"`
}

/* 表示会员关系的基本信息字段，用于会员列表的基本查询 */
type BasicMember struct {
	/* 最后一次交易的订单号.
	注:该字段从2014.4.23之后不再返回. */
	BizOrderId int64 `json:"biz_order_id"`

	/* 买家会员ID */
	BuyerId int64 `json:"buyer_id"`

	/* 会员昵称 */
	BuyerNick string `json:"buyer_nick"`

	/* 交易关闭金额 */
	CloseTradeAmount float64 `json:"close_trade_amount,string"`

	/* 交易关闭的笔数 */
	CloseTradeCount int64 `json:"close_trade_count"`

	/* 会员等级，0：店铺客户，1：普通会员，2：高级会员，3：VIP会员， 4：至尊VIP会员 */
	Grade int64 `json:"grade"`

	/* 分组的id集合字符串 */
	GroupIds string `json:"group_ids"`

	/* 购买的宝贝件数 */
	ItemNum int64 `json:"item_num"`

	/* 最后交易的日期 */
	LastTradeTime string `json:"last_trade_time"`

	/* 关系来源，1交易成功，2未交易成功 ,3 卖家主动吸纳 */
	RelationSource int64 `json:"relation_source"`

	/* 显示会员的状态，normal正常，blacklist黑名单 */
	Status string `json:"status"`

	/* 交易的金额 */
	TradeAmount float64 `json:"trade_amount,string"`

	/* 交易成功的次数 */
	TradeCount int64 `json:"trade_count"`
}

/* 卖家设置的等级优惠信息 */
type GradePromotion struct {
	/* 买家会员级别  0:店铺客户  1：普通会员 2：高级会员 3：VIP会员 4：至尊VIP */
	CurGrade string `json:"cur_grade"`

	/* 店铺客户、普通会员 、高级会员、VIP会员、至尊VIP */
	CurGradeName string `json:"cur_grade_name"`

	/* 会员级别折扣率没有小数，990代表9.9折 */
	Discount int64 `json:"discount"`

	/* 普通会员、高级会员、VIP会员、至尊VIP。空的时候代表后续等级未启用或没有下一等级 */
	NextGrade string `json:"next_grade"`

	/* 该等级对应的下一等级1:普通会员  2：高级会员 3：VIP会员 4：至尊VIP 0：后续等级都未启用或没有下一等级 */
	NextGradeName string `json:"next_grade_name"`

	/* 升级到下一个级别的需要的交易额，单位：分 */
	NextUpgradeAmount int64 `json:"next_upgrade_amount"`

	/* 升级到下一个级别的需要的交易量 */
	NextUpgradeCount int64 `json:"next_upgrade_count"`
}

/* 会员信息对象 */
type CrmMember struct {
	/* 平均客单价. */
	AvgPrice float64 `json:"avg_price,string"`

	/* 会员买家id */
	BuyerId int64 `json:"buyer_id"`

	/* 买家昵称 */
	BuyerNick string `json:"buyer_nick"`

	/* 城市.
	请注意:从2014.4.23之后,省市将采用地区标准码,请通过物流API taobao.areas.get接口获取,参考:http://api.taobao.com/apidoc/api.htm?spm=0.0.0.0.nOOF9g&path=cid:7-apiId:59.API对于老的省市代码兼容会逐步下线. */
	City string `json:"city"`

	/* 交易关闭的金额 */
	CloseTradeAmount float64 `json:"close_trade_amount,string"`

	/* 交易关闭的的笔数 */
	CloseTradeCount int64 `json:"close_trade_count"`

	/* 会员等级，0:店铺客户，1：普通会员，2：高级会员，3：VIP会员， 4：至尊VIP会员 */
	Grade int64 `json:"grade"`

	/* 会员拥有的所有分组 */
	GroupIds string `json:"group_ids"`

	/* 交易关闭的宝贝件数 */
	ItemCloseCount int64 `json:"item_close_count"`

	/* 购买的宝贝件数 */
	ItemNum int64 `json:"item_num"`

	/* 最后交易时间 */
	LastTradeTime string `json:"last_trade_time"`

	/* 北京=1,天津=2,河北省=3,山西省=4,内蒙古自治区=5,辽宁省=6,吉林省=7,黑龙江省=8,上海=9,江苏省=10,浙江省=11,安徽省=12,福建省=13,江西省=14,山东省=15,河南省=16,湖北省=17,湖南省=18, 广东省=19,广西壮族自治区=20,海南省=21,重庆=22,四川省=23,贵州省=24,云南省=25,西藏自治区26,陕西省=27,甘肃省=28,青海省=29,宁夏回族自治区=30,新疆维吾尔自治区=31,台湾省=32,香港特别行政区=33,澳门特别行政区=34,海外=35.
	注:请注意:从2014.4.23之后,省市将采用地区标准码,请通过物流API taobao.areas.get接口获取,参考:http://api.taobao.com/apidoc/api.htm?spm=0.0.0.0.nOOF9g&path=cid:7-apiId:59.API对于老的省市代码兼容会逐步下线. */
	Province int64 `json:"province"`

	/* 关系来源，1交易成功，2未成交，3卖家主动吸纳 */
	RelationSource int64 `json:"relation_source"`

	/* 显示会员的状态，normal正常，blacklist黑名单 */
	Status string `json:"status"`

	/* 交易成功的金额 */
	TradeAmount float64 `json:"trade_amount,string"`

	/* 交易成功笔数 */
	TradeCount int64 `json:"trade_count"`
}

/* 描述分组的数据结构 */
type Group struct {
	/* 分组的创建时间 */
	GroupCreate string `json:"group_create"`

	/* 分组的id */
	GroupId int64 `json:"group_id"`

	/* 分组的修改时间 */
	GroupModify string `json:"group_modify"`

	/* 分组的名称 */
	GroupName string `json:"group_name"`

	/* 分组所拥有的会员数量,如果返回值为-1，表示当前服务忙或数据在初始化。 */
	MemberCount int64 `json:"member_count"`

	/* 分组的状态，1表示正常 */
	Status string `json:"status"`
}

/* 授权关系 */
type WlbAuthorization struct {
	/* 授权结束时间 */
	AuthorizeEndTime string `json:"authorize_end_time"`

	/* 授权ID */
	AuthorizeId int64 `json:"authorize_id"`

	/* 授权开始时间 */
	AuthorizeStartTime string `json:"authorize_start_time"`

	/* 代销人用户ID */
	ConsignUserId int64 `json:"consign_user_id"`

	/* 授权商品ID */
	ItemId int64 `json:"item_id"`

	/* 授权名称 */
	Name string `json:"name"`

	/* 货主用户ID */
	OwnerUserId int64 `json:"owner_user_id"`

	/* 授权数量 */
	Quantity int64 `json:"quantity"`

	/* 授权编码 */
	RuleCode string `json:"rule_code"`

	/* 状态：
	VALID -- 1 有效
	INVALIDATION -- 2 失效 */
	Status string `json:"status"`
}

/* 通道消息 */
type WlbMessage struct {
	/* 创建时间 */
	GmtCreate string `json:"gmt_create"`

	/* 消息通道ID */
	Id int64 `json:"id"`

	/* 通知消息编码： STOCK_IN_NOT_CONSISTENT---入库单不一致 CANCEL_ORDER_SUCCESS---取消订单成功 CANCEL_ORDER_FAILURE---取消订单失败 INVENTORY_CHECK---盘点   INVENTORY_CHECK---盘点消息  ORDER_REJECT--wms拒单  ORDER_CONFIRMED--订单处理成功 WMS_FAILED--wms处理失败 */
	MsgCode string `json:"msg_code"`

	/* 通知内容：msg_code为STOCK_IN_NOT_CONSISTENT时,msg_content为:orderCode:code;orderItemId:111;itemId:123;planQuantity:10;relQuantity:6;
	msg_code为CANCEL_ORDER_SUCCESS及其它时,msg_content为：orderCode:code;
	msg_code为CANCEL_ORDER_SUCCESS及其它时,msg_content为：orderCode:code; msg_code为INVENTORY_CHECK时，msg_content为orderCode:code; */
	MsgContent string `json:"msg_content"`

	/* 消息描述 */
	MsgDescription string `json:"msg_description"`

	/* 消息状态： 不需要确认：NO_NEED_CONFIRM 已确认：CONFIRMED 待确认：TO_BE_CONFIRM */
	Status string `json:"status"`

	/* 用户ID */
	UserId int64 `json:"user_id"`
}

/* 物流宝代销关系 */
type WlbConsignMent struct {
	/* 代销关系id */
	Id int64 `json:"id"`

	/* 代销商用户前台宝贝id */
	ItemId int64 `json:"item_id"`

	/* 代销数量 */
	Number int64 `json:"number"`

	/* 供应商商品id */
	TgtItemId int64 `json:"tgt_item_id"`

	/* 供应商用户id */
	TgtUserId int64 `json:"tgt_user_id"`

	/* 代销商用户id */
	UserId int64 `json:"user_id"`
}

/* 物流宝补货统计对象 */
type WlbReplenish struct {
	/* 根据历史统计值计算出来的预期值
	WarnByLast3Days=1; ByLast3Days=3;
	ByLast10Days=10;
	ByLast30Days=30;
	ByLast90Days=90
	WarnByLast3Days为按照过去3天的发出的件数来统计的达到安全库存的天数；其它4项分别为按照过去3、10、30、90天发出的商品件数，现有库存可以维持的天数 */
	EstimateValue string `json:"estimate_value"`

	/* 订单历史统计值
	Last3DaysTotal=3; Last10DaysTotal=10;
	Last30DaysTotal=30; Last90DaysTotal=90
	分别为过去3、10、30、90天发出的商品件数 */
	HistoryValue string `json:"history_value"`

	/* 商品id */
	ItemId int64 `json:"item_id"`

	/* 补货周期（单位：天） */
	RetrievalCount int64 `json:"retrieval_count"`

	/* 可销售库存数 */
	SellCount int64 `json:"sell_count"`

	/* 仓库编码 */
	StoreCode string `json:"store_code"`

	/* 在途库存数 */
	TransportCount int64 `json:"transport_count"`

	/* 用户id */
	UserId int64 `json:"user_id"`

	/* 安全库存 */
	WarnCount int64 `json:"warn_count"`
}

/* 物流宝订单流转信息对象 */
type WlbProcessStatus struct {
	/* 状态内容 */
	Content string `json:"content"`

	/* 操作时间 */
	OperateTime string `json:"operate_time"`

	/* 操作人 */
	Operater string `json:"operater"`

	/* 物流宝订单编码 */
	OrderCode string `json:"order_code"`

	/* 备注 */
	Remark string `json:"remark"`

	/* 订单操作状态：WMS_ACCEPT;WMS_PRINT;WMS_PICK;WMS_CHECK;WMS_PACKAGE;WMS_CONSIGN;WMS_CANCEL;WMS_UNKNOWN;WMS_CONFIRMED
	TMS_ACCEPT;TMS_STATION_IN;TMS_STATION_OUT;TMS_SIGN;TMS_REJECT;TMS_CANCEL;TMS_UNKNOW;SYS_UNKNOWN */
	StatusCode string `json:"status_code"`

	/* 仓库编码 */
	StoreCode string `json:"store_code"`

	/* 仓库合作公司编码 */
	StoreTpCode string `json:"store_tp_code"`

	/* tms合作公司订单编码 */
	TmsOrderCode string `json:"tms_order_code"`

	/* tms合作公司编码 */
	TmsTpCode string `json:"tms_tp_code"`
}

/* 物流宝订单商品 */
type WlbOrderItem struct {
	/* 批次号备注 */
	BatchRemark string `json:"batch_remark"`

	/* 物流宝订单确认状态：
	NO_NEED_CONFIRM--不需要确认
	WAIT_CONFIRM--待确认
	CONFIRMED--已确认 */
	ConfirmStatus string `json:"confirm_status"`

	/* 外部实体id */
	ExtEntityId string `json:"ext_entity_id"`

	/* 外部实体类型 */
	ExtEntityType string `json:"ext_entity_type"`

	/* 第一位标示是否为赠品 */
	Flag int64 `json:"flag"`

	/* 订单商品id */
	Id int64 `json:"id"`

	/* INVENTORY_TYPE_SELL 可销库存
	INVENTORY_TYPE_IMPERFECTIONS 残次库存
	INVENTORY_TYPE_FREEZE 冻结库存
	INVENTORY_TYPE_ON_PASSAGE 在途库存
	INVENTORY_TYPE_ENGINE_DAMAGE 机损
	INVENTORY_TYPE_BOX_DAMAGE 箱损 */
	InventoryType string `json:"inventory_type"`

	/* 订单商品编码 */
	ItemCode string `json:"item_code"`

	/* 物流宝订单商品的物流宝商品id */
	ItemId int64 `json:"item_id"`

	/* 订单商品名称 */
	ItemName string `json:"item_name"`

	/* 商品价格 */
	ItemPrice int64 `json:"item_price"`

	/* 物流宝订单编码 */
	OrderCode string `json:"order_code"`

	/* 物流宝订单id */
	OrderId int64 `json:"order_id"`

	/* 子交易号 */
	OrderSub2code string `json:"order_sub_2code"`

	/* 订单号 */
	OrderSubCode string `json:"order_sub_code"`

	/* (1)其它: OTHER (2)淘宝交易: TAOBAO (3)调拨: ALLOCATION (4)盘点:CHECK (5)销售采购:PRUCHASE(6)其它交易：OTHER_TRADE */
	OrderSubType string `json:"order_sub_type"`

	/* 订单商品图片url */
	PictureUrl string `json:"picture_url"`

	/* 计划数量 */
	PlanQuantity int64 `json:"plan_quantity"`

	/* 货主id */
	ProviderTpId int64 `json:"provider_tp_id"`

	/* 货主nick */
	ProviderTpNick string `json:"provider_tp_nick"`

	/* 商品发布版本号 */
	PublishVersion int64 `json:"publish_version"`

	/* 实际数量 */
	RealQuantity int64 `json:"real_quantity"`

	/* 订单商品备注 */
	Remark string `json:"remark"`

	/* 订单商品用户id */
	UserId int64 `json:"user_id"`

	/* 订单商品用户昵称 */
	UserNick string `json:"user_nick"`
}

/* 系统自动生成 */
type QrcodeDO struct {
	/* 二维码所属渠道ID */
	ChannelId int64 `json:"channel_id"`

	/* 二维码对应的渠道名 */
	ChannelName string `json:"channel_name"`

	/* 二维码的矢量图下载地址，只有设置入参need_eps为true且style不为官方模板时，才返回值 */
	EpsUrl string `json:"eps_url"`

	/* 二维码id */
	QrcodeId int64 `json:"qrcode_id"`

	/* 二维码图片链接 */
	QrcodeUrl string `json:"qrcode_url"`

	/* 二维码的短地址，每一个原始地址都会生成一个短地址 */
	ShortUrl string `json:"short_url"`

	/* 二维码扫码后访问的目的地址 */
	Url string `json:"url"`
}

/* 物流宝订单对象 */
type WlbOrder struct {
	/* 买家nick */
	BuyerNick string `json:"buyer_nick"`

	/* 订单取消状态：
	1-取消中；
	2-取消失败；
	3-取消完成 */
	CancelOrderStatus int64 `json:"cancel_order_status"`

	/* 确认状态：
	(1)不需要确认：NO_NEED_CONFIRM
	(2)等待确认：WAIT_CONFIRM
	(3)已确认:CONFIRMED */
	ConfirmStatus string `json:"confirm_status"`

	/* 计划送达结束时间 */
	ExpectEndTime string `json:"expect_end_time"`

	/* 计划送达开始时间 */
	ExpectStartTime string `json:"expect_start_time"`

	/* 发票信息 */
	InvoiceInfo string `json:"invoice_info"`

	/* 物流宝订单中的商品种类数量 */
	ItemKindsCount int64 `json:"item_kinds_count"`

	/* 出库或者入库，IN表示入库，OUT表示出库 */
	OperateType string `json:"operate_type"`

	/* 订单编码 */
	OrderCode string `json:"order_code"`

	/* 第1位:COD,2:限时配送,3:预售,4:需要发票,5:已投诉,第6位:合单,第7位:拆单 第8位：EXCHANGE-换货， 第9位:VISIT-上门 ， 第10位: MODIFYTRANSPORT-是否可改配送方式，第11位：是否物流代理确认发货 */
	OrderFlag int64 `json:"order_flag"`

	/* 订单来源:
	产生物流订单的原因，比如:

	订单来源:1:TAOBAO;2:EXT;3:ERP;4:WMS */
	OrderSource string `json:"order_source"`

	/* 对应创建物流宝订单top接口中的的out_biz_code字段，主要是用来去重用 */
	OrderSourceCode string `json:"order_source_code"`

	/* 物流状态，
	订单已创建：0
	订单已取消: -1
	订单关闭:-3
	下发中: 10
	已下发: 11
	接收方拒签 :-100
	已发货:100
	签收成功:200 */
	OrderStatus string `json:"order_status"`

	/* 卖家取消,仓库取消标识 */
	OrderStatusReason string `json:"order_status_reason"`

	/* (1)其它:    OTHER
	(2)淘宝交易: TAOBAO
	(3)301:调拨: ALLOCATION
	(4)401:盘点:CHECK
	(5)501:销售采购:PRUCHASE */
	OrderSubType string `json:"order_sub_type"`

	/* 1:正常订单: NARMAL
	2:退货订单: RETURN
	3:换货订单: CHANGE */
	OrderType string `json:"order_type"`

	/* 原订单编码 */
	PrevOrderCode string `json:"prev_order_code"`

	/* 实际入库的种类数量 */
	RealKindsCount int64 `json:"real_kinds_count"`

	/* 应收金额 */
	ReceivableAmount int64 `json:"receivable_amount"`

	/* 收件人具体地址 */
	ReceiverAddress string `json:"receiver_address"`

	/* 区或者县 */
	ReceiverArea string `json:"receiver_area"`

	/* 收件人城市 */
	ReceiverCity string `json:"receiver_city"`

	/* 接收人电子邮箱 */
	ReceiverMail string `json:"receiver_mail"`

	/* 接收人手机号码 */
	ReceiverMobile string `json:"receiver_mobile"`

	/* 接收人姓名 */
	ReceiverName string `json:"receiver_name"`

	/* 接收人固定电话 */
	ReceiverPhone string `json:"receiver_phone"`

	/* 收件人省份 */
	ReceiverProvince string `json:"receiver_province"`

	/* 接收人旺旺 */
	ReceiverWangwang string `json:"receiver_wangwang"`

	/* 收件人邮编 */
	ReceiverZipCode string `json:"receiver_zip_code"`

	/* 订单备注 */
	Remark string `json:"remark"`

	/* 发货日期:
	(1)1 为工作日
	(2)2 为节假日 */
	ScheduleDay string `json:"schedule_day"`

	/* 配送结束时间通常是HH:MM格式 */
	ScheduleEnd string `json:"schedule_end"`

	/* 发货速度  ，
	101-当日达，
	102-次晨达，
	103-次日达 */
	ScheduleSpeed int64 `json:"schedule_speed"`

	/* 配送开始时间通常是HH:MM格式 */
	ScheduleStart string `json:"schedule_start"`

	/* 发件人地址 */
	SenderAddress string `json:"sender_address"`

	/* 发件人所在区 */
	SenderArea string `json:"sender_area"`

	/* 发件人城市 */
	SenderCity string `json:"sender_city"`

	/* 发件人电子邮箱 */
	SenderEmail string `json:"sender_email"`

	/* 发件人移动电话 */
	SenderMobile string `json:"sender_mobile"`

	/* 发件人姓名 */
	SenderName string `json:"sender_name"`

	/* 发件人联系电话 */
	SenderPhone string `json:"sender_phone"`

	/* 发件人所在省份 */
	SenderProvince string `json:"sender_province"`

	/* 发件人邮编 */
	SenderZipCode string `json:"sender_zip_code"`

	/* cod服务费 */
	ServiceFee int64 `json:"service_fee"`

	/* 物流运输方式：
	MAIL-平邮
	EXPRESS-快递
	EMS-EMS
	OTHER-其他 */
	ShippingType string `json:"shipping_type"`

	/* 仓库编码 */
	StoreCode string `json:"store_code"`

	/* 发货物流公司编码，STO,YTO,EMS等 */
	TmsTpCode string `json:"tms_tp_code"`

	/* 订单总价 */
	TotalAmount int64 `json:"total_amount"`

	/* 卖家ID */
	UserId int64 `json:"user_id"`

	/* 卖家NICK */
	UserNick string `json:"user_nick"`
}

/* 卖家定购的服务 */
type WlbSellerSubscription struct {
	/* 定购有效结束日期 */
	EndDate string `json:"end_date"`

	/* 创建时间 */
	GmtCreate string `json:"gmt_create"`

	/* 修改时间 */
	GmtModified string `json:"gmt_modified"`

	/* 定购ID */
	Id int64 `json:"id"`

	/* 判断该仓库是否是实体仓，还是虚拟仓，null是实体仓，10:代表虚拟仓 */
	IsOwnService int64 `json:"is_own_service"`

	/* 父定购ID
	可通过该字段来得之服务上下级关系。
	例定购了仓储服务，下又有TMS服务。
	该字段保存仓储服务ID。 */
	ParentId int64 `json:"parent_id"`

	/* 服务商ID */
	ProviderUserId int64 `json:"provider_user_id"`

	/* 备注 */
	Remark string `json:"remark"`

	/* 自有仓的别名，当当前订购记录为自有仓时才会有值 */
	ServiceAlias string `json:"service_alias"`

	/* 服务编码 */
	ServiceCode string `json:"service_code"`

	/* 服务ID */
	ServiceId int64 `json:"service_id"`

	/* 服务名 */
	ServiceName string `json:"service_name"`

	/* 服务类型，
	STORE 1-仓储、
	TMS 2-TMS、
	PACKAGE 3-包装服务
	SUPPLIER 4-供货
	INSTALL 5-安装
	COMPLEX_SERVICE 100-综合服务 */
	ServiceType string `json:"service_type"`

	/* 定购有效开始日期 */
	StartDate string `json:"start_date"`

	/* 状态
	AUDITING 1-待审核
	CANCEL 2-撤销
	CHECKED 3-审核通过
	FAILED 4-审核未通过
	SYNCHRONIZING 5-同步中 */
	Status string `json:"status"`

	/* 定购用户ID */
	SubscriberUserId int64 `json:"subscriber_user_id"`

	/* 定购用户NICK */
	SubscriberUserNick string `json:"subscriber_user_nick"`

	/* 联系人地址信息 */
	WlbPartnerAddress *WlbPartnerAddress `json:"wlb_partner_address"`

	/* 联系人联系详情 */
	WlbPartnerContact *WlbPartnerContact `json:"wlb_partner_contact"`
}

/* 库存变更记录 */
type WlbItemInventoryLog struct {
	/* 批次号 */
	BatchCode string `json:"batch_code"`

	/* 创建日期 */
	GmtCreate string `json:"gmt_create"`

	/* 库存变更ID */
	Id int64 `json:"id"`

	/* VENDIBLE  1-可销售;
	FREEZE  201-冻结库存;
	ONWAY  301-在途库存;
	DEFECT  101-残存品;
	ENGINE_DAMAGE 102-机损;
	BOX_DAMAGE 103-箱损 */
	InventType string `json:"invent_type"`

	/* 商品ID */
	ItemId int64 `json:"item_id"`

	/* 库存操作作类型
	CHU_KU 1-出库
	RU_KU 2-入库
	FREEZE 3-冻结
	THAW 4-解冻
	CHECK_FREEZE 5-冻结确认
	CHANGE_KU 6-库存类型变更 */
	OpType string `json:"op_type"`

	/* 库存操作者ID */
	OpUserId int64 `json:"op_user_id"`

	/* 订单号 */
	OrderCode string `json:"order_code"`

	/* 订单商品ID */
	OrderItemId int64 `json:"order_item_id"`

	/* 处理数量变化值 */
	Quantity int64 `json:"quantity"`

	/* 备注 */
	Remark string `json:"remark"`

	/* 结果值 */
	ResultQuantity int64 `json:"result_quantity"`

	/* 仓库编码 */
	StoreCode string `json:"store_code"`

	/* 用户ID */
	UserId int64 `json:"user_id"`
}

/* 外部商品实体 */
type OutEntityItem struct {
	/* entity_type对应的实体类型的id：
	当entity_type为IC_ITEM时，entity_id为ic的商品id */
	EntityId string `json:"entity_id"`

	/* 外部实体类型：
	IC_ITEM--ic商品
	IC_SKU--ic销售单元 */
	EntityType string `json:"entity_type"`
}

/* 物流订单运单信息 */
type WlbTmsOrder struct {
	/* 运单号 */
	Code string `json:"code"`

	/* 物流公司编码 */
	CompanyCode string `json:"company_code"`

	/* 物流订单编号 */
	OrderCode string `json:"order_code"`

	/* 物流订单的所有者ID,货主 */
	UserId int64 `json:"user_id"`
}

/* 用户订购信息 */
type ArticleUserSubscribe struct {
	/* 订购关系到期时间 */
	Deadline string `json:"deadline"`

	/* 收费项目代码，从合作伙伴后台（my.open.taobao.com）-收费管理-收费项目列表 能够获得收费项目代码 */
	ItemCode string `json:"item_code"`
}

/* SPU发布模板，定义了产品发布需要那些关键属性，绑定属性。 */
type SpuTemplateDO struct {
	/* 产品绑定属性，内容为属性ID(PID)的列表,绑定属性肯定在类目上有，对应属性的类目特征，子属性请根据PID到类目上去取 */
	AffectProperties []int64 `json:"affect_properties"`

	/* 类目ID */
	CategoryId int64 `json:"category_id"`

	/* 品类ID，和类目ID类似 */
	CommodityId int64 `json:"commodity_id"`

	/* 过滤属性，内容有属性ID(PID)列表，很重要的属性，filter_properties包含的属性，必须填写 */
	FilterProperties []int64 `json:"filter_properties"`

	/* 产品关键属性，内容为属性ID(PID)的列表，注意关键属性可以在类目上不存在。不存在的PID，默认为输入，没有子属性。属性名称在prop_name_str中取 */
	KeyProperties []int64 `json:"key_properties"`

	/* 预留 */
	PropFeatures string `json:"prop_features"`

	/* 属性名称扁平化结构，只保证不在类目上的CP有值 */
	PropNameStr string `json:"prop_name_str"`

	/* 模板ID，发布产品，必须放到Product中 */
	TemplateId int64 `json:"template_id"`
}

/* 应用订单信息 */
type ArticleBizOrder struct {
	/* 应用收费代码，从合作伙伴后台（my.open.taobao.com）-收费管理-收费项目列表 能够获得该应用的收费代码 */
	ArticleCode string `json:"article_code"`

	/* 商品模型名称 */
	ArticleItemName string `json:"article_item_name"`

	/* 应用名称 */
	ArticleName string `json:"article_name"`

	/* 订单号 */
	BizOrderId int64 `json:"biz_order_id"`

	/* 订单类型，1=新订 2=续订 3=升级 4=后台赠送 5=后台自动续订 6=订单审核后生成订购关系（暂时用不到） */
	BizType int64 `json:"biz_type"`

	/* 订单创建时间（订购时间） */
	Create string `json:"create"`

	/* 原价（单位为分） */
	Fee string `json:"fee"`

	/* 收费项目代码，从合作伙伴后台（my.open.taobao.com）-收费管理-收费项目列表 能够获得收费项目代码 */
	ItemCode string `json:"item_code"`

	/* 收费项目名称 */
	ItemName string `json:"item_name"`

	/* 淘宝会员名 */
	Nick string `json:"nick"`

	/* 订购周期  1表示年 ，2表示月，3表示天。 */
	OrderCycle string `json:"order_cycle"`

	/* 订购周期结束时间 */
	OrderCycleEnd string `json:"order_cycle_end"`

	/* 订购周期开始时间 */
	OrderCycleStart string `json:"order_cycle_start"`

	/* 子订单号 */
	OrderId int64 `json:"order_id"`

	/* 优惠（单位为分） */
	PromFee string `json:"prom_fee"`

	/* 退款（单位为分；升级时，系统会将升级前老版本按照剩余订购天数退还剩余金额） */
	RefundFee string `json:"refund_fee"`

	/* 实付（单位为分） */
	TotalPayFee string `json:"total_pay_fee"`
}

/* 分销商等级 */
type FenxiaoGrade struct {
	/* 记录创建时间 */
	Created string `json:"created"`

	/* 主键 */
	GradeId int64 `json:"grade_id"`

	/* 记录最后修改时间 */
	Modified string `json:"modified"`

	/* 分销商等级名称 */
	Name string `json:"name"`
}

/* 套装配置 */
type SuiteConfDO struct {
	/* 如果有值，套装子规格形如：A*n,B*n,C*n....，n只能是inputs中的 */
	Inputs []int64 `json:"inputs"`

	/* 套装，子规格限制数，当前一级类目，最多允许选择多少个产品规格，为空表示无限制。套装子规格形如:A*n,B*n,C*n....,max_size就是C的最大个数 */
	MaxSize int64 `json:"max_size"`

	/* 套装，子规格数量总和限制，当前一级类目，所以子产品规格的数量不能超过这个值，为空表示无限制,套装子规格形如：A*n,B*n,C*n....,max_total_number就是所有n和的最大限制 */
	MaxTotalNumber int64 `json:"max_total_number"`

	/* 一级类目ID，标识在该一级在类目下能布套装 */
	RootCatId int64 `json:"root_cat_id"`

	/* 套装类目ID，当前一级类目，如果发布跨叶子类目的的套装，默认放到那个类目。如果没有值表示root_cat_id类目下发布的套装产品，不支持跨类目 */
	SuiteCatId int64 `json:"suite_cat_id"`
}

/* 提示信息对象 */
type TipInfo struct {
	/* 提示信息 */
	Info string `json:"info"`

	/* 后端商品ID或者商家编码 */
	ScItemId string `json:"sc_item_id"`
}

/* 仓库对象 */
type Store struct {
	/* 仓库的物理地址 */
	Address string `json:"address"`

	/* 仓库对应的淘宝区域 */
	AddressAreaName string `json:"address_area_name"`

	/* 仓库简称 */
	AliasName string `json:"alias_name"`

	/* 联系人 */
	Contact string `json:"contact"`

	/* 联系电话 */
	Phone string `json:"phone"`

	/* 邮编 */
	PostCode string `json:"post_code"`

	/* 商家的仓库编码，不允许重复 */
	StoreCode string `json:"store_code"`

	/* 商家的仓库名称 */
	StoreName string `json:"store_name"`

	/* 仓库类型 */
	StoreType string `json:"store_type"`
}

/* 商品库存对象 */
type InventorySum struct {
	/* 库存类型：
	1：正常
	2：损坏
	3：冻结
	10：质押
	11-20:商家自定义 */
	InventoryType int64 `json:"inventory_type"`

	/* 库存类型名称 */
	InventoryTypeName string `json:"inventory_type_name"`

	/* 总占用数量 */
	OccupyQuantity int64 `json:"occupy_quantity"`

	/* 总物理库存数量 */
	Quantity int64 `json:"quantity"`

	/* 总预扣数量 */
	ReserveQuantity int64 `json:"reserve_quantity"`

	/* 商品商家编码 */
	ScItemCode string `json:"sc_item_code"`

	/* 商品后端ID，如果有传sc_item_code,参数可以为0 */
	ScItemId int64 `json:"sc_item_id"`

	/* 商家仓库编码 */
	StoreCode string `json:"store_code"`
}

/* 联系人联系详情 */
type WlbPartnerContact struct {
	/* 仓库联系人姓名 */
	Name string `json:"name"`

	/* 联系电话 */
	Phone string `json:"phone"`
}

/* 联系人地址信息 */
type WlbPartnerAddress struct {
	/* 详细地址 */
	Address string `json:"address"`

	/* 区 */
	Borough string `json:"borough"`

	/* 市 */
	City string `json:"city"`

	/* 省 */
	Province string `json:"province"`

	/* 邮编 */
	Zip string `json:"zip"`
}

/* 搭配套餐类。 */
type Meal struct {
	/* 搭配套餐商品列表。item_id为商品的id;item_show_name为商品显示名。因最多允许5个商品进行搭配，所以查询最多有5个，以json格式传出。 */
	ItemList string `json:"item_list"`

	/* 套餐id。 */
	MealId int64 `json:"meal_id"`

	/* 搭配套餐描述！ */
	MealMemo string `json:"meal_memo"`

	/* 搭配套餐名称。 */
	MealName string `json:"meal_name"`

	/* 套餐一口价(单位是：分) */
	MealPrice float64 `json:"meal_price,string"`

	/* 普通运费模板id。若这个字段为空或0时，运费是卖家负责;若这个字段不为空，说明运费模板存在，运费是买家负责。 */
	PostageId int64 `json:"postage_id"`

	/* 套餐状态。有效：VALID;失效：INVALID(有效套餐为可使用的套餐,无效套餐为套餐中有商品下架或库存为0时)。 */
	Status string `json:"status"`
}

/* 应用订购信息 */
type ArticleSub struct {
	/* 应用收费代码，从合作伙伴后台（my.open.taobao.com）-收费管理-收费项目列表 能够获得该应用的收费代码 */
	ArticleCode string `json:"article_code"`

	/* 应用名称 */
	ArticleName string `json:"article_name"`

	/* 是否自动续费 */
	Autosub bool `json:"autosub"`

	/* 订购关系到期时间 */
	Deadline string `json:"deadline"`

	/* 是否到期提醒 */
	ExpireNotice bool `json:"expire_notice"`

	/* 收费项目代码，从合作伙伴后台（my.open.taobao.com）-收费管理-收费项目列表 能够获得收费项目代码 */
	ItemCode string `json:"item_code"`

	/* 收费项目名称 */
	ItemName string `json:"item_name"`

	/* 淘宝会员名 */
	Nick string `json:"nick"`

	/* 状态，1=有效 2=过期 */
	Status int64 `json:"status"`
}

/* 等级折扣数据结构 */
type GradeDiscount struct {
	/* 等级ID或分销商ID */
	DiscountId int64 `json:"discount_id"`

	/* 折扣类型：1-等级、2-分销商折扣 */
	DiscountType int64 `json:"discount_type"`

	/* 采购价格 */
	Price float64 `json:"price,string"`

	/* skuId */
	SkuId int64 `json:"sku_id"`

	/* 模式：1-代销、2-经销 */
	TradeType int64 `json:"trade_type"`
}

/* 合作申请 */
type Requisition struct {
	/* 好评率 */
	DistAppraise int64 `json:"dist_appraise"`

	/* 主营类目 */
	DistCategory int64 `json:"dist_category"`

	/* 主营类目名称 */
	DistCategoryName string `json:"dist_category_name"`

	/* 是否消保(0-不是、1-是) */
	DistIsXiaobao int64 `json:"dist_is_xiaobao"`

	/* 店铺星级 */
	DistLevel int64 `json:"dist_level"`

	/* 分销申请加盟时，给供应商的留言 */
	DistMessage string `json:"dist_message"`

	/* 开店时间 */
	DistOpenDate string `json:"dist_open_date"`

	/* 店铺地址 */
	DistShopAddress string `json:"dist_shop_address"`

	/* 分销商id */
	DistributorId int64 `json:"distributor_id"`

	/* 分销商nick */
	DistributorNick string `json:"distributor_nick"`

	/* 申请时间 */
	GmtCreate string `json:"gmt_create"`

	/* 合作申请ID */
	RequisitionId int64 `json:"requisition_id"`

	/* 申请状态（1-申请中、2-成功、3-被退回、4-已撤消、5-过期） */
	Status int64 `json:"status"`
}

/* 后端商品映射关系对象 */
type ScItemMap struct {
	/* map_type=1时，item_id为IC商品id

	map_type=7时，item_id为分销商品id */
	ItemId int64 `json:"item_id"`

	/* 1:前台和后台关系
	7:分销和后台关系 */
	MapType int64 `json:"map_type"`

	/* 后端商品ID */
	RelItemId int64 `json:"rel_item_id"`

	/* 后端商品所有者商家编码 */
	RelOuterCode string `json:"rel_outer_code"`

	/* 后端商品所有者id */
	RelUserId int64 `json:"rel_user_id"`

	/* 后端商品所有者名称 */
	RelUserNick string `json:"rel_user_nick"`

	/* 当宝贝下没SKU时该字段为空 */
	SkuId int64 `json:"sku_id"`

	/* Ic商家id(分销商id) */
	UserId int64 `json:"user_id"`

	/* Ic商家nick(分销商nick) */
	UserNick string `json:"user_nick"`
}

/* 查询分页结构 */
type QueryPagination struct {
	/* 当前页码数 */
	PageIndex int64 `json:"page_index"`

	/* 分页记录个数，如果用户输入的记录数大于50，则一页显示50条记录 */
	PageSize int64 `json:"page_size"`
}

/* 后端商品 */
type ScItem struct {
	/* 条形码 */
	BarCode string `json:"bar_code"`

	/* 品牌id */
	BrandId int64 `json:"brand_id"`

	/* 品牌名称 */
	BrandName string `json:"brand_name"`

	/* 高 单位：mm */
	Height int64 `json:"height"`

	/* 1表示区域销售，0或是空是非区域销售 */
	IsAreaSale int64 `json:"is_area_sale"`

	/* 贵重品:false:不是 true：是 */
	IsCostly bool `json:"is_costly"`

	/* 是否危险 false：不是  true：是 */
	IsDangerous bool `json:"is_dangerous"`

	/* 是否易碎 false ：不是  true：是 */
	IsFriable bool `json:"is_friable"`

	/* 是否保质期：false:不是 true：是 */
	IsWarranty bool `json:"is_warranty"`

	/* 商品id */
	ItemId int64 `json:"item_id"`

	/* 商品名称 */
	ItemName string `json:"item_name"`

	/* 1.普通供应链商品 2.供应链组合主商品 */
	ItemType int64 `json:"item_type"`

	/* 长度 单位：mm */
	Length int64 `json:"length"`

	/* LIQUID:液体，1：粉体，SOLID：固体 */
	MatterStatus string `json:"matter_status"`

	/* 后端商品options字段 */
	Options int64 `json:"options"`

	/* 商家编码 */
	OuterCode string `json:"outer_code"`

	/* 价格：分（吊牌价） */
	Price int64 `json:"price"`

	/* 商品属性格式是  p1:v1,p2:v2,p3:v3 */
	Properties string `json:"properties"`

	/* 备注 */
	Remark string `json:"remark"`

	/* 体积：立方厘米 */
	Volume int64 `json:"volume"`

	/* 重量.单位：克 */
	Weight int64 `json:"weight"`

	/* 宽 单位：mm */
	Width int64 `json:"width"`

	/* 仓储商编码，可以支持多个，格式wmsname:code */
	WmsCode string `json:"wms_code"`
}

/* 收\发货地址 */
type WaybillAddress struct {
	/* 详细地址 */
	AddressDetail string `json:"address_detail"`

	/* -- */
	AddressFormat string `json:"address_format"`

	/* 区名称（三级地址） */
	Area string `json:"area"`

	/* 三级地址代码 */
	AreaCode int64 `json:"area_code"`

	/* 市名称（二级地址） */
	City string `json:"city"`

	/* 二级地址代码 */
	CityCode int64 `json:"city_code"`

	/* 末级地址 */
	DivisionId int64 `json:"division_id"`

	/* 省名称（一级地址） */
	Province string `json:"province"`

	/* 一级地址代码 */
	ProvinceCode int64 `json:"province_code"`

	/* 街道\镇名称（四级地址） */
	Town string `json:"town"`

	/* 四级地址代码 */
	TownCode int64 `json:"town_code"`

	/* waybill 地址记录ID(非地址库ID) */
	WaybillAddressId int64 `json:"waybill_address_id"`
}

/* 订单数据 */
type TradeOrderInfo struct {
	/* 是否阿里系订单 */
	AliOrder bool `json:"ali_order"`

	/* 收货人地址 */
	ConsigneeAddress *WaybillAddress `json:"consignee_address"`

	/* 收货人 */
	ConsigneeName string `json:"consignee_name"`

	/* 收货人联系方式 */
	ConsigneePhone string `json:"consignee_phone"`

	/* 商品名称 */
	ItemName string `json:"item_name"`

	/* 物流服务能力集合 */
	LogisticsServiceList []*LogisticsService `json:"logistics_service_list>logistics_service"`

	/* 订单渠道 */
	OrderChannelsType string `json:"order_channels_type"`

	/* 订单渠道来源 */
	OrderType int64 `json:"order_type"`

	/* 包裹中的商品类型 */
	PackageItems []*PackageItem `json:"package_items>package_item"`

	/* 快递服务产品类型编码 */
	ProductType string `json:"product_type"`

	/* 使用者ID */
	RealUserId int64 `json:"real_user_id"`

	/* 发货人姓名 */
	SendName string `json:"send_name"`

	/* 发货人联系方式 */
	SendPhone string `json:"send_phone"`

	/* 大头笔 */
	ShortAddress string `json:"short_address"`

	/* 交易订单列表 */
	TradeOrderList []string `json:"trade_order_list"`

	/* 包裹体积（立方厘米） */
	Volume int64 `json:"volume"`

	/* 面单号 */
	WaybillCode string `json:"waybill_code"`

	/* 包裹重量（克） */
	Weight int64 `json:"weight"`
}

/* 宝贝详情页面信息 */
type ItemTemplate struct {
	/* 用于区分宝贝模板属于内店和外店 */
	ShopType int64 `json:"shop_type"`

	/* 宝贝模板的id */
	TemplateId int64 `json:"template_id"`

	/* 宝贝详情模板的名称 */
	TemplateName string `json:"template_name"`
}

/* 申请面单返回数据 */
type WaybillApplyNewInfo struct {
	/* 包裹对应的派件（收件）物流服务商网点（分支机构）代码 */
	ConsigneeBranchCode string `json:"consignee_branch_code"`

	/* 包裹对应的派件（收件）物流服务商网点（分支机构）名称 */
	ConsigneeBranchName string `json:"consignee_branch_name"`

	/* 预留字段 */
	Feature string `json:"feature"`

	/* -- */
	Result int64 `json:"result"`

	/* 面单号对应的物流服务商网点（分支机构）代码 */
	ShippingBranchCode string `json:"shipping_branch_code"`

	/* 面单号对于的物流服务商网点（分支机构）名称 */
	ShippingBranchName string `json:"shipping_branch_name"`

	/* 根据收货地址返回大头笔信息 */
	ShortAddress string `json:"short_address"`

	/* 面单对应的订单列 */
	TradeOrderInfo *TradeOrderInfo `json:"trade_order_info"`

	/* 返回的面单号 */
	WaybillCode string `json:"waybill_code"`
}

/* 分销商品下载记录 */
type FenxiaoItemRecord struct {
	/* 下载时间 */
	Created string `json:"created"`

	/* 分销商ID */
	DistributorId int64 `json:"distributor_id"`

	/* 商品ID */
	ItemId int64 `json:"item_id"`

	/* 更新时间（系统时间，无业务意义） */
	Modified string `json:"modified"`

	/* 产品ID */
	ProductId int64 `json:"product_id"`

	/* 分销方式：AGENT（只做代销，默认值）、DEALER（只做经销） */
	TradeType string `json:"trade_type"`
}

/* 合作分销关系 */
type Cooperation struct {
	/* 供应商授权的支付方式：ALIPAY(支付宝)、OFFPREPAY(预付款)、OFFTRANSFER(转帐)、OFFSETTLEMENT(后期结算) */
	AuthPayway []string `json:"auth_payway"`

	/* 合作关系ID */
	CooperateId int64 `json:"cooperate_id"`

	/* 分销商ID */
	DistributorId int64 `json:"distributor_id"`

	/* 分销商nick */
	DistributorNick string `json:"distributor_nick"`

	/* 合作终止时间 */
	EndDate string `json:"end_date"`

	/* 等级ID */
	GradeId int64 `json:"grade_id"`

	/* 授权产品线 */
	ProductLine string `json:"product_line"`

	/* 授权产品线名称，和product_line中的值按序对应 */
	ProductLineName []string `json:"product_line_name"`

	/* 合作起始时间 */
	StartDate string `json:"start_date"`

	/* 合作状态： NORMAL(合作中)、 ENDING(终止中) 、END (终止) */
	Status string `json:"status"`

	/* 供应商ID */
	SupplierId int64 `json:"supplier_id"`

	/* 供应商NICK */
	SupplierNick string `json:"supplier_nick"`

	/* 分销方式： AGENT(代销) 、DEALER(经销) */
	TradeType string `json:"trade_type"`
}

/* 限时打折 */
type LimitDiscount struct {
	/* 限时打折结束时间。 */
	EndTime string `json:"end_time"`

	/* 该限时打折宝贝数量。 */
	ItemNum int64 `json:"item_num"`

	/* 限时打折ID。 */
	LimitDiscountId int64 `json:"limit_discount_id"`

	/* 限时打折名称。 */
	LimitDiscountName string `json:"limit_discount_name"`

	/* 限时打折开始时间。 */
	StartTime string `json:"start_time"`
}

/* 地址库返回数据信息 */
type AddressResult struct {
	/* 详细街道地址，不需要重复填写省/市/区 */
	Addr string `json:"addr"`

	/* 区域ID */
	AreaId int64 `json:"area_id"`

	/* 是否默认退货地址 */
	CancelDef bool `json:"cancel_def"`

	/* 市 */
	City string `json:"city"`

	/* 地址库ID */
	ContactId int64 `json:"contact_id"`

	/* 联系人姓名 */
	ContactName string `json:"contact_name"`

	/* 区、县 */
	Country string `json:"country"`

	/* 是否默认取货地址 */
	GetDef bool `json:"get_def"`

	/* 备注 */
	Memo string `json:"memo"`

	/* 手机号码，手机与电话必需有一个
	手机号码不能超过20位 */
	MobilePhone string `json:"mobile_phone"`

	/* 修改日期时间 */
	ModifyDate string `json:"modify_date"`

	/* 电话号码,手机与电话必需有一个 */
	Phone string `json:"phone"`

	/* 省 */
	Province string `json:"province"`

	/* 公司名称, */
	SellerCompany string `json:"seller_company"`

	/* 是否默认发货地址 */
	SendDef bool `json:"send_def"`

	/* 地区邮政编码 */
	ZipCode string `json:"zip_code"`
}

/* 限时打折详情 */
type LimitDiscountDetail struct {
	/* 限时打折结束时间。 */
	EndTime string `json:"end_time"`

	/* 该商品限时折扣 */
	ItemDiscount float64 `json:"item_discount,string"`

	/* 商品的id(数字类型) */
	ItemId int64 `json:"item_id"`

	/* 限时打折名称 */
	LimitDiscountName string `json:"limit_discount_name"`

	/* 每人限购数量，1、2、5、10000(不限)。 */
	LimitNum int64 `json:"limit_num"`

	/* 限时打折开始时间。 */
	StartTime string `json:"start_time"`
}

/* 客服未回复统计 */
type NonreplyStatById struct {
	/* 客服人员未回复的客户ID */
	NonReplyCustomId string `json:"non_reply_customId"`

	/* 客服未回复数 */
	NonReplyNum int64 `json:"non_reply_num"`

	/* 客服人员ID */
	ServiceStaffId string `json:"service_staff_id"`
}

/* 客服评价统计 */
type StaffEvalStatById struct {
	/* 客服评价 */
	Evaluations []*Evaluation `json:"evaluations>evaluation"`

	/* 客服人员ID */
	ServiceStaffId string `json:"service_staff_id"`
}

/* 商家与CP的订购关系 */
type WaybillApplySubscriptionInfo struct {
	/* CP网点信息及对应的商家的发货信息 */
	BranchAccountCols []*WaybillBranchAccount `json:"branch_account_cols>waybill_branch_account"`

	/* 物流服务商ID */
	CpCode string `json:"cp_code"`

	/* 1是直营，2是加盟 */
	CpType int64 `json:"cp_type"`

	/* 订购关系是否建立 */
	Result int64 `json:"result"`
}

/* CP下的网点信息 */
type WaybillBranchAccount struct {
	/* 已用单数 */
	AllocatedQuantity int64 `json:"allocated_quantity"`

	/* 网点ID */
	BranchCode string `json:"branch_code"`

	/* 取消的面对总数 */
	CancelQuantity int64 `json:"cancel_quantity"`

	/* 物流服务商ID */
	CpCode string `json:"cp_code"`

	/* 已经打印的面单总数 */
	PrintQuantity int64 `json:"print_quantity"`

	/* 可用单数 */
	Quantity int64 `json:"quantity"`

	/* 商家ID */
	SellerId int64 `json:"seller_id"`

	/* 当前网点下的发货地址 */
	ShippAddressCols []*WaybillAddress `json:"shipp_address_cols>waybill_address"`
}

/* 更新面单数据 */
type WaybillApplyUpdateInfo struct {
	/* 收货网点信息 */
	ConsigneeBranchCode string `json:"consignee_branch_code"`

	/* 收货网点编码 */
	ConsigneeBranchName string `json:"consignee_branch_name"`

	/* -- */
	Desc string `json:"desc"`

	/* -- */
	Result int64 `json:"result"`

	/* 挑拣规则（大头笔信息） */
	ShortAddress string `json:"short_address"`

	/* -- */
	TradeOrderInfo *TradeOrderInfo `json:"trade_order_info"`

	/* -- */
	WaybillCode string `json:"waybill_code"`
}

/* 拍卖商品相关信息 */
type PaimaiInfo struct {
	/* 用户自定义的固定保证金。如果用户未传则说明用户选择使用淘宝默认的保证金模式10%，此时deposit等于0. */
	Deposit int64 `json:"deposit"`

	/* 降价拍中的降价间隔 */
	Interval int64 `json:"interval"`

	/* 拍卖类型，目前包括增加拍，荷兰拍和降价拍。 */
	Mode int64 `json:"mode"`

	/* 降价拍的保留价 */
	Reserve float64 `json:"reserve,string"`

	/* 对于拍卖来说可以设定有效期，这里是有效期的小时数。 */
	ValidHour int64 `json:"valid_hour"`

	/* 对于拍卖来说可以设定有效期，这里是有效期的分钟数。 */
	ValidMinute int64 `json:"valid_minute"`
}

/* 发货时间数据结构 */
type DeliveryTime struct {
	/* 商品级别设置的发货时间。设置了商品级别的发货时间，相对发货时间，则填写相对发货时间的天数（大于3）；绝对发货时间，则填写yyyy-mm-dd格式，如2013-11-11 */
	DeliveryTime string `json:"delivery_time"`

	/* 发货时间类型：绝对发货时间或者相对发货时间 */
	DeliveryTimeType string `json:"delivery_time_type"`

	/* 设置是否使用发货时间，商品级别，sku级别 */
	NeedDeliveryTime string `json:"need_delivery_time"`
}

/* 库存详情对象。其中包括货主ID，仓库编码，库存，库存类型等属性 */
type WlbInventory struct {
	/* 商品ID */
	ItemId int64 `json:"item_id"`

	/* 冻结(锁定)数量，用来跟踪库存的中间状态，比如前台销售了1件商品，这时lock加1，当商品出库的时候lock再减回去 */
	LockQuantity int64 `json:"lock_quantity"`

	/* 库存数量(有效数量) */
	Quantity int64 `json:"quantity"`

	/* 仓库编码，关联到仓库类型服务的编码非托管库存(卖家自己管理的库存，物流宝不可见又称自有库存)的所在仓库编码: STORE_SYS_PRIVATE */
	StoreCode string `json:"store_code"`

	/* VENDIBLE--可销售库存
	FREEZE--冻结库存
	ONWAY--在途库存
	DEFECT--残次品库存 */
	Type string `json:"type"`

	/* 货主ID */
	UserId int64 `json:"user_id"`
}

/* 物流宝商品 */
type WlbItem struct {
	/* 品牌ID */
	BrandId int64 `json:"brand_id"`

	/* 颜色 */
	Color string `json:"color"`

	/* 创建人 */
	Creator string `json:"creator"`

	/* 标记，用逗号隔开的字符串。
	BIT_HAS_AUTHORIZE 第1位 是否有授权规则;
	BATCH  第2位 是否有批次规则；
	SYNCHRONIZATION 第3位 是否有同步规则。 */
	Flag string `json:"flag"`

	/* 创建日期 */
	GmtCreate string `json:"gmt_create"`

	/* 修改日期 */
	GmtModified string `json:"gmt_modified"`

	/* 货类 */
	GoodsCat string `json:"goods_cat"`

	/* 高 */
	Height int64 `json:"height"`

	/* 商品id */
	Id int64 `json:"id"`

	/* 是否危险品 */
	IsDangerous bool `json:"is_dangerous"`

	/* 是否易碎 */
	IsFriable bool `json:"is_friable"`

	/* 是不是sku商品
	值为true或false */
	IsSku bool `json:"is_sku"`

	/* 商家编码 */
	ItemCode string `json:"item_code"`

	/* 最后修改人 */
	LastModifier string `json:"last_modifier"`

	/* mm */
	Length int64 `json:"length"`

	/* 商品的名称 */
	Name string `json:"name"`

	/* 包装材料 */
	PackageMaterial string `json:"package_material"`

	/* 父item的id，当item为物流宝子商品时，parent_id必填,否则不必填
	可通过父ID来得知商品的关系。 */
	ParentId int64 `json:"parent_id"`

	/* 价格 */
	Price int64 `json:"price"`

	/* 计价货类 */
	PricingCat string `json:"pricing_cat"`

	/* 属性key:value; key:value */
	Properties string `json:"properties"`

	/* 发布版本号，用来同步商 */
	PublishVersion int64 `json:"publish_version"`

	/* 商品备注 */
	Remark string `json:"remark"`

	/* 状态，item_status_valid -- 1 表示 有效 item_status_lock -- 2 表示锁住 */
	Status string `json:"status"`

	/* 前台商品名称 */
	Title string `json:"title"`

	/* 商品类型：
	NORMAL--普通类型;
	COMBINE--组合商品;
	DISTRIBUTION--分销商品;
	默认为NORMAL */
	Type string `json:"type"`

	/* 商品所有人淘宝ID */
	UserId int64 `json:"user_id"`

	/* 商品所有人淘宝nick */
	UserNick string `json:"user_nick"`

	/* 立方mm */
	Volume int64 `json:"volume"`

	/* 重量 */
	Weight int64 `json:"weight"`

	/* 宽 */
	Width int64 `json:"width"`
}

/* 活动数据结构 */
type Activity struct {
	/* 活动id */
	ActivityId int64 `json:"activity_id"`

	/* 领用优惠券的链接 */
	ActivityUrl string `json:"activity_url"`

	/* 已经领取的优惠券的数量 */
	AppliedCount int64 `json:"applied_count"`

	/* 活动对应的优惠券ID */
	CouponId int64 `json:"coupon_id"`

	/* self代表自己创建，other他人创建 */
	CreateUser string `json:"create_user"`

	/* 每个买家限领取优惠券的数量，1～5张 */
	PersonLimitCount int64 `json:"person_limit_count"`

	/* enabled代表有效，invalid代表失效。other代表空值 */
	Status string `json:"status"`

	/* 卖家设置优惠券领取的总领用量 */
	TotalCount int64 `json:"total_count"`
}

/* 数据推送任务 */
type DataPushTask struct {
	/* 任务所属的appkey */
	AppKey string `json:"app_key"`

	/* 任务编号 */
	Id int64 `json:"id"`

	/* 下次执行时间 */
	NextExecuteTime string `json:"next_execute_time"`

	/* 上一次同步到的时间点 */
	NowSyncTime string `json:"now_sync_time"`

	/* 任务的业务参数，用json格式表示 */
	Params string `json:"params"`

	/* 任务类型代码 */
	Type int64 `json:"type"`

	/* 任务所属的用户编号 */
	UserId int64 `json:"user_id"`

	/* 用户的nick */
	UserNick string `json:"user_nick"`

	/* 任务的版本号，在更新任务时，需要用到此数据 */
	Version int64 `json:"version"`
}

/* 优惠券数据结构 */
type Coupon struct {
	/* 订单满多少分才能用这个优惠券，501就是满501分能使用。注意：返回的是“分”，不是“元” */
	Condition int64 `json:"condition"`

	/* 优惠券ID */
	CouponId int64 `json:"coupon_id"`

	/* 优惠券创建时间 */
	CreatTime string `json:"creat_time"`

	/* 优惠券的创建渠道，自己创建/他人创建 */
	CreateChannel string `json:"create_channel"`

	/* 优惠券的面值，返回的是“分”，不是“元”，500代表500分相当于5元 */
	Denominations int64 `json:"denominations"`

	/* 优惠券的截止日期 */
	EndTime string `json:"end_time"`
}

/* 宝贝主图视频数据结构 */
type MpicVideo struct {
	/* 主图视频记录所关联的商品的数字id */
	NumIid int64 `json:"num_iid"`

	/* 主图视频的时长，单位：秒 */
	VideoDuaration int64 `json:"video_duaration"`

	/* 主图视频的在淘视频中的ID */
	VideoId int64 `json:"video_id"`

	/* 主图视频的缩略图URL */
	VideoPic string `json:"video_pic"`

	/* 主图视频的状态 */
	VideoStatus int64 `json:"video_status"`
}

/* 物流公司详细信息 */
type PartnerDetail struct {
	/* 物流公司支付宝账号 */
	AccountNo string `json:"account_no"`

	/* 物流公司代码 */
	CompanyCode string `json:"company_code"`

	/* 物流公司id */
	CompanyId int64 `json:"company_id"`

	/* 物流公司简称 */
	CompanyName string `json:"company_name"`

	/* 物流公司全名 */
	FullName string `json:"full_name"`

	/* 运单号验证正则表达式 */
	RegMailNo string `json:"reg_mail_no"`

	/* 旺旺id */
	WangwangId string `json:"wangwang_id"`
}

/* 折扣信息 */
type Discount struct {
	/* 创建时间 */
	Created string `json:"created"`

	/* 折扣详情 */
	Details []*DiscountDetail `json:"details>discount_detail"`

	/* 折扣ID */
	DiscountId int64 `json:"discount_id"`

	/* 修改时间 */
	Modified string `json:"modified"`

	/* 折扣名称 */
	Name string `json:"name"`
}

/* 折扣详情信息 */
type DiscountDetail struct {
	/* 创建时间 */
	Created string `json:"created"`

	/* 折扣详情ID */
	DetailId int64 `json:"detail_id"`

	/* 优惠方式:PERCENT（按折扣优惠）、PRICE（按减价优惠） */
	DiscountType string `json:"discount_type"`

	/* 优惠比率或者优惠价格 10%或10 */
	DiscountValue int64 `json:"discount_value"`

	/* 修改时间 */
	Modified string `json:"modified"`

	/* 会员等级的id或者分销商id */
	TargetId int64 `json:"target_id"`

	/* 等级名称或者分销商名称 */
	TargetName string `json:"target_name"`

	/* 折扣类型:GRADE（按会员等级优惠）、DISTRIBUTOR（按分销商优惠） */
	TargetType string `json:"target_type"`
}

/* 客服评价 */
type Evaluation struct {
	/* 客服评价内容 */
	EvaluationName string `json:"evaluation_name"`

	/* 评价数量 */
	EvaluationNum string `json:"evaluation_num"`
}

/* 子账号基本信息 */
type SubAccountInfo struct {
	/* 子账号是否参与分流 true:参与分流 false:未参与分流 */
	SubDispatchStatus bool `json:"sub_dispatch_status"`

	/* 子账号Id */
	SubId int64 `json:"sub_id"`

	/* 子账号用户名 */
	SubNick string `json:"sub_nick"`

	/* 子账号是否已欠费 true:已欠费 false:未欠费 */
	SubOwedStatus bool `json:"sub_owed_status"`

	/* 子账号当前状态：1正常，2卖家停用，3处罚冻结 */
	SubStatus int64 `json:"sub_status"`

	/* 主账号Id */
	UserId int64 `json:"user_id"`

	/* 主账号用户名 */
	UserNick string `json:"user_nick"`
}

/* 在线时长 */
type OnlineTimeById struct {
	/* 客服在线时间长度（秒） */
	OnlineTimes int64 `json:"online_times"`

	/* 客服人员ID */
	ServiceStaffId string `json:"service_staff_id"`
}

/* 物流服务标签 */
type LogisticServiceTag struct {
	/* 物流服务下的标签属性,多个标签之间有";"分隔 */
	ServiceTag string `json:"service_tag"`

	/* 服务类型=编码
	平邮=POST
	快递=FAST
	EMS=EMS
	消费者选快递时为FAST */
	ServiceType string `json:"service_type"`
}

/* 平均等待时长 */
type WaitingTimeById struct {
	/* 平均等待时间长度（秒） */
	AvgWaitingTimes int64 `json:"avg_waiting_times"`

	/* 客服人员ID */
	ServiceStaffId string `json:"service_staff_id"`
}

/* 客服回复统计 */
type ReplyStatById struct {
	/* 客服回复数 */
	ReplyNum int64 `json:"reply_num"`

	/* 客服ID */
	UserId string `json:"user_id"`
}

/* 部门信息 */
type Department struct {
	/* 部门ID */
	DepartmentId int64 `json:"department_id"`

	/* 部门名称 */
	DepartmentName string `json:"department_name"`

	/* 当前部门的父部门ID */
	ParentId int64 `json:"parent_id"`

	/* 部门下关联的子账号id列表 */
	SubUserIds []int64 `json:"sub_user_ids"`
}

/* 子账号职务信息 */
type Duty struct {
	/* 职务ID */
	DutyId int64 `json:"duty_id"`

	/* 职务级别 */
	DutyLevel int64 `json:"duty_level"`

	/* 职务名称 */
	DutyName string `json:"duty_name"`
}

/* 子账号详细信息，其中包括账号基本信息、员工信息和部门职务信息 */
type SubUserFullInfo struct {
	/* 部门Id */
	DepartmentId int64 `json:"department_id"`

	/* 部门名称 */
	DepartmentName string `json:"department_name"`

	/* 职务Id */
	DutyId int64 `json:"duty_id"`

	/* 职务等级 */
	DutyLevel int64 `json:"duty_level"`

	/* 职务名称 */
	DutyName string `json:"duty_name"`

	/* 员工ID */
	EmployeeId int64 `json:"employee_id"`

	/* 员工姓名 */
	EmployeeName string `json:"employee_name"`

	/* 员工花名 */
	EmployeeNickname string `json:"employee_nickname"`

	/* 入职员工工号 */
	EmployeeNum string `json:"employee_num"`

	/* 员工入职时间 */
	EntryDate string `json:"entry_date"`

	/* 直接上级的员工ID */
	LeaderId int64 `json:"leader_id"`

	/* 办公电话 */
	OfficePhone string `json:"office_phone"`

	/* 父部门Id */
	ParentDepartment int64 `json:"parent_department"`

	/* 员工性别  1:男;  2:女 */
	Sex int64 `json:"sex"`

	/* 子账号是否参与分流 true:参与分流 false:未参与分流 */
	SubDispatchStatus bool `json:"sub_dispatch_status"`

	/* 子账号Id */
	SubId int64 `json:"sub_id"`

	/* 子账号用户名 */
	SubNick string `json:"sub_nick"`

	/* 子账号是否已欠费 true:已欠费 false:未欠费 */
	SubOwedStatus bool `json:"sub_owed_status"`

	/* 子账号当前状态：1正常，2卖家停用，3处罚冻结 */
	SubStatus int64 `json:"sub_status"`

	/* 子账号企业邮箱 */
	SubuserEmail string `json:"subuser_email"`

	/* 主账号企业邮箱 */
	UserEmail string `json:"user_email"`

	/* 主账号Id */
	UserId int64 `json:"user_id"`

	/* 主账号用户名 */
	UserNick string `json:"user_nick"`

	/* 工作地点 */
	WorkLocation string `json:"work_location"`
}

/* 物流的服务标签 */
type LogisticsTag struct {
	/* 服务标签 */
	LogisticServiceTagList []*LogisticServiceTag `json:"logistic_service_tag_list>logistic_service_tag"`

	/* 主订单或子订单的订单号 */
	OrderId string `json:"order_id"`
}

/* 物流服务类型 */
type WaybillServiceType struct {
	/* 物流服务能力编码 */
	Code string `json:"code"`

	/* 物流服务能力名称 */
	Name string `json:"name"`
}

/* 物流产品类型 */
type WaybillProductType struct {
	/* 物流产品类型编码 */
	Code string `json:"code"`

	/* 物流产品类型名称 */
	Name string `json:"name"`

	/* 物流服务类型列表 */
	ServiceTypes []*WaybillServiceType `json:"service_types>waybill_service_type"`
}

/* 判定服务是否可达的返回结果 */
type AddressReachableResult struct {
	/* 区域编码 */
	DivisionId int64 `json:"division_id"`

	/* 错误码 */
	ErrorCode string `json:"error_code"`

	/* 错误信息 */
	ErrorMsg string `json:"error_msg"`

	/* 物流公司代号 */
	PartnerCode string `json:"partner_code"`

	/* 物流公司编码ID */
	PartnerId int64 `json:"partner_id"`

	/* 物流公司名称 */
	PartnerName string `json:"partner_name"`

	/* 服务是否可达，    0 - 不可达    1 - 可达    2 - 不确定    3 - 未配置 */
	Reachable int64 `json:"reachable"`

	/* 服务对应的数字编码，如揽派范围对应88 */
	ServiceType int64 `json:"service_type"`

	/* 是否成功 */
	Success bool `json:"success"`
}

/* 产品规格审核信息 */
type Ticket struct {
	/* 如果产品规格，需要商家审核，为商家审核用户ID */
	AuditSellerId int64 `json:"audit_seller_id"`

	/* 创建人ID */
	CreateUserId int64 `json:"create_user_id"`

	/* 产品规格申请时间 */
	GmtCreate string `json:"gmt_create"`

	/* 产品规格审核单最后修改时间 */
	GmtModified string `json:"gmt_modified"`

	/* 关于审核原因，更详细的说明 */
	Memo string `json:"memo"`

	/* 审核原因 */
	Reason string `json:"reason"`

	/* 产品规格ID */
	SpecId int64 `json:"spec_id"`

	/* 1, "商家确认"
	2, "商家拒绝"
	3, "小二确认"
	4, "小二拒绝"
	5, "待商家处理"
	6, "商家审核超时"
	7, "待小二审核"
	9, "品牌商确认"
	10, "免审通过"
	14, "免审拒绝" */
	Status int64 `json:"status"`
}

/* 包裹里面的商品类型 */
type PackageItem struct {
	/* 商品数量 */
	Count int64 `json:"count"`

	/* 商品类型 */
	ItemName string `json:"item_name"`
}

/* 物流服务能力集合 */
type LogisticsService struct {
	/* 服务编码 */
	ServiceCode string `json:"service_code"`

	/* 服务名称 */
	ServiceName string `json:"service_name"`

	/* 服务类型值，json格式表示 */
	ServiceValue4Json string `json:"service_value4_json"`
}

/* 打印返回信息 */
type WaybillApplyPrintCheckInfo struct {
	/* 打印提示信息编码 */
	NoticeCode string `json:"notice_code"`

	/* 打印提示信息 */
	NoticeMessage string `json:"notice_message"`

	/* 打印次数 */
	PrintQuantity int64 `json:"print_quantity"`

	/* 电子面单号 */
	WaybillCode string `json:"waybill_code"`
}

/* 面单详情 */
type WaybillDetailQueryInfo struct {
	/* 收货人地址 */
	ConsigneeAddress *WaybillAddress `json:"consignee_address"`

	/* 包裹对应的派件（收件）物流服务商网点（分支机构）代码 */
	ConsigneeBranchCode string `json:"consignee_branch_code"`

	/* 包裹对应的派件（收件）物流服务商网点（分支机构）名称 */
	ConsigneeBranchName string `json:"consignee_branch_name"`

	/* 收件人姓名 */
	ConsigneeName string `json:"consignee_name"`

	/* 收件人联系方式 */
	ConsigneePhone string `json:"consignee_phone"`

	/* 物流商编码CODE */
	CpCode string `json:"cp_code"`

	/* 创建时间 */
	CreateTime string `json:"create_time"`

	/* 最后一次打印时间 */
	LastPrintTime string `json:"last_print_time"`

	/* 物流服务能力集合 */
	LogisticsServiceList []*LogisticsService `json:"logistics_service_list>logistics_service"`

	/* 包裹里面的商品类型 */
	PackageItems []*PackageItem `json:"package_items>package_item"`

	/* 揽收时间 */
	PickupTime string `json:"pickup_time"`

	/* 打印次数 */
	PrintCount int64 `json:"print_count"`

	/* 快递服务产品类型编码 */
	ProductType string `json:"product_type"`

	/* 使用者ID */
	RealUserId int64 `json:"real_user_id"`

	/* 发件人姓名 */
	SendName string `json:"send_name"`

	/* 发件人联系方式 */
	SendPhone string `json:"send_phone"`

	/* 发货地址 */
	ShippingAddress *WaybillAddress `json:"shipping_address"`

	/* 发货网点编码 */
	ShippingBranchCode string `json:"shipping_branch_code"`

	/* 发货网点信息 */
	ShippingBranchName string `json:"shipping_branch_name"`

	/* 大头笔信息 */
	ShortAddress string `json:"short_address"`

	/* 签收时间 */
	SignTime string `json:"sign_time"`

	/* 面单状态 */
	Status int64 `json:"status"`

	/* 交易订单列表 */
	TradeOrderList []string `json:"trade_order_list"`

	/* 包裹重量 单位为G(克) */
	Volume int64 `json:"volume"`

	/* 电子面单信息 */
	WaybillCode string `json:"waybill_code"`

	/* 包裹体积 单位为ML(毫升)或立方厘米 */
	Weight int64 `json:"weight"`
}

/* 本地生活垂直市场数据结构，修改宝贝时在参数empty_fields里设置locality_life可删除所有电子凭证信息 */
type LocalityLife struct {
	/* 表示是否使用邮寄 0: 代表不使用邮寄； 1：代表使用邮寄；如果不设置这个值，代表不使用邮寄 */
	ChooseLogis string `json:"choose_logis"`

	/* 电子凭证业务属性 */
	Eticket string `json:"eticket"`

	/* 电子交易凭证有效期，有三种格式：
	如果有效期为起止日期类型，此值为2012-08-06,2012-08-16
	如果有效期为【购买成功日 至】类型则格式为2012-08-16
	如果有效期为天数类型则格式为15 */
	Expirydate string `json:"expirydate"`

	/* 格式为
	码商id:nick */
	Merchant string `json:"merchant"`

	/* 网点ID,在参数empty_fields里设置locality_life.network_id可删除网点ID */
	NetworkId string `json:"network_id"`

	/* 电子凭证售中自动退款比例 */
	OnsaleAutoRefundRatio int64 `json:"onsale_auto_refund_ratio"`

	/* 退款比例，百分比%前的数字，1-100的正整数值；在参数empty_fields里设置locality_life.refund_ratio可删除退款比例 */
	RefundRatio int64 `json:"refund_ratio"`

	/* 退款码费承担方。发布电子凭证宝贝的时候会增加“退款码费承担方”配置项，可选填：(1)s（卖家承担） (2)b(买家承担) */
	Refundmafee string `json:"refundmafee"`

	/* 核销打款:1代表核销打款,0代表非核销打款;
	在参数empty_fields里设置locality_life.verification可删除核销打款 */
	Verification string `json:"verification"`
}

/* 该数据结构保存宝贝描述对应的规范化信息 */
type DescModuleInfo struct {
	/* 代表宝贝描述中规范化打标使用到的模块id列表，以逗号分隔。 */
	AnchorModuleIds string `json:"anchor_module_ids"`

	/* 类型代表规范化打标的类型：人工或自动化 */
	Type int64 `json:"type"`
}

/* ProductSpec(产品规格)结构。 */
type ProductSpec struct {
	/* 产品规格条形码，支持EAN-13格式。 */
	Barcode string `json:"barcode"`

	/* 产品品牌id */
	BrandId int64 `json:"brand_id"`

	/* 认证图片列表 */
	CertifiedPics []*CertPicInfo `json:"certified_pics>cert_pic_info"`

	/* 认证文本列表 */
	CertifiedTxts []*CertTxtInfo `json:"certified_txts>cert_txt_info"`

	/* 基础色数据 */
	ChangeProp string `json:"change_prop"`

	/* 用户输入的属性值存放位置，例如可输入的销售属性，当用户获取pid vid后，应该先从spec_props_alias中获取，然后通过类目属性获取，获取不到，可以通过这个字段获取。 */
	CustomePropsName string `json:"custome_props_name"`

	/* 产品规格吊牌价，以分为单位的整数，非负,无默认值，上限999999999 */
	LabelPrice int64 `json:"label_price"`

	/* 上市时间 */
	MarketTime string `json:"market_time"`

	/* 规格的数量，仅当通过组合产品获取时有值 */
	Number int64 `json:"number"`

	/* 产品的主图片地址。绝对地址，格式：http://host/image_path。 */
	PicUrl string `json:"pic_url"`

	/* 产品货号 */
	ProductCode string `json:"product_code"`

	/* 产品ID。 */
	ProductId int64 `json:"product_id"`

	/* 产品规格ID。 */
	SpecId int64 `json:"spec_id"`

	/* 产品规格的销售属性组合。格式为：pid:vid;pid:vid。其中：pid是销售属性id，vid是销售属性值id。如果该类目品牌下面没有销售属性，可以不用填写。销售属性通过tmall.brandcat.salespro.get接口获取。 */
	SpecProps string `json:"spec_props"`

	/* 销售属性值别名。格式为：pid1:vid1:别名1;pid2:vid2:别名2。其中：pid是销售属性id，vid是销售属性值id。别名长度不可以超过30个字符。目前只有颜色销售属性支持别名。 */
	SpecPropsAlias string `json:"spec_props_alias"`

	/* 1:表示可以使用的数据，也就是审核通过的。
	3：表示等待小二审核的产品规格，这个数据暂时还不能使用，要等待审核通过后，才能使用。 */
	Status int64 `json:"status"`
}

/* 产品资质认证图片信息，包括认证类型以及图片url */
type CertPicInfo struct {
	/* 认证类型的数值id */
	CertType int64 `json:"cert_type"`

	/* 认证图片的url地址 */
	PicUrl string `json:"pic_url"`
}

/* 用于保存宝贝描述规范化模块信息 */
type IdsModule struct {
	/* 宝贝描述规范化模块id */
	Id int64 `json:"id"`

	/* 宝贝描述规范化模块名 */
	Name string `json:"name"`

	/* 0为自动打标；
	1为人工打标； */
	Type int64 `json:"type"`
}

/* 被管控的品牌和类目的所对应的销售属性 */
type CatBrandSaleProp struct {
	/* 被管控的品牌的Id */
	BrandId int64 `json:"brand_id"`

	/* 被管控的类目ID */
	CatId int64 `json:"cat_id"`

	/* 如果该属性为营销属性，则获取默认值 */
	DefMarketPropValue int64 `json:"def_market_prop_value"`

	/* true表示：不是产品的规格属性
	false表示：是产品的规格属性。 */
	IsNotSpec bool `json:"is_not_spec"`

	/* 被管控的销售属性ID */
	PropertyId int64 `json:"property_id"`
}

/* 管控的类目以及品牌信息 */
type BrandCatControlInfo struct {
	/* 管控的品牌类目信息，一组列表 */
	BrandCatControls []*BrandCatControl `json:"brand_cat_controls>brand_cat_control"`
}

/* 管控的品牌类目信息 */
type BrandCatControl struct {
	/* 被管控的品牌的ID号码 */
	BrandId int64 `json:"brand_id"`

	/* 被管控的品牌名称 */
	BrandName string `json:"brand_name"`

	/* 被管控的类目的ID号 */
	CatId int64 `json:"cat_id"`

	/* 被管控的类目的ID号 */
	CatName string `json:"cat_name"`

	/* 以;隔开多个认证资料。:隔开资料ID与内容。如？1:产品包装图片;2:完整产品资质 */
	CertifiedData string `json:"certified_data"`
}

/* 交易回流信息 */
type TradeTrace struct {
	/* 动作发生的时间 */
	ActionTime string `json:"action_time"`

	/* 应用标题 */
	AppTitle string `json:"app_title"`

	/* 子订单的id列表,以逗号分割 */
	OrderIds string `json:"order_ids"`

	/* 备注字段 */
	Remark string `json:"remark"`

	/* 卖家的淘宝nick */
	SellerNick string `json:"seller_nick"`

	/* 回流的订单状态 */
	Status string `json:"status"`

	/* 交易tid */
	Tid int64 `json:"tid"`
}

/* 订单全链路中的订单状态统计 */
type TradeStat struct {
	/* 数量 */
	Count int64 `json:"count"`

	/* 状态名称 */
	Status string `json:"status"`
}

/* 订单全链路用户信息 */
type HlUserDO struct {
	/* 回流信息是否开通买家端展示 */
	OpenForBuyer string `json:"open_for_buyer"`
}

/* 订单全链路退款追踪详情 */
type RefundTrace struct {
	/* 动作发生的时间 */
	ActionTime string `json:"action_time"`

	/* 应用标题 */
	AppTitle string `json:"app_title"`

	/* 退款编号 */
	RefundId int64 `json:"refund_id"`

	/* 备注字段 */
	Remark string `json:"remark"`

	/* 卖家的淘宝nick */
	SellerNick string `json:"seller_nick"`

	/* 回流的订单状态 */
	Status string `json:"status"`

	/* 交易tid */
	Tid int64 `json:"tid"`
}

/* 经销订单监控记录信息 */
type TradeMonitor struct {
	/* 地区 */
	Area string `json:"area"`

	/* 交易订单的商品购买数量 */
	BuyAmount int64 `json:"buy_amount"`

	/* 收货人姓名 */
	BuyerFullName string `json:"buyer_full_name"`

	/* 买家的淘宝账号昵称 */
	BuyerNick string `json:"buyer_nick"`

	/* 城市 */
	City string `json:"city"`

	/* 经销商的淘宝账号昵称 */
	DistributorNick string `json:"distributor_nick"`

	/* 交易订单的商品id */
	ItemId int64 `json:"item_id"`

	/* 交易订单的商品的商家编码 */
	ItemNumber string `json:"item_number"`

	/* 交易订单的商品价格 */
	ItemPrice int64 `json:"item_price"`

	/* 交易订单的商品的sku名称 */
	ItemSkuName string `json:"item_sku_name"`

	/* 交易订单的商品的sku商家编码 */
	ItemSkuNumber string `json:"item_sku_number"`

	/* 交易订单的商品标题 */
	ItemTitle string `json:"item_title"`

	/* 交易订单的商品总价格（单价×数量+改价+优惠） */
	ItemTotalPrice int64 `json:"item_total_price"`

	/* 交易订单的付款时间 */
	PayTime string `json:"pay_time"`

	/* 供应商的产品id */
	ProductId int64 `json:"product_id"`

	/* 供应商的产品的商家编码 */
	ProductNumber string `json:"product_number"`

	/* 供应商的产品的sku商家编码 */
	ProductSkuNumber string `json:"product_sku_number"`

	/* 供应商的产品标题 */
	ProductTitle string `json:"product_title"`

	/* 省份 */
	Provence string `json:"provence"`

	/* 交易订单的商品最高零售价 */
	RetailPriceHigh int64 `json:"retail_price_high"`

	/* 交易订单的商品最低零售价 */
	RetailPriceLow int64 `json:"retail_price_low"`

	/* 收货人地址 */
	ShippingAddress string `json:"shipping_address"`

	/* 交易订单的状态：
	WAIT_SELLER_SEND_GOODS(已付款，待发货）<br>WAIT_BUYER_CONFIRM_GOODS(已付款，已发货)<br>TRADE_FINISHED(交易成功)
	TRADE_CLOSED(交易关闭)<br>TRADE_REFUNDING（退款中） */
	Status string `json:"status"`

	/* 交易订单的子订单号 */
	SubTcOrderId int64 `json:"sub_tc_order_id"`

	/* 供应商的淘宝账号昵称 */
	SupplierNick string `json:"supplier_nick"`

	/* 商品的卖出金额调整，金额增加时为正数，金额减少时为负数，单位是分，不带小数 */
	TcAdjustFee int64 `json:"tc_adjust_fee"`

	/* 优惠金额，始终为正数，单位是分，不带小数 */
	TcDiscountFee int64 `json:"tc_discount_fee"`

	/* 交易订单号 */
	TcOrderId int64 `json:"tc_order_id"`

	/* 商品优惠类型：聚划算、秒杀或其他 */
	TcPreferentialType string `json:"tc_preferential_type"`

	/* 主键id */
	TradeMonitorId int64 `json:"trade_monitor_id"`
}

/* 客户等待（客服）平均时长列表 */
type WaitingTimesOnDay struct {
	/* 等待时长（统计）日期 */
	WaitingDate string `json:"waiting_date"`

	/* 等待时长列表 */
	WaitingTimeByIds []*WaitingTimeById `json:"waiting_time_by_ids>waiting_time_by_id"`
}

/* 单条交易子订单语义标签对象 */
type ImprFeedIdDO struct {
	/* 存放主评和追评的评价标签信息
	biz_type(1:主评 2:追评) */
	FeedInfoList []*ImprFeedInfoDO `json:"feed_info_list>impr_feed_info_d_o"`

	/* 被评价者昵称 */
	Nick string `json:"nick"`

	/* 评价结果（1好评 0中评 -1差评） */
	Rate int64 `json:"rate"`
}

/* 未回复统计列表(按天) */
type NonReplyStatOnDay struct {
	/* （某天的）未回复统计 */
	NonreplyDate string `json:"nonreply_date"`

	/* 未回复统计列表 */
	NonreplyStatByIds []*NonreplyStatById `json:"nonreply_stat_by_ids>nonreply_stat_by_id"`
}

/* 客服评价统计列表(按天) */
type StaffEvalStatOnDay struct {
	/* 评价产生的日期 */
	EvalDate string `json:"eval_date"`

	/* 客服评价统计列表 */
	StaffEvalStatByIds []*StaffEvalStatById `json:"staff_eval_stat_by_ids>staff_eval_stat_by_id"`
}

/* 某天的客服在线时长列表 */
type OnlineTimesOnDay struct {
	/* 在线日期 */
	OnlineDate string `json:"online_date"`

	/* 在线时长列表 */
	OnlineTimeByIds []*OnlineTimeById `json:"online_time_by_ids>online_time_by_id"`
}

/* 批量异步任务结果 */
type Task struct {
	/* 下载文件的MD5校验码，通过此校验码可以检查下载的文件是否是完整的。 */
	CheckCode string `json:"check_code"`

	/* 任务创建时间 */
	Created string `json:"created"`

	/* 大任务结果下载地址。当创建的认任务是大数据量的任务时，获取结果会返回此字段，同时subtasks列表会为空。
	通过这个地址可以下载到结果的结构体，格式同普通任务下载的一样。
	每次获取到的地址只能下载一次。下载的文件加上后缀名.zip打开。 */
	DownloadUrl string `json:"download_url"`

	/* 此任务是由哪个api产生的 */
	Method string `json:"method"`

	/* 定时类型任务的执行时间点 */
	Schedule string `json:"schedule"`

	/* 异步任务处理状态。new（还未开始处理），doing（处理中），done（处理结束）。 */
	Status string `json:"status"`

	/* 子任务处理结果，如果任务还没有处理完，返回的结果列表为空。如果任务处理完毕，返回子任务结果列表 */
	Subtasks []*Subtask `json:"subtasks>subtask"`

	/* 异步任务id。创建异步任务时返回的任务id号 */
	TaskId int64 `json:"task_id"`
}

/* 图片链接 */
type PicUrl struct {
	/* 图片链接地址 */
	Url string `json:"url"`
}

/* 批量异步任务的子任务结果 */
type Subtask struct {
	/* 标记子任务是否成功。为true表示子任务成功，用户可以按照正确执行的结果格式解析sub_task_result。为false表示子任务失败了，用户需要按照失败的结果格式解析sub_task_result（里面只有sub_code和sub_msg） */
	IsSuccess bool `json:"is_success"`

	/* 子任务的有效请求参数，以json格式进行key:value的组合 */
	SubTaskRequest string `json:"sub_task_request"`

	/* 子任务返回的结果，以json格式进行key:value组合，可以和单个api请求结果解析复用。以获取交易订单详情为例：子任务执行成功返回的结果格式为：{“trade”:{"tid":123456,"seller_nick":"淘宝卖家"}}；子任务执行失败结果格式为{"sub_code":"isv.trade-not-exist","sub_msg":"交易订单不存在"} */
	SubTaskResult string `json:"sub_task_result"`
}

/* (某天)回复统计列表 */
type ReplyStatOnDay struct {
	/* 某天（的回复统计） */
	ReplyDate string `json:"reply_date"`

	/* 客服回复统计 */
	ReplyStatByIds []*ReplyStatById `json:"reply_stat_by_ids>reply_stat_by_id"`
}

/* 交易的优惠信息详情 */
type PromotionDetail struct {
	/* 优惠金额（免运费、限时打折时为空）,单位：元 */
	DiscountFee float64 `json:"discount_fee,string"`

	/* 赠品的宝贝id */
	GiftItemId string `json:"gift_item_id"`

	/* 满就送商品时，所送商品的名称 */
	GiftItemName string `json:"gift_item_name"`

	/* 满就送礼物的礼物数量 */
	GiftItemNum string `json:"gift_item_num"`

	/* 交易的主订单或子订单号 */
	Id int64 `json:"id"`

	/* 优惠活动的描述 */
	PromotionDesc string `json:"promotion_desc"`

	/* 优惠id，(由营销工具id、优惠活动id和优惠详情id组成，结构为：营销工具id-优惠活动id_优惠详情id，如mjs-123024_211143） */
	PromotionId string `json:"promotion_id"`

	/* 优惠信息的名称 */
	PromotionName string `json:"promotion_name"`
}

/* KFC 关键词过滤匹配结果 */
type KfcSearchResult struct {
	/* 过滤后的文本：
	当匹配到B等级的词时，文本中的关键词被替换为*号，content即为关键词替换后的文本；
	其他情况，content始终为null */
	Content string `json:"content"`

	/* 匹配到的关键词的等级，值为null，或为A、B、C、D。
	当匹配不到关键词时，值为null，否则值为A、B、C、D中的一个。
	A、B、C、D等级按严重程度从高至低排列。 */
	Level string `json:"level"`

	/* 是否匹配到关键词,匹配到则为true. */
	Matched bool `json:"matched"`
}

/* 组及其成员列表 */
type GroupMember struct {
	/* 组编号 */
	GroupId int64 `json:"group_id"`

	/* 组名称 */
	GroupName string `json:"group_name"`

	/* 组成员列表，逗号分隔 */
	MemberList string `json:"member_list"`
}

/* 属性输入特征DO */
type PropertyInputDO struct {
	/* property_id对应的属性是否可输入 */
	IsAllowInput bool `json:"is_allow_input"`

	/* 如果property_id对应的属性是子属性，is_root_allow_input标识该子属性的顶级父属性是否可输入。否则is_root_allow_input和is_allow_input
	值是一样的。目前只有品牌会出现父属性不可输入，子属性可输入的情况 */
	IsRootAllowInput bool `json:"is_root_allow_input"`

	/* property_id对应的属性是不是子属性 */
	IsSubProperty bool `json:"is_sub_property"`

	/* 属性ID */
	PropertyId int64 `json:"property_id"`
}

/* 物流宝订单，并且包含订单详情 */
type WlbOrderDetail struct {
	/* 如果是交易单，则显示交易中买家昵称 */
	BuyerNick string `json:"buyer_nick"`

	/* 订单创建时间 */
	CreateTime string `json:"create_time"`

	/* 是否已全部完成 */
	IsCompleted bool `json:"is_completed"`

	/* 订单最后一次修改时间 */
	ModifyTime string `json:"modify_time"`

	/* 出库或者入库，IN表示入库，OUT表示出库 */
	OperateType string `json:"operate_type"`

	/* 订单编码 */
	OrderCode string `json:"order_code"`

	/* 物流宝订单对应的商品详情 */
	OrderItemList []*WlbOrderItem `json:"order_item_list>wlb_order_item"`

	/* 订单来源:
	产生物流订单的原因，比如:

	订单来源:1:TAOBAO;2:EXT;3:ERP;4:WMS */
	OrderSource string `json:"order_source"`

	/* 对应创建物流宝订单top接口中的的out_biz_code字段，主要是用来去重用 */
	OrderSourceCode string `json:"order_source_code"`

	/* 物流状态，
	订单已创建：0
	订单已取消: -1
	订单关闭:-3
	下发中: 10
	已下发: 11
	接收方拒签 :-100
	已发货:100
	签收成功:200 */
	OrderStatus string `json:"order_status"`

	/* (1)其它:    OTHER
	(2)淘宝交易: TAOBAO
	(3)301:调拨: ALLOCATION
	(4)401:盘点:CHECK
	(5)501:销售采购:PRUCHASE */
	OrderSubType string `json:"order_sub_type"`

	/* 1:正常订单: NARMAL
	2:退货订单: RETURN
	3:换货订单: CHANGE */
	OrderType string `json:"order_type"`

	/* 订单备注 */
	Remark string `json:"remark"`

	/* 仓库编码 */
	StoreCode string `json:"store_code"`

	/* 卖家ID */
	UserId int64 `json:"user_id"`

	/* 卖家NICK */
	UserNick string `json:"user_nick"`
}

/* 同意退款返回值中一部分，退款编号与结果信息映射类 */
type RefundMappingResult struct {
	/* 结果信息 */
	Message string `json:"message"`

	/* 退款ID */
	RefundId int64 `json:"refund_id"`

	/* 是否成功 */
	Succ bool `json:"succ"`
}

/* 服务平台评价流水对象 */
type ScoreResult struct {
	/* 服务态度评分 */
	AttitudeScore string `json:"attitude_score"`

	/* 平均分 */
	AvgScore string `json:"avg_score"`

	/* 易用性评分 */
	EasyuseScore string `json:"easyuse_score"`

	/* 评价时间 */
	GmtCreate string `json:"gmt_create"`

	/* 评价id */
	Id int64 `json:"id"`

	/* 是否实际付费 1-实际付费 2-实际未付费 */
	IsPay int64 `json:"is_pay"`

	/* 是否为有效评分 1-有效评分 2-无效评分 */
	IsValid int64 `json:"is_valid"`

	/* 服务规格code */
	ItemCode string `json:"item_code"`

	/* 服务规格名称 */
	ItemName string `json:"item_name"`

	/* 描述相符 */
	MatchedScore string `json:"matched_score"`

	/* 专业性评分 */
	ProfScore string `json:"prof_score"`

	/* 交片速度 */
	RapidScore string `json:"rapid_score"`

	/* 服务code */
	ServiceCode string `json:"service_code"`

	/* 稳定性评分 */
	StabilityScore string `json:"stability_score"`

	/* 评论内容 */
	Suggestion string `json:"suggestion"`

	/* 评价人用户昵称 */
	UserNick string `json:"user_nick"`
}

/* 筛单结果集 */
type AddressReachableTopResult struct {
	/* 筛单结果l列表 */
	ReachableResultList []*AddressReachableResult `json:"reachable_result_list>address_reachable_result"`
}

/* 物流跟踪信息的一条 */
type TransitStepInfo struct {
	/* 节点说明 ，指明当前节点揽收、派送，签收。 */
	Action string `json:"action"`

	/* 描述 */
	Desc string `json:"desc"`

	/* 地址地一 */
	NodeDescription string `json:"node_description"`

	/* 状态描述 */
	StatusDesc string `json:"status_desc"`

	/* 状态发生的时间 */
	StatusTime string `json:"status_time"`

	/* 时间。。 */
	Time string `json:"time"`
}

/* 子账号角色 */
type Role struct {
	/* 创建时间 */
	CreateTime string `json:"create_time"`

	/* 角色描述 */
	Description string `json:"description"`

	/* 修改时间 */
	ModifiedTime string `json:"modified_time"`

	/* 所拥有权限 */
	Permissions []*Permission `json:"permissions>permission"`

	/* 角色id */
	RoleId int64 `json:"role_id"`

	/* 角色名 */
	RoleName string `json:"role_name"`

	/* 卖家Id */
	SellerId int64 `json:"seller_id"`
}

/* 子账号所拥有的权限对象(直接赋予的权限和通过角色赋予的权限的总和对象) */
type SubUserPermission struct {
	/* 子账号被直接赋予的权限点列表 */
	Permissions []*Permission `json:"permissions>permission"`

	/* 子账号被赋予的角色信息(Role)列表。列表中的角色对象只有role_id，role_name，permissions信息 */
	Roles []*Role `json:"roles>role"`
}

/* 店铺级优惠信息 */
type PromotionInShop struct {
	/* 优惠活动名称 */
	Name string `json:"name"`

	/* 优惠详情描述。 */
	PromotionDetailDesc string `json:"promotion_detail_desc"`

	/* idValue值 */
	PromotionId string `json:"promotion_id"`
}

/* 子订单的帐务数据结构 */
type OrderAmount struct {
	/* 卖家手工调整的子订单的优惠金额.格式为:1.01;单位:元;精确到小数点后两位. */
	AdjustFee string `json:"adjust_fee"`

	/* 子订单的系统优惠金额。精确到2位小数;单位:元。如:200.07，表示:200元7分 */
	DiscountFee string `json:"discount_fee"`

	/* 分摊之后的实付金额 */
	DivideOrderFee string `json:"divide_order_fee"`

	/* 子交易订单中购买商品的数量 */
	Num int64 `json:"num"`

	/* 子订单对应的商品数字id */
	NumIid int64 `json:"num_iid"`

	/* 子交易订单编号 */
	Oid int64 `json:"oid"`

	/* 优惠分摊 */
	PartMjzDiscount string `json:"part_mjz_discount"`

	/* 子订单实付金额。精确到2位小数，单位:元。如:200.07，表示:200元7分。计算公式如下：payment = price * num + adjust_fee - discount_fee + post_fee(邮费，单笔子订单时子订单实付金额包含邮费，多笔子订单时不包含邮费)；对于退款成功的子订单，由于主订单的优惠分摊金额，会造成该字段可能不为0.00元。建议使用退款前的实付金额减去退款单中的实际退款金额计算。 */
	Payment string `json:"payment"`

	/* 商品价格。精确到2位小数;单位:元。如:200.07，表示:200元7分 */
	Price string `json:"price"`

	/* 子订单的系统优惠的名称，对应于discount_fee的名称 */
	PromotionName string `json:"promotion_name"`

	/* 子订单对应的商品的sku_id */
	SkuId int64 `json:"sku_id"`

	/* SKU的值。如：机身颜色:黑色;手机套餐:官方标配 */
	SkuPropertiesName string `json:"sku_properties_name"`

	/* 商品标题 */
	Title string `json:"title"`
}

/* 优惠信息对象 */
type PromotionDisplayTop struct {
	/* 单品级优惠信息 */
	PromotionInItem []*PromotionInItem `json:"promotion_in_item>promotion_in_item"`

	/* 店铺级优惠信息 */
	PromotionInShop []*PromotionInShop `json:"promotion_in_shop>promotion_in_shop"`
}

/* 单品级优惠信息 */
type PromotionInItem struct {
	/* 优惠描述 */
	Desc string `json:"desc"`

	/* 优惠结束时间 */
	EndTime string `json:"end_time"`

	/* 优惠折后价格 */
	ItemPromoPrice float64 `json:"item_promo_price,string"`

	/* 优惠展示名称 */
	Name string `json:"name"`

	/* 需要支付附加物，显示为+xxx。如：+20淘金币 */
	OtherNeed string `json:"other_need"`

	/* 赠送东西。如：送10商城积分 */
	OtherSend string `json:"other_send"`

	/* idValue的值 */
	PromotionId string `json:"promotion_id"`

	/* sku价格对应的id（保证二者顺序相同） */
	SkuIdList []string `json:"sku_id_list"`

	/* sku价格列表 */
	SkuPriceList []float64 `json:"sku_price_list,string"`

	/* 优惠开始时间 */
	StartTime string `json:"start_time"`
}

/* 推荐的关联商品 */
type FavoriteItem struct {
	/* 商品名称 */
	ItemName string `json:"item_name"`

	/* 商品图片地址 */
	ItemPictrue string `json:"item_pictrue"`

	/* 商品价格 */
	ItemPrice float64 `json:"item_price,string"`

	/* 商品的详情页面地址 */
	ItemUrl string `json:"item_url"`

	/* 促销价格 */
	PromotionPrice float64 `json:"promotion_price,string"`

	/* 商品销售次数 */
	SellCount int64 `json:"sell_count"`

	/* 商品id（具有跟踪效果）代替原来的item_id  <a href="http://dev.open.taobao.com/bbs/read.php?tid=24323">详细说明</a> */
	TrackIid string `json:"track_iid"`
}

/* 推荐关联店铺信息 */
type FavoriteShop struct {
	/* 店铺卖家信用 */
	Rate int64 `json:"rate"`

	/* 卖家ID */
	SellerId int64 `json:"seller_id"`

	/* 卖家昵称 */
	SellerNick string `json:"seller_nick"`

	/* 店铺ID */
	ShopId int64 `json:"shop_id"`

	/* 店铺ID */
	ShopName string `json:"shop_name"`

	/* 店铺LOGO图片 */
	ShopPic string `json:"shop_pic"`

	/* 店铺首页链接 */
	ShopUrl string `json:"shop_url"`
}

/* 交易订单的帐务信息详情 */
type TradeAmount struct {
	/* 支付宝交易号，如：2009112081173831 */
	AlipayNo string `json:"alipay_no"`

	/* 买家货到付款服务费。精确到2位小数;单位:元。如:12.07，表示:12元7分 */
	BuyerCodFee string `json:"buyer_cod_fee"`

	/* 买家获得积分,返点的积分。格式:100;单位:个 */
	BuyerObtainPointFee int64 `json:"buyer_obtain_point_fee"`

	/* 货到付款服务费。精确到2位小数;单位:元。如:12.07，表示:12元7分 */
	CodFee string `json:"cod_fee"`

	/* 交易佣金。精确到2位小数;单位:元。如:200.07，表示:200元7分 */
	CommissionFee string `json:"commission_fee"`

	/* 交易创建时间 */
	Created string `json:"created"`

	/* 交易成功时间(更新交易状态为成功的同时更新)/确认收货时间。格式:yyyy-MM-dd HH:mm:ss */
	EndTime string `json:"end_time"`

	/* 快递代收款。精确到2位小数;单位:元。如:212.07，表示:212元7分 */
	ExpressAgencyFee string `json:"express_agency_fee"`

	/* 子订单的帐务金额详情列表 */
	OrderAmounts []*OrderAmount `json:"order_amounts>order_amount"`

	/* 付款时间。格式:yyyy-MM-dd HH:mm:ss */
	PayTime string `json:"pay_time"`

	/* 主订单实付金额。精确到2位小数;单位:元。如:200.07，表示:200元7分，计算公式如下：
	如果主订单只有一笔子订单 payment = 子订单的实付金额 + 货到付款服务费(如果是货到付款订单) - 主订单的系统级优惠；如果主订单有多笔子订单 payment = 各子订单的实付金额之和 + 货到付款服务费(如果是货到付款订单) + 邮费 - 主订单的系统级优惠 */
	Payment string `json:"payment"`

	/* 邮费。精确到2位小数;单位:元。如:200.07，表示:200元7分。目前只提供整笔交易的邮费，暂不提供各子订单的邮费 */
	PostFee string `json:"post_fee"`

	/* 主交易订单的系统级优惠详情 */
	PromotionDetails []*PromotionDetail `json:"promotion_details>promotion_detail"`

	/* 卖家货到付款服务费。精确到2位小数;单位:元。如:12.07，表示:12元7分 */
	SellerCodFee string `json:"seller_cod_fee"`

	/* 交易主订单编号 */
	Tid int64 `json:"tid"`

	/* 主订单的商品金额（各子订单中商品price * num的和，不包括任何优惠信息）。精确到2位小数;单位:元。如:200.07，表示:200元7分 */
	TotalFee string `json:"total_fee"`
}

/* 食品安全信息，包括：
生产许可证号、产品标准号、厂名、厂址等 */
type FoodSecurity struct {
	/* 厂家联系方式 */
	Contact string `json:"contact"`

	/* 产品标准号 */
	DesignCode string `json:"design_code"`

	/* 厂名 */
	Factory string `json:"factory"`

	/* 厂址 */
	FactorySite string `json:"factory_site"`

	/* 食品添加剂 */
	FoodAdditive string `json:"food_additive"`

	/* 健字号，保健品/膳食营养补充剂 这个类目下特有的信息，此类目下无需填写生产许可证编号（QS），如果填写了生产许可证编号（QS）将被忽略不保存；保存宝贝时，标题前会自动加上健字号产品名称一起作为宝贝标题； */
	HealthProductNo string `json:"health_product_no"`

	/* 配料表 */
	Mix string `json:"mix"`

	/* 保质期 */
	Period string `json:"period"`

	/* 储藏方法 */
	PlanStorage string `json:"plan_storage"`

	/* 生产许可证号 */
	PrdLicenseNo string `json:"prd_license_no"`

	/* 生产结束日期 */
	ProductDateEnd string `json:"product_date_end"`

	/* 生产开始日期 */
	ProductDateStart string `json:"product_date_start"`

	/* 进货结束日期，要在生产日期之后 */
	StockDateEnd string `json:"stock_date_end"`

	/* 进货开始日期，要在生产日期之后 */
	StockDateStart string `json:"stock_date_start"`

	/* 供货商 */
	Supplier string `json:"supplier"`
}

/* 评价信息包括标签信息 */
type TmallRateInfo struct {
	/* 追加评价内容 */
	AppendContent string `json:"append_content"`

	/* 追评中是否含有负向标签 */
	AppendHasNegtv bool `json:"append_has_negtv"`

	/* 追加评价中带有的语义标签列表 */
	AppendTags []*TmallRateTag `json:"append_tags>tmall_rate_tag"`

	/* 追加评价时间 */
	AppendTime string `json:"append_time"`

	/* 评价时间 */
	CommentTime string `json:"comment_time"`

	/* 评价内容 */
	Content string `json:"content"`

	/* 原始评价是否含有负向标签 */
	HasNegtv bool `json:"has_negtv"`

	/* 原始评价对应的标签列表 */
	Tags []*TmallRateTag `json:"tags>tmall_rate_tag"`

	/* 表示评价者的昵称 */
	UserNick string `json:"user_nick"`
}

/* 天猫评价标签对象 */
type TmallRateTag struct {
	/* 表示标签的极性，正极true，负极false */
	Posi bool `json:"posi"`

	/* 表示标签的名称 */
	TagName string `json:"tag_name"`
}

/* 天猫评价标签详细信息 */
type TmallRateTagDetail struct {
	/* 反应该标签的评价数量 */
	Count int64 `json:"count"`

	/* 标签的极性：1正极 -1负极 */
	Posi bool `json:"posi"`

	/* 标签名称 */
	TagName string `json:"tag_name"`
}

/* 卖家设置售后服务对象 */
type AfterSale struct {
	/* id */
	AfterSaleId int64 `json:"after_sale_id"`

	/* 名称 */
	AfterSaleName string `json:"after_sale_name"`

	/* tfs地址 */
	AfterSalePath string `json:"after_sale_path"`
}

/* 产品分销商属性 */
type FenxiaoPdu struct {
	/* 分销商ID */
	DistributorId int64 `json:"distributor_id"`

	/* 分销商用户名 */
	DistributorName string `json:"distributor_name"`

	/* 产品ID */
	ProductId int64 `json:"product_id"`

	/* 产品代销配额库存 */
	QuantityAgent int64 `json:"quantity_agent"`

	/* 产品销售属性值 */
	SkuProperties string `json:"sku_properties"`
}

/* 查询揽送范围之内的物流公司信息 */
type LogisticsPartner struct {
	/* 物流公司揽收和资费详细信息 */
	Carriage *CarriageDetail `json:"carriage"`

	/* 揽收说明信息 */
	CoverRemark string `json:"cover_remark"`

	/* 物流公司详细信息 */
	Partner *PartnerDetail `json:"partner"`

	/* 不可送达的说明信息 */
	UncoverRemark string `json:"uncover_remark"`
}

/* 门票商品操作结果 */
type TicketItemProcessResult struct {
	/* 商品操作流程是否被迫终止。
	如果broken==false则表示成功发布商品，则通过item_id字段获取新生成商品id；
	如果broken==true则表示发布商品遇到问题，则通过broken_reasons字段获取到失败原因。 */
	Broken bool `json:"broken"`

	/* 操作商品遇到的错误原因 */
	BrokenReasons []string `json:"broken_reasons"`

	/* 成功操作的商品标识 */
	ItemId int64 `json:"item_id"`
}

/* 商品优惠策略详情 */
type Promotion struct {
	/* 减价件数，1只减一件，0表示多件 */
	DecreaseNum int64 `json:"decrease_num"`

	/* 优惠类型，PRICE表示按价格优惠，DISCOUNT表示按折扣优惠 */
	DiscountType string `json:"discount_type"`

	/* 优惠额度 */
	DiscountValue string `json:"discount_value"`

	/* 优惠结束日期 */
	EndDate string `json:"end_date"`

	/* 商品数字ID */
	NumIid int64 `json:"num_iid"`

	/* 优惠描述 */
	PromotionDesc string `json:"promotion_desc"`

	/* 优惠ID */
	PromotionId int64 `json:"promotion_id"`

	/* 优惠标题，显示在宝贝详情页面的优惠图标的tip。 */
	PromotionTitle string `json:"promotion_title"`

	/* 优惠开始日期 */
	StartDate string `json:"start_date"`

	/* 优惠策略状态，ACTIVE表示有效，UNACTIVE表示无效 */
	Status string `json:"status"`

	/* 对应的人群标签ID */
	TagId int64 `json:"tag_id"`
}

/* 门票商品电子凭证信息 */
type TicketEtc struct {
	/* 商品电子凭证是否关联本地商户-在门票商品为电子凭证时必选 */
	AssociationStatus bool `json:"association_status"`

	/* 商品电子凭证的自动退款比例-在门票商品为电子凭证时必选 */
	AutoRefund int64 `json:"auto_refund"`

	/* 商品电子凭证的码商-在门票商品为电子凭证时必选 */
	MerchantId string `json:"merchant_id"`

	/* 商品电子凭证的码商名-在门票商品为电子凭证时必选 */
	MerchantNick string `json:"merchant_nick"`

	/* 商品电子凭证的码商-在门票商品为电子凭证时必选 */
	NetworkId string `json:"network_id"`

	/* 商品电子凭证的过期退款比例-在门票商品为电子凭证时必选 */
	OverduePay int64 `json:"overdue_pay"`

	/* 商品电子凭证是否核销打款-在门票商品为电子凭证时必选 */
	VerificationPay bool `json:"verification_pay"`
}

/* 新门票类目商品信息 */
type TicketItem struct {
	/* 商品返点比例（只对B卖家开放，单位为%） */
	AuctionPoint float64 `json:"auction_point,string"`

	/* 商品状态（onsale：上架，instock：仓库） */
	AuctionStatus string `json:"auction_status"`

	/* 类目标识 */
	CatId int64 `json:"cat_id"`

	/* 商品所在地-城市 */
	City string `json:"city"`

	/* 商品描述 */
	Description string `json:"description"`

	/* 商品对应的错误信息。针对get接口使用。 */
	ErrMsg string `json:"err_msg"`

	/* 商品电子凭证信息-在门票商品为电子凭证时必选 */
	Etc *TicketEtc `json:"etc"`

	/* 商品是否有发票（有发票为true，没有发票为false） */
	HaveInvoice bool `json:"have_invoice"`

	/* 商品主图 */
	Image1 string `json:"image_1"`

	/* 商品第一张多图 */
	Image2 string `json:"image_2"`

	/* 商品第二张多图 */
	Image3 string `json:"image_3"`

	/* 商品第三张多图 */
	Image4 string `json:"image_4"`

	/* 商品第四张多图 */
	Image5 string `json:"image_5"`

	/* 商品的标识 */
	ItemId int64 `json:"item_id"`

	/* 商品的上架时间（精确到分，格式为：yyyy-MM-dd HH:mm） */
	ListTime string `json:"list_time"`

	/* 商品的物流运费模板-在产品规格使用到物流时必选 */
	PostageId int64 `json:"postage_id"`

	/* 产品（具有产品规格的）标识 */
	ProductId int64 `json:"product_id"`

	/* 商品是否橱窗推荐（橱窗推荐；true，不推荐：false） */
	PromotedStatus bool `json:"promoted_status"`

	/* 商品所在地-省份 */
	Prov string `json:"prov"`

	/* 宝贝所属的店铺分类列表-店铺分类标识请使用店铺相关接口获取获取，多个店铺分类标识之间通过逗号进行分隔，最多包含10个分类标识 */
	ShopCats string `json:"shop_cats"`

	/* 参见taobao.ticket.item.add文档描述 */
	Skus string `json:"skus"`

	/* 商品是否为拍下减库存（拍下减库存为true，付款减库存为false） */
	SubStockAtBuy bool `json:"sub_stock_at_buy"`

	/* 商品标题 */
	Title string `json:"title"`

	/* 商品视频-视频标识由多媒体中相关接口获取 */
	VideoId int64 `json:"video_id"`

	/* 商品是否参与店铺会员打折 */
	VipPromoted bool `json:"vip_promoted"`
}

/*taobao.appstore.subscribe.get*/
type AppstoreSubscribeGetRequest struct {
	values url.Values
}

func (this *AppstoreSubscribeGetRequest) GetApiMethodName() string {
	return "taobao.appstore.subscribe.get"
}
func (this *AppstoreSubscribeGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *AppstoreSubscribeGetRequest) GetValues() url.Values {
	return this.values
}

/* 用户昵称 */
func (this *AppstoreSubscribeGetRequest) SetNick(value string) {
	this.Set("nick", value)
}

type AppstoreSubscribeGetResponse struct {
	*ErrorResponse                     `json:"error_response,omitempty"`
	AppstoreSubscribeGetResponseResult `json:"appstore_subscribe_get_response"`
}
type AppstoreSubscribeGetResponseResult struct {
	/* 用户订购信息 */
	UserSubscribe *UserSubscribe `json:"user_subscribe"`
}

/*taobao.user.get*/
type UserGetRequest struct {
	values url.Values
}

func (this *UserGetRequest) GetApiMethodName() string {
	return "taobao.user.get"
}
func (this *UserGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *UserGetRequest) GetValues() url.Values {
	return this.values
}

/* 需返回的字段列表。可选值：User结构体中的所有字段；以半角逗号(,)分隔。需要用户授权才能获取用户对应的uid和user_id。 */
func (this *UserGetRequest) SetFields(value string) {
	this.Set("fields", value)
}

/* 用户昵称，如果昵称为中文，请使用UTF-8字符集对昵称进行URL编码。
<br><font color="red">注：在传入session的情况下,可以不传nick，表示取当前用户信息；否则nick必须传.<br>
自用型应用不需要传入nick
</font> */
func (this *UserGetRequest) SetNick(value string) {
	this.Set("nick", value)
}

type UserGetResponse struct {
	*ErrorResponse        `json:"error_response,omitempty"`
	UserGetResponseResult `json:"user_get_response"`
}
type UserGetResponseResult struct {
	/* 用户信息 */
	User *User `json:"user"`
}

/*taobao.user.seller.get*/
type UserSellerGetRequest struct {
	values url.Values
}

func (this *UserSellerGetRequest) GetApiMethodName() string {
	return "taobao.user.seller.get"
}
func (this *UserSellerGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *UserSellerGetRequest) GetValues() url.Values {
	return this.values
}

/* 只返回user_id,nick,sex,seller_credit,type,has_more_pic,item_img_num,item_img_size,prop_img_num,prop_img_size,auto_repost,promoted_type,status,alipay_bind,consumer_protection,avatar,liangpin,sign_food_seller_promise,has_shop,is_lightning_consignment,has_sub_stock,is_golden_seller,vip_info,magazine_subscribe,vertical_market,online_gaming参数 */
func (this *UserSellerGetRequest) SetFields(value string) {
	this.Set("fields", value)
}

type UserSellerGetResponse struct {
	*ErrorResponse              `json:"error_response,omitempty"`
	UserSellerGetResponseResult `json:"user_seller_get_response"`
}
type UserSellerGetResponseResult struct {
	/* 只返回user_id,nick,sex,seller_credit,type,has_more_pic,item_img_num,item_img_size,prop_img_num,prop_img_size,auto_repost,promoted_type,status,alipay_bind,consumer_protection,avatar,liangpin,sign_food_seller_promise,has_shop,is_lightning_consignment,has_sub_stock,is_golden_seller,vip_info,magazine_subscribe,vertical_market,online_gaming */
	User *User `json:"user"`
}

/*taobao.users.get*/
type UsersGetRequest struct {
	values url.Values
}

func (this *UsersGetRequest) GetApiMethodName() string {
	return "taobao.users.get"
}
func (this *UsersGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *UsersGetRequest) GetValues() url.Values {
	return this.values
}

/* 查询字段：User数据结构的公开信息字段列表，以半角逗号(,)分隔 */
func (this *UsersGetRequest) SetFields(value string) {
	this.Set("fields", value)
}

/* 用户昵称，多个以半角逗号(,)分隔，最多40个 */
func (this *UsersGetRequest) SetNicks(value string) {
	this.Set("nicks", value)
}

type UsersGetResponse struct {
	*ErrorResponse         `json:"error_response,omitempty"`
	UsersGetResponseResult `json:"users_get_response"`
}
type UsersGetResponseResult struct {
	/* 用户信息列表 */
	Users struct {
		User []*User `json:"user"`
	} `json:"Users"`
}

/*taobao.itemcats.authorize.get*/
type ItemcatsAuthorizeGetRequest struct {
	values url.Values
}

func (this *ItemcatsAuthorizeGetRequest) GetApiMethodName() string {
	return "taobao.itemcats.authorize.get"
}
func (this *ItemcatsAuthorizeGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *ItemcatsAuthorizeGetRequest) GetValues() url.Values {
	return this.values
}

/* 需要返回的字段。目前支持有：
brand.vid, brand.name,
item_cat.cid, item_cat.name, item_cat.status,item_cat.sort_order,item_cat.parent_cid,item_cat.is_parent,
xinpin_item_cat.cid,
xinpin_item_cat.name,
xinpin_item_cat.status,
xinpin_item_cat.sort_order,
xinpin_item_cat.parent_cid,
xinpin_item_cat.is_parent */
func (this *ItemcatsAuthorizeGetRequest) SetFields(value string) {
	this.Set("fields", value)
}

type ItemcatsAuthorizeGetResponse struct {
	*ErrorResponse                     `json:"error_response,omitempty"`
	ItemcatsAuthorizeGetResponseResult `json:"itemcats_authorize_get_response"`
}
type ItemcatsAuthorizeGetResponseResult struct {
	/* 里面有3个数组：
	Brand[]品牌列表,
	ItemCat[] 类目列表
	XinpinItemCat[] 针对于C卖家新品类目列表 */
	SellerAuthorize *SellerAuthorize `json:"seller_authorize"`
}

/*taobao.itemcats.get*/
type ItemcatsGetRequest struct {
	values url.Values
}

func (this *ItemcatsGetRequest) GetApiMethodName() string {
	return "taobao.itemcats.get"
}
func (this *ItemcatsGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *ItemcatsGetRequest) GetValues() url.Values {
	return this.values
}

/* 商品所属类目ID列表，用半角逗号(,)分隔 例如:(18957,19562,) (cids、parent_cid至少传一个)<br /> 支持最大值为：9223372036854775807<br /> 支持最小值为：0 */
func (this *ItemcatsGetRequest) SetCids(value string) {
	this.Set("cids", value)
}

/* 需要返回的字段列表，见ItemCat，默认返回：cid,parent_cid,name,is_parent；增量类目信息,根据fields传入的参数返回相应的结果。 features字段： 1、如果存在attr_key=freeze表示该类目被冻结了，attr_value=0,5，value可能存在2个值（也可能只有1个），用逗号分割，0表示禁编辑，5表示禁止发布 */
func (this *ItemcatsGetRequest) SetFields(value string) {
	this.Set("fields", value)
}

/* 父商品类目 id，0表示根节点, 传输该参数返回所有子类目。 (cids、parent_cid至少传一个)<br /> 支持最大值为：9223372036854775807<br /> 支持最小值为：0 */
func (this *ItemcatsGetRequest) SetParentCid(value string) {
	this.Set("parent_cid", value)
}

type ItemcatsGetResponse struct {
	*ErrorResponse            `json:"error_response,omitempty"`
	ItemcatsGetResponseResult `json:"itemcats_get_response"`
}
type ItemcatsGetResponseResult struct {
	/* 增量类目信息,根据fields传入的参数返回相应的结果；增量类目信息,根据fields传入的参数返回相应的结果。 features字段： 1、如果存在attr_key=freeze表示该类目被冻结了，attr_value=0,5，value可能存在2个值（也可能只有1个），用逗号分割，0表示禁编辑，5表示禁止发布 */
	ItemCats struct {
		ItemCat []*ItemCat `json:"item_cat"`
	} `json:"ItemCats"`
	/* 最近修改时间(如果取增量，会返回该字段)。格式:yyyy-MM-dd HH:mm:ss */
	LastModified string `json:"last_modified"`
}

/*taobao.itemprops.get*/
type ItempropsGetRequest struct {
	values url.Values
}

func (this *ItempropsGetRequest) GetApiMethodName() string {
	return "taobao.itemprops.get"
}
func (this *ItempropsGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *ItempropsGetRequest) GetValues() url.Values {
	return this.values
}

/* 属性的Key，支持多条，以“,”分隔 */
func (this *ItempropsGetRequest) SetAttrKeys(value string) {
	this.Set("attr_keys", value)
}

/* 类目子属性路径,由该子属性上层的类目属性和类目属性值组成,格式pid:vid;pid:vid.取类目子属性需要传child_path,cid */
func (this *ItempropsGetRequest) SetChildPath(value string) {
	this.Set("child_path", value)
}

/* 叶子类目ID，如果只传cid，则只返回一级属性,通过taobao.itemcats.get获得叶子类目ID */
func (this *ItempropsGetRequest) SetCid(value string) {
	this.Set("cid", value)
}

/* 需要返回的字段列表，见：ItemProp，默认返回：pid, name, must, multi, prop_values */
func (this *ItempropsGetRequest) SetFields(value string) {
	this.Set("fields", value)
}

/* 是否颜色属性。可选值:true(是),false(否) (删除的属性不会匹配和返回这个条件) */
func (this *ItempropsGetRequest) SetIsColorProp(value string) {
	this.Set("is_color_prop", value)
}

/* 是否枚举属性。可选值:true(是),false(否) (删除的属性不会匹配和返回这个条件)。如果返回true，属性值是下拉框选择输入，如果返回false，属性值是用户自行手工输入。 */
func (this *ItempropsGetRequest) SetIsEnumProp(value string) {
	this.Set("is_enum_prop", value)
}

/* 在is_enum_prop是true的前提下，是否是卖家可以自行输入的属性（注：如果is_enum_prop返回false，该参数统一返回false）。可选值:true(是),false(否) (删除的属性不会匹配和返回这个条件) */
func (this *ItempropsGetRequest) SetIsInputProp(value string) {
	this.Set("is_input_prop", value)
}

/* 是否商品属性，这个属性只能放于发布商品时使用。可选值:true(是),false(否) */
func (this *ItempropsGetRequest) SetIsItemProp(value string) {
	this.Set("is_item_prop", value)
}

/* 是否关键属性。可选值:true(是),false(否) */
func (this *ItempropsGetRequest) SetIsKeyProp(value string) {
	this.Set("is_key_prop", value)
}

/* 是否销售属性。可选值:true(是),false(否) */
func (this *ItempropsGetRequest) SetIsSaleProp(value string) {
	this.Set("is_sale_prop", value)
}

/* 父属性ID */
func (this *ItempropsGetRequest) SetParentPid(value string) {
	this.Set("parent_pid", value)
}

/* 属性id (取类目属性时，传pid，不用同时传PID和parent_pid) */
func (this *ItempropsGetRequest) SetPid(value string) {
	this.Set("pid", value)
}

/* 获取类目的类型：1代表集市、2代表天猫<br /> 支持最大值为：2<br /> 支持最小值为：1 */
func (this *ItempropsGetRequest) SetType(value string) {
	this.Set("type", value)
}

type ItempropsGetResponse struct {
	*ErrorResponse             `json:"error_response,omitempty"`
	ItempropsGetResponseResult `json:"itemprops_get_response"`
}
type ItempropsGetResponseResult struct {
	/* 类目属性信息(如果是取全量或者增量，不包括属性值),根据fields传入的参数返回相应的结果 */
	ItemProps struct {
		ItemProp []*ItemProp `json:"item_prop"`
	} `json:"ItemProps"`
	/* 最近修改时间(只有取全量或增量的时候会返回该字段)。格式:yyyy-MM-dd HH:mm:ss */
	LastModified string `json:"last_modified"`
}

/*taobao.itempropvalues.get*/
type ItempropvaluesGetRequest struct {
	values url.Values
}

func (this *ItempropvaluesGetRequest) GetApiMethodName() string {
	return "taobao.itempropvalues.get"
}
func (this *ItempropvaluesGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *ItempropvaluesGetRequest) GetValues() url.Values {
	return this.values
}

/* 属性的Key，支持多条，以“,”分隔 */
func (this *ItempropvaluesGetRequest) SetAttrKeys(value string) {
	this.Set("attr_keys", value)
}

/* 叶子类目ID ,通过taobao.itemcats.get获得叶子类目ID */
func (this *ItempropvaluesGetRequest) SetCid(value string) {
	this.Set("cid", value)
}

/* 需要返回的字段。目前支持有：cid,pid,prop_name,vid,name,name_alias,status,sort_order */
func (this *ItempropvaluesGetRequest) SetFields(value string) {
	this.Set("fields", value)
}

/* 属性和属性值 id串，格式例如(pid1;pid2)或(pid1:vid1;pid2:vid2)或(pid1;pid2:vid2) */
func (this *ItempropvaluesGetRequest) SetPvs(value string) {
	this.Set("pvs", value)
}

/* 获取类目的类型：1代表集市、2代表天猫<br /> 支持最大值为：2<br /> 支持最小值为：1 */
func (this *ItempropvaluesGetRequest) SetType(value string) {
	this.Set("type", value)
}

type ItempropvaluesGetResponse struct {
	*ErrorResponse                  `json:"error_response,omitempty"`
	ItempropvaluesGetResponseResult `json:"itempropvalues_get_response"`
}
type ItempropvaluesGetResponseResult struct {
	/* 最近修改时间。格式:yyyy-MM-dd HH:mm:ss */
	LastModified string `json:"last_modified"`
	/* 属性值,根据fields传入的参数返回相应的结果 */
	PropValues struct {
		PropValue []*PropValue `json:"prop_value"`
	} `json:"PropValues"`
}

/*taobao.aftersale.get*/
type AftersaleGetRequest struct {
	values url.Values
}

func (this *AftersaleGetRequest) GetApiMethodName() string {
	return "taobao.aftersale.get"
}
func (this *AftersaleGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *AftersaleGetRequest) GetValues() url.Values {
	return this.values
}

type AftersaleGetResponse struct {
	*ErrorResponse             `json:"error_response,omitempty"`
	AftersaleGetResponseResult `json:"aftersale_get_response"`
}
type AftersaleGetResponseResult struct {
	/* 售后服务返回对象 */
	AfterSales struct {
		AfterSale []*AfterSale `json:"after_sale"`
	} `json:"AfterSales"`
}

/*taobao.item.add*/
type ItemAddRequest struct {
	values url.Values
}

func (this *ItemAddRequest) GetApiMethodName() string {
	return "taobao.item.add"
}
func (this *ItemAddRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *ItemAddRequest) GetValues() url.Values {
	return this.values
}

/* 售后说明模板id */
func (this *ItemAddRequest) SetAfterSaleId(value string) {
	this.Set("after_sale_id", value)
}

/* 商品上传后的状态。可选值:onsale(出售中),instock(仓库中);默认值:onsale */
func (this *ItemAddRequest) SetApproveStatus(value string) {
	this.Set("approve_status", value)
}

/* 商品的积分返点比例。如:5,表示:返点比例0.5%. 注意：返点比例必须是>0的整数，而且最大是90,即为9%.B商家在发布非虚拟商品时，返点必须是 5的倍数，即0.5%的倍数。其它是1的倍数，即0.1%的倍数。无名良品商家发布商品时，复用该字段记录积分宝返点比例，返点必须是对应类目的返点步长的整数倍，默认是5，即0.5%。注意此时该字段值依旧必须是>0的整数，最高值不超过500，即50% */
func (this *ItemAddRequest) SetAuctionPoint(value string) {
	this.Set("auction_point", value)
}

/* 代充商品类型。在代充商品的类目下，不传表示不标记商品类型（交易搜索中就不能通过标记搜到相关的交易了）。可选类型：
no_mark(不做类型标记)
time_card(点卡软件代充)
fee_card(话费软件代充) */
func (this *ItemAddRequest) SetAutoFill(value string) {
	this.Set("auto_fill", value)
}

/* 商品条形码 */
func (this *ItemAddRequest) SetBarcode(value string) {
	this.Set("barcode", value)
}

/* 商品基础色，数据格式为：pid:vid:rvid1,rvid2,rvid3;pid:vid:rvid1;
基础色只支持以下14种颜色：28335//绿色
28338//蓝色
90554//桔色
28324//黄色
28341//黑色
28320//白色
28326//红色
28329//紫色
3232480//粉红色
107121//透明
132069//褐色
28332//浅灰色
3232478//深灰色
130164//花色 */
func (this *ItemAddRequest) SetChangeProp(value string) {
	this.Set("change_prop", value)
}

/* 叶子类目id<br /> 支持最小值为：0 */
func (this *ItemAddRequest) SetCid(value string) {
	this.Set("cid", value)
}

/* 此为货到付款运费模板的ID，对应的JAVA类型是long,如果COD卖家应用了货到付款运费模板，此值要进行设置。
该字段已经废弃 */
func (this *ItemAddRequest) SetCodPostageId(value string) {
	this.Set("cod_postage_id", value)
}

/* 宝贝描述。字数要大于5个字符，小于25000个字符，受违禁词控制<br /> 支持最大长度为：200000<br /> 支持的最大列表长度为：200000 */
func (this *ItemAddRequest) SetDesc(value string) {
	this.Set("desc", value)
}

/* 商品描述模块化，模块列表，具体数据结构参见Item_Desc_Module。详细的使用方法可参考下面FAQ中说明。 */
func (this *ItemAddRequest) SetDescModules(value string) {
	this.Set("desc_modules", value)
}

/* ems费用。取值范围:0.01-999.00;精确到2位小数;单位:元。如:25.07，表示:25元7分 */
func (this *ItemAddRequest) SetEmsFee(value string) {
	this.Set("ems_fee", value)
}

/* 快递费用。取值范围:0.01-999.00;精确到2位小数;单位:元。如:15.07，表示:15元7分 */
func (this *ItemAddRequest) SetExpressFee(value string) {
	this.Set("express_fee", value)
}

/* 厂家联系方式 */
func (this *ItemAddRequest) SetFoodSecurityContact(value string) {
	this.Set("food_security.contact", value)
}

/* 产品标准号 */
func (this *ItemAddRequest) SetFoodSecurityDesignCode(value string) {
	this.Set("food_security.design_code", value)
}

/* 厂名 */
func (this *ItemAddRequest) SetFoodSecurityFactory(value string) {
	this.Set("food_security.factory", value)
}

/* 厂址 */
func (this *ItemAddRequest) SetFoodSecurityFactorySite(value string) {
	this.Set("food_security.factory_site", value)
}

/* 食品添加剂 */
func (this *ItemAddRequest) SetFoodSecurityFoodAdditive(value string) {
	this.Set("food_security.food_additive", value)
}

/* 健字号，保健品/膳食营养补充剂 这个类目下特有的信息，此类目下无需填写生产许可证编号（QS），如果填写了生产许可证编号（QS）将被忽略不保存；保存宝贝时，标题前会自动加上健字号产品名称一起作为宝贝标题； */
func (this *ItemAddRequest) SetFoodSecurityHealthProductNo(value string) {
	this.Set("food_security.health_product_no", value)
}

/* 配料表 */
func (this *ItemAddRequest) SetFoodSecurityMix(value string) {
	this.Set("food_security.mix", value)
}

/* 保质期 */
func (this *ItemAddRequest) SetFoodSecurityPeriod(value string) {
	this.Set("food_security.period", value)
}

/* 储藏方法 */
func (this *ItemAddRequest) SetFoodSecurityPlanStorage(value string) {
	this.Set("food_security.plan_storage", value)
}

/* 生产许可证号 */
func (this *ItemAddRequest) SetFoodSecurityPrdLicenseNo(value string) {
	this.Set("food_security.prd_license_no", value)
}

/* 生产结束日期,格式必须为yyyy-MM-dd */
func (this *ItemAddRequest) SetFoodSecurityProductDateEnd(value string) {
	this.Set("food_security.product_date_end", value)
}

/* 生产开始日期，格式必须为yyyy-MM-dd */
func (this *ItemAddRequest) SetFoodSecurityProductDateStart(value string) {
	this.Set("food_security.product_date_start", value)
}

/* 进货结束日期，要在生产日期之后，格式必须为yyyy-MM-dd */
func (this *ItemAddRequest) SetFoodSecurityStockDateEnd(value string) {
	this.Set("food_security.stock_date_end", value)
}

/* 进货开始日期，要在生产日期之后，格式必须为yyyy-MM-dd */
func (this *ItemAddRequest) SetFoodSecurityStockDateStart(value string) {
	this.Set("food_security.stock_date_start", value)
}

/* 供货商 */
func (this *ItemAddRequest) SetFoodSecuritySupplier(value string) {
	this.Set("food_security.supplier", value)
}

/* 运费承担方式。可选值:seller（卖家承担）,buyer(买家承担);默认值:seller。卖家承担不用设置邮费和postage_id.买家承担的时候，必填邮费和postage_id
如果用户设置了运费模板会优先使用运费模板，否则要同步设置邮费（post_fee,express_fee,ems_fee） */
func (this *ItemAddRequest) SetFreightPayer(value string) {
	this.Set("freight_payer", value)
}

/* 全球购商品采购地（地区/国家）,默认值只在全球购商品采购地（库存类型选择情况生效），地区国家值为（美国, 香港, 日本, 英国, 新西兰, 德国, 韩国, 荷兰, 澳洲, 法国, 意大利, 台湾, 澳门, 加拿大, 瑞士, 西班牙, 泰国, 新加坡, 马来西亚, 菲律宾, 其他）<br /> 支持最大长度为：30<br /> 支持的最大列表长度为：30 */
func (this *ItemAddRequest) SetGlobalStockCountry(value string) {
	this.Set("global_stock_country", value)
}

/* 全球购商品采购地（库存类型），
有两种库存类型：现货和代购
参数值为1时代表现货，值为2时代表代购。注意：使用时请与 全球购商品采购地（地区/国家）配合使用 */
func (this *ItemAddRequest) SetGlobalStockType(value string) {
	this.Set("global_stock_type", value)
}

/* 支持会员打折。可选值:true,false;默认值:false(不打折) */
func (this *ItemAddRequest) SetHasDiscount(value string) {
	this.Set("has_discount", value)
}

/* 是否有发票。可选值:true,false (商城卖家此字段必须为true);默认值:false(无发票) */
func (this *ItemAddRequest) SetHasInvoice(value string) {
	this.Set("has_invoice", value)
}

/* 橱窗推荐。可选值:true,false;默认值:false(不推荐) */
func (this *ItemAddRequest) SetHasShowcase(value string) {
	this.Set("has_showcase", value)
}

/* 是否有保修。可选值:true,false;默认值:false(不保修) */
func (this *ItemAddRequest) SetHasWarranty(value string) {
	this.Set("has_warranty", value)
}

/* 商品主图片。类型:JPG,GIF;最大长度:500K<br /> 支持的文件类型为：gif,jpg,jpeg,png<br /> 支持的最大列表长度为：524288 */
func (this *ItemAddRequest) SetImage(value string) {
	this.Set("image", value)
}

/* 加价(降价)幅度。如果为0，代表系统代理幅度。对于增价拍和荷兰拍来说是加价幅度，对于降价拍来说是降价幅度。 */
func (this *ItemAddRequest) SetIncrement(value string) {
	this.Set("increment", value)
}

/* 用户自行输入的类目属性ID串。结构："pid1,pid2,pid3"，如："20000"（表示品牌） 注：通常一个类目下用户可输入的关键属性不超过1个。 */
func (this *ItemAddRequest) SetInputPids(value string) {
	this.Set("input_pids", value)
}

/* 用户自行输入的子属性名和属性值，结构:"父属性值;一级子属性名;一级子属性值;二级子属性名;自定义输入值,....",如：“耐克;耐克系列;科比系列;科比系列;2K5,Nike乔丹鞋;乔丹系列;乔丹鞋系列;乔丹鞋系列;json5”，多个自定义属性用','分割，input_str需要与input_pids一一对应，注：通常一个类目下用户可输入的关键属性不超过1个。所有属性别名加起来不能超过3999字节 */
func (this *ItemAddRequest) SetInputStr(value string) {
	this.Set("input_str", value)
}

/* 是否是3D */
func (this *ItemAddRequest) SetIs3D(value string) {
	this.Set("is_3D", value)
}

/* 是否在外店显示 */
func (this *ItemAddRequest) SetIsEx(value string) {
	this.Set("is_ex", value)
}

/* 实物闪电发货 */
func (this *ItemAddRequest) SetIsLightningConsignment(value string) {
	this.Set("is_lightning_consignment", value)
}

/* 是否是线下商品。 1：线上商品（默认值）； 2：线上或线下商品； 3：线下商品。 */
func (this *ItemAddRequest) SetIsOffline(value string) {
	this.Set("is_offline", value)
}

/* 是否在淘宝上显示（如果传FALSE，则在淘宝主站无法显示该商品） */
func (this *ItemAddRequest) SetIsTaobao(value string) {
	this.Set("is_taobao", value)
}

/* 商品是否为新品。只有在当前类目开通新品,并且当前用户拥有该类目下发布新品权限时才能设置is_xinpin为true，否则设置true后会返回错误码:isv.invalid-permission:add-xinpin。同时只有一口价全新的宝贝才能设置为新品，否则会返回错误码：isv.invalid-parameter:xinpin。不设置该参数值或设置为false效果一致。 */
func (this *ItemAddRequest) SetIsXinpin(value string) {
	this.Set("is_xinpin", value)
}

/* 表示商品的体积，如果需要使用按体积计费的运费模板，一定要设置这个值。该值的单位为立方米（m3），如果是其他单位，请转换成成立方米。
该值支持两种格式的设置：格式1：bulk:3,单位为立方米(m3),表示直接设置为商品的体积。格式2：length:10;breadth:10;height:10，单位为米（m）。体积和长宽高都支持小数类型。
在传入体积或长宽高时候，不能带单位。体积的单位默认为立方米（m3），长宽高的单位默认为米(m)
该值支持两种格式的设置：格式1：bulk:3,单位为立方米(m3),表示直接设置为商品的体积。格式2：length:10;breadth:10;height:10，单位为米（m） */
func (this *ItemAddRequest) SetItemSize(value string) {
	this.Set("item_size", value)
}

/* 商品的重量，用于按重量计费的运费模板。注意：单位为kg。
只能传入数值类型（包含小数），不能带单位，单位默认为kg。 */
func (this *ItemAddRequest) SetItemWeight(value string) {
	this.Set("item_weight", value)
}

/* 商品文字的字符集。繁体传入"zh_HK"，简体传入"zh_CN"，不传默认为简体 */
func (this *ItemAddRequest) SetLang(value string) {
	this.Set("lang", value)
}

/* 定时上架时间。(时间格式：yyyy-MM-dd HH:mm:ss) */
func (this *ItemAddRequest) SetListTime(value string) {
	this.Set("list_time", value)
}

/* 发布电子凭证宝贝时候表示是否使用邮寄 0: 代表不使用邮寄； 1：代表使用邮寄；如果不设置这个值，代表不使用邮寄 */
func (this *ItemAddRequest) SetLocalityLifeChooseLogis(value string) {
	this.Set("locality_life.choose_logis", value)
}

/* 本地生活电子交易凭证业务，目前此字段只涉及到的信息为有效期;
如果有效期为起止日期类型，此值为2012-08-06,2012-08-16
如果有效期为【购买成功日 至】类型则格式为2012-08-16
如果有效期为天数类型则格式为15 */
func (this *ItemAddRequest) SetLocalityLifeExpirydate(value string) {
	this.Set("locality_life.expirydate", value)
}

/* 码商信息，格式为 码商id:nick */
func (this *ItemAddRequest) SetLocalityLifeMerchant(value string) {
	this.Set("locality_life.merchant", value)
}

/* 网点ID */
func (this *ItemAddRequest) SetLocalityLifeNetworkId(value string) {
	this.Set("locality_life.network_id", value)
}

/* 电子凭证售中自动退款比例，百分比%前的数字，介于1-100之间的整数 */
func (this *ItemAddRequest) SetLocalityLifeOnsaleAutoRefundRatio(value string) {
	this.Set("locality_life.onsale_auto_refund_ratio", value)
}

/* 退款比例，
百分比%前的数字,1-100的正整数值 */
func (this *ItemAddRequest) SetLocalityLifeRefundRatio(value string) {
	this.Set("locality_life.refund_ratio", value)
}

/* 退款码费承担方。发布电子凭证宝贝的时候会增加“退款码费承担方”配置项，可选填：(1)s（卖家承担） (2)b(买家承担) */
func (this *ItemAddRequest) SetLocalityLifeRefundmafee(value string) {
	this.Set("locality_life.refundmafee", value)
}

/* 核销打款
1代表核销打款 0代表非核销打款 */
func (this *ItemAddRequest) SetLocalityLifeVerification(value string) {
	this.Set("locality_life.verification", value)
}

/* 所在地城市。如杭州 。 */
func (this *ItemAddRequest) SetLocationCity(value string) {
	this.Set("location.city", value)
}

/* 所在地省份。如浙江 */
func (this *ItemAddRequest) SetLocationState(value string) {
	this.Set("location.state", value)
}

/* 该宝贝是否支持【7天无理由退货】，卖家选择的值只是一个因素，最终以类目和选择的属性条件来确定是否支持7天。填入字符0，表示不支持；未填写或填人字符1，表示支持7天无理由退货； */
func (this *ItemAddRequest) SetNewprepay(value string) {
	this.Set("newprepay", value)
}

/* 商品数量。取值范围:0-900000000的整数。且需要等于Sku所有数量的和。
拍卖商品中增加拍只能为1，荷兰拍要在[2,500)范围内。<br /> 支持最大值为：900000000<br /> 支持最小值为：0 */
func (this *ItemAddRequest) SetNum(value string) {
	this.Set("num", value)
}

/* 汽车O2O绑定线下服务标记，如不为空，表示关联服务，否则，不关联服务。 */
func (this *ItemAddRequest) SetO2oBindService(value string) {
	this.Set("o2o_bind_service", value)
}

/* 商品外部编码，该字段的最大长度是64个字节 */
func (this *ItemAddRequest) SetOuterId(value string) {
	this.Set("outer_id", value)
}

/* 拍卖宝贝的保证金。对于增价拍和荷兰拍来说保证金有两种模式：淘宝默认模式（首次出价金额的10%），自定义固定保证金（固定冻结金额只能输入不超过30万的正整数），并且保证金只冻结1次。对于降价拍来说保证金只有淘宝默认的（竞拍价格的10% * 竞拍数量），并且每次出价都需要冻结保证金。
对于拍卖宝贝来说，保证金是必须的，但是默认使用淘宝默认保证金模式，只有用户需要使用自定义固定保证金的时候才需要使用到这个参数，如果该参数不传或传入0则代表使用默认。 */
func (this *ItemAddRequest) SetPaimaiInfoDeposit(value string) {
	this.Set("paimai_info.deposit", value)
}

/* 降价拍宝贝的降价周期(分钟)。降价拍宝贝的价格每隔paimai_info.interval时间会下降一次increment。<br /> 支持最大值为：60<br /> 支持最小值为：1 */
func (this *ItemAddRequest) SetPaimaiInfoInterval(value string) {
	this.Set("paimai_info.interval", value)
}

/* 拍卖商品选择的拍卖类型，拍卖类型包括三种：增价拍(1)，荷兰拍(2)和降价拍(3)。<br /> 支持最大值为：3<br /> 支持最小值为：1 */
func (this *ItemAddRequest) SetPaimaiInfoMode(value string) {
	this.Set("paimai_info.mode", value)
}

/* 降价拍宝贝的保留价。对于降价拍来说，paimai_info.reserve必须大于0，且小于price-increment，而且（price-paimai_info.reserve）/increment的计算结果必须为整数 */
func (this *ItemAddRequest) SetPaimaiInfoReserve(value string) {
	this.Set("paimai_info.reserve", value)
}

/* 自定义销售周期的小时数。拍卖宝贝可以自定义销售周期，这里指定销售周期的小时数。注意，该参数只作为输入参数，不能通过taobao.item.get接口获取。<br /> 支持最大值为：48<br /> 支持最小值为：1 */
func (this *ItemAddRequest) SetPaimaiInfoValidHour(value string) {
	this.Set("paimai_info.valid_hour", value)
}

/* 自定义销售周期的分钟数。拍卖宝贝可以自定义销售周期，这里是指定销售周期的分钟数。自定义销售周期的小时数。拍卖宝贝可以自定义销售周期，这里指定销售周期的小时数。注意，该参数只作为输入参数，不能通过taobao.item.get接口获取。<br /> 支持最大值为：59<br /> 支持最小值为：0 */
func (this *ItemAddRequest) SetPaimaiInfoValidMinute(value string) {
	this.Set("paimai_info.valid_minute", value)
}

/* 商品主图需要关联的图片空间的相对url。这个url所对应的图片必须要属于当前用户。pic_path和image只需要传入一个,如果两个都传，默认选择pic_path */
func (this *ItemAddRequest) SetPicPath(value string) {
	this.Set("pic_path", value)
}

/* 平邮费用。取值范围:0.01-999.00;精确到2位小数;单位:元。如:5.07，表示:5元7分. 注:post_fee,express_fee,ems_fee需要一起填写 */
func (this *ItemAddRequest) SetPostFee(value string) {
	this.Set("post_fee", value)
}

/* 宝贝所属的运费模板ID。取值范围：整数且必须是该卖家的运费模板的ID（可通过taobao.delivery.template.get获得当前会话用户的所有邮费模板） */
func (this *ItemAddRequest) SetPostageId(value string) {
	this.Set("postage_id", value)
}

/* 商品价格。取值范围:0-100000000;精确到2位小数;单位:元。如:200.07，表示:200元7分。需要在正确的价格区间内。
拍卖商品对应的起拍价。 */
func (this *ItemAddRequest) SetPrice(value string) {
	this.Set("price", value)
}

/* 商品所属的产品ID(B商家发布商品需要用) */
func (this *ItemAddRequest) SetProductId(value string) {
	this.Set("product_id", value)
}

/* 属性值别名。如pid:vid:别名;pid1:vid1:别名1 ，其中：pid是属性id vid是属性值id。总长度不超过800字节<br /> 支持最大长度为：800<br /> 支持的最大列表长度为：800 */
func (this *ItemAddRequest) SetPropertyAlias(value string) {
	this.Set("property_alias", value)
}

/* 商品属性列表。格式:pid:vid;pid:vid。属性的pid调用taobao.itemprops.get取得，属性值的vid用taobao.itempropvalues.get取得vid。 如果该类目下面没有属性，可以不用填写。如果有属性，必选属性必填，其他非必选属性可以选择不填写.属性不能超过35对。所有属性加起来包括分割符不能超过549字节，单个属性没有限制。 如果有属性是可输入的话，则用字段input_str填入属性的值 */
func (this *ItemAddRequest) SetProps(value string) {
	this.Set("props", value)
}

/* 商品资质信息 */
func (this *ItemAddRequest) SetQualification(value string) {
	this.Set("qualification", value)
}

/* 景区门票在选择订金支付时候，需要交的预订费。传入的值是1到20之间的数值，小数点后最多可以保留两位（多余的部分将做四舍五入的处理）。这个数值表示的是预订费的比例，最终的预订费为 scenic_ticket_book_cost乘一口价除以100 */
func (this *ItemAddRequest) SetScenicTicketBookCost(value string) {
	this.Set("scenic_ticket_book_cost", value)
}

/* 景区门票类宝贝发布时候，当卖家签订了支付宝代扣协议时候，需要选择支付方式：全额支付和订金支付。当scenic_ticket_pay_way为1时表示全额支付，为2时表示订金支付 */
func (this *ItemAddRequest) SetScenicTicketPayWay(value string) {
	this.Set("scenic_ticket_pay_way", value)
}

/* 商品卖点信息，最长150个字符。天猫商家和集市卖家都可用。<br /> 支持最大长度为：150<br /> 支持的最大列表长度为：150 */
func (this *ItemAddRequest) SetSellPoint(value string) {
	this.Set("sell_point", value)
}

/* 是否承诺退换货服务!虚拟商品无须设置此项! */
func (this *ItemAddRequest) SetSellPromise(value string) {
	this.Set("sell_promise", value)
}

/* 商品所属的店铺类目列表。按逗号分隔。结构:",cid1,cid2,...,"，如果店铺类目存在二级类目，必须传入子类目cids。 */
func (this *ItemAddRequest) SetSellerCids(value string) {
	this.Set("seller_cids", value)
}

/* sku层面的条形码，多个SKU情况，与SKU价格库存格式类似，用逗号分隔 */
func (this *ItemAddRequest) SetSkuBarcode(value string) {
	this.Set("sku_barcode", value)
}

/* 家装建材类目，商品SKU的高度，单位为cm，部分类目必选。天猫商家专用。
可选值为："0-15", "15-25", "25-50", "50-60", "60-80", "80-120", "120-160", "160-200"。
数据和SKU一一对应，用,分隔，如：15-25,25-50,25-50 */
func (this *ItemAddRequest) SetSkuHdHeight(value string) {
	this.Set("sku_hd_height", value)
}

/* 家装建材类目，商品SKU的灯头数量，正整数，大于等于3，部分类目必选。天猫商家专用。
数据和SKU一一对应，用,分隔，如：3,5,7 */
func (this *ItemAddRequest) SetSkuHdLampQuantity(value string) {
	this.Set("sku_hd_lamp_quantity", value)
}

/* 家装建材类目，商品SKU的长度，正整数，单位为cm，部分类目必选。天猫商家专用。
数据和SKU一一对应，用,分隔，如：20,30,30 */
func (this *ItemAddRequest) SetSkuHdLength(value string) {
	this.Set("sku_hd_length", value)
}

/* Sku的外部id串，结构如：1234,1342,…
sku_properties, sku_quantities, sku_prices, sku_outer_ids在输入数据时要一一对应，如果没有sku_outer_ids也要写上这个参数，入参是","(这个是两个sku的示列，逗号数应该是sku个数减1)；该参数最大长度是512个字节 */
func (this *ItemAddRequest) SetSkuOuterIds(value string) {
	this.Set("sku_outer_ids", value)
}

/* Sku的价格串，结构如：10.00,5.00,… 精确到2位小数;单位:元。如:200.07，表示:200元7分 */
func (this *ItemAddRequest) SetSkuPrices(value string) {
	this.Set("sku_prices", value)
}

/* 更新的Sku的属性串，调用taobao.itemprops.get获取类目属性，如果属性是销售属性，再用taobao.itempropvalues.get取得vid。格式:一个SKU的pid1:vid;pid2:vid,多个sku之间属性用逗号分隔。该字段内的销售属性（自定义的除外）也需要在props字段填写。sku的销售属性需要一同选取，如:颜色，尺寸。如果新增商品包含了sku，则此字段一定要传入。这个字段的长度要控制在512个字节以内。
如果有自定义销售属性，则格式为pid:vid;pid2:vid2;$pText:vText , 其中$pText:vText为自定义属性。限制：其中$pText的’$’前缀不能少，且pText和vText文本中不可以存在冒号:和分号;以及逗号， */
func (this *ItemAddRequest) SetSkuProperties(value string) {
	this.Set("sku_properties", value)
}

/* Sku的数量串，结构如：num1,num2,num3 如：2,3 */
func (this *ItemAddRequest) SetSkuQuantities(value string) {
	this.Set("sku_quantities", value)
}

/* 此参数暂时不起作用 */
func (this *ItemAddRequest) SetSkuSpecIds(value string) {
	this.Set("sku_spec_ids", value)
}

/* 新旧程度。可选值：new(新)，second(二手)，unused(闲置)。B商家不能发布二手商品。
如果是二手商品，特定类目下属性里面必填新旧成色属性 */
func (this *ItemAddRequest) SetStuffStatus(value string) {
	this.Set("stuff_status", value)
}

/* 商品是否支持拍下减库存:1支持;2取消支持(付款减库存);0(默认)不更改
集市卖家默认拍下减库存;
商城卖家默认付款减库存 */
func (this *ItemAddRequest) SetSubStock(value string) {
	this.Set("sub_stock", value)
}

/* 宝贝标题。不能超过30字符，受违禁词控制。天猫图书管控类目最大允许120字符；<br /> 支持最大长度为：120<br /> 支持的最大列表长度为：120 */
func (this *ItemAddRequest) SetTitle(value string) {
	this.Set("title", value)
}

/* 发布类型。可选值:fixed(一口价),auction(拍卖)。B商家不能发布拍卖商品，而且拍卖商品是没有SKU的。
拍卖商品发布时需要附加拍卖商品信息：拍卖类型(paimai_info.mode，拍卖类型包括三种：增价拍[1]，荷兰拍[2]以及降价拍[3])，商品数量(num)，起拍价(price)，价格幅度(increament)，保证金(paimai_info.deposit)。另外拍卖商品支持自定义销售周期，通过paimai_info.valid_hour和paimai_info.valid_minute来指定。对于降价拍来说需要设置降价周期(paimai_info.interval)和拍卖保留价(paimai_info.reserve)。
注意：通过taobao.item.get接口获取拍卖信息时，会返回除了valid_hour和valid_minute之外的所有拍卖信息。 */
func (this *ItemAddRequest) SetType(value string) {
	this.Set("type", value)
}

/* 有效期。可选值:7,14;单位:天;默认值:14 */
func (this *ItemAddRequest) SetValidThru(value string) {
	this.Set("valid_thru", value)
}

/* 商品的重量(商超卖家专用字段) */
func (this *ItemAddRequest) SetWeight(value string) {
	this.Set("weight", value)
}

type ItemAddResponse struct {
	*ErrorResponse        `json:"error_response,omitempty"`
	ItemAddResponseResult `json:"item_add_response"`
}
type ItemAddResponseResult struct {
	/* 商品结构,仅有numIid和created返回 */
	Item *Item `json:"item"`
}

/*taobao.item.anchor.get*/
type ItemAnchorGetRequest struct {
	values url.Values
}

func (this *ItemAnchorGetRequest) GetApiMethodName() string {
	return "taobao.item.anchor.get"
}
func (this *ItemAnchorGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *ItemAnchorGetRequest) GetValues() url.Values {
	return this.values
}

/* 对应类目编号 */
func (this *ItemAnchorGetRequest) SetCatId(value string) {
	this.Set("cat_id", value)
}

/* 宝贝模板类型是人工打标还是自动打标：人工打标为1，自动打标为0.人工和自动打标为-1.<br /> 支持最大值为：1<br /> 支持最小值为：-1 */
func (this *ItemAnchorGetRequest) SetType(value string) {
	this.Set("type", value)
}

type ItemAnchorGetResponse struct {
	*ErrorResponse              `json:"error_response,omitempty"`
	ItemAnchorGetResponseResult `json:"item_anchor_get_response"`
}
type ItemAnchorGetResponseResult struct {
	/* 宝贝描述规范化可使用打标模块的锚点信息 */
	AnchorModules struct {
		IdsModule []*IdsModule `json:"ids_module"`
	} `json:"AnchorModules"`
	/* 返回的宝贝描述模板结果数目 */
	TotalResults int64 `json:"total_results"`
}

/*taobao.item.barcode.update*/
type ItemBarcodeUpdateRequest struct {
	values url.Values
}

func (this *ItemBarcodeUpdateRequest) GetApiMethodName() string {
	return "taobao.item.barcode.update"
}
func (this *ItemBarcodeUpdateRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *ItemBarcodeUpdateRequest) GetValues() url.Values {
	return this.values
}

/* 是否强制保存商品条码。
true：强制保存
false ：需要执行条码库校验 */
func (this *ItemBarcodeUpdateRequest) SetIsforce(value string) {
	this.Set("isforce", value)
}

/* 商品条形码，如果不用更新，可选择不填 */
func (this *ItemBarcodeUpdateRequest) SetItemBarcode(value string) {
	this.Set("item_barcode", value)
}

/* 被更新商品的ID */
func (this *ItemBarcodeUpdateRequest) SetItemId(value string) {
	this.Set("item_id", value)
}

/* SKU维度的条形码，和sku_ids字段一一对应，中间以英文逗号分隔 */
func (this *ItemBarcodeUpdateRequest) SetSkuBarcodes(value string) {
	this.Set("sku_barcodes", value)
}

/* 被更新SKU的ID列表，中间以英文逗号进行分隔。如果没有SKU或者不需要更新SKU的条形码，不需要设置 */
func (this *ItemBarcodeUpdateRequest) SetSkuIds(value string) {
	this.Set("sku_ids", value)
}

/* 访问来源，这字段提供给千牛扫码枪用，
其他调用方，不需要填写 */
func (this *ItemBarcodeUpdateRequest) SetSrc(value string) {
	this.Set("src", value)
}

type ItemBarcodeUpdateResponse struct {
	*ErrorResponse                  `json:"error_response,omitempty"`
	ItemBarcodeUpdateResponseResult `json:"item_barcode_update_response"`
}
type ItemBarcodeUpdateResponseResult struct {
	/* 商品结构里的num_iid，modified */
	Item *Item `json:"item"`
}

/*taobao.item.cspu.move*/
type ItemCspuMoveRequest struct {
	values url.Values
}

func (this *ItemCspuMoveRequest) GetApiMethodName() string {
	return "taobao.item.cspu.move"
}
func (this *ItemCspuMoveRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *ItemCspuMoveRequest) GetValues() url.Values {
	return this.values
}

/* 商品id，必填 */
func (this *ItemCspuMoveRequest) SetItemId(value string) {
	this.Set("item_id", value)
}

/* 如果商品存在sku信息，必须输入sku和cspu的映射关系，可以支持多个skuID对应一个cspuID。“-1”代替删除sku。反之不能填写 */
func (this *ItemCspuMoveRequest) SetSkuCspuMapping(value string) {
	this.Set("sku_cspu_mapping", value)
}

/* 商品有SKU信息并且类目配置了营销属性（如套餐），则必须填写sku和套餐的映射关系。并且sku套餐信息和sku_cspu_mapping传入的sku_id信息保持一致。反之，不需要填写， */
func (this *ItemCspuMoveRequest) SetSkuMealpropertyMapping(value string) {
	this.Set("sku_mealproperty_mapping", value)
}

/* 商品需要挂接的目产品ID，如果不填默认不修改产品ID,即商品本身对应的产品ID */
func (this *ItemCspuMoveRequest) SetSpuId(value string) {
	this.Set("spu_id", value)
}

type ItemCspuMoveResponse struct {
	*ErrorResponse             `json:"error_response,omitempty"`
	ItemCspuMoveResponseResult `json:"item_cspu_move_response"`
}
type ItemCspuMoveResponseResult struct {
	/* 商品结构里的num_iid，modified */
	Item *Item `json:"item"`
}

/*taobao.item.delete*/
type ItemDeleteRequest struct {
	values url.Values
}

func (this *ItemDeleteRequest) GetApiMethodName() string {
	return "taobao.item.delete"
}
func (this *ItemDeleteRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *ItemDeleteRequest) GetValues() url.Values {
	return this.values
}

/* 商品数字ID，该参数必须<br /> 支持最小值为：0 */
func (this *ItemDeleteRequest) SetNumIid(value string) {
	this.Set("num_iid", value)
}

type ItemDeleteResponse struct {
	*ErrorResponse           `json:"error_response,omitempty"`
	ItemDeleteResponseResult `json:"item_delete_response"`
}
type ItemDeleteResponseResult struct {
	/* 被删除商品的相关信息 */
	Item *Item `json:"item"`
}

/*taobao.item.get*/
type ItemGetRequest struct {
	values url.Values
}

func (this *ItemGetRequest) GetApiMethodName() string {
	return "taobao.item.get"
}
func (this *ItemGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *ItemGetRequest) GetValues() url.Values {
	return this.values
}

/* 需要返回的商品对象字段，如title,price,desc_modules等。可选值：Item商品结构体中所有字段均可返回；多个字段用“,”分隔。<br>新增返回字段：item_weight(商品的重量，格式为数字，包含小数)、item_size(商品的体积，格式为数字，包含小数)、change_prop（商品基础色数据） */
func (this *ItemGetRequest) SetFields(value string) {
	this.Set("fields", value)
}

/* 商品数字ID<br /> 支持最小值为：1 */
func (this *ItemGetRequest) SetNumIid(value string) {
	this.Set("num_iid", value)
}

/* 商品数字ID(带有跟踪效果) */
func (this *ItemGetRequest) SetTrackIid(value string) {
	this.Set("track_iid", value)
}

type ItemGetResponse struct {
	*ErrorResponse        `json:"error_response,omitempty"`
	ItemGetResponseResult `json:"item_get_response"`
}
type ItemGetResponseResult struct {
	/* 获取的商品 具体字段根据权限和设定的fields决定 */
	Item *Item `json:"item"`
}

/*taobao.item.img.delete*/
type ItemImgDeleteRequest struct {
	values url.Values
}

func (this *ItemImgDeleteRequest) GetApiMethodName() string {
	return "taobao.item.img.delete"
}
func (this *ItemImgDeleteRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *ItemImgDeleteRequest) GetValues() url.Values {
	return this.values
}

/* 商品图片ID */
func (this *ItemImgDeleteRequest) SetId(value string) {
	this.Set("id", value)
}

/* 商品数字ID，必选<br /> 支持最小值为：0 */
func (this *ItemImgDeleteRequest) SetNumIid(value string) {
	this.Set("num_iid", value)
}

type ItemImgDeleteResponse struct {
	*ErrorResponse              `json:"error_response,omitempty"`
	ItemImgDeleteResponseResult `json:"item_img_delete_response"`
}
type ItemImgDeleteResponseResult struct {
	/* 商品图片结构 */
	ItemImg *ItemImg `json:"item_img"`
}

/*taobao.item.img.upload*/
type ItemImgUploadRequest struct {
	values url.Values
}

func (this *ItemImgUploadRequest) GetApiMethodName() string {
	return "taobao.item.img.upload"
}
func (this *ItemImgUploadRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *ItemImgUploadRequest) GetValues() url.Values {
	return this.values
}

/* 商品图片id(如果是更新图片，则需要传该参数) */
func (this *ItemImgUploadRequest) SetId(value string) {
	this.Set("id", value)
}

/* 商品图片内容类型:JPG,GIF;最大:500KB 。支持的文件类型：gif,jpg,jpeg,png<br /> 支持的文件类型为：gif,jpg,jpeg,png,bmp<br /> 支持的最大列表长度为：1048576 */
func (this *ItemImgUploadRequest) SetImage(value string) {
	this.Set("image", value)
}

/* 是否将该图片设为主图,可选值:true,false;默认值:false(非主图) */
func (this *ItemImgUploadRequest) SetIsMajor(value string) {
	this.Set("is_major", value)
}

/* 商品数字ID，该参数必须<br /> 支持最小值为：0 */
func (this *ItemImgUploadRequest) SetNumIid(value string) {
	this.Set("num_iid", value)
}

/* 图片序号 */
func (this *ItemImgUploadRequest) SetPosition(value string) {
	this.Set("position", value)
}

type ItemImgUploadResponse struct {
	*ErrorResponse              `json:"error_response,omitempty"`
	ItemImgUploadResponseResult `json:"item_img_upload_response"`
}
type ItemImgUploadResponseResult struct {
	/* 商品图片结构 */
	ItemImg *ItemImg `json:"item_img"`
}

/*taobao.item.joint.img*/
type ItemJointImgRequest struct {
	values url.Values
}

func (this *ItemJointImgRequest) GetApiMethodName() string {
	return "taobao.item.joint.img"
}
func (this *ItemJointImgRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *ItemJointImgRequest) GetValues() url.Values {
	return this.values
}

/* 商品图片id(如果是更新图片，则需要传该参数) */
func (this *ItemJointImgRequest) SetId(value string) {
	this.Set("id", value)
}

/* 上传的图片是否关联为商品主图（如果需更新主图，则需要传人true） */
func (this *ItemJointImgRequest) SetIsMajor(value string) {
	this.Set("is_major", value)
}

/* 商品数字ID，必选<br /> 支持最小值为：0 */
func (this *ItemJointImgRequest) SetNumIid(value string) {
	this.Set("num_iid", value)
}

/* 图片URL,图片空间图片的相对地址 */
func (this *ItemJointImgRequest) SetPicPath(value string) {
	this.Set("pic_path", value)
}

/* 图片序号 */
func (this *ItemJointImgRequest) SetPosition(value string) {
	this.Set("position", value)
}

type ItemJointImgResponse struct {
	*ErrorResponse             `json:"error_response,omitempty"`
	ItemJointImgResponseResult `json:"item_joint_img_response"`
}
type ItemJointImgResponseResult struct {
	/* 商品图片信息 */
	ItemImg *ItemImg `json:"item_img"`
}

/*taobao.item.joint.propimg*/
type ItemJointPropimgRequest struct {
	values url.Values
}

func (this *ItemJointPropimgRequest) GetApiMethodName() string {
	return "taobao.item.joint.propimg"
}
func (this *ItemJointPropimgRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *ItemJointPropimgRequest) GetValues() url.Values {
	return this.values
}

/* 属性图片ID。如果是新增不需要填写 */
func (this *ItemJointPropimgRequest) SetId(value string) {
	this.Set("id", value)
}

/* 商品数字ID，必选<br /> 支持最小值为：0 */
func (this *ItemJointPropimgRequest) SetNumIid(value string) {
	this.Set("num_iid", value)
}

/* 图片地址(传入图片相对地址即可,即不需包含 http://img02.taobao.net/bao/uploaded ) */
func (this *ItemJointPropimgRequest) SetPicPath(value string) {
	this.Set("pic_path", value)
}

/* 图片序号 */
func (this *ItemJointPropimgRequest) SetPosition(value string) {
	this.Set("position", value)
}

/* 属性列表。调用taobao.itemprops.get获取，属性必须是颜色属性，格式:pid:vid。 */
func (this *ItemJointPropimgRequest) SetProperties(value string) {
	this.Set("properties", value)
}

type ItemJointPropimgResponse struct {
	*ErrorResponse                 `json:"error_response,omitempty"`
	ItemJointPropimgResponseResult `json:"item_joint_propimg_response"`
}
type ItemJointPropimgResponseResult struct {
	/* 属性图片对象信息 */
	PropImg *PropImg `json:"prop_img"`
}

/*taobao.item.price.update*/
type ItemPriceUpdateRequest struct {
	values url.Values
}

func (this *ItemPriceUpdateRequest) GetApiMethodName() string {
	return "taobao.item.price.update"
}
func (this *ItemPriceUpdateRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *ItemPriceUpdateRequest) GetValues() url.Values {
	return this.values
}

/* 售后服务说明模板id */
func (this *ItemPriceUpdateRequest) SetAfterSaleId(value string) {
	this.Set("after_sale_id", value)
}

/* 商品上传后的状态。可选值:onsale（出售中）,instock（库中），如果同时更新商品状态为出售中及list_time为将来的时间，则商品还是处于定时上架的状态, 此时item.is_timing为true */
func (this *ItemPriceUpdateRequest) SetApproveStatus(value string) {
	this.Set("approve_status", value)
}

/* 商品的积分返点比例。如：5 表示返点比例0.5%. 注意：返点比例必须是>0的整数，而且最大是90,即为9%.B商家在发布非虚拟商品时，返点必须是 5的倍数，即0.5%的倍数。其它是1的倍数，即0.1%的倍数。无名良品商家发布商品时，复用该字段记录积分宝返点比例，返点必须是对应类目的返点步长的整数倍，默认是5，即0.5%。注意此时该字段值依旧必须是>0的整数，注意此时该字段值依旧必须是>0的整数，最高值不超过500，即50% */
func (this *ItemPriceUpdateRequest) SetAuctionPoint(value string) {
	this.Set("auction_point", value)
}

/* 代充商品类型。只有少数类目下的商品可以标记上此字段，具体哪些类目可以上传可以通过taobao.itemcat.features.get获得。在代充商品的类目下，不传表示不标记商品类型（交易搜索中就不能通过标记搜到相关的交易了）。可选类型：
no_mark(不做类型标记)
time_card(点卡软件代充)
fee_card(话费软件代充) */
func (this *ItemPriceUpdateRequest) SetAutoFill(value string) {
	this.Set("auto_fill", value)
}

/* 叶子类目id<br /> 支持最小值为：0 */
func (this *ItemPriceUpdateRequest) SetCid(value string) {
	this.Set("cid", value)
}

/* 货到付款运费模板ID */
func (this *ItemPriceUpdateRequest) SetCodPostageId(value string) {
	this.Set("cod_postage_id", value)
}

/* 商品描述. 字数要大于5个字符，小于25000个字符 ，受违禁词控制<br /> 支持最大长度为：200000<br /> 支持的最大列表长度为：200000 */
func (this *ItemPriceUpdateRequest) SetDesc(value string) {
	this.Set("desc", value)
}

/* ems费用。取值范围:0.01-999.00;精确到2位小数;单位:元。如:25.07，表示:25元7分 */
func (this *ItemPriceUpdateRequest) SetEmsFee(value string) {
	this.Set("ems_fee", value)
}

/* 快递费用。取值范围:0.01-999.00;精确到2位小数;单位:元。如:15.07，表示:15元7分 */
func (this *ItemPriceUpdateRequest) SetExpressFee(value string) {
	this.Set("express_fee", value)
}

/* 运费承担方式。运费承担方式。可选值:seller（卖家承担）,buyer(买家承担); */
func (this *ItemPriceUpdateRequest) SetFreightPayer(value string) {
	this.Set("freight_payer", value)
}

/* 支持会员打折。可选值:true,false; */
func (this *ItemPriceUpdateRequest) SetHasDiscount(value string) {
	this.Set("has_discount", value)
}

/* 是否有发票。可选值:true,false (商城卖家此字段必须为true) */
func (this *ItemPriceUpdateRequest) SetHasInvoice(value string) {
	this.Set("has_invoice", value)
}

/* 橱窗推荐。可选值:true,false; */
func (this *ItemPriceUpdateRequest) SetHasShowcase(value string) {
	this.Set("has_showcase", value)
}

/* 是否有保修。可选值:true,false; */
func (this *ItemPriceUpdateRequest) SetHasWarranty(value string) {
	this.Set("has_warranty", value)
}

/* 忽略警告提示. */
func (this *ItemPriceUpdateRequest) SetIgnorewarning(value string) {
	this.Set("ignorewarning", value)
}

/* 商品图片。类型:JPG,GIF;最大长度:500k<br /> 支持的文件类型为：gif,jpg,jpeg,png<br /> 支持的最大列表长度为：524288 */
func (this *ItemPriceUpdateRequest) SetImage(value string) {
	this.Set("image", value)
}

/* 加价幅度 如果为0，代表系统代理幅度 */
func (this *ItemPriceUpdateRequest) SetIncrement(value string) {
	this.Set("increment", value)
}

/* 用户自行输入的类目属性ID串，结构："pid1,pid2,pid3"，如："20000"（表示品牌） 注：通常一个类目下用户可输入的关键属性不超过1个。 */
func (this *ItemPriceUpdateRequest) SetInputPids(value string) {
	this.Set("input_pids", value)
}

/* 用户自行输入的子属性名和属性值，结构:"父属性值;一级子属性名;一级子属性值;二级子属性名;自定义输入值,....",如：“耐克;耐克系列;科比系列;科比系列;2K5,Nike乔丹鞋;乔丹系列;乔丹鞋系列;乔丹鞋系列;json5”，多个自定义属性用','分割，input_str需要与input_pids一一对应，注：通常一个类目下用户可输入的关键属性不超过1个。所有属性别名加起来不能超过3999字节。此处不可以使用“其他”、“其它”和“其她”这三个词。 */
func (this *ItemPriceUpdateRequest) SetInputStr(value string) {
	this.Set("input_str", value)
}

/* 是否是3D */
func (this *ItemPriceUpdateRequest) SetIs3D(value string) {
	this.Set("is_3D", value)
}

/* 是否在外店显示 */
func (this *ItemPriceUpdateRequest) SetIsEx(value string) {
	this.Set("is_ex", value)
}

/* 实物闪电发货。注意：在售的闪电发货产品不允许取消闪电发货，需要先下架商品才能取消闪电发货标记 */
func (this *ItemPriceUpdateRequest) SetIsLightningConsignment(value string) {
	this.Set("is_lightning_consignment", value)
}

/* 是否替换sku */
func (this *ItemPriceUpdateRequest) SetIsReplaceSku(value string) {
	this.Set("is_replace_sku", value)
}

/* 是否在淘宝上显示 */
func (this *ItemPriceUpdateRequest) SetIsTaobao(value string) {
	this.Set("is_taobao", value)
}

/* 商品是否为新品。只有在当前类目开通新品,并且当前用户拥有该类目下发布新品权限时才能设置is_xinpin为true，否则设置true后会返回错误码:isv.invalid-permission:xinpin。同时只有一口价全新的宝贝才能设置为新品，否则会返回错误码：isv.invalid-parameter:xinpin。不设置参数就保持原有值。 */
func (this *ItemPriceUpdateRequest) SetIsXinpin(value string) {
	this.Set("is_xinpin", value)
}

/* 商品文字的版本，繁体传入”zh_HK”，简体传入”zh_CN” */
func (this *ItemPriceUpdateRequest) SetLang(value string) {
	this.Set("lang", value)
}

/* 上架时间。不论是更新架下的商品还是出售中的商品，如果这个字段小于当前时间则直接上架商品，并且上架的时间为更新商品的时间，此时item.is_timing为false，如果大于当前时间则宝贝会下架进入定时上架的宝贝中。 */
func (this *ItemPriceUpdateRequest) SetListTime(value string) {
	this.Set("list_time", value)
}

/* 所在地城市。如杭州 具体可以下载http://dl.open.taobao.com/sdk/商品城市列表.rar 取到 */
func (this *ItemPriceUpdateRequest) SetLocationCity(value string) {
	this.Set("location.city", value)
}

/* 所在地省份。如浙江 具体可以下载http://dl.open.taobao.com/sdk/商品城市列表.rar 取到 */
func (this *ItemPriceUpdateRequest) SetLocationState(value string) {
	this.Set("location.state", value)
}

/* 商品数量，取值范围:0-999999的整数。且需要等于Sku所有数量的和<br /> 支持最大值为：999999<br /> 支持最小值为：0 */
func (this *ItemPriceUpdateRequest) SetNum(value string) {
	this.Set("num", value)
}

/* 商品数字ID，该参数必须<br /> 支持最小值为：1 */
func (this *ItemPriceUpdateRequest) SetNumIid(value string) {
	this.Set("num_iid", value)
}

/* 商家编码 */
func (this *ItemPriceUpdateRequest) SetOuterId(value string) {
	this.Set("outer_id", value)
}

/* 商品主图需要关联的图片空间的相对url。这个url所对应的图片必须要属于当前用户。pic_path和image只需要传入一个,如果两个都传，默认选择pic_path */
func (this *ItemPriceUpdateRequest) SetPicPath(value string) {
	this.Set("pic_path", value)
}

/* 平邮费用。取值范围:0.01-999.00;精确到2位小数;单位:元。如:5.07，表示:5元7分, 注:post_fee,express_fee,ems_fee需一起填写 */
func (this *ItemPriceUpdateRequest) SetPostFee(value string) {
	this.Set("post_fee", value)
}

/* 宝贝所属的运费模板ID。取值范围：整数且必须是该卖家的运费模板的ID（可通过taobao.postages.get获得当前会话用户的所有邮费模板） */
func (this *ItemPriceUpdateRequest) SetPostageId(value string) {
	this.Set("postage_id", value)
}

/* 商品价格。取值范围:0-100000000;精确到2位小数;单位:元。如:200.07，表示:200元7分。需要在正确的价格区间内。 */
func (this *ItemPriceUpdateRequest) SetPrice(value string) {
	this.Set("price", value)
}

/* 商品所属的产品ID(B商家发布商品需要用) */
func (this *ItemPriceUpdateRequest) SetProductId(value string) {
	this.Set("product_id", value)
}

/* 属性值别名。如pid:vid:别名;pid1:vid1:别名1， pid:属性id vid:值id。总长度不超过512字节 */
func (this *ItemPriceUpdateRequest) SetPropertyAlias(value string) {
	this.Set("property_alias", value)
}

/* 商品属性列表。格式:pid:vid;pid:vid。属性的pid调用taobao.itemprops.get取得，属性值的vid用taobao.itempropvalues.get取得vid。 如果该类目下面没有属性，可以不用填写。如果有属性，必选属性必填，其他非必选属性可以选择不填写.属性不能超过35对。所有属性加起来包括分割符不能超过549字节，单个属性没有限制。 如果有属性是可输入的话，则用字段input_str填入属性的值。 */
func (this *ItemPriceUpdateRequest) SetProps(value string) {
	this.Set("props", value)
}

/* 是否承诺退换货服务!虚拟商品无须设置此项! */
func (this *ItemPriceUpdateRequest) SetSellPromise(value string) {
	this.Set("sell_promise", value)
}

/* 重新关联商品与店铺类目，结构:",cid1,cid2,...,"，如果店铺类目存在二级类目，必须传入子类目cids。 */
func (this *ItemPriceUpdateRequest) SetSellerCids(value string) {
	this.Set("seller_cids", value)
}

/* Sku的外部id串，结构如：1234,1342,… sku_properties, sku_quantities, sku_prices, sku_outer_ids在输入数据时要一一对应，如果没有sku_outer_ids也要写上这个参数，入参是","(这个是两个sku的示列，逗号数应该是sku个数减1)；该参数最大长度是512个字节 */
func (this *ItemPriceUpdateRequest) SetSkuOuterIds(value string) {
	this.Set("sku_outer_ids", value)
}

/* 更新的Sku的价格串，结构如：10.00,5.00,… 精确到2位小数;单位:元。如:200.07，表示:200元7分 */
func (this *ItemPriceUpdateRequest) SetSkuPrices(value string) {
	this.Set("sku_prices", value)
}

/* 更新的Sku的属性串，调用taobao.itemprops.get获取类目属性，如果属性是销售属性，再用taobao.itempropvalues.get取得vid。格式:pid:vid;pid:vid。该字段内的销售属性也需要在props字段填写 。如果更新时有对Sku进行操作，则Sku的properties一定要传入。 */
func (this *ItemPriceUpdateRequest) SetSkuProperties(value string) {
	this.Set("sku_properties", value)
}

/* 更新的Sku的数量串，结构如：num1,num2,num3 如:2,3,4 */
func (this *ItemPriceUpdateRequest) SetSkuQuantities(value string) {
	this.Set("sku_quantities", value)
}

/* 商品新旧程度。可选值:new（全新）,unused（闲置）,second（二手）。 */
func (this *ItemPriceUpdateRequest) SetStuffStatus(value string) {
	this.Set("stuff_status", value)
}

/* 商品是否支持拍下减库存:1支持;2取消支持(付款减库存);0(默认)不更改 集市卖家默认拍下减库存; 商城卖家默认付款减库存 */
func (this *ItemPriceUpdateRequest) SetSubStock(value string) {
	this.Set("sub_stock", value)
}

/* 宝贝标题. 不能超过60字符,受违禁词控制<br /> 支持最大长度为：60<br /> 支持的最大列表长度为：60 */
func (this *ItemPriceUpdateRequest) SetTitle(value string) {
	this.Set("title", value)
}

/* 有效期。可选值:7,14;单位:天; */
func (this *ItemPriceUpdateRequest) SetValidThru(value string) {
	this.Set("valid_thru", value)
}

/* 商品的重量(商超卖家专用字段) */
func (this *ItemPriceUpdateRequest) SetWeight(value string) {
	this.Set("weight", value)
}

type ItemPriceUpdateResponse struct {
	*ErrorResponse                `json:"error_response,omitempty"`
	ItemPriceUpdateResponseResult `json:"item_price_update_response"`
}
type ItemPriceUpdateResponseResult struct {
	/* 商品结构里的num_iid，modified */
	Item *Item `json:"item"`
}

/*taobao.item.propimg.delete*/
type ItemPropimgDeleteRequest struct {
	values url.Values
}

func (this *ItemPropimgDeleteRequest) GetApiMethodName() string {
	return "taobao.item.propimg.delete"
}
func (this *ItemPropimgDeleteRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *ItemPropimgDeleteRequest) GetValues() url.Values {
	return this.values
}

/* 商品属性图片ID */
func (this *ItemPropimgDeleteRequest) SetId(value string) {
	this.Set("id", value)
}

/* 商品数字ID，必选<br /> 支持最小值为：0 */
func (this *ItemPropimgDeleteRequest) SetNumIid(value string) {
	this.Set("num_iid", value)
}

type ItemPropimgDeleteResponse struct {
	*ErrorResponse                  `json:"error_response,omitempty"`
	ItemPropimgDeleteResponseResult `json:"item_propimg_delete_response"`
}
type ItemPropimgDeleteResponseResult struct {
	/* 属性图片结构 */
	PropImg *PropImg `json:"prop_img"`
}

/*taobao.item.propimg.upload*/
type ItemPropimgUploadRequest struct {
	values url.Values
}

func (this *ItemPropimgUploadRequest) GetApiMethodName() string {
	return "taobao.item.propimg.upload"
}
func (this *ItemPropimgUploadRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *ItemPropimgUploadRequest) GetValues() url.Values {
	return this.values
}

/* 属性图片ID。如果是新增不需要填写 */
func (this *ItemPropimgUploadRequest) SetId(value string) {
	this.Set("id", value)
}

/* 属性图片内容。类型:JPG,GIF;最大长度:500K;图片大小不超过:1M<br /> 支持的文件类型为：gif,jpg,jpeg,png,bmp<br /> 支持的最大列表长度为：1048576 */
func (this *ItemPropimgUploadRequest) SetImage(value string) {
	this.Set("image", value)
}

/* 商品数字ID，必选<br /> 支持最小值为：0 */
func (this *ItemPropimgUploadRequest) SetNumIid(value string) {
	this.Set("num_iid", value)
}

/* 图片位置 */
func (this *ItemPropimgUploadRequest) SetPosition(value string) {
	this.Set("position", value)
}

/* 属性列表。调用taobao.itemprops.get获取类目属性，属性必须是颜色属性，再用taobao.itempropvalues.get取得vid。格式:pid:vid。 */
func (this *ItemPropimgUploadRequest) SetProperties(value string) {
	this.Set("properties", value)
}

type ItemPropimgUploadResponse struct {
	*ErrorResponse                  `json:"error_response,omitempty"`
	ItemPropimgUploadResponseResult `json:"item_propimg_upload_response"`
}
type ItemPropimgUploadResponseResult struct {
	/* PropImg属性图片结构 */
	PropImg *PropImg `json:"prop_img"`
}

/*taobao.item.quantity.update*/
type ItemQuantityUpdateRequest struct {
	values url.Values
}

func (this *ItemQuantityUpdateRequest) GetApiMethodName() string {
	return "taobao.item.quantity.update"
}
func (this *ItemQuantityUpdateRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *ItemQuantityUpdateRequest) GetValues() url.Values {
	return this.values
}

/* 商品数字ID，必填参数 */
func (this *ItemQuantityUpdateRequest) SetNumIid(value string) {
	this.Set("num_iid", value)
}

/* SKU的商家编码，可选参数。如果不填则默认修改宝贝的库存，如果填了则按照商家编码搜索出对应的SKU并修改库存。当sku_id和本字段都填写时以sku_id为准搜索对应SKU */
func (this *ItemQuantityUpdateRequest) SetOuterId(value string) {
	this.Set("outer_id", value)
}

/* 库存修改值，必选。当全量更新库存时，quantity必须为大于等于0的正整数；当增量更新库存时，quantity为整数，可小于等于0。若增量更新时传入的库存为负数，则负数与实际库存之和不能小于0。比如当前实际库存为1，传入增量更新quantity=-1，库存改为0 */
func (this *ItemQuantityUpdateRequest) SetQuantity(value string) {
	this.Set("quantity", value)
}

/* 要操作的SKU的数字ID，可选。如果不填默认修改宝贝的库存，如果填上则修改该SKU的库存 */
func (this *ItemQuantityUpdateRequest) SetSkuId(value string) {
	this.Set("sku_id", value)
}

/* 库存更新方式，可选。1为全量更新，2为增量更新。如果不填，默认为全量更新 */
func (this *ItemQuantityUpdateRequest) SetType(value string) {
	this.Set("type", value)
}

type ItemQuantityUpdateResponse struct {
	*ErrorResponse                   `json:"error_response,omitempty"`
	ItemQuantityUpdateResponseResult `json:"item_quantity_update_response"`
}
type ItemQuantityUpdateResponseResult struct {
	/* iid、numIid、num和modified，skus中每个sku的skuId、quantity和modified */
	Item *Item `json:"item"`
}

/*taobao.item.recommend.add*/
type ItemRecommendAddRequest struct {
	values url.Values
}

func (this *ItemRecommendAddRequest) GetApiMethodName() string {
	return "taobao.item.recommend.add"
}
func (this *ItemRecommendAddRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *ItemRecommendAddRequest) GetValues() url.Values {
	return this.values
}

/* 商品数字ID，该参数必须<br /> 支持最小值为：0 */
func (this *ItemRecommendAddRequest) SetNumIid(value string) {
	this.Set("num_iid", value)
}

type ItemRecommendAddResponse struct {
	*ErrorResponse                 `json:"error_response,omitempty"`
	ItemRecommendAddResponseResult `json:"item_recommend_add_response"`
}
type ItemRecommendAddResponseResult struct {
	/* 被推荐的商品的信息 */
	Item *Item `json:"item"`
}

/*taobao.item.recommend.delete*/
type ItemRecommendDeleteRequest struct {
	values url.Values
}

func (this *ItemRecommendDeleteRequest) GetApiMethodName() string {
	return "taobao.item.recommend.delete"
}
func (this *ItemRecommendDeleteRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *ItemRecommendDeleteRequest) GetValues() url.Values {
	return this.values
}

/* 商品数字ID，该参数必须<br /> 支持最小值为：0 */
func (this *ItemRecommendDeleteRequest) SetNumIid(value string) {
	this.Set("num_iid", value)
}

type ItemRecommendDeleteResponse struct {
	*ErrorResponse                    `json:"error_response,omitempty"`
	ItemRecommendDeleteResponseResult `json:"item_recommend_delete_response"`
}
type ItemRecommendDeleteResponseResult struct {
	/* 被取消橱窗推荐的商品信息 */
	Item *Item `json:"item"`
}

/*taobao.item.sku.add*/
type ItemSkuAddRequest struct {
	values url.Values
}

func (this *ItemSkuAddRequest) GetApiMethodName() string {
	return "taobao.item.sku.add"
}
func (this *ItemSkuAddRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *ItemSkuAddRequest) GetValues() url.Values {
	return this.values
}

/* 忽略警告提示. */
func (this *ItemSkuAddRequest) SetIgnorewarning(value string) {
	this.Set("ignorewarning", value)
}

/* sku所属商品的价格。当用户新增sku，使商品价格不属于sku价格之间的时候，用于修改商品的价格，使sku能够添加成功 */
func (this *ItemSkuAddRequest) SetItemPrice(value string) {
	this.Set("item_price", value)
}

/* Sku文字的版本。可选值:zh_HK(繁体),zh_CN(简体);默认值:zh_CN */
func (this *ItemSkuAddRequest) SetLang(value string) {
	this.Set("lang", value)
}

/* Sku所属商品数字id，可通过 taobao.item.get 获取。必选<br /> 支持最小值为：0 */
func (this *ItemSkuAddRequest) SetNumIid(value string) {
	this.Set("num_iid", value)
}

/* Sku的商家外部id */
func (this *ItemSkuAddRequest) SetOuterId(value string) {
	this.Set("outer_id", value)
}

/* Sku的销售价格。商品的价格要在商品所有的sku的价格之间。精确到2位小数;单位:元。如:200.07，表示:200元7分 */
func (this *ItemSkuAddRequest) SetPrice(value string) {
	this.Set("price", value)
}

/* Sku属性串。格式:pid:vid;pid:vid,如:1627207:3232483;1630696:3284570,表示:机身颜色:军绿色;手机套餐:一电一充。
如果包含自定义属性则格式为pid:vid;pid2:vid2;$pText:vText , 其中$pText:vText为自定义属性。限制：其中$pText的‘$’前缀不能少，且pText和vText文本中不可以存在 冒号:和分号;以及逗号， */
func (this *ItemSkuAddRequest) SetProperties(value string) {
	this.Set("properties", value)
}

/* Sku的库存数量。sku的总数量应该小于等于商品总数量(Item的NUM)。取值范围:大于零的整数<br /> 支持最小值为：0 */
func (this *ItemSkuAddRequest) SetQuantity(value string) {
	this.Set("quantity", value)
}

/* 家装建材类目，商品SKU的高度，单位为cm，部分类目必选。天猫商家专用。
可选值为："0-15", "15-25", "25-50", "50-60", "60-80", "80-120", "120-160", "160-200"。
数据和SKU一一对应，用,分隔，如：15-25,25-50,25-50 */
func (this *ItemSkuAddRequest) SetSkuHdHeight(value string) {
	this.Set("sku_hd_height", value)
}

/* 家装建材类目，商品SKU的灯头数量，正整数，大于等于3，部分类目必选。天猫商家专用。
数据和SKU一一对应，用,分隔，如：3,5,7 */
func (this *ItemSkuAddRequest) SetSkuHdLampQuantity(value string) {
	this.Set("sku_hd_lamp_quantity", value)
}

/* 家装建材类目，商品SKU的长度，正整数，单位为cm，部分类目必选。天猫商家专用。
数据和SKU一一对应，用,分隔，如：20,30,30 */
func (this *ItemSkuAddRequest) SetSkuHdLength(value string) {
	this.Set("sku_hd_length", value)
}

/* 产品的规格信息 */
func (this *ItemSkuAddRequest) SetSpecId(value string) {
	this.Set("spec_id", value)
}

type ItemSkuAddResponse struct {
	*ErrorResponse           `json:"error_response,omitempty"`
	ItemSkuAddResponseResult `json:"item_sku_add_response"`
}
type ItemSkuAddResponseResult struct {
	/* sku */
	Sku *Sku `json:"sku"`
}

/*taobao.item.sku.delete*/
type ItemSkuDeleteRequest struct {
	values url.Values
}

func (this *ItemSkuDeleteRequest) GetApiMethodName() string {
	return "taobao.item.sku.delete"
}
func (this *ItemSkuDeleteRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *ItemSkuDeleteRequest) GetValues() url.Values {
	return this.values
}

/* 忽略警告提示. */
func (this *ItemSkuDeleteRequest) SetIgnorewarning(value string) {
	this.Set("ignorewarning", value)
}

/* sku所属商品的数量,大于0的整数。当用户删除sku，使商品数量不等于sku数量之和时候，用于修改商品的数量，使sku能够删除成功。特别是删除最后一个sku的时候，一定要设置商品数量到正常的值，否则删除失败 */
func (this *ItemSkuDeleteRequest) SetItemNum(value string) {
	this.Set("item_num", value)
}

/* sku所属商品的价格。当用户删除sku，使商品价格不属于sku价格之间的时候，用于修改商品的价格，使sku能够删除成功 */
func (this *ItemSkuDeleteRequest) SetItemPrice(value string) {
	this.Set("item_price", value)
}

/* Sku文字的版本。可选值:zh_HK(繁体),zh_CN(简体);默认值:zh_CN */
func (this *ItemSkuDeleteRequest) SetLang(value string) {
	this.Set("lang", value)
}

/* Sku所属商品数字id，可通过 taobao.item.get 获取。必选 */
func (this *ItemSkuDeleteRequest) SetNumIid(value string) {
	this.Set("num_iid", value)
}

/* Sku属性串。格式:pid:vid;pid:vid,如: 1627207:3232483;1630696:3284570,表示机身颜色:军绿色;手机套餐:一电一充 */
func (this *ItemSkuDeleteRequest) SetProperties(value string) {
	this.Set("properties", value)
}

type ItemSkuDeleteResponse struct {
	*ErrorResponse              `json:"error_response,omitempty"`
	ItemSkuDeleteResponseResult `json:"item_sku_delete_response"`
}
type ItemSkuDeleteResponseResult struct {
	/* Sku结构 */
	Sku *Sku `json:"sku"`
}

/*taobao.item.sku.get*/
type ItemSkuGetRequest struct {
	values url.Values
}

func (this *ItemSkuGetRequest) GetApiMethodName() string {
	return "taobao.item.sku.get"
}
func (this *ItemSkuGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *ItemSkuGetRequest) GetValues() url.Values {
	return this.values
}

/* 需返回的字段列表。可选值：Sku结构体中的所有字段；字段之间用“,”分隔。 */
func (this *ItemSkuGetRequest) SetFields(value string) {
	this.Set("fields", value)
}

/* 卖家nick(num_iid和nick必传一个)，只传卖家nick时候，该api返回的结果不包含cspu（SKu上的产品规格信息）。 */
func (this *ItemSkuGetRequest) SetNick(value string) {
	this.Set("nick", value)
}

/* 商品的数字IID（num_iid和nick必传一个，推荐用num_iid），传商品的数字id返回的结果里包含cspu（SKu上的产品规格信息）。<br /> 支持最小值为：0 */
func (this *ItemSkuGetRequest) SetNumIid(value string) {
	this.Set("num_iid", value)
}

/* Sku的id。可以通过taobao.item.get得到 */
func (this *ItemSkuGetRequest) SetSkuId(value string) {
	this.Set("sku_id", value)
}

type ItemSkuGetResponse struct {
	*ErrorResponse           `json:"error_response,omitempty"`
	ItemSkuGetResponseResult `json:"item_sku_get_response"`
}
type ItemSkuGetResponseResult struct {
	/* Sku */
	Sku *Sku `json:"sku"`
}

/*taobao.item.sku.price.update*/
type ItemSkuPriceUpdateRequest struct {
	values url.Values
}

func (this *ItemSkuPriceUpdateRequest) GetApiMethodName() string {
	return "taobao.item.sku.price.update"
}
func (this *ItemSkuPriceUpdateRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *ItemSkuPriceUpdateRequest) GetValues() url.Values {
	return this.values
}

/* 忽略警告提示. */
func (this *ItemSkuPriceUpdateRequest) SetIgnorewarning(value string) {
	this.Set("ignorewarning", value)
}

/* sku所属商品的价格。当用户更新sku，使商品价格不属于sku价格之间的时候，用于修改商品的价格，使sku能够更新成功 */
func (this *ItemSkuPriceUpdateRequest) SetItemPrice(value string) {
	this.Set("item_price", value)
}

/* Sku文字的版本。可选值:zh_HK(繁体),zh_CN(简体);默认值:zh_CN */
func (this *ItemSkuPriceUpdateRequest) SetLang(value string) {
	this.Set("lang", value)
}

/* Sku所属商品数字id，可通过 taobao.item.get 获取<br /> 支持最小值为：0 */
func (this *ItemSkuPriceUpdateRequest) SetNumIid(value string) {
	this.Set("num_iid", value)
}

/* Sku的商家外部id */
func (this *ItemSkuPriceUpdateRequest) SetOuterId(value string) {
	this.Set("outer_id", value)
}

/* Sku的销售价格。精确到2位小数;单位:元。如:200.07，表示:200元7分。修改后的sku价格要保证商品的价格在所有sku价格所形成的价格区间内（例如：商品价格为6元，sku价格有5元、10元两种，如果要修改5元sku的价格，那么修改的范围只能是0-6元之间；如果要修改10元的sku，那么修改的范围只能是6到无穷大的区间中） */
func (this *ItemSkuPriceUpdateRequest) SetPrice(value string) {
	this.Set("price", value)
}

/* Sku属性串。格式:pid:vid;pid:vid,如: 1627207:3232483;1630696:3284570,表示机身颜色:军绿色;手机套餐:一电一充 */
func (this *ItemSkuPriceUpdateRequest) SetProperties(value string) {
	this.Set("properties", value)
}

/* Sku的库存数量。sku的总数量应该小于等于商品总数量(Item的NUM)，sku数量变化后item的总数量也会随着变化。取值范围:大于等于零的整数<br /> 支持最小值为：0 */
func (this *ItemSkuPriceUpdateRequest) SetQuantity(value string) {
	this.Set("quantity", value)
}

type ItemSkuPriceUpdateResponse struct {
	*ErrorResponse                   `json:"error_response,omitempty"`
	ItemSkuPriceUpdateResponseResult `json:"item_sku_price_update_response"`
}
type ItemSkuPriceUpdateResponseResult struct {
	/* 商品SKU信息（只包含num_iid和modified） */
	Sku *Sku `json:"sku"`
}

/*taobao.item.sku.update*/
type ItemSkuUpdateRequest struct {
	values url.Values
}

func (this *ItemSkuUpdateRequest) GetApiMethodName() string {
	return "taobao.item.sku.update"
}
func (this *ItemSkuUpdateRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *ItemSkuUpdateRequest) GetValues() url.Values {
	return this.values
}

/* SKU条形码 */
func (this *ItemSkuUpdateRequest) SetBarcode(value string) {
	this.Set("barcode", value)
}

/* 忽略警告提示. */
func (this *ItemSkuUpdateRequest) SetIgnorewarning(value string) {
	this.Set("ignorewarning", value)
}

/* sku所属商品的价格。当用户更新sku，使商品价格不属于sku价格之间的时候，用于修改商品的价格，使sku能够更新成功 */
func (this *ItemSkuUpdateRequest) SetItemPrice(value string) {
	this.Set("item_price", value)
}

/* Sku文字的版本。可选值:zh_HK(繁体),zh_CN(简体);默认值:zh_CN */
func (this *ItemSkuUpdateRequest) SetLang(value string) {
	this.Set("lang", value)
}

/* Sku所属商品数字id，可通过 taobao.item.get 获取<br /> 支持最小值为：0 */
func (this *ItemSkuUpdateRequest) SetNumIid(value string) {
	this.Set("num_iid", value)
}

/* Sku的商家外部id */
func (this *ItemSkuUpdateRequest) SetOuterId(value string) {
	this.Set("outer_id", value)
}

/* Sku的销售价格。精确到2位小数;单位:元。如:200.07，表示:200元7分。修改后的sku价格要保证商品的价格在所有sku价格所形成的价格区间内（例如：商品价格为6元，sku价格有5元、10元两种，如果要修改5元sku的价格，那么修改的范围只能是0-6元之间；如果要修改10元的sku，那么修改的范围只能是6到无穷大的区间中） */
func (this *ItemSkuUpdateRequest) SetPrice(value string) {
	this.Set("price", value)
}

/* Sku属性串。格式:pid:vid;pid:vid,如: 1627207:3232483;1630696:3284570,表示机身颜色:军绿色;手机套餐:一电一充。
如果包含自定义属性，则格式为pid:vid;pid2:vid2;$pText:vText , 其中$pText:vText为自定义属性。限制：其中$pText的’$’前缀不能少，且pText和vText文本中不可以存在 冒号:和分号;以及逗号， */
func (this *ItemSkuUpdateRequest) SetProperties(value string) {
	this.Set("properties", value)
}

/* Sku的库存数量。sku的总数量应该小于等于商品总数量(Item的NUM)，sku数量变化后item的总数量也会随着变化。取值范围:大于等于零的整数<br /> 支持最小值为：0 */
func (this *ItemSkuUpdateRequest) SetQuantity(value string) {
	this.Set("quantity", value)
}

/* 产品的规格信息。 */
func (this *ItemSkuUpdateRequest) SetSpecId(value string) {
	this.Set("spec_id", value)
}

type ItemSkuUpdateResponse struct {
	*ErrorResponse              `json:"error_response,omitempty"`
	ItemSkuUpdateResponseResult `json:"item_sku_update_response"`
}
type ItemSkuUpdateResponseResult struct {
	/* 商品Sku */
	Sku *Sku `json:"sku"`
}

/*taobao.item.skus.get*/
type ItemSkusGetRequest struct {
	values url.Values
}

func (this *ItemSkusGetRequest) GetApiMethodName() string {
	return "taobao.item.skus.get"
}
func (this *ItemSkusGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *ItemSkusGetRequest) GetValues() url.Values {
	return this.values
}

/* 需返回的字段列表。可选值：Sku结构体中的所有字段；字段之间用“,”分隔。 */
func (this *ItemSkusGetRequest) SetFields(value string) {
	this.Set("fields", value)
}

/* sku所属商品数字id，必选。num_iid个数不能超过40个 */
func (this *ItemSkusGetRequest) SetNumIids(value string) {
	this.Set("num_iids", value)
}

type ItemSkusGetResponse struct {
	*ErrorResponse            `json:"error_response,omitempty"`
	ItemSkusGetResponseResult `json:"item_skus_get_response"`
}
type ItemSkusGetResponseResult struct {
	/* Sku列表 */
	Skus struct {
		Sku []*Sku `json:"sku"`
	} `json:"Skus"`
}

/*taobao.item.templates.get*/
type ItemTemplatesGetRequest struct {
	values url.Values
}

func (this *ItemTemplatesGetRequest) GetApiMethodName() string {
	return "taobao.item.templates.get"
}
func (this *ItemTemplatesGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *ItemTemplatesGetRequest) GetValues() url.Values {
	return this.values
}

type ItemTemplatesGetResponse struct {
	*ErrorResponse                 `json:"error_response,omitempty"`
	ItemTemplatesGetResponseResult `json:"item_templates_get_response"`
}
type ItemTemplatesGetResponseResult struct {
	/* 返回宝贝模板对象。包含模板id，模板name，还有模板的类别（0表示外店，1表示内店） */
	ItemTemplateList struct {
		ItemTemplate []*ItemTemplate `json:"item_template"`
	} `json:"ItemTemplateList"`
}

/*taobao.item.update*/
type ItemUpdateRequest struct {
	values url.Values
}

func (this *ItemUpdateRequest) GetApiMethodName() string {
	return "taobao.item.update"
}
func (this *ItemUpdateRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *ItemUpdateRequest) GetValues() url.Values {
	return this.values
}

/* 售后服务说明模板id */
func (this *ItemUpdateRequest) SetAfterSaleId(value string) {
	this.Set("after_sale_id", value)
}

/* 商品上传后的状态。可选值:onsale（出售中）,instock（库中），如果同时更新商品状态为出售中及list_time为将来的时间，则商品还是处于定时上架的状态, 此时item.is_timing为true */
func (this *ItemUpdateRequest) SetApproveStatus(value string) {
	this.Set("approve_status", value)
}

/* 商品的积分返点比例。如：5 表示返点比例0.5%. 注意：返点比例必须是>0的整数，而且最大是90,即为9%.B商家在发布非虚拟商品时，返点必须是 5的倍数，即0.5%的倍数。其它是1的倍数，即0.1%的倍数。无名良品商家发布商品时，复用该字段记录积分宝返点比例，返点必须是对应类目的返点步长的整数倍，默认是5，即0.5%。注意此时该字段值依旧必须是>0的整数，注意此时该字段值依旧必须是>0的整数，最高值不超过500，即50% */
func (this *ItemUpdateRequest) SetAuctionPoint(value string) {
	this.Set("auction_point", value)
}

/* 代充商品类型。只有少数类目下的商品可以标记上此字段，具体哪些类目可以上传可以通过taobao.itemcat.features.get获得。在代充商品的类目下，不传表示不标记商品类型（交易搜索中就不能通过标记搜到相关的交易了）。可选类型：
no_mark(不做类型标记)
time_card(点卡软件代充)
fee_card(话费软件代充) */
func (this *ItemUpdateRequest) SetAutoFill(value string) {
	this.Set("auto_fill", value)
}

/* 商品条形码 */
func (this *ItemUpdateRequest) SetBarcode(value string) {
	this.Set("barcode", value)
}

/* 商品基础色，数据格式为：pid:vid:rvid1,rvid2,rvid3;pid:vid:rvid1 */
func (this *ItemUpdateRequest) SetChangeProp(value string) {
	this.Set("change_prop", value)
}

/* 天猫超市扩展字段，天猫超市专用 */
func (this *ItemUpdateRequest) SetChaoshiExtendsInfo(value string) {
	this.Set("chaoshi_extends_info", value)
}

/* 叶子类目id<br /> 支持最小值为：0 */
func (this *ItemUpdateRequest) SetCid(value string) {
	this.Set("cid", value)
}

/* 货到付款运费模板ID
该字段已经废弃，货到付款模板已经集成到运费模板中。 */
func (this *ItemUpdateRequest) SetCodPostageId(value string) {
	this.Set("cod_postage_id", value)
}

/* 商品描述. 字数要大于5个字符，小于25000个字符 ，受违禁词控制<br /> 支持最大长度为：200000<br /> 支持的最大列表长度为：200000 */
func (this *ItemUpdateRequest) SetDesc(value string) {
	this.Set("desc", value)
}

/* 商品描述模块化，模块列表；数据结构可参考Item_Desc_Module 。详细使用说明：http://open.taobao.com/support/question_detail.htm?spm=0.0.0.0.cRcj3S&id=147498 ； */
func (this *ItemUpdateRequest) SetDescModules(value string) {
	this.Set("desc_modules", value)
}

/* 支持宝贝信息的删除,如需删除对应的食品安全信息中的储藏方法、保质期， 则应该设置此参数的值为：food_security.plan_storage,food_security.period; 各个参数的名称之间用【,】分割, 如果对应的参数有设置过值，即使在这个列表中，也不会被删除; 目前支持此功能的宝贝信息如下：食品安全信息所有字段、电子交易凭证字段（locality_life，locality_life.verification，locality_life.refund_ratio，locality_life.network_id ，locality_life.onsale_auto_refund_ratio）。支持对全球购宝贝信息的清除（字符串中包含global_stock） */
func (this *ItemUpdateRequest) SetEmptyFields(value string) {
	this.Set("empty_fields", value)
}

/* ems费用。取值范围:0.01-999.00;精确到2位小数;单位:元。如:25.07，表示:25元7分 */
func (this *ItemUpdateRequest) SetEmsFee(value string) {
	this.Set("ems_fee", value)
}

/* 快递费用。取值范围:0.01-999.00;精确到2位小数;单位:元。如:15.07，表示:15元7分 */
func (this *ItemUpdateRequest) SetExpressFee(value string) {
	this.Set("express_fee", value)
}

/* 厂家联系方式 */
func (this *ItemUpdateRequest) SetFoodSecurityContact(value string) {
	this.Set("food_security.contact", value)
}

/* 产品标准号 */
func (this *ItemUpdateRequest) SetFoodSecurityDesignCode(value string) {
	this.Set("food_security.design_code", value)
}

/* 厂名 */
func (this *ItemUpdateRequest) SetFoodSecurityFactory(value string) {
	this.Set("food_security.factory", value)
}

/* 厂址 */
func (this *ItemUpdateRequest) SetFoodSecurityFactorySite(value string) {
	this.Set("food_security.factory_site", value)
}

/* 食品添加剂 */
func (this *ItemUpdateRequest) SetFoodSecurityFoodAdditive(value string) {
	this.Set("food_security.food_additive", value)
}

/* 健字号，保健品/膳食营养补充剂 这个类目下特有的信息，此类目下无需填写生产许可证编号（QS），如果填写了生产许可证编号（QS）将被忽略不保存；保存宝贝时，标题前会自动加上健字号产品名称一起作为宝贝标题； */
func (this *ItemUpdateRequest) SetFoodSecurityHealthProductNo(value string) {
	this.Set("food_security.health_product_no", value)
}

/* 配料表 */
func (this *ItemUpdateRequest) SetFoodSecurityMix(value string) {
	this.Set("food_security.mix", value)
}

/* 保质期 */
func (this *ItemUpdateRequest) SetFoodSecurityPeriod(value string) {
	this.Set("food_security.period", value)
}

/* 储藏方法 */
func (this *ItemUpdateRequest) SetFoodSecurityPlanStorage(value string) {
	this.Set("food_security.plan_storage", value)
}

/* 生产许可证号 */
func (this *ItemUpdateRequest) SetFoodSecurityPrdLicenseNo(value string) {
	this.Set("food_security.prd_license_no", value)
}

/* 生产结束日期,格式必须为yyyy-MM-dd */
func (this *ItemUpdateRequest) SetFoodSecurityProductDateEnd(value string) {
	this.Set("food_security.product_date_end", value)
}

/* 生产开始日期，格式必须为yyyy-MM-dd */
func (this *ItemUpdateRequest) SetFoodSecurityProductDateStart(value string) {
	this.Set("food_security.product_date_start", value)
}

/* 进货结束日期，要在生产日期之后，格式必须为yyyy-MM-dd */
func (this *ItemUpdateRequest) SetFoodSecurityStockDateEnd(value string) {
	this.Set("food_security.stock_date_end", value)
}

/* 进货开始日期，要在生产日期之后，格式必须为yyyy-MM-dd */
func (this *ItemUpdateRequest) SetFoodSecurityStockDateStart(value string) {
	this.Set("food_security.stock_date_start", value)
}

/* 供货商 */
func (this *ItemUpdateRequest) SetFoodSecuritySupplier(value string) {
	this.Set("food_security.supplier", value)
}

/* 运费承担方式。运费承担方式。可选值:seller（卖家承担）,buyer(买家承担); */
func (this *ItemUpdateRequest) SetFreightPayer(value string) {
	this.Set("freight_payer", value)
}

/* 全球购商品采购地（地区/国家）,默认值只在全球购商品采购地（库存类型选择情况生效），地区国家值为（美国, 香港, 日本, 英国, 新西兰, 德国, 韩国, 荷兰, 澳洲, 法国, 意大利, 台湾, 澳门, 加拿大, 瑞士, 西班牙, 泰国, 新加坡, 马来西亚, 菲律宾, 其他）<br /> 支持最大长度为：30<br /> 支持的最大列表长度为：30 */
func (this *ItemUpdateRequest) SetGlobalStockCountry(value string) {
	this.Set("global_stock_country", value)
}

/* 全球购商品采购地（库存类型）
全球购商品有两种库存类型：现货和代购 参数值为1时代表现货，值为2时代表代购。注意：使用时请与 全球购商品采购地（地区/国家）配合使用 */
func (this *ItemUpdateRequest) SetGlobalStockType(value string) {
	this.Set("global_stock_type", value)
}

/* 支持会员打折。可选值:true,false; */
func (this *ItemUpdateRequest) SetHasDiscount(value string) {
	this.Set("has_discount", value)
}

/* 是否有发票。可选值:true,false (商城卖家此字段必须为true) */
func (this *ItemUpdateRequest) SetHasInvoice(value string) {
	this.Set("has_invoice", value)
}

/* 橱窗推荐。可选值:true,false; */
func (this *ItemUpdateRequest) SetHasShowcase(value string) {
	this.Set("has_showcase", value)
}

/* 是否有保修。可选值:true,false; */
func (this *ItemUpdateRequest) SetHasWarranty(value string) {
	this.Set("has_warranty", value)
}

/* 忽略警告提示. */
func (this *ItemUpdateRequest) SetIgnorewarning(value string) {
	this.Set("ignorewarning", value)
}

/* 商品图片。类型:JPG,GIF;最大长度:500k<br /> 支持的文件类型为：gif,jpg,jpeg,png<br /> 支持的最大列表长度为：524288 */
func (this *ItemUpdateRequest) SetImage(value string) {
	this.Set("image", value)
}

/* 加价(降价)幅度。如果为0，代表系统代理幅度。对于增价拍和荷兰拍来说是加价幅度，对于降价拍来说是降价幅度。 */
func (this *ItemUpdateRequest) SetIncrement(value string) {
	this.Set("increment", value)
}

/* 用户自行输入的类目属性ID串，结构："pid1,pid2,pid3"，如："20000"（表示品牌） 注：通常一个类目下用户可输入的关键属性不超过1个。 */
func (this *ItemUpdateRequest) SetInputPids(value string) {
	this.Set("input_pids", value)
}

/* 用户自行输入的子属性名和属性值，结构:"父属性值;一级子属性名;一级子属性值;二级子属性名;自定义输入值,....",如：“耐克;耐克系列;科比系列;科比系列;2K5,Nike乔丹鞋;乔丹系列;乔丹鞋系列;乔丹鞋系列;json5”，多个自定义属性用','分割，input_str需要与input_pids一一对应，注：通常一个类目下用户可输入的关键属性不超过1个。所有属性别名加起来不能超过3999字节。此处不可以使用“其他”、“其它”和“其她”这三个词。 */
func (this *ItemUpdateRequest) SetInputStr(value string) {
	this.Set("input_str", value)
}

/* 是否是3D */
func (this *ItemUpdateRequest) SetIs3D(value string) {
	this.Set("is_3D", value)
}

/* 是否在外店显示 */
func (this *ItemUpdateRequest) SetIsEx(value string) {
	this.Set("is_ex", value)
}

/* 实物闪电发货。注意：在售的闪电发货产品不允许取消闪电发货，需要先下架商品才能取消闪电发货标记 */
func (this *ItemUpdateRequest) SetIsLightningConsignment(value string) {
	this.Set("is_lightning_consignment", value)
}

/* 是否是线下商品。 1：线上商品（默认值）； 2：线上或线下商品； 3：线下商品。 */
func (this *ItemUpdateRequest) SetIsOffline(value string) {
	this.Set("is_offline", value)
}

/* 是否替换sku */
func (this *ItemUpdateRequest) SetIsReplaceSku(value string) {
	this.Set("is_replace_sku", value)
}

/* 是否在淘宝上显示（如果传FALSE，则在淘宝主站无法显示该商品） */
func (this *ItemUpdateRequest) SetIsTaobao(value string) {
	this.Set("is_taobao", value)
}

/* 商品是否为新品。只有在当前类目开通新品,并且当前用户拥有该类目下发布新品权限时才能设置is_xinpin为true，否则设置true后会返回错误码:isv.invalid-permission:xinpin。同时只有一口价全新的宝贝才能设置为新品，否则会返回错误码：isv.invalid-parameter:xinpin。不设置参数就保持原有值。 */
func (this *ItemUpdateRequest) SetIsXinpin(value string) {
	this.Set("is_xinpin", value)
}

/* 表示商品的体积，如果需要使用按体积计费的运费模板，一定要设置这个值。该值的单位为立方米（m3），如果是其他单位，请转换成成立方米。
该值支持两种格式的设置：格式1：bulk:3,单位为立方米(m3),表示直接设置为商品的体积。格式2：length:10;breadth:10;height:10，单位为米（m）。体积和长宽高都支持小数类型。
在传入体积或长宽高时候，不能带单位。体积的单位默认为立方米（m3），长宽高的单位默认为米(m)
在编辑的时候，如果需要删除体积属性，请设置该值为0，如bulk:0 */
func (this *ItemUpdateRequest) SetItemSize(value string) {
	this.Set("item_size", value)
}

/* 商品的重量，用于按重量计费的运费模板。注意：单位为kg。 只能传入数值类型（包含小数），不能带单位，单位默认为kg。 在编辑时候，如果需要在商品里删除重量的信息，就需要将值设置为0 */
func (this *ItemUpdateRequest) SetItemWeight(value string) {
	this.Set("item_weight", value)
}

/* 商品文字的版本，繁体传入”zh_HK”，简体传入”zh_CN” */
func (this *ItemUpdateRequest) SetLang(value string) {
	this.Set("lang", value)
}

/* 上架时间。大于当前时间则宝贝会下架进入定时上架的宝贝中。 */
func (this *ItemUpdateRequest) SetListTime(value string) {
	this.Set("list_time", value)
}

/* 编辑电子凭证宝贝时候表示是否使用邮寄
0: 代表不使用邮寄；
1：代表使用邮寄；
如果不设置这个值，代表不使用邮寄 */
func (this *ItemUpdateRequest) SetLocalityLifeChooseLogis(value string) {
	this.Set("locality_life.choose_logis", value)
}

/* 本地生活电子交易凭证业务，目前此字段只涉及到的信息为有效期;
如果有效期为起止日期类型，此值为2012-08-06,2012-08-16
如果有效期为【购买成功日 至】类型则格式为2012-08-16
如果有效期为天数类型则格式为15 */
func (this *ItemUpdateRequest) SetLocalityLifeExpirydate(value string) {
	this.Set("locality_life.expirydate", value)
}

/* 码商信息，格式为 码商id:nick */
func (this *ItemUpdateRequest) SetLocalityLifeMerchant(value string) {
	this.Set("locality_life.merchant", value)
}

/* 网点ID,在参数empty_fields里设置locality_life.network_id可删除网点ID */
func (this *ItemUpdateRequest) SetLocalityLifeNetworkId(value string) {
	this.Set("locality_life.network_id", value)
}

/* 电子凭证售中自动退款比例，百分比%前的数字，介于1-100之间的整数 */
func (this *ItemUpdateRequest) SetLocalityLifeOnsaleAutoRefundRatio(value string) {
	this.Set("locality_life.onsale_auto_refund_ratio", value)
}

/* 退款比例，百分比%前的数字,1-100的正整数值; 在参数empty_fields里设置locality_life.refund_ratio可删除退款比例 */
func (this *ItemUpdateRequest) SetLocalityLifeRefundRatio(value string) {
	this.Set("locality_life.refund_ratio", value)
}

/* 退款码费承担方。发布电子凭证宝贝的时候会增加“退款码费承担方”配置项，可选填：(1)s（卖家承担） (2)b(买家承担) */
func (this *ItemUpdateRequest) SetLocalityLifeRefundmafee(value string) {
	this.Set("locality_life.refundmafee", value)
}

/* 核销打款,1代表核销打款 0代表非核销打款; 在参数empty_fields里设置locality_life.verification可删除核销打款 */
func (this *ItemUpdateRequest) SetLocalityLifeVerification(value string) {
	this.Set("locality_life.verification", value)
}

/* 所在地城市。如杭州 */
func (this *ItemUpdateRequest) SetLocationCity(value string) {
	this.Set("location.city", value)
}

/* 所在地省份。如浙江 */
func (this *ItemUpdateRequest) SetLocationState(value string) {
	this.Set("location.state", value)
}

/* 该宝贝是否支持【7天无理由退货】，卖家选择的值只是一个因素，最终以类目和选择的属性条件来确定是否支持7天。填入字符0，表示不支持；未填写或填人字符1，表示支持7天无理由退货；<br>注意：使用该API修改商品其它属性如标题title时，如需保持商品不支持7天无理由退货状态，该字段需传入0 。 */
func (this *ItemUpdateRequest) SetNewprepay(value string) {
	this.Set("newprepay", value)
}

/* 商品数量，取值范围:0-900000000的整数。且需要等于Sku所有数量的和 拍卖商品中增加拍只能为1，荷兰拍要在[2,500)范围内。<br /> 支持最大值为：900000000<br /> 支持最小值为：0 */
func (this *ItemUpdateRequest) SetNum(value string) {
	this.Set("num", value)
}

/* 商品数字ID，该参数必须<br /> 支持最小值为：1 */
func (this *ItemUpdateRequest) SetNumIid(value string) {
	this.Set("num_iid", value)
}

/* 汽车O2O绑定线下服务标记，如不为空，表示关联服务，否则，不关联服务。 */
func (this *ItemUpdateRequest) SetO2oBindService(value string) {
	this.Set("o2o_bind_service", value)
}

/* 商家编码 */
func (this *ItemUpdateRequest) SetOuterId(value string) {
	this.Set("outer_id", value)
}

/* 拍卖宝贝的保证金。对于增价拍和荷兰拍来说保证金有两种模式：淘宝默认模式（首次出价金额的10%），自定义固定保证金（固定冻结金额只能输入不超过30万的正整数），并且保证金只冻结1次。对于降价拍来说保证金只有淘宝默认的（竞拍价格的10% * 竞拍数量），并且每次出价都需要冻结保证金。
对于拍卖宝贝来说，保证金是必须的，但是默认使用淘宝默认保证金模式，只有用户需要使用自定义固定保证金的时候才需要使用到这个参数。如果该参数不传或传入0则代表使用默认。 */
func (this *ItemUpdateRequest) SetPaimaiInfoDeposit(value string) {
	this.Set("paimai_info.deposit", value)
}

/* 降价拍宝贝的降价周期(分钟)。降价拍宝贝的价格每隔paimai_info.interval时间会下降一次increment。<br /> 支持最大值为：60<br /> 支持最小值为：1 */
func (this *ItemUpdateRequest) SetPaimaiInfoInterval(value string) {
	this.Set("paimai_info.interval", value)
}

/* 拍卖商品选择的拍卖类型，拍卖类型包括三种：增价拍(1)，荷兰拍(2)和降价拍(3)。<br /> 支持最大值为：3<br /> 支持最小值为：1 */
func (this *ItemUpdateRequest) SetPaimaiInfoMode(value string) {
	this.Set("paimai_info.mode", value)
}

/* 降价拍宝贝的保留价。对于降价拍来说，paimai_info.reserve必须大于0，且小于price-increment，而且（price-paimai_info.reserve）/increment的计算结果必须为整数 */
func (this *ItemUpdateRequest) SetPaimaiInfoReserve(value string) {
	this.Set("paimai_info.reserve", value)
}

/* 自定义销售周期的小时数。拍卖宝贝可以自定义销售周期，这里指定销售周期的小时数。自定义销售周期的小时数。拍卖宝贝可以自定义销售周期，这里指定销售周期的小时数。注意，该参数只作为输入参数，不能通过taobao.item.get接口获取。<br /> 支持最大值为：48<br /> 支持最小值为：1 */
func (this *ItemUpdateRequest) SetPaimaiInfoValidHour(value string) {
	this.Set("paimai_info.valid_hour", value)
}

/* 自定义销售周期的分钟数。拍卖宝贝可以自定义销售周期，这里是指定销售周期的分钟数。自定义销售周期的小时数。拍卖宝贝可以自定义销售周期，这里指定销售周期的小时数。注意，该参数只作为输入参数，不能通过taobao.item.get接口获取。<br /> 支持最大值为：59<br /> 支持最小值为：0 */
func (this *ItemUpdateRequest) SetPaimaiInfoValidMinute(value string) {
	this.Set("paimai_info.valid_minute", value)
}

/* 商品主图需要关联的图片空间的相对url。这个url所对应的图片必须要属于当前用户。pic_path和image只需要传入一个,如果两个都传，默认选择pic_path */
func (this *ItemUpdateRequest) SetPicPath(value string) {
	this.Set("pic_path", value)
}

/* 平邮费用。取值范围:0.01-999.00;精确到2位小数;单位:元。如:5.07，表示:5元7分, 注:post_fee,express_fee,ems_fee需一起填写 */
func (this *ItemUpdateRequest) SetPostFee(value string) {
	this.Set("post_fee", value)
}

/* 宝贝所属的运费模板ID。取值范围：整数且必须是该卖家的运费模板的ID（可通过taobao.delivery.templates.get获得当前会话用户的所有邮费模板） */
func (this *ItemUpdateRequest) SetPostageId(value string) {
	this.Set("postage_id", value)
}

/* 商品价格。取值范围:0-100000000;精确到2位小数;单位:元。如:200.07，表示:200元7分。需要在正确的价格区间内。 拍卖商品对应的起拍价。 */
func (this *ItemUpdateRequest) SetPrice(value string) {
	this.Set("price", value)
}

/* 商品所属的产品ID(B商家发布商品需要用) */
func (this *ItemUpdateRequest) SetProductId(value string) {
	this.Set("product_id", value)
}

/* 属性值别名。如pid:vid:别名;pid1:vid1:别名1， pid:属性id vid:值id。总长度不超过800字节<br /> 支持最大长度为：800<br /> 支持的最大列表长度为：800 */
func (this *ItemUpdateRequest) SetPropertyAlias(value string) {
	this.Set("property_alias", value)
}

/* 商品属性列表。格式:pid:vid;pid:vid。属性的pid调用taobao.itemprops.get取得，属性值的vid用taobao.itempropvalues.get取得vid。 如果该类目下面没有属性，可以不用填写。如果有属性，必选属性必填，其他非必选属性可以选择不填写.属性不能超过35对。所有属性加起来包括分割符不能超过549字节，单个属性没有限制。 如果有属性是可输入的话，则用字段input_str填入属性的值。 */
func (this *ItemUpdateRequest) SetProps(value string) {
	this.Set("props", value)
}

/* 商品资质信息 */
func (this *ItemUpdateRequest) SetQualification(value string) {
	this.Set("qualification", value)
}

/* 景区门票在选择订金支付时候，需要交的预订费。传入的值是1到20之间的数值，小数点后最多可以保留两位（多余的部分将做四舍五入的处理）。这个数值表示的是预订费的比例，最终的预订费为 scenic_ticket_book_cost乘一口价除以100 */
func (this *ItemUpdateRequest) SetScenicTicketBookCost(value string) {
	this.Set("scenic_ticket_book_cost", value)
}

/* 景区门票类宝贝编辑时候，当卖家签订了支付宝代扣协议时候，需要选择支付方式：全额支付和订金支付。当scenic_ticket_pay_way为1时表示全额支付，为2时表示订金支付 */
func (this *ItemUpdateRequest) SetScenicTicketPayWay(value string) {
	this.Set("scenic_ticket_pay_way", value)
}

/* 商品卖点信息，最长150个字符。天猫和集市都可用<br /> 支持最大长度为：150<br /> 支持的最大列表长度为：150 */
func (this *ItemUpdateRequest) SetSellPoint(value string) {
	this.Set("sell_point", value)
}

/* 是否承诺退换货服务!虚拟商品无须设置此项! */
func (this *ItemUpdateRequest) SetSellPromise(value string) {
	this.Set("sell_promise", value)
}

/* 重新关联商品与店铺类目，结构:",cid1,cid2,...,"，如果店铺类目存在二级类目，必须传入子类目cids。 */
func (this *ItemUpdateRequest) SetSellerCids(value string) {
	this.Set("seller_cids", value)
}

/* sku层面的条形码，多个SKU情况，与SKU价格库存格式类似，用逗号分隔 */
func (this *ItemUpdateRequest) SetSkuBarcode(value string) {
	this.Set("sku_barcode", value)
}

/* 家装建材类目，商品SKU的高度，单位为cm，部分类目必选。天猫商家专用。
可选值为："0-15", "15-25", "25-50", "50-60", "60-80", "80-120", "120-160", "160-200"。
数据和SKU一一对应，用,分隔，如：15-25,25-50,25-50 */
func (this *ItemUpdateRequest) SetSkuHdHeight(value string) {
	this.Set("sku_hd_height", value)
}

/* 家装建材类目，商品SKU的灯头数量，正整数，大于等于3，部分类目必选。天猫商家专用。
数据和SKU一一对应，用,分隔，如：3,5,7 */
func (this *ItemUpdateRequest) SetSkuHdLampQuantity(value string) {
	this.Set("sku_hd_lamp_quantity", value)
}

/* 家装建材类目，商品SKU的长度，正整数，单位为cm，部分类目必选。天猫商家专用。
数据和SKU一一对应，用,分隔，如：20,30,30 */
func (this *ItemUpdateRequest) SetSkuHdLength(value string) {
	this.Set("sku_hd_length", value)
}

/* Sku的外部id串，结构如：1234,1342,… sku_properties, sku_quantities, sku_prices, sku_outer_ids在输入数据时要一一对应，如果没有sku_outer_ids也要写上这个参数，入参是","(这个是两个sku的示列，逗号数应该是sku个数减1)；该参数最大长度是512个字节 */
func (this *ItemUpdateRequest) SetSkuOuterIds(value string) {
	this.Set("sku_outer_ids", value)
}

/* 更新的Sku的价格串，结构如：10.00,5.00,… 精确到2位小数;单位:元。如:200.07，表示:200元7分 */
func (this *ItemUpdateRequest) SetSkuPrices(value string) {
	this.Set("sku_prices", value)
}

/* 更新的Sku的属性串，调用taobao.itemprops.get获取类目属性，如果属性是销售属性，再用taobao.itempropvalues.get取得vid。格式:pid:vid;pid:vid,多个sku之间用逗号分隔。该字段内的销售属性(自定义的除外)也需要在props字段填写 . 规则：如果该SKU存在旧商品，则修改；否则新增Sku。如果更新时有对Sku进行操作，则Sku的properties一定要传入。如果存在自定义销售属性，则格式为pid:vid;pid2:vid2;$pText:vText，其中$pText:vText为自定义属性。限制：其中$pText的’$’前缀不能少，且pText和vText文本中不可以存在 冒号:和分号;以及逗号 */
func (this *ItemUpdateRequest) SetSkuProperties(value string) {
	this.Set("sku_properties", value)
}

/* 更新的Sku的数量串，结构如：num1,num2,num3 如:2,3,4 */
func (this *ItemUpdateRequest) SetSkuQuantities(value string) {
	this.Set("sku_quantities", value)
}

/* 暂时不可用 */
func (this *ItemUpdateRequest) SetSkuSpecIds(value string) {
	this.Set("sku_spec_ids", value)
}

/* 商品新旧程度。可选值:new（全新）,unused（闲置）,second（二手）。 */
func (this *ItemUpdateRequest) SetStuffStatus(value string) {
	this.Set("stuff_status", value)
}

/* 商品是否支持拍下减库存:1支持;2取消支持(付款减库存);0(默认)不更改
集市卖家默认拍下减库存;
商城卖家默认付款减库存 */
func (this *ItemUpdateRequest) SetSubStock(value string) {
	this.Set("sub_stock", value)
}

/* 宝贝标题. 不能超过30字符,受违禁词控制<br /> 支持最大长度为：120<br /> 支持的最大列表长度为：120 */
func (this *ItemUpdateRequest) SetTitle(value string) {
	this.Set("title", value)
}

/* 有效期。可选值:7,14;单位:天; */
func (this *ItemUpdateRequest) SetValidThru(value string) {
	this.Set("valid_thru", value)
}

/* 商品的重量(商超卖家专用字段) */
func (this *ItemUpdateRequest) SetWeight(value string) {
	this.Set("weight", value)
}

type ItemUpdateResponse struct {
	*ErrorResponse           `json:"error_response,omitempty"`
	ItemUpdateResponseResult `json:"item_update_response"`
}
type ItemUpdateResponseResult struct {
	/* 商品结构里的num_iid，modified */
	Item *Item `json:"item"`
}

/*taobao.item.update.delisting*/
type ItemUpdateDelistingRequest struct {
	values url.Values
}

func (this *ItemUpdateDelistingRequest) GetApiMethodName() string {
	return "taobao.item.update.delisting"
}
func (this *ItemUpdateDelistingRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *ItemUpdateDelistingRequest) GetValues() url.Values {
	return this.values
}

/* 商品数字ID，该参数必须<br /> 支持最小值为：0 */
func (this *ItemUpdateDelistingRequest) SetNumIid(value string) {
	this.Set("num_iid", value)
}

type ItemUpdateDelistingResponse struct {
	*ErrorResponse                    `json:"error_response,omitempty"`
	ItemUpdateDelistingResponseResult `json:"item_update_delisting_response"`
}
type ItemUpdateDelistingResponseResult struct {
	/* 返回商品更新信息：返回的结果是:num_iid和modified */
	Item *Item `json:"item"`
}

/*taobao.item.update.listing*/
type ItemUpdateListingRequest struct {
	values url.Values
}

func (this *ItemUpdateListingRequest) GetApiMethodName() string {
	return "taobao.item.update.listing"
}
func (this *ItemUpdateListingRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *ItemUpdateListingRequest) GetValues() url.Values {
	return this.values
}

/* 需要上架的商品的数量。取值范围:大于零的整数。如果商品有sku，则上架数量默认为所有sku数量总和，不可修改。否则商品数量根据设置数量调整为num<br /> 支持最小值为：0 */
func (this *ItemUpdateListingRequest) SetNum(value string) {
	this.Set("num", value)
}

/* 商品数字ID，该参数必须<br /> 支持最小值为：0 */
func (this *ItemUpdateListingRequest) SetNumIid(value string) {
	this.Set("num_iid", value)
}

type ItemUpdateListingResponse struct {
	*ErrorResponse                  `json:"error_response,omitempty"`
	ItemUpdateListingResponseResult `json:"item_update_listing_response"`
}
type ItemUpdateListingResponseResult struct {
	/* 上架后返回的商品信息：返回的结果就是:num_iid和modified */
	Item *Item `json:"item"`
}

/*taobao.items.custom.get*/
type ItemsCustomGetRequest struct {
	values url.Values
}

func (this *ItemsCustomGetRequest) GetApiMethodName() string {
	return "taobao.items.custom.get"
}
func (this *ItemsCustomGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *ItemsCustomGetRequest) GetValues() url.Values {
	return this.values
}

/* 需返回的字段列表。可选值：Item商品结构体中的所有字段；多个字段之间用“,”分隔。如果想返回整个子对象，那字段为item_img，如果是想返回子对象里面的字段，那字段为item_img.url。新增返回字段：one_station标记商品是否淘1站商品 */
func (this *ItemsCustomGetRequest) SetFields(value string) {
	this.Set("fields", value)
}

/* 商品的外部商品ID，支持批量，最多不超过40个。 */
func (this *ItemsCustomGetRequest) SetOuterId(value string) {
	this.Set("outer_id", value)
}

type ItemsCustomGetResponse struct {
	*ErrorResponse               `json:"error_response,omitempty"`
	ItemsCustomGetResponseResult `json:"items_custom_get_response"`
}
type ItemsCustomGetResponseResult struct {
	/* 商品列表，具体返回字段以fields决定 */
	Items struct {
		Item []*Item `json:"item"`
	} `json:"Items"`
}

/*taobao.items.inventory.get*/
type ItemsInventoryGetRequest struct {
	values url.Values
}

func (this *ItemsInventoryGetRequest) GetApiMethodName() string {
	return "taobao.items.inventory.get"
}
func (this *ItemsInventoryGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *ItemsInventoryGetRequest) GetValues() url.Values {
	return this.values
}

/* 分类字段。可选值:<br>
regular_shelved(定时上架)<br>
never_on_shelf(从未上架)<br>
off_shelf(我下架的)<br>
<font color='red'>for_shelved(等待所有上架)<br>
sold_out(全部卖完)<br>
violation_off_shelf(违规下架的)<br>
默认查询for_shelved(等待所有上架)这个状态的商品<br></font>
注：for_shelved(等待所有上架)=regular_shelved(定时上架)+never_on_shelf(从未上架)+off_shelf(我下架的) */
func (this *ItemsInventoryGetRequest) SetBanner(value string) {
	this.Set("banner", value)
}

/* 商品类目ID。ItemCat中的cid字段。可以通过taobao.itemcats.get取到<br /> 支持最小值为：0 */
func (this *ItemsInventoryGetRequest) SetCid(value string) {
	this.Set("cid", value)
}

/* 商品结束修改时间 */
func (this *ItemsInventoryGetRequest) SetEndModified(value string) {
	this.Set("end_modified", value)
}

/* 需返回的字段列表。可选值为 Item 商品结构体中的以下字段：
approve_status,num_iid,title,nick,type,cid,pic_url,num,props,valid_thru, list_time,price,has_discount,has_invoice,has_warranty,has_showcase, modified,delist_time,postage_id,seller_cids,outer_id；字段之间用“,”分隔。<br>
不支持其他字段，如果需要获取其他字段数据，调用taobao.item.get。 */
func (this *ItemsInventoryGetRequest) SetFields(value string) {
	this.Set("fields", value)
}

/* 是否参与会员折扣。可选值：true，false。默认不过滤该条件 */
func (this *ItemsInventoryGetRequest) SetHasDiscount(value string) {
	this.Set("has_discount", value)
}

/* 是否挂接了达尔文标准产品体系。 */
func (this *ItemsInventoryGetRequest) SetIsCspu(value string) {
	this.Set("is_cspu", value)
}

/* 商品是否在外部网店显示 */
func (this *ItemsInventoryGetRequest) SetIsEx(value string) {
	this.Set("is_ex", value)
}

/* 商品是否在淘宝显示 */
func (this *ItemsInventoryGetRequest) SetIsTaobao(value string) {
	this.Set("is_taobao", value)
}

/* 排序方式。格式为column:asc/desc ，column可选值:list_time(上架时间),delist_time(下架时间),num(商品数量)，modified(最近修改时间);默认上架时间降序(即最新上架排在前面)。如按照上架时间降序排序方式为list_time:desc */
func (this *ItemsInventoryGetRequest) SetOrderBy(value string) {
	this.Set("order_by", value)
}

/* 页码。取值范围:大于零小于等于101的整数;默认值为1，即返回第一页数据。当页码超过101页时系统就会报错，故请大家在用此接口获取数据时尽可能的细化自己的搜索条件，例如根据修改时间分段获取商品。<br /> 支持最大值为：101 */
func (this *ItemsInventoryGetRequest) SetPageNo(value string) {
	this.Set("page_no", value)
}

/* 每页条数。取值范围:大于零的整数;最大值：200；默认值：40。 */
func (this *ItemsInventoryGetRequest) SetPageSize(value string) {
	this.Set("page_size", value)
}

/* 搜索字段。搜索商品的title。 */
func (this *ItemsInventoryGetRequest) SetQ(value string) {
	this.Set("q", value)
}

/* 卖家店铺内自定义类目ID。多个之间用“,”分隔。可以根据taobao.sellercats.list.get获得.(<font color="red">注：目前最多支持32个ID号传入</font>) */
func (this *ItemsInventoryGetRequest) SetSellerCids(value string) {
	this.Set("seller_cids", value)
}

/* 商品起始修改时间 */
func (this *ItemsInventoryGetRequest) SetStartModified(value string) {
	this.Set("start_modified", value)
}

type ItemsInventoryGetResponse struct {
	*ErrorResponse                  `json:"error_response,omitempty"`
	ItemsInventoryGetResponseResult `json:"items_inventory_get_response"`
}
type ItemsInventoryGetResponseResult struct {
	/* 搜索到底商品列表，具体字段根据设定的fields决定，不包括desc,stuff_status字段 */
	Items struct {
		Item []*Item `json:"item"`
	} `json:"Items"`
	/* 搜索到符合条件的结果总数 */
	TotalResults int64 `json:"total_results"`
}

/*taobao.items.list.get*/
type ItemsListGetRequest struct {
	values url.Values
}

func (this *ItemsListGetRequest) GetApiMethodName() string {
	return "taobao.items.list.get"
}
func (this *ItemsListGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *ItemsListGetRequest) GetValues() url.Values {
	return this.values
}

/* 需要返回的商品对象字段。可选值：Item商品结构体中字段均可返回(除item_weight,item_size)；多个字段用“,”分隔。如果想返回整个子对象，那字段为itemimg，如果是想返回子对象里面的字段，那字段为itemimg.url。 */
func (this *ItemsListGetRequest) SetFields(value string) {
	this.Set("fields", value)
}

/* 商品数字id列表，多个num_iid用逗号隔开，一次不超过20个。 */
func (this *ItemsListGetRequest) SetNumIids(value string) {
	this.Set("num_iids", value)
}

/* 商品数字id列表，多个track_iid用逗号隔开，一次不超过20个。 */
func (this *ItemsListGetRequest) SetTrackIids(value string) {
	this.Set("track_iids", value)
}

type ItemsListGetResponse struct {
	*ErrorResponse             `json:"error_response,omitempty"`
	ItemsListGetResponseResult `json:"items_list_get_response"`
}
type ItemsListGetResponseResult struct {
	/* 获取的商品 具体字段根据权限和设定的fields决定 */
	Items struct {
		Item []*Item `json:"item"`
	} `json:"Items"`
}

/*taobao.items.onsale.get*/
type ItemsOnsaleGetRequest struct {
	values url.Values
}

func (this *ItemsOnsaleGetRequest) GetApiMethodName() string {
	return "taobao.items.onsale.get"
}
func (this *ItemsOnsaleGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *ItemsOnsaleGetRequest) GetValues() url.Values {
	return this.values
}

/* 商品类目ID。ItemCat中的cid字段。可以通过taobao.itemcats.get取到<br /> 支持最小值为：0 */
func (this *ItemsOnsaleGetRequest) SetCid(value string) {
	this.Set("cid", value)
}

/* 结束的修改时间 */
func (this *ItemsOnsaleGetRequest) SetEndModified(value string) {
	this.Set("end_modified", value)
}

/* 需返回的字段列表。可选值：Item商品结构体中的以下字段：
approve_status,num_iid,title,nick,type,cid,pic_url,num,props,valid_thru,list_time,price,has_discount,has_invoice,has_warranty,has_showcase,modified,delist_time,postage_id,seller_cids,outer_id；字段之间用“,”分隔。
不支持其他字段，如果需要获取其他字段数据，调用taobao.item.get。 */
func (this *ItemsOnsaleGetRequest) SetFields(value string) {
	this.Set("fields", value)
}

/* 是否参与会员折扣。可选值：true，false。默认不过滤该条件 */
func (this *ItemsOnsaleGetRequest) SetHasDiscount(value string) {
	this.Set("has_discount", value)
}

/* 是否橱窗推荐。 可选值：true，false。默认不过滤该条件 */
func (this *ItemsOnsaleGetRequest) SetHasShowcase(value string) {
	this.Set("has_showcase", value)
}

/* 是否挂接了达尔文标准产品体系。 */
func (this *ItemsOnsaleGetRequest) SetIsCspu(value string) {
	this.Set("is_cspu", value)
}

/* 商品是否在外部网店显示 */
func (this *ItemsOnsaleGetRequest) SetIsEx(value string) {
	this.Set("is_ex", value)
}

/* 商品是否在淘宝显示 */
func (this *ItemsOnsaleGetRequest) SetIsTaobao(value string) {
	this.Set("is_taobao", value)
}

/* 排序方式。格式为column:asc/desc ，column可选值:list_time(上架时间),delist_time(下架时间),num(商品数量)，modified(最近修改时间)，sold_quantity（商品销量）,;默认上架时间降序(即最新上架排在前面)。如按照上架时间降序排序方式为list_time:desc */
func (this *ItemsOnsaleGetRequest) SetOrderBy(value string) {
	this.Set("order_by", value)
}

/* 页码。取值范围:大于零的整数。默认值为1,即默认返回第一页数据。用此接口获取数据时，当翻页获取的条数（page_no*page_size）超过10万,为了保护后台搜索引擎，接口将报错。所以请大家尽可能的细化自己的搜索条件，例如根据修改时间分段获取商品 */
func (this *ItemsOnsaleGetRequest) SetPageNo(value string) {
	this.Set("page_no", value)
}

/* 每页条数。取值范围:大于零的整数;最大值：200；默认值：40。用此接口获取数据时，当翻页获取的条数（page_no*page_size）超过2万,为了保护后台搜索引擎，接口将报错。所以请大家尽可能的细化自己的搜索条件，例如根据修改时间分段获取商品 */
func (this *ItemsOnsaleGetRequest) SetPageSize(value string) {
	this.Set("page_size", value)
}

/* 搜索字段。搜索商品的title。 */
func (this *ItemsOnsaleGetRequest) SetQ(value string) {
	this.Set("q", value)
}

/* 卖家店铺内自定义类目ID。多个之间用“,”分隔。可以根据taobao.sellercats.list.get获得.(<font color="red">注：目前最多支持32个ID号传入</font>) */
func (this *ItemsOnsaleGetRequest) SetSellerCids(value string) {
	this.Set("seller_cids", value)
}

/* 起始的修改时间 */
func (this *ItemsOnsaleGetRequest) SetStartModified(value string) {
	this.Set("start_modified", value)
}

type ItemsOnsaleGetResponse struct {
	*ErrorResponse               `json:"error_response,omitempty"`
	ItemsOnsaleGetResponseResult `json:"items_onsale_get_response"`
}
type ItemsOnsaleGetResponseResult struct {
	/* 搜索到的商品列表，具体字段根据设定的fields决定，不包括desc字段 */
	Items struct {
		Item []*Item `json:"item"`
	} `json:"Items"`
	/* 搜索到符合条件的结果总数 */
	TotalResults int64 `json:"total_results"`
}

/*taobao.product.add*/
type ProductAddRequest struct {
	values url.Values
}

func (this *ProductAddRequest) GetApiMethodName() string {
	return "taobao.product.add"
}
func (this *ProductAddRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *ProductAddRequest) GetValues() url.Values {
	return this.values
}

/* 非关键属性结构:pid:vid;pid:vid.<br>
非关键属性<font color=red>不包含</font>关键属性、销售属性、用户自定义属性、商品属性;
<br>调用taobao.itemprops.get获取pid,调用taobao.itempropvalues.get获取vid.<br><font color=red>注:支持最大长度为512字节</font><br /> 支持最大长度为：512<br /> 支持的最大列表长度为：512 */
func (this *ProductAddRequest) SetBinds(value string) {
	this.Set("binds", value)
}

/* 商品类目ID.调用taobao.itemcats.get获取;注意:必须是叶子类目 id. */
func (this *ProductAddRequest) SetCid(value string) {
	this.Set("cid", value)
}

/* 用户自定义属性,结构：pid1:value1;pid2:value2，如果有型号，系列等子属性用: 隔开 例如：“20000:优衣库:型号:001;632501:1234”，表示“品牌:优衣库:型号:001;货号:1234”
<br><font color=red>注：包含所有自定义属性的传入</font> */
func (this *ProductAddRequest) SetCustomerProps(value string) {
	this.Set("customer_props", value)
}

/* 产品描述.最大不超过25000个字符 */
func (this *ProductAddRequest) SetDesc(value string) {
	this.Set("desc", value)
}

/* 存放产品扩展信息，由List(ProductExtraInfo)转化成jsonArray存入.<br /> 支持最大长度为：25000<br /> 支持的最大列表长度为：25000 */
func (this *ProductAddRequest) SetExtraInfo(value string) {
	this.Set("extra_info", value)
}

/* 产品主图片.最大1M,目前仅支持GIF,JPG.<br /> 支持的文件类型为：gif,jpg,png,jpeg<br /> 支持的最大列表长度为：1048576 */
func (this *ProductAddRequest) SetImage(value string) {
	this.Set("image", value)
}

/* 是否发布套装产品，和suite_items_str配合使用，is_pub_suite=true走套装SPU发布逻辑，达尔文体系下不需要再走tmall.product.spec.add发布产品规格 */
func (this *ProductAddRequest) SetIsPubSuite(value string) {
	this.Set("is_pub_suite", value)
}

/* 是不是主图 */
func (this *ProductAddRequest) SetMajor(value string) {
	this.Set("major", value)
}

/* 市场ID，1为新增C2C市场的产品信息， 2为新增B2C市场的产品信息。
不填写此值则C用户新增B2C市场的产品信息，B用户新增B2C市场的产品信息。 */
func (this *ProductAddRequest) SetMarketId(value string) {
	this.Set("market_id", value)
}

/* 上市时间。目前只支持鞋城类目传入此参数 */
func (this *ProductAddRequest) SetMarketTime(value string) {
	this.Set("market_time", value)
}

/* 产品名称,最大30个字符. */
func (this *ProductAddRequest) SetName(value string) {
	this.Set("name", value)
}

/* 外部产品ID */
func (this *ProductAddRequest) SetOuterId(value string) {
	this.Set("outer_id", value)
}

/* 包装清单。注意，在管控类目下，包装清单不能为空，同时保证清单的格式为：
名称:数字;名称:数字;
其中，名称不能违禁、不能超过60字符，数字不能超过999 */
func (this *ProductAddRequest) SetPackingList(value string) {
	this.Set("packing_list", value)
}

/* 产品市场价.精确到2位小数;单位为元.如：200.07 */
func (this *ProductAddRequest) SetPrice(value string) {
	this.Set("price", value)
}

/* 销售属性值别名。格式为pid1:vid1:alias1;pid1:vid2:alia2。只有少数销售属性值支持传入别名，比如颜色和尺寸 */
func (this *ProductAddRequest) SetPropertyAlias(value string) {
	this.Set("property_alias", value)
}

/* 关键属性 结构:pid:vid;pid:vid.调用taobao.itemprops.get获取pid,调用taobao.itempropvalues.get获取vid;如果碰到用户自定义属性,请用customer_props. */
func (this *ProductAddRequest) SetProps(value string) {
	this.Set("props", value)
}

/* 销售属性结构:pid:vid;pid:vid.调用taobao.itemprops.get获取is_sale_prop＝true的pid,调用taobao.itempropvalues.get获取vid. */
func (this *ProductAddRequest) SetSaleProps(value string) {
	this.Set("sale_props", value)
}

/* 商品卖点描述，长度限制为20个汉字 */
func (this *ProductAddRequest) SetSellPt(value string) {
	this.Set("sell_pt", value)
}

/* 发布套装产品时，套装关联的产品规格+数量的字符串，格式：specsId:number。 */
func (this *ProductAddRequest) SetSuiteItemsStr(value string) {
	this.Set("suite_items_str", value)
}

/* 在天猫，无关键属性发布产品，必须指定模板ID,模板ID通过tmall.product.template.get获取 */
func (this *ProductAddRequest) SetTemplateId(value string) {
	this.Set("template_id", value)
}

type ProductAddResponse struct {
	*ErrorResponse           `json:"error_response,omitempty"`
	ProductAddResponseResult `json:"product_add_response"`
}
type ProductAddResponseResult struct {
	/* 产品结构 */
	Product *Product `json:"product"`
}

/*taobao.product.get*/
type ProductGetRequest struct {
	values url.Values
}

func (this *ProductGetRequest) GetApiMethodName() string {
	return "taobao.product.get"
}
func (this *ProductGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *ProductGetRequest) GetValues() url.Values {
	return this.values
}

/* 商品类目id.调用taobao.itemcats.get获取;必须是叶子类目id,如果没有传product_id,那么cid和props必须要传. */
func (this *ProductGetRequest) SetCid(value string) {
	this.Set("cid", value)
}

/* 用户自定义关键属性,结构：pid1:value1;pid2:value2，如果有型号，系列等子属性用: 隔开 例如：“20000:优衣库:型号:001;632501:1234”，表示“品牌:优衣库:型号:001;货号:1234” */
func (this *ProductGetRequest) SetCustomerProps(value string) {
	this.Set("customer_props", value)
}

/* 需返回的字段列表.可选值:Product数据结构中的所有字段;多个字段之间用","分隔. */
func (this *ProductGetRequest) SetFields(value string) {
	this.Set("fields", value)
}

/* 市场ID，1为取C2C市场的产品信息， 2为取B2C市场的产品信息。
不填写此值则默认取C2C的产品信息。 */
func (this *ProductGetRequest) SetMarketId(value string) {
	this.Set("market_id", value)
}

/* Product的id.两种方式来查看一个产品:1.传入product_id来查询 2.传入cid和props来查询 */
func (this *ProductGetRequest) SetProductId(value string) {
	this.Set("product_id", value)
}

/* 比如:诺基亚N73这个产品的关键属性列表就是:品牌:诺基亚;型号:N73,对应的PV值就是10005:10027;10006:29729. */
func (this *ProductGetRequest) SetProps(value string) {
	this.Set("props", value)
}

type ProductGetResponse struct {
	*ErrorResponse           `json:"error_response,omitempty"`
	ProductGetResponseResult `json:"product_get_response"`
}
type ProductGetResponseResult struct {
	/* 返回具体信息为入参fields请求的字段信息 */
	Product *Product `json:"product"`
}

/*taobao.product.img.delete*/
type ProductImgDeleteRequest struct {
	values url.Values
}

func (this *ProductImgDeleteRequest) GetApiMethodName() string {
	return "taobao.product.img.delete"
}
func (this *ProductImgDeleteRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *ProductImgDeleteRequest) GetValues() url.Values {
	return this.values
}

/* 非主图ID */
func (this *ProductImgDeleteRequest) SetId(value string) {
	this.Set("id", value)
}

/* 产品ID.Product的id,通过taobao.product.add接口新增产品的时候会返回id. */
func (this *ProductImgDeleteRequest) SetProductId(value string) {
	this.Set("product_id", value)
}

type ProductImgDeleteResponse struct {
	*ErrorResponse                 `json:"error_response,omitempty"`
	ProductImgDeleteResponseResult `json:"product_img_delete_response"`
}
type ProductImgDeleteResponseResult struct {
	/* 返回productimg中的：id,product_id */
	ProductImg *ProductImg `json:"product_img"`
}

/*taobao.product.img.upload*/
type ProductImgUploadRequest struct {
	values url.Values
}

func (this *ProductImgUploadRequest) GetApiMethodName() string {
	return "taobao.product.img.upload"
}
func (this *ProductImgUploadRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *ProductImgUploadRequest) GetValues() url.Values {
	return this.values
}

/* 产品图片ID.修改图片时需要传入 */
func (this *ProductImgUploadRequest) SetId(value string) {
	this.Set("id", value)
}

/* 图片内容.图片最大为500K,只支持JPG,GIF格式.<br /> 支持的文件类型为：gif,jpg,png,jpeg<br /> 支持的最大列表长度为：1048576 */
func (this *ProductImgUploadRequest) SetImage(value string) {
	this.Set("image", value)
}

/* 是否将该图片设为主图.可选值:true,false;默认值:false. */
func (this *ProductImgUploadRequest) SetIsMajor(value string) {
	this.Set("is_major", value)
}

/* 图片序号 */
func (this *ProductImgUploadRequest) SetPosition(value string) {
	this.Set("position", value)
}

/* 产品ID.Product的id */
func (this *ProductImgUploadRequest) SetProductId(value string) {
	this.Set("product_id", value)
}

type ProductImgUploadResponse struct {
	*ErrorResponse                 `json:"error_response,omitempty"`
	ProductImgUploadResponseResult `json:"product_img_upload_response"`
}
type ProductImgUploadResponseResult struct {
	/* 返回产品图片结构中的：url,id,created,modified */
	ProductImg *ProductImg `json:"product_img"`
}

/*taobao.product.propimg.delete*/
type ProductPropimgDeleteRequest struct {
	values url.Values
}

func (this *ProductPropimgDeleteRequest) GetApiMethodName() string {
	return "taobao.product.propimg.delete"
}
func (this *ProductPropimgDeleteRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *ProductPropimgDeleteRequest) GetValues() url.Values {
	return this.values
}

/* 属性图片ID */
func (this *ProductPropimgDeleteRequest) SetId(value string) {
	this.Set("id", value)
}

/* 产品ID.Product的id. */
func (this *ProductPropimgDeleteRequest) SetProductId(value string) {
	this.Set("product_id", value)
}

type ProductPropimgDeleteResponse struct {
	*ErrorResponse                     `json:"error_response,omitempty"`
	ProductPropimgDeleteResponseResult `json:"product_propimg_delete_response"`
}
type ProductPropimgDeleteResponseResult struct {
	/* 返回product_prop_img数据结构中的：product_id,id */
	ProductPropImg *ProductPropImg `json:"product_prop_img"`
}

/*taobao.product.propimg.upload*/
type ProductPropimgUploadRequest struct {
	values url.Values
}

func (this *ProductPropimgUploadRequest) GetApiMethodName() string {
	return "taobao.product.propimg.upload"
}
func (this *ProductPropimgUploadRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *ProductPropimgUploadRequest) GetValues() url.Values {
	return this.values
}

/* 产品属性图片ID */
func (this *ProductPropimgUploadRequest) SetId(value string) {
	this.Set("id", value)
}

/* 图片内容.图片最大为2M,只支持JPG,GIF.<br /> 支持的文件类型为：gif,jpg,png,jpeg<br /> 支持的最大列表长度为：1048576 */
func (this *ProductPropimgUploadRequest) SetImage(value string) {
	this.Set("image", value)
}

/* 图片序号 */
func (this *ProductPropimgUploadRequest) SetPosition(value string) {
	this.Set("position", value)
}

/* 产品ID.Product的id */
func (this *ProductPropimgUploadRequest) SetProductId(value string) {
	this.Set("product_id", value)
}

/* 属性串.目前仅支持颜色属性.调用taobao.itemprops.get获取类目属性,取得颜色属性pid,再用taobao.itempropvalues.get取得vid;格式:pid:vid,只能传入一个颜色pid:vid串; */
func (this *ProductPropimgUploadRequest) SetProps(value string) {
	this.Set("props", value)
}

type ProductPropimgUploadResponse struct {
	*ErrorResponse                     `json:"error_response,omitempty"`
	ProductPropimgUploadResponseResult `json:"product_propimg_upload_response"`
}
type ProductPropimgUploadResponseResult struct {
	/* 支持返回产品属性图片中的：url,id,created,modified */
	ProductPropImg *ProductPropImg `json:"product_prop_img"`
}

/*taobao.product.update*/
type ProductUpdateRequest struct {
	values url.Values
}

func (this *ProductUpdateRequest) GetApiMethodName() string {
	return "taobao.product.update"
}
func (this *ProductUpdateRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *ProductUpdateRequest) GetValues() url.Values {
	return this.values
}

/* 非关键属性.调用taobao.itemprops.get获取pid,调用taobao.itempropvalues.get获取vid;格式:pid:vid;pid:vid */
func (this *ProductUpdateRequest) SetBinds(value string) {
	this.Set("binds", value)
}

/* 产品描述.最大不超过25000个字符 */
func (this *ProductUpdateRequest) SetDesc(value string) {
	this.Set("desc", value)
}

/* 存放产品扩展信息，由List(ProductExtraInfo)转化成jsonArray存入.<br /> 支持最大长度为：25000<br /> 支持的最大列表长度为：25000 */
func (this *ProductUpdateRequest) SetExtraInfo(value string) {
	this.Set("extra_info", value)
}

/* 产品主图.最大500K,目前仅支持GIF,JPG<br /> 支持的文件类型为：gif,jpg,png,jpeg<br /> 支持的最大列表长度为：1048576 */
func (this *ProductUpdateRequest) SetImage(value string) {
	this.Set("image", value)
}

/* 是否是主图 */
func (this *ProductUpdateRequest) SetMajor(value string) {
	this.Set("major", value)
}

/* 市场ID，1为更新C2C市场的产品信息， 2为更新B2C市场的产品信息。
不填写此值则C用户更新B2C市场的产品信息，B用户更新B2C市场的产品信息。 */
func (this *ProductUpdateRequest) SetMarketId(value string) {
	this.Set("market_id", value)
}

/* 产品名称.最大不超过30个字符 */
func (this *ProductUpdateRequest) SetName(value string) {
	this.Set("name", value)
}

/* 自定义非关键属性 */
func (this *ProductUpdateRequest) SetNativeUnkeyprops(value string) {
	this.Set("native_unkeyprops", value)
}

/* 外部产品ID */
func (this *ProductUpdateRequest) SetOuterId(value string) {
	this.Set("outer_id", value)
}

/* 保证清单。 */
func (this *ProductUpdateRequest) SetPackingList(value string) {
	this.Set("packing_list", value)
}

/* 产品市场价.精确到2位小数;单位为元.如:200.07 */
func (this *ProductUpdateRequest) SetPrice(value string) {
	this.Set("price", value)
}

/* 产品ID */
func (this *ProductUpdateRequest) SetProductId(value string) {
	this.Set("product_id", value)
}

/* 销售属性.调用taobao.itemprops.get获取pid,调用taobao.itempropvalues.get获取vid;格式:pid:vid;pid:vid */
func (this *ProductUpdateRequest) SetSaleProps(value string) {
	this.Set("sale_props", value)
}

/* 产品卖点描述，最长40个字节 */
func (this *ProductUpdateRequest) SetSellPt(value string) {
	this.Set("sell_pt", value)
}

type ProductUpdateResponse struct {
	*ErrorResponse              `json:"error_response,omitempty"`
	ProductUpdateResponseResult `json:"product_update_response"`
}
type ProductUpdateResponseResult struct {
	/* 返回product数据结构中的：product_id,modified */
	Product *Product `json:"product"`
}

/*taobao.products.get*/
type ProductsGetRequest struct {
	values url.Values
}

func (this *ProductsGetRequest) GetApiMethodName() string {
	return "taobao.products.get"
}
func (this *ProductsGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *ProductsGetRequest) GetValues() url.Values {
	return this.values
}

/* 需返回的字段列表.可选值:Product数据结构中的所有字段;多个字段之间用","分隔 */
func (this *ProductsGetRequest) SetFields(value string) {
	this.Set("fields", value)
}

/* 用户昵称 */
func (this *ProductsGetRequest) SetNick(value string) {
	this.Set("nick", value)
}

/* 页码.传入值为1代表第一页,传入值为2代表第二页,依此类推.默认返回的数据是从第一页开始. */
func (this *ProductsGetRequest) SetPageNo(value string) {
	this.Set("page_no", value)
}

/* 每页条数.每页返回最多返回100条,默认值为40 */
func (this *ProductsGetRequest) SetPageSize(value string) {
	this.Set("page_size", value)
}

type ProductsGetResponse struct {
	*ErrorResponse            `json:"error_response,omitempty"`
	ProductsGetResponseResult `json:"products_get_response"`
}
type ProductsGetResponseResult struct {
	/* 返回具体信息为入参fields请求的字段信息 */
	Products struct {
		Product []*Product `json:"product"`
	} `json:"Products"`
}

/*taobao.products.search*/
type ProductsSearchRequest struct {
	values url.Values
}

func (this *ProductsSearchRequest) GetApiMethodName() string {
	return "taobao.products.search"
}
func (this *ProductsSearchRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *ProductsSearchRequest) GetValues() url.Values {
	return this.values
}

/* 按条码搜索产品信息,多个逗号隔开，不支持条码为全零的方式 */
func (this *ProductsSearchRequest) SetBarcodeStr(value string) {
	this.Set("barcode_str", value)
}

/* 商品类目ID.调用taobao.itemcats.get获取. */
func (this *ProductsSearchRequest) SetCid(value string) {
	this.Set("cid", value)
}

/* 用户自定义关键属性,结构：pid1:value1;pid2:value2，如果有型号，系列等子属性用: 隔开 例如：“20000:优衣库:型号:001;632501:1234”，表示“品牌:优衣库:型号:001;货号:1234” */
func (this *ProductsSearchRequest) SetCustomerProps(value string) {
	this.Set("customer_props", value)
}

/* 需返回的字段列表.可选值:Product数据结构中的以下字段:product_id,name,pic_url,cid,props,price,tsc;多个字段之间用","分隔.新增字段status(product的当前状态) */
func (this *ProductsSearchRequest) SetFields(value string) {
	this.Set("fields", value)
}

/* 市场ID，1为取C2C市场的产品信息， 2为取B2C市场的产品信息。
不填写此值则默认取C2C的产品信息。 */
func (this *ProductsSearchRequest) SetMarketId(value string) {
	this.Set("market_id", value)
}

/* 页码.传入值为1代表第一页,传入值为2代表第二页,依此类推.默认返回的数据是从第一页开始. */
func (this *ProductsSearchRequest) SetPageNo(value string) {
	this.Set("page_no", value)
}

/* 每页条数.每页返回最多返回100条,默认值为40. */
func (this *ProductsSearchRequest) SetPageSize(value string) {
	this.Set("page_size", value)
}

/* 属性,属性值的组合.格式:pid:vid;pid:vid;调用taobao.itemprops.get获取类目属性pid
,再用taobao.itempropvalues.get取得vid. */
func (this *ProductsSearchRequest) SetProps(value string) {
	this.Set("props", value)
}

/* 搜索的关键词是用来搜索产品的title.　注:q,cid和props至少传入一个 */
func (this *ProductsSearchRequest) SetQ(value string) {
	this.Set("q", value)
}

/* 想要获取的产品的状态列表，支持多个状态并列获取，多个状态之间用","分隔，最多同时指定5种状态。例如，只获取小二确认的spu传入"3",只要商家确认的传入"0"，既要小二确认又要商家确认的传入"0,3"。目前只支持者两种类型的状态搜索，输入其他状态无效。<br /> 支持最大长度为：20<br /> 支持的最大列表长度为：20 */
func (this *ProductsSearchRequest) SetStatus(value string) {
	this.Set("status", value)
}

/* 按关联产品规格specs搜索套装产品 */
func (this *ProductsSearchRequest) SetSuiteItemsStr(value string) {
	this.Set("suite_items_str", value)
}

/* 传入值为：3表示3C表示3C垂直市场产品，4表示鞋城垂直市场产品，8表示网游垂直市场产品。一次只能指定一种垂直市场类型<br /> 支持最小值为：0 */
func (this *ProductsSearchRequest) SetVerticalMarket(value string) {
	this.Set("vertical_market", value)
}

type ProductsSearchResponse struct {
	*ErrorResponse               `json:"error_response,omitempty"`
	ProductsSearchResponseResult `json:"products_search_response"`
}
type ProductsSearchResponseResult struct {
	/* 返回具体信息为入参fields请求的字段信息 */
	Products struct {
		Product []*Product `json:"product"`
	} `json:"Products"`
	/* 结果总数 */
	TotalResults int64 `json:"total_results"`
}

/*taobao.skus.custom.get*/
type SkusCustomGetRequest struct {
	values url.Values
}

func (this *SkusCustomGetRequest) GetApiMethodName() string {
	return "taobao.skus.custom.get"
}
func (this *SkusCustomGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *SkusCustomGetRequest) GetValues() url.Values {
	return this.values
}

/* 需返回的字段列表。可选值：Sku结构体中的所有字段；字段之间用“,”隔开 */
func (this *SkusCustomGetRequest) SetFields(value string) {
	this.Set("fields", value)
}

/* Sku的外部商家ID */
func (this *SkusCustomGetRequest) SetOuterId(value string) {
	this.Set("outer_id", value)
}

type SkusCustomGetResponse struct {
	*ErrorResponse              `json:"error_response,omitempty"`
	SkusCustomGetResponseResult `json:"skus_custom_get_response"`
}
type SkusCustomGetResponseResult struct {
	/* Sku对象，具体字段以fields决定 */
	Skus struct {
		Sku []*Sku `json:"sku"`
	} `json:"Skus"`
}

/*taobao.skus.quantity.update*/
type SkusQuantityUpdateRequest struct {
	values url.Values
}

func (this *SkusQuantityUpdateRequest) GetApiMethodName() string {
	return "taobao.skus.quantity.update"
}
func (this *SkusQuantityUpdateRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *SkusQuantityUpdateRequest) GetValues() url.Values {
	return this.values
}

/* 商品数字ID，必填参数 */
func (this *SkusQuantityUpdateRequest) SetNumIid(value string) {
	this.Set("num_iid", value)
}

/* 特殊可选，skuIdQuantities为空的时候用该字段通过outerId来指定sku和其库存修改值。格式为outerId:库存修改值;outerId:库存修改值。当skuIdQuantities不为空的时候该字段失效。当一个outerId对应多个sku时，所有匹配到的sku都会被修改库存。最多支持20个SKU同时修改。 */
func (this *SkusQuantityUpdateRequest) SetOuteridQuantities(value string) {
	this.Set("outerid_quantities", value)
}

/* sku库存批量修改入参，用于指定一批sku和每个sku的库存修改值，特殊可填。格式为skuId:库存修改值;skuId:库存修改值。最多支持20个SKU同时修改。 */
func (this *SkusQuantityUpdateRequest) SetSkuidQuantities(value string) {
	this.Set("skuid_quantities", value)
}

/* 库存更新方式，可选。1为全量更新，2为增量更新。如果不填，默认为全量更新。当选择全量更新时，如果库存更新值传入的是负数，会出错并返回错误码；当选择增量更新时，如果库存更新值为负数且绝对值大于当前库存，则sku库存会设置为0. */
func (this *SkusQuantityUpdateRequest) SetType(value string) {
	this.Set("type", value)
}

type SkusQuantityUpdateResponse struct {
	*ErrorResponse                   `json:"error_response,omitempty"`
	SkusQuantityUpdateResponseResult `json:"skus_quantity_update_response"`
}
type SkusQuantityUpdateResponseResult struct {
	/* iid、numIid、num和modified，skus中每个sku的skuId、quantity和modified */
	Item *Item `json:"item"`
}

/*taobao.ticket.item.add*/
type TicketItemAddRequest struct {
	values url.Values
}

func (this *TicketItemAddRequest) GetApiMethodName() string {
	return "taobao.ticket.item.add"
}
func (this *TicketItemAddRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *TicketItemAddRequest) GetValues() url.Values {
	return this.values
}

/* 门票商品返点比例（只对B卖家开放，单位为%） */
func (this *TicketItemAddRequest) SetAuctionPoint(value string) {
	this.Set("auction_point", value)
}

/* 门票商品状态（onsale：上架，instock：仓库）<br /> 支持最大长度为：10<br /> 支持的最大列表长度为：10 */
func (this *TicketItemAddRequest) SetAuctionStatus(value string) {
	this.Set("auction_status", value)
}

/* 门票商品所在地-城市<br /> 支持最大长度为：20<br /> 支持的最大列表长度为：20 */
func (this *TicketItemAddRequest) SetCity(value string) {
	this.Set("city", value)
}

/* 门票宝贝描述<br /> 支持最大长度为：25000<br /> 支持的最大列表长度为：25000 */
func (this *TicketItemAddRequest) SetDescription(value string) {
	this.Set("description", value)
}

/* 门票商品电子凭证是否关联本地商户-在门票商品为电子凭证时必选 */
func (this *TicketItemAddRequest) SetEtcAssociationStatus(value string) {
	this.Set("etc.association_status", value)
}

/* 门票商品电子凭证的自动退款比例-在门票商品为电子凭证时必选<br /> 支持最大值为：100<br /> 支持最小值为：0 */
func (this *TicketItemAddRequest) SetEtcAutoRefund(value string) {
	this.Set("etc.auto_refund", value)
}

/* 门票商品电子凭证的码商-在门票商品为电子凭证时必选 */
func (this *TicketItemAddRequest) SetEtcMerchantId(value string) {
	this.Set("etc.merchant_id", value)
}

/* 门票商品电子凭证的码商名-在门票商品为电子凭证时必选 */
func (this *TicketItemAddRequest) SetEtcMerchantNick(value string) {
	this.Set("etc.merchant_nick", value)
}

/* 门票商品电子凭证网店ID-在门票商品为电子凭证时必选 */
func (this *TicketItemAddRequest) SetEtcNetworkId(value string) {
	this.Set("etc.network_id", value)
}

/* 门票商品电子凭证的过期退款比例-在门票商品为电子凭证时必选<br /> 支持最大值为：100<br /> 支持最小值为：0 */
func (this *TicketItemAddRequest) SetEtcOverduePay(value string) {
	this.Set("etc.overdue_pay", value)
}

/* 门票商品电子凭证是否核销打款-在门票商品为电子凭证时必选 */
func (this *TicketItemAddRequest) SetEtcVerificationPay(value string) {
	this.Set("etc.verification_pay", value)
}

/* 门票商品是否有发票（有发票：true，没有发票：false）.
对于B卖家来说只能选择有发票。 */
func (this *TicketItemAddRequest) SetHaveInvoice(value string) {
	this.Set("have_invoice", value)
}

/* 商品主图-该值来自图片空间接口返回的图片链接（不包含域名和前缀） */
func (this *TicketItemAddRequest) SetImage1(value string) {
	this.Set("image_1", value)
}

/* 第一张多图-该值来自图片空间接口返回的图片链接（不包含域名和前缀） */
func (this *TicketItemAddRequest) SetImage2(value string) {
	this.Set("image_2", value)
}

/* 第二张多图-该值来自图片空间接口返回的图片链接（不包含域名和前缀） */
func (this *TicketItemAddRequest) SetImage3(value string) {
	this.Set("image_3", value)
}

/* 第三张多图-该值来自图片空间接口返回的图片链接（不包含域名和前缀） */
func (this *TicketItemAddRequest) SetImage4(value string) {
	this.Set("image_4", value)
}

/* 第四张多图-该值来自图片空间接口返回的图片链接（不包含域名和前缀） */
func (this *TicketItemAddRequest) SetImage5(value string) {
	this.Set("image_5", value)
}

/* 门票商品的上架时间（精确到分，格式为：yyyy-MM-dd HH:mm）<br /> 支持最大长度为：20<br /> 支持的最大列表长度为：20 */
func (this *TicketItemAddRequest) SetListTime(value string) {
	this.Set("list_time", value)
}

/* 门票商品的物流运费模板-在产品规格使用到物流时必选 */
func (this *TicketItemAddRequest) SetPostageId(value string) {
	this.Set("postage_id", value)
}

/* 门票对应产品标识<br>
门票商品只能使用达尔文体系下的产品。该参数可通过taobao.products.search接口获取，注意返回产品的cspu_feature属性，该属性表明是否为达尔文体系下产品。 */
func (this *TicketItemAddRequest) SetProductId(value string) {
	this.Set("product_id", value)
}

/* 门票商品是否橱窗推荐（橱窗推荐；true，不推荐则可不用设置该值） */
func (this *TicketItemAddRequest) SetPromotedStatus(value string) {
	this.Set("promoted_status", value)
}

/* 门票商品所在地-省份<br /> 支持最大长度为：20<br /> 支持的最大列表长度为：20 */
func (this *TicketItemAddRequest) SetProv(value string) {
	this.Set("prov", value)
}

/* 门票宝贝所属的店铺分类列表-店铺分类标识请使用店铺相关接口获取获取，多个店铺分类标识之间通过逗号进行分隔，最多包含10个分类标识 */
func (this *TicketItemAddRequest) SetShopCats(value string) {
	this.Set("shop_cats", value)
}

/* 门票商品产品规格信息，包括门票的库存价格等信息（产品规格可通过tmall.product.specs.get接口获取）。<br>请通过门票商品使用的产品所对应的产品规格信息进行设置，注意由于产品规格信息比较复杂，所以直接使用json作为传输，请按照约定传入正确的json格式）<br>
格式定义：<br>
<pre>
{
  "SKU":
	{
		"20890017-121840019_20394-121290067":{// 产品下的某个产品规格的属性属性值信息（pid-vid_pid-vid）
			"effDates":{
				"1":{
					"type":"0",// 有效期类型，0-非指定日票，1-指定日票，2-年卡
					"effDate" :{ // 有效期，如下几块，type已经去掉
						"startDate":"2013-01-01",// 有效期时间段开始时间,null代表未设置
						"endDate":"2013-12-31",// 有效期时间段结束时间,null代表未设置
						"weeks":["1","4"],// 有效期周,1~7代表周一到周日，null代表未设置
						"startHour":"12",// 有效期开始时间小时,null代表未设置
						"startMinute":"59",// 有效期开始时间分钟,null代表未设置
						"endHour":"18"// 有效期结束时间小时,null代表未设置
						"endMinute":"30"// 有效期结束时间分钟,null代表未设置
						"effDays":"1",// xx后n天内有效的天数，购买后、出票后、开卡后n天内有效,null代表未设置
					},
					"timeLimit":{// 入园时间限制数据
						"limit":{
							"type":"0",// 入园时间限制类型,0-不限，1-提前n天的n点n分，2-提前n小时n分钟
							"aheadDays":"3",// 入园时间限制类型为提前n天的n点n分时才使用，入园时间提前n天的天数
							"aheadAtHour":"12",// 入园时间限制类型为提前n天的n点n分时才使用，入园时间在n点n分之前的小时
							"aheadAtMinute":"30",// 入园时间限制类型为提前n天的n点n分时才使用，入园时间在n点n分之前的分钟
							"aheadHours":"48",// 入园时间限制类型为提前n小时n分时才使用，入园时间提前n小时的小时
							"aheadMinutes":"30",// 入园时间限制类型为提前n小时n分时才使用，入园时间提前n分钟的分钟
						}
						"autoActivate":{ // 有效期类型为年卡时才使用
							"type":"0",// 是否有最晚自动开卡时间,0-无，1-有最晚自动开卡
							"time":"30" // 天数,无最晚自动开发时间时为null
						}
					}

					"price" : 10000,
					"inventory" : 100,
					"outerId" : "商家编码"
				},
				"222":{}
			},
			"save":true, // 必须，代表是保存的有效数据

				"returnRule":{
					"type":"0",      // 0,1,2
					"value":""      // type 0:""，1:""，2:"卖家输入的文本"
				},
				"changeRule":{
					"type":"1",      // 0,1,2
					"value":""      // type 0:""，1:""，2:"卖家输入的文本"
				},
				"guide": "卖家输入的文本"
		}
	}
}
</pre><br> */
func (this *TicketItemAddRequest) SetSkus(value string) {
	this.Set("skus", value)
}

/* 门票商品库存技术方式（拍下减库存：true，付款减库存：false） */
func (this *TicketItemAddRequest) SetSubStockAtBuy(value string) {
	this.Set("sub_stock_at_buy", value)
}

/* 门票商品标题<br /> 支持最大长度为：60<br /> 支持的最大列表长度为：60 */
func (this *TicketItemAddRequest) SetTitle(value string) {
	this.Set("title", value)
}

/* 门票商品视频-视频标识由多媒体中相关接口获取 */
func (this *TicketItemAddRequest) SetVideoId(value string) {
	this.Set("video_id", value)
}

/* 门票商品是否参与店铺会员打折 */
func (this *TicketItemAddRequest) SetVipPromoted(value string) {
	this.Set("vip_promoted", value)
}

type TicketItemAddResponse struct {
	*ErrorResponse              `json:"error_response,omitempty"`
	TicketItemAddResponseResult `json:"ticket_item_add_response"`
}
type TicketItemAddResponseResult struct {
	/* 门票商品操作结果，具体请参见TicketItemProcessResult数据结构 */
	TicketItemProcessResult *TicketItemProcessResult `json:"ticket_item_process_result"`
}

/*taobao.ticket.item.get*/
type TicketItemGetRequest struct {
	values url.Values
}

func (this *TicketItemGetRequest) GetApiMethodName() string {
	return "taobao.ticket.item.get"
}
func (this *TicketItemGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *TicketItemGetRequest) GetValues() url.Values {
	return this.values
}

/* 需要返回的门票商品（TicketItem）对象字段，如title,price,skus等。<br>可选值：TicketItem商品结构体中所有字段均可返回；多个字段用“,”分隔。 */
func (this *TicketItemGetRequest) SetFields(value string) {
	this.Set("fields", value)
}

/* 新门票类目商品的标识（非日历价格库存商品） */
func (this *TicketItemGetRequest) SetItemId(value string) {
	this.Set("item_id", value)
}

type TicketItemGetResponse struct {
	*ErrorResponse              `json:"error_response,omitempty"`
	TicketItemGetResponseResult `json:"ticket_item_get_response"`
}
type TicketItemGetResponseResult struct {
	/* 参见TicketItem数据结构文档 */
	TicketItem *TicketItem `json:"ticket_item"`
}

/*taobao.ticket.item.update*/
type TicketItemUpdateRequest struct {
	values url.Values
}

func (this *TicketItemUpdateRequest) GetApiMethodName() string {
	return "taobao.ticket.item.update"
}
func (this *TicketItemUpdateRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *TicketItemUpdateRequest) GetValues() url.Values {
	return this.values
}

/* 门票商品返点比例（只对B卖家开放，单位为%） */
func (this *TicketItemUpdateRequest) SetAuctionPoint(value string) {
	this.Set("auction_point", value)
}

/* 门票商品状态（onsale：上架，instock：仓库）<br /> 支持最大长度为：10<br /> 支持的最大列表长度为：10 */
func (this *TicketItemUpdateRequest) SetAuctionStatus(value string) {
	this.Set("auction_status", value)
}

/* 门票商品所在地-城市<br /> 支持最大长度为：20<br /> 支持的最大列表长度为：20 */
func (this *TicketItemUpdateRequest) SetCity(value string) {
	this.Set("city", value)
}

/* 门票宝贝描述<br /> 支持最大长度为：25000<br /> 支持的最大列表长度为：25000 */
func (this *TicketItemUpdateRequest) SetDescription(value string) {
	this.Set("description", value)
}

/* 门票商品电子凭证是否关联本地商户-在门票商品为电子凭证时必选 */
func (this *TicketItemUpdateRequest) SetEtcAssociationStatus(value string) {
	this.Set("etc.association_status", value)
}

/* 门票商品电子凭证的自动退款比例-在门票商品为电子凭证时必选<br /> 支持最大值为：100<br /> 支持最小值为：0 */
func (this *TicketItemUpdateRequest) SetEtcAutoRefund(value string) {
	this.Set("etc.auto_refund", value)
}

/* 门票商品电子凭证的码商-在门票商品为电子凭证时必选 */
func (this *TicketItemUpdateRequest) SetEtcMerchantId(value string) {
	this.Set("etc.merchant_id", value)
}

/* 门票商品电子凭证的码商名-在门票商品为电子凭证时必选 */
func (this *TicketItemUpdateRequest) SetEtcMerchantNick(value string) {
	this.Set("etc.merchant_nick", value)
}

/* 门票商品电子凭证网店ID-在门票商品为电子凭证时必选 */
func (this *TicketItemUpdateRequest) SetEtcNetworkId(value string) {
	this.Set("etc.network_id", value)
}

/* 门票商品电子凭证的过期退款比例-在门票商品为电子凭证时必选<br /> 支持最大值为：100<br /> 支持最小值为：0 */
func (this *TicketItemUpdateRequest) SetEtcOverduePay(value string) {
	this.Set("etc.overdue_pay", value)
}

/* 门票商品电子凭证是否核销打款-在门票商品为电子凭证时必选 */
func (this *TicketItemUpdateRequest) SetEtcVerificationPay(value string) {
	this.Set("etc.verification_pay", value)
}

/* 门票商品是否有发票（有发票：true，没有发票：false）.
对于B卖家来说只能选择有发票。 */
func (this *TicketItemUpdateRequest) SetHaveInvoice(value string) {
	this.Set("have_invoice", value)
}

/* 商品主图-该值来自图片空间接口返回的图片链接（不包含域名和前缀） */
func (this *TicketItemUpdateRequest) SetImage1(value string) {
	this.Set("image_1", value)
}

/* 第一张多图-该值来自图片空间接口返回的图片链接（不包含域名和前缀） */
func (this *TicketItemUpdateRequest) SetImage2(value string) {
	this.Set("image_2", value)
}

/* 第二张多图-该值来自图片空间接口返回的图片链接（不包含域名和前缀） */
func (this *TicketItemUpdateRequest) SetImage3(value string) {
	this.Set("image_3", value)
}

/* 第三张多图-该值来自图片空间接口返回的图片链接（不包含域名和前缀） */
func (this *TicketItemUpdateRequest) SetImage4(value string) {
	this.Set("image_4", value)
}

/* 第四张多图-该值来自图片空间接口返回的图片链接（不包含域名和前缀） */
func (this *TicketItemUpdateRequest) SetImage5(value string) {
	this.Set("image_5", value)
}

/* 需要更新的门票商品标识（只支持门票二期商品） */
func (this *TicketItemUpdateRequest) SetItemId(value string) {
	this.Set("item_id", value)
}

/* 门票商品的上架时间（精确到分，格式为：yyyy-MM-dd HH:mm）<br /> 支持最大长度为：20<br /> 支持的最大列表长度为：20 */
func (this *TicketItemUpdateRequest) SetListTime(value string) {
	this.Set("list_time", value)
}

/* 门票商品的物流运费模板-在产品规格使用到物流时必选 */
func (this *TicketItemUpdateRequest) SetPostageId(value string) {
	this.Set("postage_id", value)
}

/* 门票商品挂载到的产品<br>
通过taobao.products.search获取产品信息，并且对应产品必须是小二确认的达尔文体系产品。 */
func (this *TicketItemUpdateRequest) SetProductId(value string) {
	this.Set("product_id", value)
}

/* 门票商品是否橱窗推荐（橱窗推荐；true，不推荐则可不用设置该值） */
func (this *TicketItemUpdateRequest) SetPromotedStatus(value string) {
	this.Set("promoted_status", value)
}

/* 门票商品所在地-省份<br /> 支持最大长度为：20<br /> 支持的最大列表长度为：20 */
func (this *TicketItemUpdateRequest) SetProv(value string) {
	this.Set("prov", value)
}

/* 需要删除的字段列表 */
func (this *TicketItemUpdateRequest) SetRemoveFields(value string) {
	this.Set("remove_fields", value)
}

/* 门票宝贝所属的店铺分类列表-店铺分类标识请使用店铺相关接口获取获取，多个店铺分类标识之间通过逗号进行分隔，最多包含10个分类标识 */
func (this *TicketItemUpdateRequest) SetShopCats(value string) {
	this.Set("shop_cats", value)
}

/* 门票商品产品规格信息（门票的产品规格信息可以通过tmall.product.specs.get接口获取）<br>。请通过门票商品使用的产品所对应的产品规格信息进行设置，注意由于产品规格信息比较复杂，所以直接使用json作为传输，请按照约定传入正确的json格式）<br>
格式定义：<br>
<pre>
{
  "SKU":
	{
		"20890017-121840019_20394-121290067":{// 产品下的某个产品规格的属性属性值信息（pid-vid_pid-vid）
			"effDates":{
				"1":{
					"type":"0",// 有效期类型，0-非指定日票，1-指定日票，2-年卡
					"effDate" :{ // 有效期，如下几块，type已经去掉
						"startDate":"2013-01-01",// 有效期时间段开始时间,null代表未设置
						"endDate":"2013-12-31",// 有效期时间段结束时间,null代表未设置
						"weeks":["1","4"],// 有效期周,1~7代表周一到周日，null代表未设置
						"startHour":"12",// 有效期开始时间小时,null代表未设置
						"startMinute":"59",// 有效期开始时间分钟,null代表未设置
						"endHour":"18"// 有效期结束时间小时,null代表未设置
						"endMinute":"30"// 有效期结束时间分钟,null代表未设置
						"effDays":"1",// xx后n天内有效的天数，购买后、出票后、开卡后n天内有效,null代表未设置
					},
					"timeLimit":{// 入园时间限制数据
						"limit":{
							"type":"0",// 入园时间限制类型,0-不限，1-提前n天的n点n分，2-提前n小时n分钟
							"aheadDays":"3",// 入园时间限制类型为提前n天的n点n分时才使用，入园时间提前n天的天数
							"aheadAtHour":"12",// 入园时间限制类型为提前n天的n点n分时才使用，入园时间在n点n分之前的小时
							"aheadAtMinute":"30",// 入园时间限制类型为提前n天的n点n分时才使用，入园时间在n点n分之前的分钟
							"aheadHours":"48",// 入园时间限制类型为提前n小时n分时才使用，入园时间提前n小时的小时
							"aheadMinutes":"30",// 入园时间限制类型为提前n小时n分时才使用，入园时间提前n分钟的分钟
						}
						"autoActivate":{ // 有效期类型为年卡时才使用
							"type":"0",// 是否有最晚自动开卡时间,0-无，1-有最晚自动开卡
							"time":"30" // 天数,无最晚自动开发时间时为null
						}
					}

					"price" : 10000,
					"inventory" : 100,
					"outerId" : "商家编码"
				},
				"222":{}
			},
			"save":true, // 必须，代表是保存的有效数据

				"returnRule":{
					"type":"0",      // 0,1,2
					"value":""      // type 0:""，1:""，2:"卖家输入的文本"
				},
				"changeRule":{
					"type":"1",      // 0,1,2
					"value":""      // type 0:""，1:""，2:"卖家输入的文本"
				},
				"guide": "卖家输入的文本"
		}
	}
}
</pre><br> */
func (this *TicketItemUpdateRequest) SetSkus(value string) {
	this.Set("skus", value)
}

/* 门票商品库存技术方式（拍下减库存：true，付款减库存：false） */
func (this *TicketItemUpdateRequest) SetSubStockAtBuy(value string) {
	this.Set("sub_stock_at_buy", value)
}

/* 门票商品标题<br /> 支持最大长度为：60<br /> 支持的最大列表长度为：60 */
func (this *TicketItemUpdateRequest) SetTitle(value string) {
	this.Set("title", value)
}

/* 门票商品视频-视频标识由多媒体中相关接口获取 */
func (this *TicketItemUpdateRequest) SetVideoId(value string) {
	this.Set("video_id", value)
}

/* 门票商品是否参与店铺会员打折 */
func (this *TicketItemUpdateRequest) SetVipPromoted(value string) {
	this.Set("vip_promoted", value)
}

type TicketItemUpdateResponse struct {
	*ErrorResponse                 `json:"error_response,omitempty"`
	TicketItemUpdateResponseResult `json:"ticket_item_update_response"`
}
type TicketItemUpdateResponseResult struct {
	/* 门票商品操作结果，具体请参见TicketItemProcessResult数据结构 */
	TicketItemProcessResult *TicketItemProcessResult `json:"ticket_item_process_result"`
}

/*taobao.ticket.items.get*/
type TicketItemsGetRequest struct {
	values url.Values
}

func (this *TicketItemsGetRequest) GetApiMethodName() string {
	return "taobao.ticket.items.get"
}
func (this *TicketItemsGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *TicketItemsGetRequest) GetValues() url.Values {
	return this.values
}

/* 需要返回的门票商品（TicketItem）对象字段，如title,price,skus等。<br>可选值：TicketItem商品结构体中所有字段均可返回；多个字段用“,”分隔。 */
func (this *TicketItemsGetRequest) SetFields(value string) {
	this.Set("fields", value)
}

/* 批量获取信息的商品标识。最多不能超过20个。 */
func (this *TicketItemsGetRequest) SetItemIds(value string) {
	this.Set("item_ids", value)
}

type TicketItemsGetResponse struct {
	*ErrorResponse               `json:"error_response,omitempty"`
	TicketItemsGetResponseResult `json:"ticket_items_get_response"`
}
type TicketItemsGetResponseResult struct {
	/* 商品信息 */
	TicketItems struct {
		TicketItem []*TicketItem `json:"ticket_item"`
	} `json:"TicketItems"`
}

/*taobao.ump.promotion.get*/
type UmpPromotionGetRequest struct {
	values url.Values
}

func (this *UmpPromotionGetRequest) GetApiMethodName() string {
	return "taobao.ump.promotion.get"
}
func (this *UmpPromotionGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *UmpPromotionGetRequest) GetValues() url.Values {
	return this.values
}

/* 商品id */
func (this *UmpPromotionGetRequest) SetItemId(value string) {
	this.Set("item_id", value)
}

type UmpPromotionGetResponse struct {
	*ErrorResponse                `json:"error_response,omitempty"`
	UmpPromotionGetResponseResult `json:"ump_promotion_get_response"`
}
type UmpPromotionGetResponseResult struct {
	/* 优惠详细信息 */
	Promotions *PromotionDisplayTop `json:"promotions"`
}

/*tmall.brandcat.control.get*/
type TmallBrandcatControlGetRequest struct {
	values url.Values
}

func (this *TmallBrandcatControlGetRequest) GetApiMethodName() string {
	return "tmall.brandcat.control.get"
}
func (this *TmallBrandcatControlGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *TmallBrandcatControlGetRequest) GetValues() url.Values {
	return this.values
}

type TmallBrandcatControlGetResponse struct {
	*ErrorResponse                        `json:"error_response,omitempty"`
	TmallBrandcatControlGetResponseResult `json:"tmall_brandcat_control_get_response"`
}
type TmallBrandcatControlGetResponseResult struct {
	/* 被管控的类目和品牌信息。如何判断一个产品是否被管控，是去品牌的信息和类目信息做一个交集。 */
	BrandCatControlInfo *BrandCatControlInfo `json:"brand_cat_control_info"`
}

/*tmall.brandcat.propinput.get*/
type TmallBrandcatPropinputGetRequest struct {
	values url.Values
}

func (this *TmallBrandcatPropinputGetRequest) GetApiMethodName() string {
	return "tmall.brandcat.propinput.get"
}
func (this *TmallBrandcatPropinputGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *TmallBrandcatPropinputGetRequest) GetValues() url.Values {
	return this.values
}

/* 品牌ID，如果类目没有品牌，指定null */
func (this *TmallBrandcatPropinputGetRequest) SetBrandId(value string) {
	this.Set("brand_id", value)
}

/* 类目ID */
func (this *TmallBrandcatPropinputGetRequest) SetCid(value string) {
	this.Set("cid", value)
}

/* 属性ID，如果属性有子属性，请指定最后一级子属性ID，tmall.brandcat.propinput.get返回的即为的该属性ID对应的输入特征，对于有子属性模板的情况指定顶级属性ID即可 */
func (this *TmallBrandcatPropinputGetRequest) SetPid(value string) {
	this.Set("pid", value)
}

type TmallBrandcatPropinputGetResponse struct {
	*ErrorResponse                          `json:"error_response,omitempty"`
	TmallBrandcatPropinputGetResponseResult `json:"tmall_brandcat_propinput_get_response"`
}
type TmallBrandcatPropinputGetResponseResult struct {
	/* 属性输入特征 */
	PropertyInput *PropertyInputDO `json:"property_input"`
}

/*tmall.brandcat.salespro.get*/
type TmallBrandcatSalesproGetRequest struct {
	values url.Values
}

func (this *TmallBrandcatSalesproGetRequest) GetApiMethodName() string {
	return "tmall.brandcat.salespro.get"
}
func (this *TmallBrandcatSalesproGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *TmallBrandcatSalesproGetRequest) GetValues() url.Values {
	return this.values
}

/* 被管控的品牌Id */
func (this *TmallBrandcatSalesproGetRequest) SetBrandId(value string) {
	this.Set("brand_id", value)
}

/* 被管控的类目Id */
func (this *TmallBrandcatSalesproGetRequest) SetCatId(value string) {
	this.Set("cat_id", value)
}

type TmallBrandcatSalesproGetResponse struct {
	*ErrorResponse                         `json:"error_response,omitempty"`
	TmallBrandcatSalesproGetResponseResult `json:"tmall_brandcat_salespro_get_response"`
}
type TmallBrandcatSalesproGetResponseResult struct {
	/* 获取被管控品牌和类目下的销售属性，注意，一个管控品牌和类目下可能有多个销售属性。 */
	CatBrandSaleProps struct {
		CatBrandSaleProp []*CatBrandSaleProp `json:"cat_brand_sale_prop"`
	} `json:"CatBrandSaleProps"`
}

/*tmall.brandcat.suiteconf.get*/
type TmallBrandcatSuiteconfGetRequest struct {
	values url.Values
}

func (this *TmallBrandcatSuiteconfGetRequest) GetApiMethodName() string {
	return "tmall.brandcat.suiteconf.get"
}
func (this *TmallBrandcatSuiteconfGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *TmallBrandcatSuiteconfGetRequest) GetValues() url.Values {
	return this.values
}

type TmallBrandcatSuiteconfGetResponse struct {
	*ErrorResponse                          `json:"error_response,omitempty"`
	TmallBrandcatSuiteconfGetResponseResult `json:"tmall_brandcat_suiteconf_get_response"`
}
type TmallBrandcatSuiteconfGetResponseResult struct {
	/* 套装的配置信息 */
	SuiteConfList struct {
		SuiteConfDO []*SuiteConfDO `json:"suite_conf_d_o"`
	} `json:"SuiteConfList"`
}

/*tmall.item.desc.modules.get*/
type TmallItemDescModulesGetRequest struct {
	values url.Values
}

func (this *TmallItemDescModulesGetRequest) GetApiMethodName() string {
	return "tmall.item.desc.modules.get"
}
func (this *TmallItemDescModulesGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *TmallItemDescModulesGetRequest) GetValues() url.Values {
	return this.values
}

/* 淘宝后台发布商品的叶子类目id，可通过taobao.itemcats.get查到。api 访问地址http://api.taobao.com/apidoc/api.htm?spm=0.0.0.0.CFhhk4&path=cid:3-apiId:122 */
func (this *TmallItemDescModulesGetRequest) SetCatId(value string) {
	this.Set("cat_id", value)
}

/* 商家主帐号id */
func (this *TmallItemDescModulesGetRequest) SetUsrId(value string) {
	this.Set("usr_id", value)
}

type TmallItemDescModulesGetResponse struct {
	*ErrorResponse                        `json:"error_response,omitempty"`
	TmallItemDescModulesGetResponseResult `json:"tmall_item_desc_modules_get_response"`
}
type TmallItemDescModulesGetResponseResult struct {
	/* 返回描述模块信息 */
	ModularDescInfo *ModularDescInfo `json:"modular_desc_info"`
}

/*tmall.product.books.add*/
type TmallProductBooksAddRequest struct {
	values url.Values
}

func (this *TmallProductBooksAddRequest) GetApiMethodName() string {
	return "tmall.product.books.add"
}
func (this *TmallProductBooksAddRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *TmallProductBooksAddRequest) GetValues() url.Values {
	return this.values
}

/* 摘要信息，不支持HTML代码，长度限制5000 */
func (this *TmallProductBooksAddRequest) SetAbstractMsg(value string) {
	this.Set("abstract_msg", value)
}

/* 用户的Id,需要确定自己id的可以联系接口负责人<br /> 支持最大长度为：40<br /> 支持的最大列表长度为：40 */
func (this *TmallProductBooksAddRequest) SetAppInfo(value string) {
	this.Set("app_info", value)
}

/* 作者/著者，最多填写三个人名，超出三个人的以“等”结束<br /> 支持最大长度为：80<br /> 支持的最大列表长度为：80 */
func (this *TmallProductBooksAddRequest) SetAuthor(value string) {
	this.Set("author", value)
}

/* 作者国别/地区，到国家级别<br /> 支持最大长度为：80<br /> 支持的最大列表长度为：80 */
func (this *TmallProductBooksAddRequest) SetAuthorArea(value string) {
	this.Set("author_area", value)
}

/* 条形码，数字，共13位，9787开头<br /> 支持最大长度为：20<br /> 支持的最大列表长度为：20 */
func (this *TmallProductBooksAddRequest) SetBarCode(value string) {
	this.Set("bar_code", value)
}

/* 装帧<br /> 支持最大长度为：10<br /> 支持的最大列表长度为：10 */
func (this *TmallProductBooksAddRequest) SetBookBind(value string) {
	this.Set("book_bind", value)
}

/* 开本，如：16  表示16开本 */
func (this *TmallProductBooksAddRequest) SetBookSize(value string) {
	this.Set("book_size", value)
}

/* 版本<br /> 支持的最大列表长度为：10 */
func (this *TmallProductBooksAddRequest) SetBookVersion(value string) {
	this.Set("book_version", value)
}

/* 目录，不支持HTML代码，长度限制为8000<br /> 支持最大长度为：8000<br /> 支持的最大列表长度为：8000 */
func (this *TmallProductBooksAddRequest) SetCatalog(value string) {
	this.Set("catalog", value)
}

/* 类目id */
func (this *TmallProductBooksAddRequest) SetCategoryId(value string) {
	this.Set("category_id", value)
}

/* 中图分类号，英文字母加数字组成<br /> 支持最大长度为：30<br /> 支持的最大列表长度为：30 */
func (this *TmallProductBooksAddRequest) SetChinaClassifyNo(value string) {
	this.Set("china_classify_no", value)
}

/* cip数据编号，真实格式：（XXX）第***号，
导入时，格式为XXX-***<br /> 支持最大长度为：50<br /> 支持的最大列表长度为：50 */
func (this *TmallProductBooksAddRequest) SetCip(value string) {
	this.Set("cip", value)
}

/* 点评者，只写一个人的名字，超出的以“等”结束<br /> 支持最大长度为：40<br /> 支持的最大列表长度为：40 */
func (this *TmallProductBooksAddRequest) SetCommentator(value string) {
	this.Set("commentator", value)
}

/* 用户自己的Id号 */
func (this *TmallProductBooksAddRequest) SetCustomId(value string) {
	this.Set("custom_id", value)
}

/* 图书名称(正副书名)<br /> 支持最大长度为：150<br /> 支持的最大列表长度为：150 */
func (this *TmallProductBooksAddRequest) SetDeputyName(value string) {
	this.Set("deputy_name", value)
}

/* 绘图者，只写一个人的名字，超出的以“等”结束<br /> 支持最大长度为：40<br /> 支持的最大列表长度为：40 */
func (this *TmallProductBooksAddRequest) SetDrawor(value string) {
	this.Set("drawor", value)
}

/* 编者，只写一个人的名字，超出的以“等”结束<br /> 支持最大长度为：40<br /> 支持的最大列表长度为：40 */
func (this *TmallProductBooksAddRequest) SetEditor(value string) {
	this.Set("editor", value)
}

/* 摄影者，只写一个人的名字，超出的以“等”结束<br /> 支持最大长度为：40<br /> 支持的最大列表长度为：40 */
func (this *TmallProductBooksAddRequest) SetGraphor(value string) {
	this.Set("graphor", value)
}

/* 上传图片，图片上传使用byte[]类型 */
func (this *TmallProductBooksAddRequest) SetImage(value string) {
	this.Set("image", value)
}

/* 不带‘-’的图书ISBN号
1. 位数限定：10位 末尾校验,7开头
2. 位数限定：13位、开头数字限定： 9787开头<br /> 支持最大长度为：13<br /> 支持的最大列表长度为：13 */
func (this *TmallProductBooksAddRequest) SetIsbn(value string) {
	this.Set("isbn", value)
}

/* 这条记录对应的淘宝的Ids */
func (this *TmallProductBooksAddRequest) SetItemIds(value string) {
	this.Set("item_ids", value)
}

/* 口述者，只写一个人的名字，超出的以“等”结束<br /> 支持最大长度为：40<br /> 支持的最大列表长度为：40 */
func (this *TmallProductBooksAddRequest) SetNarrator(value string) {
	this.Set("narrator", value)
}

/* 分册名<br /> 支持最大长度为：150<br /> 支持的最大列表长度为：150 */
func (this *TmallProductBooksAddRequest) SetPartName(value string) {
	this.Set("part_name", value)
}

/* 分册号<br /> 支持的最大列表长度为：10 */
func (this *TmallProductBooksAddRequest) SetPartNo(value string) {
	this.Set("part_no", value)
}

/* 策划者，只写一个人的名字，超出的以“等”结束<br /> 支持最大长度为：40<br /> 支持的最大列表长度为：40 */
func (this *TmallProductBooksAddRequest) SetPlottor(value string) {
	this.Set("plottor", value)
}

/* 价格支持整数和小数，小数保留小数点后两位，若有多个价格，价格之间通过','号分隔<br /> 支持最大长度为：40<br /> 支持的最大列表长度为：40 */
func (this *TmallProductBooksAddRequest) SetPrice(value string) {
	this.Set("price", value)
}

/* 出版社名称<br /> 支持最大长度为：80<br /> 支持的最大列表长度为：80 */
func (this *TmallProductBooksAddRequest) SetPublishCompany(value string) {
	this.Set("publish_company", value)
}

/* 出版时间，格式必须注意：年份/月份 */
func (this *TmallProductBooksAddRequest) SetPublishYm(value string) {
	this.Set("publish_ym", value)
}

/* 年号，必须是完整的年份，如：2013，当输入的值不是四位或者首位大于2时，调接口失败<br /> 支持最大长度为：10<br /> 支持的最大列表长度为：10 */
func (this *TmallProductBooksAddRequest) SetReginYear(value string) {
	this.Set("regin_year", value)
}

/* 注释者，只写一个人的名字，超出的以“等”结束<br /> 支持最大长度为：40<br /> 支持的最大列表长度为：40 */
func (this *TmallProductBooksAddRequest) SetScholisat(value string) {
	this.Set("scholisat", value)
}

/* 丛书名<br /> 支持最大长度为：150<br /> 支持的最大列表长度为：150 */
func (this *TmallProductBooksAddRequest) SetSeriesBooksName(value string) {
	this.Set("series_books_name", value)
}

/* 此字段废弃<br /> 支持最大长度为：150<br /> 支持的最大列表长度为：150 */
func (this *TmallProductBooksAddRequest) SetSpuImg(value string) {
	this.Set("spu_img", value)
}

/* 简介，不支持HTML代码，长度不能超过5000<br /> 支持最大长度为：5000<br /> 支持的最大列表长度为：5000 */
func (this *TmallProductBooksAddRequest) SetSummary(value string) {
	this.Set("summary", value)
}

/* 译者，只写一个人的名字，超出的以“等”结束<br /> 支持最大长度为：40<br /> 支持的最大列表长度为：40 */
func (this *TmallProductBooksAddRequest) SetTranslator(value string) {
	this.Set("translator", value)
}

type TmallProductBooksAddResponse struct {
	*ErrorResponse                     `json:"error_response,omitempty"`
	TmallProductBooksAddResponseResult `json:"tmall_product_books_add_response"`
}
type TmallProductBooksAddResponseResult struct {
	/* 请求相应结构 */
	ProductBooks *ProductBooks `json:"product_books"`
}

/*tmall.product.spec.add*/
type TmallProductSpecAddRequest struct {
	values url.Values
}

func (this *TmallProductSpecAddRequest) GetApiMethodName() string {
	return "tmall.product.spec.add"
}
func (this *TmallProductSpecAddRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *TmallProductSpecAddRequest) GetValues() url.Values {
	return this.values
}

/* 产品二维码 */
func (this *TmallProductSpecAddRequest) SetBarcode(value string) {
	this.Set("barcode", value)
}

/* 存放产品规格认证类型-认证图片url映射信息，格式为k:v;k:v;，其中key为认证类型数字id，value为调用tmall.product.spec.pic.upload返回的认证图片url文本 */
func (this *TmallProductSpecAddRequest) SetCertifiedPicStr(value string) {
	this.Set("certified_pic_str", value)
}

/* 存放产品规格认证类型-认证文本映射信息，格式为k:v;k:v;，其中key为认证类型数字id，value为认证文本值 */
func (this *TmallProductSpecAddRequest) SetCertifiedTxtStr(value string) {
	this.Set("certified_txt_str", value)
}

/* 产品基础色，数据格式为：pid:vid:rvid1,rvid2,rvid3;pid:vid:rvid1 */
func (this *TmallProductSpecAddRequest) SetChangeProp(value string) {
	this.Set("change_prop", value)
}

/* 用户自定义销售属性，结构：pid1:value1;pid2:value2。在 */
func (this *TmallProductSpecAddRequest) SetCustomerSpecProps(value string) {
	this.Set("customer_spec_props", value)
}

/* 产品图片 */
func (this *TmallProductSpecAddRequest) SetImage(value string) {
	this.Set("image", value)
}

/* 产品规格吊牌价，以分为单位，无默认值，上限999999999<br /> 支持最大值为：999999999<br /> 支持最小值为：0 */
func (this *TmallProductSpecAddRequest) SetLabelPrice(value string) {
	this.Set("label_price", value)
}

/* 产品上市时间 */
func (this *TmallProductSpecAddRequest) SetMarketTime(value string) {
	this.Set("market_time", value)
}

/* 产品货号 */
func (this *TmallProductSpecAddRequest) SetProductCode(value string) {
	this.Set("product_code", value)
}

/* 产品ID */
func (this *TmallProductSpecAddRequest) SetProductId(value string) {
	this.Set("product_id", value)
}

/* 产品的规格属性 */
func (this *TmallProductSpecAddRequest) SetSpecProps(value string) {
	this.Set("spec_props", value)
}

/* 规格属性别名,只允许传颜色别名<br /> 支持最大长度为：60<br /> 支持的最大列表长度为：60 */
func (this *TmallProductSpecAddRequest) SetSpecPropsAlias(value string) {
	this.Set("spec_props_alias", value)
}

type TmallProductSpecAddResponse struct {
	*ErrorResponse                    `json:"error_response,omitempty"`
	TmallProductSpecAddResponseResult `json:"tmall_product_spec_add_response"`
}
type TmallProductSpecAddResponseResult struct {
	/* 产品规格对象 */
	ProductSpec *ProductSpec `json:"product_spec"`
}

/*tmall.product.spec.get*/
type TmallProductSpecGetRequest struct {
	values url.Values
}

func (this *TmallProductSpecGetRequest) GetApiMethodName() string {
	return "tmall.product.spec.get"
}
func (this *TmallProductSpecGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *TmallProductSpecGetRequest) GetValues() url.Values {
	return this.values
}

/* 要获取信息的产品规格信息。 */
func (this *TmallProductSpecGetRequest) SetSpecId(value string) {
	this.Set("spec_id", value)
}

type TmallProductSpecGetResponse struct {
	*ErrorResponse                    `json:"error_response,omitempty"`
	TmallProductSpecGetResponseResult `json:"tmall_product_spec_get_response"`
}
type TmallProductSpecGetResponseResult struct {
	/* 返回的产品规格信息，注意，这个产品规格信息可能是等待审核的，不一定可用。根据状态判断1：表示审核通过
	3：表示等待审核。 */
	ProductSpec *ProductSpec `json:"product_spec"`
}

/*tmall.product.spec.pic.upload*/
type TmallProductSpecPicUploadRequest struct {
	values url.Values
}

func (this *TmallProductSpecPicUploadRequest) GetApiMethodName() string {
	return "tmall.product.spec.pic.upload"
}
func (this *TmallProductSpecPicUploadRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *TmallProductSpecPicUploadRequest) GetValues() url.Values {
	return this.values
}

/* 上传的认证图片文件 */
func (this *TmallProductSpecPicUploadRequest) SetCertifyPic(value string) {
	this.Set("certify_pic", value)
}

/* 上传的认证图片的认证类型<br>
1：代表产品包装正面图<br>
2：代表完整产品资质<br>
3：代表产品包装反面图<br>
4：代表产品包装侧面图<br>
5：代表产品包装条形码特写<br>
6：代表特殊用途化妆品批准文号<br>
7：代表3C认证图标<br> */
func (this *TmallProductSpecPicUploadRequest) SetCertifyType(value string) {
	this.Set("certify_type", value)
}

type TmallProductSpecPicUploadResponse struct {
	*ErrorResponse                          `json:"error_response,omitempty"`
	TmallProductSpecPicUploadResponseResult `json:"tmall_product_spec_pic_upload_response"`
}
type TmallProductSpecPicUploadResponseResult struct {
	/* 上传成功的产品规格认证图片url */
	SpecPicUrl string `json:"spec_pic_url"`
}

/*tmall.product.specs.get*/
type TmallProductSpecsGetRequest struct {
	values url.Values
}

func (this *TmallProductSpecsGetRequest) GetApiMethodName() string {
	return "tmall.product.specs.get"
}
func (this *TmallProductSpecsGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *TmallProductSpecsGetRequest) GetValues() url.Values {
	return this.values
}

/* 类目的ID号，该id必须和properties同时传入。
而且只有当product_id不传入的时候才起效果。 */
func (this *TmallProductSpecsGetRequest) SetCatId(value string) {
	this.Set("cat_id", value)
}

/* 产品的ID。这个不能和properties和cat_id同时起效果 */
func (this *TmallProductSpecsGetRequest) SetProductId(value string) {
	this.Set("product_id", value)
}

/* 关键属性的字符串，pid:vid;pid:vid
该字段必须和cat_id同时传入才起效果。 而且只有当product_id不传入的时候才起效果。 */
func (this *TmallProductSpecsGetRequest) SetProperties(value string) {
	this.Set("properties", value)
}

type TmallProductSpecsGetResponse struct {
	*ErrorResponse                     `json:"error_response,omitempty"`
	TmallProductSpecsGetResponseResult `json:"tmall_product_specs_get_response"`
}
type TmallProductSpecsGetResponseResult struct {
	/* 返回一组产品规格信息。 */
	ProductSpecs struct {
		ProductSpec []*ProductSpec `json:"product_spec"`
	} `json:"ProductSpecs"`
}

/*tmall.product.specs.ticket.get*/
type TmallProductSpecsTicketGetRequest struct {
	values url.Values
}

func (this *TmallProductSpecsTicketGetRequest) GetApiMethodName() string {
	return "tmall.product.specs.ticket.get"
}
func (this *TmallProductSpecsTicketGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *TmallProductSpecsTicketGetRequest) GetValues() url.Values {
	return this.values
}

/* 产品规格ID，多个用逗号分隔 */
func (this *TmallProductSpecsTicketGetRequest) SetSpecIds(value string) {
	this.Set("spec_ids", value)
}

type TmallProductSpecsTicketGetResponse struct {
	*ErrorResponse                           `json:"error_response,omitempty"`
	TmallProductSpecsTicketGetResponseResult `json:"tmall_product_specs_ticket_get_response"`
}
type TmallProductSpecsTicketGetResponseResult struct {
	/* 产品规格审核单信息 */
	Tickets struct {
		Ticket []*Ticket `json:"ticket"`
	} `json:"Tickets"`
}

/*tmall.product.suitespecs.get*/
type TmallProductSuitespecsGetRequest struct {
	values url.Values
}

func (this *TmallProductSuitespecsGetRequest) GetApiMethodName() string {
	return "tmall.product.suitespecs.get"
}
func (this *TmallProductSuitespecsGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *TmallProductSuitespecsGetRequest) GetValues() url.Values {
	return this.values
}

/* 类目的ID号，该id必须和properties同时传入。
而且只有当product_id不传入的时候才起效果。 */
func (this *TmallProductSuitespecsGetRequest) SetCatId(value string) {
	this.Set("cat_id", value)
}

/* 产品的ID。这个不能和properties和cat_id同时起效果 */
func (this *TmallProductSuitespecsGetRequest) SetProductId(value string) {
	this.Set("product_id", value)
}

/* 关键属性的字符串，pid:vid;pid:vid
该字段必须和cat_id同时传入才起效果。 而且只有当product_id不传入的时候才起效果。 */
func (this *TmallProductSuitespecsGetRequest) SetProperties(value string) {
	this.Set("properties", value)
}

type TmallProductSuitespecsGetResponse struct {
	*ErrorResponse                          `json:"error_response,omitempty"`
	TmallProductSuitespecsGetResponseResult `json:"tmall_product_suitespecs_get_response"`
}
type TmallProductSuitespecsGetResponseResult struct {
	/* 返回一组产品规格信息。 */
	ProductSpecs struct {
		ProductSpec []*ProductSpec `json:"product_spec"`
	} `json:"ProductSpecs"`
}

/*tmall.product.template.get*/
type TmallProductTemplateGetRequest struct {
	values url.Values
}

func (this *TmallProductTemplateGetRequest) GetApiMethodName() string {
	return "tmall.product.template.get"
}
func (this *TmallProductTemplateGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *TmallProductTemplateGetRequest) GetValues() url.Values {
	return this.values
}

/* 类目ID */
func (this *TmallProductTemplateGetRequest) SetCid(value string) {
	this.Set("cid", value)
}

type TmallProductTemplateGetResponse struct {
	*ErrorResponse                        `json:"error_response,omitempty"`
	TmallProductTemplateGetResponseResult `json:"tmall_product_template_get_response"`
}
type TmallProductTemplateGetResponseResult struct {
	/* 见SpuTemplateDO说明 */
	Template *SpuTemplateDO `json:"template"`
}

/*taobao.trade.amount.get*/
type TradeAmountGetRequest struct {
	values url.Values
}

func (this *TradeAmountGetRequest) GetApiMethodName() string {
	return "taobao.trade.amount.get"
}
func (this *TradeAmountGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *TradeAmountGetRequest) GetValues() url.Values {
	return this.values
}

/* 订单帐务详情需要返回的字段信息，可选值如下：
1. TradeAmount中可指定的fields：
tid,alipay_no,created,pay_time,end_time,total_fee,payment,post_fee,cod_fee,commission_fee,buyer_obtain_point_fee
2. OrderAmount中可指定的fields：order_amounts.oid,order_amounts.title,order_amounts.num_iid,
order_amounts.sku_properties_name,order_amounts.sku_id,order_amounts.num,order_amounts.price,order_amounts.discount_fee,order_amounts.adjust_fee,order_amounts.payment,order_amounts.promotion_name
3. order_amounts(返回OrderAmount的所有内容)
4. promotion_details(指定该值会返回主订单的promotion_details中除id之外的所有字段) */
func (this *TradeAmountGetRequest) SetFields(value string) {
	this.Set("fields", value)
}

/* 订单交易编号<br /> 支持最大值为：9223372036854775807<br /> 支持最小值为：-9223372036854775808 */
func (this *TradeAmountGetRequest) SetTid(value string) {
	this.Set("tid", value)
}

type TradeAmountGetResponse struct {
	*ErrorResponse               `json:"error_response,omitempty"`
	TradeAmountGetResponseResult `json:"trade_amount_get_response"`
}
type TradeAmountGetResponseResult struct {
	/* 主订单的财务信息详情 */
	TradeAmount *TradeAmount `json:"trade_amount"`
}

/*taobao.trade.close*/
type TradeCloseRequest struct {
	values url.Values
}

func (this *TradeCloseRequest) GetApiMethodName() string {
	return "taobao.trade.close"
}
func (this *TradeCloseRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *TradeCloseRequest) GetValues() url.Values {
	return this.values
}

/* 交易关闭原因。可以选择的理由有：
1.未及时付款
2.买家联系不上
3.谢绝还价
4.商品瑕疵
5.协商不一致
6.买家不想买
7.与买家协商一致 */
func (this *TradeCloseRequest) SetCloseReason(value string) {
	this.Set("close_reason", value)
}

/* 主订单或子订单编号。 */
func (this *TradeCloseRequest) SetTid(value string) {
	this.Set("tid", value)
}

type TradeCloseResponse struct {
	*ErrorResponse           `json:"error_response,omitempty"`
	TradeCloseResponseResult `json:"trade_close_response"`
}
type TradeCloseResponseResult struct {
	/* 关闭交易时返回的Trade信息，可用字段有tid和modified */
	Trade *Trade `json:"trade"`
}

/*taobao.trade.confirmfee.get*/
type TradeConfirmfeeGetRequest struct {
	values url.Values
}

func (this *TradeConfirmfeeGetRequest) GetApiMethodName() string {
	return "taobao.trade.confirmfee.get"
}
func (this *TradeConfirmfeeGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *TradeConfirmfeeGetRequest) GetValues() url.Values {
	return this.values
}

/* 是否是子订单。可选值:IS_FATHER(父订单),IS_CHILD(子订单) */
func (this *TradeConfirmfeeGetRequest) SetIsDetail(value string) {
	this.Set("is_detail", value)
}

/* 交易编号，或子订单编号 */
func (this *TradeConfirmfeeGetRequest) SetTid(value string) {
	this.Set("tid", value)
}

type TradeConfirmfeeGetResponse struct {
	*ErrorResponse                   `json:"error_response,omitempty"`
	TradeConfirmfeeGetResponseResult `json:"trade_confirmfee_get_response"`
}
type TradeConfirmfeeGetResponseResult struct {
	/* 获取到的交易确认收货费用 */
	TradeConfirmFee *TradeConfirmFee `json:"trade_confirm_fee"`
}

/*taobao.trade.fullinfo.get*/
type TradeFullinfoGetRequest struct {
	values url.Values
}

func (this *TradeFullinfoGetRequest) GetApiMethodName() string {
	return "taobao.trade.fullinfo.get"
}
func (this *TradeFullinfoGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *TradeFullinfoGetRequest) GetValues() url.Values {
	return this.values
}

/* 1.Trade中可以指定返回的fields：seller_nick, buyer_nick, title, type, created, tid, seller_rate,buyer_flag, buyer_rate, status, payment, adjust_fee, post_fee, total_fee, pay_time, end_time, modified, consign_time, buyer_obtain_point_fee, point_fee, real_point_fee, received_payment, commission_fee, buyer_memo, seller_memo, alipay_no,alipay_id,buyer_message, pic_path, num_iid, num, price, buyer_alipay_no, receiver_name, receiver_state, receiver_city, receiver_district, receiver_address, receiver_zip, receiver_mobile, receiver_phone,seller_flag, seller_alipay_no, seller_mobile, seller_phone, seller_name, seller_email, available_confirm_fee, has_post_fee, timeout_action_time, snapshot_url, cod_fee, cod_status, shipping_type, trade_memo, is_3D,buyer_email,buyer_area, trade_from,is_lgtype,is_force_wlb,is_brand_sale,buyer_cod_fee,discount_fee,seller_cod_fee,express_agency_fee,invoice_name,service_orders,credit_cardfee,step_trade_status,step_paid_fee,mark_desc,has_yfx,yfx_fee,yfx_id,yfx_type,trade_source(注：当该授权用户为卖家时不能查看买家buyer_memo,buyer_flag),eticket_ext,send_time, is_daixiao,is_part_consign, arrive_interval, arrive_cut_time, consign_interval,zero_purchase,alipay_point,pcc_af,2.Order中可以指定返回fields：orders.title, orders.pic_path, orders.price, orders.num, orders.num_iid, orders.sku_id, orders.refund_status, orders.status, orders.oid, orders.total_fee, orders.payment, orders.discount_fee, orders.adjust_fee, orders.snapshot_url, orders.timeout_action_time，orders.sku_properties_name, orders.item_meal_name, orders.item_meal_id,orders.buyer_rate, orders.seller_rate, orders.outer_iid, orders.outer_sku_id, orders.refund_id, orders.seller_type, orders.is_oversold,orders.end_time,orders.order_from,orders.consign_time,orders.shipping_type,orders.logistics_company,orders.invoice_no, orders.is_daixiao
3.fields：orders（返回Order的所有内容）
4.flelds：promotion_details(返回promotion_details所有内容，优惠详情),invoice_name(发票抬头),orders.is_www(子订单是否是www订单,orders.store_code(发货的仓库编码)<br>
5. field:service_tags(返回物流标签) */
func (this *TradeFullinfoGetRequest) SetFields(value string) {
	this.Set("fields", value)
}

/* 交易编号<br /> 支持最大值为：9223372036854775807<br /> 支持最小值为：1 */
func (this *TradeFullinfoGetRequest) SetTid(value string) {
	this.Set("tid", value)
}

type TradeFullinfoGetResponse struct {
	*ErrorResponse                 `json:"error_response,omitempty"`
	TradeFullinfoGetResponseResult `json:"trade_fullinfo_get_response"`
}
type TradeFullinfoGetResponseResult struct {
	/* 搜索到的交易信息列表，返回的Trade和Order中包含的具体信息为入参fields请求的字段信息 */
	Trade *Trade `json:"trade"`
}

/*taobao.trade.get*/
type TradeGetRequest struct {
	values url.Values
}

func (this *TradeGetRequest) GetApiMethodName() string {
	return "taobao.trade.get"
}
func (this *TradeGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *TradeGetRequest) GetValues() url.Values {
	return this.values
}

/* 需要返回的字段。目前支持有：<br>

1. Trade中可以指定返回的fields:seller_nick, buyer_nick, title, type, created, tid, seller_rate, buyer_rate, status, payment, discount_fee, adjust_fee, post_fee, total_fee, pay_time, end_time, modified, consign_time, buyer_obtain_point_fee, point_fee, real_point_fee, received_payment, commission_fee, buyer_memo, seller_memo, alipay_no, buyer_message, pic_path, num_iid, num, price, cod_fee, cod_status, shipping_type， is_daixiao，consign_interval，arrive_interval，arrive_cut_time <br>
2. Order中可以指定返回fields:orders.title, orders.pic_path, orders.price, orders.num, orders.num_iid, orders.sku_id, orders.refund_status, orders.status, orders.oid, orders.total_fee, orders.payment, orders.discount_fee, orders.adjust_fee, orders.sku_properties_name, orders.item_meal_name, orders.outer_sku_id, orders.outer_iid, orders.buyer_rate, orders.seller_rate， orders.is_daixiao <br>
3. fields：orders（返回Order中的所有允许返回的字段）,orders.is_wwww(是否是www订单),orders.store_code(仓库代码）<br>
4. field:service_tags(返回物流标签) */
func (this *TradeGetRequest) SetFields(value string) {
	this.Set("fields", value)
}

/* 交易编号 */
func (this *TradeGetRequest) SetTid(value string) {
	this.Set("tid", value)
}

type TradeGetResponse struct {
	*ErrorResponse         `json:"error_response,omitempty"`
	TradeGetResponseResult `json:"trade_get_response"`
}
type TradeGetResponseResult struct {
	/* 搜索到的交易信息列表，返回的Trade和Order中包含的具体信息为入参fields请求的字段信息 */
	Trade *Trade `json:"trade"`
}

/*taobao.trade.memo.add*/
type TradeMemoAddRequest struct {
	values url.Values
}

func (this *TradeMemoAddRequest) GetApiMethodName() string {
	return "taobao.trade.memo.add"
}
func (this *TradeMemoAddRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *TradeMemoAddRequest) GetValues() url.Values {
	return this.values
}

/* 交易备注旗帜，可选值为：0(灰色), 1(红色), 2(黄色), 3(绿色), 4(蓝色), 5(粉红色)，默认值为0<br /> 支持最大值为：5<br /> 支持最小值为：0 */
func (this *TradeMemoAddRequest) SetFlag(value string) {
	this.Set("flag", value)
}

/* 交易备注。最大长度: 1000个字节 */
func (this *TradeMemoAddRequest) SetMemo(value string) {
	this.Set("memo", value)
}

/* 交易编号 */
func (this *TradeMemoAddRequest) SetTid(value string) {
	this.Set("tid", value)
}

type TradeMemoAddResponse struct {
	*ErrorResponse             `json:"error_response,omitempty"`
	TradeMemoAddResponseResult `json:"trade_memo_add_response"`
}
type TradeMemoAddResponseResult struct {
	/* 对一笔交易添加备注后返回其对应的Trade，Trade中可用的返回字段有tid和created */
	Trade *Trade `json:"trade"`
}

/*taobao.trade.memo.update*/
type TradeMemoUpdateRequest struct {
	values url.Values
}

func (this *TradeMemoUpdateRequest) GetApiMethodName() string {
	return "taobao.trade.memo.update"
}
func (this *TradeMemoUpdateRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *TradeMemoUpdateRequest) GetValues() url.Values {
	return this.values
}

/* 卖家交易备注旗帜，可选值为：0(灰色), 1(红色), 2(黄色), 3(绿色), 4(蓝色), 5(粉红色)，默认值为0<br /> 支持最大值为：5<br /> 支持最小值为：0 */
func (this *TradeMemoUpdateRequest) SetFlag(value string) {
	this.Set("flag", value)
}

/* 卖家交易备注。最大长度: 1000个字节 */
func (this *TradeMemoUpdateRequest) SetMemo(value string) {
	this.Set("memo", value)
}

/* 是否对memo的值置空
若为true，则不管传入的memo字段的值是否为空，都将会对已有的memo值清空，慎用；
若用false，则会根据memo是否为空来修改memo的值：若memo为空则忽略对已有memo字段的修改，若memo非空，则使用新传入的memo覆盖已有的memo的值 */
func (this *TradeMemoUpdateRequest) SetReset(value string) {
	this.Set("reset", value)
}

/* 交易编号 */
func (this *TradeMemoUpdateRequest) SetTid(value string) {
	this.Set("tid", value)
}

type TradeMemoUpdateResponse struct {
	*ErrorResponse                `json:"error_response,omitempty"`
	TradeMemoUpdateResponseResult `json:"trade_memo_update_response"`
}
type TradeMemoUpdateResponseResult struct {
	/* 更新交易的备注信息后返回的Trade，其中可用字段为tid和modified */
	Trade *Trade `json:"trade"`
}

/*taobao.trade.ordersku.update*/
type TradeOrderskuUpdateRequest struct {
	values url.Values
}

func (this *TradeOrderskuUpdateRequest) GetApiMethodName() string {
	return "taobao.trade.ordersku.update"
}
func (this *TradeOrderskuUpdateRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *TradeOrderskuUpdateRequest) GetValues() url.Values {
	return this.values
}

/* 子订单编号（对于单笔订单的交易可以传交易编号）。 */
func (this *TradeOrderskuUpdateRequest) SetOid(value string) {
	this.Set("oid", value)
}

/* 销售属性编号，可以通过taobao.item.skus.get获取订单对应的商品的所有销售属性。 */
func (this *TradeOrderskuUpdateRequest) SetSkuId(value string) {
	this.Set("sku_id", value)
}

/* 销售属性组合串，格式：p1:v1;p2:v2，如：1627207:28329;20509:28314。可以通过taobao.item.skus.get获取订单对应的商品的所有销售属性。 */
func (this *TradeOrderskuUpdateRequest) SetSkuProps(value string) {
	this.Set("sku_props", value)
}

type TradeOrderskuUpdateResponse struct {
	*ErrorResponse                    `json:"error_response,omitempty"`
	TradeOrderskuUpdateResponseResult `json:"trade_ordersku_update_response"`
}
type TradeOrderskuUpdateResponseResult struct {
	/* 只返回oid和modified */
	Order *Order `json:"order"`
}

/*taobao.trade.postage.update*/
type TradePostageUpdateRequest struct {
	values url.Values
}

func (this *TradePostageUpdateRequest) GetApiMethodName() string {
	return "taobao.trade.postage.update"
}
func (this *TradePostageUpdateRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *TradePostageUpdateRequest) GetValues() url.Values {
	return this.values
}

/* 邮费价格(邮费单位是元） */
func (this *TradePostageUpdateRequest) SetPostFee(value string) {
	this.Set("post_fee", value)
}

/* 主订单编号 */
func (this *TradePostageUpdateRequest) SetTid(value string) {
	this.Set("tid", value)
}

type TradePostageUpdateResponse struct {
	*ErrorResponse                   `json:"error_response,omitempty"`
	TradePostageUpdateResponseResult `json:"trade_postage_update_response"`
}
type TradePostageUpdateResponseResult struct {
	/* 返回trade类型，其中包含修改时间modified，修改邮费post_fee，修改后的总费用total_fee和买家实付款payment */
	Trade *Trade `json:"trade"`
}

/*taobao.trade.receivetime.delay*/
type TradeReceivetimeDelayRequest struct {
	values url.Values
}

func (this *TradeReceivetimeDelayRequest) GetApiMethodName() string {
	return "taobao.trade.receivetime.delay"
}
func (this *TradeReceivetimeDelayRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *TradeReceivetimeDelayRequest) GetValues() url.Values {
	return this.values
}

/* 延长收货的天数，可选值为：3, 5, 7, 10。<br /> 支持最大值为：10<br /> 支持最小值为：3 */
func (this *TradeReceivetimeDelayRequest) SetDays(value string) {
	this.Set("days", value)
}

/* 主订单号 */
func (this *TradeReceivetimeDelayRequest) SetTid(value string) {
	this.Set("tid", value)
}

type TradeReceivetimeDelayResponse struct {
	*ErrorResponse                      `json:"error_response,omitempty"`
	TradeReceivetimeDelayResponseResult `json:"trade_receivetime_delay_response"`
}
type TradeReceivetimeDelayResponseResult struct {
	/* 更新后的交易数据，只包括tid和modified两个字段。 */
	Trade *Trade `json:"trade"`
}

/*taobao.trade.shippingaddress.update*/
type TradeShippingaddressUpdateRequest struct {
	values url.Values
}

func (this *TradeShippingaddressUpdateRequest) GetApiMethodName() string {
	return "taobao.trade.shippingaddress.update"
}
func (this *TradeShippingaddressUpdateRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *TradeShippingaddressUpdateRequest) GetValues() url.Values {
	return this.values
}

/* 收货地址。最大长度为228个字节。<br /> 支持最大长度为：228<br /> 支持的最大列表长度为：228 */
func (this *TradeShippingaddressUpdateRequest) SetReceiverAddress(value string) {
	this.Set("receiver_address", value)
}

/* 城市。最大长度为32个字节。如：杭州<br /> 支持最大长度为：32<br /> 支持的最大列表长度为：32 */
func (this *TradeShippingaddressUpdateRequest) SetReceiverCity(value string) {
	this.Set("receiver_city", value)
}

/* 区/县。最大长度为32个字节。如：西湖区<br /> 支持最大长度为：32<br /> 支持的最大列表长度为：32 */
func (this *TradeShippingaddressUpdateRequest) SetReceiverDistrict(value string) {
	this.Set("receiver_district", value)
}

/* 移动电话。最大长度为30个字节。<br /> 支持最大长度为：30<br /> 支持的最大列表长度为：30 */
func (this *TradeShippingaddressUpdateRequest) SetReceiverMobile(value string) {
	this.Set("receiver_mobile", value)
}

/* 收货人全名。最大长度为50个字节。<br /> 支持最大长度为：50<br /> 支持的最大列表长度为：50 */
func (this *TradeShippingaddressUpdateRequest) SetReceiverName(value string) {
	this.Set("receiver_name", value)
}

/* 固定电话。最大长度为30个字节。<br /> 支持最大长度为：30<br /> 支持的最大列表长度为：30 */
func (this *TradeShippingaddressUpdateRequest) SetReceiverPhone(value string) {
	this.Set("receiver_phone", value)
}

/* 省份。最大长度为32个字节。如：浙江<br /> 支持最大长度为：32<br /> 支持的最大列表长度为：32 */
func (this *TradeShippingaddressUpdateRequest) SetReceiverState(value string) {
	this.Set("receiver_state", value)
}

/* 邮政编码。必须由6个数字组成。<br /> 支持最大长度为：6<br /> 支持的最大列表长度为：6 */
func (this *TradeShippingaddressUpdateRequest) SetReceiverZip(value string) {
	this.Set("receiver_zip", value)
}

/* 交易编号。 */
func (this *TradeShippingaddressUpdateRequest) SetTid(value string) {
	this.Set("tid", value)
}

type TradeShippingaddressUpdateResponse struct {
	*ErrorResponse                           `json:"error_response,omitempty"`
	TradeShippingaddressUpdateResponseResult `json:"trade_shippingaddress_update_response"`
}
type TradeShippingaddressUpdateResponseResult struct {
	/* 交易结构 */
	Trade *Trade `json:"trade"`
}

/*taobao.trade.snapshot.get*/
type TradeSnapshotGetRequest struct {
	values url.Values
}

func (this *TradeSnapshotGetRequest) GetApiMethodName() string {
	return "taobao.trade.snapshot.get"
}
func (this *TradeSnapshotGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *TradeSnapshotGetRequest) GetValues() url.Values {
	return this.values
}

/* 需要返回的字段列表。现只支持："snapshot"字段 */
func (this *TradeSnapshotGetRequest) SetFields(value string) {
	this.Set("fields", value)
}

/* 交易编号 */
func (this *TradeSnapshotGetRequest) SetTid(value string) {
	this.Set("tid", value)
}

type TradeSnapshotGetResponse struct {
	*ErrorResponse                 `json:"error_response,omitempty"`
	TradeSnapshotGetResponseResult `json:"trade_snapshot_get_response"`
}
type TradeSnapshotGetResponseResult struct {
	/* 只包含Trade中的tid和snapshot、子订单Order中的oid和snapshot */
	Trade *Trade `json:"trade"`
}

/*taobao.trades.sold.get*/
type TradesSoldGetRequest struct {
	values url.Values
}

func (this *TradesSoldGetRequest) GetApiMethodName() string {
	return "taobao.trades.sold.get"
}
func (this *TradesSoldGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *TradesSoldGetRequest) GetValues() url.Values {
	return this.values
}

/* 买家昵称 */
func (this *TradesSoldGetRequest) SetBuyerNick(value string) {
	this.Set("buyer_nick", value)
}

/* 查询交易创建时间结束。格式:yyyy-MM-dd HH:mm:ss */
func (this *TradesSoldGetRequest) SetEndCreated(value string) {
	this.Set("end_created", value)
}

/* 可选值有ershou(二手市场的订单）,service（商城服务子订单）mark（双十一大促活动异常订单）作为扩展类型筛选只能做单个ext_type查询，不能全部查询所有的ext_type类型 */
func (this *TradesSoldGetRequest) SetExtType(value string) {
	this.Set("ext_type", value)
}

/* 需要返回的字段。目前支持有：<br>
1. Trade中可以指定返回的fields:<br>
seller_nick, buyer_nick, title, type, created,  tid, seller_rate,seller_can_rate, buyer_rate,can_rate, status, payment, discount_fee, adjust_fee, post_fee, total_fee, pay_time, end_time, modified, consign_time, buyer_obtain_point_fee, point_fee, real_point_fee, received_payment,  pic_path, num_iid, num, price, cod_fee, cod_status, shipping_type, receiver_name, receiver_state, receiver_city, receiver_district, receiver_address, receiver_zip, receiver_mobile, receiver_phone,seller_flag,alipay_id,alipay_no,is_lgtype,is_force_wlb,is_brand_sale,buyer_area,has_buyer_message, credit_card_fee, lg_aging_type, lg_aging, step_trade_status,step_paid_fee,mark_desc,has_yfx,yfx_fee,yfx_id,yfx_type,trade_source,send_time,is_daixiao,is_wt,is_part_consign,zero_purchase
<br>
2. Order中可以指定返回fields：orders.title, orders.pic_path, orders.price, orders.num, orders.num_iid, orders.sku_id, orders.refund_status, orders.status, orders.oid, orders.total_fee, orders.payment, orders.discount_fee, orders.adjust_fee, orders.sku_properties_name, orders.item_meal_name, orders.buyer_rate, orders.seller_rate, orders.outer_iid, orders.outer_sku_id, orders.refund_id, orders.seller_type, orders.end_time,orders.order_from,orders.consign_time,orders.shipping_type,orders.logistics_company,orders.invoice_no,orders.is_daixiao<br>
3. fields：orders（返回2中Order的所有内容）
4.fields:service_orders(返回service_order中所有内容) */
func (this *TradesSoldGetRequest) SetFields(value string) {
	this.Set("fields", value)
}

/* 页码。取值范围:大于零的整数; 默认值:1 */
func (this *TradesSoldGetRequest) SetPageNo(value string) {
	this.Set("page_no", value)
}

/* 每页条数。取值范围:大于零的整数; 默认值:40;最大值:100 */
func (this *TradesSoldGetRequest) SetPageSize(value string) {
	this.Set("page_size", value)
}

/* 评价状态，默认查询所有评价状态的数据，除了默认值外每次只能查询一种状态。<br>
可选值：
RATE_UNBUYER(买家未评)
RATE_UNSELLER(卖家未评)
RATE_BUYER_UNSELLER(买家已评，卖家未评)
RATE_UNBUYER_SELLER(买家未评，卖家已评)
RATE_BUYER_SELLER(买家已评,卖家已评) */
func (this *TradesSoldGetRequest) SetRateStatus(value string) {
	this.Set("rate_status", value)
}

/* 查询三个月内交易创建时间开始。格式:yyyy-MM-dd HH:mm:ss */
func (this *TradesSoldGetRequest) SetStartCreated(value string) {
	this.Set("start_created", value)
}

/* 交易状态，默认查询所有交易状态的数据，除了默认值外每次只能查询一种状态。
可选值：
TRADE_NO_CREATE_PAY(没有创建支付宝交易)
WAIT_BUYER_PAY(等待买家付款)
WAIT_SELLER_SEND_GOODS(等待卖家发货,即:买家已付款)
SELLER_CONSIGNED_PART（卖家部分发货）
WAIT_BUYER_CONFIRM_GOODS(等待买家确认收货,即:卖家已发货)
TRADE_BUYER_SIGNED(买家已签收,货到付款专用)
TRADE_FINISHED(交易成功)
TRADE_CLOSED(交易关闭)
TRADE_CLOSED_BY_TAOBAO(交易被淘宝关闭)
ALL_WAIT_PAY(包含：WAIT_BUYER_PAY、TRADE_NO_CREATE_PAY)
ALL_CLOSED(包含：TRADE_CLOSED、TRADE_CLOSED_BY_TAOBAO)
WAIT_PRE_AUTH_CONFIRM(余额宝0元购合约中) */
func (this *TradesSoldGetRequest) SetStatus(value string) {
	this.Set("status", value)
}

/* 卖家对交易的自定义分组标签，目前可选值为：time_card（点卡软件代充），fee_card（话费软件代充） */
func (this *TradesSoldGetRequest) SetTag(value string) {
	this.Set("tag", value)
}

/* 交易类型列表，同时查询多种交易类型可用逗号分隔。<span style="color:red;font-weight: bold;">默认同时查询guarantee_trade, auto_delivery, ec, cod,step的5种交易类型的数据；查询所有交易类型的数据，需要设置下面全部可选值。</span>
可选值：
fixed(一口价)
auction(拍卖)
guarantee_trade(一口价、拍卖)
step(分阶段付款，万人团，阶梯团订单）
independent_simple_trade(旺店入门版交易)
independent_shop_trade(旺店标准版交易)
auto_delivery(自动发货)
ec(直冲)
cod(货到付款)
game_equipment(游戏装备)
shopex_trade(ShopEX交易)
netcn_trade(万网交易)
external_trade(统一外部交易)
instant_trade (即时到账)
b2c_cod(大商家货到付款)
hotel_trade(酒店类型交易)
super_market_trade(商超交易)
super_market_cod_trade(商超货到付款交易)
taohua(淘花网交易类型）
waimai(外卖交易类型）
o2o_offlinetrade（O2O交易）
nopaid（即时到帐/趣味猜交易类型）
step (万人团) eticket(电子凭证)
tmall_i18n（天猫国际）;nopaid （无付款交易）insurance_plus（保险）finance（基金）
注：guarantee_trade是一个组合查询条件，并不是一种交易类型，获取批量或单个订单中不会返回此种类型的订单。
pre_auth_type(预授权0元购) */
func (this *TradesSoldGetRequest) SetType(value string) {
	this.Set("type", value)
}

/* 是否启用has_next的分页方式，如果指定true,则返回的结果中不包含总记录数，但是会新增一个是否存在下一页的的字段，通过此种方式获取增量交易，接口调用成功率在原有的基础上有所提升。 */
func (this *TradesSoldGetRequest) SetUseHasNext(value string) {
	this.Set("use_has_next", value)
}

type TradesSoldGetResponse struct {
	*ErrorResponse              `json:"error_response,omitempty"`
	TradesSoldGetResponseResult `json:"trades_sold_get_response"`
}
type TradesSoldGetResponseResult struct {
	/* 是否存在下一页 */
	HasNext bool `json:"has_next"`
	/* 搜索到的交易信息总数 */
	TotalResults int64 `json:"total_results"`
	/* 搜索到的交易信息列表，返回的Trade和Order中包含的具体信息为入参fields请求的字段信息 */
	Trades struct {
		Trade []*Trade `json:"trade"`
	} `json:"Trades"`
}

/*taobao.trades.sold.increment.get*/
type TradesSoldIncrementGetRequest struct {
	values url.Values
}

func (this *TradesSoldIncrementGetRequest) GetApiMethodName() string {
	return "taobao.trades.sold.increment.get"
}
func (this *TradesSoldIncrementGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *TradesSoldIncrementGetRequest) GetValues() url.Values {
	return this.values
}

/* 查询修改结束时间，必须大于修改开始时间(修改时间跨度不能大于一天)，格式:yyyy-MM-dd HH:mm:ss。<span style="color:red;font-weight: bold;">建议使用30分钟以内的时间跨度，能大大提高响应速度和成功率</span>。 */
func (this *TradesSoldIncrementGetRequest) SetEndModified(value string) {
	this.Set("end_modified", value)
}

/* 可选值有ershou(二手市场的订单）,service（商城服务子订单）mark（双十一大促活动异常订单）作为扩展类型筛选只能做单个ext_type查询，不能全部查询所有的ext_type类型 */
func (this *TradesSoldIncrementGetRequest) SetExtType(value string) {
	this.Set("ext_type", value)
}

/* 需要返回的字段。目前支持有：
1.Trade中可以指定返回的fields:seller_nick, buyer_nick, title, type, created, tid, seller_rate,seller_can_rate, buyer_rate,can_rate,status, payment, discount_fee, adjust_fee, post_fee, total_fee, pay_time, end_time, modified, consign_time, buyer_obtain_point_fee, point_fee, real_point_fee, received_payment,pic_path, num_iid, num, price, cod_fee, cod_status, shipping_type, receiver_name, receiver_state, receiver_city, receiver_district, receiver_address, receiver_zip, receiver_mobile, receiver_phone,alipay_id,alipay_no,is_lgtype,is_force_wlb,is_brand_sale,has_buyer_message,credit_card_fee,step_trade_status,step_paid_fee,mark_desc,send_time,,has_yfx,yfx_fee,yfx_id,yfx_type,trade_source,seller_flag,is_daixiao,is_part_consign,zero_purchase
2.Order中可以指定返回fields：
orders.title, orders.pic_path, orders.price, orders.num, orders.num_iid, orders.sku_id, orders.refund_status, orders.status, orders.oid, orders.total_fee, orders.payment, orders.discount_fee, orders.adjust_fee, orders.sku_properties_name, orders.item_meal_name, orders.buyer_rate, orders.seller_rate, orders.outer_iid, orders.outer_sku_id, orders.refund_id, orders.seller_type，orders.end_time,,orders.consign_time,orders.shipping_type,orders.logistics_company,orders.invoice_no,orders.is_daixiao
3.fields：orders（返回Order的所有内容）
4.fields:service_orders(返回service_order中所有内容) */
func (this *TradesSoldIncrementGetRequest) SetFields(value string) {
	this.Set("fields", value)
}

/* 页码。取值范围:大于零的整数;默认值:1。<span style="color:red;font-weight: bold;">注：必须采用倒序的分页方式（从最后一页往回取）才能避免漏单问题。</span> */
func (this *TradesSoldIncrementGetRequest) SetPageNo(value string) {
	this.Set("page_no", value)
}

/* 每页条数。取值范围：1~100，默认值：40。<span style="color:red;font-weight: bold;">建议使用40~50，可以提高成功率，减少超时数量</span>。 */
func (this *TradesSoldIncrementGetRequest) SetPageSize(value string) {
	this.Set("page_size", value)
}

/* 查询修改开始时间(修改时间跨度不能大于一天)。格式:yyyy-MM-dd HH:mm:ss */
func (this *TradesSoldIncrementGetRequest) SetStartModified(value string) {
	this.Set("start_modified", value)
}

/* 交易状态，默认查询所有交易状态的数据，除了默认值外每次只能查询一种状态。 可选值 TRADE_NO_CREATE_PAY(没有创建支付宝交易) WAIT_BUYER_PAY(等待买家付款)
SELLER_CONSIGNED_PART（卖家部分发货）
WAIT_SELLER_SEND_GOODS(等待卖家发货,即:买家已付款) WAIT_BUYER_CONFIRM_GOODS(等待买家确认收货,即:卖家已发货) TRADE_BUYER_SIGNED(买家已签收,货到付款专用) TRADE_FINISHED(交易成功) TRADE_CLOSED(交易关闭) TRADE_CLOSED_BY_TAOBAO(交易被淘宝关闭) ALL_WAIT_PAY(包含：WAIT_BUYER_PAY、TRADE_NO_CREATE_PAY) ALL_CLOSED(包含：TRADE_CLOSED、TRADE_CLOSED_BY_TAOBAO)
WAIT_PRE_AUTH_CONFIRM(余额宝0元购合约中) */
func (this *TradesSoldIncrementGetRequest) SetStatus(value string) {
	this.Set("status", value)
}

/* 卖家对交易的自定义分组标签，目前可选值为：time_card（点卡软件代充），fee_card（话费软件代充） */
func (this *TradesSoldIncrementGetRequest) SetTag(value string) {
	this.Set("tag", value)
}

/* 交易类型列表，同时查询多种交易类型可用逗号分隔。<span style="color:red;font-weight: bold;">默认同时查询guarantee_trade, auto_delivery, ec, cod,step的5种交易类型的数据；查询所有交易类型的数据，需要设置下面全部可选值。</span>
可选值：
fixed(一口价)
auction(拍卖)
step（分阶段付款，万人团，阶梯团订单）
guarantee_trade(一口价、拍卖)
independent_simple_trade(旺店入门版交易)
independent_shop_trade(旺店标准版交易)
auto_delivery(自动发货)
ec(直冲) cod(货到付款)
fenxiao(分销)
game_equipment(游戏装备)
shopex_trade(ShopEX交易)
netcn_trade(万网交易)
external_trade(统一外部交易)
instant_trade (即时到账)
b2c_cod(大商家货到付款)
hotel_trade(酒店类型交易)
super_market_trade(商超交易),
super_market_cod_trade(商超货到付款交易)
taohua(桃花网交易类型）
waimai(外卖交易类型）
o2o_offlinetrade（O2O交易）
nopaid（即时到帐/趣味猜交易类型）
 eticket(电子凭证)
tmall_i18n（天猫国际）;nopaid（无付款交易）insurance_plus（保险）finance（基金）
注：guarantee_trade是一个组合查询条件，并不是一种交易类型，获取批量或单个订单中不会返回此种类型的订单。
pre_auth_type(预授权0元购) */
func (this *TradesSoldIncrementGetRequest) SetType(value string) {
	this.Set("type", value)
}

/* 是否启用has_next的分页方式，如果指定true,则返回的结果中不包含总记录数，但是会新增一个是否存在下一页的的字段，<span style="color:red;font-weight: bold;">通过此种方式获取增量交易，效率在原有的基础上有80%的提升</span>。 */
func (this *TradesSoldIncrementGetRequest) SetUseHasNext(value string) {
	this.Set("use_has_next", value)
}

type TradesSoldIncrementGetResponse struct {
	*ErrorResponse                       `json:"error_response,omitempty"`
	TradesSoldIncrementGetResponseResult `json:"trades_sold_increment_get_response"`
}
type TradesSoldIncrementGetResponseResult struct {
	/* 是否存在下一页 */
	HasNext bool `json:"has_next"`
	/* 搜索到的交易信息总数 */
	TotalResults int64 `json:"total_results"`
	/* 搜索到的交易信息列表，返回的Trade和Order中包含的具体信息为入参fields请求的字段信息 */
	Trades struct {
		Trade []*Trade `json:"trade"`
	} `json:"Trades"`
}

/*taobao.trades.sold.incrementv.get*/
type TradesSoldIncrementvGetRequest struct {
	values url.Values
}

func (this *TradesSoldIncrementvGetRequest) GetApiMethodName() string {
	return "taobao.trades.sold.incrementv.get"
}
func (this *TradesSoldIncrementvGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *TradesSoldIncrementvGetRequest) GetValues() url.Values {
	return this.values
}

/* 查询入库结束时间，必须大于入库开始时间(修改时间跨度不能大于一天)，格式:yyyy-MM-dd HH:mm:ss。<span style="color:red;font-weight: bold;">建议使用30分钟以内的时间跨度，能大大提高响应速度和成功率</span>。 */
func (this *TradesSoldIncrementvGetRequest) SetEndCreate(value string) {
	this.Set("end_create", value)
}

/* 可选值有ershou(二手市场的订单）,service（商城服务子订单）mark（双十一大促活动异常订单）作为扩展类型筛选只能做单个ext_type查询，不能全部查询所有的ext_type类型 */
func (this *TradesSoldIncrementvGetRequest) SetExtType(value string) {
	this.Set("ext_type", value)
}

/* 需要返回的字段。目前支持有：
1.Trade中可以指定返回的fields:seller_nick, buyer_nick, title, type, created, tid, seller_rate, buyer_rate, status, payment, discount_fee, adjust_fee, post_fee, total_fee, pay_time, end_time, modified, consign_time, buyer_obtain_point_fee, point_fee, real_point_fee, received_payment,pic_path, num_iid, num, price, cod_fee, cod_status, shipping_type, receiver_name, receiver_state, receiver_city, receiver_district, receiver_address, receiver_zip, receiver_mobile, receiver_phone,alipay_id,alipay_no,is_lgtype,is_force_wlb,is_brand_sale,has_buyer_message,credit_card_fee,step_trade_status,step_paid_fee,mark_desc，is_daixiao,is_part_consign
2.Order中可以指定返回fields：
orders.title, orders.pic_path, orders.price, orders.num, orders.num_iid, orders.sku_id, orders.refund_status, orders.status, orders.oid, orders.total_fee, orders.payment, orders.discount_fee, orders.adjust_fee, orders.sku_properties_name, orders.item_meal_name, orders.buyer_rate, orders.seller_rate, orders.outer_iid, orders.outer_sku_id, orders.refund_id, orders.seller_type，orders.end_time, orders.order_from,orders.consign_time,orders.shipping_type,orders.logistics_company,orders.invoice_no，orders.is_daixiao
3.fields：orders（返回Order的所有内容）
4.fields:service_orders(返回service_order中所有内容),,orders.is_www(是否是www订单）,orders.store_code（仓库地址） */
func (this *TradesSoldIncrementvGetRequest) SetFields(value string) {
	this.Set("fields", value)
}

/* 页码。取值范围:大于零的整数;默认值:1。<span style="color:red;font-weight: bold;">注：必须采用倒序的分页方式（从最后一页往回取）才能避免漏单问题。</span> */
func (this *TradesSoldIncrementvGetRequest) SetPageNo(value string) {
	this.Set("page_no", value)
}

/* 每页条数。取值范围：1~100，默认值：40。<span style="color:red;font-weight: bold;">建议使用40~50，可以提高成功率，减少超时数量</span>。 */
func (this *TradesSoldIncrementvGetRequest) SetPageSize(value string) {
	this.Set("page_size", value)
}

/* 查询入库开始时间(修改时间跨度不能大于一天)。格式:yyyy-MM-dd HH:mm:ss */
func (this *TradesSoldIncrementvGetRequest) SetStartCreate(value string) {
	this.Set("start_create", value)
}

/* 交易状态，默认查询所有交易状态的数据，除了默认值外每次只能查询一种状态。 可选值 TRADE_NO_CREATE_PAY(没有创建支付宝交易) WAIT_BUYER_PAY(等待买家付款) WAIT_SELLER_SEND_GOODS(等待卖家发货,即:买家已付款)
SELLER_CONSIGNED_PART（卖家部分发货）
WAIT_BUYER_CONFIRM_GOODS(等待买家确认收货,即:卖家已发货) TRADE_BUYER_SIGNED(买家已签收,货到付款专用) TRADE_FINISHED(交易成功) TRADE_CLOSED(交易关闭) TRADE_CLOSED_BY_TAOBAO(交易被淘宝关闭) ALL_WAIT_PAY(包含：WAIT_BUYER_PAY、TRADE_NO_CREATE_PAY) ALL_CLOSED(包含：TRADE_CLOSED、TRADE_CLOSED_BY_TAOBAO) */
func (this *TradesSoldIncrementvGetRequest) SetStatus(value string) {
	this.Set("status", value)
}

/* 卖家对交易的自定义分组标签，目前可选值为：time_card（点卡软件代充），fee_card（话费软件代充） */
func (this *TradesSoldIncrementvGetRequest) SetTag(value string) {
	this.Set("tag", value)
}

/* 交易类型列表，同时查询多种交易类型可用逗号分隔。<span style="color:red;font-weight: bold;">默认同时查询guarantee_trade, auto_delivery, ec, cod,step的5种交易类型的数据；查询所有交易类型的数据，需要设置下面全部可选值。</span>
可选值：
fixed(一口价)
auction(拍卖)
step（分阶段付款，万人团，阶梯团订单）
guarantee_trade(一口价、拍卖)
independent_simple_trade(旺店入门版交易)
independent_shop_trade(旺店标准版交易)
auto_delivery(自动发货)
ec(直冲) cod(货到付款)
fenxiao(分销)
game_equipment(游戏装备)
shopex_trade(ShopEX交易)
netcn_trade(万网交易)
external_trade(统一外部交易)
instant_trade (即时到账)
b2c_cod(大商家货到付款)
hotel_trade(酒店类型交易)
super_market_trade(商超交易),
super_market_cod_trade(商超货到付款交易)
taohua(桃花网交易类型）
waimai(外卖交易类型）
o2o_offlinetrade（O2O交易）
nopaid（无付款订单）
eticket(电子凭证)
tmall_i18n（天猫国际）
注：guarantee_trade是一个组合查询条件，并不是一种交易类型，获取批量或单个订单中不会返回此种类型的订单。 */
func (this *TradesSoldIncrementvGetRequest) SetType(value string) {
	this.Set("type", value)
}

/* 是否启用has_next的分页方式，如果指定true,则返回的结果中不包含总记录数，但是会新增一个是否存在下一页的的字段，<span style="color:red;font-weight: bold;">通过此种方式获取增量交易，效率在原有的基础上有80%的提升</span>。 */
func (this *TradesSoldIncrementvGetRequest) SetUseHasNext(value string) {
	this.Set("use_has_next", value)
}

type TradesSoldIncrementvGetResponse struct {
	*ErrorResponse                        `json:"error_response,omitempty"`
	TradesSoldIncrementvGetResponseResult `json:"trades_sold_incrementv_get_response"`
}
type TradesSoldIncrementvGetResponseResult struct {
	/* 是否存在下一页 */
	HasNext bool `json:"has_next"`
	/* 搜索到的交易信息总数 */
	TotalResults int64 `json:"total_results"`
	/* 搜索到的交易信息列表，返回的Trade和Order中包含的具体信息为入参fields请求的字段信息 */
	Trades struct {
		Trade []*Trade `json:"trade"`
	} `json:"Trades"`
}

/*taobao.traderate.add*/
type TraderateAddRequest struct {
	values url.Values
}

func (this *TraderateAddRequest) GetApiMethodName() string {
	return "taobao.traderate.add"
}
func (this *TraderateAddRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *TraderateAddRequest) GetValues() url.Values {
	return this.values
}

/* 是否匿名,卖家评不能匿名。可选值:true(匿名),false(非匿名)。注意：如果交易订单匿名，则评价也匿名 */
func (this *TraderateAddRequest) SetAnony(value string) {
	this.Set("anony", value)
}

/* 评价内容,最大长度: 500个汉字 .注意：当评价结果为good时就不用输入评价内容.评价内容为neutral/bad的时候需要输入评价内容 */
func (this *TraderateAddRequest) SetContent(value string) {
	this.Set("content", value)
}

/* 子订单ID */
func (this *TraderateAddRequest) SetOid(value string) {
	this.Set("oid", value)
}

/* 评价结果,可选值:good(好评),neutral(中评),bad(差评) */
func (this *TraderateAddRequest) SetResult(value string) {
	this.Set("result", value)
}

/* 评价者角色,可选值:seller(卖家),buyer(买家) */
func (this *TraderateAddRequest) SetRole(value string) {
	this.Set("role", value)
}

/* 交易ID */
func (this *TraderateAddRequest) SetTid(value string) {
	this.Set("tid", value)
}

type TraderateAddResponse struct {
	*ErrorResponse             `json:"error_response,omitempty"`
	TraderateAddResponseResult `json:"traderate_add_response"`
}
type TraderateAddResponseResult struct {
	/* 返回tid、oid、create */
	TradeRate *TradeRate `json:"trade_rate"`
}

/*taobao.traderate.explain.add*/
type TraderateExplainAddRequest struct {
	values url.Values
}

func (this *TraderateExplainAddRequest) GetApiMethodName() string {
	return "taobao.traderate.explain.add"
}
func (this *TraderateExplainAddRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *TraderateExplainAddRequest) GetValues() url.Values {
	return this.values
}

/* 子订单ID */
func (this *TraderateExplainAddRequest) SetOid(value string) {
	this.Set("oid", value)
}

/* 评价解释内容,最大长度: 500个汉字 */
func (this *TraderateExplainAddRequest) SetReply(value string) {
	this.Set("reply", value)
}

type TraderateExplainAddResponse struct {
	*ErrorResponse                    `json:"error_response,omitempty"`
	TraderateExplainAddResponseResult `json:"traderate_explain_add_response"`
}
type TraderateExplainAddResponseResult struct {
	/* 商城卖家给评价解释是否成功。 */
	IsSuccess bool `json:"is_success"`
}

/*taobao.traderate.impr.imprword.byaucid.get*/
type TraderateImprImprwordByaucidGetRequest struct {
	values url.Values
}

func (this *TraderateImprImprwordByaucidGetRequest) GetApiMethodName() string {
	return "taobao.traderate.impr.imprword.byaucid.get"
}
func (this *TraderateImprImprwordByaucidGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *TraderateImprImprwordByaucidGetRequest) GetValues() url.Values {
	return this.values
}

/* 淘宝的商品id */
func (this *TraderateImprImprwordByaucidGetRequest) SetAuctionId(value string) {
	this.Set("auction_id", value)
}

type TraderateImprImprwordByaucidGetResponse struct {
	*ErrorResponse                                `json:"error_response,omitempty"`
	TraderateImprImprwordByaucidGetResponseResult `json:"traderate_impr_imprword_byaucid_get_response"`
}
type TraderateImprImprwordByaucidGetResponseResult struct {
	/* 返回的大家印象的结构体 */
	ImprWordList struct {
		ImprItemDO []*ImprItemDO `json:"impr_item_d_o"`
	} `json:"ImprWordList"`
}

/*taobao.traderate.impr.imprword.byfeedid.get*/
type TraderateImprImprwordByfeedidGetRequest struct {
	values url.Values
}

func (this *TraderateImprImprwordByfeedidGetRequest) GetApiMethodName() string {
	return "taobao.traderate.impr.imprword.byfeedid.get"
}
func (this *TraderateImprImprwordByfeedidGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *TraderateImprImprwordByfeedidGetRequest) GetValues() url.Values {
	return this.values
}

/* 交易订单id号（如果包含子订单，请使用子订单id号） */
func (this *TraderateImprImprwordByfeedidGetRequest) SetChildTradeId(value string) {
	this.Set("child_trade_id", value)
}

type TraderateImprImprwordByfeedidGetResponse struct {
	*ErrorResponse                                 `json:"error_response,omitempty"`
	TraderateImprImprwordByfeedidGetResponseResult `json:"traderate_impr_imprword_byfeedid_get_response"`
}
type TraderateImprImprwordByfeedidGetResponseResult struct {
	/* 根据子订单和买家昵称找到的评价和印象词结果 */
	ImprFeed *ImprFeedIdDO `json:"impr_feed"`
}

/*taobao.traderate.impr.imprwords.get*/
type TraderateImprImprwordsGetRequest struct {
	values url.Values
}

func (this *TraderateImprImprwordsGetRequest) GetApiMethodName() string {
	return "taobao.traderate.impr.imprwords.get"
}
func (this *TraderateImprImprwordsGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *TraderateImprImprwordsGetRequest) GetValues() url.Values {
	return this.values
}

/* 淘宝叶子类目id */
func (this *TraderateImprImprwordsGetRequest) SetCatLeafId(value string) {
	this.Set("cat_leaf_id", value)
}

/* 淘宝一级类目id */
func (this *TraderateImprImprwordsGetRequest) SetCatRootId(value string) {
	this.Set("cat_root_id", value)
}

type TraderateImprImprwordsGetResponse struct {
	*ErrorResponse                          `json:"error_response,omitempty"`
	TraderateImprImprwordsGetResponseResult `json:"traderate_impr_imprwords_get_response"`
}
type TraderateImprImprwordsGetResponseResult struct {
	/* 返回类目下所有大家印象的标签 */
	ImprWords []string `json:"impr_words"`
}

/*taobao.traderate.list.add*/
type TraderateListAddRequest struct {
	values url.Values
}

func (this *TraderateListAddRequest) GetApiMethodName() string {
	return "taobao.traderate.list.add"
}
func (this *TraderateListAddRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *TraderateListAddRequest) GetValues() url.Values {
	return this.values
}

/* 是否匿名，卖家评不能匿名。可选值:true(匿名),false(非匿名)。 注意：如果买家匿名购买，那么买家的评价默认匿名 */
func (this *TraderateListAddRequest) SetAnony(value string) {
	this.Set("anony", value)
}

/* 评价内容,最大长度: 500个汉字 .注意：当评价结果为good时就不用输入评价内容.评价内容为neutral/bad的时候需要输入评价内容 */
func (this *TraderateListAddRequest) SetContent(value string) {
	this.Set("content", value)
}

/* 评价结果。可选值:good(好评),neutral(中评),bad(差评) */
func (this *TraderateListAddRequest) SetResult(value string) {
	this.Set("result", value)
}

/* 评价者角色。可选值:seller(卖家),buyer(买家) */
func (this *TraderateListAddRequest) SetRole(value string) {
	this.Set("role", value)
}

/* 交易ID */
func (this *TraderateListAddRequest) SetTid(value string) {
	this.Set("tid", value)
}

type TraderateListAddResponse struct {
	*ErrorResponse                 `json:"error_response,omitempty"`
	TraderateListAddResponseResult `json:"traderate_list_add_response"`
}
type TraderateListAddResponseResult struct {
	/* 返回的评论的信息，仅返回tid和created字段 */
	TradeRate *TradeRate `json:"trade_rate"`
}

/*taobao.traderates.get*/
type TraderatesGetRequest struct {
	values url.Values
}

func (this *TraderatesGetRequest) GetApiMethodName() string {
	return "taobao.traderates.get"
}
func (this *TraderatesGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *TraderatesGetRequest) GetValues() url.Values {
	return this.values
}

/* 评价结束时间。如果只输入结束时间，那么全部返回所有评价数据。 */
func (this *TraderatesGetRequest) SetEndDate(value string) {
	this.Set("end_date", value)
}

/* 需返回的字段列表。可选值：TradeRate 结构中的所有字段，多个字段之间用“,”分隔 */
func (this *TraderatesGetRequest) SetFields(value string) {
	this.Set("fields", value)
}

/* 商品的数字ID */
func (this *TraderatesGetRequest) SetNumIid(value string) {
	this.Set("num_iid", value)
}

/* 页码。取值范围:大于零的整数最大限制为200; 默认值:1 */
func (this *TraderatesGetRequest) SetPageNo(value string) {
	this.Set("page_no", value)
}

/* 每页获取条数。默认值40，最小值1，最大值150。<br /> 支持最大值为：150<br /> 支持最小值为：1 */
func (this *TraderatesGetRequest) SetPageSize(value string) {
	this.Set("page_size", value)
}

/* 评价类型。可选值:get(得到),give(给出) */
func (this *TraderatesGetRequest) SetRateType(value string) {
	this.Set("rate_type", value)
}

/* 评价结果。可选值:good(好评),neutral(中评),bad(差评) */
func (this *TraderatesGetRequest) SetResult(value string) {
	this.Set("result", value)
}

/* 评价者角色即评价的发起方。可选值:seller(卖家),buyer(买家)。 当 give buyer 以买家身份给卖家的评价； 当 get seller 以买家身份得到卖家给的评价； 当 give seller 以卖家身份给买家的评价； 当 get buyer 以卖家身份得到买家给的评价。 */
func (this *TraderatesGetRequest) SetRole(value string) {
	this.Set("role", value)
}

/* 评价开始时。如果只输入开始时间，那么能返回开始时间之后的评价数据。 */
func (this *TraderatesGetRequest) SetStartDate(value string) {
	this.Set("start_date", value)
}

/* 交易订单id，可以是父订单id号，也可以是子订单id号 */
func (this *TraderatesGetRequest) SetTid(value string) {
	this.Set("tid", value)
}

/* 是否启用has_next的分页方式，如果指定true,则返回的结果中不包含总记录数，但是会新增一个是否存在下一页的的字段，通过此种方式获取评价信息，效率在原有的基础上有80%的提升。 */
func (this *TraderatesGetRequest) SetUseHasNext(value string) {
	this.Set("use_has_next", value)
}

type TraderatesGetResponse struct {
	*ErrorResponse              `json:"error_response,omitempty"`
	TraderatesGetResponseResult `json:"traderates_get_response"`
}
type TraderatesGetResponseResult struct {
	/* 当使用use_has_next时返回信息，如果还有下一页则返回true */
	HasNext bool `json:"has_next"`
	/* 搜索到的评价总数。相同的查询时间段条件下，最大只能获取总共1500条评价记录。所以当评价大于等于1500时 ISV可以通过start_date及end_date来进行拆分，以保证可以查询到全部数据 */
	TotalResults int64 `json:"total_results"`
	/* 评价列表。返回的TradeRate包含的具体信息为入参fields请求的字段信息 */
	TradeRates struct {
		TradeRate []*TradeRate `json:"trade_rate"`
	} `json:"TradeRates"`
}

/*tmall.traderate.feeds.get*/
type TmallTraderateFeedsGetRequest struct {
	values url.Values
}

func (this *TmallTraderateFeedsGetRequest) GetApiMethodName() string {
	return "tmall.traderate.feeds.get"
}
func (this *TmallTraderateFeedsGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *TmallTraderateFeedsGetRequest) GetValues() url.Values {
	return this.values
}

/* 交易子订单ID */
func (this *TmallTraderateFeedsGetRequest) SetChildTradeId(value string) {
	this.Set("child_trade_id", value)
}

type TmallTraderateFeedsGetResponse struct {
	*ErrorResponse                       `json:"error_response,omitempty"`
	TmallTraderateFeedsGetResponseResult `json:"tmall_traderate_feeds_get_response"`
}
type TmallTraderateFeedsGetResponseResult struct {
	/* 返回评价信息 */
	TmallRateInfo *TmallRateInfo `json:"tmall_rate_info"`
}

/*tmall.traderate.itemtags.get*/
type TmallTraderateItemtagsGetRequest struct {
	values url.Values
}

func (this *TmallTraderateItemtagsGetRequest) GetApiMethodName() string {
	return "tmall.traderate.itemtags.get"
}
func (this *TmallTraderateItemtagsGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *TmallTraderateItemtagsGetRequest) GetValues() url.Values {
	return this.values
}

/* 商品ID */
func (this *TmallTraderateItemtagsGetRequest) SetItemId(value string) {
	this.Set("item_id", value)
}

type TmallTraderateItemtagsGetResponse struct {
	*ErrorResponse                          `json:"error_response,omitempty"`
	TmallTraderateItemtagsGetResponseResult `json:"tmall_traderate_itemtags_get_response"`
}
type TmallTraderateItemtagsGetResponseResult struct {
	/* 标签列表 */
	Tags struct {
		TmallRateTagDetail []*TmallRateTagDetail `json:"tmall_rate_tag_detail"`
	} `json:"Tags"`
}

/*taobao.areas.get*/
type AreasGetRequest struct {
	values url.Values
}

func (this *AreasGetRequest) GetApiMethodName() string {
	return "taobao.areas.get"
}
func (this *AreasGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *AreasGetRequest) GetValues() url.Values {
	return this.values
}

/* 需返回的字段列表.可选值:Area 结构中的所有字段;多个字段之间用","分隔.如:id,type,name,parent_id,zip. */
func (this *AreasGetRequest) SetFields(value string) {
	this.Set("fields", value)
}

type AreasGetResponse struct {
	*ErrorResponse         `json:"error_response,omitempty"`
	AreasGetResponseResult `json:"areas_get_response"`
}
type AreasGetResponseResult struct {
	/* 地址区域信息列表.返回的Area包含的具体信息为入参fields请求的字段信息. */
	Areas struct {
		Area []*Area `json:"area"`
	} `json:"Areas"`
}

/*taobao.delivery.template.add*/
type DeliveryTemplateAddRequest struct {
	values url.Values
}

func (this *DeliveryTemplateAddRequest) GetApiMethodName() string {
	return "taobao.delivery.template.add"
}
func (this *DeliveryTemplateAddRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *DeliveryTemplateAddRequest) GetValues() url.Values {
	return this.values
}

/* 可选值：0、1 ，说明如下<br>0:表示买家承担服务费;<br>1:表示卖家承担服务费 */
func (this *DeliveryTemplateAddRequest) SetAssumer(value string) {
	this.Set("assumer", value)
}

/* 卖家发货地址区域ID
<br/><br/><font color=blue>可以不填，如果没有填写取卖家默认发货地址的区域ID，如果需要输入必须用taobao.areas.get接口获取.或者参考：http://www.stats.gov.cn/tjbz/xzqhdm/t20080215_402462675.htm
</font>

<br/><br/><font color=red>注意：填入该值时必须取您的发货地址最小区域级别ID，比如您的发货地址是：某省XX市xx区（县）时需要填入区(县)的ID，当然有些地方没有区或县可以直接填市级别的ID</font> */
func (this *DeliveryTemplateAddRequest) SetConsignAreaId(value string) {
	this.Set("consign_area_id", value)
}

/* 运费模板的名称，长度不能超过50个字节<br /> 支持最大长度为：50<br /> 支持的最大列表长度为：50 */
func (this *DeliveryTemplateAddRequest) SetName(value string) {
	this.Set("name", value)
}

/* 增费：输入0.00-999.99（最多包含两位小数）

<br/><br/><font color=blue>增费必须小于等于首费，但是当首费为0时增费可以大于首费</font>


<br/><br/><font color=red>输入的格式分号个数和template_types数量一致，逗号个数必须与template_dests以分号隔开之后一一对应的数量一致</font> */
func (this *DeliveryTemplateAddRequest) SetTemplateAddFees(value string) {
	this.Set("template_add_fees", value)
}

/* 增费标准：当valuation(记价方式)为0时输入1-9999范围内的整数<br/><font color=red>当valuation(记价方式)为1时输入0.1-9999.9范围内的小数只能包含以为小数（单位为千克）<br/><font color=blue>当valuation(记价方式)为3时输入0.1-999.9范围内的数值，数值只能包含一位小数（单位为 立方米）
<br/>
<br><font color=red>输入的格式分号个数和template_types数量一致，逗号个数必须与template_dests以分号隔开之后一一对应的数量一致</font> */
func (this *DeliveryTemplateAddRequest) SetTemplateAddStandards(value string) {
	this.Set("template_add_standards", value)
}

/* 邮费子项涉及的地区.结构: value1;value2;value3,value4
<br>如:1,110000;1,110000;1,310000;1,320000,330000。 aredId解释(1=全国,110000=北京,310000=上海,320000=江苏,330000=浙江)
如果template_types设置为post;ems;exmpress;cod则表示为平邮(post)指定默认地区(全国)和北京地区的运费;其他的类似以分号区分一一对应
<br/>可以用taobao.areas.get接口获取.或者参考：http://www.stats.gov.cn/tjbz/xzqhdm/t20080215_402462675.htm<br/>
<br/><font color=red>每个运费方式设置涉及的地区中必须包含全国地区（areaId=1）表示默认运费,可以只设置默认运费</font>
<br><font color=blue>注意:为多个地区指定指定不同（首费标准、首费、增费标准、增费一项不一样就算不同）的运费以逗号","区分，
template_start_standards(首费标准)、template_start_fees(首费)、
template_add_standards(增费标准)、
template_add_fees(增费)必须与template_types分号数量相同。如果为需要为多个地区指定相同运费则地区之间用“|”隔开即可。</font>
<font color=red>如果卖家没有传入发货地址则默认地址库的发货地址</font> */
func (this *DeliveryTemplateAddRequest) SetTemplateDests(value string) {
	this.Set("template_dests", value)
}

/* 首费：输入0.00-999.99（最多包含两位小数）
<br/><font color=red>输入的格式分号个数和template_types数量一致，逗号个数必须与template_dests以分号隔开之后一一对应的数量一致</font> */
func (this *DeliveryTemplateAddRequest) SetTemplateStartFees(value string) {
	this.Set("template_start_fees", value)
}

/* 首费标准：当valuation(记价方式)为0时输入1-9999范围内的整数<br/><font color=red>当valuation(记价方式)为1时输入0.1-9999.9范围内的小数只能包含以为小数（单位为千克）<br/><font color=blue>当valuation(记价方式)为3时输入0.1-999.9范围内的数值，数值只能包含一位小数（单位为 立方米）
<br/>
<font color=red>输入的格式分号个数和template_types数量一致，逗号个数必须与template_dests以分号隔开之后一一对应的数量一致</font> */
func (this *DeliveryTemplateAddRequest) SetTemplateStartStandards(value string) {
	this.Set("template_start_standards", value)
}

/* 运费方式:平邮 (post),快递公司(express),EMS (ems),货到付款(cod),物流宝保障速递(wlb),家装物流(furniture)结构:value1;value2;value3;value4
如: post;express;ems;cod
<br/><font color=red>
注意:在添加多个运费方式时,字符串中使用 ";" 分号区分
。template_dests(指定地区)
template_start_standards(首费标准)、template_start_fees(首费)、template_add_standards(增费标准)、template_add_fees(增费)必须与template_types的分号数量相同. </font>
<br>
<font color=blue>
注意：<br/>
1、post,ems,express三种运费方式必须填写一个<br/>
2、只有订购了货到付款用户可以填cod;只有家装物流用户可以填写furniture
只有订购了保障速递的用户可以填写bzsd,只有物流宝用户可以填写wlb<br/>
3、如果是货到付款用户当没有填写cod运送方式时，运费模板会默认继承express的费用为cod的费用<br/>
4、如果是保障速递户当没有填写bzsd运送方式时，运费模板会默认继承express的费用为bzsd的费用<br/>
5、由于3和4的原因所以当是货到付款用户或保障速递用户时如果没填写对应的运送方式express是必须填写的
</font> */
func (this *DeliveryTemplateAddRequest) SetTemplateTypes(value string) {
	this.Set("template_types", value)
}

/* 可选值：0、1、3，说明如下。<br>0:表示按宝贝件数计算运费 <br>1:表示按宝贝重量计算运费
<br>3:表示按宝贝体积计算运费 */
func (this *DeliveryTemplateAddRequest) SetValuation(value string) {
	this.Set("valuation", value)
}

type DeliveryTemplateAddResponse struct {
	*ErrorResponse                    `json:"error_response,omitempty"`
	DeliveryTemplateAddResponseResult `json:"delivery_template_add_response"`
}
type DeliveryTemplateAddResponseResult struct {
	/* 模板对象 */
	DeliveryTemplate *DeliveryTemplate `json:"delivery_template"`
}

/*taobao.delivery.template.delete*/
type DeliveryTemplateDeleteRequest struct {
	values url.Values
}

func (this *DeliveryTemplateDeleteRequest) GetApiMethodName() string {
	return "taobao.delivery.template.delete"
}
func (this *DeliveryTemplateDeleteRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *DeliveryTemplateDeleteRequest) GetValues() url.Values {
	return this.values
}

/* 运费模板ID */
func (this *DeliveryTemplateDeleteRequest) SetTemplateId(value string) {
	this.Set("template_id", value)
}

type DeliveryTemplateDeleteResponse struct {
	*ErrorResponse                       `json:"error_response,omitempty"`
	DeliveryTemplateDeleteResponseResult `json:"delivery_template_delete_response"`
}
type DeliveryTemplateDeleteResponseResult struct {
	/* 表示删除成功还是失败 */
	Complete bool `json:"complete"`
}

/*taobao.delivery.template.get*/
type DeliveryTemplateGetRequest struct {
	values url.Values
}

func (this *DeliveryTemplateGetRequest) GetApiMethodName() string {
	return "taobao.delivery.template.get"
}
func (this *DeliveryTemplateGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *DeliveryTemplateGetRequest) GetValues() url.Values {
	return this.values
}

/* 需返回的字段列表。 <br/>
<B>
可选值:示例中的值;各字段之间用","隔开
</B>
<br/><br/>
<font color=blue>
template_id：查询模板ID <br/>
template_name:查询模板名称 <br/>
assumer：查询服务承担放 <br/>
valuation:查询计价规则 <br/>
supports:查询增值服务列表 <br/>
created:查询模板创建时间 <br/>
modified:查询修改时间<br/>
consign_area_id:运费模板上设置的卖家发货地址最小级别ID<br/>
address:卖家地址信息
</font>
<br/><br/>
<font color=bule>
query_cod:查询货到付款运费信息<br/>
query_post:查询平邮运费信息 <br/>
query_express:查询快递公司运费信息 <br/>
query_ems:查询EMS运费信息<br/>
query_bzsd:查询普通保障速递运费信息<br/>
query_wlb:查询物流宝保障速递运费信息<br/>
query_furniture:查询家装物流运费信息
<font color=blue>
<br/><br/>
<font color=red>不能有空格</font> */
func (this *DeliveryTemplateGetRequest) SetFields(value string) {
	this.Set("fields", value)
}

/* 需要查询的运费模板ID列表 */
func (this *DeliveryTemplateGetRequest) SetTemplateIds(value string) {
	this.Set("template_ids", value)
}

/* 在没有登录的情况下根据淘宝用户昵称查询指定的模板 */
func (this *DeliveryTemplateGetRequest) SetUserNick(value string) {
	this.Set("user_nick", value)
}

type DeliveryTemplateGetResponse struct {
	*ErrorResponse                    `json:"error_response,omitempty"`
	DeliveryTemplateGetResponseResult `json:"delivery_template_get_response"`
}
type DeliveryTemplateGetResponseResult struct {
	/* 运费模板列表 */
	DeliveryTemplates struct {
		DeliveryTemplate []*DeliveryTemplate `json:"delivery_template"`
	} `json:"DeliveryTemplates"`
	/* 获得到符合条件的结果总数 */
	TotalResults int64 `json:"total_results"`
}

/*taobao.delivery.template.update*/
type DeliveryTemplateUpdateRequest struct {
	values url.Values
}

func (this *DeliveryTemplateUpdateRequest) GetApiMethodName() string {
	return "taobao.delivery.template.update"
}
func (this *DeliveryTemplateUpdateRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *DeliveryTemplateUpdateRequest) GetValues() url.Values {
	return this.values
}

/* 可选值：0,1 <br>  说明<br>0:表示买家承担服务费;<br>1:表示卖家承担服务费 */
func (this *DeliveryTemplateUpdateRequest) SetAssumer(value string) {
	this.Set("assumer", value)
}

/* 模板名称，长度不能大于50个字节 */
func (this *DeliveryTemplateUpdateRequest) SetName(value string) {
	this.Set("name", value)
}

/* 增费：输入0.00-999.99（最多包含两位小数）<br/><font color=blue>增费可以为0</font><br/><font color=red>输入的格式分号个数和template_types数量一致，逗号个数必须与template_dests以分号隔开之后一一对应的数量一致</font> */
func (this *DeliveryTemplateUpdateRequest) SetTemplateAddFees(value string) {
	this.Set("template_add_fees", value)
}

/* 增费标准：当valuation(记价方式)为0时输入1-9999范围内的整数<br><font color=blue>增费标准目前只能为1</font>
<br><font color=red>输入的格式分号个数和template_types数量一致，逗号个数必须与template_dests以分号隔开之后一一对应的数量一致</font> */
func (this *DeliveryTemplateUpdateRequest) SetTemplateAddStandards(value string) {
	this.Set("template_add_standards", value)
}

/* 邮费子项涉及的地区.结构: value1;value2;value3,value4
<br>如:1,110000;1,110000;1,310000;1,320000,330000。 aredId解释(1=全国,110000=北京,310000=上海,320000=江苏,330000=浙江)
如果template_types设置为post;ems;exmpress;cod则表示为平邮(post)指定默认地区(全国)和北京地区的运费;其他的类似以分号区分一一对应
<br/>可以用taobao.areas.get接口获取.或者参考：http://www.stats.gov.cn/tjbz/xzqhdm/t20080215_402462675.htm
<br/><font color=red>每个运费方式设置的设涉及地区中必须包含全国地区（areaId=1）表示默认运费,可以只设置默认运费</font>
<br><font color=blue>注意:为多个地区指定指定不同（首费标准、首费、增费标准、增费一项不一样就算不同）的运费以逗号","区分，
template_start_standards(首费标准)、template_start_fees(首费)、
template_add_standards(增费标准)、
template_add_fees(增费)必须与template_types分号数量相同。如果为需要为多个地区指定相同运费则地区之间用“|”隔开即可。</font> */
func (this *DeliveryTemplateUpdateRequest) SetTemplateDests(value string) {
	this.Set("template_dests", value)
}

/* 需要修改的模板对应的模板ID */
func (this *DeliveryTemplateUpdateRequest) SetTemplateId(value string) {
	this.Set("template_id", value)
}

/* 首费：输入0.01-999.99（最多包含两位小数）
<br/><font color=blue> 首费不能为0</font><br><font color=red>输入的格式分号个数和template_types数量一致，逗号个数必须与template_dests以分号隔开之后一一对应的数量一致</font> */
func (this *DeliveryTemplateUpdateRequest) SetTemplateStartFees(value string) {
	this.Set("template_start_fees", value)
}

/* 首费标准：当valuation(记价方式)为0时输入1-9999范围内的整数<br><font color=blue>首费标准目前只能为1</br>
<br><font color=red>输入的格式分号个数和template_types数量一致，逗号个数必须与template_dests以分号隔开之后一一对应的数量一致</font> */
func (this *DeliveryTemplateUpdateRequest) SetTemplateStartStandards(value string) {
	this.Set("template_start_standards", value)
}

/* 运费方式:平邮 (post),快递公司(express),EMS (ems),货到付款(cod)结构:value1;value2;value3;value4
如: post;express;ems;cod
<br/><font color=red>
注意:在添加多个运费方式时,字符串中使用 ";" 分号区分。template_dests(指定地区) template_start_standards(首费标准)、template_start_fees(首费)、template_add_standards(增费标准)、template_add_fees(增费)必须与template_types的分号数量相同.
 </font>
<br/>
<font color=blue>
普通用户：post,ems,express三种运费方式必须填写一个，不能填写cod。
货到付款用户：如果填写了cod运费方式，则post,ems,express三种运费方式也必须填写一个，如果没有填写cod则填写的运费方式中必须存在express</font> */
func (this *DeliveryTemplateUpdateRequest) SetTemplateTypes(value string) {
	this.Set("template_types", value)
}

type DeliveryTemplateUpdateResponse struct {
	*ErrorResponse                       `json:"error_response,omitempty"`
	DeliveryTemplateUpdateResponseResult `json:"delivery_template_update_response"`
}
type DeliveryTemplateUpdateResponseResult struct {
	/* 表示修改是否成功 */
	Complete bool `json:"complete"`
}

/*taobao.delivery.templates.get*/
type DeliveryTemplatesGetRequest struct {
	values url.Values
}

func (this *DeliveryTemplatesGetRequest) GetApiMethodName() string {
	return "taobao.delivery.templates.get"
}
func (this *DeliveryTemplatesGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *DeliveryTemplatesGetRequest) GetValues() url.Values {
	return this.values
}

/* 需返回的字段列表。 <br/>
<B>
可选值:示例中的值;各字段之间用","隔开
</B>
<br/><br/>
<font color=blue>
template_id：查询模板ID <br/>
template_name:查询模板名称 <br/>
assumer：查询服务承担放 <br/>
valuation:查询计价规则 <br/>
supports:查询增值服务列表 <br/>
created:查询模板创建时间 <br/>
modified:查询修改时间<br/>
consign_area_id:运费模板上设置的卖家发货地址最小级别ID<br/>
address:卖家地址信息
</font>
<br/><br/>
<font color=bule>
query_cod:查询货到付款运费信息<br/>
query_post:查询平邮运费信息 <br/>
query_express:查询快递公司运费信息 <br/>
query_ems:查询EMS运费信息<br/>
query_bzsd:查询普通保障速递运费信息<br/>
query_wlb:查询物流宝保障速递运费信息<br/>
query_furniture:查询家装物流运费信息
<font color=blue>
<br/><br/>
<font color=red>不能有空格</font> */
func (this *DeliveryTemplatesGetRequest) SetFields(value string) {
	this.Set("fields", value)
}

type DeliveryTemplatesGetResponse struct {
	*ErrorResponse                     `json:"error_response,omitempty"`
	DeliveryTemplatesGetResponseResult `json:"delivery_templates_get_response"`
}
type DeliveryTemplatesGetResponseResult struct {
	/* 运费模板列表 */
	DeliveryTemplates struct {
		DeliveryTemplate []*DeliveryTemplate `json:"delivery_template"`
	} `json:"DeliveryTemplates"`
	/* 获得到符合条件的结果总数 */
	TotalResults int64 `json:"total_results"`
}

/*taobao.logistics.address.add*/
type LogisticsAddressAddRequest struct {
	values url.Values
}

func (this *LogisticsAddressAddRequest) GetApiMethodName() string {
	return "taobao.logistics.address.add"
}
func (this *LogisticsAddressAddRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *LogisticsAddressAddRequest) GetValues() url.Values {
	return this.values
}

/* 详细街道地址，不需要重复填写省/市/区 */
func (this *LogisticsAddressAddRequest) SetAddr(value string) {
	this.Set("addr", value)
}

/* 默认退货地址。<br>
<font color='red'>选择此项(true)，将当前地址设为默认退货地址，撤消原默认退货地址</font> */
func (this *LogisticsAddressAddRequest) SetCancelDef(value string) {
	this.Set("cancel_def", value)
}

/* 所在市 */
func (this *LogisticsAddressAddRequest) SetCity(value string) {
	this.Set("city", value)
}

/* 联系人姓名 <font color='red'>长度不可超过20个字节</font> */
func (this *LogisticsAddressAddRequest) SetContactName(value string) {
	this.Set("contact_name", value)
}

/* 区、县
<br><font color='red'>如果所在地区是海外的可以为空，否则为必参</font> */
func (this *LogisticsAddressAddRequest) SetCountry(value string) {
	this.Set("country", value)
}

/* 默认取货地址。<br>
<font color='red'>选择此项(true)，将当前地址设为默认取货地址，撤消原默认取货地址</font> */
func (this *LogisticsAddressAddRequest) SetGetDef(value string) {
	this.Set("get_def", value)
}

/* 备注,<br><font color='red'>备注不能超过256字节</font><br /> 支持最大长度为：256<br /> 支持的最大列表长度为：256 */
func (this *LogisticsAddressAddRequest) SetMemo(value string) {
	this.Set("memo", value)
}

/* 手机号码，手机与电话必需有一个
<br><font color='red'>手机号码不能超过20位</font> */
func (this *LogisticsAddressAddRequest) SetMobilePhone(value string) {
	this.Set("mobile_phone", value)
}

/* 电话号码,手机与电话必需有一个 */
func (this *LogisticsAddressAddRequest) SetPhone(value string) {
	this.Set("phone", value)
}

/* 所在省 */
func (this *LogisticsAddressAddRequest) SetProvince(value string) {
	this.Set("province", value)
}

/* 公司名称,<br><font color="red">公司名称长能不能超过20字节</font><br /> 支持最大长度为：20<br /> 支持的最大列表长度为：20 */
func (this *LogisticsAddressAddRequest) SetSellerCompany(value string) {
	this.Set("seller_company", value)
}

/* 地区邮政编码
<br><font color='red'>如果所在地区是海外的可以为空，否则为必参</font> */
func (this *LogisticsAddressAddRequest) SetZipCode(value string) {
	this.Set("zip_code", value)
}

type LogisticsAddressAddResponse struct {
	*ErrorResponse                    `json:"error_response,omitempty"`
	LogisticsAddressAddResponseResult `json:"logistics_address_add_response"`
}
type LogisticsAddressAddResponseResult struct {
	/* 只返回修改日期modify_date */
	AddressResult *AddressResult `json:"address_result"`
}

/*taobao.logistics.address.modify*/
type LogisticsAddressModifyRequest struct {
	values url.Values
}

func (this *LogisticsAddressModifyRequest) GetApiMethodName() string {
	return "taobao.logistics.address.modify"
}
func (this *LogisticsAddressModifyRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *LogisticsAddressModifyRequest) GetValues() url.Values {
	return this.values
}

/* 详细街道地址，不需要重复填写省/市/区 */
func (this *LogisticsAddressModifyRequest) SetAddr(value string) {
	this.Set("addr", value)
}

/* 默认退货地址。<br>
<font color='red'>选择此项(true)，将当前地址设为默认退货地址，撤消原默认退货地址</font> */
func (this *LogisticsAddressModifyRequest) SetCancelDef(value string) {
	this.Set("cancel_def", value)
}

/* 所在市 */
func (this *LogisticsAddressModifyRequest) SetCity(value string) {
	this.Set("city", value)
}

/* 地址库ID */
func (this *LogisticsAddressModifyRequest) SetContactId(value string) {
	this.Set("contact_id", value)
}

/* 联系人姓名
<font color='red'>长度不可超过20个字节</font> */
func (this *LogisticsAddressModifyRequest) SetContactName(value string) {
	this.Set("contact_name", value)
}

/* 区、县
<br><font color='red'>如果所在地区是海外的可以为空，否则为必参</font> */
func (this *LogisticsAddressModifyRequest) SetCountry(value string) {
	this.Set("country", value)
}

/* 默认取货地址。<br>
<font color='red'>选择此项(true)，将当前地址设为默认取货地址，撤消原默认取货地址</font> */
func (this *LogisticsAddressModifyRequest) SetGetDef(value string) {
	this.Set("get_def", value)
}

/* 备注,<br><font color='red'>备注不能超过256字节</font> */
func (this *LogisticsAddressModifyRequest) SetMemo(value string) {
	this.Set("memo", value)
}

/* 手机号码，手机与电话必需有一个 <br><font color='red'>手机号码不能超过20位</font> */
func (this *LogisticsAddressModifyRequest) SetMobilePhone(value string) {
	this.Set("mobile_phone", value)
}

/* 电话号码,手机与电话必需有一个 */
func (this *LogisticsAddressModifyRequest) SetPhone(value string) {
	this.Set("phone", value)
}

/* 所在省 */
func (this *LogisticsAddressModifyRequest) SetProvince(value string) {
	this.Set("province", value)
}

/* 公司名称,
<br><font color='red'>公司名称长能不能超过20字节</font> */
func (this *LogisticsAddressModifyRequest) SetSellerCompany(value string) {
	this.Set("seller_company", value)
}

/* 地区邮政编码
<br><font color='red'>如果所在地区是海外的可以为空，否则为必参</font> */
func (this *LogisticsAddressModifyRequest) SetZipCode(value string) {
	this.Set("zip_code", value)
}

type LogisticsAddressModifyResponse struct {
	*ErrorResponse                       `json:"error_response,omitempty"`
	LogisticsAddressModifyResponseResult `json:"logistics_address_modify_response"`
}
type LogisticsAddressModifyResponseResult struct {
	/* 只返回修改时间modify_date */
	AddressResult *AddressResult `json:"address_result"`
}

/*taobao.logistics.address.remove*/
type LogisticsAddressRemoveRequest struct {
	values url.Values
}

func (this *LogisticsAddressRemoveRequest) GetApiMethodName() string {
	return "taobao.logistics.address.remove"
}
func (this *LogisticsAddressRemoveRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *LogisticsAddressRemoveRequest) GetValues() url.Values {
	return this.values
}

/* 地址库ID */
func (this *LogisticsAddressRemoveRequest) SetContactId(value string) {
	this.Set("contact_id", value)
}

type LogisticsAddressRemoveResponse struct {
	*ErrorResponse                       `json:"error_response,omitempty"`
	LogisticsAddressRemoveResponseResult `json:"logistics_address_remove_response"`
}
type LogisticsAddressRemoveResponseResult struct {
	/* 只返回修改日期modify_date */
	AddressResult *AddressResult `json:"address_result"`
}

/*taobao.logistics.address.search*/
type LogisticsAddressSearchRequest struct {
	values url.Values
}

func (this *LogisticsAddressSearchRequest) GetApiMethodName() string {
	return "taobao.logistics.address.search"
}
func (this *LogisticsAddressSearchRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *LogisticsAddressSearchRequest) GetValues() url.Values {
	return this.values
}

/* 可选，参数列表如下：<br><font color='red'>
no_def:查询非默认地址<br>
get_def:查询默认取货地址<br>
cancel_def:查询默认退货地址<br>
缺省此参数时，查询所有当前用户的地址库
</font> */
func (this *LogisticsAddressSearchRequest) SetRdef(value string) {
	this.Set("rdef", value)
}

type LogisticsAddressSearchResponse struct {
	*ErrorResponse                       `json:"error_response,omitempty"`
	LogisticsAddressSearchResponseResult `json:"logistics_address_search_response"`
}
type LogisticsAddressSearchResponseResult struct {
	/* 一组地址库数据， */
	Addresses struct {
		AddressResult []*AddressResult `json:"address_result"`
	} `json:"Addresses"`
}

/*taobao.logistics.companies.get*/
type LogisticsCompaniesGetRequest struct {
	values url.Values
}

func (this *LogisticsCompaniesGetRequest) GetApiMethodName() string {
	return "taobao.logistics.companies.get"
}
func (this *LogisticsCompaniesGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *LogisticsCompaniesGetRequest) GetValues() url.Values {
	return this.values
}

/* 需返回的字段列表。可选值:LogisticCompany 结构中的所有字段;多个字段间用","逗号隔开.
如:id,code,name,reg_mail_no
<br><font color='red'>说明：</font>
<br>id：物流公司ID
<br>code：物流公司code
<br>name：物流公司名称
<br>reg_mail_no：物流公司对应的运单规则 */
func (this *LogisticsCompaniesGetRequest) SetFields(value string) {
	this.Set("fields", value)
}

/* 是否查询推荐物流公司.可选值:true,false.如果不提供此参数,将会返回所有支持电话联系的物流公司. */
func (this *LogisticsCompaniesGetRequest) SetIsRecommended(value string) {
	this.Set("is_recommended", value)
}

/* 推荐物流公司的下单方式.可选值:offline(电话联系/自己联系),online(在线下单),all(即电话联系又在线下单). 此参数仅仅用于is_recommended 为ture时。就是说对于推荐物流公司才可用.如果不选择此参数将会返回推荐物流中支持电话联系的物流公司. */
func (this *LogisticsCompaniesGetRequest) SetOrderMode(value string) {
	this.Set("order_mode", value)
}

type LogisticsCompaniesGetResponse struct {
	*ErrorResponse                      `json:"error_response,omitempty"`
	LogisticsCompaniesGetResponseResult `json:"logistics_companies_get_response"`
}
type LogisticsCompaniesGetResponseResult struct {
	/* 物流公司信息。返回的LogisticCompany包含的具体信息为入参fields请求的字段信息。 */
	LogisticsCompanies struct {
		LogisticsCompany []*LogisticsCompany `json:"logistics_company"`
	} `json:"LogisticsCompanies"`
}

/*taobao.logistics.consign.order.createandsend*/
type LogisticsConsignOrderCreateandsendRequest struct {
	values url.Values
}

func (this *LogisticsConsignOrderCreateandsendRequest) GetApiMethodName() string {
	return "taobao.logistics.consign.order.createandsend"
}
func (this *LogisticsConsignOrderCreateandsendRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *LogisticsConsignOrderCreateandsendRequest) GetValues() url.Values {
	return this.values
}

/* 物流公司ID */
func (this *LogisticsConsignOrderCreateandsendRequest) SetCompanyId(value string) {
	this.Set("company_id", value)
}

/* 物品的json数据。 */
func (this *LogisticsConsignOrderCreateandsendRequest) SetItemJsonString(value string) {
	this.Set("item_json_string", value)
}

/* 物流订单物流类型，值固定选择：2 */
func (this *LogisticsConsignOrderCreateandsendRequest) SetLogisType(value string) {
	this.Set("logis_type", value)
}

/* 运单号 */
func (this *LogisticsConsignOrderCreateandsendRequest) SetMailNo(value string) {
	this.Set("mail_no", value)
}

/* 订单来源，值选择：30 */
func (this *LogisticsConsignOrderCreateandsendRequest) SetOrderSource(value string) {
	this.Set("order_source", value)
}

/* 订单类型，值固定选择：30 */
func (this *LogisticsConsignOrderCreateandsendRequest) SetOrderType(value string) {
	this.Set("order_type", value)
}

/* 收件人街道地址 */
func (this *LogisticsConsignOrderCreateandsendRequest) SetRAddress(value string) {
	this.Set("r_address", value)
}

/* 收件人区域ID */
func (this *LogisticsConsignOrderCreateandsendRequest) SetRAreaId(value string) {
	this.Set("r_area_id", value)
}

/* 市 */
func (this *LogisticsConsignOrderCreateandsendRequest) SetRCityName(value string) {
	this.Set("r_city_name", value)
}

/* 区 */
func (this *LogisticsConsignOrderCreateandsendRequest) SetRDistName(value string) {
	this.Set("r_dist_name", value)
}

/* 手机号码 */
func (this *LogisticsConsignOrderCreateandsendRequest) SetRMobilePhone(value string) {
	this.Set("r_mobile_phone", value)
}

/* 收件人名称 */
func (this *LogisticsConsignOrderCreateandsendRequest) SetRName(value string) {
	this.Set("r_name", value)
}

/* 省 */
func (this *LogisticsConsignOrderCreateandsendRequest) SetRProvName(value string) {
	this.Set("r_prov_name", value)
}

/* 电话号码 */
func (this *LogisticsConsignOrderCreateandsendRequest) SetRTelephone(value string) {
	this.Set("r_telephone", value)
}

/* 收件人邮编 */
func (this *LogisticsConsignOrderCreateandsendRequest) SetRZipCode(value string) {
	this.Set("r_zip_code", value)
}

/* 发件人街道地址 */
func (this *LogisticsConsignOrderCreateandsendRequest) SetSAddress(value string) {
	this.Set("s_address", value)
}

/* 发件人区域ID */
func (this *LogisticsConsignOrderCreateandsendRequest) SetSAreaId(value string) {
	this.Set("s_area_id", value)
}

/* 市 */
func (this *LogisticsConsignOrderCreateandsendRequest) SetSCityName(value string) {
	this.Set("s_city_name", value)
}

/* 区 */
func (this *LogisticsConsignOrderCreateandsendRequest) SetSDistName(value string) {
	this.Set("s_dist_name", value)
}

/* 手机号码 */
func (this *LogisticsConsignOrderCreateandsendRequest) SetSMobilePhone(value string) {
	this.Set("s_mobile_phone", value)
}

/* 发件人名称 */
func (this *LogisticsConsignOrderCreateandsendRequest) SetSName(value string) {
	this.Set("s_name", value)
}

/* 省 */
func (this *LogisticsConsignOrderCreateandsendRequest) SetSProvName(value string) {
	this.Set("s_prov_name", value)
}

/* 电话号码 */
func (this *LogisticsConsignOrderCreateandsendRequest) SetSTelephone(value string) {
	this.Set("s_telephone", value)
}

/* 发件人出编 */
func (this *LogisticsConsignOrderCreateandsendRequest) SetSZipCode(value string) {
	this.Set("s_zip_code", value)
}

/* 费用承担方式 1买家承担运费 2卖家承担运费 */
func (this *LogisticsConsignOrderCreateandsendRequest) SetShipping(value string) {
	this.Set("shipping", value)
}

/* 交易流水号，淘外订单号或者商家内部交易流水号 */
func (this *LogisticsConsignOrderCreateandsendRequest) SetTradeId(value string) {
	this.Set("trade_id", value)
}

/* 用户ID */
func (this *LogisticsConsignOrderCreateandsendRequest) SetUserId(value string) {
	this.Set("user_id", value)
}

type LogisticsConsignOrderCreateandsendResponse struct {
	*ErrorResponse                                   `json:"error_response,omitempty"`
	LogisticsConsignOrderCreateandsendResponseResult `json:"logistics_consign_order_createandsend_response"`
}
type LogisticsConsignOrderCreateandsendResponseResult struct {
	/* 是否成功 */
	IsSuccess bool `json:"is_success"`
	/* 订单ID */
	OrderId int64 `json:"order_id"`
	/* 结果描述 */
	ResultDesc string `json:"result_desc"`
}

/*taobao.logistics.consign.resend*/
type LogisticsConsignResendRequest struct {
	values url.Values
}

func (this *LogisticsConsignResendRequest) GetApiMethodName() string {
	return "taobao.logistics.consign.resend"
}
func (this *LogisticsConsignResendRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *LogisticsConsignResendRequest) GetValues() url.Values {
	return this.values
}

/* 物流公司代码.如"POST"代表中国邮政,"ZJS"代表宅急送。调用 taobao.logistics.companies.get 获取。
<br><font color='red'>如果是货到付款订单，选择的物流公司必须支持货到付款发货方式</font> */
func (this *LogisticsConsignResendRequest) SetCompanyCode(value string) {
	this.Set("company_code", value)
}

/* feature参数格式<br>
范例: identCode=tid1:识别码1,识别码2|tid2:识别码3;machineCode=tid3:3C机器号A,3C机器号B<br>
identCode为识别码的KEY,machineCode为3C的KEY,多个key之间用”;”分隔<br>
“tid1:识别码1,识别码2|tid2:识别码3”为identCode对应的value。
"|"不同商品间的分隔符。<br>
例1商品和2商品，之间就用"|"分开。<br>
TID就是商品代表的子订单号，对应taobao.trade.fullinfo.get 接口获得的oid字段。(通过OID可以唯一定位到当前商品上)<br>
":"TID和具体传入参数间的分隔符。冒号前表示TID,之后代表该商品的参数属性。<br>
"," 属性间分隔符。（对应商品数量，当存在一个商品的数量超过1个时，用逗号分开）。<br>
具体:当订单中A商品的数量为2个，其中手机串号分别为"12345","67890"。<br>
参数格式：identCode=TIDA:12345,67890。
TIDA对应了A宝贝，冒号后用逗号分隔的"12345","67890".说明本订单A宝贝的数量为2，值分别为"12345","67890"。<br>
当存在"|"时，就说明订单中存在多个商品，商品间用"|"分隔了开来。|"之后的内容含义同上。 */
func (this *LogisticsConsignResendRequest) SetFeature(value string) {
	this.Set("feature", value)
}

/* 表明是否是拆单，默认值0，1表示拆单 */
func (this *LogisticsConsignResendRequest) SetIsSplit(value string) {
	this.Set("is_split", value)
}

/* 运单号.具体一个物流公司的真实运单号码。淘宝官方物流会校验，请谨慎传入； */
func (this *LogisticsConsignResendRequest) SetOutSid(value string) {
	this.Set("out_sid", value)
}

/* 商家的IP地址 */
func (this *LogisticsConsignResendRequest) SetSellerIp(value string) {
	this.Set("seller_ip", value)
}

/* 拆单子订单列表，对应的数据是：子订单号列表。可以不传，但是如果传了则必须符合传递的规则。子订单必须是操作的物流订单的子订单的真子集！ */
func (this *LogisticsConsignResendRequest) SetSubTid(value string) {
	this.Set("sub_tid", value)
}

/* 淘宝交易ID<br /> 支持最小值为：1000 */
func (this *LogisticsConsignResendRequest) SetTid(value string) {
	this.Set("tid", value)
}

type LogisticsConsignResendResponse struct {
	*ErrorResponse                       `json:"error_response,omitempty"`
	LogisticsConsignResendResponseResult `json:"logistics_consign_resend_response"`
}
type LogisticsConsignResendResponseResult struct {
	/* 返回发货是否成功is_success */
	Shipping *Shipping `json:"shipping"`
}

/*taobao.logistics.dummy.send*/
type LogisticsDummySendRequest struct {
	values url.Values
}

func (this *LogisticsDummySendRequest) GetApiMethodName() string {
	return "taobao.logistics.dummy.send"
}
func (this *LogisticsDummySendRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *LogisticsDummySendRequest) GetValues() url.Values {
	return this.values
}

/* feature参数格式<br>
范例: identCode=tid1:识别码1,识别码2|tid2:识别码3;machineCode=tid3:3C机器号A,3C机器号B<br>
identCode为识别码的KEY,machineCode为3C的KEY,多个key之间用”;”分隔<br>
“tid1:识别码1,识别码2|tid2:识别码3”为identCode对应的value。
"|"不同商品间的分隔符。<br>
例1商品和2商品，之间就用"|"分开。<br>
TID就是商品代表的子订单号，对应taobao.trade.fullinfo.get 接口获得的oid字段。(通过OID可以唯一定位到当前商品上)<br>
":"TID和具体传入参数间的分隔符。冒号前表示TID,之后代表该商品的参数属性。<br>
"," 属性间分隔符。（对应商品数量，当存在一个商品的数量超过1个时，用逗号分开）。<br>
具体:当订单中A商品的数量为2个，其中手机串号分别为"12345","67890"。<br>
参数格式：identCode=TIDA:12345,67890。
TIDA对应了A宝贝，冒号后用逗号分隔的"12345","67890".说明本订单A宝贝的数量为2，值分别为"12345","67890"。<br>
当存在"|"时，就说明订单中存在多个商品，商品间用"|"分隔了开来。|"之后的内容含义同上。 */
func (this *LogisticsDummySendRequest) SetFeature(value string) {
	this.Set("feature", value)
}

/* 商家的IP地址 */
func (this *LogisticsDummySendRequest) SetSellerIp(value string) {
	this.Set("seller_ip", value)
}

/* 淘宝交易ID */
func (this *LogisticsDummySendRequest) SetTid(value string) {
	this.Set("tid", value)
}

type LogisticsDummySendResponse struct {
	*ErrorResponse                   `json:"error_response,omitempty"`
	LogisticsDummySendResponseResult `json:"logistics_dummy_send_response"`
}
type LogisticsDummySendResponseResult struct {
	/* 返回发货是否成功is_success */
	Shipping *Shipping `json:"shipping"`
}

/*taobao.logistics.offline.send*/
type LogisticsOfflineSendRequest struct {
	values url.Values
}

func (this *LogisticsOfflineSendRequest) GetApiMethodName() string {
	return "taobao.logistics.offline.send"
}
func (this *LogisticsOfflineSendRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *LogisticsOfflineSendRequest) GetValues() url.Values {
	return this.values
}

/* 卖家联系人地址库ID，可以通过taobao.logistics.address.search接口查询到地址库ID。<br><font color='red'>如果为空，取的卖家的默认退货地址</font><br> */
func (this *LogisticsOfflineSendRequest) SetCancelId(value string) {
	this.Set("cancel_id", value)
}

/* 物流公司代码.如"POST"就代表中国邮政,"ZJS"就代表宅急送.调用 taobao.logistics.companies.get 获取。非淘宝官方物流合作公司，填写具体的物流公司名称，如“顺丰”。 */
func (this *LogisticsOfflineSendRequest) SetCompanyCode(value string) {
	this.Set("company_code", value)
}

/* feature参数格式<br>
范例: identCode=tid1:识别码1,识别码2|tid2:识别码3;machineCode=tid3:3C机器号A,3C机器号B<br>
identCode为识别码的KEY,machineCode为3C的KEY,多个key之间用”;”分隔<br>
“tid1:识别码1,识别码2|tid2:识别码3”为identCode对应的value。
"|"不同商品间的分隔符。<br>
例1商品和2商品，之间就用"|"分开。<br>
TID就是商品代表的子订单号，对应taobao.trade.fullinfo.get 接口获得的oid字段。(通过OID可以唯一定位到当前商品上)<br>
":"TID和具体传入参数间的分隔符。冒号前表示TID,之后代表该商品的参数属性。<br>
"," 属性间分隔符。（对应商品数量，当存在一个商品的数量超过1个时，用逗号分开）。<br>
具体:当订单中A商品的数量为2个，其中手机串号分别为"12345","67890"。<br>
参数格式：identCode=TIDA:12345,67890。
TIDA对应了A宝贝，冒号后用逗号分隔的"12345","67890".说明本订单A宝贝的数量为2，值分别为"12345","67890"。<br>
当存在"|"时，就说明订单中存在多个商品，商品间用"|"分隔了开来。|"之后的内容含义同上。 */
func (this *LogisticsOfflineSendRequest) SetFeature(value string) {
	this.Set("feature", value)
}

/* 表明是否是拆单
1表示拆单
0表示不拆单，默认值0 */
func (this *LogisticsOfflineSendRequest) SetIsSplit(value string) {
	this.Set("is_split", value)
}

/* 运单号.具体一个物流公司的真实运单号码。淘宝官方物流会校验，请谨慎传入；若company_code中传入的代码非淘宝官方物流合作公司，此处运单号不校验。 */
func (this *LogisticsOfflineSendRequest) SetOutSid(value string) {
	this.Set("out_sid", value)
}

/* 商家的IP地址 */
func (this *LogisticsOfflineSendRequest) SetSellerIp(value string) {
	this.Set("seller_ip", value)
}

/* 卖家联系人地址库ID，可以通过taobao.logistics.address.search接口查询到地址库ID。<font color='red'>如果为空，取的卖家的默认取货地址</font> */
func (this *LogisticsOfflineSendRequest) SetSenderId(value string) {
	this.Set("sender_id", value)
}

/* 需要拆单发货的子订单集合，为空表示不做拆单发货。 */
func (this *LogisticsOfflineSendRequest) SetSubTid(value string) {
	this.Set("sub_tid", value)
}

/* 淘宝交易ID<br /> 支持最小值为：1000 */
func (this *LogisticsOfflineSendRequest) SetTid(value string) {
	this.Set("tid", value)
}

type LogisticsOfflineSendResponse struct {
	*ErrorResponse                     `json:"error_response,omitempty"`
	LogisticsOfflineSendResponseResult `json:"logistics_offline_send_response"`
}
type LogisticsOfflineSendResponseResult struct {
	/* 返回发货是否成功is_success */
	Shipping *Shipping `json:"shipping"`
}

/*taobao.logistics.online.cancel*/
type LogisticsOnlineCancelRequest struct {
	values url.Values
}

func (this *LogisticsOnlineCancelRequest) GetApiMethodName() string {
	return "taobao.logistics.online.cancel"
}
func (this *LogisticsOnlineCancelRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *LogisticsOnlineCancelRequest) GetValues() url.Values {
	return this.values
}

/* 淘宝交易ID */
func (this *LogisticsOnlineCancelRequest) SetTid(value string) {
	this.Set("tid", value)
}

type LogisticsOnlineCancelResponse struct {
	*ErrorResponse                      `json:"error_response,omitempty"`
	LogisticsOnlineCancelResponseResult `json:"logistics_online_cancel_response"`
}
type LogisticsOnlineCancelResponseResult struct {
	/* 成功与失败 */
	IsSuccess bool `json:"is_success"`
	/* 修改时间 */
	ModifyTime string `json:"modify_time"`
	/* 重新创建物流订单id */
	RecreatedOrderId int64 `json:"recreated_order_id"`
}

/*taobao.logistics.online.confirm*/
type LogisticsOnlineConfirmRequest struct {
	values url.Values
}

func (this *LogisticsOnlineConfirmRequest) GetApiMethodName() string {
	return "taobao.logistics.online.confirm"
}
func (this *LogisticsOnlineConfirmRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *LogisticsOnlineConfirmRequest) GetValues() url.Values {
	return this.values
}

/* 表明是否是拆单，默认值0，1表示拆单 */
func (this *LogisticsOnlineConfirmRequest) SetIsSplit(value string) {
	this.Set("is_split", value)
}

/* 运单号.具体一个物流公司的真实运单号码。淘宝官方物流会校验，请谨慎传入；若company_code中传入的代码非淘宝官方物流合作公司，此处运单号不校验。<br> */
func (this *LogisticsOnlineConfirmRequest) SetOutSid(value string) {
	this.Set("out_sid", value)
}

/* 商家的IP地址 */
func (this *LogisticsOnlineConfirmRequest) SetSellerIp(value string) {
	this.Set("seller_ip", value)
}

/* 拆单子订单列表，对应的数据是：子订单号的列表。可以不传，但是如果传了则必须符合传递的规则。子订单必须是操作的物流订单的子订单的真子集！ */
func (this *LogisticsOnlineConfirmRequest) SetSubTid(value string) {
	this.Set("sub_tid", value)
}

/* 淘宝交易ID<br /> 支持最小值为：1000 */
func (this *LogisticsOnlineConfirmRequest) SetTid(value string) {
	this.Set("tid", value)
}

type LogisticsOnlineConfirmResponse struct {
	*ErrorResponse                       `json:"error_response,omitempty"`
	LogisticsOnlineConfirmResponseResult `json:"logistics_online_confirm_response"`
}
type LogisticsOnlineConfirmResponseResult struct {
	/* 只返回is_success：是否成功。 */
	Shipping *Shipping `json:"shipping"`
}

/*taobao.logistics.online.send*/
type LogisticsOnlineSendRequest struct {
	values url.Values
}

func (this *LogisticsOnlineSendRequest) GetApiMethodName() string {
	return "taobao.logistics.online.send"
}
func (this *LogisticsOnlineSendRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *LogisticsOnlineSendRequest) GetValues() url.Values {
	return this.values
}

/* 卖家联系人地址库ID，可以通过taobao.logistics.address.search接口查询到地址库ID。<br><font color='red'>如果为空，取的卖家的默认退货地址</font><br> */
func (this *LogisticsOnlineSendRequest) SetCancelId(value string) {
	this.Set("cancel_id", value)
}

/* 物流公司代码.如"POST"就代表中国邮政,"ZJS"就代表宅急送.调用 taobao.logistics.companies.get 获取。
<br><font color='red'>如果是货到付款订单，选择的物流公司必须支持货到付款发货方式</font> */
func (this *LogisticsOnlineSendRequest) SetCompanyCode(value string) {
	this.Set("company_code", value)
}

/* feature参数格式<br>
范例: identCode=tid1:识别码1,识别码2|tid2:识别码3;machineCode=tid3:3C机器号A,3C机器号B<br>
identCode为识别码的KEY,machineCode为3C的KEY,多个key之间用”;”分隔<br>
“tid1:识别码1,识别码2|tid2:识别码3”为identCode对应的value。
"|"不同商品间的分隔符。<br>
例1商品和2商品，之间就用"|"分开。<br>
TID就是商品代表的子订单号，对应taobao.trade.fullinfo.get 接口获得的oid字段。(通过OID可以唯一定位到当前商品上)<br>
":"TID和具体传入参数间的分隔符。冒号前表示TID,之后代表该商品的参数属性。<br>
"," 属性间分隔符。（对应商品数量，当存在一个商品的数量超过1个时，用逗号分开）。<br>
具体:当订单中A商品的数量为2个，其中手机串号分别为"12345","67890"。<br>
参数格式：identCode=TIDA:12345,67890。
TIDA对应了A宝贝，冒号后用逗号分隔的"12345","67890".说明本订单A宝贝的数量为2，值分别为"12345","67890"。<br>
当存在"|"时，就说明订单中存在多个商品，商品间用"|"分隔了开来。|"之后的内容含义同上。 */
func (this *LogisticsOnlineSendRequest) SetFeature(value string) {
	this.Set("feature", value)
}

/* 表明是否是拆单，默认值0，1表示拆单 */
func (this *LogisticsOnlineSendRequest) SetIsSplit(value string) {
	this.Set("is_split", value)
}

/* 运单号.具体一个物流公司的真实运单号码。淘宝官方物流会校验，请谨慎传入； */
func (this *LogisticsOnlineSendRequest) SetOutSid(value string) {
	this.Set("out_sid", value)
}

/* 商家的IP地址 */
func (this *LogisticsOnlineSendRequest) SetSellerIp(value string) {
	this.Set("seller_ip", value)
}

/* 卖家联系人地址库ID，可以通过taobao.logistics.address.search接口查询到地址库ID。<font color='red'>如果为空，取的卖家的默认取货地址</font> */
func (this *LogisticsOnlineSendRequest) SetSenderId(value string) {
	this.Set("sender_id", value)
}

/* 拆单子订单列表，对应的数据是：子订单号的列表。可以不传，但是如果传了则必须符合传递的规则。子订单必须是操作的物流订单的子订单的真子集！ */
func (this *LogisticsOnlineSendRequest) SetSubTid(value string) {
	this.Set("sub_tid", value)
}

/* 淘宝交易ID<br /> 支持最小值为：1000 */
func (this *LogisticsOnlineSendRequest) SetTid(value string) {
	this.Set("tid", value)
}

type LogisticsOnlineSendResponse struct {
	*ErrorResponse                    `json:"error_response,omitempty"`
	LogisticsOnlineSendResponseResult `json:"logistics_online_send_response"`
}
type LogisticsOnlineSendResponseResult struct {
	/* 返回发货是否成功is_success */
	Shipping *Shipping `json:"shipping"`
}

/*taobao.logistics.orders.detail.get*/
type LogisticsOrdersDetailGetRequest struct {
	values url.Values
}

func (this *LogisticsOrdersDetailGetRequest) GetApiMethodName() string {
	return "taobao.logistics.orders.detail.get"
}
func (this *LogisticsOrdersDetailGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *LogisticsOrdersDetailGetRequest) GetValues() url.Values {
	return this.values
}

/* 买家昵称 */
func (this *LogisticsOrdersDetailGetRequest) SetBuyerNick(value string) {
	this.Set("buyer_nick", value)
}

/* 创建时间结束.格式:yyyy-MM-dd HH:mm:ss */
func (this *LogisticsOrdersDetailGetRequest) SetEndCreated(value string) {
	this.Set("end_created", value)
}

/* 需返回的字段列表.可选值:Shipping 物流数据结构中所有字段.fileds中可以指定返回以上任意一个或者多个字段,以","分隔. */
func (this *LogisticsOrdersDetailGetRequest) SetFields(value string) {
	this.Set("fields", value)
}

/* 谁承担运费.可选值:buyer(买家),seller(卖家).如:buyer */
func (this *LogisticsOrdersDetailGetRequest) SetFreightPayer(value string) {
	this.Set("freight_payer", value)
}

/* 页码.该字段没传 或 值<1 ,则默认page_no为1 */
func (this *LogisticsOrdersDetailGetRequest) SetPageNo(value string) {
	this.Set("page_no", value)
}

/* 每页条数.该字段没传 或 值<1 ，则默认page_size为40<br /> 支持最大值为：100 */
func (this *LogisticsOrdersDetailGetRequest) SetPageSize(value string) {
	this.Set("page_size", value)
}

/* 收货人姓名 */
func (this *LogisticsOrdersDetailGetRequest) SetReceiverName(value string) {
	this.Set("receiver_name", value)
}

/* 卖家是否发货.可选值:yes(是),no(否).如:yes. */
func (this *LogisticsOrdersDetailGetRequest) SetSellerConfirm(value string) {
	this.Set("seller_confirm", value)
}

/* 创建时间开始.格式:yyyy-MM-dd HH:mm:ss */
func (this *LogisticsOrdersDetailGetRequest) SetStartCreated(value string) {
	this.Set("start_created", value)
}

/* 物流状态.可查看数据结构 Shipping 中的status字段. */
func (this *LogisticsOrdersDetailGetRequest) SetStatus(value string) {
	this.Set("status", value)
}

/* 交易ID.如果加入tid参数的话,不用传其他的参数,但是仅会返回一条物流订单信息. */
func (this *LogisticsOrdersDetailGetRequest) SetTid(value string) {
	this.Set("tid", value)
}

/* 物流方式.可选值:post(平邮),express(快递),ems(EMS).如:post */
func (this *LogisticsOrdersDetailGetRequest) SetType(value string) {
	this.Set("type", value)
}

type LogisticsOrdersDetailGetResponse struct {
	*ErrorResponse                         `json:"error_response,omitempty"`
	LogisticsOrdersDetailGetResponseResult `json:"logistics_orders_detail_get_response"`
}
type LogisticsOrdersDetailGetResponseResult struct {
	/* 获取的物流订单详情列表.返回的Shipping包含的具体信息为入参fields请求的字段信息. */
	Shippings struct {
		Shipping []*Shipping `json:"shipping"`
	} `json:"Shippings"`
	/* 搜索到的物流订单列表总数 */
	TotalResults int64 `json:"total_results"`
}

/*taobao.logistics.orders.get*/
type LogisticsOrdersGetRequest struct {
	values url.Values
}

func (this *LogisticsOrdersGetRequest) GetApiMethodName() string {
	return "taobao.logistics.orders.get"
}
func (this *LogisticsOrdersGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *LogisticsOrdersGetRequest) GetValues() url.Values {
	return this.values
}

/* 买家昵称 */
func (this *LogisticsOrdersGetRequest) SetBuyerNick(value string) {
	this.Set("buyer_nick", value)
}

/* 创建时间结束 */
func (this *LogisticsOrdersGetRequest) SetEndCreated(value string) {
	this.Set("end_created", value)
}

/* 需返回的字段列表.可选值:Shipping 物流数据结构中的以下字段: <br>
tid,order_code,seller_nick,buyer_nick,delivery_start, delivery_end,out_sid,item_title,receiver_name, created,modified,status,type,freight_payer,seller_confirm,company_name,sub_tids,is_spilt；<br>多个字段之间用","分隔。如tid,seller_nick,buyer_nick,delivery_start。 */
func (this *LogisticsOrdersGetRequest) SetFields(value string) {
	this.Set("fields", value)
}

/* 谁承担运费.可选值:buyer(买家),seller(卖家).如:buyer */
func (this *LogisticsOrdersGetRequest) SetFreightPayer(value string) {
	this.Set("freight_payer", value)
}

/* 页码.该字段没传 或 值<1 ,则默认page_no为1 */
func (this *LogisticsOrdersGetRequest) SetPageNo(value string) {
	this.Set("page_no", value)
}

/* 每页条数.该字段没传 或 值<1 ,则默认page_size为40<br /> 支持最大值为：100 */
func (this *LogisticsOrdersGetRequest) SetPageSize(value string) {
	this.Set("page_size", value)
}

/* 收货人姓名 */
func (this *LogisticsOrdersGetRequest) SetReceiverName(value string) {
	this.Set("receiver_name", value)
}

/* 卖家是否发货.可选值:yes(是),no(否).如:yes */
func (this *LogisticsOrdersGetRequest) SetSellerConfirm(value string) {
	this.Set("seller_confirm", value)
}

/* 创建时间开始 */
func (this *LogisticsOrdersGetRequest) SetStartCreated(value string) {
	this.Set("start_created", value)
}

/* 物流状态.查看数据结构 Shipping 中的status字段. */
func (this *LogisticsOrdersGetRequest) SetStatus(value string) {
	this.Set("status", value)
}

/* 交易ID.如果加入tid参数的话,不用传其他的参数,若传入tid：非拆单场景，仅会返回一条物流订单信息；拆单场景，会返回多条物流订单信息 */
func (this *LogisticsOrdersGetRequest) SetTid(value string) {
	this.Set("tid", value)
}

/* 物流方式.可选值:post(平邮),express(快递),ems(EMS).如:post */
func (this *LogisticsOrdersGetRequest) SetType(value string) {
	this.Set("type", value)
}

type LogisticsOrdersGetResponse struct {
	*ErrorResponse                   `json:"error_response,omitempty"`
	LogisticsOrdersGetResponseResult `json:"logistics_orders_get_response"`
}
type LogisticsOrdersGetResponseResult struct {
	/* 获取的物流订单详情列表
	返回的Shipping包含的具体信息为入参fields请求的字段信息 */
	Shippings struct {
		Shipping []*Shipping `json:"shipping"`
	} `json:"Shippings"`
	/* 搜索到的物流订单列表总数 */
	TotalResults int64 `json:"total_results"`
}

/*taobao.logistics.partners.get*/
type LogisticsPartnersGetRequest struct {
	values url.Values
}

func (this *LogisticsPartnersGetRequest) GetApiMethodName() string {
	return "taobao.logistics.partners.get"
}
func (this *LogisticsPartnersGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *LogisticsPartnersGetRequest) GetValues() url.Values {
	return this.values
}

/* 货物价格.只有当选择货到付款此参数才会有效 */
func (this *LogisticsPartnersGetRequest) SetGoodsValue(value string) {
	this.Set("goods_value", value)
}

/* 是否需要揽收资费信息，默认false。在此值为false时，返回内容中将无carriage。在未设置source_id或target_id的情况下，无法查询揽收资费信息。自己联系无揽收资费记录。 */
func (this *LogisticsPartnersGetRequest) SetIsNeedCarriage(value string) {
	this.Set("is_need_carriage", value)
}

/* 服务类型，根据此参数可查出提供相应服务类型的物流公司信息(物流公司状态正常)，可选值：cod(货到付款)、online(在线下单)、 offline(自己联系)、limit(限时物流)。然后再根据source_id,target_id,goods_value这三个条件来过滤物流公司. 目前输入自己联系服务类型将会返回空，因为自己联系并没有具体的服务信息，所以不会有记录。 */
func (this *LogisticsPartnersGetRequest) SetServiceType(value string) {
	this.Set("service_type", value)
}

/* 物流公司揽货地地区码（必须是区、县一级的）.参考:<a href="http://www.stats.gov.cn/tjsj/tjbz/xzqhdm/201401/t20140116_501070.html">http://www.stats.gov.cn/tjsj/tjbz/xzqhdm/201401/t20140116_501070.html</a>或者调用 taobao.areas.get 获取 */
func (this *LogisticsPartnersGetRequest) SetSourceId(value string) {
	this.Set("source_id", value)
}

/* 物流公司派送地地区码（必须是区、县一级的）.参考:<a href="http://www.stats.gov.cn/tjsj/tjbz/xzqhdm/201401/t20140116_501070.html">http://www.stats.gov.cn/tjsj/tjbz/xzqhdm/201401/t20140116_501070.html</a>或者调用 taobao.areas.get 获取 */
func (this *LogisticsPartnersGetRequest) SetTargetId(value string) {
	this.Set("target_id", value)
}

type LogisticsPartnersGetResponse struct {
	*ErrorResponse                     `json:"error_response,omitempty"`
	LogisticsPartnersGetResponseResult `json:"logistics_partners_get_response"`
}
type LogisticsPartnersGetResponseResult struct {
	/* 物流公司信息。 */
	LogisticsPartners struct {
		LogisticsPartner []*LogisticsPartner `json:"logistics_partner"`
	} `json:"LogisticsPartners"`
}

/*taobao.logistics.trace.search*/
type LogisticsTraceSearchRequest struct {
	values url.Values
}

func (this *LogisticsTraceSearchRequest) GetApiMethodName() string {
	return "taobao.logistics.trace.search"
}
func (this *LogisticsTraceSearchRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *LogisticsTraceSearchRequest) GetValues() url.Values {
	return this.values
}

/* 表明是否是拆单，默认值0，1表示拆单 */
func (this *LogisticsTraceSearchRequest) SetIsSplit(value string) {
	this.Set("is_split", value)
}

/* 卖家昵称 */
func (this *LogisticsTraceSearchRequest) SetSellerNick(value string) {
	this.Set("seller_nick", value)
}

/* 拆单子订单列表，对应的数据是：子订单号的列表。可以不传，但是如果传了则必须符合传递的规则。子订单必须是操作的物流订单的子订单的真子集 */
func (this *LogisticsTraceSearchRequest) SetSubTid(value string) {
	this.Set("sub_tid", value)
}

/* 淘宝交易号，请勿传非淘宝交易号 */
func (this *LogisticsTraceSearchRequest) SetTid(value string) {
	this.Set("tid", value)
}

type LogisticsTraceSearchResponse struct {
	*ErrorResponse                     `json:"error_response,omitempty"`
	LogisticsTraceSearchResponseResult `json:"logistics_trace_search_response"`
}
type LogisticsTraceSearchResponseResult struct {
	/* 物流公司名称 */
	CompanyName string `json:"company_name"`
	/* 运单号 */
	OutSid string `json:"out_sid"`
	/* 订单的物流状态（仅支持线上发货online订单，线下发货offline发出后直接变为已签收）
	* 等候发送给物流公司
	*已提交给物流公司,等待物流公司接单
	*已经确认消息接收，等待物流公司接单
	*物流公司已接单
	*物流公司不接单
	*物流公司揽收失败
	*物流公司揽收成功
	*签收失败
	*对方已签收
	*对方拒绝签收 */
	Status string `json:"status"`
	/* 交易号 */
	Tid int64 `json:"tid"`
	/* 流转信息列表 */
	TraceList struct {
		TransitStepInfo []*TransitStepInfo `json:"transit_step_info"`
	} `json:"TraceList"`
}

/*taobao.sellercats.list.add*/
type SellercatsListAddRequest struct {
	values url.Values
}

func (this *SellercatsListAddRequest) GetApiMethodName() string {
	return "taobao.sellercats.list.add"
}
func (this *SellercatsListAddRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *SellercatsListAddRequest) GetValues() url.Values {
	return this.values
}

/* 卖家自定义类目名称。不超过20个字符 */
func (this *SellercatsListAddRequest) SetName(value string) {
	this.Set("name", value)
}

/* 父类目编号，如果类目为店铺下的一级类目：值等于0，如果类目为子类目，调用获取taobao.sellercats.list.get父类目编号 */
func (this *SellercatsListAddRequest) SetParentCid(value string) {
	this.Set("parent_cid", value)
}

/* 链接图片URL地址。(绝对地址，格式：http://host/image_path) */
func (this *SellercatsListAddRequest) SetPictUrl(value string) {
	this.Set("pict_url", value)
}

/* 该类目在页面上的排序位置,取值范围:大于零的整数 */
func (this *SellercatsListAddRequest) SetSortOrder(value string) {
	this.Set("sort_order", value)
}

type SellercatsListAddResponse struct {
	*ErrorResponse                  `json:"error_response,omitempty"`
	SellercatsListAddResponseResult `json:"sellercats_list_add_response"`
}
type SellercatsListAddResponseResult struct {
	/* 返回seller_cat数据结构中的：cid,created */
	SellerCat *SellerCat `json:"seller_cat"`
}

/*taobao.sellercats.list.get*/
type SellercatsListGetRequest struct {
	values url.Values
}

func (this *SellercatsListGetRequest) GetApiMethodName() string {
	return "taobao.sellercats.list.get"
}
func (this *SellercatsListGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *SellercatsListGetRequest) GetValues() url.Values {
	return this.values
}

/* 卖家昵称 */
func (this *SellercatsListGetRequest) SetNick(value string) {
	this.Set("nick", value)
}

type SellercatsListGetResponse struct {
	*ErrorResponse                  `json:"error_response,omitempty"`
	SellercatsListGetResponseResult `json:"sellercats_list_get_response"`
}
type SellercatsListGetResponseResult struct {
	/* 卖家自定义类目 */
	SellerCats struct {
		SellerCat []*SellerCat `json:"seller_cat"`
	} `json:"SellerCats"`
}

/*taobao.sellercats.list.update*/
type SellercatsListUpdateRequest struct {
	values url.Values
}

func (this *SellercatsListUpdateRequest) GetApiMethodName() string {
	return "taobao.sellercats.list.update"
}
func (this *SellercatsListUpdateRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *SellercatsListUpdateRequest) GetValues() url.Values {
	return this.values
}

/* 卖家自定义类目编号 */
func (this *SellercatsListUpdateRequest) SetCid(value string) {
	this.Set("cid", value)
}

/* 卖家自定义类目名称。不超过20个字符 */
func (this *SellercatsListUpdateRequest) SetName(value string) {
	this.Set("name", value)
}

/* 链接图片URL地址 */
func (this *SellercatsListUpdateRequest) SetPictUrl(value string) {
	this.Set("pict_url", value)
}

/* 该类目在页面上的排序位置,取值范围:大于零的整数 */
func (this *SellercatsListUpdateRequest) SetSortOrder(value string) {
	this.Set("sort_order", value)
}

type SellercatsListUpdateResponse struct {
	*ErrorResponse                     `json:"error_response,omitempty"`
	SellercatsListUpdateResponseResult `json:"sellercats_list_update_response"`
}
type SellercatsListUpdateResponseResult struct {
	/* 返回sellercat数据结构中的：cid,modified */
	SellerCat *SellerCat `json:"seller_cat"`
}

/*taobao.shop.get*/
type ShopGetRequest struct {
	values url.Values
}

func (this *ShopGetRequest) GetApiMethodName() string {
	return "taobao.shop.get"
}
func (this *ShopGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *ShopGetRequest) GetValues() url.Values {
	return this.values
}

/* 需返回的字段列表。可选值：Shop 结构中的所有字段；多个字段之间用逗号(,)分隔 */
func (this *ShopGetRequest) SetFields(value string) {
	this.Set("fields", value)
}

/* 卖家昵称 */
func (this *ShopGetRequest) SetNick(value string) {
	this.Set("nick", value)
}

type ShopGetResponse struct {
	*ErrorResponse        `json:"error_response,omitempty"`
	ShopGetResponseResult `json:"shop_get_response"`
}
type ShopGetResponseResult struct {
	/* 店铺信息 */
	Shop *Shop `json:"shop"`
}

/*taobao.shop.remainshowcase.get*/
type ShopRemainshowcaseGetRequest struct {
	values url.Values
}

func (this *ShopRemainshowcaseGetRequest) GetApiMethodName() string {
	return "taobao.shop.remainshowcase.get"
}
func (this *ShopRemainshowcaseGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *ShopRemainshowcaseGetRequest) GetValues() url.Values {
	return this.values
}

type ShopRemainshowcaseGetResponse struct {
	*ErrorResponse                      `json:"error_response,omitempty"`
	ShopRemainshowcaseGetResponseResult `json:"shop_remainshowcase_get_response"`
}
type ShopRemainshowcaseGetResponseResult struct {
	/* 支持返回剩余橱窗数量，已用橱窗数量，总橱窗数量 */
	Shop *Shop `json:"shop"`
}

/*taobao.shop.update*/
type ShopUpdateRequest struct {
	values url.Values
}

func (this *ShopUpdateRequest) GetApiMethodName() string {
	return "taobao.shop.update"
}
func (this *ShopUpdateRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *ShopUpdateRequest) GetValues() url.Values {
	return this.values
}

/* 店铺公告。不超过1024个字符 */
func (this *ShopUpdateRequest) SetBulletin(value string) {
	this.Set("bulletin", value)
}

/* 店铺描述。10～2000个字符以内 */
func (this *ShopUpdateRequest) SetDesc(value string) {
	this.Set("desc", value)
}

/* 店铺标题。不超过30个字符；过滤敏感词，如淘咖啡、阿里巴巴等。title, bulletin和desc至少必须传一个 */
func (this *ShopUpdateRequest) SetTitle(value string) {
	this.Set("title", value)
}

type ShopUpdateResponse struct {
	*ErrorResponse           `json:"error_response,omitempty"`
	ShopUpdateResponseResult `json:"shop_update_response"`
}
type ShopUpdateResponseResult struct {
	/* 店铺信息 */
	Shop *Shop `json:"shop"`
}

/*taobao.shopcats.list.get*/
type ShopcatsListGetRequest struct {
	values url.Values
}

func (this *ShopcatsListGetRequest) GetApiMethodName() string {
	return "taobao.shopcats.list.get"
}
func (this *ShopcatsListGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *ShopcatsListGetRequest) GetValues() url.Values {
	return this.values
}

/* 需要返回的字段列表，见ShopCat，默认返回：cid,parent_cid,name,is_parent */
func (this *ShopcatsListGetRequest) SetFields(value string) {
	this.Set("fields", value)
}

type ShopcatsListGetResponse struct {
	*ErrorResponse                `json:"error_response,omitempty"`
	ShopcatsListGetResponseResult `json:"shopcats_list_get_response"`
}
type ShopcatsListGetResponseResult struct {
	/* 店铺类目列表信息 */
	ShopCats struct {
		ShopCat []*ShopCat `json:"shop_cat"`
	} `json:"ShopCats"`
}

/*taobao.fenxiao.cooperation.audit*/
type FenxiaoCooperationAuditRequest struct {
	values url.Values
}

func (this *FenxiaoCooperationAuditRequest) GetApiMethodName() string {
	return "taobao.fenxiao.cooperation.audit"
}
func (this *FenxiaoCooperationAuditRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *FenxiaoCooperationAuditRequest) GetValues() url.Values {
	return this.values
}

/* 1:审批通过，审批通过后要加入授权产品线列表；
2:审批拒绝 */
func (this *FenxiaoCooperationAuditRequest) SetAuditResult(value string) {
	this.Set("audit_result", value)
}

/* 当审批通过时需要指定授权产品线id列表(代销)，如果没传则表示不开通，且经销和代销的授权产品线列表至少传入一种，同时传入则表示都开通。 */
func (this *FenxiaoCooperationAuditRequest) SetProductLineListAgent(value string) {
	this.Set("product_line_list_agent", value)
}

/* 当审批通过时需要指定授权产品线id列表(经销)，如果没传则表示不开通，且经销和代销的授权产品线列表至少传入一种，同时传入则表示都开通。 */
func (this *FenxiaoCooperationAuditRequest) SetProductLineListDealer(value string) {
	this.Set("product_line_list_dealer", value)
}

/* 备注 */
func (this *FenxiaoCooperationAuditRequest) SetRemark(value string) {
	this.Set("remark", value)
}

/* 合作申请Id */
func (this *FenxiaoCooperationAuditRequest) SetRequisitionId(value string) {
	this.Set("requisition_id", value)
}

type FenxiaoCooperationAuditResponse struct {
	*ErrorResponse                        `json:"error_response,omitempty"`
	FenxiaoCooperationAuditResponseResult `json:"fenxiao_cooperation_audit_response"`
}
type FenxiaoCooperationAuditResponseResult struct {
	/* 操作是否成功 */
	IsSuccess bool `json:"is_success"`
}

/*taobao.fenxiao.cooperation.get*/
type FenxiaoCooperationGetRequest struct {
	values url.Values
}

func (this *FenxiaoCooperationGetRequest) GetApiMethodName() string {
	return "taobao.fenxiao.cooperation.get"
}
func (this *FenxiaoCooperationGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *FenxiaoCooperationGetRequest) GetValues() url.Values {
	return this.values
}

/* 合作结束时间yyyy-MM-dd HH:mm:ss */
func (this *FenxiaoCooperationGetRequest) SetEndDate(value string) {
	this.Set("end_date", value)
}

/* 页码（大于0的整数，默认1） */
func (this *FenxiaoCooperationGetRequest) SetPageNo(value string) {
	this.Set("page_no", value)
}

/* 每页记录数（默认20，最大50） */
func (this *FenxiaoCooperationGetRequest) SetPageSize(value string) {
	this.Set("page_size", value)
}

/* 合作开始时间yyyy-MM-dd HH:mm:ss */
func (this *FenxiaoCooperationGetRequest) SetStartDate(value string) {
	this.Set("start_date", value)
}

/* 合作状态： NORMAL(合作中)、 ENDING(终止中) 、END (终止) */
func (this *FenxiaoCooperationGetRequest) SetStatus(value string) {
	this.Set("status", value)
}

/* 分销方式：AGENT(代销) 、DEALER（经销） */
func (this *FenxiaoCooperationGetRequest) SetTradeType(value string) {
	this.Set("trade_type", value)
}

type FenxiaoCooperationGetResponse struct {
	*ErrorResponse                      `json:"error_response,omitempty"`
	FenxiaoCooperationGetResponseResult `json:"fenxiao_cooperation_get_response"`
}
type FenxiaoCooperationGetResponseResult struct {
	/* 合作分销关系 */
	Cooperations struct {
		Cooperation []*Cooperation `json:"cooperation"`
	} `json:"Cooperations"`
	/* 结果记录数 */
	TotalResults int64 `json:"total_results"`
}

/*taobao.fenxiao.cooperation.productcat.add*/
type FenxiaoCooperationProductcatAddRequest struct {
	values url.Values
}

func (this *FenxiaoCooperationProductcatAddRequest) GetApiMethodName() string {
	return "taobao.fenxiao.cooperation.productcat.add"
}
func (this *FenxiaoCooperationProductcatAddRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *FenxiaoCooperationProductcatAddRequest) GetValues() url.Values {
	return this.values
}

/* 合作关系id */
func (this *FenxiaoCooperationProductcatAddRequest) SetCooperateId(value string) {
	this.Set("cooperate_id", value)
}

/* 等级ID（为空则不修改） */
func (this *FenxiaoCooperationProductcatAddRequest) SetGradeId(value string) {
	this.Set("grade_id", value)
}

/* 产品线id列表，若有多个，以逗号分隔 */
func (this *FenxiaoCooperationProductcatAddRequest) SetProductLineList(value string) {
	this.Set("product_line_list", value)
}

type FenxiaoCooperationProductcatAddResponse struct {
	*ErrorResponse                                `json:"error_response,omitempty"`
	FenxiaoCooperationProductcatAddResponseResult `json:"fenxiao_cooperation_productcat_add_response"`
}
type FenxiaoCooperationProductcatAddResponseResult struct {
	/* 操作是否成功 */
	IsSuccess bool `json:"is_success"`
}

/*taobao.fenxiao.cooperation.terminate*/
type FenxiaoCooperationTerminateRequest struct {
	values url.Values
}

func (this *FenxiaoCooperationTerminateRequest) GetApiMethodName() string {
	return "taobao.fenxiao.cooperation.terminate"
}
func (this *FenxiaoCooperationTerminateRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *FenxiaoCooperationTerminateRequest) GetValues() url.Values {
	return this.values
}

/* 合作编号 */
func (this *FenxiaoCooperationTerminateRequest) SetCooperateId(value string) {
	this.Set("cooperate_id", value)
}

/* 终止天数，可选1,2,3,5,7,15，在多少天数内终止 */
func (this *FenxiaoCooperationTerminateRequest) SetEndRemainDays(value string) {
	this.Set("end_remain_days", value)
}

/* 终止说明（5-2000字）<br /> 支持最大长度为：2000<br /> 支持的最大列表长度为：2000 */
func (this *FenxiaoCooperationTerminateRequest) SetEndRemark(value string) {
	this.Set("end_remark", value)
}

type FenxiaoCooperationTerminateResponse struct {
	*ErrorResponse                            `json:"error_response,omitempty"`
	FenxiaoCooperationTerminateResponseResult `json:"fenxiao_cooperation_terminate_response"`
}
type FenxiaoCooperationTerminateResponseResult struct {
	/* 操作是否成功 */
	IsSuccess bool `json:"is_success"`
}

/*taobao.fenxiao.cooperation.update*/
type FenxiaoCooperationUpdateRequest struct {
	values url.Values
}

func (this *FenxiaoCooperationUpdateRequest) GetApiMethodName() string {
	return "taobao.fenxiao.cooperation.update"
}
func (this *FenxiaoCooperationUpdateRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *FenxiaoCooperationUpdateRequest) GetValues() url.Values {
	return this.values
}

/* 分销商ID */
func (this *FenxiaoCooperationUpdateRequest) SetDistributorId(value string) {
	this.Set("distributor_id", value)
}

/* 等级ID(0代表取消) */
func (this *FenxiaoCooperationUpdateRequest) SetGradeId(value string) {
	this.Set("grade_id", value)
}

/* 分销方式(新增)： AGENT(代销)、DEALER(经销)(默认为代销) */
func (this *FenxiaoCooperationUpdateRequest) SetTradeType(value string) {
	this.Set("trade_type", value)
}

type FenxiaoCooperationUpdateResponse struct {
	*ErrorResponse                         `json:"error_response,omitempty"`
	FenxiaoCooperationUpdateResponseResult `json:"fenxiao_cooperation_update_response"`
}
type FenxiaoCooperationUpdateResponseResult struct {
	/* 更新结果成功失败 */
	IsSuccess bool `json:"is_success"`
}

/*taobao.fenxiao.dealer.requisitionorder.agree*/
type FenxiaoDealerRequisitionorderAgreeRequest struct {
	values url.Values
}

func (this *FenxiaoDealerRequisitionorderAgreeRequest) GetApiMethodName() string {
	return "taobao.fenxiao.dealer.requisitionorder.agree"
}
func (this *FenxiaoDealerRequisitionorderAgreeRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *FenxiaoDealerRequisitionorderAgreeRequest) GetValues() url.Values {
	return this.values
}

/* 采购申请/经销采购单编号 */
func (this *FenxiaoDealerRequisitionorderAgreeRequest) SetDealerOrderId(value string) {
	this.Set("dealer_order_id", value)
}

type FenxiaoDealerRequisitionorderAgreeResponse struct {
	*ErrorResponse                                   `json:"error_response,omitempty"`
	FenxiaoDealerRequisitionorderAgreeResponseResult `json:"fenxiao_dealer_requisitionorder_agree_response"`
}
type FenxiaoDealerRequisitionorderAgreeResponseResult struct {
	/* 操作是否成功。true：成功；false：失败。 */
	IsSuccess bool `json:"is_success"`
}

/*taobao.fenxiao.dealer.requisitionorder.close*/
type FenxiaoDealerRequisitionorderCloseRequest struct {
	values url.Values
}

func (this *FenxiaoDealerRequisitionorderCloseRequest) GetApiMethodName() string {
	return "taobao.fenxiao.dealer.requisitionorder.close"
}
func (this *FenxiaoDealerRequisitionorderCloseRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *FenxiaoDealerRequisitionorderCloseRequest) GetValues() url.Values {
	return this.values
}

/* 经销采购单编号 */
func (this *FenxiaoDealerRequisitionorderCloseRequest) SetDealerOrderId(value string) {
	this.Set("dealer_order_id", value)
}

/* 关闭原因：
1：长时间无法联系到分销商，取消交易。
2：分销商错误提交申请，取消交易。
3：缺货，近段时间都无法发货。
4：分销商恶意提交申请单。
5：其他原因。 */
func (this *FenxiaoDealerRequisitionorderCloseRequest) SetReason(value string) {
	this.Set("reason", value)
}

/* 关闭详细原因，字数5-200字 */
func (this *FenxiaoDealerRequisitionorderCloseRequest) SetReasonDetail(value string) {
	this.Set("reason_detail", value)
}

type FenxiaoDealerRequisitionorderCloseResponse struct {
	*ErrorResponse                                   `json:"error_response,omitempty"`
	FenxiaoDealerRequisitionorderCloseResponseResult `json:"fenxiao_dealer_requisitionorder_close_response"`
}
type FenxiaoDealerRequisitionorderCloseResponseResult struct {
	/* 操作是否成功。true：成功；false：失败。 */
	IsSuccess bool `json:"is_success"`
}

/*taobao.fenxiao.dealer.requisitionorder.create*/
type FenxiaoDealerRequisitionorderCreateRequest struct {
	values url.Values
}

func (this *FenxiaoDealerRequisitionorderCreateRequest) GetApiMethodName() string {
	return "taobao.fenxiao.dealer.requisitionorder.create"
}
func (this *FenxiaoDealerRequisitionorderCreateRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *FenxiaoDealerRequisitionorderCreateRequest) GetValues() url.Values {
	return this.values
}

/* 收货人所在街道地址 */
func (this *FenxiaoDealerRequisitionorderCreateRequest) SetAddress(value string) {
	this.Set("address", value)
}

/* 买家姓名（自提方式填写提货人姓名） */
func (this *FenxiaoDealerRequisitionorderCreateRequest) SetBuyerName(value string) {
	this.Set("buyer_name", value)
}

/* 收货人所在市 */
func (this *FenxiaoDealerRequisitionorderCreateRequest) SetCity(value string) {
	this.Set("city", value)
}

/* 收货人所在区 */
func (this *FenxiaoDealerRequisitionorderCreateRequest) SetDistrict(value string) {
	this.Set("district", value)
}

/* 身份证号（自提方式必填，填写提货人身份证号码，提货时用于确认提货人身份） */
func (this *FenxiaoDealerRequisitionorderCreateRequest) SetIdCardNumber(value string) {
	this.Set("id_card_number", value)
}

/* 配送方式。SELF_PICKUP：自提；LOGISTICS：仓库发货 */
func (this *FenxiaoDealerRequisitionorderCreateRequest) SetLogisticsType(value string) {
	this.Set("logistics_type", value)
}

/* 买家的手机号码（1、此字段与phone字段至少填写一个。2、自提方式下此字段必填，保存提货人联系电话） */
func (this *FenxiaoDealerRequisitionorderCreateRequest) SetMobile(value string) {
	this.Set("mobile", value)
}

/* 采购清单，存放多个采购明细，每个采购明细内部以‘:’隔开，多个采购明细之间以‘,’隔开. 例(分销产品id:skuid:购买数量:申请单价,分销产品id:skuid:购买数量:申请单价)，申请单价的单位为分。不存在sku请留空skuid，如（分销产品id::购买数量:申请单价） */
func (this *FenxiaoDealerRequisitionorderCreateRequest) SetOrderDetail(value string) {
	this.Set("order_detail", value)
}

/* 买家联系电话（此字段和mobile字段至少填写一个） */
func (this *FenxiaoDealerRequisitionorderCreateRequest) SetPhone(value string) {
	this.Set("phone", value)
}

/* 收货人所在地区邮政编码 */
func (this *FenxiaoDealerRequisitionorderCreateRequest) SetPostCode(value string) {
	this.Set("post_code", value)
}

/* 收货人所在省份 */
func (this *FenxiaoDealerRequisitionorderCreateRequest) SetProvince(value string) {
	this.Set("province", value)
}

type FenxiaoDealerRequisitionorderCreateResponse struct {
	*ErrorResponse                                    `json:"error_response,omitempty"`
	FenxiaoDealerRequisitionorderCreateResponseResult `json:"fenxiao_dealer_requisitionorder_create_response"`
}
type FenxiaoDealerRequisitionorderCreateResponseResult struct {
	/* 经销采购申请编号 */
	DealerOrderId int64 `json:"dealer_order_id"`
}

/*taobao.fenxiao.dealer.requisitionorder.get*/
type FenxiaoDealerRequisitionorderGetRequest struct {
	values url.Values
}

func (this *FenxiaoDealerRequisitionorderGetRequest) GetApiMethodName() string {
	return "taobao.fenxiao.dealer.requisitionorder.get"
}
func (this *FenxiaoDealerRequisitionorderGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *FenxiaoDealerRequisitionorderGetRequest) GetValues() url.Values {
	return this.values
}

/* 采购申请/经销采购单最迟修改时间。与start_date字段的最大时间间隔不能超过30天 */
func (this *FenxiaoDealerRequisitionorderGetRequest) SetEndDate(value string) {
	this.Set("end_date", value)
}

/* 多个字段用","分隔。 fields 如果为空：返回所有采购申请/经销采购单对象(dealer_orders)字段。 如果不为空：返回指定采购单对象(dealer_orders)字段。 例1： dealer_order_details.product_id 表示只返回product_id 例2： dealer_order_details 表示只返回明细列表 */
func (this *FenxiaoDealerRequisitionorderGetRequest) SetFields(value string) {
	this.Set("fields", value)
}

/* 查询者自己在所要查询的采购申请/经销采购单中的身份。
1：供应商。
2：分销商。
注：填写其他值当做错误处理。 */
func (this *FenxiaoDealerRequisitionorderGetRequest) SetIdentity(value string) {
	this.Set("identity", value)
}

/* 采购申请/经销采购单状态。
0：全部状态。
1：分销商提交申请，待供应商审核。
2：供应商驳回申请，待分销商确认。
3：供应商修改后，待分销商确认。
4：分销商拒绝修改，待供应商再审核。
5：审核通过下单成功，待分销商付款。
7：付款成功，待供应商发货。
8：供应商发货，待分销商收货。
9：分销商收货，交易成功。
10：采购申请/经销采购单关闭。

注：无值按默认值0计，超出状态范围返回错误信息。 */
func (this *FenxiaoDealerRequisitionorderGetRequest) SetOrderStatus(value string) {
	this.Set("order_status", value)
}

/* 页码（大于0的整数。无值或小于1的值按默认值1计） */
func (this *FenxiaoDealerRequisitionorderGetRequest) SetPageNo(value string) {
	this.Set("page_no", value)
}

/* 每页条数（大于0但小于等于50的整数。无值或大于50或小于1的值按默认值50计） */
func (this *FenxiaoDealerRequisitionorderGetRequest) SetPageSize(value string) {
	this.Set("page_size", value)
}

/* 采购申请/经销采购单最早修改时间 */
func (this *FenxiaoDealerRequisitionorderGetRequest) SetStartDate(value string) {
	this.Set("start_date", value)
}

type FenxiaoDealerRequisitionorderGetResponse struct {
	*ErrorResponse                                 `json:"error_response,omitempty"`
	FenxiaoDealerRequisitionorderGetResponseResult `json:"fenxiao_dealer_requisitionorder_get_response"`
}
type FenxiaoDealerRequisitionorderGetResponseResult struct {
	/* 采购申请/经销采购单结果列表 */
	DealerOrders struct {
		DealerOrder []*DealerOrder `json:"dealer_order"`
	} `json:"DealerOrders"`
	/* 按查询条件查到的记录总数 */
	TotalResults int64 `json:"total_results"`
}

/*taobao.fenxiao.dealer.requisitionorder.modify*/
type FenxiaoDealerRequisitionorderModifyRequest struct {
	values url.Values
}

func (this *FenxiaoDealerRequisitionorderModifyRequest) GetApiMethodName() string {
	return "taobao.fenxiao.dealer.requisitionorder.modify"
}
func (this *FenxiaoDealerRequisitionorderModifyRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *FenxiaoDealerRequisitionorderModifyRequest) GetValues() url.Values {
	return this.values
}

/* 经销采购单编号 */
func (this *FenxiaoDealerRequisitionorderModifyRequest) SetDealerOrderId(value string) {
	this.Set("dealer_order_id", value)
}

/* 要删除的商品明细id列表，多个id使用英文符号的逗号隔开 */
func (this *FenxiaoDealerRequisitionorderModifyRequest) SetDeleteDetailIds(value string) {
	this.Set("delete_detail_ids", value)
}

/* 经销采购单的商品明细的新的采购价格。格式为商品明细id:价格修改值,商品明细id:价格修改值 */
func (this *FenxiaoDealerRequisitionorderModifyRequest) SetDetailIdPrices(value string) {
	this.Set("detail_id_prices", value)
}

/* 修改经销采购单的商品明细的新的库存。格式为商品明细id:库存修改值,商品明细id:库存修改值 */
func (this *FenxiaoDealerRequisitionorderModifyRequest) SetDetailIdQuantities(value string) {
	this.Set("detail_id_quantities", value)
}

/* 新邮费（单位：分，示例值1005表示10.05元）。必须大于等于0。自提方式不可修改邮费。不提交该参数表示不修改邮费。 */
func (this *FenxiaoDealerRequisitionorderModifyRequest) SetNewPostFee(value string) {
	this.Set("new_post_fee", value)
}

type FenxiaoDealerRequisitionorderModifyResponse struct {
	*ErrorResponse                                    `json:"error_response,omitempty"`
	FenxiaoDealerRequisitionorderModifyResponseResult `json:"fenxiao_dealer_requisitionorder_modify_response"`
}
type FenxiaoDealerRequisitionorderModifyResponseResult struct {
	/* 操作是否成功。true：成功；false：失败。 */
	IsSuccess bool `json:"is_success"`
}

/*taobao.fenxiao.dealer.requisitionorder.query*/
type FenxiaoDealerRequisitionorderQueryRequest struct {
	values url.Values
}

func (this *FenxiaoDealerRequisitionorderQueryRequest) GetApiMethodName() string {
	return "taobao.fenxiao.dealer.requisitionorder.query"
}
func (this *FenxiaoDealerRequisitionorderQueryRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *FenxiaoDealerRequisitionorderQueryRequest) GetValues() url.Values {
	return this.values
}

/* 经销采购单编号。
多个编号用英文符号的逗号隔开。最多支持50个经销采购单编号的查询。 */
func (this *FenxiaoDealerRequisitionorderQueryRequest) SetDealerOrderIds(value string) {
	this.Set("dealer_order_ids", value)
}

/* 多个字段用","分隔。 fields 如果为空：返回所有经销采购单对象(dealer_orders)字段。 如果不为空：返回指定采购单对象(dealer_orders)字段。 例1： dealer_order_details.product_id 表示只返回product_id 例2： dealer_order_details 表示只返回明细列表 */
func (this *FenxiaoDealerRequisitionorderQueryRequest) SetFields(value string) {
	this.Set("fields", value)
}

type FenxiaoDealerRequisitionorderQueryResponse struct {
	*ErrorResponse                                   `json:"error_response,omitempty"`
	FenxiaoDealerRequisitionorderQueryResponseResult `json:"fenxiao_dealer_requisitionorder_query_response"`
}
type FenxiaoDealerRequisitionorderQueryResponseResult struct {
	/* 经销采购单结果列表 */
	DealerOrders struct {
		DealerOrder []*DealerOrder `json:"dealer_order"`
	} `json:"DealerOrders"`
}

/*taobao.fenxiao.dealer.requisitionorder.refuse*/
type FenxiaoDealerRequisitionorderRefuseRequest struct {
	values url.Values
}

func (this *FenxiaoDealerRequisitionorderRefuseRequest) GetApiMethodName() string {
	return "taobao.fenxiao.dealer.requisitionorder.refuse"
}
func (this *FenxiaoDealerRequisitionorderRefuseRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *FenxiaoDealerRequisitionorderRefuseRequest) GetValues() url.Values {
	return this.values
}

/* 采购申请/经销采购单编号 */
func (this *FenxiaoDealerRequisitionorderRefuseRequest) SetDealerOrderId(value string) {
	this.Set("dealer_order_id", value)
}

/* 驳回原因（1：价格不合理；2：采购数量不合理；3：其他原因） */
func (this *FenxiaoDealerRequisitionorderRefuseRequest) SetReason(value string) {
	this.Set("reason", value)
}

/* 驳回详细原因，字数范围5-200字 */
func (this *FenxiaoDealerRequisitionorderRefuseRequest) SetReasonDetail(value string) {
	this.Set("reason_detail", value)
}

type FenxiaoDealerRequisitionorderRefuseResponse struct {
	*ErrorResponse                                    `json:"error_response,omitempty"`
	FenxiaoDealerRequisitionorderRefuseResponseResult `json:"fenxiao_dealer_requisitionorder_refuse_response"`
}
type FenxiaoDealerRequisitionorderRefuseResponseResult struct {
	/* 操作是否成功。true：成功；false：失败。 */
	IsSuccess bool `json:"is_success"`
}

/*taobao.fenxiao.dealer.requisitionorder.remark.update*/
type FenxiaoDealerRequisitionorderRemarkUpdateRequest struct {
	values url.Values
}

func (this *FenxiaoDealerRequisitionorderRemarkUpdateRequest) GetApiMethodName() string {
	return "taobao.fenxiao.dealer.requisitionorder.remark.update"
}
func (this *FenxiaoDealerRequisitionorderRemarkUpdateRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *FenxiaoDealerRequisitionorderRemarkUpdateRequest) GetValues() url.Values {
	return this.values
}

/* 经销采购单ID */
func (this *FenxiaoDealerRequisitionorderRemarkUpdateRequest) SetDealerOrderId(value string) {
	this.Set("dealer_order_id", value)
}

/* 备注留言，可为空 */
func (this *FenxiaoDealerRequisitionorderRemarkUpdateRequest) SetSupplierMemo(value string) {
	this.Set("supplier_memo", value)
}

/* 旗子的标记，必选。
1-5之间的数字。
非1-5之间，都采用1作为默认。
1:红色
2:黄色
3:绿色
4:蓝色
5:粉红色<br /> 支持最大值为：5<br /> 支持最小值为：1 */
func (this *FenxiaoDealerRequisitionorderRemarkUpdateRequest) SetSupplierMemoFlag(value string) {
	this.Set("supplier_memo_flag", value)
}

type FenxiaoDealerRequisitionorderRemarkUpdateResponse struct {
	*ErrorResponse                                          `json:"error_response,omitempty"`
	FenxiaoDealerRequisitionorderRemarkUpdateResponseResult `json:"fenxiao_dealer_requisitionorder_remark_update_response"`
}
type FenxiaoDealerRequisitionorderRemarkUpdateResponseResult struct {
	/* 操作是否成功 */
	IsSuccess bool `json:"is_success"`
}

/*taobao.fenxiao.discount.add*/
type FenxiaoDiscountAddRequest struct {
	values url.Values
}

func (this *FenxiaoDiscountAddRequest) GetApiMethodName() string {
	return "taobao.fenxiao.discount.add"
}
func (this *FenxiaoDiscountAddRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *FenxiaoDiscountAddRequest) GetValues() url.Values {
	return this.values
}

/* 折扣名称,长度不能超过25字节 */
func (this *FenxiaoDiscountAddRequest) SetDiscountName(value string) {
	this.Set("discount_name", value)
}

/* PERCENT（按折扣优惠）、PRICE（按减价优惠），例如"PERCENT,PRICE,PERCENT" */
func (this *FenxiaoDiscountAddRequest) SetDiscountTypes(value string) {
	this.Set("discount_types", value)
}

/* 优惠比率或者优惠价格，例如：”8000,-2300,7000”,大小为-100000000到100000000之间（单位：分） */
func (this *FenxiaoDiscountAddRequest) SetDiscountValues(value string) {
	this.Set("discount_values", value)
}

/* 会员等级的id或者分销商id，例如：”1001,2001,1002” */
func (this *FenxiaoDiscountAddRequest) SetTargetIds(value string) {
	this.Set("target_ids", value)
}

/* GRADE（按会员等级优惠）、DISTRIBUTOR（按分销商优惠），例如"GRADE,DISTRIBUTOR" */
func (this *FenxiaoDiscountAddRequest) SetTargetTypes(value string) {
	this.Set("target_types", value)
}

type FenxiaoDiscountAddResponse struct {
	*ErrorResponse                   `json:"error_response,omitempty"`
	FenxiaoDiscountAddResponseResult `json:"fenxiao_discount_add_response"`
}
type FenxiaoDiscountAddResponseResult struct {
	/* 折扣ID */
	DiscountId int64 `json:"discount_id"`
}

/*taobao.fenxiao.discount.update*/
type FenxiaoDiscountUpdateRequest struct {
	values url.Values
}

func (this *FenxiaoDiscountUpdateRequest) GetApiMethodName() string {
	return "taobao.fenxiao.discount.update"
}
func (this *FenxiaoDiscountUpdateRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *FenxiaoDiscountUpdateRequest) GetValues() url.Values {
	return this.values
}

/* 详情ID，例如：”0,1002,1003” */
func (this *FenxiaoDiscountUpdateRequest) SetDetailIds(value string) {
	this.Set("detail_ids", value)
}

/* ADD(新增)、UPDATE（更新）、DEL（删除，对应的target_type等信息填NULL），例如：”UPDATE,DEL,DEL” */
func (this *FenxiaoDiscountUpdateRequest) SetDetailStatuss(value string) {
	this.Set("detail_statuss", value)
}

/* 折扣ID */
func (this *FenxiaoDiscountUpdateRequest) SetDiscountId(value string) {
	this.Set("discount_id", value)
}

/* 折扣名称，长度不能超过25字节 */
func (this *FenxiaoDiscountUpdateRequest) SetDiscountName(value string) {
	this.Set("discount_name", value)
}

/* 状态DEL（删除）UPDATE(更新) */
func (this *FenxiaoDiscountUpdateRequest) SetDiscountStatus(value string) {
	this.Set("discount_status", value)
}

/* PERCENT（按折扣优惠）、PRICE（按减价优惠），例如"PERCENT,PRICE,PERCENT" */
func (this *FenxiaoDiscountUpdateRequest) SetDiscountTypes(value string) {
	this.Set("discount_types", value)
}

/* 优惠比率或者优惠价格，例如：”8000,-2300,7000”,大小为-100000000到100000000之间（单位：分） */
func (this *FenxiaoDiscountUpdateRequest) SetDiscountValues(value string) {
	this.Set("discount_values", value)
}

/* 会员等级的id或者分销商id，例如：”1001,2001,1002” */
func (this *FenxiaoDiscountUpdateRequest) SetTargetIds(value string) {
	this.Set("target_ids", value)
}

/* GRADE（按会员等级优惠）、DISTRIBUTOR（按分销商优惠），例如"GRADE,DISTRIBUTOR" */
func (this *FenxiaoDiscountUpdateRequest) SetTargetTypes(value string) {
	this.Set("target_types", value)
}

type FenxiaoDiscountUpdateResponse struct {
	*ErrorResponse                      `json:"error_response,omitempty"`
	FenxiaoDiscountUpdateResponseResult `json:"fenxiao_discount_update_response"`
}
type FenxiaoDiscountUpdateResponseResult struct {
	/* 成功状态 */
	IsSuccess bool `json:"is_success"`
}

/*taobao.fenxiao.discounts.get*/
type FenxiaoDiscountsGetRequest struct {
	values url.Values
}

func (this *FenxiaoDiscountsGetRequest) GetApiMethodName() string {
	return "taobao.fenxiao.discounts.get"
}
func (this *FenxiaoDiscountsGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *FenxiaoDiscountsGetRequest) GetValues() url.Values {
	return this.values
}

/* 折扣ID */
func (this *FenxiaoDiscountsGetRequest) SetDiscountId(value string) {
	this.Set("discount_id", value)
}

/* 指定查询额外的信息，可选值：DETAIL（查询折扣详情），多个可选值用逗号分割。（只允许指定折扣ID情况下使用） */
func (this *FenxiaoDiscountsGetRequest) SetExtFields(value string) {
	this.Set("ext_fields", value)
}

type FenxiaoDiscountsGetResponse struct {
	*ErrorResponse                    `json:"error_response,omitempty"`
	FenxiaoDiscountsGetResponseResult `json:"fenxiao_discounts_get_response"`
}
type FenxiaoDiscountsGetResponseResult struct {
	/* 折扣数据结构 */
	Discounts struct {
		Discount []*Discount `json:"discount"`
	} `json:"Discounts"`
	/* 记录数 */
	TotalResults int64 `json:"total_results"`
}

/*taobao.fenxiao.distributor.items.get*/
type FenxiaoDistributorItemsGetRequest struct {
	values url.Values
}

func (this *FenxiaoDistributorItemsGetRequest) GetApiMethodName() string {
	return "taobao.fenxiao.distributor.items.get"
}
func (this *FenxiaoDistributorItemsGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *FenxiaoDistributorItemsGetRequest) GetValues() url.Values {
	return this.values
}

/* 分销商ID 。 */
func (this *FenxiaoDistributorItemsGetRequest) SetDistributorId(value string) {
	this.Set("distributor_id", value)
}

/* 设置结束时间,空为不设置。 */
func (this *FenxiaoDistributorItemsGetRequest) SetEndModified(value string) {
	this.Set("end_modified", value)
}

/* 页码（大于0的整数，默认1） */
func (this *FenxiaoDistributorItemsGetRequest) SetPageNo(value string) {
	this.Set("page_no", value)
}

/* 每页记录数（默认20，最大50） */
func (this *FenxiaoDistributorItemsGetRequest) SetPageSize(value string) {
	this.Set("page_size", value)
}

/* 产品ID */
func (this *FenxiaoDistributorItemsGetRequest) SetProductId(value string) {
	this.Set("product_id", value)
}

/* 设置开始时间。空为不设置。 */
func (this *FenxiaoDistributorItemsGetRequest) SetStartModified(value string) {
	this.Set("start_modified", value)
}

type FenxiaoDistributorItemsGetResponse struct {
	*ErrorResponse                           `json:"error_response,omitempty"`
	FenxiaoDistributorItemsGetResponseResult `json:"fenxiao_distributor_items_get_response"`
}
type FenxiaoDistributorItemsGetResponseResult struct {
	/* 下载记录对象 */
	Records struct {
		FenxiaoItemRecord []*FenxiaoItemRecord `json:"fenxiao_item_record"`
	} `json:"Records"`
	/* 查询结果记录数 */
	TotalResults int64 `json:"total_results"`
}

/*taobao.fenxiao.distributor.products.get*/
type FenxiaoDistributorProductsGetRequest struct {
	values url.Values
}

func (this *FenxiaoDistributorProductsGetRequest) GetApiMethodName() string {
	return "taobao.fenxiao.distributor.products.get"
}
func (this *FenxiaoDistributorProductsGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *FenxiaoDistributorProductsGetRequest) GetValues() url.Values {
	return this.values
}

/* 下载状态，默认未下载。UNDOWNLOAD：未下载，DOWNLOADED：已下载。 */
func (this *FenxiaoDistributorProductsGetRequest) SetDownloadStatus(value string) {
	this.Set("download_status", value)
}

/* 结束时间 */
func (this *FenxiaoDistributorProductsGetRequest) SetEndTime(value string) {
	this.Set("end_time", value)
}

/* 指定查询额外的信息，可选值：skus（sku数据）、images（多图），多个可选值用逗号分割。 */
func (this *FenxiaoDistributorProductsGetRequest) SetFields(value string) {
	this.Set("fields", value)
}

/* 根据商品ID列表查询，优先级次于产品ID列表，高于其他分页查询条件。如果商品不是分销商品，自动过滤。最大限制20，用逗号分割，例如：“1001,1002,1003,1004,1005” */
func (this *FenxiaoDistributorProductsGetRequest) SetItemIds(value string) {
	this.Set("item_ids", value)
}

/* 排序。QUANTITY_DESC：库存降序，CREATE_TIME_DESC，创建时间降序。 */
func (this *FenxiaoDistributorProductsGetRequest) SetOrderBy(value string) {
	this.Set("order_by", value)
}

/* 页码（大于0的整数，默认1） */
func (this *FenxiaoDistributorProductsGetRequest) SetPageNo(value string) {
	this.Set("page_no", value)
}

/* 每页记录数（默认20，最大50） */
func (this *FenxiaoDistributorProductsGetRequest) SetPageSize(value string) {
	this.Set("page_size", value)
}

/* 产品ID列表，优先级最高，传了忽略其他查询条件。用逗号分割，例如：“1001,1002,1003,1004,1005” */
func (this *FenxiaoDistributorProductsGetRequest) SetPids(value string) {
	this.Set("pids", value)
}

/* 产品线ID */
func (this *FenxiaoDistributorProductsGetRequest) SetProductcatId(value string) {
	this.Set("productcat_id", value)
}

/* 开始时间 */
func (this *FenxiaoDistributorProductsGetRequest) SetStartTime(value string) {
	this.Set("start_time", value)
}

/* 供应商nick，分页查询时，必传 */
func (this *FenxiaoDistributorProductsGetRequest) SetSupplierNick(value string) {
	this.Set("supplier_nick", value)
}

/* 查询时间类型，默认更新时间。MODIFIED:更新时间，CREATE:创建时间 */
func (this *FenxiaoDistributorProductsGetRequest) SetTimeType(value string) {
	this.Set("time_type", value)
}

/* 分销方式，分页查询时，必传。AGENT：代销，DEALER：经销 */
func (this *FenxiaoDistributorProductsGetRequest) SetTradeType(value string) {
	this.Set("trade_type", value)
}

type FenxiaoDistributorProductsGetResponse struct {
	*ErrorResponse                              `json:"error_response,omitempty"`
	FenxiaoDistributorProductsGetResponseResult `json:"fenxiao_distributor_products_get_response"`
}
type FenxiaoDistributorProductsGetResponseResult struct {
	/* 是否存在下一页 */
	HasNext bool `json:"has_next"`
	/* 产品对象记录集。返回 FenxiaoProduct 包含的字段信息。 */
	Products struct {
		FenxiaoProduct []*FenxiaoProduct `json:"fenxiao_product"`
	} `json:"Products"`
}

/*taobao.fenxiao.distributors.get*/
type FenxiaoDistributorsGetRequest struct {
	values url.Values
}

func (this *FenxiaoDistributorsGetRequest) GetApiMethodName() string {
	return "taobao.fenxiao.distributors.get"
}
func (this *FenxiaoDistributorsGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *FenxiaoDistributorsGetRequest) GetValues() url.Values {
	return this.values
}

/* 分销商用户名列表。多个之间以“,”分隔;最多支持50个分销商用户名。 */
func (this *FenxiaoDistributorsGetRequest) SetNicks(value string) {
	this.Set("nicks", value)
}

type FenxiaoDistributorsGetResponse struct {
	*ErrorResponse                       `json:"error_response,omitempty"`
	FenxiaoDistributorsGetResponseResult `json:"fenxiao_distributors_get_response"`
}
type FenxiaoDistributorsGetResponseResult struct {
	/* 分销商详细信息 */
	Distributors struct {
		Distributor []*Distributor `json:"distributor"`
	} `json:"Distributors"`
}

/*taobao.fenxiao.grade.add*/
type FenxiaoGradeAddRequest struct {
	values url.Values
}

func (this *FenxiaoGradeAddRequest) GetApiMethodName() string {
	return "taobao.fenxiao.grade.add"
}
func (this *FenxiaoGradeAddRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *FenxiaoGradeAddRequest) GetValues() url.Values {
	return this.values
}

/* 等级名称，等级名称不可重复<br /> 支持最大长度为：20<br /> 支持的最大列表长度为：20 */
func (this *FenxiaoGradeAddRequest) SetName(value string) {
	this.Set("name", value)
}

type FenxiaoGradeAddResponse struct {
	*ErrorResponse                `json:"error_response,omitempty"`
	FenxiaoGradeAddResponseResult `json:"fenxiao_grade_add_response"`
}
type FenxiaoGradeAddResponseResult struct {
	/* 等级ID */
	GradeId int64 `json:"grade_id"`
	/* 操作是否成功 */
	IsSuccess bool `json:"is_success"`
}

/*taobao.fenxiao.grade.delete*/
type FenxiaoGradeDeleteRequest struct {
	values url.Values
}

func (this *FenxiaoGradeDeleteRequest) GetApiMethodName() string {
	return "taobao.fenxiao.grade.delete"
}
func (this *FenxiaoGradeDeleteRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *FenxiaoGradeDeleteRequest) GetValues() url.Values {
	return this.values
}

/* 等级ID */
func (this *FenxiaoGradeDeleteRequest) SetGradeId(value string) {
	this.Set("grade_id", value)
}

type FenxiaoGradeDeleteResponse struct {
	*ErrorResponse                   `json:"error_response,omitempty"`
	FenxiaoGradeDeleteResponseResult `json:"fenxiao_grade_delete_response"`
}
type FenxiaoGradeDeleteResponseResult struct {
	/* 操作是否成功 */
	IsSuccess bool `json:"is_success"`
}

/*taobao.fenxiao.grade.update*/
type FenxiaoGradeUpdateRequest struct {
	values url.Values
}

func (this *FenxiaoGradeUpdateRequest) GetApiMethodName() string {
	return "taobao.fenxiao.grade.update"
}
func (this *FenxiaoGradeUpdateRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *FenxiaoGradeUpdateRequest) GetValues() url.Values {
	return this.values
}

/* 等级ID */
func (this *FenxiaoGradeUpdateRequest) SetGradeId(value string) {
	this.Set("grade_id", value)
}

/* 等级名称，等级名称不可重复<br /> 支持最大长度为：20<br /> 支持的最大列表长度为：20 */
func (this *FenxiaoGradeUpdateRequest) SetName(value string) {
	this.Set("name", value)
}

type FenxiaoGradeUpdateResponse struct {
	*ErrorResponse                   `json:"error_response,omitempty"`
	FenxiaoGradeUpdateResponseResult `json:"fenxiao_grade_update_response"`
}
type FenxiaoGradeUpdateResponseResult struct {
	/* 操作是否成功 */
	IsSuccess bool `json:"is_success"`
}

/*taobao.fenxiao.grades.get*/
type FenxiaoGradesGetRequest struct {
	values url.Values
}

func (this *FenxiaoGradesGetRequest) GetApiMethodName() string {
	return "taobao.fenxiao.grades.get"
}
func (this *FenxiaoGradesGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *FenxiaoGradesGetRequest) GetValues() url.Values {
	return this.values
}

type FenxiaoGradesGetResponse struct {
	*ErrorResponse                 `json:"error_response,omitempty"`
	FenxiaoGradesGetResponseResult `json:"fenxiao_grades_get_response"`
}
type FenxiaoGradesGetResponseResult struct {
	/* 分销商等级信息 */
	FenxiaoGrades struct {
		FenxiaoGrade []*FenxiaoGrade `json:"fenxiao_grade"`
	} `json:"FenxiaoGrades"`
	/* 记录数 */
	TotalResults int64 `json:"total_results"`
}

/*taobao.fenxiao.login.user.get*/
type FenxiaoLoginUserGetRequest struct {
	values url.Values
}

func (this *FenxiaoLoginUserGetRequest) GetApiMethodName() string {
	return "taobao.fenxiao.login.user.get"
}
func (this *FenxiaoLoginUserGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *FenxiaoLoginUserGetRequest) GetValues() url.Values {
	return this.values
}

type FenxiaoLoginUserGetResponse struct {
	*ErrorResponse                    `json:"error_response,omitempty"`
	FenxiaoLoginUserGetResponseResult `json:"fenxiao_login_user_get_response"`
}
type FenxiaoLoginUserGetResponseResult struct {
	/* 登录用户信息 */
	LoginUser *LoginUser `json:"login_user"`
}

/*taobao.fenxiao.order.close*/
type FenxiaoOrderCloseRequest struct {
	values url.Values
}

func (this *FenxiaoOrderCloseRequest) GetApiMethodName() string {
	return "taobao.fenxiao.order.close"
}
func (this *FenxiaoOrderCloseRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *FenxiaoOrderCloseRequest) GetValues() url.Values {
	return this.values
}

/* 关闭理由,特殊字符会被转义，会改变长度，有特殊字符是请注意<br /> 支持最大长度为：200<br /> 支持的最大列表长度为：200 */
func (this *FenxiaoOrderCloseRequest) SetMessage(value string) {
	this.Set("message", value)
}

/* 采购单编号 */
func (this *FenxiaoOrderCloseRequest) SetPurchaseOrderId(value string) {
	this.Set("purchase_order_id", value)
}

/* 子采购单ID，可传多笔子单ID，逗号分隔 */
func (this *FenxiaoOrderCloseRequest) SetSubOrderIds(value string) {
	this.Set("sub_order_ids", value)
}

type FenxiaoOrderCloseResponse struct {
	*ErrorResponse                  `json:"error_response,omitempty"`
	FenxiaoOrderCloseResponseResult `json:"fenxiao_order_close_response"`
}
type FenxiaoOrderCloseResponseResult struct {
	/* 操作是否成功 */
	IsSuccess bool `json:"is_success"`
}

/*taobao.fenxiao.order.confirm.paid*/
type FenxiaoOrderConfirmPaidRequest struct {
	values url.Values
}

func (this *FenxiaoOrderConfirmPaidRequest) GetApiMethodName() string {
	return "taobao.fenxiao.order.confirm.paid"
}
func (this *FenxiaoOrderConfirmPaidRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *FenxiaoOrderConfirmPaidRequest) GetValues() url.Values {
	return this.values
}

/* 确认支付信息（字数小于100） */
func (this *FenxiaoOrderConfirmPaidRequest) SetConfirmRemark(value string) {
	this.Set("confirm_remark", value)
}

/* 采购单编号。 */
func (this *FenxiaoOrderConfirmPaidRequest) SetPurchaseOrderId(value string) {
	this.Set("purchase_order_id", value)
}

type FenxiaoOrderConfirmPaidResponse struct {
	*ErrorResponse                        `json:"error_response,omitempty"`
	FenxiaoOrderConfirmPaidResponseResult `json:"fenxiao_order_confirm_paid_response"`
}
type FenxiaoOrderConfirmPaidResponseResult struct {
	/* 确认结果成功与否 */
	IsSuccess bool `json:"is_success"`
}

/*taobao.fenxiao.order.create.dealer*/
type FenxiaoOrderCreateDealerRequest struct {
	values url.Values
}

func (this *FenxiaoOrderCreateDealerRequest) GetApiMethodName() string {
	return "taobao.fenxiao.order.create.dealer"
}
func (this *FenxiaoOrderCreateDealerRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *FenxiaoOrderCreateDealerRequest) GetValues() url.Values {
	return this.values
}

/* 街道 */
func (this *FenxiaoOrderCreateDealerRequest) SetAddr(value string) {
	this.Set("addr", value)
}

/* 买家姓名<br /> 支持最大长度为：20<br /> 支持的最大列表长度为：20 */
func (this *FenxiaoOrderCreateDealerRequest) SetBuyerName(value string) {
	this.Set("buyer_name", value)
}

/* 市 */
func (this *FenxiaoOrderCreateDealerRequest) SetCity(value string) {
	this.Set("city", value)
}

/* 区 */
func (this *FenxiaoOrderCreateDealerRequest) SetCountry(value string) {
	this.Set("country", value)
}

/* 运费，单位为分 */
func (this *FenxiaoOrderCreateDealerRequest) SetLogisticFee(value string) {
	this.Set("logistic_fee", value)
}

/* 运输方式，快递,平邮等 */
func (this *FenxiaoOrderCreateDealerRequest) SetLogisticType(value string) {
	this.Set("logistic_type", value)
}

/* 留言<br /> 支持最大长度为：200<br /> 支持的最大列表长度为：200 */
func (this *FenxiaoOrderCreateDealerRequest) SetMessage(value string) {
	this.Set("message", value)
}

/* 买家手机号码和电话号码两者中必须有一个 */
func (this *FenxiaoOrderCreateDealerRequest) SetMobilePhone(value string) {
	this.Set("mobile_phone", value)
}

/* erp主订单号，用于去重。当传入号已存在将返回原来的采购单 */
func (this *FenxiaoOrderCreateDealerRequest) SetOuterId(value string) {
	this.Set("outer_id", value)
}

/* 支付类型,需要供应支持该支付类型 */
func (this *FenxiaoOrderCreateDealerRequest) SetPayType(value string) {
	this.Set("pay_type", value)
}

/* 买家电话号码 */
func (this *FenxiaoOrderCreateDealerRequest) SetPhone(value string) {
	this.Set("phone", value)
}

/* 省 */
func (this *FenxiaoOrderCreateDealerRequest) SetProvince(value string) {
	this.Set("province", value)
}

/* 子单信息,子单内部以‘,’隔开，多个子单以‘;’隔开.
例(分销产品id,skuid,购买数量,单价;分销产品id:,skuid,购买数量,单价)
单价的单位位分 */
func (this *FenxiaoOrderCreateDealerRequest) SetSubOrderDetail(value string) {
	this.Set("sub_order_detail", value)
}

/* 邮编 */
func (this *FenxiaoOrderCreateDealerRequest) SetZipCode(value string) {
	this.Set("zip_code", value)
}

type FenxiaoOrderCreateDealerResponse struct {
	*ErrorResponse                         `json:"error_response,omitempty"`
	FenxiaoOrderCreateDealerResponseResult `json:"fenxiao_order_create_dealer_response"`
}
type FenxiaoOrderCreateDealerResponseResult struct {
	/* 采购单号 */
	GetModule int64 `json:"get_module"`
}

/*taobao.fenxiao.order.message.add*/
type FenxiaoOrderMessageAddRequest struct {
	values url.Values
}

func (this *FenxiaoOrderMessageAddRequest) GetApiMethodName() string {
	return "taobao.fenxiao.order.message.add"
}
func (this *FenxiaoOrderMessageAddRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *FenxiaoOrderMessageAddRequest) GetValues() url.Values {
	return this.values
}

/* 留言内容 */
func (this *FenxiaoOrderMessageAddRequest) SetMessage(value string) {
	this.Set("message", value)
}

/* 采购单id */
func (this *FenxiaoOrderMessageAddRequest) SetPurchaseOrderId(value string) {
	this.Set("purchase_order_id", value)
}

type FenxiaoOrderMessageAddResponse struct {
	*ErrorResponse                       `json:"error_response,omitempty"`
	FenxiaoOrderMessageAddResponseResult `json:"fenxiao_order_message_add_response"`
}
type FenxiaoOrderMessageAddResponseResult struct {
	/* 是否成功 */
	IsSuccess bool `json:"is_success"`
}

/*taobao.fenxiao.order.remark.update*/
type FenxiaoOrderRemarkUpdateRequest struct {
	values url.Values
}

func (this *FenxiaoOrderRemarkUpdateRequest) GetApiMethodName() string {
	return "taobao.fenxiao.order.remark.update"
}
func (this *FenxiaoOrderRemarkUpdateRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *FenxiaoOrderRemarkUpdateRequest) GetValues() url.Values {
	return this.values
}

/* 采购单编号 */
func (this *FenxiaoOrderRemarkUpdateRequest) SetPurchaseOrderId(value string) {
	this.Set("purchase_order_id", value)
}

/* 备注内容(供应商操作) */
func (this *FenxiaoOrderRemarkUpdateRequest) SetSupplierMemo(value string) {
	this.Set("supplier_memo", value)
}

/* 旗子的标记，1-5之间的数字。非1-5之间，都采用1作为默认。
1:红色
2:黄色
3:绿色
4:蓝色
5:粉红色<br /> 支持最大值为：5<br /> 支持最小值为：1 */
func (this *FenxiaoOrderRemarkUpdateRequest) SetSupplierMemoFlag(value string) {
	this.Set("supplier_memo_flag", value)
}

type FenxiaoOrderRemarkUpdateResponse struct {
	*ErrorResponse                         `json:"error_response,omitempty"`
	FenxiaoOrderRemarkUpdateResponseResult `json:"fenxiao_order_remark_update_response"`
}
type FenxiaoOrderRemarkUpdateResponseResult struct {
	/* 操作是否成功 */
	IsSuccess bool `json:"is_success"`
}

/*taobao.fenxiao.orders.get*/
type FenxiaoOrdersGetRequest struct {
	values url.Values
}

func (this *FenxiaoOrdersGetRequest) GetApiMethodName() string {
	return "taobao.fenxiao.orders.get"
}
func (this *FenxiaoOrdersGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *FenxiaoOrdersGetRequest) GetValues() url.Values {
	return this.values
}

/* 结束时间 格式 yyyy-MM-dd HH:mm:ss.支持到秒的查询。若不传时分秒，默认为0时0分0秒。若purchase_order_id没传，则此参数必传。 */
func (this *FenxiaoOrdersGetRequest) SetEndCreated(value string) {
	this.Set("end_created", value)
}

/* 多个字段用","分隔。

fields
如果为空：返回所有采购单对象(purchase_orders)字段。
如果不为空：返回指定采购单对象(purchase_orders)字段。

例1：
sub_purchase_orders.tc_order_id 表示只返回tc_order_id
例2：
sub_purchase_orders表示只返回子采购单列表 */
func (this *FenxiaoOrdersGetRequest) SetFields(value string) {
	this.Set("fields", value)
}

/* 页码。（大于0的整数。默认为1）<br /> 支持最大值为：2147483647<br /> 支持最小值为：-21474836478 */
func (this *FenxiaoOrdersGetRequest) SetPageNo(value string) {
	this.Set("page_no", value)
}

/* 每页条数。（每页条数不超过50条）<br /> 支持最大值为：2147483647<br /> 支持最小值为：-21474836478 */
func (this *FenxiaoOrdersGetRequest) SetPageSize(value string) {
	this.Set("page_size", value)
}

/* 采购单编号或分销流水号，若其它参数没传，则此参数必传。<br /> 支持最大值为：9223372036854775807<br /> 支持最小值为：-9223372036854775808 */
func (this *FenxiaoOrdersGetRequest) SetPurchaseOrderId(value string) {
	this.Set("purchase_order_id", value)
}

/* 起始时间 格式 yyyy-MM-dd HH:mm:ss.支持到秒的查询。若不传时分秒，默认为0时0分0秒。若purchase_order_id没传，则此参数必传。 */
func (this *FenxiaoOrdersGetRequest) SetStartCreated(value string) {
	this.Set("start_created", value)
}

/* 交易状态，不传默认查询所有采购单根据身份选择自身状态可选值:<br> *供应商：<br> WAIT_SELLER_SEND_GOODS(等待发货)<br> WAIT_SELLER_CONFIRM_PAY(待确认收款)<br> WAIT_BUYER_PAY(等待付款)<br> WAIT_BUYER_CONFIRM_GOODS(已发货)<br> TRADE_REFUNDING(退款中)<br> TRADE_FINISHED(采购成功)<br> TRADE_CLOSED(已关闭)<br> *分销商：<br> WAIT_BUYER_PAY(等待付款)<br> WAIT_BUYER_CONFIRM_GOODS(待收货确认)<br> TRADE_FOR_PAY(已付款)<br> TRADE_REFUNDING(退款中)<br> TRADE_FINISHED(采购成功)<br> TRADE_CLOSED(已关闭)<br> */
func (this *FenxiaoOrdersGetRequest) SetStatus(value string) {
	this.Set("status", value)
}

/* 采购单下游买家订单id */
func (this *FenxiaoOrdersGetRequest) SetTcOrderId(value string) {
	this.Set("tc_order_id", value)
}

/* 可选值：trade_time_type(采购单按照成交时间范围查询),update_time_type(采购单按照更新时间范围查询) */
func (this *FenxiaoOrdersGetRequest) SetTimeType(value string) {
	this.Set("time_type", value)
}

type FenxiaoOrdersGetResponse struct {
	*ErrorResponse                 `json:"error_response,omitempty"`
	FenxiaoOrdersGetResponseResult `json:"fenxiao_orders_get_response"`
}
type FenxiaoOrdersGetResponseResult struct {
	/* 采购单及子采购单信息。返回 PurchaseOrder 包含的字段信息。 */
	PurchaseOrders struct {
		PurchaseOrder []*PurchaseOrder `json:"purchase_order"`
	} `json:"PurchaseOrders"`
	/* 搜索到的采购单记录总数 */
	TotalResults int64 `json:"total_results"`
}

/*taobao.fenxiao.product.add*/
type FenxiaoProductAddRequest struct {
	values url.Values
}

func (this *FenxiaoProductAddRequest) GetApiMethodName() string {
	return "taobao.fenxiao.product.add"
}
func (this *FenxiaoProductAddRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *FenxiaoProductAddRequest) GetValues() url.Values {
	return this.values
}

/* 警戒库存必须是0到29999。 */
func (this *FenxiaoProductAddRequest) SetAlarmNumber(value string) {
	this.Set("alarm_number", value)
}

/* 所属类目id，参考Taobao.itemcats.get，不支持成人等类目，输入成人类目id保存提示类目属性错误。 */
func (this *FenxiaoProductAddRequest) SetCategoryId(value string) {
	this.Set("category_id", value)
}

/* 所在地：市，例：“杭州” */
func (this *FenxiaoProductAddRequest) SetCity(value string) {
	this.Set("city", value)
}

/* 代销采购价格，单位：元。例：“10.56”。必须在0.01元到10000000元之间。 */
func (this *FenxiaoProductAddRequest) SetCostPrice(value string) {
	this.Set("cost_price", value)
}

/* 经销采购价，单位：元。例：“10.56”。必须在0.01元到10000000元之间。 */
func (this *FenxiaoProductAddRequest) SetDealerCostPrice(value string) {
	this.Set("dealer_cost_price", value)
}

/* 产品描述，长度为5到25000字符。 */
func (this *FenxiaoProductAddRequest) SetDesc(value string) {
	this.Set("desc", value)
}

/* 折扣ID */
func (this *FenxiaoProductAddRequest) SetDiscountId(value string) {
	this.Set("discount_id", value)
}

/* 是否有保修，可选值：false（否）、true（是），默认false。 */
func (this *FenxiaoProductAddRequest) SetHaveGuarantee(value string) {
	this.Set("have_guarantee", value)
}

/* 是否有发票，可选值：false（否）、true（是），默认false。 */
func (this *FenxiaoProductAddRequest) SetHaveInvoice(value string) {
	this.Set("have_invoice", value)
}

/* 产品主图，大小不超过500k，格式为gif,jpg,jpeg,png,bmp等图片 */
func (this *FenxiaoProductAddRequest) SetImage(value string) {
	this.Set("image", value)
}

/* 自定义属性。格式为pid:value;pid:value */
func (this *FenxiaoProductAddRequest) SetInputProperties(value string) {
	this.Set("input_properties", value)
}

/* 添加产品时，添加入参isAuthz:yes|no
yes:需要授权
no:不需要授权
默认是需要授权 */
func (this *FenxiaoProductAddRequest) SetIsAuthz(value string) {
	this.Set("is_authz", value)
}

/* 导入的商品ID */
func (this *FenxiaoProductAddRequest) SetItemId(value string) {
	this.Set("item_id", value)
}

/* 产品名称，长度不超过60个字节。 */
func (this *FenxiaoProductAddRequest) SetName(value string) {
	this.Set("name", value)
}

/* 商家编码，长度不能超过60个字节。 */
func (this *FenxiaoProductAddRequest) SetOuterId(value string) {
	this.Set("outer_id", value)
}

/* 产品主图图片空间相对路径或绝对路径 */
func (this *FenxiaoProductAddRequest) SetPicPath(value string) {
	this.Set("pic_path", value)
}

/* ems费用，单位：元。例：“10.56”。 大小为0.00元到999999元之间。 */
func (this *FenxiaoProductAddRequest) SetPostageEms(value string) {
	this.Set("postage_ems", value)
}

/* 快递费用，单位：元。例：“10.56”。 大小为0.01元到999999元之间。 */
func (this *FenxiaoProductAddRequest) SetPostageFast(value string) {
	this.Set("postage_fast", value)
}

/* 运费模板ID，参考taobao.postages.get。 */
func (this *FenxiaoProductAddRequest) SetPostageId(value string) {
	this.Set("postage_id", value)
}

/* 平邮费用，单位：元。例：“10.56”。 大小为0.01元到999999元之间。 */
func (this *FenxiaoProductAddRequest) SetPostageOrdinary(value string) {
	this.Set("postage_ordinary", value)
}

/* 运费类型，可选值：seller（供应商承担运费）、buyer（分销商承担运费）,默认seller。 */
func (this *FenxiaoProductAddRequest) SetPostageType(value string) {
	this.Set("postage_type", value)
}

/* 产品线ID */
func (this *FenxiaoProductAddRequest) SetProductcatId(value string) {
	this.Set("productcat_id", value)
}

/* 产品属性，格式为pid:vid;pid:vid */
func (this *FenxiaoProductAddRequest) SetProperties(value string) {
	this.Set("properties", value)
}

/* 属性别名，格式为：pid:vid:alias;pid:vid:alias（alias为别名） */
func (this *FenxiaoProductAddRequest) SetPropertyAlias(value string) {
	this.Set("property_alias", value)
}

/* 所在地：省，例：“浙江” */
func (this *FenxiaoProductAddRequest) SetProv(value string) {
	this.Set("prov", value)
}

/* 产品库存必须是1到999999。 */
func (this *FenxiaoProductAddRequest) SetQuantity(value string) {
	this.Set("quantity", value)
}

/* 最高零售价，单位：元。例：“10.56”。必须在0.01元到10000000元之间，最高零售价必须大于最低零售价。 */
func (this *FenxiaoProductAddRequest) SetRetailPriceHigh(value string) {
	this.Set("retail_price_high", value)
}

/* 最低零售价，单位：元。例：“10.56”。必须在0.01元到10000000元之间。 */
func (this *FenxiaoProductAddRequest) SetRetailPriceLow(value string) {
	this.Set("retail_price_low", value)
}

/* sku的采购价。如果多个，用逗号分隔，并与其他sku信息保持相同顺序 */
func (this *FenxiaoProductAddRequest) SetSkuCostPrices(value string) {
	this.Set("sku_cost_prices", value)
}

/* sku的经销采购价。如果多个，用逗号分隔，并与其他sku信息保持相同顺序。其中每个值的单位：元。例：“10.56,12.3”。必须在0.01元到10000000元之间。 */
func (this *FenxiaoProductAddRequest) SetSkuDealerCostPrices(value string) {
	this.Set("sku_dealer_cost_prices", value)
}

/* sku的商家编码。如果多个，用逗号分隔，并与其他sku信息保持相同顺序 */
func (this *FenxiaoProductAddRequest) SetSkuOuterIds(value string) {
	this.Set("sku_outer_ids", value)
}

/* sku的属性。如果多个，用逗号分隔，并与其他sku信息保持相同顺序 */
func (this *FenxiaoProductAddRequest) SetSkuProperties(value string) {
	this.Set("sku_properties", value)
}

/* sku的库存。如果多个，用逗号分隔，并与其他sku信息保持相同顺序 */
func (this *FenxiaoProductAddRequest) SetSkuQuantitys(value string) {
	this.Set("sku_quantitys", value)
}

/* sku的采购基准价。如果多个，用逗号分隔，并与其他sku信息保持相同顺序 */
func (this *FenxiaoProductAddRequest) SetSkuStandardPrices(value string) {
	this.Set("sku_standard_prices", value)
}

/* 产品spuID，达尔文产品必须要传spuID，否则不能发布。其他非达尔文产品，看情况传 */
func (this *FenxiaoProductAddRequest) SetSpuId(value string) {
	this.Set("spu_id", value)
}

/* 采购基准价格，单位：元。例：“10.56”。必须在0.01元到10000000元之间。 */
func (this *FenxiaoProductAddRequest) SetStandardPrice(value string) {
	this.Set("standard_price", value)
}

/* 零售基准价，单位：元。例：“10.56”。必须在0.01元到10000000元之间。 */
func (this *FenxiaoProductAddRequest) SetStandardRetailPrice(value string) {
	this.Set("standard_retail_price", value)
}

/* 分销方式：AGENT（只做代销，默认值）、DEALER（只做经销）、ALL（代销和经销都做） */
func (this *FenxiaoProductAddRequest) SetTradeType(value string) {
	this.Set("trade_type", value)
}

type FenxiaoProductAddResponse struct {
	*ErrorResponse                  `json:"error_response,omitempty"`
	FenxiaoProductAddResponseResult `json:"fenxiao_product_add_response"`
}
type FenxiaoProductAddResponseResult struct {
	/* 产品创建时间 时间格式：yyyy-MM-dd HH:mm:ss */
	Created string `json:"created"`
	/* 产品ID */
	Pid int64 `json:"pid"`
}

/*taobao.fenxiao.product.gradeprice.get*/
type FenxiaoProductGradepriceGetRequest struct {
	values url.Values
}

func (this *FenxiaoProductGradepriceGetRequest) GetApiMethodName() string {
	return "taobao.fenxiao.product.gradeprice.get"
}
func (this *FenxiaoProductGradepriceGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *FenxiaoProductGradepriceGetRequest) GetValues() url.Values {
	return this.values
}

/* 产品id */
func (this *FenxiaoProductGradepriceGetRequest) SetProductId(value string) {
	this.Set("product_id", value)
}

/* skuId */
func (this *FenxiaoProductGradepriceGetRequest) SetSkuId(value string) {
	this.Set("sku_id", value)
}

/* 经、代销模式（1：代销、2：经销） */
func (this *FenxiaoProductGradepriceGetRequest) SetTradeType(value string) {
	this.Set("trade_type", value)
}

type FenxiaoProductGradepriceGetResponse struct {
	*ErrorResponse                            `json:"error_response,omitempty"`
	FenxiaoProductGradepriceGetResponseResult `json:"fenxiao_product_gradeprice_get_response"`
}
type FenxiaoProductGradepriceGetResponseResult struct {
	/* 等级折扣列表 */
	GradeDiscounts struct {
		GradeDiscount []*GradeDiscount `json:"grade_discount"`
	} `json:"GradeDiscounts"`
	/* 操作是否成功 */
	IsSuccess bool `json:"is_success"`
}

/*taobao.fenxiao.product.gradeprice.update*/
type FenxiaoProductGradepriceUpdateRequest struct {
	values url.Values
}

func (this *FenxiaoProductGradepriceUpdateRequest) GetApiMethodName() string {
	return "taobao.fenxiao.product.gradeprice.update"
}
func (this *FenxiaoProductGradepriceUpdateRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *FenxiaoProductGradepriceUpdateRequest) GetValues() url.Values {
	return this.values
}

/* 会员等级的id或者分销商id，例如：”1001,2001,1002” */
func (this *FenxiaoProductGradepriceUpdateRequest) SetIds(value string) {
	this.Set("ids", value)
}

/* 优惠价格,大小为0到100000000之间的整数或两位小数，例：优惠价格为：100元2角5分，传入的参数应写成：100.25 */
func (this *FenxiaoProductGradepriceUpdateRequest) SetPrices(value string) {
	this.Set("prices", value)
}

/* 产品Id */
func (this *FenxiaoProductGradepriceUpdateRequest) SetProductId(value string) {
	this.Set("product_id", value)
}

/* skuId，如果产品有skuId,必须要输入skuId;没有skuId的时候不必选 */
func (this *FenxiaoProductGradepriceUpdateRequest) SetSkuId(value string) {
	this.Set("sku_id", value)
}

/* 选择折扣方式：GRADE（按等级进行设置）;DISCITUTOR(按分销商进行设置）。例如"GRADE,DISTRIBUTOR" */
func (this *FenxiaoProductGradepriceUpdateRequest) SetTargetType(value string) {
	this.Set("target_type", value)
}

/* 交易类型： AGENT(代销)、DEALER(经销)，ALL(代销和经销) */
func (this *FenxiaoProductGradepriceUpdateRequest) SetTradeType(value string) {
	this.Set("trade_type", value)
}

type FenxiaoProductGradepriceUpdateResponse struct {
	*ErrorResponse                               `json:"error_response,omitempty"`
	FenxiaoProductGradepriceUpdateResponseResult `json:"fenxiao_product_gradeprice_update_response"`
}
type FenxiaoProductGradepriceUpdateResponseResult struct {
	/* 返回操作结果：成功或失败 */
	IsSuccess bool `json:"is_success"`
}

/*taobao.fenxiao.product.image.delete*/
type FenxiaoProductImageDeleteRequest struct {
	values url.Values
}

func (this *FenxiaoProductImageDeleteRequest) GetApiMethodName() string {
	return "taobao.fenxiao.product.image.delete"
}
func (this *FenxiaoProductImageDeleteRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *FenxiaoProductImageDeleteRequest) GetValues() url.Values {
	return this.values
}

/* 图片位置 */
func (this *FenxiaoProductImageDeleteRequest) SetPosition(value string) {
	this.Set("position", value)
}

/* 产品ID */
func (this *FenxiaoProductImageDeleteRequest) SetProductId(value string) {
	this.Set("product_id", value)
}

/* properties表示sku图片的属性。key:value形式，key是pid，value是vid。如果position是0的话，则properties需要是必传项 */
func (this *FenxiaoProductImageDeleteRequest) SetProperties(value string) {
	this.Set("properties", value)
}

type FenxiaoProductImageDeleteResponse struct {
	*ErrorResponse                          `json:"error_response,omitempty"`
	FenxiaoProductImageDeleteResponseResult `json:"fenxiao_product_image_delete_response"`
}
type FenxiaoProductImageDeleteResponseResult struct {
	/* 操作时间 */
	Created string `json:"created"`
	/* 操作结果 */
	Result bool `json:"result"`
}

/*taobao.fenxiao.product.image.upload*/
type FenxiaoProductImageUploadRequest struct {
	values url.Values
}

func (this *FenxiaoProductImageUploadRequest) GetApiMethodName() string {
	return "taobao.fenxiao.product.image.upload"
}
func (this *FenxiaoProductImageUploadRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *FenxiaoProductImageUploadRequest) GetValues() url.Values {
	return this.values
}

/* 产品图片 */
func (this *FenxiaoProductImageUploadRequest) SetImage(value string) {
	this.Set("image", value)
}

/* 产品主图图片空间相对路径或绝对路径 */
func (this *FenxiaoProductImageUploadRequest) SetPicPath(value string) {
	this.Set("pic_path", value)
}

/* 图片位置，0-14之间。0：操作sku属性图片，1：主图，2-5：细节图，6-14：额外主图 */
func (this *FenxiaoProductImageUploadRequest) SetPosition(value string) {
	this.Set("position", value)
}

/* 产品ID */
func (this *FenxiaoProductImageUploadRequest) SetProductId(value string) {
	this.Set("product_id", value)
}

/* properties表示sku图片的属性。key:value形式，key是pid，value是vid。如果position是0的话，则properties需要是必传项 */
func (this *FenxiaoProductImageUploadRequest) SetProperties(value string) {
	this.Set("properties", value)
}

type FenxiaoProductImageUploadResponse struct {
	*ErrorResponse                          `json:"error_response,omitempty"`
	FenxiaoProductImageUploadResponseResult `json:"fenxiao_product_image_upload_response"`
}
type FenxiaoProductImageUploadResponseResult struct {
	/* 操作时间 */
	Created string `json:"created"`
	/* 操作是否成功 */
	Result bool `json:"result"`
}

/*taobao.fenxiao.product.map.add*/
type FenxiaoProductMapAddRequest struct {
	values url.Values
}

func (this *FenxiaoProductMapAddRequest) GetApiMethodName() string {
	return "taobao.fenxiao.product.map.add"
}
func (this *FenxiaoProductMapAddRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *FenxiaoProductMapAddRequest) GetValues() url.Values {
	return this.values
}

/* 是否需要校验商家编码，true不校验，false校验。 */
func (this *FenxiaoProductMapAddRequest) SetNotCheckOuterCode(value string) {
	this.Set("not_check_outer_code", value)
}

/* 分销产品id。 */
func (this *FenxiaoProductMapAddRequest) SetProductId(value string) {
	this.Set("product_id", value)
}

/* 后端商品id（如果当前分销产品没有sku和后端商品时需要指定）。 */
func (this *FenxiaoProductMapAddRequest) SetScItemId(value string) {
	this.Set("sc_item_id", value)
}

/* 在有sku的情况下，与各个sku对应的后端商品id列表。逗号分隔，顺序需要保证与sku_ids一致。 */
func (this *FenxiaoProductMapAddRequest) SetScItemIds(value string) {
	this.Set("sc_item_ids", value)
}

/* 分销产品的sku id。逗号分隔，顺序需要保证与sc_item_ids一致（没有sku就不传）。 */
func (this *FenxiaoProductMapAddRequest) SetSkuIds(value string) {
	this.Set("sku_ids", value)
}

type FenxiaoProductMapAddResponse struct {
	*ErrorResponse                     `json:"error_response,omitempty"`
	FenxiaoProductMapAddResponseResult `json:"fenxiao_product_map_add_response"`
}
type FenxiaoProductMapAddResponseResult struct {
	/* 操作结果 */
	IsSuccess bool `json:"is_success"`
}

/*taobao.fenxiao.product.map.delete*/
type FenxiaoProductMapDeleteRequest struct {
	values url.Values
}

func (this *FenxiaoProductMapDeleteRequest) GetApiMethodName() string {
	return "taobao.fenxiao.product.map.delete"
}
func (this *FenxiaoProductMapDeleteRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *FenxiaoProductMapDeleteRequest) GetValues() url.Values {
	return this.values
}

/* 分销产品id。 */
func (this *FenxiaoProductMapDeleteRequest) SetProductId(value string) {
	this.Set("product_id", value)
}

/* 分销产品的sku id列表，逗号分隔，在有sku时需要指定。 */
func (this *FenxiaoProductMapDeleteRequest) SetSkuIds(value string) {
	this.Set("sku_ids", value)
}

type FenxiaoProductMapDeleteResponse struct {
	*ErrorResponse                        `json:"error_response,omitempty"`
	FenxiaoProductMapDeleteResponseResult `json:"fenxiao_product_map_delete_response"`
}
type FenxiaoProductMapDeleteResponseResult struct {
	/* 操作结果 */
	IsSuccess bool `json:"is_success"`
}

/*taobao.fenxiao.product.quantity.update*/
type FenxiaoProductQuantityUpdateRequest struct {
	values url.Values
}

func (this *FenxiaoProductQuantityUpdateRequest) GetApiMethodName() string {
	return "taobao.fenxiao.product.quantity.update"
}
func (this *FenxiaoProductQuantityUpdateRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *FenxiaoProductQuantityUpdateRequest) GetValues() url.Values {
	return this.values
}

/* 产品ID */
func (this *FenxiaoProductQuantityUpdateRequest) SetProductId(value string) {
	this.Set("product_id", value)
}

/* sku属性值，产品有sku时填写，多个sku用,分隔。为空时默认该产品无sku，则只修改产品的库存。 */
func (this *FenxiaoProductQuantityUpdateRequest) SetProperties(value string) {
	this.Set("properties", value)
}

/* 库存修改值。产品有sku时，与sku属性顺序对应，用,分隔。产品无sku时，只写库存值。
当全量更新库存时，quantity必须为大于等于0的正整数；当增量更新库存时，quantity为整数，可小于等于0。若增量更新时传入的库存为负数，则负数与实际库存之和不能小于0。比如当前实际库存为1，传入增量更新quantity=-1，库存改为0 */
func (this *FenxiaoProductQuantityUpdateRequest) SetQuantity(value string) {
	this.Set("quantity", value)
}

/* 库存更新方式，可选。1为全量更新，2为增量更新。如果不填，默认为全量更新。当选择全量更新时，如果库存更新值传入的是负数，会出错并返回错误码；当选择增量更新时，如果库存更新值为负数且绝对值大于当前库存，则sku库存会设置为0 */
func (this *FenxiaoProductQuantityUpdateRequest) SetType(value string) {
	this.Set("type", value)
}

type FenxiaoProductQuantityUpdateResponse struct {
	*ErrorResponse                             `json:"error_response,omitempty"`
	FenxiaoProductQuantityUpdateResponseResult `json:"fenxiao_product_quantity_update_response"`
}
type FenxiaoProductQuantityUpdateResponseResult struct {
	/* 操作时间 */
	Created string `json:"created"`
	/* 操作结果 */
	Result bool `json:"result"`
}

/*taobao.fenxiao.product.sku.add*/
type FenxiaoProductSkuAddRequest struct {
	values url.Values
}

func (this *FenxiaoProductSkuAddRequest) GetApiMethodName() string {
	return "taobao.fenxiao.product.sku.add"
}
func (this *FenxiaoProductSkuAddRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *FenxiaoProductSkuAddRequest) GetValues() url.Values {
	return this.values
}

/* 代销采购价 */
func (this *FenxiaoProductSkuAddRequest) SetAgentCostPrice(value string) {
	this.Set("agent_cost_price", value)
}

/* 经销采购价 */
func (this *FenxiaoProductSkuAddRequest) SetDealerCostPrice(value string) {
	this.Set("dealer_cost_price", value)
}

/* 产品ID<br /> 支持的最大列表长度为：64 */
func (this *FenxiaoProductSkuAddRequest) SetProductId(value string) {
	this.Set("product_id", value)
}

/* sku属性 */
func (this *FenxiaoProductSkuAddRequest) SetProperties(value string) {
	this.Set("properties", value)
}

/* sku产品库存，在0到1000000之间，如果不传，则库存为0<br /> 支持最大值为：1000000<br /> 支持最小值为：0 */
func (this *FenxiaoProductSkuAddRequest) SetQuantity(value string) {
	this.Set("quantity", value)
}

/* 商家编码 */
func (this *FenxiaoProductSkuAddRequest) SetSkuNumber(value string) {
	this.Set("sku_number", value)
}

/* 采购基准价，最大值1000000000 */
func (this *FenxiaoProductSkuAddRequest) SetStandardPrice(value string) {
	this.Set("standard_price", value)
}

type FenxiaoProductSkuAddResponse struct {
	*ErrorResponse                     `json:"error_response,omitempty"`
	FenxiaoProductSkuAddResponseResult `json:"fenxiao_product_sku_add_response"`
}
type FenxiaoProductSkuAddResponseResult struct {
	/* 操作时间 */
	Created string `json:"created"`
	/* 操作结果 */
	Result bool `json:"result"`
}

/*taobao.fenxiao.product.sku.delete*/
type FenxiaoProductSkuDeleteRequest struct {
	values url.Values
}

func (this *FenxiaoProductSkuDeleteRequest) GetApiMethodName() string {
	return "taobao.fenxiao.product.sku.delete"
}
func (this *FenxiaoProductSkuDeleteRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *FenxiaoProductSkuDeleteRequest) GetValues() url.Values {
	return this.values
}

/* 产品id */
func (this *FenxiaoProductSkuDeleteRequest) SetProductId(value string) {
	this.Set("product_id", value)
}

/* sku属性 */
func (this *FenxiaoProductSkuDeleteRequest) SetProperties(value string) {
	this.Set("properties", value)
}

type FenxiaoProductSkuDeleteResponse struct {
	*ErrorResponse                        `json:"error_response,omitempty"`
	FenxiaoProductSkuDeleteResponseResult `json:"fenxiao_product_sku_delete_response"`
}
type FenxiaoProductSkuDeleteResponseResult struct {
	/* 操作时间 */
	Created string `json:"created"`
	/* 操作结果 */
	Result bool `json:"result"`
}

/*taobao.fenxiao.product.sku.update*/
type FenxiaoProductSkuUpdateRequest struct {
	values url.Values
}

func (this *FenxiaoProductSkuUpdateRequest) GetApiMethodName() string {
	return "taobao.fenxiao.product.sku.update"
}
func (this *FenxiaoProductSkuUpdateRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *FenxiaoProductSkuUpdateRequest) GetValues() url.Values {
	return this.values
}

/* 代销采购价 */
func (this *FenxiaoProductSkuUpdateRequest) SetAgentCostPrice(value string) {
	this.Set("agent_cost_price", value)
}

/* 经销采购价 */
func (this *FenxiaoProductSkuUpdateRequest) SetDealerCostPrice(value string) {
	this.Set("dealer_cost_price", value)
}

/* 产品ID */
func (this *FenxiaoProductSkuUpdateRequest) SetProductId(value string) {
	this.Set("product_id", value)
}

/* sku属性 */
func (this *FenxiaoProductSkuUpdateRequest) SetProperties(value string) {
	this.Set("properties", value)
}

/* 产品SKU库存 */
func (this *FenxiaoProductSkuUpdateRequest) SetQuantity(value string) {
	this.Set("quantity", value)
}

/* 商家编码 */
func (this *FenxiaoProductSkuUpdateRequest) SetSkuNumber(value string) {
	this.Set("sku_number", value)
}

/* 采购基准价 */
func (this *FenxiaoProductSkuUpdateRequest) SetStandardPrice(value string) {
	this.Set("standard_price", value)
}

type FenxiaoProductSkuUpdateResponse struct {
	*ErrorResponse                        `json:"error_response,omitempty"`
	FenxiaoProductSkuUpdateResponseResult `json:"fenxiao_product_sku_update_response"`
}
type FenxiaoProductSkuUpdateResponseResult struct {
	/* 操作时间 */
	Created string `json:"created"`
	/* 操作结果 */
	Result bool `json:"result"`
}

/*taobao.fenxiao.product.skus.get*/
type FenxiaoProductSkusGetRequest struct {
	values url.Values
}

func (this *FenxiaoProductSkusGetRequest) GetApiMethodName() string {
	return "taobao.fenxiao.product.skus.get"
}
func (this *FenxiaoProductSkusGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *FenxiaoProductSkusGetRequest) GetValues() url.Values {
	return this.values
}

/* 产品ID<br /> 支持的最大列表长度为：64 */
func (this *FenxiaoProductSkusGetRequest) SetProductId(value string) {
	this.Set("product_id", value)
}

type FenxiaoProductSkusGetResponse struct {
	*ErrorResponse                      `json:"error_response,omitempty"`
	FenxiaoProductSkusGetResponseResult `json:"fenxiao_product_skus_get_response"`
}
type FenxiaoProductSkusGetResponseResult struct {
	/* sku信息 */
	Skus struct {
		FenxiaoSku []*FenxiaoSku `json:"fenxiao_sku"`
	} `json:"Skus"`
	/* 记录数量 */
	TotalResults int64 `json:"total_results"`
}

/*taobao.fenxiao.product.update*/
type FenxiaoProductUpdateRequest struct {
	values url.Values
}

func (this *FenxiaoProductUpdateRequest) GetApiMethodName() string {
	return "taobao.fenxiao.product.update"
}
func (this *FenxiaoProductUpdateRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *FenxiaoProductUpdateRequest) GetValues() url.Values {
	return this.values
}

/* 警戒库存必须是0到29999。 */
func (this *FenxiaoProductUpdateRequest) SetAlarmNumber(value string) {
	this.Set("alarm_number", value)
}

/* 所属类目id，参考Taobao.itemcats.get，不支持成人等类目，输入成人类目id保存提示类目属性错误。 */
func (this *FenxiaoProductUpdateRequest) SetCategoryId(value string) {
	this.Set("category_id", value)
}

/* 所在地：市，例：“杭州” */
func (this *FenxiaoProductUpdateRequest) SetCity(value string) {
	this.Set("city", value)
}

/* 代销采购价格，单位：元。例：“10.56”。必须在0.01元到10000000元之间。 */
func (this *FenxiaoProductUpdateRequest) SetCostPrice(value string) {
	this.Set("cost_price", value)
}

/* 经销采购价，单位：元。例：“10.56”。必须在0.01元到10000000元之间。 */
func (this *FenxiaoProductUpdateRequest) SetDealerCostPrice(value string) {
	this.Set("dealer_cost_price", value)
}

/* 产品描述，长度为5到25000字符。 */
func (this *FenxiaoProductUpdateRequest) SetDesc(value string) {
	this.Set("desc", value)
}

/* 折扣ID */
func (this *FenxiaoProductUpdateRequest) SetDiscountId(value string) {
	this.Set("discount_id", value)
}

/* 是否有保修，可选值：false（否）、true（是），默认false。 */
func (this *FenxiaoProductUpdateRequest) SetHaveGuarantee(value string) {
	this.Set("have_guarantee", value)
}

/* 是否有发票，可选值：false（否）、true（是），默认false。 */
func (this *FenxiaoProductUpdateRequest) SetHaveInvoice(value string) {
	this.Set("have_invoice", value)
}

/* 主图图片，如果pic_path参数不空，则优先使用pic_path，忽略该参数<br /> 支持的文件类型为：gif,jpg,jpeg,png,bmp<br /> 支持的最大列表长度为：512000 */
func (this *FenxiaoProductUpdateRequest) SetImage(value string) {
	this.Set("image", value)
}

/* 自定义属性。格式为pid:value;pid:value */
func (this *FenxiaoProductUpdateRequest) SetInputProperties(value string) {
	this.Set("input_properties", value)
}

/* 产品是否需要授权isAuthz:yes|no
yes:需要授权
no:不需要授权 */
func (this *FenxiaoProductUpdateRequest) SetIsAuthz(value string) {
	this.Set("is_authz", value)
}

/* 产品名称，长度不超过60个字节。 */
func (this *FenxiaoProductUpdateRequest) SetName(value string) {
	this.Set("name", value)
}

/* 商家编码，长度不能超过60个字节。 */
func (this *FenxiaoProductUpdateRequest) SetOuterId(value string) {
	this.Set("outer_id", value)
}

/* 产品主图图片空间相对路径或绝对路径 */
func (this *FenxiaoProductUpdateRequest) SetPicPath(value string) {
	this.Set("pic_path", value)
}

/* 产品ID */
func (this *FenxiaoProductUpdateRequest) SetPid(value string) {
	this.Set("pid", value)
}

/* ems费用，单位：元。例：“10.56”。大小为0.01元到999999元之间。更新时必须指定运费类型为buyer，否则不更新。 */
func (this *FenxiaoProductUpdateRequest) SetPostageEms(value string) {
	this.Set("postage_ems", value)
}

/* 快递费用，单位：元。例：“10.56”。大小为0.01元到999999元之间。更新时必须指定运费类型为buyer，否则不更新。 */
func (this *FenxiaoProductUpdateRequest) SetPostageFast(value string) {
	this.Set("postage_fast", value)
}

/* 运费模板ID，参考taobao.postages.get。更新时必须指定运费类型为 buyer，否则不更新。 */
func (this *FenxiaoProductUpdateRequest) SetPostageId(value string) {
	this.Set("postage_id", value)
}

/* 平邮费用，单位：元。例：“10.56”。大小为0.01元到999999元之间。更新时必须指定运费类型为buyer，否则不更新。 */
func (this *FenxiaoProductUpdateRequest) SetPostageOrdinary(value string) {
	this.Set("postage_ordinary", value)
}

/* 运费类型，可选值：seller（供应商承担运费）、buyer（分销商承担运费）。 */
func (this *FenxiaoProductUpdateRequest) SetPostageType(value string) {
	this.Set("postage_type", value)
}

/* 产品属性 */
func (this *FenxiaoProductUpdateRequest) SetProperties(value string) {
	this.Set("properties", value)
}

/* 属性别名 */
func (this *FenxiaoProductUpdateRequest) SetPropertyAlias(value string) {
	this.Set("property_alias", value)
}

/* 所在地：省，例：“浙江” */
func (this *FenxiaoProductUpdateRequest) SetProv(value string) {
	this.Set("prov", value)
}

/* 产品库存必须是1到999999。 */
func (this *FenxiaoProductUpdateRequest) SetQuantity(value string) {
	this.Set("quantity", value)
}

/* 最高零售价，单位：元。例：“10.56”。必须在0.01元到10000000元之间，最高零售价必须大于最低零售价。 */
func (this *FenxiaoProductUpdateRequest) SetRetailPriceHigh(value string) {
	this.Set("retail_price_high", value)
}

/* 最低零售价，单位：元。例：“10.56”。必须在0.01元到10000000元之间。 */
func (this *FenxiaoProductUpdateRequest) SetRetailPriceLow(value string) {
	this.Set("retail_price_low", value)
}

/* sku采购价格，单位元，例："10.50,11.00,20.50"，字段必须和上面的sku_ids或sku_properties保持一致。 */
func (this *FenxiaoProductUpdateRequest) SetSkuCostPrices(value string) {
	this.Set("sku_cost_prices", value)
}

/* sku的经销采购价。如果多个，用逗号分隔，并与其他sku信息保持相同顺序。其中每个值的单位：元。例：“10.56,12.3”。必须在0.01元到10000000元之间。 */
func (this *FenxiaoProductUpdateRequest) SetSkuDealerCostPrices(value string) {
	this.Set("sku_dealer_cost_prices", value)
}

/* sku id列表，例：1001,1002,1003。如果传入sku_properties将忽略此参数。 */
func (this *FenxiaoProductUpdateRequest) SetSkuIds(value string) {
	this.Set("sku_ids", value)
}

/* sku商家编码 ，单位元，例："S1000,S1002,S1003"，字段必须和上面的id或sku_properties保持一致，如果没有可以写成",," */
func (this *FenxiaoProductUpdateRequest) SetSkuOuterIds(value string) {
	this.Set("sku_outer_ids", value)
}

/* sku属性。格式:pid:vid;pid:vid,表示一组属性如:1627207:3232483;1630696:3284570,表示一组:机身颜色:军绿色;手机套餐:一电一充。多组之间用逗号“,”区分。(属性的pid调用taobao.itemprops.get取得，属性值的vid用taobao.itempropvalues.get取得vid)
通过此字段可新增和更新sku。若传入此值将忽略sku_ids字段。sku其他字段与此值保持一致。 */
func (this *FenxiaoProductUpdateRequest) SetSkuProperties(value string) {
	this.Set("sku_properties", value)
}

/* 根据sku属性删除sku信息。需要按组删除属性。 */
func (this *FenxiaoProductUpdateRequest) SetSkuPropertiesDel(value string) {
	this.Set("sku_properties_del", value)
}

/* sku库存，单位元，例："10,20,30"，字段必须和sku_ids或sku_properties保持一致。 */
func (this *FenxiaoProductUpdateRequest) SetSkuQuantitys(value string) {
	this.Set("sku_quantitys", value)
}

/* sku采购基准价，单位元，例："10.50,11.00,20.50"，字段必须和上面的sku_ids或sku_properties保持一致。 */
func (this *FenxiaoProductUpdateRequest) SetSkuStandardPrices(value string) {
	this.Set("sku_standard_prices", value)
}

/* 采购基准价，单位：元。例：“10.56”。必须在0.01元到10000000元之间。 */
func (this *FenxiaoProductUpdateRequest) SetStandardPrice(value string) {
	this.Set("standard_price", value)
}

/* 零售基准价，单位：元。例：“10.56”。必须在0.01元到10000000元之间。 */
func (this *FenxiaoProductUpdateRequest) SetStandardRetailPrice(value string) {
	this.Set("standard_retail_price", value)
}

/* 发布状态，可选值：up（上架）、down（下架）、delete（删除），输入非法字符则忽略。 */
func (this *FenxiaoProductUpdateRequest) SetStatus(value string) {
	this.Set("status", value)
}

type FenxiaoProductUpdateResponse struct {
	*ErrorResponse                     `json:"error_response,omitempty"`
	FenxiaoProductUpdateResponseResult `json:"fenxiao_product_update_response"`
}
type FenxiaoProductUpdateResponseResult struct {
	/* 更新时间，时间格式：yyyy-MM-dd HH:mm:ss */
	Modified string `json:"modified"`
	/* 产品ID */
	Pid int64 `json:"pid"`
}

/*taobao.fenxiao.productcat.add*/
type FenxiaoProductcatAddRequest struct {
	values url.Values
}

func (this *FenxiaoProductcatAddRequest) GetApiMethodName() string {
	return "taobao.fenxiao.productcat.add"
}
func (this *FenxiaoProductcatAddRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *FenxiaoProductcatAddRequest) GetValues() url.Values {
	return this.values
}

/* 代销默认采购价比例，注意：100.00%，则输入为10000<br /> 支持最大值为：99999<br /> 支持最小值为：100 */
func (this *FenxiaoProductcatAddRequest) SetAgentCostPercent(value string) {
	this.Set("agent_cost_percent", value)
}

/* 经销默认采购价比例，注意：100.00%，则输入为10000<br /> 支持最大值为：99999<br /> 支持最小值为：100 */
func (this *FenxiaoProductcatAddRequest) SetDealerCostPercent(value string) {
	this.Set("dealer_cost_percent", value)
}

/* 产品线名称<br /> 支持最大长度为：10<br /> 支持的最大列表长度为：10 */
func (this *FenxiaoProductcatAddRequest) SetName(value string) {
	this.Set("name", value)
}

/* 最高零售价比例，注意：100.00%，则输入为10000<br /> 支持最大值为：99999<br /> 支持最小值为：100 */
func (this *FenxiaoProductcatAddRequest) SetRetailHighPercent(value string) {
	this.Set("retail_high_percent", value)
}

/* 最低零售价比例，注意：100.00%，则输入为10000<br /> 支持最大值为：99999<br /> 支持最小值为：100 */
func (this *FenxiaoProductcatAddRequest) SetRetailLowPercent(value string) {
	this.Set("retail_low_percent", value)
}

type FenxiaoProductcatAddResponse struct {
	*ErrorResponse                     `json:"error_response,omitempty"`
	FenxiaoProductcatAddResponseResult `json:"fenxiao_productcat_add_response"`
}
type FenxiaoProductcatAddResponseResult struct {
	/* 操作是否成功 */
	IsSuccess bool `json:"is_success"`
	/* 产品线ID */
	ProductLineId int64 `json:"product_line_id"`
}

/*taobao.fenxiao.productcat.delete*/
type FenxiaoProductcatDeleteRequest struct {
	values url.Values
}

func (this *FenxiaoProductcatDeleteRequest) GetApiMethodName() string {
	return "taobao.fenxiao.productcat.delete"
}
func (this *FenxiaoProductcatDeleteRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *FenxiaoProductcatDeleteRequest) GetValues() url.Values {
	return this.values
}

/* 产品线ID */
func (this *FenxiaoProductcatDeleteRequest) SetProductLineId(value string) {
	this.Set("product_line_id", value)
}

type FenxiaoProductcatDeleteResponse struct {
	*ErrorResponse                        `json:"error_response,omitempty"`
	FenxiaoProductcatDeleteResponseResult `json:"fenxiao_productcat_delete_response"`
}
type FenxiaoProductcatDeleteResponseResult struct {
	/* 操作是否成功 */
	IsSuccess bool `json:"is_success"`
}

/*taobao.fenxiao.productcat.update*/
type FenxiaoProductcatUpdateRequest struct {
	values url.Values
}

func (this *FenxiaoProductcatUpdateRequest) GetApiMethodName() string {
	return "taobao.fenxiao.productcat.update"
}
func (this *FenxiaoProductcatUpdateRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *FenxiaoProductcatUpdateRequest) GetValues() url.Values {
	return this.values
}

/* 代销默认采购价比例，注意：100.00%，则输入为10000<br /> 支持最大值为：99999<br /> 支持最小值为：100 */
func (this *FenxiaoProductcatUpdateRequest) SetAgentCostPercent(value string) {
	this.Set("agent_cost_percent", value)
}

/* 经销默认采购价比例，注意：100.00%，则输入为10000<br /> 支持最大值为：99999<br /> 支持最小值为：100 */
func (this *FenxiaoProductcatUpdateRequest) SetDealerCostPercent(value string) {
	this.Set("dealer_cost_percent", value)
}

/* 产品线名称<br /> 支持最大长度为：10<br /> 支持的最大列表长度为：10 */
func (this *FenxiaoProductcatUpdateRequest) SetName(value string) {
	this.Set("name", value)
}

/* 产品线ID */
func (this *FenxiaoProductcatUpdateRequest) SetProductLineId(value string) {
	this.Set("product_line_id", value)
}

/* 最高零售价比例，注意：100.00%，则输入为10000<br /> 支持最大值为：99999<br /> 支持最小值为：100 */
func (this *FenxiaoProductcatUpdateRequest) SetRetailHighPercent(value string) {
	this.Set("retail_high_percent", value)
}

/* 最低零售价比例，注意：100.00%，则输入为10000<br /> 支持最大值为：99999<br /> 支持最小值为：100 */
func (this *FenxiaoProductcatUpdateRequest) SetRetailLowPercent(value string) {
	this.Set("retail_low_percent", value)
}

type FenxiaoProductcatUpdateResponse struct {
	*ErrorResponse                        `json:"error_response,omitempty"`
	FenxiaoProductcatUpdateResponseResult `json:"fenxiao_productcat_update_response"`
}
type FenxiaoProductcatUpdateResponseResult struct {
	/* 操作是否成功 */
	IsSuccess bool `json:"is_success"`
}

/*taobao.fenxiao.productcats.get*/
type FenxiaoProductcatsGetRequest struct {
	values url.Values
}

func (this *FenxiaoProductcatsGetRequest) GetApiMethodName() string {
	return "taobao.fenxiao.productcats.get"
}
func (this *FenxiaoProductcatsGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *FenxiaoProductcatsGetRequest) GetValues() url.Values {
	return this.values
}

/* 返回字段列表 */
func (this *FenxiaoProductcatsGetRequest) SetFields(value string) {
	this.Set("fields", value)
}

type FenxiaoProductcatsGetResponse struct {
	*ErrorResponse                      `json:"error_response,omitempty"`
	FenxiaoProductcatsGetResponseResult `json:"fenxiao_productcats_get_response"`
}
type FenxiaoProductcatsGetResponseResult struct {
	/* 产品线列表。返回 ProductCat 包含的字段信息。 */
	Productcats struct {
		ProductCat []*ProductCat `json:"product_cat"`
	} `json:"Productcats"`
	/* 查询结果记录数 */
	TotalResults int64 `json:"total_results"`
}

/*taobao.fenxiao.products.get*/
type FenxiaoProductsGetRequest struct {
	values url.Values
}

func (this *FenxiaoProductsGetRequest) GetApiMethodName() string {
	return "taobao.fenxiao.products.get"
}
func (this *FenxiaoProductsGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *FenxiaoProductsGetRequest) GetValues() url.Values {
	return this.values
}

/* 结束修改时间 */
func (this *FenxiaoProductsGetRequest) SetEndModified(value string) {
	this.Set("end_modified", value)
}

/* 指定查询额外的信息，可选值：skus（sku数据）、images（多图），多个可选值用逗号分割。 */
func (this *FenxiaoProductsGetRequest) SetFields(value string) {
	this.Set("fields", value)
}

/* 查询产品列表时，查询入参“是否需要授权”
yes:需要授权
no:不需要授权 */
func (this *FenxiaoProductsGetRequest) SetIsAuthz(value string) {
	this.Set("is_authz", value)
}

/* 可根据导入的商品ID列表查询，优先级次于产品ID、sku_numbers，高于其他分页查询条件。最大限制20，用逗号分割，例如：“1001,1002,1003,1004,1005” */
func (this *FenxiaoProductsGetRequest) SetItemIds(value string) {
	this.Set("item_ids", value)
}

/* 商家编码 */
func (this *FenxiaoProductsGetRequest) SetOuterId(value string) {
	this.Set("outer_id", value)
}

/* 页码（大于0的整数，默认1） */
func (this *FenxiaoProductsGetRequest) SetPageNo(value string) {
	this.Set("page_no", value)
}

/* 每页记录数（默认20，最大50） */
func (this *FenxiaoProductsGetRequest) SetPageSize(value string) {
	this.Set("page_size", value)
}

/* 产品ID列表（最大限制30），用逗号分割，例如：“1001,1002,1003,1004,1005” */
func (this *FenxiaoProductsGetRequest) SetPids(value string) {
	this.Set("pids", value)
}

/* 产品线ID */
func (this *FenxiaoProductsGetRequest) SetProductcatId(value string) {
	this.Set("productcat_id", value)
}

/* sku商家编码 */
func (this *FenxiaoProductsGetRequest) SetSkuNumber(value string) {
	this.Set("sku_number", value)
}

/* 开始修改时间 */
func (this *FenxiaoProductsGetRequest) SetStartModified(value string) {
	this.Set("start_modified", value)
}

/* 产品状态，可选值：up（上架）、down（下架），不传默认查询所有 */
func (this *FenxiaoProductsGetRequest) SetStatus(value string) {
	this.Set("status", value)
}

type FenxiaoProductsGetResponse struct {
	*ErrorResponse                   `json:"error_response,omitempty"`
	FenxiaoProductsGetResponseResult `json:"fenxiao_products_get_response"`
}
type FenxiaoProductsGetResponseResult struct {
	/* 产品对象记录集。返回 FenxiaoProduct 包含的字段信息。 */
	Products struct {
		FenxiaoProduct []*FenxiaoProduct `json:"fenxiao_product"`
	} `json:"Products"`
	/* 查询结果记录数 */
	TotalResults int64 `json:"total_results"`
}

/*taobao.fenxiao.refund.create*/
type FenxiaoRefundCreateRequest struct {
	values url.Values
}

func (this *FenxiaoRefundCreateRequest) GetApiMethodName() string {
	return "taobao.fenxiao.refund.create"
}
func (this *FenxiaoRefundCreateRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *FenxiaoRefundCreateRequest) GetValues() url.Values {
	return this.values
}

/* 是否退货，只有子采购单发货后，才能申请退货 */
func (this *FenxiaoRefundCreateRequest) SetIsReturnGoods(value string) {
	this.Set("is_return_goods", value)
}

/* 此条子采购单是否退邮费，邮费要记在最后一笔申请退款的子单上面 */
func (this *FenxiaoRefundCreateRequest) SetIsReturnPostFee(value string) {
	this.Set("is_return_post_fee", value)
}

/* 退款说明,2-200字 */
func (this *FenxiaoRefundCreateRequest) SetRefundDesc(value string) {
	this.Set("refund_desc", value)
}

/* 发货前：
1、缺货，2、拍错商品，3、商品缺少所需样式，4、协商一致退款，5、未及时发货，0、其它
发货后：
经销：
1、商品质量问题，2、收到的商品不符，3、协商一致退款，4、一直未收到货，5、退还多付邮费，6、折扣、赠品、发票等问题，0、其它
代销：
1、商品质量问题，2、收到的商品不符，3、协商一致退款，4、买家一直未收到货，5、退还多付邮费，6、折扣、赠品、发票等问题，0、其它 */
func (this *FenxiaoRefundCreateRequest) SetRefundReasonId(value string) {
	this.Set("refund_reason_id", value)
}

/* 需要退款的金额，单位为分 */
func (this *FenxiaoRefundCreateRequest) SetReturnFee(value string) {
	this.Set("return_fee", value)
}

/* 需要创建退款的子采购单id */
func (this *FenxiaoRefundCreateRequest) SetSubOrderId(value string) {
	this.Set("sub_order_id", value)
}

type FenxiaoRefundCreateResponse struct {
	*ErrorResponse                    `json:"error_response,omitempty"`
	FenxiaoRefundCreateResponseResult `json:"fenxiao_refund_create_response"`
}
type FenxiaoRefundCreateResponseResult struct {
	/* 退款是否创建成功 */
	IsSuccess bool `json:"is_success"`
}

/*taobao.fenxiao.refund.get*/
type FenxiaoRefundGetRequest struct {
	values url.Values
}

func (this *FenxiaoRefundGetRequest) GetApiMethodName() string {
	return "taobao.fenxiao.refund.get"
}
func (this *FenxiaoRefundGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *FenxiaoRefundGetRequest) GetValues() url.Values {
	return this.values
}

/* 是否查询下游买家的退款信息 */
func (this *FenxiaoRefundGetRequest) SetQuerySellerRefund(value string) {
	this.Set("query_seller_refund", value)
}

/* 要查询的退款子单的id */
func (this *FenxiaoRefundGetRequest) SetSubOrderId(value string) {
	this.Set("sub_order_id", value)
}

type FenxiaoRefundGetResponse struct {
	*ErrorResponse                 `json:"error_response,omitempty"`
	FenxiaoRefundGetResponseResult `json:"fenxiao_refund_get_response"`
}
type FenxiaoRefundGetResponseResult struct {
	/* 退款详情 */
	RefundDetail *RefundDetail `json:"refund_detail"`
}

/*taobao.fenxiao.refund.message.add*/
type FenxiaoRefundMessageAddRequest struct {
	values url.Values
}

func (this *FenxiaoRefundMessageAddRequest) GetApiMethodName() string {
	return "taobao.fenxiao.refund.message.add"
}
func (this *FenxiaoRefundMessageAddRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *FenxiaoRefundMessageAddRequest) GetValues() url.Values {
	return this.values
}

/* 留言凭证 */
func (this *FenxiaoRefundMessageAddRequest) SetImage(value string) {
	this.Set("image", value)
}

/* 留言内容 */
func (this *FenxiaoRefundMessageAddRequest) SetMessageContent(value string) {
	this.Set("message_content", value)
}

/* 发生退款的子采购单id */
func (this *FenxiaoRefundMessageAddRequest) SetSubOrderId(value string) {
	this.Set("sub_order_id", value)
}

type FenxiaoRefundMessageAddResponse struct {
	*ErrorResponse                        `json:"error_response,omitempty"`
	FenxiaoRefundMessageAddResponseResult `json:"fenxiao_refund_message_add_response"`
}
type FenxiaoRefundMessageAddResponseResult struct {
	/* 退款是否创建成功 */
	IsSuccess bool `json:"is_success"`
}

/*taobao.fenxiao.refund.message.get*/
type FenxiaoRefundMessageGetRequest struct {
	values url.Values
}

func (this *FenxiaoRefundMessageGetRequest) GetApiMethodName() string {
	return "taobao.fenxiao.refund.message.get"
}
func (this *FenxiaoRefundMessageGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *FenxiaoRefundMessageGetRequest) GetValues() url.Values {
	return this.values
}

/* 页码。（大于0的整数。默认为1） */
func (this *FenxiaoRefundMessageGetRequest) SetPageNo(value string) {
	this.Set("page_no", value)
}

/* 每页条数。（默认10） */
func (this *FenxiaoRefundMessageGetRequest) SetPageSize(value string) {
	this.Set("page_size", value)
}

/* 子采购单id */
func (this *FenxiaoRefundMessageGetRequest) SetSubOrderId(value string) {
	this.Set("sub_order_id", value)
}

type FenxiaoRefundMessageGetResponse struct {
	*ErrorResponse                        `json:"error_response,omitempty"`
	FenxiaoRefundMessageGetResponseResult `json:"fenxiao_refund_message_get_response"`
}
type FenxiaoRefundMessageGetResponseResult struct {
	/* 退款留言 */
	OrderMessages struct {
		OrderMessage []*OrderMessage `json:"order_message"`
	} `json:"OrderMessages"`
	/* 总留言条数 */
	TotalResults int64 `json:"total_results"`
}

/*taobao.fenxiao.refund.query*/
type FenxiaoRefundQueryRequest struct {
	values url.Values
}

func (this *FenxiaoRefundQueryRequest) GetApiMethodName() string {
	return "taobao.fenxiao.refund.query"
}
func (this *FenxiaoRefundQueryRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *FenxiaoRefundQueryRequest) GetValues() url.Values {
	return this.values
}

/* 代销采购退款最迟修改时间。与start_date的最大时间间隔不能超过30天 */
func (this *FenxiaoRefundQueryRequest) SetEndDate(value string) {
	this.Set("end_date", value)
}

/* 页码（大于0的整数。无值或小于1的值按默认值1计） */
func (this *FenxiaoRefundQueryRequest) SetPageNo(value string) {
	this.Set("page_no", value)
}

/* 每页条数（大于0但小于等于50的整数。无值或大于50或小于1的值按默认值50计） */
func (this *FenxiaoRefundQueryRequest) SetPageSize(value string) {
	this.Set("page_size", value)
}

/* 是否查询下游买家的退款信息 */
func (this *FenxiaoRefundQueryRequest) SetQuerySellerRefund(value string) {
	this.Set("query_seller_refund", value)
}

/* 代销采购退款单最早修改时间 */
func (this *FenxiaoRefundQueryRequest) SetStartDate(value string) {
	this.Set("start_date", value)
}

type FenxiaoRefundQueryResponse struct {
	*ErrorResponse                   `json:"error_response,omitempty"`
	FenxiaoRefundQueryResponseResult `json:"fenxiao_refund_query_response"`
}
type FenxiaoRefundQueryResponseResult struct {
	/* 代销采购退款列表 */
	RefundList struct {
		RefundDetail []*RefundDetail `json:"refund_detail"`
	} `json:"RefundList"`
	/* 按查询条件查到的记录总数 */
	TotalResults int64 `json:"total_results"`
}

/*taobao.fenxiao.refund.update*/
type FenxiaoRefundUpdateRequest struct {
	values url.Values
}

func (this *FenxiaoRefundUpdateRequest) GetApiMethodName() string {
	return "taobao.fenxiao.refund.update"
}
func (this *FenxiaoRefundUpdateRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *FenxiaoRefundUpdateRequest) GetValues() url.Values {
	return this.values
}

/* 是否退货，只有子采购单发货后，才能申请退货 */
func (this *FenxiaoRefundUpdateRequest) SetIsReturnGoods(value string) {
	this.Set("is_return_goods", value)
}

/* 退款说明 */
func (this *FenxiaoRefundUpdateRequest) SetRefundDesc(value string) {
	this.Set("refund_desc", value)
}

/* 发货前：
1、缺货，2、拍错商品，3、商品缺少所需样式，4、协商一致退款，5、未及时发货，0、其它
发货后：
经销：
1、商品质量问题，2、收到的商品不符，3、协商一致退款，4、一直未收到货，5、退还多付邮费，6、折扣、赠品、发票等问题，0、其它
代销：
1、商品质量问题，2、收到的商品不符，3、协商一致退款，4、买家一直未收到货，5、退还多付邮费，6、折扣、赠品、发票等问题，0、其它 */
func (this *FenxiaoRefundUpdateRequest) SetRefundReasonId(value string) {
	this.Set("refund_reason_id", value)
}

/* 需要退款的金额 */
func (this *FenxiaoRefundUpdateRequest) SetReturnFee(value string) {
	this.Set("return_fee", value)
}

/* 需要修改退款的子采购单id */
func (this *FenxiaoRefundUpdateRequest) SetSubOrderId(value string) {
	this.Set("sub_order_id", value)
}

type FenxiaoRefundUpdateResponse struct {
	*ErrorResponse                    `json:"error_response,omitempty"`
	FenxiaoRefundUpdateResponseResult `json:"fenxiao_refund_update_response"`
}
type FenxiaoRefundUpdateResponseResult struct {
	/* 退款是否修改成功 */
	IsSuccess bool `json:"is_success"`
}

/*taobao.fenxiao.requisitions.get*/
type FenxiaoRequisitionsGetRequest struct {
	values url.Values
}

func (this *FenxiaoRequisitionsGetRequest) GetApiMethodName() string {
	return "taobao.fenxiao.requisitions.get"
}
func (this *FenxiaoRequisitionsGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *FenxiaoRequisitionsGetRequest) GetValues() url.Values {
	return this.values
}

/* 申请结束时间yyyy-MM-dd */
func (this *FenxiaoRequisitionsGetRequest) SetApplyEnd(value string) {
	this.Set("apply_end", value)
}

/* 申请开始时间yyyy-MM-dd */
func (this *FenxiaoRequisitionsGetRequest) SetApplyStart(value string) {
	this.Set("apply_start", value)
}

/* 页码（大于0的整数，默认1） */
func (this *FenxiaoRequisitionsGetRequest) SetPageNo(value string) {
	this.Set("page_no", value)
}

/* 每页记录数（默认20，最大50） */
func (this *FenxiaoRequisitionsGetRequest) SetPageSize(value string) {
	this.Set("page_size", value)
}

/* 申请状态（1-申请中、2-成功、3-被退回、4-已撤消、5-过期） */
func (this *FenxiaoRequisitionsGetRequest) SetStatus(value string) {
	this.Set("status", value)
}

type FenxiaoRequisitionsGetResponse struct {
	*ErrorResponse                       `json:"error_response,omitempty"`
	FenxiaoRequisitionsGetResponseResult `json:"fenxiao_requisitions_get_response"`
}
type FenxiaoRequisitionsGetResponseResult struct {
	/* 操作是否成功 */
	IsSuccess bool `json:"is_success"`
	/* 合作申请 */
	Requisitions struct {
		Requisition []*Requisition `json:"requisition"`
	} `json:"Requisitions"`
	/* 结果记录数 */
	TotalResults int64 `json:"total_results"`
}

/*taobao.fenxiao.trademonitor.get*/
type FenxiaoTrademonitorGetRequest struct {
	values url.Values
}

func (this *FenxiaoTrademonitorGetRequest) GetApiMethodName() string {
	return "taobao.fenxiao.trademonitor.get"
}
func (this *FenxiaoTrademonitorGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *FenxiaoTrademonitorGetRequest) GetValues() url.Values {
	return this.values
}

/* 经销商的淘宝账号 */
func (this *FenxiaoTrademonitorGetRequest) SetDistributorNick(value string) {
	this.Set("distributor_nick", value)
}

/* 结束时间 格式 yyyy-MM-dd HH:mm:ss.支持到秒的查询。若不传时分秒，默认为0时0分0秒。 */
func (this *FenxiaoTrademonitorGetRequest) SetEndCreated(value string) {
	this.Set("end_created", value)
}

/* 多个字段用","分隔。 fields 如果为空：返回所有采购单对象(purchase_orders)字段。 如果不为空：返回指定采购单对象(purchase_orders)字段。例如：trade_monitors.item_title表示只返回item_title */
func (this *FenxiaoTrademonitorGetRequest) SetFields(value string) {
	this.Set("fields", value)
}

/* 页码。（大于0的整数。小于1按1计） */
func (this *FenxiaoTrademonitorGetRequest) SetPageNo(value string) {
	this.Set("page_no", value)
}

/* 每页条数。（每页条数不超过50条，超过50或小于0均按50计） */
func (this *FenxiaoTrademonitorGetRequest) SetPageSize(value string) {
	this.Set("page_size", value)
}

/* 产品id */
func (this *FenxiaoTrademonitorGetRequest) SetProductId(value string) {
	this.Set("product_id", value)
}

/* 起始时间 格式 yyyy-MM-dd HH:mm:ss.支持到秒的查询。若不传时分秒，默认为0时0分0秒。 */
func (this *FenxiaoTrademonitorGetRequest) SetStartCreated(value string) {
	this.Set("start_created", value)
}

type FenxiaoTrademonitorGetResponse struct {
	*ErrorResponse                       `json:"error_response,omitempty"`
	FenxiaoTrademonitorGetResponseResult `json:"fenxiao_trademonitor_get_response"`
}
type FenxiaoTrademonitorGetResponseResult struct {
	/* 搜索到的经销商品订单数量 */
	TotalResults int64 `json:"total_results"`
	/* 经销商品订单监控信息 */
	TradeMonitors struct {
		TradeMonitor []*TradeMonitor `json:"trade_monitor"`
	} `json:"TradeMonitors"`
}

/*taobao.inventory.adjust.external*/
type InventoryAdjustExternalRequest struct {
	values url.Values
}

func (this *InventoryAdjustExternalRequest) GetApiMethodName() string {
	return "taobao.inventory.adjust.external"
}
func (this *InventoryAdjustExternalRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *InventoryAdjustExternalRequest) GetValues() url.Values {
	return this.values
}

/* 外部订单类型：ALLOCATE:调拨、 RETURN:退货、PURCHACE：采购、BALANCE:盘点、NON_TAOBAO_TRADE：非淘宝交易、OTHERS：其他 */
func (this *InventoryAdjustExternalRequest) SetBizType(value string) {
	this.Set("biz_type", value)
}

/* 商家外部定单号 */
func (this *InventoryAdjustExternalRequest) SetBizUniqueCode(value string) {
	this.Set("biz_unique_code", value)
}

/* 商品初始库存信息： [{"scItemId":"商品后端ID，如果有传scItemCode,参数可以为0","scItemCode":"商品商家编码","inventoryType":"库存类型  1：正常,”direction”: 1: 盘盈 -1: 盘亏,参数可选,"quantity":"数量(正数)"}] */
func (this *InventoryAdjustExternalRequest) SetItems(value string) {
	this.Set("items", value)
}

/* 库存占用返回的操作码.operate_type 为OUTBOUND时，如果是确认事先进行过的库存占用，需要传入当时返回的操作码，并且明细必须与申请时保持一致 */
func (this *InventoryAdjustExternalRequest) SetOccupyOperateCode(value string) {
	this.Set("occupy_operate_code", value)
}

/* 业务操作时间 */
func (this *InventoryAdjustExternalRequest) SetOperateTime(value string) {
	this.Set("operate_time", value)
}

/* 库存操作类别：INBOUND：入库 OUTBOUND: 出库BALANCE: 盘点 */
func (this *InventoryAdjustExternalRequest) SetOperateType(value string) {
	this.Set("operate_type", value)
}

/* 出库时库存扣减类型，operate_type为OUTBOUND时有效。
AUTO_CALCULATE:自动计算可供扣减，如果库存不够返回失败 ClIENT_FORCE：强制要求最大化扣减，有多少扣多少 */
func (this *InventoryAdjustExternalRequest) SetReduceType(value string) {
	this.Set("reduce_type", value)
}

/* 商家仓库编码 */
func (this *InventoryAdjustExternalRequest) SetStoreCode(value string) {
	this.Set("store_code", value)
}

type InventoryAdjustExternalResponse struct {
	*ErrorResponse                        `json:"error_response,omitempty"`
	InventoryAdjustExternalResponseResult `json:"inventory_adjust_external_response"`
}
type InventoryAdjustExternalResponseResult struct {
	/* 操作返回码 */
	OperateCode string `json:"operate_code"`
	/* 提示信息 */
	TipInfos struct {
		TipInfo []*TipInfo `json:"tip_info"`
	} `json:"TipInfos"`
}

/*taobao.inventory.adjust.trade*/
type InventoryAdjustTradeRequest struct {
	values url.Values
}

func (this *InventoryAdjustTradeRequest) GetApiMethodName() string {
	return "taobao.inventory.adjust.trade"
}
func (this *InventoryAdjustTradeRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *InventoryAdjustTradeRequest) GetValues() url.Values {
	return this.values
}

/* 商家外部定单号 */
func (this *InventoryAdjustTradeRequest) SetBizUniqueCode(value string) {
	this.Set("biz_unique_code", value)
}

/* 商品初始库存信息： [{ "TBOrderCode”:”淘宝交易号”,"TBSubOrderCode ":"淘宝子交易单号,赠品可以不填","”isGift”:”TRUE或者FALSE,是否赠品”,storeCode":"商家仓库编码"," scItemId ":"商品后端ID","scItemCode":"商品商家编码"," originScItemId ":"原商品ID","inventoryType":"","quantity":"111","isComplete":"TRUE或者FALSE，是否全部确认出库"}] */
func (this *InventoryAdjustTradeRequest) SetItems(value string) {
	this.Set("items", value)
}

/* 业务操作时间 */
func (this *InventoryAdjustTradeRequest) SetOperateTime(value string) {
	this.Set("operate_time", value)
}

/* 订单类型：B2C、B2B */
func (this *InventoryAdjustTradeRequest) SetTbOrderType(value string) {
	this.Set("tb_order_type", value)
}

type InventoryAdjustTradeResponse struct {
	*ErrorResponse                     `json:"error_response,omitempty"`
	InventoryAdjustTradeResponseResult `json:"inventory_adjust_trade_response"`
}
type InventoryAdjustTradeResponseResult struct {
	/* 操作返回码 */
	OperateCode string `json:"operate_code"`
	/* 提示信息 */
	TipInfos struct {
		TipInfo []*TipInfo `json:"tip_info"`
	} `json:"TipInfos"`
}

/*taobao.inventory.initial*/
type InventoryInitialRequest struct {
	values url.Values
}

func (this *InventoryInitialRequest) GetApiMethodName() string {
	return "taobao.inventory.initial"
}
func (this *InventoryInitialRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *InventoryInitialRequest) GetValues() url.Values {
	return this.values
}

/* 商品初始库存信息： [{"scItemId":"商品后端ID，如果有传scItemCode,参数可以为0","scItemCode":"商品商家编码","inventoryType":"库存类型  1：正常,2：损坏,3：冻结,10：质押,11-20:用户自定义","quantity":"数量"}] */
func (this *InventoryInitialRequest) SetItems(value string) {
	this.Set("items", value)
}

/* 商家仓库编码 */
func (this *InventoryInitialRequest) SetStoreCode(value string) {
	this.Set("store_code", value)
}

type InventoryInitialResponse struct {
	*ErrorResponse                 `json:"error_response,omitempty"`
	InventoryInitialResponseResult `json:"inventory_initial_response"`
}
type InventoryInitialResponseResult struct {
	/* 提示信息 */
	TipInfos struct {
		TipInfo []*TipInfo `json:"tip_info"`
	} `json:"TipInfos"`
}

/*taobao.inventory.initial.item*/
type InventoryInitialItemRequest struct {
	values url.Values
}

func (this *InventoryInitialItemRequest) GetApiMethodName() string {
	return "taobao.inventory.initial.item"
}
func (this *InventoryInitialItemRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *InventoryInitialItemRequest) GetValues() url.Values {
	return this.values
}

/* 后端商品id */
func (this *InventoryInitialItemRequest) SetScItemId(value string) {
	this.Set("sc_item_id", value)
}

/* 商品初始库存信息： [{"storeCode":"必选,商家仓库编号","inventoryType":"可选，库存类型 1：正常,2：损坏,3：冻结,10：质押,11-20:用户自定义,默认为1","quantity":"必选,数量"}] */
func (this *InventoryInitialItemRequest) SetStoreInventorys(value string) {
	this.Set("store_inventorys", value)
}

type InventoryInitialItemResponse struct {
	*ErrorResponse                     `json:"error_response,omitempty"`
	InventoryInitialItemResponseResult `json:"inventory_initial_item_response"`
}
type InventoryInitialItemResponseResult struct {
	/* 提示信息 */
	TipInfos struct {
		TipInfo []*TipInfo `json:"tip_info"`
	} `json:"TipInfos"`
}

/*taobao.inventory.occupy.adjust*/
type InventoryOccupyAdjustRequest struct {
	values url.Values
}

func (this *InventoryOccupyAdjustRequest) GetApiMethodName() string {
	return "taobao.inventory.occupy.adjust"
}
func (this *InventoryOccupyAdjustRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *InventoryOccupyAdjustRequest) GetValues() url.Values {
	return this.values
}

/* 商家外部订单号 */
func (this *InventoryOccupyAdjustRequest) SetBizUniqueCode(value string) {
	this.Set("biz_unique_code", value)
}

/* 商品初始库存信息： [{ "TBOrderCode":"淘宝交易号","TBSubOrderCode ":"淘宝子交易单号","originalStoreCode":"交易下单的仓库编码","storeCode":"要切换占用到的商家仓库编码"," scItemId ":"商品后端ID","scItemCode":"商品商家编码","inventoryType":"仓储类型","quantity":"新仓库的占用数量，如果不传，则取用原先的占用数"}] */
func (this *InventoryOccupyAdjustRequest) SetItems(value string) {
	this.Set("items", value)
}

/* 业务操作时间 */
func (this *InventoryOccupyAdjustRequest) SetOperateTime(value string) {
	this.Set("operate_time", value)
}

/* 订单类型：B2C、B2B */
func (this *InventoryOccupyAdjustRequest) SetTbOrderType(value string) {
	this.Set("tb_order_type", value)
}

type InventoryOccupyAdjustResponse struct {
	*ErrorResponse                      `json:"error_response,omitempty"`
	InventoryOccupyAdjustResponseResult `json:"inventory_occupy_adjust_response"`
}
type InventoryOccupyAdjustResponseResult struct {
	/* 操作返回码 */
	OperateCode string `json:"operate_code"`
	/* 提示信息 */
	TipInfos struct {
		TipInfo []*TipInfo `json:"tip_info"`
	} `json:"TipInfos"`
}

/*taobao.inventory.query*/
type InventoryQueryRequest struct {
	values url.Values
}

func (this *InventoryQueryRequest) GetApiMethodName() string {
	return "taobao.inventory.query"
}
func (this *InventoryQueryRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *InventoryQueryRequest) GetValues() url.Values {
	return this.values
}

/* 后端商品的商家编码列表，控制到50个 */
func (this *InventoryQueryRequest) SetScItemCodes(value string) {
	this.Set("sc_item_codes", value)
}

/* 后端商品ID 列表，控制到50个 */
func (this *InventoryQueryRequest) SetScItemIds(value string) {
	this.Set("sc_item_ids", value)
}

/* 卖家昵称 */
func (this *InventoryQueryRequest) SetSellerNick(value string) {
	this.Set("seller_nick", value)
}

/* 仓库列表:GLY001^GLY002 */
func (this *InventoryQueryRequest) SetStoreCodes(value string) {
	this.Set("store_codes", value)
}

type InventoryQueryResponse struct {
	*ErrorResponse               `json:"error_response,omitempty"`
	InventoryQueryResponseResult `json:"inventory_query_response"`
}
type InventoryQueryResponseResult struct {
	/* 商品总体库存信息 */
	ItemInventorys struct {
		InventorySum []*InventorySum `json:"inventory_sum"`
	} `json:"ItemInventorys"`
	/* 提示信息，提示不存在的后端商品 */
	TipInfos struct {
		TipInfo []*TipInfo `json:"tip_info"`
	} `json:"TipInfos"`
}

/*taobao.inventory.store.manage*/
type InventoryStoreManageRequest struct {
	values url.Values
}

func (this *InventoryStoreManageRequest) GetApiMethodName() string {
	return "taobao.inventory.store.manage"
}
func (this *InventoryStoreManageRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *InventoryStoreManageRequest) GetValues() url.Values {
	return this.values
}

/* 仓库的物理地址，可更新 */
func (this *InventoryStoreManageRequest) SetAddress(value string) {
	this.Set("address", value)
}

/* 仓库区域名，可更新 */
func (this *InventoryStoreManageRequest) SetAddressAreaName(value string) {
	this.Set("address_area_name", value)
}

/* 仓库简称，可更新 */
func (this *InventoryStoreManageRequest) SetAliasName(value string) {
	this.Set("alias_name", value)
}

/* 联系人，可更新 */
func (this *InventoryStoreManageRequest) SetContact(value string) {
	this.Set("contact", value)
}

/* 参数定义，ADD：新建; UPDATE：更新 */
func (this *InventoryStoreManageRequest) SetOperateType(value string) {
	this.Set("operate_type", value)
}

/* 联系电话，可更新 */
func (this *InventoryStoreManageRequest) SetPhone(value string) {
	this.Set("phone", value)
}

/* 邮编，可更新 */
func (this *InventoryStoreManageRequest) SetPostcode(value string) {
	this.Set("postcode", value)
}

/* 商家的仓库编码，不允许重复，不允许更新 */
func (this *InventoryStoreManageRequest) SetStoreCode(value string) {
	this.Set("store_code", value)
}

/* 商家的仓库名称，可更新 */
func (this *InventoryStoreManageRequest) SetStoreName(value string) {
	this.Set("store_name", value)
}

/* 仓库类型，可更新。目前只支持自有仓，TYPE_OWN：自有物理仓 */
func (this *InventoryStoreManageRequest) SetStoreType(value string) {
	this.Set("store_type", value)
}

type InventoryStoreManageResponse struct {
	*ErrorResponse                     `json:"error_response,omitempty"`
	InventoryStoreManageResponseResult `json:"inventory_store_manage_response"`
}
type InventoryStoreManageResponseResult struct {
	/* 返回结果 */
	StoreList struct {
		Store []*Store `json:"store"`
	} `json:"StoreList"`
}

/*taobao.inventory.store.query*/
type InventoryStoreQueryRequest struct {
	values url.Values
}

func (this *InventoryStoreQueryRequest) GetApiMethodName() string {
	return "taobao.inventory.store.query"
}
func (this *InventoryStoreQueryRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *InventoryStoreQueryRequest) GetValues() url.Values {
	return this.values
}

/* 商家的仓库编码 */
func (this *InventoryStoreQueryRequest) SetStoreCode(value string) {
	this.Set("store_code", value)
}

type InventoryStoreQueryResponse struct {
	*ErrorResponse                    `json:"error_response,omitempty"`
	InventoryStoreQueryResponseResult `json:"inventory_store_query_response"`
}
type InventoryStoreQueryResponseResult struct {
	/* 仓库列表 */
	StoreList struct {
		Store []*Store `json:"store"`
	} `json:"StoreList"`
}

/*taobao.scitem.add*/
type ScitemAddRequest struct {
	values url.Values
}

func (this *ScitemAddRequest) GetApiMethodName() string {
	return "taobao.scitem.add"
}
func (this *ScitemAddRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *ScitemAddRequest) GetValues() url.Values {
	return this.values
}

/* 条形码 */
func (this *ScitemAddRequest) SetBarCode(value string) {
	this.Set("bar_code", value)
}

/* 品牌id */
func (this *ScitemAddRequest) SetBrandId(value string) {
	this.Set("brand_id", value)
}

/* brand_Name */
func (this *ScitemAddRequest) SetBrandName(value string) {
	this.Set("brand_name", value)
}

/* 高 单位：mm */
func (this *ScitemAddRequest) SetHeight(value string) {
	this.Set("height", value)
}

/* 1表示区域销售，0或是空是非区域销售 */
func (this *ScitemAddRequest) SetIsAreaSale(value string) {
	this.Set("is_area_sale", value)
}

/* 是否是贵重品 0:不是 1：是 */
func (this *ScitemAddRequest) SetIsCostly(value string) {
	this.Set("is_costly", value)
}

/* 是否危险 0：不是  1：是 */
func (this *ScitemAddRequest) SetIsDangerous(value string) {
	this.Set("is_dangerous", value)
}

/* 是否易碎 0：不是  1：是 */
func (this *ScitemAddRequest) SetIsFriable(value string) {
	this.Set("is_friable", value)
}

/* 是否保质期：0:不是 1：是 */
func (this *ScitemAddRequest) SetIsWarranty(value string) {
	this.Set("is_warranty", value)
}

/* 商品名称 */
func (this *ScitemAddRequest) SetItemName(value string) {
	this.Set("item_name", value)
}

/* 0.普通供应链商品 1.供应链组合主商品 */
func (this *ScitemAddRequest) SetItemType(value string) {
	this.Set("item_type", value)
}

/* 长度 单位：mm */
func (this *ScitemAddRequest) SetLength(value string) {
	this.Set("length", value)
}

/* 0:液体，1：粉体，2：固体 */
func (this *ScitemAddRequest) SetMatterStatus(value string) {
	this.Set("matter_status", value)
}

/* 商家编码 */
func (this *ScitemAddRequest) SetOuterCode(value string) {
	this.Set("outer_code", value)
}

/* 价格 单位：分 */
func (this *ScitemAddRequest) SetPrice(value string) {
	this.Set("price", value)
}

/* 商品属性格式是  p1:v1,p2:v2,p3:v3 */
func (this *ScitemAddRequest) SetProperties(value string) {
	this.Set("properties", value)
}

/* remark */
func (this *ScitemAddRequest) SetRemark(value string) {
	this.Set("remark", value)
}

/* spuId或是cspuid */
func (this *ScitemAddRequest) SetSpuId(value string) {
	this.Set("spu_id", value)
}

/* 体积：立方厘米 */
func (this *ScitemAddRequest) SetVolume(value string) {
	this.Set("volume", value)
}

/* 重量 单位：g */
func (this *ScitemAddRequest) SetWeight(value string) {
	this.Set("weight", value)
}

/* 宽 单位：mm */
func (this *ScitemAddRequest) SetWidth(value string) {
	this.Set("width", value)
}

/* 仓储商编码 */
func (this *ScitemAddRequest) SetWmsCode(value string) {
	this.Set("wms_code", value)
}

type ScitemAddResponse struct {
	*ErrorResponse          `json:"error_response,omitempty"`
	ScitemAddResponseResult `json:"scitem_add_response"`
}
type ScitemAddResponseResult struct {
	/* 后台商品信息 */
	ScItem *ScItem `json:"sc_item"`
}

/*taobao.scitem.get*/
type ScitemGetRequest struct {
	values url.Values
}

func (this *ScitemGetRequest) GetApiMethodName() string {
	return "taobao.scitem.get"
}
func (this *ScitemGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *ScitemGetRequest) GetValues() url.Values {
	return this.values
}

/* 商品id */
func (this *ScitemGetRequest) SetItemId(value string) {
	this.Set("item_id", value)
}

type ScitemGetResponse struct {
	*ErrorResponse          `json:"error_response,omitempty"`
	ScitemGetResponseResult `json:"scitem_get_response"`
}
type ScitemGetResponseResult struct {
	/* 后端商品 */
	ScItem *ScItem `json:"sc_item"`
}

/*taobao.scitem.map.add*/
type ScitemMapAddRequest struct {
	values url.Values
}

func (this *ScitemMapAddRequest) GetApiMethodName() string {
	return "taobao.scitem.map.add"
}
func (this *ScitemMapAddRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *ScitemMapAddRequest) GetValues() url.Values {
	return this.values
}

/* 前台ic商品id */
func (this *ScitemMapAddRequest) SetItemId(value string) {
	this.Set("item_id", value)
}

/* 默认值为false
true:进行高级校验,前端商品或SKU的商家编码必须与后端商品的商家编码一致，否则会拒绝关联
false:不进行高级校验 */
func (this *ScitemMapAddRequest) SetNeedCheck(value string) {
	this.Set("need_check", value)
}

/* sc_item_id和outer_code 其中一个不能为空 */
func (this *ScitemMapAddRequest) SetOuterCode(value string) {
	this.Set("outer_code", value)
}

/* sc_item_id和outer_code 其中一个不能为空 */
func (this *ScitemMapAddRequest) SetScItemId(value string) {
	this.Set("sc_item_id", value)
}

/* 前台ic商品skuid */
func (this *ScitemMapAddRequest) SetSkuId(value string) {
	this.Set("sku_id", value)
}

type ScitemMapAddResponse struct {
	*ErrorResponse             `json:"error_response,omitempty"`
	ScitemMapAddResponseResult `json:"scitem_map_add_response"`
}
type ScitemMapAddResponseResult struct {
	/* 接口调用返回结果信息：商家编码 */
	OuterCode string `json:"outer_code"`
}

/*taobao.scitem.map.batch.query*/
type ScitemMapBatchQueryRequest struct {
	values url.Values
}

func (this *ScitemMapBatchQueryRequest) GetApiMethodName() string {
	return "taobao.scitem.map.batch.query"
}
func (this *ScitemMapBatchQueryRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *ScitemMapBatchQueryRequest) GetValues() url.Values {
	return this.values
}

/* 后端商品的商家编码 */
func (this *ScitemMapBatchQueryRequest) SetOuterCode(value string) {
	this.Set("outer_code", value)
}

/* 当前页码数 */
func (this *ScitemMapBatchQueryRequest) SetPageIndex(value string) {
	this.Set("page_index", value)
}

/* 分页记录个数，如果用户输入的记录数大于50，则一页显示50条记录 */
func (this *ScitemMapBatchQueryRequest) SetPageSize(value string) {
	this.Set("page_size", value)
}

/* 后端商品id */
func (this *ScitemMapBatchQueryRequest) SetScItemId(value string) {
	this.Set("sc_item_id", value)
}

type ScitemMapBatchQueryResponse struct {
	*ErrorResponse                    `json:"error_response,omitempty"`
	ScitemMapBatchQueryResponseResult `json:"scitem_map_batch_query_response"`
}
type ScitemMapBatchQueryResponseResult struct {
	/* 后端商品映射列表 */
	ScItemMaps struct {
		ScItemMap []*ScItemMap `json:"sc_item_map"`
	} `json:"ScItemMaps"`
}

/*taobao.scitem.map.delete*/
type ScitemMapDeleteRequest struct {
	values url.Values
}

func (this *ScitemMapDeleteRequest) GetApiMethodName() string {
	return "taobao.scitem.map.delete"
}
func (this *ScitemMapDeleteRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *ScitemMapDeleteRequest) GetValues() url.Values {
	return this.values
}

/* 后台商品ID */
func (this *ScitemMapDeleteRequest) SetScItemId(value string) {
	this.Set("sc_item_id", value)
}

/* 店铺用户nick。 如果该参数为空则删除后端商品与当前调用人的商品映射关系;如果不为空则删除指定用户与后端商品的映射关系 */
func (this *ScitemMapDeleteRequest) SetUserNick(value string) {
	this.Set("user_nick", value)
}

type ScitemMapDeleteResponse struct {
	*ErrorResponse                `json:"error_response,omitempty"`
	ScitemMapDeleteResponseResult `json:"scitem_map_delete_response"`
}
type ScitemMapDeleteResponseResult struct {
	/* 失效条数 */
	Module int64 `json:"module"`
}

/*taobao.scitem.map.query*/
type ScitemMapQueryRequest struct {
	values url.Values
}

func (this *ScitemMapQueryRequest) GetApiMethodName() string {
	return "taobao.scitem.map.query"
}
func (this *ScitemMapQueryRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *ScitemMapQueryRequest) GetValues() url.Values {
	return this.values
}

/* map_type为1：前台ic商品id
map_type为2：分销productid */
func (this *ScitemMapQueryRequest) SetItemId(value string) {
	this.Set("item_id", value)
}

/* map_type为1：前台ic商品skuid
map_type为2：分销商品的skuid */
func (this *ScitemMapQueryRequest) SetSkuId(value string) {
	this.Set("sku_id", value)
}

type ScitemMapQueryResponse struct {
	*ErrorResponse               `json:"error_response,omitempty"`
	ScitemMapQueryResponseResult `json:"scitem_map_query_response"`
}
type ScitemMapQueryResponseResult struct {
	/* 后端商品映射列表 */
	ScItemMaps struct {
		ScItemMap []*ScItemMap `json:"sc_item_map"`
	} `json:"ScItemMaps"`
}

/*taobao.scitem.outercode.get*/
type ScitemOutercodeGetRequest struct {
	values url.Values
}

func (this *ScitemOutercodeGetRequest) GetApiMethodName() string {
	return "taobao.scitem.outercode.get"
}
func (this *ScitemOutercodeGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *ScitemOutercodeGetRequest) GetValues() url.Values {
	return this.values
}

/* 商品编码 */
func (this *ScitemOutercodeGetRequest) SetOuterCode(value string) {
	this.Set("outer_code", value)
}

type ScitemOutercodeGetResponse struct {
	*ErrorResponse                   `json:"error_response,omitempty"`
	ScitemOutercodeGetResponseResult `json:"scitem_outercode_get_response"`
}
type ScitemOutercodeGetResponseResult struct {
	/* 后台商品 */
	ScItem *ScItem `json:"sc_item"`
}

/*taobao.scitem.query*/
type ScitemQueryRequest struct {
	values url.Values
}

func (this *ScitemQueryRequest) GetApiMethodName() string {
	return "taobao.scitem.query"
}
func (this *ScitemQueryRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *ScitemQueryRequest) GetValues() url.Values {
	return this.values
}

/* 条形码 */
func (this *ScitemQueryRequest) SetBarCode(value string) {
	this.Set("bar_code", value)
}

/* 商品名称 */
func (this *ScitemQueryRequest) SetItemName(value string) {
	this.Set("item_name", value)
}

/* ITEM类型(只允许输入以下英文或空) NORMAL 0:普通商品; COMBINE 1:是否是组合商品 DISTRIBUTION */
func (this *ScitemQueryRequest) SetItemType(value string) {
	this.Set("item_type", value)
}

/* 商家给商品的一个编码 */
func (this *ScitemQueryRequest) SetOuterCode(value string) {
	this.Set("outer_code", value)
}

/* 当前页码数 */
func (this *ScitemQueryRequest) SetPageIndex(value string) {
	this.Set("page_index", value)
}

/* 分页记录个数，如果用户输入的记录数大于50，则一页显示50条记录 */
func (this *ScitemQueryRequest) SetPageSize(value string) {
	this.Set("page_size", value)
}

/* 仓库编码 */
func (this *ScitemQueryRequest) SetWmsCode(value string) {
	this.Set("wms_code", value)
}

type ScitemQueryResponse struct {
	*ErrorResponse            `json:"error_response,omitempty"`
	ScitemQueryResponseResult `json:"scitem_query_response"`
}
type ScitemQueryResponseResult struct {
	/* 当前所在页数 */
	PageIndex int64 `json:"page_index"`
	/* 返回商品列表的页数 */
	PageSize int64 `json:"page_size"`
	/* 分页 */
	QueryPagination *QueryPagination `json:"query_pagination"`
	/* List<ScItemDO> */
	ScItemList struct {
		ScItem []*ScItem `json:"sc_item"`
	} `json:"ScItemList"`
	/* 商品条数 */
	TotalPage int64 `json:"total_page"`
}

/*taobao.scitem.update*/
type ScitemUpdateRequest struct {
	values url.Values
}

func (this *ScitemUpdateRequest) GetApiMethodName() string {
	return "taobao.scitem.update"
}
func (this *ScitemUpdateRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *ScitemUpdateRequest) GetValues() url.Values {
	return this.values
}

/* 条形码 */
func (this *ScitemUpdateRequest) SetBarCode(value string) {
	this.Set("bar_code", value)
}

/* 品牌id */
func (this *ScitemUpdateRequest) SetBrandId(value string) {
	this.Set("brand_id", value)
}

/* brand_Name */
func (this *ScitemUpdateRequest) SetBrandName(value string) {
	this.Set("brand_name", value)
}

/* 高 单位：mm */
func (this *ScitemUpdateRequest) SetHeight(value string) {
	this.Set("height", value)
}

/* 1表示区域销售，0或是空是非区域销售 */
func (this *ScitemUpdateRequest) SetIsAreaSale(value string) {
	this.Set("is_area_sale", value)
}

/* 是否是贵重品 0:不是 1：是 */
func (this *ScitemUpdateRequest) SetIsCostly(value string) {
	this.Set("is_costly", value)
}

/* 是否危险 0：不是  0：是 */
func (this *ScitemUpdateRequest) SetIsDangerous(value string) {
	this.Set("is_dangerous", value)
}

/* 是否易碎 0：不是  1：是 */
func (this *ScitemUpdateRequest) SetIsFriable(value string) {
	this.Set("is_friable", value)
}

/* 是否保质期：0:不是 1：是 */
func (this *ScitemUpdateRequest) SetIsWarranty(value string) {
	this.Set("is_warranty", value)
}

/* 后端商品ID，跟outer_code必须指定一个 */
func (this *ScitemUpdateRequest) SetItemId(value string) {
	this.Set("item_id", value)
}

/* 商品名称 */
func (this *ScitemUpdateRequest) SetItemName(value string) {
	this.Set("item_name", value)
}

/* 0.普通供应链商品 1.供应链组合主商品 */
func (this *ScitemUpdateRequest) SetItemType(value string) {
	this.Set("item_type", value)
}

/* 长度 单位：mm */
func (this *ScitemUpdateRequest) SetLength(value string) {
	this.Set("length", value)
}

/* 0:液体，1：粉体，2：固体 */
func (this *ScitemUpdateRequest) SetMatterStatus(value string) {
	this.Set("matter_status", value)
}

/* 商家编码，跟item_id必须指定一个 */
func (this *ScitemUpdateRequest) SetOuterCode(value string) {
	this.Set("outer_code", value)
}

/* price */
func (this *ScitemUpdateRequest) SetPrice(value string) {
	this.Set("price", value)
}

/* remark */
func (this *ScitemUpdateRequest) SetRemark(value string) {
	this.Set("remark", value)
}

/* 移除商品属性P列表,P由系统分配：p1；p2 */
func (this *ScitemUpdateRequest) SetRemoveProperties(value string) {
	this.Set("remove_properties", value)
}

/* 淘宝SKU产品级编码CSPU ID */
func (this *ScitemUpdateRequest) SetSpuId(value string) {
	this.Set("spu_id", value)
}

/* 需要更新的商品属性格式是  p1:v1,p2:v2,p3:v3 */
func (this *ScitemUpdateRequest) SetUpdateProperties(value string) {
	this.Set("update_properties", value)
}

/* 体积：立方厘米 */
func (this *ScitemUpdateRequest) SetVolume(value string) {
	this.Set("volume", value)
}

/* weight */
func (this *ScitemUpdateRequest) SetWeight(value string) {
	this.Set("weight", value)
}

/* 宽 单位：mm */
func (this *ScitemUpdateRequest) SetWidth(value string) {
	this.Set("width", value)
}

/* 仓储商编码 */
func (this *ScitemUpdateRequest) SetWmsCode(value string) {
	this.Set("wms_code", value)
}

type ScitemUpdateResponse struct {
	*ErrorResponse             `json:"error_response,omitempty"`
	ScitemUpdateResponseResult `json:"scitem_update_response"`
}
type ScitemUpdateResponseResult struct {
	/* 更新商品数量,1表示成功更新了一条数据，0：表示未找到匹配的数据 */
	UpdateRows int64 `json:"update_rows"`
}

/*taobao.wangwang.eservice.accountstatus.get*/
type WangwangEserviceAccountstatusGetRequest struct {
	values url.Values
}

func (this *WangwangEserviceAccountstatusGetRequest) GetApiMethodName() string {
	return "taobao.wangwang.eservice.accountstatus.get"
}
func (this *WangwangEserviceAccountstatusGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *WangwangEserviceAccountstatusGetRequest) GetValues() url.Values {
	return this.values
}

/* 的起始时间。</br>
格式为： YYYY-mm-dd_HH:MM:SS </br>
注：</br>
1. 查询时间必须在30天之内。
2. 起始时间和结束时间间隔不超过7天。<br /> 支持最大长度为：19<br /> 支持的最大列表长度为：19 */
func (this *WangwangEserviceAccountstatusGetRequest) SetBtime(value string) {
	this.Set("btime", value)
}

/* 结束时间 。 </br>
格式为：YYYY-mm-dd_HH:MM:SS </br>
注：</br>
1. 查询时间必须在30天之内。
2. 起始时间和结束时间间隔不超过7天。<br /> 支持最大长度为：19<br /> 支持的最大列表长度为：19 */
func (this *WangwangEserviceAccountstatusGetRequest) SetEtime(value string) {
	this.Set("etime", value)
}

/* 查询账号（序列）。 需要带前缀（如cntaobao）。</br>

注： </br>
1. uid中所有账号必须为授权店铺的店铺内账号。</br>
2. 如果传入多个账号，账号与账号之间以","分割。最多30个账号。每个账号的最大长度是64。</br>
3. 由于后端依然是顺序请求，所以该接口的响应时长是会随uid的数量线性增长，即uid数量越多，调用时间越长。请合理使用<br /> 支持最大长度为：1920<br /> 支持的最大列表长度为：1920 */
func (this *WangwangEserviceAccountstatusGetRequest) SetUid(value string) {
	this.Set("uid", value)
}

type WangwangEserviceAccountstatusGetResponse struct {
	*ErrorResponse                                 `json:"error_response,omitempty"`
	WangwangEserviceAccountstatusGetResponseResult `json:"wangwang_eservice_accountstatus_get_response"`
}
type WangwangEserviceAccountstatusGetResponseResult struct {
	/* 账号操作记录列表。 */
	AccountInfos struct {
		AccountInfo []*AccountInfo `json:"account_info"`
	} `json:"AccountInfos"`
	/* 接口调用返回码。 0：成功。100：部分成功。 */
	RetCode int64 `json:"ret_code"`
}

/*taobao.wangwang.eservice.avgwaittime.get*/
type WangwangEserviceAvgwaittimeGetRequest struct {
	values url.Values
}

func (this *WangwangEserviceAvgwaittimeGetRequest) GetApiMethodName() string {
	return "taobao.wangwang.eservice.avgwaittime.get"
}
func (this *WangwangEserviceAvgwaittimeGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *WangwangEserviceAvgwaittimeGetRequest) GetValues() url.Values {
	return this.values
}

/* 结束时间 */
func (this *WangwangEserviceAvgwaittimeGetRequest) SetEndDate(value string) {
	this.Set("end_date", value)
}

/* 客服人员id：cntaobao+淘宝nick，例如cntaobaotest<br /> 支持最大长度为：1900<br /> 支持的最大列表长度为：1900 */
func (this *WangwangEserviceAvgwaittimeGetRequest) SetServiceStaffId(value string) {
	this.Set("service_staff_id", value)
}

/* 开始时间 */
func (this *WangwangEserviceAvgwaittimeGetRequest) SetStartDate(value string) {
	this.Set("start_date", value)
}

type WangwangEserviceAvgwaittimeGetResponse struct {
	*ErrorResponse                               `json:"error_response,omitempty"`
	WangwangEserviceAvgwaittimeGetResponseResult `json:"wangwang_eservice_avgwaittime_get_response"`
}
type WangwangEserviceAvgwaittimeGetResponseResult struct {
	/* 平均等待时长 */
	WaitingTimeListOnDays struct {
		WaitingTimesOnDay []*WaitingTimesOnDay `json:"waiting_times_on_day"`
	} `json:"WaitingTimeListOnDays"`
}

/*taobao.wangwang.eservice.evals.get*/
type WangwangEserviceEvalsGetRequest struct {
	values url.Values
}

func (this *WangwangEserviceEvalsGetRequest) GetApiMethodName() string {
	return "taobao.wangwang.eservice.evals.get"
}
func (this *WangwangEserviceEvalsGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *WangwangEserviceEvalsGetRequest) GetValues() url.Values {
	return this.values
}

/* 结束时间 */
func (this *WangwangEserviceEvalsGetRequest) SetEndDate(value string) {
	this.Set("end_date", value)
}

/* 想要查询的账号列表 */
func (this *WangwangEserviceEvalsGetRequest) SetServiceStaffId(value string) {
	this.Set("service_staff_id", value)
}

/* 开始时间 */
func (this *WangwangEserviceEvalsGetRequest) SetStartDate(value string) {
	this.Set("start_date", value)
}

type WangwangEserviceEvalsGetResponse struct {
	*ErrorResponse                         `json:"error_response,omitempty"`
	WangwangEserviceEvalsGetResponseResult `json:"wangwang_eservice_evals_get_response"`
}
type WangwangEserviceEvalsGetResponseResult struct {
	/* 0表示成功 */
	ResultCode int64 `json:"result_code"`
	/* 总条数 */
	ResultCount int64 `json:"result_count"`
	/* 评价具体数据 */
	StaffEvalDetails struct {
		EvalDetail []*EvalDetail `json:"eval_detail"`
	} `json:"StaffEvalDetails"`
}

/*taobao.wangwang.eservice.evaluation.get*/
type WangwangEserviceEvaluationGetRequest struct {
	values url.Values
}

func (this *WangwangEserviceEvaluationGetRequest) GetApiMethodName() string {
	return "taobao.wangwang.eservice.evaluation.get"
}
func (this *WangwangEserviceEvaluationGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *WangwangEserviceEvaluationGetRequest) GetValues() url.Values {
	return this.values
}

/* 查询结束日期 */
func (this *WangwangEserviceEvaluationGetRequest) SetEndDate(value string) {
	this.Set("end_date", value)
}

/* 客服人员id：cntaobao+淘宝nick，例如cntaobaotest<br /> 支持最大长度为：1900<br /> 支持的最大列表长度为：1900 */
func (this *WangwangEserviceEvaluationGetRequest) SetServiceStaffId(value string) {
	this.Set("service_staff_id", value)
}

/* 查询开始日期 */
func (this *WangwangEserviceEvaluationGetRequest) SetStartDate(value string) {
	this.Set("start_date", value)
}

type WangwangEserviceEvaluationGetResponse struct {
	*ErrorResponse                              `json:"error_response,omitempty"`
	WangwangEserviceEvaluationGetResponseResult `json:"wangwang_eservice_evaluation_get_response"`
}
type WangwangEserviceEvaluationGetResponseResult struct {
	/* 客服平均统计列表 */
	StaffEvalStatOnDays struct {
		StaffEvalStatOnDay []*StaffEvalStatOnDay `json:"staff_eval_stat_on_day"`
	} `json:"StaffEvalStatOnDays"`
}

/*taobao.wangwang.eservice.groupmember.get*/
type WangwangEserviceGroupmemberGetRequest struct {
	values url.Values
}

func (this *WangwangEserviceGroupmemberGetRequest) GetApiMethodName() string {
	return "taobao.wangwang.eservice.groupmember.get"
}
func (this *WangwangEserviceGroupmemberGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *WangwangEserviceGroupmemberGetRequest) GetValues() url.Values {
	return this.values
}

/* 主帐号ID：cntaobao+淘宝nick，例如cntaobaotest<br /> 支持最大长度为：128<br /> 支持的最大列表长度为：128 */
func (this *WangwangEserviceGroupmemberGetRequest) SetManagerId(value string) {
	this.Set("manager_id", value)
}

type WangwangEserviceGroupmemberGetResponse struct {
	*ErrorResponse                               `json:"error_response,omitempty"`
	WangwangEserviceGroupmemberGetResponseResult `json:"wangwang_eservice_groupmember_get_response"`
}
type WangwangEserviceGroupmemberGetResponseResult struct {
	/* 组及其成员列表 */
	GroupMemberList struct {
		GroupMember []*GroupMember `json:"group_member"`
	} `json:"GroupMemberList"`
}

/*taobao.wangwang.eservice.loginlogs.get*/
type WangwangEserviceLoginlogsGetRequest struct {
	values url.Values
}

func (this *WangwangEserviceLoginlogsGetRequest) GetApiMethodName() string {
	return "taobao.wangwang.eservice.loginlogs.get"
}
func (this *WangwangEserviceLoginlogsGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *WangwangEserviceLoginlogsGetRequest) GetValues() url.Values {
	return this.values
}

/* 查询登录日志的结束时间，必须按示例的格式，否则会返回错误 */
func (this *WangwangEserviceLoginlogsGetRequest) SetEndDate(value string) {
	this.Set("end_date", value)
}

/* 需要查询登录日志的账号列表 */
func (this *WangwangEserviceLoginlogsGetRequest) SetServiceStaffId(value string) {
	this.Set("service_staff_id", value)
}

/* 查询登录日志的开始日期，必须按示例的格式，否则会返回错误 */
func (this *WangwangEserviceLoginlogsGetRequest) SetStartDate(value string) {
	this.Set("start_date", value)
}

type WangwangEserviceLoginlogsGetResponse struct {
	*ErrorResponse                             `json:"error_response,omitempty"`
	WangwangEserviceLoginlogsGetResponseResult `json:"wangwang_eservice_loginlogs_get_response"`
}
type WangwangEserviceLoginlogsGetResponseResult struct {
	/* 在指定时间段登录日志的条数 */
	Count int64 `json:"count"`
	/* 登录日志具体信息 */
	Loginlogs struct {
		LoginLog []*LoginLog `json:"login_log"`
	} `json:"Loginlogs"`
	/* 被查询的用户id */
	UserId string `json:"user_id"`
}

/*taobao.wangwang.eservice.newevals.get*/
type WangwangEserviceNewevalsGetRequest struct {
	values url.Values
}

func (this *WangwangEserviceNewevalsGetRequest) GetApiMethodName() string {
	return "taobao.wangwang.eservice.newevals.get"
}
func (this *WangwangEserviceNewevalsGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *WangwangEserviceNewevalsGetRequest) GetValues() url.Values {
	return this.values
}

/* 查询登录日志的开始日期，必须按示例的格式，否则会返回错误；
开始时间不能在当前时间30天前，开始时间和结束时间的间隔不能超过7天；
开始时间不能超过当前系统时间 */
func (this *WangwangEserviceNewevalsGetRequest) SetBtime(value string) {
	this.Set("btime", value)
}

/* 查询登录日志的结束时间，必须按示例的格式，否则会返回错误；
结束时间不能小于开始时间，结束时间和开始时间的间隔不能超过7天 */
func (this *WangwangEserviceNewevalsGetRequest) SetEtime(value string) {
	this.Set("etime", value)
}

/* 需要查询登录日志的账号列表，多个id之间用逗号隔开，每次查询的id数不能超过30个 */
func (this *WangwangEserviceNewevalsGetRequest) SetQueryIds(value string) {
	this.Set("query_ids", value)
}

type WangwangEserviceNewevalsGetResponse struct {
	*ErrorResponse                            `json:"error_response,omitempty"`
	WangwangEserviceNewevalsGetResponseResult `json:"wangwang_eservice_newevals_get_response"`
}
type WangwangEserviceNewevalsGetResponseResult struct {
	/* 0表示成功 */
	ResultCode int64 `json:"result_code"`
	/* 总条数 */
	ResultCount int64 `json:"result_count"`
	/* 评价具体数据 */
	StaffEvalDetails struct {
		EvalDetail []*EvalDetail `json:"eval_detail"`
	} `json:"StaffEvalDetails"`
}

/*taobao.wangwang.eservice.noreplynum.get*/
type WangwangEserviceNoreplynumGetRequest struct {
	values url.Values
}

func (this *WangwangEserviceNoreplynumGetRequest) GetApiMethodName() string {
	return "taobao.wangwang.eservice.noreplynum.get"
}
func (this *WangwangEserviceNoreplynumGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *WangwangEserviceNoreplynumGetRequest) GetValues() url.Values {
	return this.values
}

/* 结束日期 */
func (this *WangwangEserviceNoreplynumGetRequest) SetEndDate(value string) {
	this.Set("end_date", value)
}

/* 客服人员id：cntaobao+淘宝nick，例如cntaobaotest<br /> 支持最大长度为：1900<br /> 支持的最大列表长度为：1900 */
func (this *WangwangEserviceNoreplynumGetRequest) SetServiceStaffId(value string) {
	this.Set("service_staff_id", value)
}

/* 开始日期 */
func (this *WangwangEserviceNoreplynumGetRequest) SetStartDate(value string) {
	this.Set("start_date", value)
}

type WangwangEserviceNoreplynumGetResponse struct {
	*ErrorResponse                              `json:"error_response,omitempty"`
	WangwangEserviceNoreplynumGetResponseResult `json:"wangwang_eservice_noreplynum_get_response"`
}
type WangwangEserviceNoreplynumGetResponseResult struct {
	/* 未回复统计列表 */
	NonReplyStatOnDays struct {
		NonReplyStatOnDay []*NonReplyStatOnDay `json:"non_reply_stat_on_day"`
	} `json:"NonReplyStatOnDays"`
}

/*taobao.wangwang.eservice.onlinetime.get*/
type WangwangEserviceOnlinetimeGetRequest struct {
	values url.Values
}

func (this *WangwangEserviceOnlinetimeGetRequest) GetApiMethodName() string {
	return "taobao.wangwang.eservice.onlinetime.get"
}
func (this *WangwangEserviceOnlinetimeGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *WangwangEserviceOnlinetimeGetRequest) GetValues() url.Values {
	return this.values
}

/* 结束日期 */
func (this *WangwangEserviceOnlinetimeGetRequest) SetEndDate(value string) {
	this.Set("end_date", value)
}

/* 客服人员id：cntaobao+淘宝nick，例如cntaobaotest */
func (this *WangwangEserviceOnlinetimeGetRequest) SetServiceStaffId(value string) {
	this.Set("service_staff_id", value)
}

/* 开始日期 */
func (this *WangwangEserviceOnlinetimeGetRequest) SetStartDate(value string) {
	this.Set("start_date", value)
}

type WangwangEserviceOnlinetimeGetResponse struct {
	*ErrorResponse                              `json:"error_response,omitempty"`
	WangwangEserviceOnlinetimeGetResponseResult `json:"wangwang_eservice_onlinetime_get_response"`
}
type WangwangEserviceOnlinetimeGetResponseResult struct {
	/* 客服在线时长（按天统计，排列） */
	OnlineTimesListOnDays struct {
		OnlineTimesOnDay []*OnlineTimesOnDay `json:"online_times_on_day"`
	} `json:"OnlineTimesListOnDays"`
}

/*taobao.wangwang.eservice.receivenum.get*/
type WangwangEserviceReceivenumGetRequest struct {
	values url.Values
}

func (this *WangwangEserviceReceivenumGetRequest) GetApiMethodName() string {
	return "taobao.wangwang.eservice.receivenum.get"
}
func (this *WangwangEserviceReceivenumGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *WangwangEserviceReceivenumGetRequest) GetValues() url.Values {
	return this.values
}

/* 查询接待人数的结束日期 时间精确到日 时分秒可以直接传00:00:00 */
func (this *WangwangEserviceReceivenumGetRequest) SetEndDate(value string) {
	this.Set("end_date", value)
}

/* 客服人员id：cntaobao+淘宝nick，例如cntaobaotest<br /> 支持最大长度为：1900<br /> 支持的最大列表长度为：1900 */
func (this *WangwangEserviceReceivenumGetRequest) SetServiceStaffId(value string) {
	this.Set("service_staff_id", value)
}

/* 查询接待人数的开始日期 时间精确到日 时分秒可直接传 00:00:00 */
func (this *WangwangEserviceReceivenumGetRequest) SetStartDate(value string) {
	this.Set("start_date", value)
}

type WangwangEserviceReceivenumGetResponse struct {
	*ErrorResponse                              `json:"error_response,omitempty"`
	WangwangEserviceReceivenumGetResponseResult `json:"wangwang_eservice_receivenum_get_response"`
}
type WangwangEserviceReceivenumGetResponseResult struct {
	/* 客服回复列表，按天统计，排列 */
	ReplyStatListOnDays struct {
		ReplyStatOnDay []*ReplyStatOnDay `json:"reply_stat_on_day"`
	} `json:"ReplyStatListOnDays"`
}

/*taobao.wangwang.eservice.streamweigths.get*/
type WangwangEserviceStreamweigthsGetRequest struct {
	values url.Values
}

func (this *WangwangEserviceStreamweigthsGetRequest) GetApiMethodName() string {
	return "taobao.wangwang.eservice.streamweigths.get"
}
func (this *WangwangEserviceStreamweigthsGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *WangwangEserviceStreamweigthsGetRequest) GetValues() url.Values {
	return this.values
}

type WangwangEserviceStreamweigthsGetResponse struct {
	*ErrorResponse                                 `json:"error_response,omitempty"`
	WangwangEserviceStreamweigthsGetResponseResult `json:"wangwang_eservice_streamweigths_get_response"`
}
type WangwangEserviceStreamweigthsGetResponseResult struct {
	/* 0表示返回正确 */
	ResultCode int64 `json:"result_code"`
	/* 返回数据条数 */
	ResultCount int64 `json:"result_count"`
	/* 分流权重数据 */
	StaffStreamWeights struct {
		StreamWeight []*StreamWeight `json:"stream_weight"`
	} `json:"StaffStreamWeights"`
	/* 返回数据的总权重，返回数据为空的时候没有这个数字 */
	TotalWeight int64 `json:"total_weight"`
}

/*alibaba.geoip.get*/
type AlibabaGeoipGetRequest struct {
	values url.Values
}

func (this *AlibabaGeoipGetRequest) GetApiMethodName() string {
	return "alibaba.geoip.get"
}
func (this *AlibabaGeoipGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *AlibabaGeoipGetRequest) GetValues() url.Values {
	return this.values
}

/* 要查询的IP地址,与language一起使用，与iplist二选一使用，提供单个IP查询 */
func (this *AlibabaGeoipGetRequest) SetIp(value string) {
	this.Set("ip", value)
}

/* 返回结果的文字语言，cn中文；en英文 */
func (this *AlibabaGeoipGetRequest) SetLanguage(value string) {
	this.Set("language", value)
}

type AlibabaGeoipGetResponse struct {
	*ErrorResponse                `json:"error_response,omitempty"`
	AlibabaGeoipGetResponseResult `json:"alibaba_geoip_get_response"`
}
type AlibabaGeoipGetResponseResult struct {
	/* 城市名称 */
	City string `json:"city"`
	/* 城市id */
	Cityid string `json:"cityid"`
	/* 国家 */
	Country string `json:"country"`
	/* 国家编码 */
	Countryid string `json:"countryid"`
	/* 县 */
	County string `json:"county"`
	/* 县ID */
	Countyid string `json:"countyid"`
	/* IP地址字符串 */
	Ip string `json:"ip"`
	/* 运营商名称 */
	Isp string `json:"isp"`
	/* 运营商id */
	Ispid string `json:"ispid"`
	/* 语言 */
	Lang string `json:"lang"`
	/* 经度 */
	Latitude string `json:"latitude"`
	/* IP地址的整数值 */
	Lip int64 `json:"lip"`
	/* 纬度 */
	Longitude string `json:"longitude"`
	/* 地区名称 */
	Region string `json:"region"`
	/* 地区编码 */
	Regionid string `json:"regionid"`
}

/*taobao.appip.get*/
type AppipGetRequest struct {
	values url.Values
}

func (this *AppipGetRequest) GetApiMethodName() string {
	return "taobao.appip.get"
}
func (this *AppipGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *AppipGetRequest) GetValues() url.Values {
	return this.values
}

type AppipGetResponse struct {
	*ErrorResponse         `json:"error_response,omitempty"`
	AppipGetResponseResult `json:"appip_get_response"`
}
type AppipGetResponseResult struct {
	/* ISV发起请求服务器IP */
	Ip string `json:"ip"`
}

/*taobao.kfc.keyword.search*/
type KfcKeywordSearchRequest struct {
	values url.Values
}

func (this *KfcKeywordSearchRequest) GetApiMethodName() string {
	return "taobao.kfc.keyword.search"
}
func (this *KfcKeywordSearchRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *KfcKeywordSearchRequest) GetValues() url.Values {
	return this.values
}

/* 应用点，分为一级应用点、二级应用点。其中一级应用点通常是指某一个系统或产品，比如淘宝的商品应用（taobao_auction）；二级应用点，是指一级应用点下的具体的分类，比如商品标题(title)、商品描述(content)。不同的二级应用可以设置不同关键词。

这里的apply参数是由一级应用点与二级应用点合起来的字符（一级应用点+"."+二级应用点），如taobao_auction.title。


通常apply参数是不需要传递的。如有特殊需求（比如特殊的过滤需求，需要自己维护一套自己词库），需传递此参数。 */
func (this *KfcKeywordSearchRequest) SetApply(value string) {
	this.Set("apply", value)
}

/* 需要过滤的文本信息 */
func (this *KfcKeywordSearchRequest) SetContent(value string) {
	this.Set("content", value)
}

/* 发布信息的淘宝会员名，可以不传 */
func (this *KfcKeywordSearchRequest) SetNick(value string) {
	this.Set("nick", value)
}

type KfcKeywordSearchResponse struct {
	*ErrorResponse                 `json:"error_response,omitempty"`
	KfcKeywordSearchResponseResult `json:"kfc_keyword_search_response"`
}
type KfcKeywordSearchResponseResult struct {
	/* KFC关键词匹配返回的结果信息 */
	KfcSearchResult *KfcSearchResult `json:"kfc_search_result"`
}

/*taobao.time.get*/
type TimeGetRequest struct {
	values url.Values
}

func (this *TimeGetRequest) GetApiMethodName() string {
	return "taobao.time.get"
}
func (this *TimeGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *TimeGetRequest) GetValues() url.Values {
	return this.values
}

type TimeGetResponse struct {
	*ErrorResponse        `json:"error_response,omitempty"`
	TimeGetResponseResult `json:"time_get_response"`
}
type TimeGetResponseResult struct {
	/* 淘宝系统当前时间。格式:yyyy-MM-dd HH:mm:ss */
	Time string `json:"time"`
}

/*taobao.topats.result.get*/
type TopatsResultGetRequest struct {
	values url.Values
}

func (this *TopatsResultGetRequest) GetApiMethodName() string {
	return "taobao.topats.result.get"
}
func (this *TopatsResultGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *TopatsResultGetRequest) GetValues() url.Values {
	return this.values
}

/* 任务id号，创建任务时返回的task_id */
func (this *TopatsResultGetRequest) SetTaskId(value string) {
	this.Set("task_id", value)
}

type TopatsResultGetResponse struct {
	*ErrorResponse                `json:"error_response,omitempty"`
	TopatsResultGetResponseResult `json:"topats_result_get_response"`
}
type TopatsResultGetResponseResult struct {
	/* 任务结果信息 */
	Task *Task `json:"task"`
}

/*taobao.topats.task.delete*/
type TopatsTaskDeleteRequest struct {
	values url.Values
}

func (this *TopatsTaskDeleteRequest) GetApiMethodName() string {
	return "taobao.topats.task.delete"
}
func (this *TopatsTaskDeleteRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *TopatsTaskDeleteRequest) GetValues() url.Values {
	return this.values
}

/* 需要取消的任务ID */
func (this *TopatsTaskDeleteRequest) SetTaskId(value string) {
	this.Set("task_id", value)
}

type TopatsTaskDeleteResponse struct {
	*ErrorResponse                 `json:"error_response,omitempty"`
	TopatsTaskDeleteResponseResult `json:"topats_task_delete_response"`
}
type TopatsTaskDeleteResponseResult struct {
	/* 表示操作是否成功，是为true，否为false。 */
	IsSuccess bool `json:"is_success"`
}

/*taobao.inventory.ipc.inventorydetail.get*/
type InventoryIpcInventorydetailGetRequest struct {
	values url.Values
}

func (this *InventoryIpcInventorydetailGetRequest) GetApiMethodName() string {
	return "taobao.inventory.ipc.inventorydetail.get"
}
func (this *InventoryIpcInventorydetailGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *InventoryIpcInventorydetailGetRequest) GetValues() url.Values {
	return this.values
}

/* 主订单号，可选<br /> 支持最小值为：0 */
func (this *InventoryIpcInventorydetailGetRequest) SetBizOrderId(value string) {
	this.Set("biz_order_id", value)
}

/* 子订单号，可选<br /> 支持最小值为：0 */
func (this *InventoryIpcInventorydetailGetRequest) SetBizSubOrderId(value string) {
	this.Set("biz_sub_order_id", value)
}

/* 当前页数<br /> 支持最小值为：0 */
func (this *InventoryIpcInventorydetailGetRequest) SetPageIndex(value string) {
	this.Set("page_index", value)
}

/* 一页显示多少条 */
func (this *InventoryIpcInventorydetailGetRequest) SetPageSize(value string) {
	this.Set("page_size", value)
}

/* 仓储商品id */
func (this *InventoryIpcInventorydetailGetRequest) SetScItemId(value string) {
	this.Set("sc_item_id", value)
}

/* 1:查询预扣  4：查询占用<br /> 支持最小值为：1 */
func (this *InventoryIpcInventorydetailGetRequest) SetStatusQuery(value string) {
	this.Set("status_query", value)
}

type InventoryIpcInventorydetailGetResponse struct {
	*ErrorResponse                               `json:"error_response,omitempty"`
	InventoryIpcInventorydetailGetResponseResult `json:"inventory_ipc_inventorydetail_get_response"`
}
type InventoryIpcInventorydetailGetResponseResult struct {
	/* 库存明细列表 */
	InventoryDetails struct {
		IpcInventoryDetailDo []*IpcInventoryDetailDo `json:"ipc_inventory_detail_do"`
	} `json:"InventoryDetails"`
}

/*taobao.logistics.address.reachable*/
type LogisticsAddressReachableRequest struct {
	values url.Values
}

func (this *LogisticsAddressReachableRequest) GetApiMethodName() string {
	return "taobao.logistics.address.reachable"
}
func (this *LogisticsAddressReachableRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *LogisticsAddressReachableRequest) GetValues() url.Values {
	return this.values
}

/* 详细地址 */
func (this *LogisticsAddressReachableRequest) SetAddress(value string) {
	this.Set("address", value)
}

/* 标准区域编码(三级行政区编码或是四级行政区)，可以通过taobao.areas.get获取，如北京市朝阳区为110105 */
func (this *LogisticsAddressReachableRequest) SetAreaCode(value string) {
	this.Set("area_code", value)
}

/* 物流公司编码ID，可以从这个接口获取所有物流公司的标准编码taobao.logistics.companies.get，可以传入多个值，以英文逗号分隔，如“1000000952,1000000953” */
func (this *LogisticsAddressReachableRequest) SetPartnerIds(value string) {
	this.Set("partner_ids", value)
}

/* 服务对应的数字编码，如揽派范围对应88 */
func (this *LogisticsAddressReachableRequest) SetServiceType(value string) {
	this.Set("service_type", value)
}

/* 发货地，标准区域编码(四级行政)，可以通过taobao.areas.get获取，如浙江省杭州市余杭区闲林街道为 330110011 */
func (this *LogisticsAddressReachableRequest) SetSourceAreaCode(value string) {
	this.Set("source_area_code", value)
}

type LogisticsAddressReachableResponse struct {
	*ErrorResponse                          `json:"error_response,omitempty"`
	LogisticsAddressReachableResponseResult `json:"logistics_address_reachable_response"`
}
type LogisticsAddressReachableResponseResult struct {
	/* 地址可达返回结果，每个TP对应一个 */
	ReachableResultList struct {
		AddressReachableResult []*AddressReachableResult `json:"address_reachable_result"`
	} `json:"ReachableResultList"`
}

/*taobao.logistics.address.reachablebatch.get*/
type LogisticsAddressReachablebatchGetRequest struct {
	values url.Values
}

func (this *LogisticsAddressReachablebatchGetRequest) GetApiMethodName() string {
	return "taobao.logistics.address.reachablebatch.get"
}
func (this *LogisticsAddressReachablebatchGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *LogisticsAddressReachablebatchGetRequest) GetValues() url.Values {
	return this.values
}

/* 筛单用户输入地址结构 */
func (this *LogisticsAddressReachablebatchGetRequest) SetAddressList(value string) {
	this.Set("address_list", value)
}

type LogisticsAddressReachablebatchGetResponse struct {
	*ErrorResponse                                  `json:"error_response,omitempty"`
	LogisticsAddressReachablebatchGetResponseResult `json:"logistics_address_reachablebatch_get_response"`
}
type LogisticsAddressReachablebatchGetResponseResult struct {
	/* 物流是否可达结果列表 */
	ReachableResults struct {
		AddressReachableTopResult []*AddressReachableTopResult `json:"address_reachable_top_result"`
	} `json:"ReachableResults"`
}

/*taobao.wlb.icitem.regionsale.operate*/
type WlbIcitemRegionsaleOperateRequest struct {
	values url.Values
}

func (this *WlbIcitemRegionsaleOperateRequest) GetApiMethodName() string {
	return "taobao.wlb.icitem.regionsale.operate"
}
func (this *WlbIcitemRegionsaleOperateRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *WlbIcitemRegionsaleOperateRequest) GetValues() url.Values {
	return this.values
}

/* 前台宝贝id */
func (this *WlbIcitemRegionsaleOperateRequest) SetItemId(value string) {
	this.Set("item_id", value)
}

/* 1:表示添加区域化销售服务
-1:表示去除区域化销售服务 */
func (this *WlbIcitemRegionsaleOperateRequest) SetOpType(value string) {
	this.Set("op_type", value)
}

type WlbIcitemRegionsaleOperateResponse struct {
	*ErrorResponse                           `json:"error_response,omitempty"`
	WlbIcitemRegionsaleOperateResponseResult `json:"wlb_icitem_regionsale_operate_response"`
}
type WlbIcitemRegionsaleOperateResponseResult struct {
	/* 修改时间 */
	GmtModified string `json:"gmt_modified"`
}

/*taobao.wlb.inventory.detail.get*/
type WlbInventoryDetailGetRequest struct {
	values url.Values
}

func (this *WlbInventoryDetailGetRequest) GetApiMethodName() string {
	return "taobao.wlb.inventory.detail.get"
}
func (this *WlbInventoryDetailGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *WlbInventoryDetailGetRequest) GetValues() url.Values {
	return this.values
}

/* 库存类型列表，值包括：
VENDIBLE--可销售库存
FREEZE--冻结库存
ONWAY--在途库存
DEFECT--残次品库存
ENGINE_DAMAGE--机损
BOX_DAMAGE--箱损
EXPIRATION--过保 */
func (this *WlbInventoryDetailGetRequest) SetInventoryTypeList(value string) {
	this.Set("inventory_type_list", value)
}

/* 商品ID */
func (this *WlbInventoryDetailGetRequest) SetItemId(value string) {
	this.Set("item_id", value)
}

/* 仓库编码 */
func (this *WlbInventoryDetailGetRequest) SetStoreCode(value string) {
	this.Set("store_code", value)
}

type WlbInventoryDetailGetResponse struct {
	*ErrorResponse                      `json:"error_response,omitempty"`
	WlbInventoryDetailGetResponseResult `json:"wlb_inventory_detail_get_response"`
}
type WlbInventoryDetailGetResponseResult struct {
	/* 库存详情列表 */
	InventoryList struct {
		WlbInventory []*WlbInventory `json:"wlb_inventory"`
	} `json:"InventoryList"`
	/* 商品ID */
	ItemId int64 `json:"item_id"`
}

/*taobao.wlb.inventory.sync*/
type WlbInventorySyncRequest struct {
	values url.Values
}

func (this *WlbInventorySyncRequest) GetApiMethodName() string {
	return "taobao.wlb.inventory.sync"
}
func (this *WlbInventorySyncRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *WlbInventorySyncRequest) GetValues() url.Values {
	return this.values
}

/* 商品ID */
func (this *WlbInventorySyncRequest) SetItemId(value string) {
	this.Set("item_id", value)
}

/* 外部实体类型.存如下值
IC_ITEM --表示IC商品;
IC_SKU --表示IC最小单位商品;
WLB_ITEM  --表示WLB商品.
若值不在范围内，则按WLB_ITEM处理 */
func (this *WlbInventorySyncRequest) SetItemType(value string) {
	this.Set("item_type", value)
}

/* 库存数量 */
func (this *WlbInventorySyncRequest) SetQuantity(value string) {
	this.Set("quantity", value)
}

type WlbInventorySyncResponse struct {
	*ErrorResponse                 `json:"error_response,omitempty"`
	WlbInventorySyncResponseResult `json:"wlb_inventory_sync_response"`
}
type WlbInventorySyncResponseResult struct {
	/* 修改时间 */
	GmtModified string `json:"gmt_modified"`
}

/*taobao.wlb.inventorylog.query*/
type WlbInventorylogQueryRequest struct {
	values url.Values
}

func (this *WlbInventorylogQueryRequest) GetApiMethodName() string {
	return "taobao.wlb.inventorylog.query"
}
func (this *WlbInventorylogQueryRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *WlbInventorylogQueryRequest) GetValues() url.Values {
	return this.values
}

/* 结束修改时间,小于等于该时间 */
func (this *WlbInventorylogQueryRequest) SetGmtEnd(value string) {
	this.Set("gmt_end", value)
}

/* 起始修改时间,大于等于该时间 */
func (this *WlbInventorylogQueryRequest) SetGmtStart(value string) {
	this.Set("gmt_start", value)
}

/* 商品ID */
func (this *WlbInventorylogQueryRequest) SetItemId(value string) {
	this.Set("item_id", value)
}

/* 库存操作作类型(可以为空)
CHU_KU 1-出库
RU_KU 2-入库
FREEZE 3-冻结
THAW 4-解冻
CHECK_FREEZE 5-冻结确认
CHANGE_KU 6-库存类型变更
若值不在范围内，则按CHU_KU处理 */
func (this *WlbInventorylogQueryRequest) SetOpType(value string) {
	this.Set("op_type", value)
}

/* 可指定授权的用户来查询 */
func (this *WlbInventorylogQueryRequest) SetOpUserId(value string) {
	this.Set("op_user_id", value)
}

/* 单号 */
func (this *WlbInventorylogQueryRequest) SetOrderCode(value string) {
	this.Set("order_code", value)
}

/* 当前页 */
func (this *WlbInventorylogQueryRequest) SetPageNo(value string) {
	this.Set("page_no", value)
}

/* 分页记录个数 */
func (this *WlbInventorylogQueryRequest) SetPageSize(value string) {
	this.Set("page_size", value)
}

/* 仓库编码 */
func (this *WlbInventorylogQueryRequest) SetStoreCode(value string) {
	this.Set("store_code", value)
}

type WlbInventorylogQueryResponse struct {
	*ErrorResponse                     `json:"error_response,omitempty"`
	WlbInventorylogQueryResponseResult `json:"wlb_inventorylog_query_response"`
}
type WlbInventorylogQueryResponseResult struct {
	/* 库存变更记录列表 */
	InventoryLogList struct {
		WlbItemInventoryLog []*WlbItemInventoryLog `json:"wlb_item_inventory_log"`
	} `json:"InventoryLogList"`
	/* 记录数 */
	TotalCount int64 `json:"total_count"`
}

/*taobao.wlb.item.add*/
type WlbItemAddRequest struct {
	values url.Values
}

func (this *WlbItemAddRequest) GetApiMethodName() string {
	return "taobao.wlb.item.add"
}
func (this *WlbItemAddRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *WlbItemAddRequest) GetValues() url.Values {
	return this.values
}

/* 商品颜色 */
func (this *WlbItemAddRequest) SetColor(value string) {
	this.Set("color", value)
}

/* 货类 */
func (this *WlbItemAddRequest) SetGoodsCat(value string) {
	this.Set("goods_cat", value)
}

/* 商品高度，单位毫米 */
func (this *WlbItemAddRequest) SetHeight(value string) {
	this.Set("height", value)
}

/* 是否危险品 */
func (this *WlbItemAddRequest) SetIsDangerous(value string) {
	this.Set("is_dangerous", value)
}

/* 是否易碎品 */
func (this *WlbItemAddRequest) SetIsFriable(value string) {
	this.Set("is_friable", value)
}

/* 是否sku */
func (this *WlbItemAddRequest) SetIsSku(value string) {
	this.Set("is_sku", value)
}

/* 商品编码 */
func (this *WlbItemAddRequest) SetItemCode(value string) {
	this.Set("item_code", value)
}

/* 商品长度，单位毫米 */
func (this *WlbItemAddRequest) SetLength(value string) {
	this.Set("length", value)
}

/* 商品名称 */
func (this *WlbItemAddRequest) SetName(value string) {
	this.Set("name", value)
}

/* 商品包装材料类型 */
func (this *WlbItemAddRequest) SetPackageMaterial(value string) {
	this.Set("package_material", value)
}

/* 商品价格，单位：分 */
func (this *WlbItemAddRequest) SetPrice(value string) {
	this.Set("price", value)
}

/* 计价货类 */
func (this *WlbItemAddRequest) SetPricingCat(value string) {
	this.Set("pricing_cat", value)
}

/* 属性名列表,目前支持的属性：
毛重:GWeight
净重:Nweight
皮重: Tweight
自定义属性：
paramkey1
paramkey2
paramkey3
paramkey4 */
func (this *WlbItemAddRequest) SetProNameList(value string) {
	this.Set("pro_name_list", value)
}

/* 属性值列表：
10,8 */
func (this *WlbItemAddRequest) SetProValueList(value string) {
	this.Set("pro_value_list", value)
}

/* 商品备注 */
func (this *WlbItemAddRequest) SetRemark(value string) {
	this.Set("remark", value)
}

/* 是否支持批次 */
func (this *WlbItemAddRequest) SetSupportBatch(value string) {
	this.Set("support_batch", value)
}

/* 商品标题 */
func (this *WlbItemAddRequest) SetTitle(value string) {
	this.Set("title", value)
}

/* NORMAL--普通商品
COMBINE--组合商品
DISTRIBUTION--分销 */
func (this *WlbItemAddRequest) SetType(value string) {
	this.Set("type", value)
}

/* 商品体积，单位立方厘米 */
func (this *WlbItemAddRequest) SetVolume(value string) {
	this.Set("volume", value)
}

/* 商品重量，单位G */
func (this *WlbItemAddRequest) SetWeight(value string) {
	this.Set("weight", value)
}

/* 商品宽度，单位毫米 */
func (this *WlbItemAddRequest) SetWidth(value string) {
	this.Set("width", value)
}

type WlbItemAddResponse struct {
	*ErrorResponse           `json:"error_response,omitempty"`
	WlbItemAddResponseResult `json:"wlb_item_add_response"`
}
type WlbItemAddResponseResult struct {
	/* 新增的商品 */
	ItemId int64 `json:"item_id"`
}

/*taobao.wlb.item.authorization.add*/
type WlbItemAuthorizationAddRequest struct {
	values url.Values
}

func (this *WlbItemAuthorizationAddRequest) GetApiMethodName() string {
	return "taobao.wlb.item.authorization.add"
}
func (this *WlbItemAuthorizationAddRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *WlbItemAuthorizationAddRequest) GetValues() url.Values {
	return this.values
}

/* 授权类型：1=全量授权，0=部分授权
当部分授权时，需要指定授权数量quantity */
func (this *WlbItemAuthorizationAddRequest) SetAuthType(value string) {
	this.Set("auth_type", value)
}

/* 授权结束时间 */
func (this *WlbItemAuthorizationAddRequest) SetAuthorizeEndTime(value string) {
	this.Set("authorize_end_time", value)
}

/* 授权开始时间 */
func (this *WlbItemAuthorizationAddRequest) SetAuthorizeStartTime(value string) {
	this.Set("authorize_start_time", value)
}

/* 被授权人用户id */
func (this *WlbItemAuthorizationAddRequest) SetConsignUserNick(value string) {
	this.Set("consign_user_nick", value)
}

/* 商品id列表，以英文逗号,分隔，最多可放入50个商品ID。 */
func (this *WlbItemAuthorizationAddRequest) SetItemIdList(value string) {
	this.Set("item_id_list", value)
}

/* 规则名称<br /> 支持最大长度为：255<br /> 支持的最大列表长度为：255 */
func (this *WlbItemAuthorizationAddRequest) SetName(value string) {
	this.Set("name", value)
}

/* 授权数量 */
func (this *WlbItemAuthorizationAddRequest) SetQuantity(value string) {
	this.Set("quantity", value)
}

/* 授权规则：目前只有GRANT_FIX，按照数量授权 */
func (this *WlbItemAuthorizationAddRequest) SetRuleCode(value string) {
	this.Set("rule_code", value)
}

type WlbItemAuthorizationAddResponse struct {
	*ErrorResponse                        `json:"error_response,omitempty"`
	WlbItemAuthorizationAddResponseResult `json:"wlb_item_authorization_add_response"`
}
type WlbItemAuthorizationAddResponseResult struct {
	/* 授权规则ID列表 */
	RuleIdList []int64 `json:"rule_id_list"`
}

/*taobao.wlb.item.authorization.delete*/
type WlbItemAuthorizationDeleteRequest struct {
	values url.Values
}

func (this *WlbItemAuthorizationDeleteRequest) GetApiMethodName() string {
	return "taobao.wlb.item.authorization.delete"
}
func (this *WlbItemAuthorizationDeleteRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *WlbItemAuthorizationDeleteRequest) GetValues() url.Values {
	return this.values
}

/* 授权关系ID */
func (this *WlbItemAuthorizationDeleteRequest) SetAuthorizeId(value string) {
	this.Set("authorize_id", value)
}

type WlbItemAuthorizationDeleteResponse struct {
	*ErrorResponse                           `json:"error_response,omitempty"`
	WlbItemAuthorizationDeleteResponseResult `json:"wlb_item_authorization_delete_response"`
}
type WlbItemAuthorizationDeleteResponseResult struct {
	/* 修改时间 */
	GmtModified string `json:"gmt_modified"`
}

/*taobao.wlb.item.authorization.query*/
type WlbItemAuthorizationQueryRequest struct {
	values url.Values
}

func (this *WlbItemAuthorizationQueryRequest) GetApiMethodName() string {
	return "taobao.wlb.item.authorization.query"
}
func (this *WlbItemAuthorizationQueryRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *WlbItemAuthorizationQueryRequest) GetValues() url.Values {
	return this.values
}

/* 授权商品ID */
func (this *WlbItemAuthorizationQueryRequest) SetItemId(value string) {
	this.Set("item_id", value)
}

/* 授权名称<br /> 支持最大长度为：255<br /> 支持的最大列表长度为：255 */
func (this *WlbItemAuthorizationQueryRequest) SetName(value string) {
	this.Set("name", value)
}

/* 当前页 */
func (this *WlbItemAuthorizationQueryRequest) SetPageNo(value string) {
	this.Set("page_no", value)
}

/* 分页记录个数，如果用户输入的记录数大于50，则一页显示50条记录 */
func (this *WlbItemAuthorizationQueryRequest) SetPageSize(value string) {
	this.Set("page_size", value)
}

/* 授权编码 */
func (this *WlbItemAuthorizationQueryRequest) SetRuleCode(value string) {
	this.Set("rule_code", value)
}

/* 状态： 只能输入如下值,范围外的默认按VALID处理;不选则查询所有;
VALID -- 1 有效； INVALIDATION -- 2 失效 */
func (this *WlbItemAuthorizationQueryRequest) SetStatus(value string) {
	this.Set("status", value)
}

/* 类型：可由不同角色来查询，默认值OWNER,
OWNER -- 授权人,
ON_COMMISSION -- 被授权人 */
func (this *WlbItemAuthorizationQueryRequest) SetType(value string) {
	this.Set("type", value)
}

type WlbItemAuthorizationQueryResponse struct {
	*ErrorResponse                          `json:"error_response,omitempty"`
	WlbItemAuthorizationQueryResponseResult `json:"wlb_item_authorization_query_response"`
}
type WlbItemAuthorizationQueryResponseResult struct {
	/* 授权关系列表 */
	AuthorizationList struct {
		WlbAuthorization []*WlbAuthorization `json:"wlb_authorization"`
	} `json:"AuthorizationList"`
	/* 结果总数 */
	TotalCount int64 `json:"total_count"`
}

/*taobao.wlb.item.batch.query*/
type WlbItemBatchQueryRequest struct {
	values url.Values
}

func (this *WlbItemBatchQueryRequest) GetApiMethodName() string {
	return "taobao.wlb.item.batch.query"
}
func (this *WlbItemBatchQueryRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *WlbItemBatchQueryRequest) GetValues() url.Values {
	return this.values
}

/* 需要查询的商品ID列表，以字符串表示，ID间以;隔开 */
func (this *WlbItemBatchQueryRequest) SetItemIds(value string) {
	this.Set("item_ids", value)
}

/* 分页查询参数，指定查询页数，默认为1 */
func (this *WlbItemBatchQueryRequest) SetPageNo(value string) {
	this.Set("page_no", value)
}

/* 分页查询参数，每页查询数量，默认20，最大值50,大于50时按照50条查询 */
func (this *WlbItemBatchQueryRequest) SetPageSize(value string) {
	this.Set("page_size", value)
}

/* 仓库编号 */
func (this *WlbItemBatchQueryRequest) SetStoreCode(value string) {
	this.Set("store_code", value)
}

type WlbItemBatchQueryResponse struct {
	*ErrorResponse                  `json:"error_response,omitempty"`
	WlbItemBatchQueryResponseResult `json:"wlb_item_batch_query_response"`
}
type WlbItemBatchQueryResponseResult struct {
	/* 商品库存及批次信息查询结果 */
	ItemInventoryBatchList struct {
		WlbItemBatchInventory []*WlbItemBatchInventory `json:"wlb_item_batch_inventory"`
	} `json:"ItemInventoryBatchList"`
	/* 返回结果记录的数量 */
	TotalCount int64 `json:"total_count"`
}

/*taobao.wlb.item.combination.create*/
type WlbItemCombinationCreateRequest struct {
	values url.Values
}

func (this *WlbItemCombinationCreateRequest) GetApiMethodName() string {
	return "taobao.wlb.item.combination.create"
}
func (this *WlbItemCombinationCreateRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *WlbItemCombinationCreateRequest) GetValues() url.Values {
	return this.values
}

/* 组合商品的id列表 */
func (this *WlbItemCombinationCreateRequest) SetDestItemList(value string) {
	this.Set("dest_item_list", value)
}

/* 要建立组合关系的商品id */
func (this *WlbItemCombinationCreateRequest) SetItemId(value string) {
	this.Set("item_id", value)
}

/* 组成组合商品的比例列表，描述组合商品的组合比例，默认为1,1,1 */
func (this *WlbItemCombinationCreateRequest) SetProportionList(value string) {
	this.Set("proportion_list", value)
}

type WlbItemCombinationCreateResponse struct {
	*ErrorResponse                         `json:"error_response,omitempty"`
	WlbItemCombinationCreateResponseResult `json:"wlb_item_combination_create_response"`
}
type WlbItemCombinationCreateResponseResult struct {
	/* 组合关系创建时间 */
	GmtCreate string `json:"gmt_create"`
}

/*taobao.wlb.item.combination.delete*/
type WlbItemCombinationDeleteRequest struct {
	values url.Values
}

func (this *WlbItemCombinationDeleteRequest) GetApiMethodName() string {
	return "taobao.wlb.item.combination.delete"
}
func (this *WlbItemCombinationDeleteRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *WlbItemCombinationDeleteRequest) GetValues() url.Values {
	return this.values
}

/* 组合商品的id列表 */
func (this *WlbItemCombinationDeleteRequest) SetDestItemList(value string) {
	this.Set("dest_item_list", value)
}

/* 组合关系的商品id */
func (this *WlbItemCombinationDeleteRequest) SetItemId(value string) {
	this.Set("item_id", value)
}

type WlbItemCombinationDeleteResponse struct {
	*ErrorResponse                         `json:"error_response,omitempty"`
	WlbItemCombinationDeleteResponseResult `json:"wlb_item_combination_delete_response"`
}
type WlbItemCombinationDeleteResponseResult struct {
	/* 修改时间 */
	GmtModified string `json:"gmt_modified"`
}

/*taobao.wlb.item.combination.get*/
type WlbItemCombinationGetRequest struct {
	values url.Values
}

func (this *WlbItemCombinationGetRequest) GetApiMethodName() string {
	return "taobao.wlb.item.combination.get"
}
func (this *WlbItemCombinationGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *WlbItemCombinationGetRequest) GetValues() url.Values {
	return this.values
}

/* 要查询的组合商品id */
func (this *WlbItemCombinationGetRequest) SetItemId(value string) {
	this.Set("item_id", value)
}

type WlbItemCombinationGetResponse struct {
	*ErrorResponse                      `json:"error_response,omitempty"`
	WlbItemCombinationGetResponseResult `json:"wlb_item_combination_get_response"`
}
type WlbItemCombinationGetResponseResult struct {
	/* 组合子商品id列表 */
	ItemIdList []int64 `json:"item_id_list"`
}

/*taobao.wlb.item.consignment.create*/
type WlbItemConsignmentCreateRequest struct {
	values url.Values
}

func (this *WlbItemConsignmentCreateRequest) GetApiMethodName() string {
	return "taobao.wlb.item.consignment.create"
}
func (this *WlbItemConsignmentCreateRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *WlbItemConsignmentCreateRequest) GetValues() url.Values {
	return this.values
}

/* 商品id */
func (this *WlbItemConsignmentCreateRequest) SetItemId(value string) {
	this.Set("item_id", value)
}

/* 代销数量 */
func (this *WlbItemConsignmentCreateRequest) SetNumber(value string) {
	this.Set("number", value)
}

/* 货主商品id */
func (this *WlbItemConsignmentCreateRequest) SetOwnerItemId(value string) {
	this.Set("owner_item_id", value)
}

/* 货主id */
func (this *WlbItemConsignmentCreateRequest) SetOwnerUserId(value string) {
	this.Set("owner_user_id", value)
}

/* 通过taobao.wlb.item.authorization.add接口创建后得到的rule_id，规则中设定了代销商可以代销的商品数量 */
func (this *WlbItemConsignmentCreateRequest) SetRuleId(value string) {
	this.Set("rule_id", value)
}

type WlbItemConsignmentCreateResponse struct {
	*ErrorResponse                         `json:"error_response,omitempty"`
	WlbItemConsignmentCreateResponseResult `json:"wlb_item_consignment_create_response"`
}
type WlbItemConsignmentCreateResponseResult struct {
	/* 代销关系唯一标识 */
	ConsignmentId int64 `json:"consignment_id"`
}

/*taobao.wlb.item.consignment.delete*/
type WlbItemConsignmentDeleteRequest struct {
	values url.Values
}

func (this *WlbItemConsignmentDeleteRequest) GetApiMethodName() string {
	return "taobao.wlb.item.consignment.delete"
}
func (this *WlbItemConsignmentDeleteRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *WlbItemConsignmentDeleteRequest) GetValues() url.Values {
	return this.values
}

/* 代销商前台宝贝ID */
func (this *WlbItemConsignmentDeleteRequest) SetIcItemId(value string) {
	this.Set("ic_item_id", value)
}

/* 货主的物流宝商品ID */
func (this *WlbItemConsignmentDeleteRequest) SetOwnerItemId(value string) {
	this.Set("owner_item_id", value)
}

/* 授权关系id */
func (this *WlbItemConsignmentDeleteRequest) SetRuleId(value string) {
	this.Set("rule_id", value)
}

type WlbItemConsignmentDeleteResponse struct {
	*ErrorResponse                         `json:"error_response,omitempty"`
	WlbItemConsignmentDeleteResponseResult `json:"wlb_item_consignment_delete_response"`
}
type WlbItemConsignmentDeleteResponseResult struct {
	/* 修改时间 */
	GmtModified string `json:"gmt_modified"`
}

/*taobao.wlb.item.consignment.page.get*/
type WlbItemConsignmentPageGetRequest struct {
	values url.Values
}

func (this *WlbItemConsignmentPageGetRequest) GetApiMethodName() string {
	return "taobao.wlb.item.consignment.page.get"
}
func (this *WlbItemConsignmentPageGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *WlbItemConsignmentPageGetRequest) GetValues() url.Values {
	return this.values
}

/* 代销商宝贝id */
func (this *WlbItemConsignmentPageGetRequest) SetIcItemId(value string) {
	this.Set("ic_item_id", value)
}

/* 供应商商品id */
func (this *WlbItemConsignmentPageGetRequest) SetOwnerItemId(value string) {
	this.Set("owner_item_id", value)
}

/* 供应商用户昵称 */
func (this *WlbItemConsignmentPageGetRequest) SetOwnerUserNick(value string) {
	this.Set("owner_user_nick", value)
}

type WlbItemConsignmentPageGetResponse struct {
	*ErrorResponse                          `json:"error_response,omitempty"`
	WlbItemConsignmentPageGetResponseResult `json:"wlb_item_consignment_page_get_response"`
}
type WlbItemConsignmentPageGetResponseResult struct {
	/* 条件查询结果总数 */
	TotalCount int64 `json:"total_count"`
	/* 代销关系列表 */
	WlbConsignMents struct {
		WlbConsignMent []*WlbConsignMent `json:"wlb_consign_ment"`
	} `json:"WlbConsignMents"`
}

/*taobao.wlb.item.consignment.query*/
type WlbItemConsignmentQueryRequest struct {
	values url.Values
}

func (this *WlbItemConsignmentQueryRequest) GetApiMethodName() string {
	return "taobao.wlb.item.consignment.query"
}
func (this *WlbItemConsignmentQueryRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *WlbItemConsignmentQueryRequest) GetValues() url.Values {
	return this.values
}

/* 授权结束时间 */
func (this *WlbItemConsignmentQueryRequest) SetAuthorizeEndTime(value string) {
	this.Set("authorize_end_time", value)
}

/* 授权开始时间 */
func (this *WlbItemConsignmentQueryRequest) SetAuthorizeStartTime(value string) {
	this.Set("authorize_start_time", value)
}

/* 货主的用户昵称，未设置则查询全部 */
func (this *WlbItemConsignmentQueryRequest) SetOwnerUserNick(value string) {
	this.Set("owner_user_nick", value)
}

/* 当前页 */
func (this *WlbItemConsignmentQueryRequest) SetPageNo(value string) {
	this.Set("page_no", value)
}

/* 分页记录个数，如果用户输入的记录数大于50，则一页显示50条记录 */
func (this *WlbItemConsignmentQueryRequest) SetPageSize(value string) {
	this.Set("page_size", value)
}

type WlbItemConsignmentQueryResponse struct {
	*ErrorResponse                        `json:"error_response,omitempty"`
	WlbItemConsignmentQueryResponseResult `json:"wlb_item_consignment_query_response"`
}
type WlbItemConsignmentQueryResponseResult struct {
	/* 商品信息列表 */
	ItemList struct {
		WlbItem []*WlbItem `json:"wlb_item"`
	} `json:"ItemList"`
	/* 结果总数 */
	TotalCount int64 `json:"total_count"`
}

/*taobao.wlb.item.delete*/
type WlbItemDeleteRequest struct {
	values url.Values
}

func (this *WlbItemDeleteRequest) GetApiMethodName() string {
	return "taobao.wlb.item.delete"
}
func (this *WlbItemDeleteRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *WlbItemDeleteRequest) GetValues() url.Values {
	return this.values
}

/* 商品ID */
func (this *WlbItemDeleteRequest) SetItemId(value string) {
	this.Set("item_id", value)
}

/* 商品所有人淘宝nick */
func (this *WlbItemDeleteRequest) SetUserNick(value string) {
	this.Set("user_nick", value)
}

type WlbItemDeleteResponse struct {
	*ErrorResponse              `json:"error_response,omitempty"`
	WlbItemDeleteResponseResult `json:"wlb_item_delete_response"`
}
type WlbItemDeleteResponseResult struct {
	/* 修改时间 */
	GmtModified string `json:"gmt_modified"`
}

/*taobao.wlb.item.get*/
type WlbItemGetRequest struct {
	values url.Values
}

func (this *WlbItemGetRequest) GetApiMethodName() string {
	return "taobao.wlb.item.get"
}
func (this *WlbItemGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *WlbItemGetRequest) GetValues() url.Values {
	return this.values
}

/* 商品ID<br /> 支持的最大列表长度为：20 */
func (this *WlbItemGetRequest) SetItemId(value string) {
	this.Set("item_id", value)
}

type WlbItemGetResponse struct {
	*ErrorResponse           `json:"error_response,omitempty"`
	WlbItemGetResponseResult `json:"wlb_item_get_response"`
}
type WlbItemGetResponseResult struct {
	/* 商品信息 */
	Item *WlbItem `json:"item"`
}

/*taobao.wlb.item.map.get*/
type WlbItemMapGetRequest struct {
	values url.Values
}

func (this *WlbItemMapGetRequest) GetApiMethodName() string {
	return "taobao.wlb.item.map.get"
}
func (this *WlbItemMapGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *WlbItemMapGetRequest) GetValues() url.Values {
	return this.values
}

/* 要查询映射关系的物流宝商品id */
func (this *WlbItemMapGetRequest) SetItemId(value string) {
	this.Set("item_id", value)
}

type WlbItemMapGetResponse struct {
	*ErrorResponse              `json:"error_response,omitempty"`
	WlbItemMapGetResponseResult `json:"wlb_item_map_get_response"`
}
type WlbItemMapGetResponseResult struct {
	/* 外部商品实体列表 */
	OutEntityItemList struct {
		OutEntityItem []*OutEntityItem `json:"out_entity_item"`
	} `json:"OutEntityItemList"`
}

/*taobao.wlb.item.map.get.by.extentity*/
type WlbItemMapGetByExtentityRequest struct {
	values url.Values
}

func (this *WlbItemMapGetByExtentityRequest) GetApiMethodName() string {
	return "taobao.wlb.item.map.get.by.extentity"
}
func (this *WlbItemMapGetByExtentityRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *WlbItemMapGetByExtentityRequest) GetValues() url.Values {
	return this.values
}

/* 外部实体类型对应的商品id */
func (this *WlbItemMapGetByExtentityRequest) SetExtEntityId(value string) {
	this.Set("ext_entity_id", value)
}

/* 外部实体类型： IC_ITEM--ic商品 IC_SKU--ic销售单元 */
func (this *WlbItemMapGetByExtentityRequest) SetExtEntityType(value string) {
	this.Set("ext_entity_type", value)
}

type WlbItemMapGetByExtentityResponse struct {
	*ErrorResponse                         `json:"error_response,omitempty"`
	WlbItemMapGetByExtentityResponseResult `json:"wlb_item_map_get_by_extentity_response"`
}
type WlbItemMapGetByExtentityResponseResult struct {
	/* 物流宝商品id */
	ItemId int64 `json:"item_id"`
}

/*taobao.wlb.item.query*/
type WlbItemQueryRequest struct {
	values url.Values
}

func (this *WlbItemQueryRequest) GetApiMethodName() string {
	return "taobao.wlb.item.query"
}
func (this *WlbItemQueryRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *WlbItemQueryRequest) GetValues() url.Values {
	return this.values
}

/* 是否是最小库存单元，只有最小库存单元的商品才可以有库存,值只能给"true","false"来表示;
若值不在范围内，则按true处理 */
func (this *WlbItemQueryRequest) SetIsSku(value string) {
	this.Set("is_sku", value)
}

/* 商家编码<br /> 支持最大长度为：64<br /> 支持的最大列表长度为：64 */
func (this *WlbItemQueryRequest) SetItemCode(value string) {
	this.Set("item_code", value)
}

/* ITEM类型(只允许输入以下英文或空)
NORMAL  0:普通商品;
COMBINE  1:是否是组合商品
DISTRIBUTION  2:是否是分销商品(货主是别人)
若值不在范围内，则按NORMAL处理 */
func (this *WlbItemQueryRequest) SetItemType(value string) {
	this.Set("item_type", value)
}

/* 商品名称 */
func (this *WlbItemQueryRequest) SetName(value string) {
	this.Set("name", value)
}

/* 当前页<br /> 支持最小值为：1 */
func (this *WlbItemQueryRequest) SetPageNo(value string) {
	this.Set("page_no", value)
}

/* 分页记录个数，如果用户输入的记录数大于50，则一页显示50条记录<br /> 支持最大值为：50<br /> 支持最小值为：1 */
func (this *WlbItemQueryRequest) SetPageSize(value string) {
	this.Set("page_size", value)
}

/* 父ID,只有is_sku=1时才能有父ID，商品也可以没有付商品 */
func (this *WlbItemQueryRequest) SetParentId(value string) {
	this.Set("parent_id", value)
}

/* 只能输入以下值或空：
ITEM_STATUS_VALID -- 1 表示 有效；
ITEM_STATUS_LOCK  -- 2 表示锁住。
若值不在范围内，按ITEM_STATUS_VALID处理 */
func (this *WlbItemQueryRequest) SetStatus(value string) {
	this.Set("status", value)
}

/* 商品前台销售名字<br /> 支持最大长度为：255<br /> 支持的最大列表长度为：255 */
func (this *WlbItemQueryRequest) SetTitle(value string) {
	this.Set("title", value)
}

type WlbItemQueryResponse struct {
	*ErrorResponse             `json:"error_response,omitempty"`
	WlbItemQueryResponseResult `json:"wlb_item_query_response"`
}
type WlbItemQueryResponseResult struct {
	/* 商品信息列表 */
	ItemList struct {
		WlbItem []*WlbItem `json:"wlb_item"`
	} `json:"ItemList"`
	/* 结果总数 */
	TotalCount int64 `json:"total_count"`
}

/*taobao.wlb.item.synchronize*/
type WlbItemSynchronizeRequest struct {
	values url.Values
}

func (this *WlbItemSynchronizeRequest) GetApiMethodName() string {
	return "taobao.wlb.item.synchronize"
}
func (this *WlbItemSynchronizeRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *WlbItemSynchronizeRequest) GetValues() url.Values {
	return this.values
}

/* 外部实体ID<br /> 支持的最大列表长度为：64 */
func (this *WlbItemSynchronizeRequest) SetExtEntityId(value string) {
	this.Set("ext_entity_id", value)
}

/* 外部实体类型.存如下值
IC_ITEM   --表示IC商品
IC_SKU    --表示IC最小单位商品
若输入其他值，则按IC_ITEM处理 */
func (this *WlbItemSynchronizeRequest) SetExtEntityType(value string) {
	this.Set("ext_entity_type", value)
}

/* 商品ID<br /> 支持的最大列表长度为：20 */
func (this *WlbItemSynchronizeRequest) SetItemId(value string) {
	this.Set("item_id", value)
}

/* 商品所有人淘宝nick<br /> 支持最大长度为：64<br /> 支持的最大列表长度为：64 */
func (this *WlbItemSynchronizeRequest) SetUserNick(value string) {
	this.Set("user_nick", value)
}

type WlbItemSynchronizeResponse struct {
	*ErrorResponse                   `json:"error_response,omitempty"`
	WlbItemSynchronizeResponseResult `json:"wlb_item_synchronize_response"`
}
type WlbItemSynchronizeResponseResult struct {
	/* 修改时间 */
	GmtModified string `json:"gmt_modified"`
}

/*taobao.wlb.item.synchronize.delete*/
type WlbItemSynchronizeDeleteRequest struct {
	values url.Values
}

func (this *WlbItemSynchronizeDeleteRequest) GetApiMethodName() string {
	return "taobao.wlb.item.synchronize.delete"
}
func (this *WlbItemSynchronizeDeleteRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *WlbItemSynchronizeDeleteRequest) GetValues() url.Values {
	return this.values
}

/* 外部实体ID<br /> 支持的最大列表长度为：64 */
func (this *WlbItemSynchronizeDeleteRequest) SetExtEntityId(value string) {
	this.Set("ext_entity_id", value)
}

/* 外部实体类型.存如下值 IC_ITEM --表示IC商品; IC_SKU --表示IC最小单位商品;若输入其他值，则按IC_ITEM处理 */
func (this *WlbItemSynchronizeDeleteRequest) SetExtEntityType(value string) {
	this.Set("ext_entity_type", value)
}

/* 物流宝商品ID<br /> 支持的最大列表长度为：20 */
func (this *WlbItemSynchronizeDeleteRequest) SetItemId(value string) {
	this.Set("item_id", value)
}

type WlbItemSynchronizeDeleteResponse struct {
	*ErrorResponse                         `json:"error_response,omitempty"`
	WlbItemSynchronizeDeleteResponseResult `json:"wlb_item_synchronize_delete_response"`
}
type WlbItemSynchronizeDeleteResponseResult struct {
	/* 修改时间 */
	GmtModified string `json:"gmt_modified"`
}

/*taobao.wlb.item.update*/
type WlbItemUpdateRequest struct {
	values url.Values
}

func (this *WlbItemUpdateRequest) GetApiMethodName() string {
	return "taobao.wlb.item.update"
}
func (this *WlbItemUpdateRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *WlbItemUpdateRequest) GetValues() url.Values {
	return this.values
}

/* 商品颜色 */
func (this *WlbItemUpdateRequest) SetColor(value string) {
	this.Set("color", value)
}

/* 需要删除的商品属性key列表 */
func (this *WlbItemUpdateRequest) SetDeletePropertyKeyList(value string) {
	this.Set("delete_property_key_list", value)
}

/* 商品货类 */
func (this *WlbItemUpdateRequest) SetGoodsCat(value string) {
	this.Set("goods_cat", value)
}

/* 商品高度，单位厘米 */
func (this *WlbItemUpdateRequest) SetHeight(value string) {
	this.Set("height", value)
}

/* 要修改的商品id */
func (this *WlbItemUpdateRequest) SetId(value string) {
	this.Set("id", value)
}

/* 是否危险品 */
func (this *WlbItemUpdateRequest) SetIsDangerous(value string) {
	this.Set("is_dangerous", value)
}

/* 是否易碎品 */
func (this *WlbItemUpdateRequest) SetIsFriable(value string) {
	this.Set("is_friable", value)
}

/* 商品长度，单位厘米 */
func (this *WlbItemUpdateRequest) SetLength(value string) {
	this.Set("length", value)
}

/* 要修改的商品名称 */
func (this *WlbItemUpdateRequest) SetName(value string) {
	this.Set("name", value)
}

/* 商品包装材料类型 */
func (this *WlbItemUpdateRequest) SetPackageMaterial(value string) {
	this.Set("package_material", value)
}

/* 商品计价货类 */
func (this *WlbItemUpdateRequest) SetPricingCat(value string) {
	this.Set("pricing_cat", value)
}

/* 要修改的商品备注 */
func (this *WlbItemUpdateRequest) SetRemark(value string) {
	this.Set("remark", value)
}

/* 要修改的商品标题 */
func (this *WlbItemUpdateRequest) SetTitle(value string) {
	this.Set("title", value)
}

/* 需要修改的商品属性值的列表，如果属性不存在，则新增属性 */
func (this *WlbItemUpdateRequest) SetUpdatePropertyKeyList(value string) {
	this.Set("update_property_key_list", value)
}

/* 需要修改的属性值的列表 */
func (this *WlbItemUpdateRequest) SetUpdatePropertyValueList(value string) {
	this.Set("update_property_value_list", value)
}

/* 商品体积，单位立方厘米 */
func (this *WlbItemUpdateRequest) SetVolume(value string) {
	this.Set("volume", value)
}

/* 商品重量，单位G */
func (this *WlbItemUpdateRequest) SetWeight(value string) {
	this.Set("weight", value)
}

/* 商品宽度，单位厘米 */
func (this *WlbItemUpdateRequest) SetWidth(value string) {
	this.Set("width", value)
}

type WlbItemUpdateResponse struct {
	*ErrorResponse              `json:"error_response,omitempty"`
	WlbItemUpdateResponseResult `json:"wlb_item_update_response"`
}
type WlbItemUpdateResponseResult struct {
	/* 修改时间 */
	GmtModified string `json:"gmt_modified"`
}

/*taobao.wlb.notify.message.confirm*/
type WlbNotifyMessageConfirmRequest struct {
	values url.Values
}

func (this *WlbNotifyMessageConfirmRequest) GetApiMethodName() string {
	return "taobao.wlb.notify.message.confirm"
}
func (this *WlbNotifyMessageConfirmRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *WlbNotifyMessageConfirmRequest) GetValues() url.Values {
	return this.values
}

/* 物流宝通知消息的id，通过taobao.wlb.notify.message.page.get接口得到的WlbMessage数据结构中的id字段 */
func (this *WlbNotifyMessageConfirmRequest) SetMessageId(value string) {
	this.Set("message_id", value)
}

type WlbNotifyMessageConfirmResponse struct {
	*ErrorResponse                        `json:"error_response,omitempty"`
	WlbNotifyMessageConfirmResponseResult `json:"wlb_notify_message_confirm_response"`
}
type WlbNotifyMessageConfirmResponseResult struct {
	/* 物流宝消息确认时间 */
	GmtModified string `json:"gmt_modified"`
}

/*taobao.wlb.notify.message.page.get*/
type WlbNotifyMessagePageGetRequest struct {
	values url.Values
}

func (this *WlbNotifyMessagePageGetRequest) GetApiMethodName() string {
	return "taobao.wlb.notify.message.page.get"
}
func (this *WlbNotifyMessagePageGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *WlbNotifyMessagePageGetRequest) GetValues() url.Values {
	return this.values
}

/* 查询条件：记录截至时间 */
func (this *WlbNotifyMessagePageGetRequest) SetEndDate(value string) {
	this.Set("end_date", value)
}

/* 通知消息编码：
STOCK_IN_NOT_CONSISTENT---入库单不一致
CANCEL_ORDER_SUCCESS---取消订单成功
INVENTORY_CHECK---盘点
CANCEL_ORDER_FAILURE---取消订单失败
ORDER_REJECT--wms拒单
ORDER_CONFIRMED--订单处理成功 */
func (this *WlbNotifyMessagePageGetRequest) SetMsgCode(value string) {
	this.Set("msg_code", value)
}

/* 分页查询页数 */
func (this *WlbNotifyMessagePageGetRequest) SetPageNo(value string) {
	this.Set("page_no", value)
}

/* 分页查询的每页页数 */
func (this *WlbNotifyMessagePageGetRequest) SetPageSize(value string) {
	this.Set("page_size", value)
}

/* 查询条件：记录开始时间 */
func (this *WlbNotifyMessagePageGetRequest) SetStartDate(value string) {
	this.Set("start_date", value)
}

/* 消息状态：
不需要确认：NO_NEED_CONFIRM
已确认：CONFIRMED
待确认：TO_BE_CONFIRM */
func (this *WlbNotifyMessagePageGetRequest) SetStatus(value string) {
	this.Set("status", value)
}

type WlbNotifyMessagePageGetResponse struct {
	*ErrorResponse                        `json:"error_response,omitempty"`
	WlbNotifyMessagePageGetResponseResult `json:"wlb_notify_message_page_get_response"`
}
type WlbNotifyMessagePageGetResponseResult struct {
	/* 条件查询结果总数量 */
	TotalCount int64 `json:"total_count"`
	/* 消息结果列表 */
	WlbMessages struct {
		WlbMessage []*WlbMessage `json:"wlb_message"`
	} `json:"WlbMessages"`
}

/*taobao.wlb.order.cancel*/
type WlbOrderCancelRequest struct {
	values url.Values
}

func (this *WlbOrderCancelRequest) GetApiMethodName() string {
	return "taobao.wlb.order.cancel"
}
func (this *WlbOrderCancelRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *WlbOrderCancelRequest) GetValues() url.Values {
	return this.values
}

/* 物流宝订单编号 */
func (this *WlbOrderCancelRequest) SetWlbOrderCode(value string) {
	this.Set("wlb_order_code", value)
}

type WlbOrderCancelResponse struct {
	*ErrorResponse               `json:"error_response,omitempty"`
	WlbOrderCancelResponseResult `json:"wlb_order_cancel_response"`
}
type WlbOrderCancelResponseResult struct {
	/* 错误编码列表 */
	ErrorCodeList string `json:"error_code_list"`
	/* 修改时间，只有在取消成功的情况下，才可以做 */
	ModifyTime string `json:"modify_time"`
}

/*taobao.wlb.order.consign*/
type WlbOrderConsignRequest struct {
	values url.Values
}

func (this *WlbOrderConsignRequest) GetApiMethodName() string {
	return "taobao.wlb.order.consign"
}
func (this *WlbOrderConsignRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *WlbOrderConsignRequest) GetValues() url.Values {
	return this.values
}

/* 物流宝订单编号<br /> 支持最大长度为：64<br /> 支持的最大列表长度为：64 */
func (this *WlbOrderConsignRequest) SetWlbOrderCode(value string) {
	this.Set("wlb_order_code", value)
}

type WlbOrderConsignResponse struct {
	*ErrorResponse                `json:"error_response,omitempty"`
	WlbOrderConsignResponseResult `json:"wlb_order_consign_response"`
}
type WlbOrderConsignResponseResult struct {
	/* 修改时间 */
	ModifyTime string `json:"modify_time"`
}

/*taobao.wlb.order.create*/
type WlbOrderCreateRequest struct {
	values url.Values
}

func (this *WlbOrderCreateRequest) GetApiMethodName() string {
	return "taobao.wlb.order.create"
}
func (this *WlbOrderCreateRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *WlbOrderCreateRequest) GetValues() url.Values {
	return this.values
}

/* 支付宝交易号 */
func (this *WlbOrderCreateRequest) SetAlipayNo(value string) {
	this.Set("alipay_no", value)
}

/* 该字段暂时保留 */
func (this *WlbOrderCreateRequest) SetAttributes(value string) {
	this.Set("attributes", value)
}

/* 买家呢称<br /> 支持最大长度为：64<br /> 支持的最大列表长度为：64 */
func (this *WlbOrderCreateRequest) SetBuyerNick(value string) {
	this.Set("buyer_nick", value)
}

/* 期望结束时间，在入库单会使用到 */
func (this *WlbOrderCreateRequest) SetExpectEndTime(value string) {
	this.Set("expect_end_time", value)
}

/* 计划开始送达时间  在入库单中可能会使用 */
func (this *WlbOrderCreateRequest) SetExpectStartTime(value string) {
	this.Set("expect_start_time", value)
}

/* {"invoince_info": [{"bill_type":"发票类型，必选", "bill_title":"发票抬头，必选", "bill_amount":"发票金额(单位是分)，必选","bill_content":"发票内容，可选"}]} */
func (this *WlbOrderCreateRequest) SetInvoinceInfo(value string) {
	this.Set("invoince_info", value)
}

/* 该物流宝订单是否已完成，如果完成则设置为true，如果为false，则需要等待继续创建订单商品信息。 */
func (this *WlbOrderCreateRequest) SetIsFinished(value string) {
	this.Set("is_finished", value)
}

/* 物流宝订单编号，该接口约定每次最多只能传50条order_item_list，如果一个物流宝订单超过50条商品的时候，需要批量来调用该接口，第一次调用的时候，wlb_order_code为空，如果第一次创建成功，该接口返回wlb_order_code，其后继续再该订单上添加商品条目，需要带上wlb_order_code，out_biz_code，order_item_list,is_finished四个字段。 */
func (this *WlbOrderCreateRequest) SetOrderCode(value string) {
	this.Set("order_code", value)
}

/* 用字符串格式来表示订单标记列表：比如COD^PRESELL^SPLIT^LIMIT 等，中间用“^”来隔开 ---------------------------------------- 订单标记list（所有字母全部大写）： 1: COD –货到付款 2: LIMIT-限时配送 3: PRESELL-预售 5:COMPLAIN-已投诉 7:SPLIT-拆单， 8:EXCHANGE-换货， 9:VISIT-上门 ， 10: MODIFYTRANSPORT-是否可改配送方式，
: 是否可改配送方式  默认可更改
11 CONSIGN 物流宝代理发货,自动更改发货状态
12: SELLER_AFFORD 是否卖家承担运费 默认是，即没 13: SYNC_RETURN_BILL，同时退回发票 */
func (this *WlbOrderCreateRequest) SetOrderFlag(value string) {
	this.Set("order_flag", value)
}

/* 订单商品列表： {"order_item_list":[{"trade_code":"可选,淘宝交易订单，并且不是赠品，必须要传订单来源编号"," sub_trade_code ":"可选,淘宝子交易号","item_id":"必须,商品Id","item_code":"必须,商家编码","item_name":"可选,物流宝商品名称","item_quantity":"必选,计划数量","item_price":"必选,物品价格,单位为分","owner_user_nick
":"可选,货主nick 代销模式下会存在","flag":"判断是否为赠品0 不是1是","remarks":"可选,备注","batch_remark":"可选，批次描述信息会把这个信息带给WMS，但不会跟物流宝库存相关联"，"inventory_type":"库存类型1 可销售库存 101 类型用来定义残次品 201 冻结类型库存 301 在途库存","picture_url":"图片Url","distributor_user_nick": "分销商NICK",必选"ext_order_item_code":"可选，外部商品的商家编码"]} ======================================== 如果订单中的商品条目数大于50条的时候，我们会校验，不能创建成功，需要你按照50个一批的数量传，需要分批调用该接口，第二次传的时候，需要带上wlb_order_code和is_finished和order_item_list三个字段是必传的，is_finished为true表示传输完毕，为false表示还没完全传完。 */
func (this *WlbOrderCreateRequest) SetOrderItemList(value string) {
	this.Set("order_item_list", value)
}

/* 订单子类型：
（1）OTHER： 其他
（2）TAOBAO_TRADE： 淘宝交易
（3）OTHER_TRADE：其他交易
（4）ALLCOATE： 调拨
（5）PURCHASE:采购 */
func (this *WlbOrderCreateRequest) SetOrderSubType(value string) {
	this.Set("order_sub_type", value)
}

/* 订单类型:
（1）NORMAL_OUT ：正常出库
（2）NORMAL_IN：正常入库
（3）RETURN_IN：退货入库
（4）EXCHANGE_OUT：换货出库 */
func (this *WlbOrderCreateRequest) SetOrderType(value string) {
	this.Set("order_type", value)
}

/* 外部订单业务ID，该编号在isv中是唯一编号， 用来控制并发，去重用<br /> 支持最大长度为：128<br /> 支持的最大列表长度为：128 */
func (this *WlbOrderCreateRequest) SetOutBizCode(value string) {
	this.Set("out_biz_code", value)
}

/* 包裹件数，入库单和出库单中会用到 */
func (this *WlbOrderCreateRequest) SetPackageCount(value string) {
	this.Set("package_count", value)
}

/* 应收金额，cod订单必选 */
func (this *WlbOrderCreateRequest) SetPayableAmount(value string) {
	this.Set("payable_amount", value)
}

/* 源订单编号 */
func (this *WlbOrderCreateRequest) SetPrevOrderCode(value string) {
	this.Set("prev_order_code", value)
}

/* 收货方信息，必须传， 手机和电话必选其一。
收货方信息：
邮编^^^省^^^市^^^区^^^具体地址^^^收件方名称^^^手机^^^电话

如果某一个字段的数据为空时，必须传NA */
func (this *WlbOrderCreateRequest) SetReceiverInfo(value string) {
	this.Set("receiver_info", value)
}

/* 备注<br /> 支持最大长度为：4000<br /> 支持的最大列表长度为：4000 */
func (this *WlbOrderCreateRequest) SetRemark(value string) {
	this.Set("remark", value)
}

/* 投递时间范围要求,格式'15:20'用分号隔开 */
func (this *WlbOrderCreateRequest) SetScheduleEnd(value string) {
	this.Set("schedule_end", value)
}

/* 投递时间范围要求,格式'13:20'用分号隔开 */
func (this *WlbOrderCreateRequest) SetScheduleStart(value string) {
	this.Set("schedule_start", value)
}

/* 投递时延要求:
（1）INSTANT_ARRIVED： 当日达
（2）TOMMORROY_MORNING_ARRIVED：次晨达
（3）TOMMORROY_ARRIVED：次日达
（4）工作日：WORK_DAY
（5）节假日：WEEKED_DAY */
func (this *WlbOrderCreateRequest) SetScheduleType(value string) {
	this.Set("schedule_type", value)
}

/* 发货方信息，发货方信息必须传， 手机和电话必选其一。 发货方信息：
邮编^^^省^^^市^^^区^^^具体地址^^^收件方名称^^^手机^^^电话
如果某一个字段的数据为空时，必须传NA */
func (this *WlbOrderCreateRequest) SetSenderInfo(value string) {
	this.Set("sender_info", value)
}

/* cod服务费，只有cod订单的时候，才需要这个字段 */
func (this *WlbOrderCreateRequest) SetServiceFee(value string) {
	this.Set("service_fee", value)
}

/* 仓库编码<br /> 支持最大长度为：64<br /> 支持的最大列表长度为：64 */
func (this *WlbOrderCreateRequest) SetStoreCode(value string) {
	this.Set("store_code", value)
}

/* 出库单中可能会用到
运输公司名称^^^运输公司联系人^^^运输公司运单号^^^运输公司电话^^^运输公司联系人身份证号

========================================
如果某一个字段的数据为空时，必须传NA */
func (this *WlbOrderCreateRequest) SetTmsInfo(value string) {
	this.Set("tms_info", value)
}

/* 运单编号，退货单时可能会使用 */
func (this *WlbOrderCreateRequest) SetTmsOrderCode(value string) {
	this.Set("tms_order_code", value)
}

/* 物流公司编码<br /> 支持最大长度为：64<br /> 支持的最大列表长度为：64 */
func (this *WlbOrderCreateRequest) SetTmsServiceCode(value string) {
	this.Set("tms_service_code", value)
}

/* 总金额 */
func (this *WlbOrderCreateRequest) SetTotalAmount(value string) {
	this.Set("total_amount", value)
}

type WlbOrderCreateResponse struct {
	*ErrorResponse               `json:"error_response,omitempty"`
	WlbOrderCreateResponseResult `json:"wlb_order_create_response"`
}
type WlbOrderCreateResponseResult struct {
	/* 订单创建时间 */
	CreateTime string `json:"create_time"`
	/* 物流宝订单创建成功后，返回物流宝的订单编号；如果订单创建失败，该字段为空。 */
	OrderCode string `json:"order_code"`
}

/*taobao.wlb.order.page.get*/
type WlbOrderPageGetRequest struct {
	values url.Values
}

func (this *WlbOrderPageGetRequest) GetApiMethodName() string {
	return "taobao.wlb.order.page.get"
}
func (this *WlbOrderPageGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *WlbOrderPageGetRequest) GetValues() url.Values {
	return this.values
}

/* 查询截止时间 */
func (this *WlbOrderPageGetRequest) SetEndTime(value string) {
	this.Set("end_time", value)
}

/* 物流订单编号 */
func (this *WlbOrderPageGetRequest) SetOrderCode(value string) {
	this.Set("order_code", value)
}

/* 订单状态 */
func (this *WlbOrderPageGetRequest) SetOrderStatus(value string) {
	this.Set("order_status", value)
}

/* 订单子类型：
（1）OTHER： 其他
（2）TAOBAO_TRADE： 淘宝交易
（3）OTHER_TRADE：其他交易
（4）ALLCOATE： 调拨
（5）CHECK:  盘点单
（6）PURCHASE: 采购单 */
func (this *WlbOrderPageGetRequest) SetOrderSubType(value string) {
	this.Set("order_sub_type", value)
}

/* 订单类型:
（1）NORMAL_OUT ：正常出库
（2）NORMAL_IN：正常入库
（3）RETURN_IN：退货入库
（4）EXCHANGE_OUT：换货出库 */
func (this *WlbOrderPageGetRequest) SetOrderType(value string) {
	this.Set("order_type", value)
}

/* 分页的第几页 */
func (this *WlbOrderPageGetRequest) SetPageNo(value string) {
	this.Set("page_no", value)
}

/* 每页多少条 */
func (this *WlbOrderPageGetRequest) SetPageSize(value string) {
	this.Set("page_size", value)
}

/* 查询开始时间 */
func (this *WlbOrderPageGetRequest) SetStartTime(value string) {
	this.Set("start_time", value)
}

type WlbOrderPageGetResponse struct {
	*ErrorResponse                `json:"error_response,omitempty"`
	WlbOrderPageGetResponseResult `json:"wlb_order_page_get_response"`
}
type WlbOrderPageGetResponseResult struct {
	/* 分页查询返回结果 */
	OrderList struct {
		WlbOrder []*WlbOrder `json:"wlb_order"`
	} `json:"OrderList"`
	/* 总条数 */
	TotalCount int64 `json:"total_count"`
}

/*taobao.wlb.orderdetail.date.get*/
type WlbOrderdetailDateGetRequest struct {
	values url.Values
}

func (this *WlbOrderdetailDateGetRequest) GetApiMethodName() string {
	return "taobao.wlb.orderdetail.date.get"
}
func (this *WlbOrderdetailDateGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *WlbOrderdetailDateGetRequest) GetValues() url.Values {
	return this.values
}

/* 查询条件截止日期 */
func (this *WlbOrderdetailDateGetRequest) SetEndTime(value string) {
	this.Set("end_time", value)
}

/* 分页查询参数，指定查询页数，默认为1 */
func (this *WlbOrderdetailDateGetRequest) SetPageNo(value string) {
	this.Set("page_no", value)
}

/* 分页查询参数，每页查询数量，默认20，最大值50,大于50时按照50条查询 */
func (this *WlbOrderdetailDateGetRequest) SetPageSize(value string) {
	this.Set("page_size", value)
}

/* 查询起始日期 */
func (this *WlbOrderdetailDateGetRequest) SetStartTime(value string) {
	this.Set("start_time", value)
}

type WlbOrderdetailDateGetResponse struct {
	*ErrorResponse                      `json:"error_response,omitempty"`
	WlbOrderdetailDateGetResponseResult `json:"wlb_orderdetail_date_get_response"`
}
type WlbOrderdetailDateGetResponseResult struct {
	/* 物流宝详细订单 */
	OrderDetailList struct {
		WlbOrderDetail []*WlbOrderDetail `json:"wlb_order_detail"`
	} `json:"OrderDetailList"`
	/* 总数量 */
	TotalCount int64 `json:"total_count"`
}

/*taobao.wlb.orderitem.page.get*/
type WlbOrderitemPageGetRequest struct {
	values url.Values
}

func (this *WlbOrderitemPageGetRequest) GetApiMethodName() string {
	return "taobao.wlb.orderitem.page.get"
}
func (this *WlbOrderitemPageGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *WlbOrderitemPageGetRequest) GetValues() url.Values {
	return this.values
}

/* 物流宝订单编码 */
func (this *WlbOrderitemPageGetRequest) SetOrderCode(value string) {
	this.Set("order_code", value)
}

/* 分页查询参数，指定查询页数，默认为1 */
func (this *WlbOrderitemPageGetRequest) SetPageNo(value string) {
	this.Set("page_no", value)
}

/* 分页查询参数，每页查询数量，默认20，最大值50,大于50时按照50条查询 */
func (this *WlbOrderitemPageGetRequest) SetPageSize(value string) {
	this.Set("page_size", value)
}

type WlbOrderitemPageGetResponse struct {
	*ErrorResponse                    `json:"error_response,omitempty"`
	WlbOrderitemPageGetResponseResult `json:"wlb_orderitem_page_get_response"`
}
type WlbOrderitemPageGetResponseResult struct {
	/* 物流宝订单商品列表 */
	OrderItemList struct {
		WlbOrderItem []*WlbOrderItem `json:"wlb_order_item"`
	} `json:"OrderItemList"`
	/* 总数量 */
	TotalCount int64 `json:"total_count"`
}

/*taobao.wlb.orderstatus.get*/
type WlbOrderstatusGetRequest struct {
	values url.Values
}

func (this *WlbOrderstatusGetRequest) GetApiMethodName() string {
	return "taobao.wlb.orderstatus.get"
}
func (this *WlbOrderstatusGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *WlbOrderstatusGetRequest) GetValues() url.Values {
	return this.values
}

/* 物流宝订单编码 */
func (this *WlbOrderstatusGetRequest) SetOrderCode(value string) {
	this.Set("order_code", value)
}

type WlbOrderstatusGetResponse struct {
	*ErrorResponse                  `json:"error_response,omitempty"`
	WlbOrderstatusGetResponseResult `json:"wlb_orderstatus_get_response"`
}
type WlbOrderstatusGetResponseResult struct {
	/* 订单流转信息状态列表 */
	WlbOrderStatus struct {
		WlbProcessStatus []*WlbProcessStatus `json:"wlb_process_status"`
	} `json:"WlbOrderStatus"`
}

/*taobao.wlb.out.inventory.change.notify*/
type WlbOutInventoryChangeNotifyRequest struct {
	values url.Values
}

func (this *WlbOutInventoryChangeNotifyRequest) GetApiMethodName() string {
	return "taobao.wlb.out.inventory.change.notify"
}
func (this *WlbOutInventoryChangeNotifyRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *WlbOutInventoryChangeNotifyRequest) GetValues() url.Values {
	return this.values
}

/* 库存变化数量 */
func (this *WlbOutInventoryChangeNotifyRequest) SetChangeCount(value string) {
	this.Set("change_count", value)
}

/* 物流宝商品id或前台宝贝id（由type类型决定） */
func (this *WlbOutInventoryChangeNotifyRequest) SetItemId(value string) {
	this.Set("item_id", value)
}

/* OUT--出库
IN--入库 */
func (this *WlbOutInventoryChangeNotifyRequest) SetOpType(value string) {
	this.Set("op_type", value)
}

/* 订单号，如果source为TAOBAO_TRADE,order_source_code必须为淘宝交易号 */
func (this *WlbOutInventoryChangeNotifyRequest) SetOrderSourceCode(value string) {
	this.Set("order_source_code", value)
}

/* 库存变化唯一标识，用于去重，一个外部唯一编码唯一标识一次库存变化。 */
func (this *WlbOutInventoryChangeNotifyRequest) SetOutBizCode(value string) {
	this.Set("out_biz_code", value)
}

/* 本次库存变化后库存余额 */
func (this *WlbOutInventoryChangeNotifyRequest) SetResultCount(value string) {
	this.Set("result_count", value)
}

/* （1）OTHER： 其他
（2）TAOBAO_TRADE： 淘宝交易
（3）OTHER_TRADE：其他交易
（4）ALLCOATE： 调拨
（5）CHECK:盘点
（6）PURCHASE:采购 */
func (this *WlbOutInventoryChangeNotifyRequest) SetSource(value string) {
	this.Set("source", value)
}

/* 目前非必须，系统会选择默认值 */
func (this *WlbOutInventoryChangeNotifyRequest) SetStoreCode(value string) {
	this.Set("store_code", value)
}

/* WLB_ITEM--物流宝商品
IC_ITEM--淘宝商品
IC_SKU--淘宝sku */
func (this *WlbOutInventoryChangeNotifyRequest) SetType(value string) {
	this.Set("type", value)
}

type WlbOutInventoryChangeNotifyResponse struct {
	*ErrorResponse                            `json:"error_response,omitempty"`
	WlbOutInventoryChangeNotifyResponseResult `json:"wlb_out_inventory_change_notify_response"`
}
type WlbOutInventoryChangeNotifyResponseResult struct {
	/* 库存变化通知成功时间 */
	GmtModified string `json:"gmt_modified"`
}

/*taobao.wlb.replenish.statistics*/
type WlbReplenishStatisticsRequest struct {
	values url.Values
}

func (this *WlbReplenishStatisticsRequest) GetApiMethodName() string {
	return "taobao.wlb.replenish.statistics"
}
func (this *WlbReplenishStatisticsRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *WlbReplenishStatisticsRequest) GetValues() url.Values {
	return this.values
}

/* 商品编码 */
func (this *WlbReplenishStatisticsRequest) SetItemCode(value string) {
	this.Set("item_code", value)
}

/* 商品名称 */
func (this *WlbReplenishStatisticsRequest) SetName(value string) {
	this.Set("name", value)
}

/* 分页参数，默认第一页 */
func (this *WlbReplenishStatisticsRequest) SetPageNo(value string) {
	this.Set("page_no", value)
}

/* 分页每页页数，默认20，最大50<br /> 支持最大值为：50 */
func (this *WlbReplenishStatisticsRequest) SetPageSize(value string) {
	this.Set("page_size", value)
}

/* 仓库编码 */
func (this *WlbReplenishStatisticsRequest) SetStoreCode(value string) {
	this.Set("store_code", value)
}

type WlbReplenishStatisticsResponse struct {
	*ErrorResponse                       `json:"error_response,omitempty"`
	WlbReplenishStatisticsResponseResult `json:"wlb_replenish_statistics_response"`
}
type WlbReplenishStatisticsResponseResult struct {
	/* 补货统计列表 */
	ReplenishList struct {
		WlbReplenish []*WlbReplenish `json:"wlb_replenish"`
	} `json:"ReplenishList"`
	/* 查询记录总数 */
	TotalCount int64 `json:"total_count"`
}

/*taobao.wlb.subscription.query*/
type WlbSubscriptionQueryRequest struct {
	values url.Values
}

func (this *WlbSubscriptionQueryRequest) GetApiMethodName() string {
	return "taobao.wlb.subscription.query"
}
func (this *WlbSubscriptionQueryRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *WlbSubscriptionQueryRequest) GetValues() url.Values {
	return this.values
}

/* 当前页 */
func (this *WlbSubscriptionQueryRequest) SetPageNo(value string) {
	this.Set("page_no", value)
}

/* 分页记录个数，如果用户输入的记录数大于50，则一页显示50条记录 */
func (this *WlbSubscriptionQueryRequest) SetPageSize(value string) {
	this.Set("page_size", value)
}

/* 状态
AUDITING 1-待审核;
CANCEL 2-撤销 ;
CHECKED 3-审核通过 ;
FAILED 4-审核未通过 ;
SYNCHRONIZING 5-同步中;
只允许输入上面指定的值，且可以为空，为空时查询所有状态。若输错了，则按AUDITING处理。 */
func (this *WlbSubscriptionQueryRequest) SetStatus(value string) {
	this.Set("status", value)
}

type WlbSubscriptionQueryResponse struct {
	*ErrorResponse                     `json:"error_response,omitempty"`
	WlbSubscriptionQueryResponseResult `json:"wlb_subscription_query_response"`
}
type WlbSubscriptionQueryResponseResult struct {
	/* 卖家定购的服务列表 */
	SellerSubscriptionList struct {
		WlbSellerSubscription []*WlbSellerSubscription `json:"wlb_seller_subscription"`
	} `json:"SellerSubscriptionList"`
	/* 结果总数 */
	TotalCount int64 `json:"total_count"`
}

/*taobao.wlb.tmsorder.query*/
type WlbTmsorderQueryRequest struct {
	values url.Values
}

func (this *WlbTmsorderQueryRequest) GetApiMethodName() string {
	return "taobao.wlb.tmsorder.query"
}
func (this *WlbTmsorderQueryRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *WlbTmsorderQueryRequest) GetValues() url.Values {
	return this.values
}

/* 物流订单编号<br /> 支持最大长度为：64<br /> 支持的最大列表长度为：64 */
func (this *WlbTmsorderQueryRequest) SetOrderCode(value string) {
	this.Set("order_code", value)
}

/* 当前页 */
func (this *WlbTmsorderQueryRequest) SetPageNo(value string) {
	this.Set("page_no", value)
}

/* 分页记录个数，如果用户输入的记录数大于50，则一页显示50条记录 */
func (this *WlbTmsorderQueryRequest) SetPageSize(value string) {
	this.Set("page_size", value)
}

type WlbTmsorderQueryResponse struct {
	*ErrorResponse                 `json:"error_response,omitempty"`
	WlbTmsorderQueryResponseResult `json:"wlb_tmsorder_query_response"`
}
type WlbTmsorderQueryResponseResult struct {
	/* 物流订单运单信息列表 */
	TmsOrderList struct {
		WlbTmsOrder []*WlbTmsOrder `json:"wlb_tms_order"`
	} `json:"TmsOrderList"`
	/* 结果总数 */
	TotalCount int64 `json:"total_count"`
}

/*taobao.wlb.tradeorder.get*/
type WlbTradeorderGetRequest struct {
	values url.Values
}

func (this *WlbTradeorderGetRequest) GetApiMethodName() string {
	return "taobao.wlb.tradeorder.get"
}
func (this *WlbTradeorderGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *WlbTradeorderGetRequest) GetValues() url.Values {
	return this.values
}

/* 子交易号 */
func (this *WlbTradeorderGetRequest) SetSubTradeId(value string) {
	this.Set("sub_trade_id", value)
}

/* 指定交易类型的交易号 */
func (this *WlbTradeorderGetRequest) SetTradeId(value string) {
	this.Set("trade_id", value)
}

/* 交易类型:
TAOBAO--淘宝交易
OTHER_TRADE--其它交易 */
func (this *WlbTradeorderGetRequest) SetTradeType(value string) {
	this.Set("trade_type", value)
}

type WlbTradeorderGetResponse struct {
	*ErrorResponse                 `json:"error_response,omitempty"`
	WlbTradeorderGetResponseResult `json:"wlb_tradeorder_get_response"`
}
type WlbTradeorderGetResponseResult struct {
	/* 物流宝订单列表信息 */
	WlbOrderList struct {
		WlbOrder []*WlbOrder `json:"wlb_order"`
	} `json:"WlbOrderList"`
}

/*taobao.wlb.waybill.get*/
type WlbWaybillGetRequest struct {
	values url.Values
}

func (this *WlbWaybillGetRequest) GetApiMethodName() string {
	return "taobao.wlb.waybill.get"
}
func (this *WlbWaybillGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *WlbWaybillGetRequest) GetValues() url.Values {
	return this.values
}

/* 顺丰(SF)、EMS(标准快递：EMS；经济快件：EYB)、宅急送(ZJS)、圆通(YTO)、中通(ZTO)、百世汇通(HTKY)、优速(UC)、申通(STO)、天天快递 (TTKDEX)、全峰 (QFKD)、快捷(FAST) */
func (this *WlbWaybillGetRequest) SetCpCode(value string) {
	this.Set("cp_code", value)
}

/* 发货 地址 */
func (this *WlbWaybillGetRequest) SetShippingAddress(value string) {
	this.Set("shipping_address", value)
}

/* order_channels_type 订单来源：订单产生的电商平台
淘宝(TB)、天猫(TM)、京东(JD)、当当(DD)、拍拍(PP)、易讯(YX)、ebay(EBAY)、QQ网购(QQ)、亚马逊(AMAZON)、苏宁(SN)、国美(GM)、唯品会(WPH)、聚美(JM)、乐蜂(LF)、蘑菇街(MGJ)、聚尚(JS)、拍鞋(PX)、银泰(YT)、1号店(YHD)、凡客(VANCL)、邮乐(YL)、优购(YG)、阿里巴巴(1688)、其他(OTHERS) */
func (this *WlbWaybillGetRequest) SetTradeOrderInfoCols(value string) {
	this.Set("trade_order_info_cols", value)
}

type WlbWaybillGetResponse struct {
	*ErrorResponse              `json:"error_response,omitempty"`
	WlbWaybillGetResponseResult `json:"wlb_waybill_get_response"`
}
type WlbWaybillGetResponseResult struct {
	/* 结果 */
	Results struct {
		WaybillApplyNewInfo []*WaybillApplyNewInfo `json:"waybill_apply_new_info"`
	} `json:"Results"`
}

/*taobao.wlb.waybill.i.cancel*/
type WlbWaybillICancelRequest struct {
	values url.Values
}

func (this *WlbWaybillICancelRequest) GetApiMethodName() string {
	return "taobao.wlb.waybill.i.cancel"
}
func (this *WlbWaybillICancelRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *WlbWaybillICancelRequest) GetValues() url.Values {
	return this.values
}

/* 取消接口入参 */
func (this *WlbWaybillICancelRequest) SetWaybillApplyCancelRequest(value string) {
	this.Set("waybill_apply_cancel_request", value)
}

type WlbWaybillICancelResponse struct {
	*ErrorResponse                  `json:"error_response,omitempty"`
	WlbWaybillICancelResponseResult `json:"wlb_waybill_i_cancel_response"`
}
type WlbWaybillICancelResponseResult struct {
	/* 调用取消是否成功 */
	CancelResult bool `json:"cancel_result"`
}

/*taobao.wlb.waybill.i.fullupdate*/
type WlbWaybillIFullupdateRequest struct {
	values url.Values
}

func (this *WlbWaybillIFullupdateRequest) GetApiMethodName() string {
	return "taobao.wlb.waybill.i.fullupdate"
}
func (this *WlbWaybillIFullupdateRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *WlbWaybillIFullupdateRequest) GetValues() url.Values {
	return this.values
}

/* 更新面单信息请求 */
func (this *WlbWaybillIFullupdateRequest) SetWaybillApplyFullUpdateRequest(value string) {
	this.Set("waybill_apply_full_update_request", value)
}

type WlbWaybillIFullupdateResponse struct {
	*ErrorResponse                      `json:"error_response,omitempty"`
	WlbWaybillIFullupdateResponseResult `json:"wlb_waybill_i_fullupdate_response"`
}
type WlbWaybillIFullupdateResponseResult struct {
	/* 更新接口出参 */
	WaybillApplyUpdateInfo *WaybillApplyUpdateInfo `json:"waybill_apply_update_info"`
}

/*taobao.wlb.waybill.i.get*/
type WlbWaybillIGetRequest struct {
	values url.Values
}

func (this *WlbWaybillIGetRequest) GetApiMethodName() string {
	return "taobao.wlb.waybill.i.get"
}
func (this *WlbWaybillIGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *WlbWaybillIGetRequest) GetValues() url.Values {
	return this.values
}

/* 面单申请 */
func (this *WlbWaybillIGetRequest) SetWaybillApplyNewRequest(value string) {
	this.Set("waybill_apply_new_request", value)
}

type WlbWaybillIGetResponse struct {
	*ErrorResponse               `json:"error_response,omitempty"`
	WlbWaybillIGetResponseResult `json:"wlb_waybill_i_get_response"`
}
type WlbWaybillIGetResponseResult struct {
	/* 面单申请接口返回信息 */
	WaybillApplyNewCols struct {
		WaybillApplyNewInfo []*WaybillApplyNewInfo `json:"waybill_apply_new_info"`
	} `json:"WaybillApplyNewCols"`
}

/*taobao.wlb.waybill.i.print*/
type WlbWaybillIPrintRequest struct {
	values url.Values
}

func (this *WlbWaybillIPrintRequest) GetApiMethodName() string {
	return "taobao.wlb.waybill.i.print"
}
func (this *WlbWaybillIPrintRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *WlbWaybillIPrintRequest) GetValues() url.Values {
	return this.values
}

/* 打印请求 */
func (this *WlbWaybillIPrintRequest) SetWaybillApplyPrintCheckRequest(value string) {
	this.Set("waybill_apply_print_check_request", value)
}

type WlbWaybillIPrintResponse struct {
	*ErrorResponse                 `json:"error_response,omitempty"`
	WlbWaybillIPrintResponseResult `json:"wlb_waybill_i_print_response"`
}
type WlbWaybillIPrintResponseResult struct {
	/* 错误的面单号 */
	ErrorWaybillCode string `json:"error_waybill_code"`
	/* 面单打印信息 */
	WaybillApplyPrintCheckInfos struct {
		WaybillApplyPrintCheckInfo []*WaybillApplyPrintCheckInfo `json:"waybill_apply_print_check_info"`
	} `json:"WaybillApplyPrintCheckInfos"`
}

/*taobao.wlb.waybill.i.product*/
type WlbWaybillIProductRequest struct {
	values url.Values
}

func (this *WlbWaybillIProductRequest) GetApiMethodName() string {
	return "taobao.wlb.waybill.i.product"
}
func (this *WlbWaybillIProductRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *WlbWaybillIProductRequest) GetValues() url.Values {
	return this.values
}

/* 查询物流商电子面单产品类型入参 */
func (this *WlbWaybillIProductRequest) SetWaybillProductTypeRequest(value string) {
	this.Set("waybill_product_type_request", value)
}

type WlbWaybillIProductResponse struct {
	*ErrorResponse                   `json:"error_response,omitempty"`
	WlbWaybillIProductResponseResult `json:"wlb_waybill_i_product_response"`
}
type WlbWaybillIProductResponseResult struct {
	/* 产品类型返回 */
	ProductTypes struct {
		WaybillProductType []*WaybillProductType `json:"waybill_product_type"`
	} `json:"ProductTypes"`
}

/*taobao.wlb.waybill.i.querydetail*/
type WlbWaybillIQuerydetailRequest struct {
	values url.Values
}

func (this *WlbWaybillIQuerydetailRequest) GetApiMethodName() string {
	return "taobao.wlb.waybill.i.querydetail"
}
func (this *WlbWaybillIQuerydetailRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *WlbWaybillIQuerydetailRequest) GetValues() url.Values {
	return this.values
}

/* 面单查询请求 */
func (this *WlbWaybillIQuerydetailRequest) SetWaybillDetailQueryRequest(value string) {
	this.Set("waybill_detail_query_request", value)
}

type WlbWaybillIQuerydetailResponse struct {
	*ErrorResponse                       `json:"error_response,omitempty"`
	WlbWaybillIQuerydetailResponseResult `json:"wlb_waybill_i_querydetail_response"`
}
type WlbWaybillIQuerydetailResponseResult struct {
	/* 面单查询错误编码 */
	ErrorCodes []string `json:"error_codes"`
	/* 不存在的面单号 */
	InexistentWaybillCodes []string `json:"inexistent_waybill_codes"`
	/* 查询是否成功 */
	QuerySuccess bool `json:"query_success"`
	/* 面单详情 */
	WaybillDetails struct {
		WaybillDetailQueryInfo []*WaybillDetailQueryInfo `json:"waybill_detail_query_info"`
	} `json:"WaybillDetails"`
}

/*taobao.wlb.waybill.i.search*/
type WlbWaybillISearchRequest struct {
	values url.Values
}

func (this *WlbWaybillISearchRequest) GetApiMethodName() string {
	return "taobao.wlb.waybill.i.search"
}
func (this *WlbWaybillISearchRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *WlbWaybillISearchRequest) GetValues() url.Values {
	return this.values
}

/* 查询网点信息 */
func (this *WlbWaybillISearchRequest) SetWaybillApplyRequest(value string) {
	this.Set("waybill_apply_request", value)
}

type WlbWaybillISearchResponse struct {
	*ErrorResponse                  `json:"error_response,omitempty"`
	WlbWaybillISearchResponseResult `json:"wlb_waybill_i_search_response"`
}
type WlbWaybillISearchResponseResult struct {
	/* 订购关系 */
	Subscribtions struct {
		WaybillApplySubscriptionInfo []*WaybillApplySubscriptionInfo `json:"waybill_apply_subscription_info"`
	} `json:"Subscribtions"`
}

/*taobao.wlb.waybill.search*/
type WlbWaybillSearchRequest struct {
	values url.Values
}

func (this *WlbWaybillSearchRequest) GetApiMethodName() string {
	return "taobao.wlb.waybill.search"
}
func (this *WlbWaybillSearchRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *WlbWaybillSearchRequest) GetValues() url.Values {
	return this.values
}

/* 顺丰(SF)、EMS(标准快递：EMS；经济快件：EYB)、宅急送(ZJS)、圆通(YTO)、中通(ZTO)、百世汇通(HTKY)、优速(UC)、申通(STO)、天天快递 (TTKDEX)、全峰 (QFKD)、快捷(FAST)、邮政小包(POSTB) */
func (this *WlbWaybillSearchRequest) SetCpCode(value string) {
	this.Set("cp_code", value)
}

type WlbWaybillSearchResponse struct {
	*ErrorResponse                 `json:"error_response,omitempty"`
	WlbWaybillSearchResponseResult `json:"wlb_waybill_search_response"`
}
type WlbWaybillSearchResponseResult struct {
	/* 查询相关数据(cp_type 1是直营，2是加盟
	) */
	Results struct {
		WaybillApplySubscriptionInfo []*WaybillApplySubscriptionInfo `json:"waybill_apply_subscription_info"`
	} `json:"Results"`
}

/*taobao.wlb.waybill.update*/
type WlbWaybillUpdateRequest struct {
	values url.Values
}

func (this *WlbWaybillUpdateRequest) GetApiMethodName() string {
	return "taobao.wlb.waybill.update"
}
func (this *WlbWaybillUpdateRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *WlbWaybillUpdateRequest) GetValues() url.Values {
	return this.values
}

/* 顺丰(SF)、EMS(标准快递：EMS；经济快件：EYB)、宅急送(ZJS)、圆通(YTO)、中通(ZTO)、百世汇通(HTKY)、优速(UC)、申通(STO)、天天快递 (TTKDEX)、全峰 (QFKD)、快捷(FAST)、邮政小包(POSTB) */
func (this *WlbWaybillUpdateRequest) SetCpCode(value string) {
	this.Set("cp_code", value)
}

/* 订单数据 */
func (this *WlbWaybillUpdateRequest) SetTradeOrderInfo(value string) {
	this.Set("trade_order_info", value)
}

type WlbWaybillUpdateResponse struct {
	*ErrorResponse                 `json:"error_response,omitempty"`
	WlbWaybillUpdateResponseResult `json:"wlb_waybill_update_response"`
}
type WlbWaybillUpdateResponseResult struct {
	/* 更新返回 */
	Results struct {
		WaybillApplyUpdateInfo []*WaybillApplyUpdateInfo `json:"waybill_apply_update_info"`
	} `json:"Results"`
}

/*taobao.wlb.wlborder.get*/
type WlbWlborderGetRequest struct {
	values url.Values
}

func (this *WlbWlborderGetRequest) GetApiMethodName() string {
	return "taobao.wlb.wlborder.get"
}
func (this *WlbWlborderGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *WlbWlborderGetRequest) GetValues() url.Values {
	return this.values
}

/* 物流宝订单编码 */
func (this *WlbWlborderGetRequest) SetWlbOrderCode(value string) {
	this.Set("wlb_order_code", value)
}

type WlbWlborderGetResponse struct {
	*ErrorResponse               `json:"error_response,omitempty"`
	WlbWlborderGetResponseResult `json:"wlb_wlborder_get_response"`
}
type WlbWlborderGetResponseResult struct {
	/* 物流宝订单详情 */
	WlbOrder *WlbOrder `json:"wlb_order"`
}

/*taobao.marketing.promotion.kfc*/
type MarketingPromotionKfcRequest struct {
	values url.Values
}

func (this *MarketingPromotionKfcRequest) GetApiMethodName() string {
	return "taobao.marketing.promotion.kfc"
}
func (this *MarketingPromotionKfcRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *MarketingPromotionKfcRequest) GetValues() url.Values {
	return this.values
}

/* 活动描述 */
func (this *MarketingPromotionKfcRequest) SetPromotionDesc(value string) {
	this.Set("promotion_desc", value)
}

/* 活动名称 */
func (this *MarketingPromotionKfcRequest) SetPromotionTitle(value string) {
	this.Set("promotion_title", value)
}

type MarketingPromotionKfcResponse struct {
	*ErrorResponse                      `json:"error_response,omitempty"`
	MarketingPromotionKfcResponseResult `json:"marketing_promotion_kfc_response"`
}
type MarketingPromotionKfcResponseResult struct {
	/* 是否成功 */
	IsSuccess bool `json:"is_success"`
}

/*taobao.marketing.promotions.get*/
type MarketingPromotionsGetRequest struct {
	values url.Values
}

func (this *MarketingPromotionsGetRequest) GetApiMethodName() string {
	return "taobao.marketing.promotions.get"
}
func (this *MarketingPromotionsGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *MarketingPromotionsGetRequest) GetValues() url.Values {
	return this.values
}

/* 需返回的优惠策略结构字段列表。可选值为Promotion中所有字段，如：promotion_id, promotion_title, item_id, status, tag_id等等 */
func (this *MarketingPromotionsGetRequest) SetFields(value string) {
	this.Set("fields", value)
}

/* 是否新标签标识 */
func (this *MarketingPromotionsGetRequest) SetIsNewTag(value string) {
	this.Set("is_new_tag", value)
}

/* 商品数字ID。根据该ID查询商品下通过第三方工具设置的所有优惠策略 */
func (this *MarketingPromotionsGetRequest) SetNumIid(value string) {
	this.Set("num_iid", value)
}

/* 优惠策略状态。可选值：ACTIVE(有效)，UNACTIVE(无效)，若不传或者传入其他值，则默认查询全部 */
func (this *MarketingPromotionsGetRequest) SetStatus(value string) {
	this.Set("status", value)
}

/* 标签ID */
func (this *MarketingPromotionsGetRequest) SetTagId(value string) {
	this.Set("tag_id", value)
}

type MarketingPromotionsGetResponse struct {
	*ErrorResponse                       `json:"error_response,omitempty"`
	MarketingPromotionsGetResponseResult `json:"marketing_promotions_get_response"`
}
type MarketingPromotionsGetResponseResult struct {
	/* 商品对应的所有优惠列表 */
	Promotions struct {
		Promotion []*Promotion `json:"promotion"`
	} `json:"Promotions"`
	/* 结果总数 */
	TotalResults int64 `json:"total_results"`
}

/*taobao.promotion.activity.get*/
type PromotionActivityGetRequest struct {
	values url.Values
}

func (this *PromotionActivityGetRequest) GetApiMethodName() string {
	return "taobao.promotion.activity.get"
}
func (this *PromotionActivityGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *PromotionActivityGetRequest) GetValues() url.Values {
	return this.values
}

/* 活动的id */
func (this *PromotionActivityGetRequest) SetActivityId(value string) {
	this.Set("activity_id", value)
}

type PromotionActivityGetResponse struct {
	*ErrorResponse                     `json:"error_response,omitempty"`
	PromotionActivityGetResponseResult `json:"promotion_activity_get_response"`
}
type PromotionActivityGetResponseResult struct {
	/* 活动列表 */
	Activitys struct {
		Activity []*Activity `json:"activity"`
	} `json:"Activitys"`
}

/*taobao.promotion.coupons.get*/
type PromotionCouponsGetRequest struct {
	values url.Values
}

func (this *PromotionCouponsGetRequest) GetApiMethodName() string {
	return "taobao.promotion.coupons.get"
}
func (this *PromotionCouponsGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *PromotionCouponsGetRequest) GetValues() url.Values {
	return this.values
}

/* 优惠券的id，唯一标识这个优惠券 */
func (this *PromotionCouponsGetRequest) SetCouponId(value string) {
	this.Set("coupon_id", value)
}

/* 优惠券的面额，必须是3，5，10，20，50,100<br /> 支持最大值为：100<br /> 支持最小值为：3 */
func (this *PromotionCouponsGetRequest) SetDenominations(value string) {
	this.Set("denominations", value)
}

/* 优惠券的截止日期 */
func (this *PromotionCouponsGetRequest) SetEndTime(value string) {
	this.Set("end_time", value)
}

/* 查询的页号，结果集是分页返回的，每页20条<br /> 支持最小值为：1 */
func (this *PromotionCouponsGetRequest) SetPageNo(value string) {
	this.Set("page_no", value)
}

/* 每页条数<br /> 支持的最大列表长度为：20 */
func (this *PromotionCouponsGetRequest) SetPageSize(value string) {
	this.Set("page_size", value)
}

type PromotionCouponsGetResponse struct {
	*ErrorResponse                    `json:"error_response,omitempty"`
	PromotionCouponsGetResponseResult `json:"promotion_coupons_get_response"`
}
type PromotionCouponsGetResponseResult struct {
	/* 优惠券列表 */
	Coupons struct {
		Coupon []*Coupon `json:"coupon"`
	} `json:"Coupons"`
	/* 查询的总数量 */
	TotalResults int64 `json:"total_results"`
}

/*taobao.promotion.limitdiscount.detail.get*/
type PromotionLimitdiscountDetailGetRequest struct {
	values url.Values
}

func (this *PromotionLimitdiscountDetailGetRequest) GetApiMethodName() string {
	return "taobao.promotion.limitdiscount.detail.get"
}
func (this *PromotionLimitdiscountDetailGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *PromotionLimitdiscountDetailGetRequest) GetValues() url.Values {
	return this.values
}

/* 限时打折ID。这个针对查询唯一限时打折情况。 */
func (this *PromotionLimitdiscountDetailGetRequest) SetLimitDiscountId(value string) {
	this.Set("limit_discount_id", value)
}

type PromotionLimitdiscountDetailGetResponse struct {
	*ErrorResponse                                `json:"error_response,omitempty"`
	PromotionLimitdiscountDetailGetResponseResult `json:"promotion_limitdiscount_detail_get_response"`
}
type PromotionLimitdiscountDetailGetResponseResult struct {
	/* 限时打折对应的商品详情列表。 */
	ItemDiscountDetailList struct {
		LimitDiscountDetail []*LimitDiscountDetail `json:"limit_discount_detail"`
	} `json:"ItemDiscountDetailList"`
}

/*taobao.promotion.limitdiscount.get*/
type PromotionLimitdiscountGetRequest struct {
	values url.Values
}

func (this *PromotionLimitdiscountGetRequest) GetApiMethodName() string {
	return "taobao.promotion.limitdiscount.get"
}
func (this *PromotionLimitdiscountGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *PromotionLimitdiscountGetRequest) GetValues() url.Values {
	return this.values
}

/* 限时打折结束时间。输入的时间会被截取，年月日有效，时分秒忽略。 */
func (this *PromotionLimitdiscountGetRequest) SetEndTime(value string) {
	this.Set("end_time", value)
}

/* 限时打折ID。这个针对查询唯一限时打折情况。若此字段不为空，则说明操作为单条限时打折记录查询，其他字段忽略。若想分页按条件查询，这个字段置为空。 */
func (this *PromotionLimitdiscountGetRequest) SetLimitDiscountId(value string) {
	this.Set("limit_discount_id", value)
}

/* 分页页号。默认1。当页数大于最大页数时，结果为最大页数的数据。 */
func (this *PromotionLimitdiscountGetRequest) SetPageNumber(value string) {
	this.Set("page_number", value)
}

/* 限时打折开始时间。输入的时间会被截取，年月日有效，时分秒忽略。 */
func (this *PromotionLimitdiscountGetRequest) SetStartTime(value string) {
	this.Set("start_time", value)
}

/* 限时打折活动状态。ALL:全部状态;OVER:已结束;DOING:进行中;PROPARE:未开始(只支持大写)。当limit_discount_id为空时，为空时，默认为全部的状态。 */
func (this *PromotionLimitdiscountGetRequest) SetStatus(value string) {
	this.Set("status", value)
}

type PromotionLimitdiscountGetResponse struct {
	*ErrorResponse                          `json:"error_response,omitempty"`
	PromotionLimitdiscountGetResponseResult `json:"promotion_limitdiscount_get_response"`
}
type PromotionLimitdiscountGetResponseResult struct {
	/* 限时打折列表。 */
	LimitDiscountList struct {
		LimitDiscount []*LimitDiscount `json:"limit_discount"`
	} `json:"LimitDiscountList"`
	/* 满足该查询条件的限时打折总数量。 */
	TotalCount int64 `json:"total_count"`
}

/*taobao.promotion.meal.get*/
type PromotionMealGetRequest struct {
	values url.Values
}

func (this *PromotionMealGetRequest) GetApiMethodName() string {
	return "taobao.promotion.meal.get"
}
func (this *PromotionMealGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *PromotionMealGetRequest) GetValues() url.Values {
	return this.values
}

/* 搭配套餐id */
func (this *PromotionMealGetRequest) SetMealId(value string) {
	this.Set("meal_id", value)
}

/* 套餐状态。有效：VALID;失效：INVALID(有效套餐为可使用的套餐,无效套餐为套餐中有商品下架或库存为0时)。默认时两种情况都会查询。 */
func (this *PromotionMealGetRequest) SetStatus(value string) {
	this.Set("status", value)
}

type PromotionMealGetResponse struct {
	*ErrorResponse                 `json:"error_response,omitempty"`
	PromotionMealGetResponseResult `json:"promotion_meal_get_response"`
}
type PromotionMealGetResponseResult struct {
	/* 搭配套餐列表。 */
	MealList struct {
		Meal []*Meal `json:"meal"`
	} `json:"MealList"`
}

/*taobao.promotion.mjs.get*/
type PromotionMjsGetRequest struct {
	values url.Values
}

func (this *PromotionMjsGetRequest) GetApiMethodName() string {
	return "taobao.promotion.mjs.get"
}
func (this *PromotionMjsGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *PromotionMjsGetRequest) GetValues() url.Values {
	return this.values
}

type PromotionMjsGetResponse struct {
	*ErrorResponse                `json:"error_response,omitempty"`
	PromotionMjsGetResponseResult `json:"promotion_mjs_get_response"`
}
type PromotionMjsGetResponseResult struct {
	/* 活动结束时间。 */
	EndTime string `json:"end_time"`
	/* 活动备注。可为空，100汉字以内。 */
	Memo string `json:"memo"`
	/* 满就送活动优惠列表。 */
	MjsList struct {
		PromotionMjs []*PromotionMjs `json:"promotion_mjs"`
	} `json:"MjsList"`
	/* 满就送名称。 */
	MjsName string `json:"mjs_name"`
	/* 上不封顶标识！ */
	NoCap bool `json:"no_cap"`
	/* 淘宝店铺或外店的名字。 */
	ShopName string `json:"shop_name"`
	/* 活动开始时间。 */
	StartTime string `json:"start_time"`
}

/*taobao.crm.grade.get*/
type CrmGradeGetRequest struct {
	values url.Values
}

func (this *CrmGradeGetRequest) GetApiMethodName() string {
	return "taobao.crm.grade.get"
}
func (this *CrmGradeGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *CrmGradeGetRequest) GetValues() url.Values {
	return this.values
}

type CrmGradeGetResponse struct {
	*ErrorResponse            `json:"error_response,omitempty"`
	CrmGradeGetResponseResult `json:"crm_grade_get_response"`
}
type CrmGradeGetResponseResult struct {
	/* 等级信息集合 */
	GradePromotions struct {
		GradePromotion []*GradePromotion `json:"grade_promotion"`
	} `json:"GradePromotions"`
}

/*taobao.crm.group.add*/
type CrmGroupAddRequest struct {
	values url.Values
}

func (this *CrmGroupAddRequest) GetApiMethodName() string {
	return "taobao.crm.group.add"
}
func (this *CrmGroupAddRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *CrmGroupAddRequest) GetValues() url.Values {
	return this.values
}

/* 分组名称，每个卖家最多可以拥有100个分组<br /> 支持最大长度为：15<br /> 支持的最大列表长度为：15 */
func (this *CrmGroupAddRequest) SetGroupName(value string) {
	this.Set("group_name", value)
}

type CrmGroupAddResponse struct {
	*ErrorResponse            `json:"error_response,omitempty"`
	CrmGroupAddResponseResult `json:"crm_group_add_response"`
}
type CrmGroupAddResponseResult struct {
	/* 新增分组的id */
	GroupId int64 `json:"group_id"`
	/* 添加分组是否成功 */
	IsSuccess bool `json:"is_success"`
}

/*taobao.crm.group.append*/
type CrmGroupAppendRequest struct {
	values url.Values
}

func (this *CrmGroupAppendRequest) GetApiMethodName() string {
	return "taobao.crm.group.append"
}
func (this *CrmGroupAppendRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *CrmGroupAppendRequest) GetValues() url.Values {
	return this.values
}

/* 添加的来源分组<br /> 支持最小值为：1<br /> 支持的最大列表长度为：19 */
func (this *CrmGroupAppendRequest) SetFromGroupId(value string) {
	this.Set("from_group_id", value)
}

/* 添加的目标分组<br /> 支持最小值为：1<br /> 支持的最大列表长度为：19 */
func (this *CrmGroupAppendRequest) SetToGroupId(value string) {
	this.Set("to_group_id", value)
}

type CrmGroupAppendResponse struct {
	*ErrorResponse               `json:"error_response,omitempty"`
	CrmGroupAppendResponseResult `json:"crm_group_append_response"`
}
type CrmGroupAppendResponseResult struct {
	/* 异步任务请求成功，添加任务是否完成通过taobao.crm.grouptask.check检测 */
	IsSuccess bool `json:"is_success"`
}

/*taobao.crm.group.delete*/
type CrmGroupDeleteRequest struct {
	values url.Values
}

func (this *CrmGroupDeleteRequest) GetApiMethodName() string {
	return "taobao.crm.group.delete"
}
func (this *CrmGroupDeleteRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *CrmGroupDeleteRequest) GetValues() url.Values {
	return this.values
}

/* 要删除的分组id<br /> 支持最小值为：1<br /> 支持的最大列表长度为：19 */
func (this *CrmGroupDeleteRequest) SetGroupId(value string) {
	this.Set("group_id", value)
}

type CrmGroupDeleteResponse struct {
	*ErrorResponse               `json:"error_response,omitempty"`
	CrmGroupDeleteResponseResult `json:"crm_group_delete_response"`
}
type CrmGroupDeleteResponseResult struct {
	/* 异步任务请求成功，是否执行完毕需要通过taobao.crm.grouptask.check检测 */
	IsSuccess bool `json:"is_success"`
}

/*taobao.crm.group.move*/
type CrmGroupMoveRequest struct {
	values url.Values
}

func (this *CrmGroupMoveRequest) GetApiMethodName() string {
	return "taobao.crm.group.move"
}
func (this *CrmGroupMoveRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *CrmGroupMoveRequest) GetValues() url.Values {
	return this.values
}

/* 需要移动的分组<br /> 支持最小值为：1<br /> 支持的最大列表长度为：19 */
func (this *CrmGroupMoveRequest) SetFromGroupId(value string) {
	this.Set("from_group_id", value)
}

/* 目的分组<br /> 支持最小值为：1<br /> 支持的最大列表长度为：19 */
func (this *CrmGroupMoveRequest) SetToGroupId(value string) {
	this.Set("to_group_id", value)
}

type CrmGroupMoveResponse struct {
	*ErrorResponse             `json:"error_response,omitempty"`
	CrmGroupMoveResponseResult `json:"crm_group_move_response"`
}
type CrmGroupMoveResponseResult struct {
	/* 异步任务请求成功，是否执行完毕需要通过taobao.crm.grouptask.check检测 */
	IsSuccess bool `json:"is_success"`
}

/*taobao.crm.group.update*/
type CrmGroupUpdateRequest struct {
	values url.Values
}

func (this *CrmGroupUpdateRequest) GetApiMethodName() string {
	return "taobao.crm.group.update"
}
func (this *CrmGroupUpdateRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *CrmGroupUpdateRequest) GetValues() url.Values {
	return this.values
}

/* 分组的id<br /> 支持最小值为：1<br /> 支持的最大列表长度为：19 */
func (this *CrmGroupUpdateRequest) SetGroupId(value string) {
	this.Set("group_id", value)
}

/* 新的分组名，分组名称不能包含|或者：<br /> 支持最大长度为：15<br /> 支持的最大列表长度为：15 */
func (this *CrmGroupUpdateRequest) SetNewGroupName(value string) {
	this.Set("new_group_name", value)
}

type CrmGroupUpdateResponse struct {
	*ErrorResponse               `json:"error_response,omitempty"`
	CrmGroupUpdateResponseResult `json:"crm_group_update_response"`
}
type CrmGroupUpdateResponseResult struct {
	/* 分组修改是否成功 */
	IsSuccess bool `json:"is_success"`
}

/*taobao.crm.groups.get*/
type CrmGroupsGetRequest struct {
	values url.Values
}

func (this *CrmGroupsGetRequest) GetApiMethodName() string {
	return "taobao.crm.groups.get"
}
func (this *CrmGroupsGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *CrmGroupsGetRequest) GetValues() url.Values {
	return this.values
}

/* 显示第几页的分组，如果输入的页码大于总共的页码数，例如总共10页，但是current_page的值为11，则返回空白页，最小页码为1<br /> 支持最大值为：1000000<br /> 支持最小值为：1<br /> 支持的最大列表长度为：3 */
func (this *CrmGroupsGetRequest) SetCurrentPage(value string) {
	this.Set("current_page", value)
}

/* 每页显示的记录数，其最大值不能超过100条，最小值为1，默认20条<br /> 支持最大值为：100<br /> 支持最小值为：1<br /> 支持的最大列表长度为：3 */
func (this *CrmGroupsGetRequest) SetPageSize(value string) {
	this.Set("page_size", value)
}

type CrmGroupsGetResponse struct {
	*ErrorResponse             `json:"error_response,omitempty"`
	CrmGroupsGetResponseResult `json:"crm_groups_get_response"`
}
type CrmGroupsGetResponseResult struct {
	/* 查询到的当前卖家的当前页的会员 */
	Groups struct {
		Group []*Group `json:"group"`
	} `json:"Groups"`
	/* 记录总数 */
	TotalResult int64 `json:"total_result"`
}

/*taobao.crm.grouptask.check*/
type CrmGrouptaskCheckRequest struct {
	values url.Values
}

func (this *CrmGrouptaskCheckRequest) GetApiMethodName() string {
	return "taobao.crm.grouptask.check"
}
func (this *CrmGrouptaskCheckRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *CrmGrouptaskCheckRequest) GetValues() url.Values {
	return this.values
}

/* 分组id */
func (this *CrmGrouptaskCheckRequest) SetGroupId(value string) {
	this.Set("group_id", value)
}

type CrmGrouptaskCheckResponse struct {
	*ErrorResponse                  `json:"error_response,omitempty"`
	CrmGrouptaskCheckResponseResult `json:"crm_grouptask_check_response"`
}
type CrmGrouptaskCheckResponseResult struct {
	/* 异步任务是否完成，true表示完成 */
	IsFinished bool `json:"is_finished"`
}

/*taobao.crm.member.group.get*/
type CrmMemberGroupGetRequest struct {
	values url.Values
}

func (this *CrmMemberGroupGetRequest) GetApiMethodName() string {
	return "taobao.crm.member.group.get"
}
func (this *CrmMemberGroupGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *CrmMemberGroupGetRequest) GetValues() url.Values {
	return this.values
}

/* 会员Nick */
func (this *CrmMemberGroupGetRequest) SetBuyerNick(value string) {
	this.Set("buyer_nick", value)
}

type CrmMemberGroupGetResponse struct {
	*ErrorResponse                  `json:"error_response,omitempty"`
	CrmMemberGroupGetResponseResult `json:"crm_member_group_get_response"`
}
type CrmMemberGroupGetResponseResult struct {
	/* 查询到的当前卖家的当前页的会员 */
	Groups struct {
		Group []*Group `json:"group"`
	} `json:"Groups"`
}

/*taobao.crm.memberinfo.update*/
type CrmMemberinfoUpdateRequest struct {
	values url.Values
}

func (this *CrmMemberinfoUpdateRequest) GetApiMethodName() string {
	return "taobao.crm.memberinfo.update"
}
func (this *CrmMemberinfoUpdateRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *CrmMemberinfoUpdateRequest) GetValues() url.Values {
	return this.values
}

/* 买家昵称<br /> 支持最大长度为：32<br /> 支持的最大列表长度为：32 */
func (this *CrmMemberinfoUpdateRequest) SetBuyerNick(value string) {
	this.Set("buyer_nick", value)
}

/* 城市.
请注意:从2014.4.15之后,省市将采用地区标准码,请通过物流API taobao.areas.get接口获取,参考:http://api.taobao.com/apidoc/api.htm?spm=0.0.0.0.nOOF9g&path=cid:7-apiId:59.API对于老的省市代码兼容会逐步下线 */
func (this *CrmMemberinfoUpdateRequest) SetCity(value string) {
	this.Set("city", value)
}

/* 交易关闭金额，单位：分 */
func (this *CrmMemberinfoUpdateRequest) SetCloseTradeAmount(value string) {
	this.Set("close_trade_amount", value)
}

/* 交易关闭次数 */
func (this *CrmMemberinfoUpdateRequest) SetCloseTradeCount(value string) {
	this.Set("close_trade_count", value)
}

/* 会员等级，1：普通客户，2：高级会员，3：高级会员 ，4：至尊vip

只有正常会员才给予升级，对于status blacklist的会员升级无效<br /> 支持最大值为：4<br /> 支持最小值为：1<br /> 支持的最大列表长度为：32 */
func (this *CrmMemberinfoUpdateRequest) SetGrade(value string) {
	this.Set("grade", value)
}

/* 分组的id集合字符串 */
func (this *CrmMemberinfoUpdateRequest) SetGroupIds(value string) {
	this.Set("group_ids", value)
}

/* 宝贝件数 */
func (this *CrmMemberinfoUpdateRequest) SetItemNum(value string) {
	this.Set("item_num", value)
}

/* 北京=1,天津=2,河北省=3,山西省=4,内蒙古自治区=5,辽宁省=6,吉林省=7,黑龙江省=8,上海=9,江苏省=10,浙江省=11,安徽省=12,福建省=13,江西省=14,山东省=15,河南省=16,湖北省=17,湖南省=18, 广东省=19,广西壮族自治区=20,海南省=21,重庆=22,四川省=23,贵州省=24,云南省=25,西藏自治区=26,陕西省=27,甘肃省=28,青海省=29,宁夏回族自治区=30,新疆维吾尔自治区=31,台湾省=32,香港特别行政区=33,澳门特别行政区=34,海外=35，约定36为清除Province设置.
请注意:从2014.4.15之后,省市将采用地区标准码,请通过物流API taobao.areas.get接口获取,参考:http://api.taobao.com/apidoc/api.htm?spm=0.0.0.0.nOOF9g&path=cid:7-apiId:59.API对于老的省市代码兼容会逐步下线. */
func (this *CrmMemberinfoUpdateRequest) SetProvince(value string) {
	this.Set("province", value)
}

/* 用于描述会员的状态，normal表示正常，blacklist表示黑名单(不享受会员折扣).<br /> 支持最大长度为：32<br /> 支持的最大列表长度为：32 */
func (this *CrmMemberinfoUpdateRequest) SetStatus(value string) {
	this.Set("status", value)
}

/* 交易金额，单位：分 */
func (this *CrmMemberinfoUpdateRequest) SetTradeAmount(value string) {
	this.Set("trade_amount", value)
}

/* 交易笔数 */
func (this *CrmMemberinfoUpdateRequest) SetTradeCount(value string) {
	this.Set("trade_count", value)
}

type CrmMemberinfoUpdateResponse struct {
	*ErrorResponse                    `json:"error_response,omitempty"`
	CrmMemberinfoUpdateResponseResult `json:"crm_memberinfo_update_response"`
}
type CrmMemberinfoUpdateResponseResult struct {
	/* 会员信息修改是否成功 */
	IsSuccess bool `json:"is_success"`
}

/*taobao.crm.members.get*/
type CrmMembersGetRequest struct {
	values url.Values
}

func (this *CrmMembersGetRequest) GetApiMethodName() string {
	return "taobao.crm.members.get"
}
func (this *CrmMembersGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *CrmMembersGetRequest) GetValues() url.Values {
	return this.values
}

/* 买家的昵称<br /> 支持最大长度为：1000<br /> 支持的最大列表长度为：1000 */
func (this *CrmMembersGetRequest) SetBuyerNick(value string) {
	this.Set("buyer_nick", value)
}

/* 显示第几页的会员，如果输入的页码大于总共的页码数，例如总共10页，但是current_page的值为11，则返回空白页，最小页数为1，最大页数为1000<br /> 支持最大值为：1000<br /> 支持最小值为：1 */
func (this *CrmMembersGetRequest) SetCurrentPage(value string) {
	this.Set("current_page", value)
}

/* 会员等级，0：店铺客户，1：普通会员，2：高级会员，3：VIP会员， 4：至尊VIP会员。如果不传入值则默认为全部等级。<br /> 支持最大值为：4<br /> 支持最小值为：-1<br /> 支持的最大列表长度为：32 */
func (this *CrmMembersGetRequest) SetGrade(value string) {
	this.Set("grade", value)
}

/* 最迟上次交易时间 */
func (this *CrmMembersGetRequest) SetMaxLastTradeTime(value string) {
	this.Set("max_last_trade_time", value)
}

/* 最大交易额，单位为元 */
func (this *CrmMembersGetRequest) SetMaxTradeAmount(value string) {
	this.Set("max_trade_amount", value)
}

/* 最大交易量<br /> 支持最小值为：0 */
func (this *CrmMembersGetRequest) SetMaxTradeCount(value string) {
	this.Set("max_trade_count", value)
}

/* 最早上次交易时间 */
func (this *CrmMembersGetRequest) SetMinLastTradeTime(value string) {
	this.Set("min_last_trade_time", value)
}

/* 最小交易额,单位为元 */
func (this *CrmMembersGetRequest) SetMinTradeAmount(value string) {
	this.Set("min_trade_amount", value)
}

/* 最小交易量<br /> 支持最小值为：0 */
func (this *CrmMembersGetRequest) SetMinTradeCount(value string) {
	this.Set("min_trade_count", value)
}

/* 表示每页显示的会员数量,page_size的最大值不能超过100条,最小值不能低于1，<br /> 支持最大值为：100<br /> 支持最小值为：1 */
func (this *CrmMembersGetRequest) SetPageSize(value string) {
	this.Set("page_size", value)
}

type CrmMembersGetResponse struct {
	*ErrorResponse              `json:"error_response,omitempty"`
	CrmMembersGetResponseResult `json:"crm_members_get_response"`
}
type CrmMembersGetResponseResult struct {
	/* 根据一定条件查询到卖家的会员 */
	Members struct {
		BasicMember []*BasicMember `json:"basic_member"`
	} `json:"Members"`
	/* 记录总数 */
	TotalResult int64 `json:"total_result"`
}

/*taobao.crm.members.group.batchadd*/
type CrmMembersGroupBatchaddRequest struct {
	values url.Values
}

func (this *CrmMembersGroupBatchaddRequest) GetApiMethodName() string {
	return "taobao.crm.members.group.batchadd"
}
func (this *CrmMembersGroupBatchaddRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *CrmMembersGroupBatchaddRequest) GetValues() url.Values {
	return this.values
}

/* 会员的id（一次最多传入10个）<br /> 支持最小值为：1 */
func (this *CrmMembersGroupBatchaddRequest) SetBuyerIds(value string) {
	this.Set("buyer_ids", value)
}

/* 分组id<br /> 支持最小值为：1 */
func (this *CrmMembersGroupBatchaddRequest) SetGroupIds(value string) {
	this.Set("group_ids", value)
}

type CrmMembersGroupBatchaddResponse struct {
	*ErrorResponse                        `json:"error_response,omitempty"`
	CrmMembersGroupBatchaddResponseResult `json:"crm_members_group_batchadd_response"`
}
type CrmMembersGroupBatchaddResponseResult struct {
	/* 添加操作是否成功 */
	IsSuccess bool `json:"is_success"`
}

/*taobao.crm.members.groups.batchdelete*/
type CrmMembersGroupsBatchdeleteRequest struct {
	values url.Values
}

func (this *CrmMembersGroupsBatchdeleteRequest) GetApiMethodName() string {
	return "taobao.crm.members.groups.batchdelete"
}
func (this *CrmMembersGroupsBatchdeleteRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *CrmMembersGroupsBatchdeleteRequest) GetValues() url.Values {
	return this.values
}

/* 买家的Id集合<br /> 支持最小值为：1 */
func (this *CrmMembersGroupsBatchdeleteRequest) SetBuyerIds(value string) {
	this.Set("buyer_ids", value)
}

/* 会员需要删除的分组<br /> 支持最小值为：1 */
func (this *CrmMembersGroupsBatchdeleteRequest) SetGroupIds(value string) {
	this.Set("group_ids", value)
}

type CrmMembersGroupsBatchdeleteResponse struct {
	*ErrorResponse                            `json:"error_response,omitempty"`
	CrmMembersGroupsBatchdeleteResponseResult `json:"crm_members_groups_batchdelete_response"`
}
type CrmMembersGroupsBatchdeleteResponseResult struct {
	/* 删除是否成功 */
	IsSuccess bool `json:"is_success"`
}

/*taobao.crm.members.increment.get*/
type CrmMembersIncrementGetRequest struct {
	values url.Values
}

func (this *CrmMembersIncrementGetRequest) GetApiMethodName() string {
	return "taobao.crm.members.increment.get"
}
func (this *CrmMembersIncrementGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *CrmMembersIncrementGetRequest) GetValues() url.Values {
	return this.values
}

/* 显示第几页的会员，如果输入的页码大于总共的页码数，例如总共10页，但是current_page的值为11，则返回空白页，最小页数为1<br /> 支持最小值为：1 */
func (this *CrmMembersIncrementGetRequest) SetCurrentPage(value string) {
	this.Set("current_page", value)
}

/* 卖家修改会员信息的时间终点.如果不填写此字段,默认为当前时间. */
func (this *CrmMembersIncrementGetRequest) SetEndModify(value string) {
	this.Set("end_modify", value)
}

/* 会员等级，0：店铺客户，1：普通会员，2：高级会员，3：VIP会员， 4：至尊VIP会员<br /> 支持最大值为：4<br /> 支持最小值为：-1<br /> 支持的最大列表长度为：32 */
func (this *CrmMembersIncrementGetRequest) SetGrade(value string) {
	this.Set("grade", value)
}

/* 每页显示的会员数，page_size的值不能超过100，最小值要大于1<br /> 支持最大值为：100<br /> 支持最小值为：1 */
func (this *CrmMembersIncrementGetRequest) SetPageSize(value string) {
	this.Set("page_size", value)
}

/* 卖家修改会员信息的时间起点. */
func (this *CrmMembersIncrementGetRequest) SetStartModify(value string) {
	this.Set("start_modify", value)
}

type CrmMembersIncrementGetResponse struct {
	*ErrorResponse                       `json:"error_response,omitempty"`
	CrmMembersIncrementGetResponseResult `json:"crm_members_increment_get_response"`
}
type CrmMembersIncrementGetResponseResult struct {
	/* 返回当前页的会员列表 */
	Members struct {
		BasicMember []*BasicMember `json:"basic_member"`
	} `json:"Members"`
	/* 记录的总条数 */
	TotalResult int64 `json:"total_result"`
}

/*taobao.crm.members.search*/
type CrmMembersSearchRequest struct {
	values url.Values
}

func (this *CrmMembersSearchRequest) GetApiMethodName() string {
	return "taobao.crm.members.search"
}
func (this *CrmMembersSearchRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *CrmMembersSearchRequest) GetValues() url.Values {
	return this.values
}

/* 买家昵称<br /> 支持最大长度为：1000<br /> 支持的最大列表长度为：1000 */
func (this *CrmMembersSearchRequest) SetBuyerNick(value string) {
	this.Set("buyer_nick", value)
}

/* 城市.
请注意:该字段从2014-4-23之后不再支持作为搜索条件检索. */
func (this *CrmMembersSearchRequest) SetCity(value string) {
	this.Set("city", value)
}

/* 显示第几页的会员，如果输入的页码大于总共的页码数，例如总共10页，但是current_page的值为11，则返回空白页，最小页数为1.最大1000页<br /> 支持最大值为：1000<br /> 支持最小值为：1 */
func (this *CrmMembersSearchRequest) SetCurrentPage(value string) {
	this.Set("current_page", value)
}

/* 会员等级，0：店铺客户，1：普通客户，2：高级会员，3：VIP会员, 4：至尊VIP会员<br /> 支持最大值为：4<br /> 支持最小值为：-1<br /> 支持的最大列表长度为：32 */
func (this *CrmMembersSearchRequest) SetGrade(value string) {
	this.Set("grade", value)
}

/* 分组id<br /> 支持的最大列表长度为：19 */
func (this *CrmMembersSearchRequest) SetGroupId(value string) {
	this.Set("group_id", value)
}

/* 最大平均客单价，单位为元.
请注意:该字段从2014-4-23之后不再支持作为搜索条件检索. */
func (this *CrmMembersSearchRequest) SetMaxAvgPrice(value string) {
	this.Set("max_avg_price", value)
}

/* 最大交易关闭笔数.
请注意:该字段从2014-4-23之后不再支持作为搜索条件检索.<br /> 支持最小值为：0 */
func (this *CrmMembersSearchRequest) SetMaxCloseTradeNum(value string) {
	this.Set("max_close_trade_num", value)
}

/* 最大交易宝贝件数<br /> 支持最小值为：0 */
func (this *CrmMembersSearchRequest) SetMaxItemNum(value string) {
	this.Set("max_item_num", value)
}

/* 最迟上次交易时间 */
func (this *CrmMembersSearchRequest) SetMaxLastTradeTime(value string) {
	this.Set("max_last_trade_time", value)
}

/* 最大交易额，单位为元 */
func (this *CrmMembersSearchRequest) SetMaxTradeAmount(value string) {
	this.Set("max_trade_amount", value)
}

/* 最大交易量<br /> 支持最小值为：0 */
func (this *CrmMembersSearchRequest) SetMaxTradeCount(value string) {
	this.Set("max_trade_count", value)
}

/* 最少平均客单价，单位为元.
请注意:该字段从2014-4-23之后不再支持作为搜索条件检索. */
func (this *CrmMembersSearchRequest) SetMinAvgPrice(value string) {
	this.Set("min_avg_price", value)
}

/* 最小交易关闭的笔数.
请注意:该字段从2014-4-23之后不再支持作为搜索条件检索.<br /> 支持最小值为：0 */
func (this *CrmMembersSearchRequest) SetMinCloseTradeNum(value string) {
	this.Set("min_close_trade_num", value)
}

/* 最小交易宝贝件数<br /> 支持最小值为：0 */
func (this *CrmMembersSearchRequest) SetMinItemNum(value string) {
	this.Set("min_item_num", value)
}

/* 最早上次交易时间 */
func (this *CrmMembersSearchRequest) SetMinLastTradeTime(value string) {
	this.Set("min_last_trade_time", value)
}

/* 最小交易额，单位为元 */
func (this *CrmMembersSearchRequest) SetMinTradeAmount(value string) {
	this.Set("min_trade_amount", value)
}

/* 最小交易量<br /> 支持最小值为：0 */
func (this *CrmMembersSearchRequest) SetMinTradeCount(value string) {
	this.Set("min_trade_count", value)
}

/* 每页显示的会员数量，page_size的最大值不能超过100，最小值不能小于1<br /> 支持最大值为：100<br /> 支持最小值为：1 */
func (this *CrmMembersSearchRequest) SetPageSize(value string) {
	this.Set("page_size", value)
}

/* 北京=1,天津=2,河北省=3,山西省=4,内蒙古自治区=5,辽宁省=6,吉林省=7,黑龙江省=8,上海=9,江苏省=10,浙江省=11,安徽省=12,福建省=13,江西省=14,山东省=15,河南省=16,湖北省=17,湖南省=18, 广东省=19,广西壮族自治区=20,海南省=21,重庆=22,四川省=23,贵州省=24,云南省=25,西藏自治区26,陕西省=27,甘肃省=28,青海省=29,宁夏回族自治区=30,新疆维吾尔自治区=31,台湾省=32,香港特别行政区=33,澳门特别行政区=34,海外=35.
注:请注意:从2014.4.23之后,省市将采用地区标准码,请通过物流API taobao.areas.get接口获取,参考:http://api.taobao.com/apidoc/api.htm?spm=0.0.0.0.nOOF9g&path=cid:7-apiId:59.API对于老的省市代码兼容会逐步下线.<br /> 支持最大值为：1000000<br /> 支持最小值为：1 */
func (this *CrmMembersSearchRequest) SetProvince(value string) {
	this.Set("province", value)
}

/* 关系来源，1交易成功，2未成交，3卖家手动吸纳<br /> 支持最大值为：3<br /> 支持最小值为：1<br /> 支持的最大列表长度为：32 */
func (this *CrmMembersSearchRequest) SetRelationSource(value string) {
	this.Set("relation_source", value)
}

type CrmMembersSearchResponse struct {
	*ErrorResponse                 `json:"error_response,omitempty"`
	CrmMembersSearchResponseResult `json:"crm_members_search_response"`
}
type CrmMembersSearchResponseResult struct {
	/* 根据一定条件查询的卖家会员 */
	Members struct {
		CrmMember []*CrmMember `json:"crm_member"`
	} `json:"Members"`
	/* 记录的总条数 */
	TotalResult int64 `json:"total_result"`
}

/*taobao.crm.qn.labels.seller.get*/
type CrmQnLabelsSellerGetRequest struct {
	values url.Values
}

func (this *CrmQnLabelsSellerGetRequest) GetApiMethodName() string {
	return "taobao.crm.qn.labels.seller.get"
}
func (this *CrmQnLabelsSellerGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *CrmQnLabelsSellerGetRequest) GetValues() url.Values {
	return this.values
}

type CrmQnLabelsSellerGetResponse struct {
	*ErrorResponse                     `json:"error_response,omitempty"`
	CrmQnLabelsSellerGetResponseResult `json:"crm_qn_labels_seller_get_response"`
}
type CrmQnLabelsSellerGetResponseResult struct {
	/* 查询到的当前卖家的当前页的会员 */
	Groups struct {
		Group []*Group `json:"group"`
	} `json:"Groups"`
}

/*taobao.crm.shopvip.cancel*/
type CrmShopvipCancelRequest struct {
	values url.Values
}

func (this *CrmShopvipCancelRequest) GetApiMethodName() string {
	return "taobao.crm.shopvip.cancel"
}
func (this *CrmShopvipCancelRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *CrmShopvipCancelRequest) GetValues() url.Values {
	return this.values
}

type CrmShopvipCancelResponse struct {
	*ErrorResponse                 `json:"error_response,omitempty"`
	CrmShopvipCancelResponseResult `json:"crm_shopvip_cancel_response"`
}
type CrmShopvipCancelResponseResult struct {
	/* 返回操作是否成功 */
	IsSuccess bool `json:"is_success"`
}

/*tmall.crm.equity.get*/
type TmallCrmEquityGetRequest struct {
	values url.Values
}

func (this *TmallCrmEquityGetRequest) GetApiMethodName() string {
	return "tmall.crm.equity.get"
}
func (this *TmallCrmEquityGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *TmallCrmEquityGetRequest) GetValues() url.Values {
	return this.values
}

type TmallCrmEquityGetResponse struct {
	*ErrorResponse                  `json:"error_response,omitempty"`
	TmallCrmEquityGetResponseResult `json:"tmall_crm_equity_get_response"`
}
type TmallCrmEquityGetResponseResult struct {
	/* 天猫卖家设置的等级权益 */
	GradeEquitys struct {
		GradeEquity []*GradeEquity `json:"grade_equity"`
	} `json:"GradeEquitys"`
}

/*taobao.picture.category.add*/
type PictureCategoryAddRequest struct {
	values url.Values
}

func (this *PictureCategoryAddRequest) GetApiMethodName() string {
	return "taobao.picture.category.add"
}
func (this *PictureCategoryAddRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *PictureCategoryAddRequest) GetValues() url.Values {
	return this.values
}

/* 图片分类的父分类,一级分类的parent_id为0,二级分类的则为其父分类的picture_category_id */
func (this *PictureCategoryAddRequest) SetParentId(value string) {
	this.Set("parent_id", value)
}

/* 图片分类名称，最大长度20字符，中文字符算2个字符，不能为空 */
func (this *PictureCategoryAddRequest) SetPictureCategoryName(value string) {
	this.Set("picture_category_name", value)
}

type PictureCategoryAddResponse struct {
	*ErrorResponse                   `json:"error_response,omitempty"`
	PictureCategoryAddResponseResult `json:"picture_category_add_response"`
}
type PictureCategoryAddResponseResult struct {
	/* 图片分类信息 */
	PictureCategory *PictureCategory `json:"picture_category"`
}

/*taobao.picture.category.get*/
type PictureCategoryGetRequest struct {
	values url.Values
}

func (this *PictureCategoryGetRequest) GetApiMethodName() string {
	return "taobao.picture.category.get"
}
func (this *PictureCategoryGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *PictureCategoryGetRequest) GetValues() url.Values {
	return this.values
}

/* 图片分类的修改时间点，格式:yyyy-MM-dd HH:mm:ss。查询此修改时间点之后到目前的图片分类。 */
func (this *PictureCategoryGetRequest) SetModifiedTime(value string) {
	this.Set("modified_time", value)
}

/* 取二级分类时设置为对应父分类id
取一级分类时父分类id设为0
取全部分类的时候不设或设为-1 */
func (this *PictureCategoryGetRequest) SetParentId(value string) {
	this.Set("parent_id", value)
}

/* 图片分类ID */
func (this *PictureCategoryGetRequest) SetPictureCategoryId(value string) {
	this.Set("picture_category_id", value)
}

/* 图片分类名，不支持模糊查询 */
func (this *PictureCategoryGetRequest) SetPictureCategoryName(value string) {
	this.Set("picture_category_name", value)
}

/* 分类类型,fixed代表店铺装修分类类别，auction代表宝贝分类类别，user-define代表用户自定义分类类别 */
func (this *PictureCategoryGetRequest) SetType(value string) {
	this.Set("type", value)
}

type PictureCategoryGetResponse struct {
	*ErrorResponse                   `json:"error_response,omitempty"`
	PictureCategoryGetResponseResult `json:"picture_category_get_response"`
}
type PictureCategoryGetResponseResult struct {
	/* 图片分类 */
	PictureCategories struct {
		PictureCategory []*PictureCategory `json:"picture_category"`
	} `json:"PictureCategories"`
}

/*taobao.picture.category.update*/
type PictureCategoryUpdateRequest struct {
	values url.Values
}

func (this *PictureCategoryUpdateRequest) GetApiMethodName() string {
	return "taobao.picture.category.update"
}
func (this *PictureCategoryUpdateRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *PictureCategoryUpdateRequest) GetValues() url.Values {
	return this.values
}

/* 要更新的图片分类的id */
func (this *PictureCategoryUpdateRequest) SetCategoryId(value string) {
	this.Set("category_id", value)
}

/* 图片分类的新名字，最大长度20字符，不能为空<br /> 支持最大长度为：20<br /> 支持的最大列表长度为：20 */
func (this *PictureCategoryUpdateRequest) SetCategoryName(value string) {
	this.Set("category_name", value)
}

/* 图片分类的新父分类id */
func (this *PictureCategoryUpdateRequest) SetParentId(value string) {
	this.Set("parent_id", value)
}

type PictureCategoryUpdateResponse struct {
	*ErrorResponse                      `json:"error_response,omitempty"`
	PictureCategoryUpdateResponseResult `json:"picture_category_update_response"`
}
type PictureCategoryUpdateResponseResult struct {
	/* 更新图片分类是否成功 */
	Done bool `json:"done"`
}

/*taobao.picture.delete*/
type PictureDeleteRequest struct {
	values url.Values
}

func (this *PictureDeleteRequest) GetApiMethodName() string {
	return "taobao.picture.delete"
}
func (this *PictureDeleteRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *PictureDeleteRequest) GetValues() url.Values {
	return this.values
}

/* 图片ID字符串,可以一个也可以一组,用英文逗号间隔,如450,120,155.限制数量是100 */
func (this *PictureDeleteRequest) SetPictureIds(value string) {
	this.Set("picture_ids", value)
}

type PictureDeleteResponse struct {
	*ErrorResponse              `json:"error_response,omitempty"`
	PictureDeleteResponseResult `json:"picture_delete_response"`
}
type PictureDeleteResponseResult struct {
	/* 是否删除 */
	Success bool `json:"success"`
}

/*taobao.picture.get*/
type PictureGetRequest struct {
	values url.Values
}

func (this *PictureGetRequest) GetApiMethodName() string {
	return "taobao.picture.get"
}
func (this *PictureGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *PictureGetRequest) GetValues() url.Values {
	return this.values
}

/* 图片使用，如果是pc宝贝detail使用，设置为client:computer，查询出来的图片是符合pc的宝贝detail显示的
如果是手机宝贝detail使用，设置为client:phone，查询出来的图片是符合手机的宝贝detail显示的,默认值是全部 */
func (this *PictureGetRequest) SetClientType(value string) {
	this.Set("client_type", value)
}

/* 是否删除,unfroze代表没有删除 */
func (this *PictureGetRequest) SetDeleted(value string) {
	this.Set("deleted", value)
}

/* 查询上传结束时间点,格式:yyyy-MM-dd HH:mm:ss */
func (this *PictureGetRequest) SetEndDate(value string) {
	this.Set("end_date", value)
}

/* 图片被修改的时间点，格式:yyyy-MM-dd HH:mm:ss。查询此修改时间点之后到目前的图片。 */
func (this *PictureGetRequest) SetModifiedTime(value string) {
	this.Set("modified_time", value)
}

/* 图片查询结果排序,time:desc按上传时间从晚到早(默认), time:asc按上传时间从早到晚,sizes:desc按图片从大到小，sizes:asc按图片从小到大,默认time:desc */
func (this *PictureGetRequest) SetOrderBy(value string) {
	this.Set("order_by", value)
}

/* 页码.传入值为1代表第一页,传入值为2代表第二页,依此类推,默认值为1 */
func (this *PictureGetRequest) SetPageNo(value string) {
	this.Set("page_no", value)
}

/* 每页条数.每页返回最多返回100条,默认值40 */
func (this *PictureGetRequest) SetPageSize(value string) {
	this.Set("page_size", value)
}

/* 图片分类ID */
func (this *PictureGetRequest) SetPictureCategoryId(value string) {
	this.Set("picture_category_id", value)
}

/* 图片ID */
func (this *PictureGetRequest) SetPictureId(value string) {
	this.Set("picture_id", value)
}

/* 查询上传开始时间点,格式:yyyy-MM-dd HH:mm:ss */
func (this *PictureGetRequest) SetStartDate(value string) {
	this.Set("start_date", value)
}

/* 图片标题,最大长度50字符,中英文都算一字符 */
func (this *PictureGetRequest) SetTitle(value string) {
	this.Set("title", value)
}

/* 图片url查询接口 */
func (this *PictureGetRequest) SetUrls(value string) {
	this.Set("urls", value)
}

type PictureGetResponse struct {
	*ErrorResponse           `json:"error_response,omitempty"`
	PictureGetResponseResult `json:"picture_get_response"`
}
type PictureGetResponseResult struct {
	/* 图片信息列表 */
	Pictures struct {
		Picture []*Picture `json:"picture"`
	} `json:"Pictures"`
	/* 总的结果数 */
	TotalResults int64 `json:"totalResults"`
}

/*taobao.picture.isreferenced.get*/
type PictureIsreferencedGetRequest struct {
	values url.Values
}

func (this *PictureIsreferencedGetRequest) GetApiMethodName() string {
	return "taobao.picture.isreferenced.get"
}
func (this *PictureIsreferencedGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *PictureIsreferencedGetRequest) GetValues() url.Values {
	return this.values
}

/* 图片id */
func (this *PictureIsreferencedGetRequest) SetPictureId(value string) {
	this.Set("picture_id", value)
}

type PictureIsreferencedGetResponse struct {
	*ErrorResponse                       `json:"error_response,omitempty"`
	PictureIsreferencedGetResponseResult `json:"picture_isreferenced_get_response"`
}
type PictureIsreferencedGetResponseResult struct {
	/* 图片是否被引用 */
	IsReferenced bool `json:"is_referenced"`
}

/*taobao.picture.replace*/
type PictureReplaceRequest struct {
	values url.Values
}

func (this *PictureReplaceRequest) GetApiMethodName() string {
	return "taobao.picture.replace"
}
func (this *PictureReplaceRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *PictureReplaceRequest) GetValues() url.Values {
	return this.values
}

/* 图片二进制文件流,不能为空,允许png、jpg、gif图片格式 */
func (this *PictureReplaceRequest) SetImageData(value string) {
	this.Set("image_data", value)
}

/* 要替换的图片的id，必须大于0 */
func (this *PictureReplaceRequest) SetPictureId(value string) {
	this.Set("picture_id", value)
}

type PictureReplaceResponse struct {
	*ErrorResponse               `json:"error_response,omitempty"`
	PictureReplaceResponseResult `json:"picture_replace_response"`
}
type PictureReplaceResponseResult struct {
	/* 图片替换是否成功 */
	Done bool `json:"done"`
}

/*taobao.picture.update*/
type PictureUpdateRequest struct {
	values url.Values
}

func (this *PictureUpdateRequest) GetApiMethodName() string {
	return "taobao.picture.update"
}
func (this *PictureUpdateRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *PictureUpdateRequest) GetValues() url.Values {
	return this.values
}

/* 新的图片名，最大长度50字符，不能为空<br /> 支持最大长度为：50<br /> 支持的最大列表长度为：50 */
func (this *PictureUpdateRequest) SetNewName(value string) {
	this.Set("new_name", value)
}

/* 要更改名字的图片的id */
func (this *PictureUpdateRequest) SetPictureId(value string) {
	this.Set("picture_id", value)
}

type PictureUpdateResponse struct {
	*ErrorResponse              `json:"error_response,omitempty"`
	PictureUpdateResponseResult `json:"picture_update_response"`
}
type PictureUpdateResponseResult struct {
	/* 更新是否成功 */
	Done bool `json:"done"`
}

/*taobao.picture.upload*/
type PictureUploadRequest struct {
	values url.Values
}

func (this *PictureUploadRequest) GetApiMethodName() string {
	return "taobao.picture.upload"
}
func (this *PictureUploadRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *PictureUploadRequest) GetValues() url.Values {
	return this.values
}

/* 图片上传的来源，有电脑版本宝贝发布，手机版本宝贝发布
client:computer电脑版本宝贝使用
client:phone手机版本宝贝使用 */
func (this *PictureUploadRequest) SetClientType(value string) {
	this.Set("client_type", value)
}

/* 包括后缀名的图片标题,不能为空，如Bule.jpg,有些卖家希望图片上传后取图片文件的默认名 */
func (this *PictureUploadRequest) SetImageInputTitle(value string) {
	this.Set("image_input_title", value)
}

/* 图片二进制文件流,不能为空,允许png、jpg、gif图片格式,3M以内。 */
func (this *PictureUploadRequest) SetImg(value string) {
	this.Set("img", value)
}

/* 图片分类ID，设置具体某个分类ID或设置0上传到默认分类，只能传入一个分类<br /> 支持最大值为：9223372036854775807<br /> 支持最小值为：0 */
func (this *PictureUploadRequest) SetPictureCategoryId(value string) {
	this.Set("picture_category_id", value)
}

/* 图片标题,如果为空,传的图片标题就取去掉后缀名的image_input_title,超过50字符长度会截取50字符,重名会在标题末尾加"(1)";标题末尾已经有"(数字)"了，则数字加1 */
func (this *PictureUploadRequest) SetTitle(value string) {
	this.Set("title", value)
}

type PictureUploadResponse struct {
	*ErrorResponse              `json:"error_response,omitempty"`
	PictureUploadResponseResult `json:"picture_upload_response"`
}
type PictureUploadResponseResult struct {
	/* 当前上传的一张图片信息 */
	Picture *Picture `json:"picture"`
}

/*taobao.picture.userinfo.get*/
type PictureUserinfoGetRequest struct {
	values url.Values
}

func (this *PictureUserinfoGetRequest) GetApiMethodName() string {
	return "taobao.picture.userinfo.get"
}
func (this *PictureUserinfoGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *PictureUserinfoGetRequest) GetValues() url.Values {
	return this.values
}

type PictureUserinfoGetResponse struct {
	*ErrorResponse                   `json:"error_response,omitempty"`
	PictureUserinfoGetResponseResult `json:"picture_userinfo_get_response"`
}
type PictureUserinfoGetResponseResult struct {
	/* 用户使用图片空间的信息 */
	UserInfo *UserInfo `json:"user_info"`
}

/*taobao.video.add*/
type VideoAddRequest struct {
	values url.Values
}

func (this *VideoAddRequest) GetApiMethodName() string {
	return "taobao.video.add"
}
func (this *VideoAddRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *VideoAddRequest) GetValues() url.Values {
	return this.values
}

/* 视频封面url,不能为空且不能超过512个英文字母<br /> 支持最大长度为：512<br /> 支持的最大列表长度为：512 */
func (this *VideoAddRequest) SetCoverUrl(value string) {
	this.Set("cover_url", value)
}

/* 视频描述信息，不能为空且不能超过256个汉字<br /> 支持最大长度为：512<br /> 支持的最大列表长度为：512 */
func (this *VideoAddRequest) SetDescription(value string) {
	this.Set("description", value)
}

/* 视频标签，以','隔开，不能为空且总长度不超过128个汉字<br /> 支持最大长度为：256<br /> 支持的最大列表长度为：256 */
func (this *VideoAddRequest) SetTags(value string) {
	this.Set("tags", value)
}

/* 视频标题，不能为空且不超过128个汉字<br /> 支持最大长度为：256<br /> 支持的最大列表长度为：256 */
func (this *VideoAddRequest) SetTitle(value string) {
	this.Set("title", value)
}

/* 上传唯一识别符,上传id */
func (this *VideoAddRequest) SetUploadId(value string) {
	this.Set("upload_id", value)
}

/* 视频上传者数字id */
func (this *VideoAddRequest) SetUploaderId(value string) {
	this.Set("uploader_id", value)
}

/* 在淘宝视频中的应用key，该值向淘宝视频申请产生 */
func (this *VideoAddRequest) SetVideoAppKey(value string) {
	this.Set("video_app_key", value)
}

type VideoAddResponse struct {
	*ErrorResponse         `json:"error_response,omitempty"`
	VideoAddResponseResult `json:"video_add_response"`
}
type VideoAddResponseResult struct {
	/* 发布成功返回视频id */
	Model int64 `json:"model"`
}

/*taobao.video.update*/
type VideoUpdateRequest struct {
	values url.Values
}

func (this *VideoUpdateRequest) GetApiMethodName() string {
	return "taobao.video.update"
}
func (this *VideoUpdateRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *VideoUpdateRequest) GetValues() url.Values {
	return this.values
}

/* 视频封面url,不能超过512个英文字母<br /> 支持最大长度为：512<br /> 支持的最大列表长度为：512 */
func (this *VideoUpdateRequest) SetCoverUrl(value string) {
	this.Set("cover_url", value)
}

/* 视频描述信息，不能超过256个汉字<br /> 支持最大长度为：512<br /> 支持的最大列表长度为：512 */
func (this *VideoUpdateRequest) SetDescription(value string) {
	this.Set("description", value)
}

/* 视频标签，以','隔开，且总长度不超过128个汉字<br /> 支持最大长度为：256<br /> 支持的最大列表长度为：256 */
func (this *VideoUpdateRequest) SetTags(value string) {
	this.Set("tags", value)
}

/* 视频标题，不超过128个汉字。title, tags,cover_url和description至少必须传一个<br /> 支持最大长度为：256<br /> 支持的最大列表长度为：256 */
func (this *VideoUpdateRequest) SetTitle(value string) {
	this.Set("title", value)
}

/* 在淘宝视频中的应用key，该值向淘宝视频申请产生 */
func (this *VideoUpdateRequest) SetVideoAppKey(value string) {
	this.Set("video_app_key", value)
}

/* 视频id */
func (this *VideoUpdateRequest) SetVideoId(value string) {
	this.Set("video_id", value)
}

type VideoUpdateResponse struct {
	*ErrorResponse            `json:"error_response,omitempty"`
	VideoUpdateResponseResult `json:"video_update_response"`
}
type VideoUpdateResponseResult struct {
	/* 更新视频成功返回true */
	Model bool `json:"model"`
}

/*taobao.videos.delete*/
type VideosDeleteRequest struct {
	values url.Values
}

func (this *VideosDeleteRequest) GetApiMethodName() string {
	return "taobao.videos.delete"
}
func (this *VideosDeleteRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *VideosDeleteRequest) GetValues() url.Values {
	return this.values
}

/* 在淘宝视频中的应用key，该值向淘宝视频申请产生 */
func (this *VideosDeleteRequest) SetVideoAppKey(value string) {
	this.Set("video_app_key", value)
}

/* 视频id列表 */
func (this *VideosDeleteRequest) SetVideoIds(value string) {
	this.Set("video_ids", value)
}

type VideosDeleteResponse struct {
	*ErrorResponse             `json:"error_response,omitempty"`
	VideosDeleteResponseResult `json:"videos_delete_response"`
}
type VideosDeleteResponseResult struct {
	/* 全部删除成功返回true，不支持部分删除 */
	Model bool `json:"model"`
}

/*taobao.videos.query*/
type VideosQueryRequest struct {
	values url.Values
}

func (this *VideosQueryRequest) GetApiMethodName() string {
	return "taobao.videos.query"
}
func (this *VideosQueryRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *VideosQueryRequest) GetValues() url.Values {
	return this.values
}

/* 需要返回的视频对象字段。VideoItem结构体中所有字段均可返回；多个字段用“,”分隔；其中video_play_info中的播放url可选择性返回，其余属性全部返回；如果想返回整个子对象中所有url，那字段为video_play_info，如果是想返回子对象里面的字段，那字段为video_play_info.web_url。 */
func (this *VideosQueryRequest) SetFields(value string) {
	this.Set("fields", value)
}

/* 在淘宝视频中的应用key，该值向淘宝视频申请产生 */
func (this *VideosQueryRequest) SetVideoAppKey(value string) {
	this.Set("video_app_key", value)
}

/* 视频id列表 */
func (this *VideosQueryRequest) SetVideoIds(value string) {
	this.Set("video_ids", value)
}

type VideosQueryResponse struct {
	*ErrorResponse            `json:"error_response,omitempty"`
	VideosQueryResponseResult `json:"videos_query_response"`
}
type VideosQueryResponseResult struct {
	/* 批量查询返回的视频列表，返回具体信息为入参fields请求的字段信息 */
	VideoItems struct {
		VideoItem []*VideoItem `json:"video_item"`
	} `json:"VideoItems"`
}

/*taobao.videos.search*/
type VideosSearchRequest struct {
	values url.Values
}

func (this *VideosSearchRequest) GetApiMethodName() string {
	return "taobao.videos.search"
}
func (this *VideosSearchRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *VideosSearchRequest) GetValues() url.Values {
	return this.values
}

/* 页码。默认返回的数据是从第一页开始 */
func (this *VideosSearchRequest) SetCurrentPage(value string) {
	this.Set("current_page", value)
}

/* 需要返回的视频对象字段。VideoItem结构体中所有字段均可返回；多个字段用“,”分隔；其中video_play_info中的播放url可选择性返回，其余属性全部返回；如果想返回整个子对象中所有url，那字段为video_play_info，如果是想返回子对象里面的字段，那字段为video_play_info.web_url。 */
func (this *VideosSearchRequest) SetFields(value string) {
	this.Set("fields", value)
}

/* 关键字(标题or标签，不能同时设置title,tag，否则冲突) */
func (this *VideosSearchRequest) SetKeywords(value string) {
	this.Set("keywords", value)
}

/* 每页条数，默认值是12 */
func (this *VideosSearchRequest) SetPageSize(value string) {
	this.Set("page_size", value)
}

/* 视频状态列表；视频状态：等待转码（1），转码中（2），转码失败（3），等待审核（4），未通过审核（5），通过审核（6） */
func (this *VideosSearchRequest) SetStates(value string) {
	this.Set("states", value)
}

/* 视频标签<br /> 支持最大长度为：256<br /> 支持的最大列表长度为：256 */
func (this *VideosSearchRequest) SetTag(value string) {
	this.Set("tag", value)
}

/* 视频标题<br /> 支持最大长度为：256<br /> 支持的最大列表长度为：256 */
func (this *VideosSearchRequest) SetTitle(value string) {
	this.Set("title", value)
}

/* 视频上传者数字id */
func (this *VideosSearchRequest) SetUploaderId(value string) {
	this.Set("uploader_id", value)
}

/* 在淘宝视频中的应用key，该值向淘宝视频申请产生 */
func (this *VideosSearchRequest) SetVideoAppKey(value string) {
	this.Set("video_app_key", value)
}

type VideosSearchResponse struct {
	*ErrorResponse             `json:"error_response,omitempty"`
	VideosSearchResponseResult `json:"videos_search_response"`
}
type VideosSearchResponseResult struct {
	/* 搜索返回类，返回具体信息为入参fields请求的字段信息 */
	SearchResult *TOPSearchResult `json:"search_result"`
}

/*taobao.sellercenter.role.add*/
type SellercenterRoleAddRequest struct {
	values url.Values
}

func (this *SellercenterRoleAddRequest) GetApiMethodName() string {
	return "taobao.sellercenter.role.add"
}
func (this *SellercenterRoleAddRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *SellercenterRoleAddRequest) GetValues() url.Values {
	return this.values
}

/* 角色描述<br /> 支持最大长度为：20<br /> 支持的最大列表长度为：20 */
func (this *SellercenterRoleAddRequest) SetDescription(value string) {
	this.Set("description", value)
}

/* 角色名<br /> 支持最大长度为：20<br /> 支持的最大列表长度为：20 */
func (this *SellercenterRoleAddRequest) SetName(value string) {
	this.Set("name", value)
}

/* 表示卖家昵称<br /> 支持最大长度为：500<br /> 支持的最大列表长度为：500 */
func (this *SellercenterRoleAddRequest) SetNick(value string) {
	this.Set("nick", value)
}

/* 需要授权的权限点permission_code列表,以","分割.其code值可以通过调用taobao.sellercenter.user.permissions.get返回，其中permission.is_authorize=1的权限点可以通过本接口授权给对应角色。 */
func (this *SellercenterRoleAddRequest) SetPermissionCodes(value string) {
	this.Set("permission_codes", value)
}

type SellercenterRoleAddResponse struct {
	*ErrorResponse                    `json:"error_response,omitempty"`
	SellercenterRoleAddResponseResult `json:"sellercenter_role_add_response"`
}
type SellercenterRoleAddResponseResult struct {
	/* 子账号角色 */
	Role *Role `json:"role"`
}

/*taobao.sellercenter.role.info.get*/
type SellercenterRoleInfoGetRequest struct {
	values url.Values
}

func (this *SellercenterRoleInfoGetRequest) GetApiMethodName() string {
	return "taobao.sellercenter.role.info.get"
}
func (this *SellercenterRoleInfoGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *SellercenterRoleInfoGetRequest) GetValues() url.Values {
	return this.values
}

/* 角色id<br /> 支持的最大列表长度为：100 */
func (this *SellercenterRoleInfoGetRequest) SetRoleId(value string) {
	this.Set("role_id", value)
}

type SellercenterRoleInfoGetResponse struct {
	*ErrorResponse                        `json:"error_response,omitempty"`
	SellercenterRoleInfoGetResponseResult `json:"sellercenter_role_info_get_response"`
}
type SellercenterRoleInfoGetResponseResult struct {
	/* 角色具体信息 */
	Role *Role `json:"role"`
}

/*taobao.sellercenter.rolemembers.get*/
type SellercenterRolemembersGetRequest struct {
	values url.Values
}

func (this *SellercenterRolemembersGetRequest) GetApiMethodName() string {
	return "taobao.sellercenter.rolemembers.get"
}
func (this *SellercenterRolemembersGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *SellercenterRolemembersGetRequest) GetValues() url.Values {
	return this.values
}

/* 角色id<br /> 支持的最大列表长度为：100 */
func (this *SellercenterRolemembersGetRequest) SetRoleId(value string) {
	this.Set("role_id", value)
}

type SellercenterRolemembersGetResponse struct {
	*ErrorResponse                           `json:"error_response,omitempty"`
	SellercenterRolemembersGetResponseResult `json:"sellercenter_rolemembers_get_response"`
}
type SellercenterRolemembersGetResponseResult struct {
	/* 子账号基本信息列表。具体信息为id、子账号用户名、主账号id、主账号昵称、当前状态值、是否分流 */
	Subusers struct {
		SubUserInfo []*SubUserInfo `json:"sub_user_info"`
	} `json:"Subusers"`
}

/*taobao.sellercenter.roles.get*/
type SellercenterRolesGetRequest struct {
	values url.Values
}

func (this *SellercenterRolesGetRequest) GetApiMethodName() string {
	return "taobao.sellercenter.roles.get"
}
func (this *SellercenterRolesGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *SellercenterRolesGetRequest) GetValues() url.Values {
	return this.values
}

/* 卖家昵称(只允许查询自己的信息：当前登陆者)<br /> 支持最大长度为：500<br /> 支持的最大列表长度为：500 */
func (this *SellercenterRolesGetRequest) SetNick(value string) {
	this.Set("nick", value)
}

type SellercenterRolesGetResponse struct {
	*ErrorResponse                     `json:"error_response,omitempty"`
	SellercenterRolesGetResponseResult `json:"sellercenter_roles_get_response"`
}
type SellercenterRolesGetResponseResult struct {
	/* 卖家子账号角色列表。<br/>返回对象为 role数据对象中的role_id，role_name，description，seller_id，create_time，modified_time。不包含permissions(权限点) */
	Roles struct {
		Role []*Role `json:"role"`
	} `json:"Roles"`
}

/*taobao.sellercenter.subuser.permissions.roles.get*/
type SellercenterSubuserPermissionsRolesGetRequest struct {
	values url.Values
}

func (this *SellercenterSubuserPermissionsRolesGetRequest) GetApiMethodName() string {
	return "taobao.sellercenter.subuser.permissions.roles.get"
}
func (this *SellercenterSubuserPermissionsRolesGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *SellercenterSubuserPermissionsRolesGetRequest) GetValues() url.Values {
	return this.values
}

/* 子账号昵称(子账号标识)<br /> 支持最大长度为：100<br /> 支持的最大列表长度为：100 */
func (this *SellercenterSubuserPermissionsRolesGetRequest) SetNick(value string) {
	this.Set("nick", value)
}

type SellercenterSubuserPermissionsRolesGetResponse struct {
	*ErrorResponse                                       `json:"error_response,omitempty"`
	SellercenterSubuserPermissionsRolesGetResponseResult `json:"sellercenter_subuser_permissions_roles_get_response"`
}
type SellercenterSubuserPermissionsRolesGetResponseResult struct {
	/* 子账号被所拥有的权限 */
	SubuserPermission *SubUserPermission `json:"subuser_permission"`
}

/*taobao.sellercenter.subusers.get*/
type SellercenterSubusersGetRequest struct {
	values url.Values
}

func (this *SellercenterSubusersGetRequest) GetApiMethodName() string {
	return "taobao.sellercenter.subusers.get"
}
func (this *SellercenterSubusersGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *SellercenterSubusersGetRequest) GetValues() url.Values {
	return this.values
}

/* 表示卖家昵称 */
func (this *SellercenterSubusersGetRequest) SetNick(value string) {
	this.Set("nick", value)
}

type SellercenterSubusersGetResponse struct {
	*ErrorResponse                        `json:"error_response,omitempty"`
	SellercenterSubusersGetResponseResult `json:"sellercenter_subusers_get_response"`
}
type SellercenterSubusersGetResponseResult struct {
	/* 子账号基本信息列表。具体信息为id、子账号用户名、主账号id、主账号昵称、当前状态值、是否分流 */
	Subusers struct {
		SubUserInfo []*SubUserInfo `json:"sub_user_info"`
	} `json:"Subusers"`
}

/*taobao.sellercenter.user.permissions.get*/
type SellercenterUserPermissionsGetRequest struct {
	values url.Values
}

func (this *SellercenterUserPermissionsGetRequest) GetApiMethodName() string {
	return "taobao.sellercenter.user.permissions.get"
}
func (this *SellercenterUserPermissionsGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *SellercenterUserPermissionsGetRequest) GetValues() url.Values {
	return this.values
}

/* 用户标识，次入参必须为子账号比如zhangsan:cool。如果只输入主账号zhangsan，将报错。 */
func (this *SellercenterUserPermissionsGetRequest) SetNick(value string) {
	this.Set("nick", value)
}

type SellercenterUserPermissionsGetResponse struct {
	*ErrorResponse                               `json:"error_response,omitempty"`
	SellercenterUserPermissionsGetResponseResult `json:"sellercenter_user_permissions_get_response"`
}
type SellercenterUserPermissionsGetResponseResult struct {
	/* 权限列表 */
	Permissions struct {
		Permission []*Permission `json:"permission"`
	} `json:"Permissions"`
}

/*taobao.subuser.department.add*/
type SubuserDepartmentAddRequest struct {
	values url.Values
}

func (this *SubuserDepartmentAddRequest) GetApiMethodName() string {
	return "taobao.subuser.department.add"
}
func (this *SubuserDepartmentAddRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *SubuserDepartmentAddRequest) GetValues() url.Values {
	return this.values
}

/* 部门名称 */
func (this *SubuserDepartmentAddRequest) SetDepartmentName(value string) {
	this.Set("department_name", value)
}

/* 父部门ID 如果是最高部门则传入0 */
func (this *SubuserDepartmentAddRequest) SetParentId(value string) {
	this.Set("parent_id", value)
}

/* 主账号用户名 */
func (this *SubuserDepartmentAddRequest) SetUserNick(value string) {
	this.Set("user_nick", value)
}

type SubuserDepartmentAddResponse struct {
	*ErrorResponse                     `json:"error_response,omitempty"`
	SubuserDepartmentAddResponseResult `json:"subuser_department_add_response"`
}
type SubuserDepartmentAddResponseResult struct {
	/* 操作是否成功 true:操作成功; false:操作失败 */
	IsSuccess bool `json:"is_success"`
}

/*taobao.subuser.department.delete*/
type SubuserDepartmentDeleteRequest struct {
	values url.Values
}

func (this *SubuserDepartmentDeleteRequest) GetApiMethodName() string {
	return "taobao.subuser.department.delete"
}
func (this *SubuserDepartmentDeleteRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *SubuserDepartmentDeleteRequest) GetValues() url.Values {
	return this.values
}

/* 部门ID */
func (this *SubuserDepartmentDeleteRequest) SetDepartmentId(value string) {
	this.Set("department_id", value)
}

/* 主账号用户名 */
func (this *SubuserDepartmentDeleteRequest) SetUserNick(value string) {
	this.Set("user_nick", value)
}

type SubuserDepartmentDeleteResponse struct {
	*ErrorResponse                        `json:"error_response,omitempty"`
	SubuserDepartmentDeleteResponseResult `json:"subuser_department_delete_response"`
}
type SubuserDepartmentDeleteResponseResult struct {
	/* 操作是否成功 true:操作成功; false:操作失败 */
	IsSuccess bool `json:"is_success"`
}

/*taobao.subuser.department.update*/
type SubuserDepartmentUpdateRequest struct {
	values url.Values
}

func (this *SubuserDepartmentUpdateRequest) GetApiMethodName() string {
	return "taobao.subuser.department.update"
}
func (this *SubuserDepartmentUpdateRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *SubuserDepartmentUpdateRequest) GetValues() url.Values {
	return this.values
}

/* 部门ID */
func (this *SubuserDepartmentUpdateRequest) SetDepartmentId(value string) {
	this.Set("department_id", value)
}

/* 部门名称 */
func (this *SubuserDepartmentUpdateRequest) SetDepartmentName(value string) {
	this.Set("department_name", value)
}

/* 父部门ID 如果是最高部门则传入0 */
func (this *SubuserDepartmentUpdateRequest) SetParentId(value string) {
	this.Set("parent_id", value)
}

/* 主账号用户名 */
func (this *SubuserDepartmentUpdateRequest) SetUserNick(value string) {
	this.Set("user_nick", value)
}

type SubuserDepartmentUpdateResponse struct {
	*ErrorResponse                        `json:"error_response,omitempty"`
	SubuserDepartmentUpdateResponseResult `json:"subuser_department_update_response"`
}
type SubuserDepartmentUpdateResponseResult struct {
	/* 操作是否成功 true:操作成功; false:操作失败 */
	IsSuccess bool `json:"is_success"`
}

/*taobao.subuser.departments.get*/
type SubuserDepartmentsGetRequest struct {
	values url.Values
}

func (this *SubuserDepartmentsGetRequest) GetApiMethodName() string {
	return "taobao.subuser.departments.get"
}
func (this *SubuserDepartmentsGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *SubuserDepartmentsGetRequest) GetValues() url.Values {
	return this.values
}

/* 主账号用户名 */
func (this *SubuserDepartmentsGetRequest) SetUserNick(value string) {
	this.Set("user_nick", value)
}

type SubuserDepartmentsGetResponse struct {
	*ErrorResponse                      `json:"error_response,omitempty"`
	SubuserDepartmentsGetResponseResult `json:"subuser_departments_get_response"`
}
type SubuserDepartmentsGetResponseResult struct {
	/* 部门信息 */
	Departments struct {
		Department []*Department `json:"department"`
	} `json:"Departments"`
}

/*taobao.subuser.duty.add*/
type SubuserDutyAddRequest struct {
	values url.Values
}

func (this *SubuserDutyAddRequest) GetApiMethodName() string {
	return "taobao.subuser.duty.add"
}
func (this *SubuserDutyAddRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *SubuserDutyAddRequest) GetValues() url.Values {
	return this.values
}

/* 职务级别 */
func (this *SubuserDutyAddRequest) SetDutyLevel(value string) {
	this.Set("duty_level", value)
}

/* 职务名称 */
func (this *SubuserDutyAddRequest) SetDutyName(value string) {
	this.Set("duty_name", value)
}

/* 主账号用户名 */
func (this *SubuserDutyAddRequest) SetUserNick(value string) {
	this.Set("user_nick", value)
}

type SubuserDutyAddResponse struct {
	*ErrorResponse               `json:"error_response,omitempty"`
	SubuserDutyAddResponseResult `json:"subuser_duty_add_response"`
}
type SubuserDutyAddResponseResult struct {
	/* 操作是否成功 true:操作成功; false:操作失败 */
	IsSuccess bool `json:"is_success"`
}

/*taobao.subuser.duty.delete*/
type SubuserDutyDeleteRequest struct {
	values url.Values
}

func (this *SubuserDutyDeleteRequest) GetApiMethodName() string {
	return "taobao.subuser.duty.delete"
}
func (this *SubuserDutyDeleteRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *SubuserDutyDeleteRequest) GetValues() url.Values {
	return this.values
}

/* 职务ID */
func (this *SubuserDutyDeleteRequest) SetDutyId(value string) {
	this.Set("duty_id", value)
}

/* 主账号用户名 */
func (this *SubuserDutyDeleteRequest) SetUserNick(value string) {
	this.Set("user_nick", value)
}

type SubuserDutyDeleteResponse struct {
	*ErrorResponse                  `json:"error_response,omitempty"`
	SubuserDutyDeleteResponseResult `json:"subuser_duty_delete_response"`
}
type SubuserDutyDeleteResponseResult struct {
	/* 操作是否成功 true:操作成功; false:操作失败 */
	IsSuccess bool `json:"is_success"`
}

/*taobao.subuser.duty.update*/
type SubuserDutyUpdateRequest struct {
	values url.Values
}

func (this *SubuserDutyUpdateRequest) GetApiMethodName() string {
	return "taobao.subuser.duty.update"
}
func (this *SubuserDutyUpdateRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *SubuserDutyUpdateRequest) GetValues() url.Values {
	return this.values
}

/* 职务ID */
func (this *SubuserDutyUpdateRequest) SetDutyId(value string) {
	this.Set("duty_id", value)
}

/* 职务级别 */
func (this *SubuserDutyUpdateRequest) SetDutyLevel(value string) {
	this.Set("duty_level", value)
}

/* 职务名称 */
func (this *SubuserDutyUpdateRequest) SetDutyName(value string) {
	this.Set("duty_name", value)
}

/* 主账号用户名 */
func (this *SubuserDutyUpdateRequest) SetUserNick(value string) {
	this.Set("user_nick", value)
}

type SubuserDutyUpdateResponse struct {
	*ErrorResponse                  `json:"error_response,omitempty"`
	SubuserDutyUpdateResponseResult `json:"subuser_duty_update_response"`
}
type SubuserDutyUpdateResponseResult struct {
	/* 操作是否成功 true:操作成功; false:操作失败 */
	IsSuccess bool `json:"is_success"`
}

/*taobao.subuser.dutys.get*/
type SubuserDutysGetRequest struct {
	values url.Values
}

func (this *SubuserDutysGetRequest) GetApiMethodName() string {
	return "taobao.subuser.dutys.get"
}
func (this *SubuserDutysGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *SubuserDutysGetRequest) GetValues() url.Values {
	return this.values
}

/* 主账号用户名 */
func (this *SubuserDutysGetRequest) SetUserNick(value string) {
	this.Set("user_nick", value)
}

type SubuserDutysGetResponse struct {
	*ErrorResponse                `json:"error_response,omitempty"`
	SubuserDutysGetResponseResult `json:"subuser_dutys_get_response"`
}
type SubuserDutysGetResponseResult struct {
	/* 职务信息 */
	Dutys struct {
		Duty []*Duty `json:"duty"`
	} `json:"Dutys"`
}

/*taobao.subuser.employee.add*/
type SubuserEmployeeAddRequest struct {
	values url.Values
}

func (this *SubuserEmployeeAddRequest) GetApiMethodName() string {
	return "taobao.subuser.employee.add"
}
func (this *SubuserEmployeeAddRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *SubuserEmployeeAddRequest) GetValues() url.Values {
	return this.values
}

/* 当前员工所属部门ID */
func (this *SubuserEmployeeAddRequest) SetDepartmentId(value string) {
	this.Set("department_id", value)
}

/* 当前员工担任职务ID */
func (this *SubuserEmployeeAddRequest) SetDutyId(value string) {
	this.Set("duty_id", value)
}

/* 员工姓名 */
func (this *SubuserEmployeeAddRequest) SetEmployeeName(value string) {
	this.Set("employee_name", value)
}

/* 员工花名 */
func (this *SubuserEmployeeAddRequest) SetEmployeeNickname(value string) {
	this.Set("employee_nickname", value)
}

/* 员工工号(可以用大小写英文字母和数字) */
func (this *SubuserEmployeeAddRequest) SetEmployeeNum(value string) {
	this.Set("employee_num", value)
}

/* 员工入职时间 */
func (this *SubuserEmployeeAddRequest) SetEntryDate(value string) {
	this.Set("entry_date", value)
}

/* 员工身份证号码 */
func (this *SubuserEmployeeAddRequest) SetIdCardNum(value string) {
	this.Set("id_card_num", value)
}

/* 直接上级的员工ID */
func (this *SubuserEmployeeAddRequest) SetLeaderId(value string) {
	this.Set("leader_id", value)
}

/* 办公电话 */
func (this *SubuserEmployeeAddRequest) SetOfficePhone(value string) {
	this.Set("office_phone", value)
}

/* 员工私人邮箱 */
func (this *SubuserEmployeeAddRequest) SetPersonalEmail(value string) {
	this.Set("personal_email", value)
}

/* 员工手机号码 */
func (this *SubuserEmployeeAddRequest) SetPersonalMobile(value string) {
	this.Set("personal_mobile", value)
}

/* 员工性别 1：男; 2:女 */
func (this *SubuserEmployeeAddRequest) SetSex(value string) {
	this.Set("sex", value)
}

/* 子账号ID */
func (this *SubuserEmployeeAddRequest) SetSubId(value string) {
	this.Set("sub_id", value)
}

/* 工作地点 */
func (this *SubuserEmployeeAddRequest) SetWorkLocation(value string) {
	this.Set("work_location", value)
}

type SubuserEmployeeAddResponse struct {
	*ErrorResponse                   `json:"error_response,omitempty"`
	SubuserEmployeeAddResponseResult `json:"subuser_employee_add_response"`
}
type SubuserEmployeeAddResponseResult struct {
	/* 操作是否成功 true:操作成功; false:操作失败 */
	IsSuccess bool `json:"is_success"`
}

/*taobao.subuser.employee.update*/
type SubuserEmployeeUpdateRequest struct {
	values url.Values
}

func (this *SubuserEmployeeUpdateRequest) GetApiMethodName() string {
	return "taobao.subuser.employee.update"
}
func (this *SubuserEmployeeUpdateRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *SubuserEmployeeUpdateRequest) GetValues() url.Values {
	return this.values
}

/* 当前员工所属部门ID */
func (this *SubuserEmployeeUpdateRequest) SetDepartmentId(value string) {
	this.Set("department_id", value)
}

/* 当前员工担任职务ID(若需要将该字段的值置为空，请传入-1） */
func (this *SubuserEmployeeUpdateRequest) SetDutyId(value string) {
	this.Set("duty_id", value)
}

/* 员工姓名 */
func (this *SubuserEmployeeUpdateRequest) SetEmployeeName(value string) {
	this.Set("employee_name", value)
}

/* 员工花名(若需要将该字段的值置为空，请传入“-1”） */
func (this *SubuserEmployeeUpdateRequest) SetEmployeeNickname(value string) {
	this.Set("employee_nickname", value)
}

/* 员工工号(卖家自定义)(若需要将该字段的值置为空，请传入“-1”） */
func (this *SubuserEmployeeUpdateRequest) SetEmployeeNum(value string) {
	this.Set("employee_num", value)
}

/* 登记员工离职  true:登记员工离职 */
func (this *SubuserEmployeeUpdateRequest) SetEmployeeTurnover(value string) {
	this.Set("employee_turnover", value)
}

/* 员工入职时间(若需要将该字段的值置为空，请传入1900-01-01 00:00:00） */
func (this *SubuserEmployeeUpdateRequest) SetEntryDate(value string) {
	this.Set("entry_date", value)
}

/* 员工身份证号码(若需要将该字段的值置为空，请传入“-1”） */
func (this *SubuserEmployeeUpdateRequest) SetIdCardNum(value string) {
	this.Set("id_card_num", value)
}

/* 直接上级的员工ID(若需要将该字段的值置为空，请传入-1） */
func (this *SubuserEmployeeUpdateRequest) SetLeaderId(value string) {
	this.Set("leader_id", value)
}

/* 办公电话(若需要将该字段的值置为空，请传入“-1”） */
func (this *SubuserEmployeeUpdateRequest) SetOfficePhone(value string) {
	this.Set("office_phone", value)
}

/* 员工私人邮箱(若需要将该字段的值置为空，请传入“-1”） */
func (this *SubuserEmployeeUpdateRequest) SetPersonalEmail(value string) {
	this.Set("personal_email", value)
}

/* 员工手机号码(若需要将该字段的值置为空，请传入“-1”） */
func (this *SubuserEmployeeUpdateRequest) SetPersonalMobile(value string) {
	this.Set("personal_mobile", value)
}

/* 员工性别  1：男;  2:女 */
func (this *SubuserEmployeeUpdateRequest) SetSex(value string) {
	this.Set("sex", value)
}

/* 子账号ID */
func (this *SubuserEmployeeUpdateRequest) SetSubId(value string) {
	this.Set("sub_id", value)
}

/* 杭州大厦(若需要将该字段的值置为空，请传入“-1”） */
func (this *SubuserEmployeeUpdateRequest) SetWorkLocation(value string) {
	this.Set("work_location", value)
}

type SubuserEmployeeUpdateResponse struct {
	*ErrorResponse                      `json:"error_response,omitempty"`
	SubuserEmployeeUpdateResponseResult `json:"subuser_employee_update_response"`
}
type SubuserEmployeeUpdateResponseResult struct {
	/* 操作是否成功 true:操作成功; false:操作失败 */
	IsSuccess bool `json:"is_success"`
}

/*taobao.subuser.fullinfo.get*/
type SubuserFullinfoGetRequest struct {
	values url.Values
}

func (this *SubuserFullinfoGetRequest) GetApiMethodName() string {
	return "taobao.subuser.fullinfo.get"
}
func (this *SubuserFullinfoGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *SubuserFullinfoGetRequest) GetValues() url.Values {
	return this.values
}

/* 传入所需要的参数信息（若不需要获取子账号或主账号的企业邮箱地址，则无需传入该参数；若需要获取子账号或主账号的企业邮箱地址，则需要传入fields；可选参数值为subuser_email和user_email，传入其他参数值均无效；两个参数都需要则以逗号隔开传入即可，例如：subuser_email,user_email） */
func (this *SubuserFullinfoGetRequest) SetFields(value string) {
	this.Set("fields", value)
}

/* 子账号ID（传参中sub_id和sub_nick至少需要其中一个，若sub_id与sub_nick同时传入并且合法，那么sub_nick优先，以sub_nick查询子账号） */
func (this *SubuserFullinfoGetRequest) SetSubId(value string) {
	this.Set("sub_id", value)
}

/* 子账号用户名（传参中sub_id和sub_nick至少需要其中一个，若sub_id与sub_nick同时传入并且合法，那么sub_nick优先，以sub_nick查询子账号） */
func (this *SubuserFullinfoGetRequest) SetSubNick(value string) {
	this.Set("sub_nick", value)
}

type SubuserFullinfoGetResponse struct {
	*ErrorResponse                   `json:"error_response,omitempty"`
	SubuserFullinfoGetResponseResult `json:"subuser_fullinfo_get_response"`
}
type SubuserFullinfoGetResponseResult struct {
	/* 子账号详细信息，其中包括账号基本信息、员工信息和部门职务信息 */
	SubFullinfo *SubUserFullInfo `json:"sub_fullinfo"`
}

/*taobao.subuser.info.update*/
type SubuserInfoUpdateRequest struct {
	values url.Values
}

func (this *SubuserInfoUpdateRequest) GetApiMethodName() string {
	return "taobao.subuser.info.update"
}
func (this *SubuserInfoUpdateRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *SubuserInfoUpdateRequest) GetValues() url.Values {
	return this.values
}

/* 是否停用子账号 true:表示停用该子账号false:表示开启该子账号 */
func (this *SubuserInfoUpdateRequest) SetIsDisableSubaccount(value string) {
	this.Set("is_disable_subaccount", value)
}

/* 子账号是否参与分流 true:参与分流 false:不参与分流 */
func (this *SubuserInfoUpdateRequest) SetIsDispatch(value string) {
	this.Set("is_dispatch", value)
}

/* 子账号ID */
func (this *SubuserInfoUpdateRequest) SetSubId(value string) {
	this.Set("sub_id", value)
}

type SubuserInfoUpdateResponse struct {
	*ErrorResponse                  `json:"error_response,omitempty"`
	SubuserInfoUpdateResponseResult `json:"subuser_info_update_response"`
}
type SubuserInfoUpdateResponseResult struct {
	/* 操作是否成功 true:操作成功; false:操作失败 */
	IsSuccess bool `json:"is_success"`
}

/*taobao.subusers.get*/
type SubusersGetRequest struct {
	values url.Values
}

func (this *SubusersGetRequest) GetApiMethodName() string {
	return "taobao.subusers.get"
}
func (this *SubusersGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *SubusersGetRequest) GetValues() url.Values {
	return this.values
}

/* 主账号用户名 */
func (this *SubusersGetRequest) SetUserNick(value string) {
	this.Set("user_nick", value)
}

type SubusersGetResponse struct {
	*ErrorResponse            `json:"error_response,omitempty"`
	SubusersGetResponseResult `json:"subusers_get_response"`
}
type SubusersGetResponseResult struct {
	/* 子账号基本信息 */
	Subaccounts struct {
		SubAccountInfo []*SubAccountInfo `json:"sub_account_info"`
	} `json:"Subaccounts"`
}

/*taobao.fuwu.sale.link.gen*/
type FuwuSaleLinkGenRequest struct {
	values url.Values
}

func (this *FuwuSaleLinkGenRequest) GetApiMethodName() string {
	return "taobao.fuwu.sale.link.gen"
}
func (this *FuwuSaleLinkGenRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *FuwuSaleLinkGenRequest) GetValues() url.Values {
	return this.values
}

/* 用户需要营销的目标人群中的用户nick */
func (this *FuwuSaleLinkGenRequest) SetNick(value string) {
	this.Set("nick", value)
}

/* 从服务商后台，营销链接功能中生成的参数串直接复制使用。不要修改，否则抛错。 */
func (this *FuwuSaleLinkGenRequest) SetParamStr(value string) {
	this.Set("param_str", value)
}

type FuwuSaleLinkGenResponse struct {
	*ErrorResponse                `json:"error_response,omitempty"`
	FuwuSaleLinkGenResponseResult `json:"fuwu_sale_link_gen_response"`
}
type FuwuSaleLinkGenResponseResult struct {
	/* 通过营销链接接口生成的营销链接短地址 */
	Url string `json:"url"`
}

/*taobao.fuwu.scores.get*/
type FuwuScoresGetRequest struct {
	values url.Values
}

func (this *FuwuScoresGetRequest) GetApiMethodName() string {
	return "taobao.fuwu.scores.get"
}
func (this *FuwuScoresGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *FuwuScoresGetRequest) GetValues() url.Values {
	return this.values
}

/* 当前页 */
func (this *FuwuScoresGetRequest) SetCurrentPage(value string) {
	this.Set("current_page", value)
}

/* 评价日期，查询某一天的评价 */
func (this *FuwuScoresGetRequest) SetDate(value string) {
	this.Set("date", value)
}

/* 每页获取条数。默认值40，最小值1，最大值100。<br /> 支持最大值为：100<br /> 支持最小值为：1 */
func (this *FuwuScoresGetRequest) SetPageSize(value string) {
	this.Set("page_size", value)
}

type FuwuScoresGetResponse struct {
	*ErrorResponse              `json:"error_response,omitempty"`
	FuwuScoresGetResponseResult `json:"fuwu_scores_get_response"`
}
type FuwuScoresGetResponseResult struct {
	/* 评价流水记录 */
	ScoreResult struct {
		ScoreResult []*ScoreResult `json:"score_result"`
	} `json:"ScoreResult"`
}

/*taobao.vas.order.search*/
type VasOrderSearchRequest struct {
	values url.Values
}

func (this *VasOrderSearchRequest) GetApiMethodName() string {
	return "taobao.vas.order.search"
}
func (this *VasOrderSearchRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *VasOrderSearchRequest) GetValues() url.Values {
	return this.values
}

/* 应用收费代码，从合作伙伴后台（my.open.taobao.com）-收费管理-收费项目列表 能够获得该应用的收费代码 */
func (this *VasOrderSearchRequest) SetArticleCode(value string) {
	this.Set("article_code", value)
}

/* 订单号 */
func (this *VasOrderSearchRequest) SetBizOrderId(value string) {
	this.Set("biz_order_id", value)
}

/* 订单类型，1=新订 2=续订 3=升级 4=后台赠送 5=后台自动续订 6=订单审核后生成订购关系（暂时用不到） 空=全部 */
func (this *VasOrderSearchRequest) SetBizType(value string) {
	this.Set("biz_type", value)
}

/* 订单创建时间（订购时间）结束值 */
func (this *VasOrderSearchRequest) SetEndCreated(value string) {
	this.Set("end_created", value)
}

/* 收费项目代码，从合作伙伴后台（my.open.taobao.com）-收费管理-收费项目列表 能够获得收费项目代码 */
func (this *VasOrderSearchRequest) SetItemCode(value string) {
	this.Set("item_code", value)
}

/* 淘宝会员名 */
func (this *VasOrderSearchRequest) SetNick(value string) {
	this.Set("nick", value)
}

/* 子订单号 */
func (this *VasOrderSearchRequest) SetOrderId(value string) {
	this.Set("order_id", value)
}

/* 页码 */
func (this *VasOrderSearchRequest) SetPageNo(value string) {
	this.Set("page_no", value)
}

/* 一页包含的记录数<br /> 支持最大值为：200 */
func (this *VasOrderSearchRequest) SetPageSize(value string) {
	this.Set("page_size", value)
}

/* 订单创建时间（订购时间）起始值（当start_created和end_created都不填写时，默认返回最近90天的数据） */
func (this *VasOrderSearchRequest) SetStartCreated(value string) {
	this.Set("start_created", value)
}

type VasOrderSearchResponse struct {
	*ErrorResponse               `json:"error_response,omitempty"`
	VasOrderSearchResponseResult `json:"vas_order_search_response"`
}
type VasOrderSearchResponseResult struct {
	/* 商品订单对象 */
	ArticleBizOrders struct {
		ArticleBizOrder []*ArticleBizOrder `json:"article_biz_order"`
	} `json:"ArticleBizOrders"`
	/* 总记录数 */
	TotalItem int64 `json:"total_item"`
}

/*taobao.vas.subsc.search*/
type VasSubscSearchRequest struct {
	values url.Values
}

func (this *VasSubscSearchRequest) GetApiMethodName() string {
	return "taobao.vas.subsc.search"
}
func (this *VasSubscSearchRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *VasSubscSearchRequest) GetValues() url.Values {
	return this.values
}

/* 应用收费代码，从合作伙伴后台（my.open.taobao.com）-收费管理-收费项目列表 能够获得该应用的收费代码 */
func (this *VasSubscSearchRequest) SetArticleCode(value string) {
	this.Set("article_code", value)
}

/* 是否自动续费，true=自动续费 false=非自动续费 空=全部 */
func (this *VasSubscSearchRequest) SetAutosub(value string) {
	this.Set("autosub", value)
}

/* 到期时间结束值 */
func (this *VasSubscSearchRequest) SetEndDeadline(value string) {
	this.Set("end_deadline", value)
}

/* 是否到期提醒，true=到期提醒 false=非到期提醒 空=全部 */
func (this *VasSubscSearchRequest) SetExpireNotice(value string) {
	this.Set("expire_notice", value)
}

/* 收费项目代码，从合作伙伴后台（my.open.taobao.com）-收费管理-收费项目列表 能够获得收费项目代码 */
func (this *VasSubscSearchRequest) SetItemCode(value string) {
	this.Set("item_code", value)
}

/* 淘宝会员名 */
func (this *VasSubscSearchRequest) SetNick(value string) {
	this.Set("nick", value)
}

/* 页码 */
func (this *VasSubscSearchRequest) SetPageNo(value string) {
	this.Set("page_no", value)
}

/* 一页包含的记录数<br /> 支持最大值为：200 */
func (this *VasSubscSearchRequest) SetPageSize(value string) {
	this.Set("page_size", value)
}

/* 到期时间起始值（当start_deadline和end_deadline都不填写时，默认返回最近90天的数据） */
func (this *VasSubscSearchRequest) SetStartDeadline(value string) {
	this.Set("start_deadline", value)
}

/* 订购记录状态，1=有效 2=过期 空=全部 */
func (this *VasSubscSearchRequest) SetStatus(value string) {
	this.Set("status", value)
}

type VasSubscSearchResponse struct {
	*ErrorResponse               `json:"error_response,omitempty"`
	VasSubscSearchResponseResult `json:"vas_subsc_search_response"`
}
type VasSubscSearchResponseResult struct {
	/* 订购关系对象 */
	ArticleSubs struct {
		ArticleSub []*ArticleSub `json:"article_sub"`
	} `json:"ArticleSubs"`
	/* 总记录数 */
	TotalItem int64 `json:"total_item"`
}

/*taobao.vas.subscribe.get*/
type VasSubscribeGetRequest struct {
	values url.Values
}

func (this *VasSubscribeGetRequest) GetApiMethodName() string {
	return "taobao.vas.subscribe.get"
}
func (this *VasSubscribeGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *VasSubscribeGetRequest) GetValues() url.Values {
	return this.values
}

/* 商品编码，从合作伙伴后台（my.open.taobao.com）-收费管理-收费项目列表 能够获得该应用的商品代码 */
func (this *VasSubscribeGetRequest) SetArticleCode(value string) {
	this.Set("article_code", value)
}

/* 淘宝会员名 */
func (this *VasSubscribeGetRequest) SetNick(value string) {
	this.Set("nick", value)
}

type VasSubscribeGetResponse struct {
	*ErrorResponse                `json:"error_response,omitempty"`
	VasSubscribeGetResponseResult `json:"vas_subscribe_get_response"`
}
type VasSubscribeGetResponseResult struct {
	/* 用户订购信息 */
	ArticleUserSubscribes struct {
		ArticleUserSubscribe []*ArticleUserSubscribe `json:"article_user_subscribe"`
	} `json:"ArticleUserSubscribes"`
}

/*taobao.refund.get*/
type RefundGetRequest struct {
	values url.Values
}

func (this *RefundGetRequest) GetApiMethodName() string {
	return "taobao.refund.get"
}
func (this *RefundGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *RefundGetRequest) GetValues() url.Values {
	return this.values
}

/* 需要返回的字段。目前支持有：refund_id, alipay_no, tid, oid, buyer_nick, seller_nick, total_fee, status, created, refund_fee, good_status, has_good_return, payment, reason, desc, num_iid, title, price, num, good_return_time, company_name, sid, address, shipping_type, refund_remind_timeout, refund_phase, refund_version, operation_contraint, attribute, outer_id, sku */
func (this *RefundGetRequest) SetFields(value string) {
	this.Set("fields", value)
}

/* 退款单号<br /> 支持最大值为：9223372036854775807<br /> 支持最小值为：1 */
func (this *RefundGetRequest) SetRefundId(value string) {
	this.Set("refund_id", value)
}

type RefundGetResponse struct {
	*ErrorResponse          `json:"error_response,omitempty"`
	RefundGetResponseResult `json:"refund_get_response"`
}
type RefundGetResponseResult struct {
	/* 退款详情 */
	Refund *Refund `json:"refund"`
}

/*taobao.refund.message.add*/
type RefundMessageAddRequest struct {
	values url.Values
}

func (this *RefundMessageAddRequest) GetApiMethodName() string {
	return "taobao.refund.message.add"
}
func (this *RefundMessageAddRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *RefundMessageAddRequest) GetValues() url.Values {
	return this.values
}

/* 留言内容。最大长度: 400个字节 */
func (this *RefundMessageAddRequest) SetContent(value string) {
	this.Set("content", value)
}

/* 图片（凭证）。类型: JPG,GIF,PNG;最大为: 500K */
func (this *RefundMessageAddRequest) SetImage(value string) {
	this.Set("image", value)
}

/* 退款编号。 */
func (this *RefundMessageAddRequest) SetRefundId(value string) {
	this.Set("refund_id", value)
}

type RefundMessageAddResponse struct {
	*ErrorResponse                 `json:"error_response,omitempty"`
	RefundMessageAddResponseResult `json:"refund_message_add_response"`
}
type RefundMessageAddResponseResult struct {
	/* 退款信息。包含id和created */
	RefundMessage *RefundMessage `json:"refund_message"`
}

/*taobao.refund.messages.get*/
type RefundMessagesGetRequest struct {
	values url.Values
}

func (this *RefundMessagesGetRequest) GetApiMethodName() string {
	return "taobao.refund.messages.get"
}
func (this *RefundMessagesGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *RefundMessagesGetRequest) GetValues() url.Values {
	return this.values
}

/* 需返回的字段列表。可选值：RefundMessage结构体中的所有字段，以半角逗号(,)分隔。 */
func (this *RefundMessagesGetRequest) SetFields(value string) {
	this.Set("fields", value)
}

/* 页码<br /> 支持最小值为：1 */
func (this *RefundMessagesGetRequest) SetPageNo(value string) {
	this.Set("page_no", value)
}

/* 每页条数<br /> 支持最大值为：100<br /> 支持最小值为：1 */
func (this *RefundMessagesGetRequest) SetPageSize(value string) {
	this.Set("page_size", value)
}

/* 退款单号<br /> 支持最小值为：1 */
func (this *RefundMessagesGetRequest) SetRefundId(value string) {
	this.Set("refund_id", value)
}

/* 退款阶段，可选值：onsale（售中），aftersale（售后），天猫退款为必传。 */
func (this *RefundMessagesGetRequest) SetRefundPhase(value string) {
	this.Set("refund_phase", value)
}

type RefundMessagesGetResponse struct {
	*ErrorResponse                  `json:"error_response,omitempty"`
	RefundMessagesGetResponseResult `json:"refund_messages_get_response"`
}
type RefundMessagesGetResponseResult struct {
	/* 查询到的退款留言/凭证列表 */
	RefundMessages struct {
		RefundMessage []*RefundMessage `json:"refund_message"`
	} `json:"RefundMessages"`
	/* 查询到的退款留言/凭证总数 */
	TotalResults int64 `json:"total_results"`
}

/*taobao.refund.refuse*/
type RefundRefuseRequest struct {
	values url.Values
}

func (this *RefundRefuseRequest) GetApiMethodName() string {
	return "taobao.refund.refuse"
}
func (this *RefundRefuseRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *RefundRefuseRequest) GetValues() url.Values {
	return this.values
}

/* 退款单号 */
func (this *RefundRefuseRequest) SetRefundId(value string) {
	this.Set("refund_id", value)
}

/* 可选值为：售中：onsale，售后：aftersale，天猫退款为必填项。 */
func (this *RefundRefuseRequest) SetRefundPhase(value string) {
	this.Set("refund_phase", value)
}

/* 退款版本号，天猫退款为必填项。 */
func (this *RefundRefuseRequest) SetRefundVersion(value string) {
	this.Set("refund_version", value)
}

/* 拒绝退款时的说明信息，长度2-200<br /> 支持最大长度为：200<br /> 支持的最大列表长度为：200 */
func (this *RefundRefuseRequest) SetRefuseMessage(value string) {
	this.Set("refuse_message", value)
}

/* 拒绝退款时的退款凭证，一般是卖家拒绝退款时使用的发货凭证，最大长度130000字节，支持的图片格式：GIF, JPG, PNG。天猫退款为必填项。<br /> 支持的文件类型为：gif,jpg,png<br /> 支持的最大列表长度为：130000 */
func (this *RefundRefuseRequest) SetRefuseProof(value string) {
	this.Set("refuse_proof", value)
}

type RefundRefuseResponse struct {
	*ErrorResponse             `json:"error_response,omitempty"`
	RefundRefuseResponseResult `json:"refund_refuse_response"`
}
type RefundRefuseResponseResult struct {
	/* 拒绝退款操作是否成功 */
	IsSuccess bool `json:"is_success"`
	/* 拒绝退款成功后，会返回Refund数据结构中的refund_id, status, modified字段 */
	Refund *Refund `json:"refund"`
}

/*taobao.refunds.apply.get*/
type RefundsApplyGetRequest struct {
	values url.Values
}

func (this *RefundsApplyGetRequest) GetApiMethodName() string {
	return "taobao.refunds.apply.get"
}
func (this *RefundsApplyGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *RefundsApplyGetRequest) GetValues() url.Values {
	return this.values
}

/* 需要返回的字段。目前支持有：refund_id, tid, title, buyer_nick, seller_nick, total_fee, status, created, refund_fee */
func (this *RefundsApplyGetRequest) SetFields(value string) {
	this.Set("fields", value)
}

/* 页码。传入值为 1 代表第一页，传入值为 2 代表第二页，依此类推。默认返回的数据是从第一页开始<br /> 支持最小值为：1 */
func (this *RefundsApplyGetRequest) SetPageNo(value string) {
	this.Set("page_no", value)
}

/* 每页条数。取值范围:大于零的整数; 默认值:40;最大值:100<br /> 支持最大值为：100<br /> 支持最小值为：1 */
func (this *RefundsApplyGetRequest) SetPageSize(value string) {
	this.Set("page_size", value)
}

/* 卖家昵称 */
func (this *RefundsApplyGetRequest) SetSellerNick(value string) {
	this.Set("seller_nick", value)
}

/* 退款状态，默认查询所有退款状态的数据，除了默认值外每次只能查询一种状态。
WAIT_SELLER_AGREE(买家已经申请退款，等待卖家同意)
WAIT_BUYER_RETURN_GOODS(卖家已经同意退款，等待买家退货)
WAIT_SELLER_CONFIRM_GOODS(买家已经退货，等待卖家确认收货)
SELLER_REFUSE_BUYER(卖家拒绝退款)
CLOSED(退款关闭)
SUCCESS(退款成功) */
func (this *RefundsApplyGetRequest) SetStatus(value string) {
	this.Set("status", value)
}

/* 交易类型列表，一次查询多种类型可用半角逗号分隔，默认同时查询guarantee_trade, auto_delivery的2种类型的数据。
fixed(一口价)
auction(拍卖)
guarantee_trade(一口价、拍卖)
independent_simple_trade(旺店入门版交易)
independent_shop_trade(旺店标准版交易)
auto_delivery(自动发货)
ec(直冲)
cod(货到付款)
fenxiao(分销)
game_equipment(游戏装备)
shopex_trade(ShopEX交易)
netcn_trade(万网交易)
external_trade(统一外部交易) */
func (this *RefundsApplyGetRequest) SetType(value string) {
	this.Set("type", value)
}

type RefundsApplyGetResponse struct {
	*ErrorResponse                `json:"error_response,omitempty"`
	RefundsApplyGetResponseResult `json:"refunds_apply_get_response"`
}
type RefundsApplyGetResponseResult struct {
	/* 搜索到的退款信息列表 */
	Refunds struct {
		Refund []*Refund `json:"refund"`
	} `json:"Refunds"`
	/* 搜索到的交易信息总数 */
	TotalResults int64 `json:"total_results"`
}

/*taobao.refunds.receive.get*/
type RefundsReceiveGetRequest struct {
	values url.Values
}

func (this *RefundsReceiveGetRequest) GetApiMethodName() string {
	return "taobao.refunds.receive.get"
}
func (this *RefundsReceiveGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *RefundsReceiveGetRequest) GetValues() url.Values {
	return this.values
}

/* 买家昵称 */
func (this *RefundsReceiveGetRequest) SetBuyerNick(value string) {
	this.Set("buyer_nick", value)
}

/* 查询修改时间结束。格式: yyyy-MM-dd HH:mm:ss */
func (this *RefundsReceiveGetRequest) SetEndModified(value string) {
	this.Set("end_modified", value)
}

/* 需要返回的字段。目前支持有：refund_id, tid, title, buyer_nick, seller_nick, total_fee, status, created, refund_fee, oid, good_status, company_name, sid, payment, reason, desc, has_good_return, modified, order_status,refund_phase */
func (this *RefundsReceiveGetRequest) SetFields(value string) {
	this.Set("fields", value)
}

/* 页码。取值范围:大于零的整数; 默认值:1<br /> 支持最小值为：1 */
func (this *RefundsReceiveGetRequest) SetPageNo(value string) {
	this.Set("page_no", value)
}

/* 每页条数。取值范围:大于零的整数; 默认值:40;最大值:100<br /> 支持最大值为：100<br /> 支持最小值为：1 */
func (this *RefundsReceiveGetRequest) SetPageSize(value string) {
	this.Set("page_size", value)
}

/* 查询修改时间开始。格式: yyyy-MM-dd HH:mm:ss */
func (this *RefundsReceiveGetRequest) SetStartModified(value string) {
	this.Set("start_modified", value)
}

/* 退款状态，默认查询所有退款状态的数据，除了默认值外每次只能查询一种状态。WAIT_SELLER_AGREE(买家已经申请退款，等待卖家同意) WAIT_BUYER_RETURN_GOODS(卖家已经同意退款，等待买家退货) WAIT_SELLER_CONFIRM_GOODS(买家已经退货，等待卖家确认收货) SELLER_REFUSE_BUYER(卖家拒绝退款) CLOSED(退款关闭) SUCCESS(退款成功) */
func (this *RefundsReceiveGetRequest) SetStatus(value string) {
	this.Set("status", value)
}

/* 交易类型列表，一次查询多种类型可用半角逗号分隔，默认同时查询guarantee_trade, auto_delivery的2种类型的数据。fixed(一口价) auction(拍卖) guarantee_trade(一口价、拍卖) independent_simple_trade(旺店入门版交易) independent_shop_trade(旺店标准版交易) auto_delivery(自动发货) ec(直冲) cod(货到付款) fenxiao(分销) game_equipment(游戏装备) shopex_trade(ShopEX交易) netcn_trade(万网交易) external_trade(统一外部交易) */
func (this *RefundsReceiveGetRequest) SetType(value string) {
	this.Set("type", value)
}

/* 是否启用has_next的分页方式，如果指定true,则返回的结果中不包含总记录数，但是会新增一个是否存在下一页的的字段，通过此种方式获取增量退款，接口调用成功率在原有的基础上有所提升。 */
func (this *RefundsReceiveGetRequest) SetUseHasNext(value string) {
	this.Set("use_has_next", value)
}

type RefundsReceiveGetResponse struct {
	*ErrorResponse                  `json:"error_response,omitempty"`
	RefundsReceiveGetResponseResult `json:"refunds_receive_get_response"`
}
type RefundsReceiveGetResponseResult struct {
	/* 是否存在下一页 */
	HasNext bool `json:"has_next"`
	/* 搜索到的退款信息列表 */
	Refunds struct {
		Refund []*Refund `json:"refund"`
	} `json:"Refunds"`
	/* 搜索到的退款信息总数 */
	TotalResults int64 `json:"total_results"`
}

/*taobao.rp.refund.review*/
type RpRefundReviewRequest struct {
	values url.Values
}

func (this *RpRefundReviewRequest) GetApiMethodName() string {
	return "taobao.rp.refund.review"
}
func (this *RpRefundReviewRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *RpRefundReviewRequest) GetValues() url.Values {
	return this.values
}

/* 审核留言 */
func (this *RpRefundReviewRequest) SetMessage(value string) {
	this.Set("message", value)
}

/* 审核人姓名 */
func (this *RpRefundReviewRequest) SetOperator(value string) {
	this.Set("operator", value)
}

/* 退款单编号 */
func (this *RpRefundReviewRequest) SetRefundId(value string) {
	this.Set("refund_id", value)
}

/* 退款阶段，可选值：售中：onsale，售后：aftersale */
func (this *RpRefundReviewRequest) SetRefundPhase(value string) {
	this.Set("refund_phase", value)
}

/* 退款最后更新时间，以时间戳的方式表示 */
func (this *RpRefundReviewRequest) SetRefundVersion(value string) {
	this.Set("refund_version", value)
}

/* 审核是否可用于批量退款，可选值：true（审核通过），false（审核不通过或反审核） */
func (this *RpRefundReviewRequest) SetResult(value string) {
	this.Set("result", value)
}

type RpRefundReviewResponse struct {
	*ErrorResponse               `json:"error_response,omitempty"`
	RpRefundReviewResponseResult `json:"rp_refund_review_response"`
}
type RpRefundReviewResponseResult struct {
	/* 审核操作是否成功 */
	IsSuccess bool `json:"is_success"`
}

/*taobao.rp.refunds.agree*/
type RpRefundsAgreeRequest struct {
	values url.Values
}

func (this *RpRefundsAgreeRequest) GetApiMethodName() string {
	return "taobao.rp.refunds.agree"
}
func (this *RpRefundsAgreeRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *RpRefundsAgreeRequest) GetValues() url.Values {
	return this.values
}

/* 短信验证码，如果退款金额达到一定的数量，后端会返回调用失败，并同时往卖家的手机发送一条短信验证码。接下来用收到的短信验证码再次发起API调用即可完成退款操作。 */
func (this *RpRefundsAgreeRequest) SetCode(value string) {
	this.Set("code", value)
}

/* 退款信息，格式：refund_id|amount|version|phase，其中refund_id为退款编号，amount为退款金额（以分为单位），version为退款最后更新时间（时间戳格式），phase为退款阶段（可选值为：onsale, aftersale，天猫退款必值，淘宝退款不需要传），多个退款以半角逗号分隔。 */
func (this *RpRefundsAgreeRequest) SetRefundInfos(value string) {
	this.Set("refund_infos", value)
}

type RpRefundsAgreeResponse struct {
	*ErrorResponse               `json:"error_response,omitempty"`
	RpRefundsAgreeResponseResult `json:"rp_refunds_agree_response"`
}
type RpRefundsAgreeResponseResult struct {
	/* 信息 */
	Message string `json:"message"`
	/* 批量退款操作情况，可选值：OP_SUCC（全部成功），SOME_OP_SUCC（部分成功），OP_FAILURE_UE（全部失败） */
	MsgCode string `json:"msg_code"`
	/* 退款操作结果列表 */
	Results struct {
		RefundMappingResult []*RefundMappingResult `json:"refund_mapping_result"`
	} `json:"Results"`
	/* 操作成功 */
	Succ bool `json:"succ"`
}

/*taobao.rp.returngoods.agree*/
type RpReturngoodsAgreeRequest struct {
	values url.Values
}

func (this *RpReturngoodsAgreeRequest) GetApiMethodName() string {
	return "taobao.rp.returngoods.agree"
}
func (this *RpReturngoodsAgreeRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *RpReturngoodsAgreeRequest) GetValues() url.Values {
	return this.values
}

/* 卖家提供的退货地址，淘宝退款为必填项。 */
func (this *RpReturngoodsAgreeRequest) SetAddress(value string) {
	this.Set("address", value)
}

/* 卖家手机，淘宝退款为必填项。 */
func (this *RpReturngoodsAgreeRequest) SetMobile(value string) {
	this.Set("mobile", value)
}

/* 卖家姓名，淘宝退款为必填项。 */
func (this *RpReturngoodsAgreeRequest) SetName(value string) {
	this.Set("name", value)
}

/* 卖家提供的退货地址的邮编，淘宝退款为必填项。 */
func (this *RpReturngoodsAgreeRequest) SetPost(value string) {
	this.Set("post", value)
}

/* 退款编号 */
func (this *RpReturngoodsAgreeRequest) SetRefundId(value string) {
	this.Set("refund_id", value)
}

/* 售中：onsale，售后：aftersale，天猫退款为必填项。 */
func (this *RpReturngoodsAgreeRequest) SetRefundPhase(value string) {
	this.Set("refund_phase", value)
}

/* 退款版本号，天猫退款为必填项。 */
func (this *RpReturngoodsAgreeRequest) SetRefundVersion(value string) {
	this.Set("refund_version", value)
}

/* 卖家退货留言，天猫退款为必填项。 */
func (this *RpReturngoodsAgreeRequest) SetRemark(value string) {
	this.Set("remark", value)
}

/* 卖家收货地址编号，天猫退款为必填项。 */
func (this *RpReturngoodsAgreeRequest) SetSellerAddressId(value string) {
	this.Set("seller_address_id", value)
}

/* 卖家座机，淘宝退款为必填项。 */
func (this *RpReturngoodsAgreeRequest) SetTel(value string) {
	this.Set("tel", value)
}

type RpReturngoodsAgreeResponse struct {
	*ErrorResponse                   `json:"error_response,omitempty"`
	RpReturngoodsAgreeResponseResult `json:"rp_returngoods_agree_response"`
}
type RpReturngoodsAgreeResponseResult struct {
	/* 操作成功 */
	IsSuccess bool `json:"is_success"`
}

/*taobao.rp.returngoods.refill*/
type RpReturngoodsRefillRequest struct {
	values url.Values
}

func (this *RpReturngoodsRefillRequest) GetApiMethodName() string {
	return "taobao.rp.returngoods.refill"
}
func (this *RpReturngoodsRefillRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *RpReturngoodsRefillRequest) GetValues() url.Values {
	return this.values
}

/* 物流公司编号 */
func (this *RpReturngoodsRefillRequest) SetLogisticsCompanyCode(value string) {
	this.Set("logistics_company_code", value)
}

/* 物流公司运单号 */
func (this *RpReturngoodsRefillRequest) SetLogisticsWaybillNo(value string) {
	this.Set("logistics_waybill_no", value)
}

/* 退款单编号 */
func (this *RpReturngoodsRefillRequest) SetRefundId(value string) {
	this.Set("refund_id", value)
}

/* 退款阶段，可选值：售中：onsale，售后：aftersale */
func (this *RpReturngoodsRefillRequest) SetRefundPhase(value string) {
	this.Set("refund_phase", value)
}

type RpReturngoodsRefillResponse struct {
	*ErrorResponse                    `json:"error_response,omitempty"`
	RpReturngoodsRefillResponseResult `json:"rp_returngoods_refill_response"`
}
type RpReturngoodsRefillResponseResult struct {
	/* 验货操作是否成功 */
	IsSuccess bool `json:"is_success"`
}

/*taobao.rp.returngoods.refuse*/
type RpReturngoodsRefuseRequest struct {
	values url.Values
}

func (this *RpReturngoodsRefuseRequest) GetApiMethodName() string {
	return "taobao.rp.returngoods.refuse"
}
func (this *RpReturngoodsRefuseRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *RpReturngoodsRefuseRequest) GetValues() url.Values {
	return this.values
}

/* 退款编号 */
func (this *RpReturngoodsRefuseRequest) SetRefundId(value string) {
	this.Set("refund_id", value)
}

/* 退款服务状态，售后或者售中 */
func (this *RpReturngoodsRefuseRequest) SetRefundPhase(value string) {
	this.Set("refund_phase", value)
}

/* 退款版本号 */
func (this *RpReturngoodsRefuseRequest) SetRefundVersion(value string) {
	this.Set("refund_version", value)
}

/* 拒绝退货凭证图片，必须图片格式，大小不能超过5M */
func (this *RpReturngoodsRefuseRequest) SetRefuseProof(value string) {
	this.Set("refuse_proof", value)
}

type RpReturngoodsRefuseResponse struct {
	*ErrorResponse                    `json:"error_response,omitempty"`
	RpReturngoodsRefuseResponseResult `json:"rp_returngoods_refuse_response"`
}
type RpReturngoodsRefuseResponseResult struct {
	/* asdf */
	Result bool `json:"result"`
}

/*taobao.categoryrecommend.items.get*/
type CategoryrecommendItemsGetRequest struct {
	values url.Values
}

func (this *CategoryrecommendItemsGetRequest) GetApiMethodName() string {
	return "taobao.categoryrecommend.items.get"
}
func (this *CategoryrecommendItemsGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *CategoryrecommendItemsGetRequest) GetValues() url.Values {
	return this.values
}

/* 传入叶子类目ID */
func (this *CategoryrecommendItemsGetRequest) SetCategoryId(value string) {
	this.Set("category_id", value)
}

/* 请求个数，建议获取20个 */
func (this *CategoryrecommendItemsGetRequest) SetCount(value string) {
	this.Set("count", value)
}

/* 额外参数 */
func (this *CategoryrecommendItemsGetRequest) SetExt(value string) {
	this.Set("ext", value)
}

/* 请求类型，1：类目下热门商品推荐。其他值当非法值处理 */
func (this *CategoryrecommendItemsGetRequest) SetRecommendType(value string) {
	this.Set("recommend_type", value)
}

type CategoryrecommendItemsGetResponse struct {
	*ErrorResponse                          `json:"error_response,omitempty"`
	CategoryrecommendItemsGetResponseResult `json:"categoryrecommend_items_get_response"`
}
type CategoryrecommendItemsGetResponseResult struct {
	/* 返回关联的商品集合 */
	FavoriteItems struct {
		FavoriteItem []*FavoriteItem `json:"favorite_item"`
	} `json:"FavoriteItems"`
}

/*taobao.itemrecommend.items.get*/
type ItemrecommendItemsGetRequest struct {
	values url.Values
}

func (this *ItemrecommendItemsGetRequest) GetApiMethodName() string {
	return "taobao.itemrecommend.items.get"
}
func (this *ItemrecommendItemsGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *ItemrecommendItemsGetRequest) GetValues() url.Values {
	return this.values
}

/* 请求返回宝贝的个数，建议取20个 */
func (this *ItemrecommendItemsGetRequest) SetCount(value string) {
	this.Set("count", value)
}

/* 额外的参数信息 */
func (this *ItemrecommendItemsGetRequest) SetExt(value string) {
	this.Set("ext", value)
}

/* 商品ID */
func (this *ItemrecommendItemsGetRequest) SetItemId(value string) {
	this.Set("item_id", value)
}

/* 查询类型标识符，可传入1-3，1：同类商品推荐，2：异类商品推荐， 3：同店商品推荐。其他值当非法值处理 */
func (this *ItemrecommendItemsGetRequest) SetRecommendType(value string) {
	this.Set("recommend_type", value)
}

type ItemrecommendItemsGetResponse struct {
	*ErrorResponse                      `json:"error_response,omitempty"`
	ItemrecommendItemsGetResponseResult `json:"itemrecommend_items_get_response"`
}
type ItemrecommendItemsGetResponseResult struct {
	/* 返回的推荐商品列表结果集 */
	Values struct {
		FavoriteItem []*FavoriteItem `json:"favorite_item"`
	} `json:"Values"`
}

/*taobao.shoprecommend.items.get*/
type ShoprecommendItemsGetRequest struct {
	values url.Values
}

func (this *ShoprecommendItemsGetRequest) GetApiMethodName() string {
	return "taobao.shoprecommend.items.get"
}
func (this *ShoprecommendItemsGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *ShoprecommendItemsGetRequest) GetValues() url.Values {
	return this.values
}

/* 请求个数，最大只能获取10个 */
func (this *ShoprecommendItemsGetRequest) SetCount(value string) {
	this.Set("count", value)
}

/* 额外参数 */
func (this *ShoprecommendItemsGetRequest) SetExt(value string) {
	this.Set("ext", value)
}

/* 请求类型，1：店内热门商品推荐。其他值当非法值处理 */
func (this *ShoprecommendItemsGetRequest) SetRecommendType(value string) {
	this.Set("recommend_type", value)
}

/* <p>传入卖家ID。这里的seller_id得通过<a href="http://api.taobao.com/apidoc/api.htm?path=cid:38-apiId:10449">taobao.taobaoke.shops.get</a>
跟<a href="http://api.taobao.com/apidoc/api.htm?path=cid:38-apiId:21419">taobao.taobaoke.widget.shops.convert</a>这两个接口去获取user_id字段。</p>
<p>如果是非淘客卖家，则无法获取，暂无替代方案。</p> */
func (this *ShoprecommendItemsGetRequest) SetSellerId(value string) {
	this.Set("seller_id", value)
}

type ShoprecommendItemsGetResponse struct {
	*ErrorResponse                      `json:"error_response,omitempty"`
	ShoprecommendItemsGetResponseResult `json:"shoprecommend_items_get_response"`
}
type ShoprecommendItemsGetResponseResult struct {
	/* 返回与店铺关联的宝贝集合 */
	FavoriteItems struct {
		FavoriteItem []*FavoriteItem `json:"favorite_item"`
	} `json:"FavoriteItems"`
}

/*taobao.shoprecommend.shops.get*/
type ShoprecommendShopsGetRequest struct {
	values url.Values
}

func (this *ShoprecommendShopsGetRequest) GetApiMethodName() string {
	return "taobao.shoprecommend.shops.get"
}
func (this *ShoprecommendShopsGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *ShoprecommendShopsGetRequest) GetValues() url.Values {
	return this.values
}

/* 请求个数，建议获取16个 */
func (this *ShoprecommendShopsGetRequest) SetCount(value string) {
	this.Set("count", value)
}

/* 额外参数 */
func (this *ShoprecommendShopsGetRequest) SetExt(value string) {
	this.Set("ext", value)
}

/* 请求类型，1：关联店铺推荐。其他值当非法值处理 */
func (this *ShoprecommendShopsGetRequest) SetRecommendType(value string) {
	this.Set("recommend_type", value)
}

/* 传入卖家ID */
func (this *ShoprecommendShopsGetRequest) SetSellerId(value string) {
	this.Set("seller_id", value)
}

type ShoprecommendShopsGetResponse struct {
	*ErrorResponse                      `json:"error_response,omitempty"`
	ShoprecommendShopsGetResponseResult `json:"shoprecommend_shops_get_response"`
}
type ShoprecommendShopsGetResponseResult struct {
	/* 返回与店铺关联的店铺集 */
	FavoriteShops struct {
		FavoriteShop []*FavoriteShop `json:"favorite_shop"`
	} `json:"FavoriteShops"`
}

/*taobao.userrecommend.items.get*/
type UserrecommendItemsGetRequest struct {
	values url.Values
}

func (this *UserrecommendItemsGetRequest) GetApiMethodName() string {
	return "taobao.userrecommend.items.get"
}
func (this *UserrecommendItemsGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *UserrecommendItemsGetRequest) GetValues() url.Values {
	return this.values
}

/* 请求个数，建议取20个 */
func (this *UserrecommendItemsGetRequest) SetCount(value string) {
	this.Set("count", value)
}

/* 额外参数 */
func (this *UserrecommendItemsGetRequest) SetExt(value string) {
	this.Set("ext", value)
}

/* 请求类型，1：用户购买意图。其他值当非法值处理 */
func (this *UserrecommendItemsGetRequest) SetRecommendType(value string) {
	this.Set("recommend_type", value)
}

type UserrecommendItemsGetResponse struct {
	*ErrorResponse                      `json:"error_response,omitempty"`
	UserrecommendItemsGetResponseResult `json:"userrecommend_items_get_response"`
}
type UserrecommendItemsGetResponseResult struct {
	/* 返回用户相关的关联宝贝集合 */
	FavoriteItems struct {
		FavoriteItem []*FavoriteItem `json:"favorite_item"`
	} `json:"FavoriteItems"`
}

/*taobao.jds.hluser.get*/
type JdsHluserGetRequest struct {
	values url.Values
}

func (this *JdsHluserGetRequest) GetApiMethodName() string {
	return "taobao.jds.hluser.get"
}
func (this *JdsHluserGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *JdsHluserGetRequest) GetValues() url.Values {
	return this.values
}

type JdsHluserGetResponse struct {
	*ErrorResponse             `json:"error_response,omitempty"`
	JdsHluserGetResponseResult `json:"jds_hluser_get_response"`
}
type JdsHluserGetResponseResult struct {
	/* 回流用户信息 */
	HlUser *HlUserDO `json:"hl_user"`
}

/*taobao.jds.hluser.update*/
type JdsHluserUpdateRequest struct {
	values url.Values
}

func (this *JdsHluserUpdateRequest) GetApiMethodName() string {
	return "taobao.jds.hluser.update"
}
func (this *JdsHluserUpdateRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *JdsHluserUpdateRequest) GetValues() url.Values {
	return this.values
}

/* 回流信息是否开通买家端展示 */
func (this *JdsHluserUpdateRequest) SetOpenForBuyer(value string) {
	this.Set("open_for_buyer", value)
}

type JdsHluserUpdateResponse struct {
	*ErrorResponse                `json:"error_response,omitempty"`
	JdsHluserUpdateResponseResult `json:"jds_hluser_update_response"`
}
type JdsHluserUpdateResponseResult struct {
	/* 是否成功 */
	Result bool `json:"result"`
}

/*taobao.jds.refund.traces.get*/
type JdsRefundTracesGetRequest struct {
	values url.Values
}

func (this *JdsRefundTracesGetRequest) GetApiMethodName() string {
	return "taobao.jds.refund.traces.get"
}
func (this *JdsRefundTracesGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *JdsRefundTracesGetRequest) GetValues() url.Values {
	return this.values
}

/* 淘宝的退款编号<br /> 支持最小值为：1 */
func (this *JdsRefundTracesGetRequest) SetRefundId(value string) {
	this.Set("refund_id", value)
}

type JdsRefundTracesGetResponse struct {
	*ErrorResponse                   `json:"error_response,omitempty"`
	JdsRefundTracesGetResponseResult `json:"jds_refund_traces_get_response"`
}
type JdsRefundTracesGetResponseResult struct {
	/* 退款跟踪列表 */
	Traces struct {
		RefundTrace []*RefundTrace `json:"refund_trace"`
	} `json:"Traces"`
}

/*taobao.jds.trade.traces.get*/
type JdsTradeTracesGetRequest struct {
	values url.Values
}

func (this *JdsTradeTracesGetRequest) GetApiMethodName() string {
	return "taobao.jds.trade.traces.get"
}
func (this *JdsTradeTracesGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *JdsTradeTracesGetRequest) GetValues() url.Values {
	return this.values
}

/* 淘宝的订单tid */
func (this *JdsTradeTracesGetRequest) SetTid(value string) {
	this.Set("tid", value)
}

type JdsTradeTracesGetResponse struct {
	*ErrorResponse                  `json:"error_response,omitempty"`
	JdsTradeTracesGetResponseResult `json:"jds_trade_traces_get_response"`
}
type JdsTradeTracesGetResponseResult struct {
	/* 跟踪信息列表 */
	Traces struct {
		TradeTrace []*TradeTrace `json:"trade_trace"`
	} `json:"Traces"`
}

/*taobao.jds.trades.statistics.diff*/
type JdsTradesStatisticsDiffRequest struct {
	values url.Values
}

func (this *JdsTradesStatisticsDiffRequest) GetApiMethodName() string {
	return "taobao.jds.trades.statistics.diff"
}
func (this *JdsTradesStatisticsDiffRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *JdsTradesStatisticsDiffRequest) GetValues() url.Values {
	return this.values
}

/* 查询的日期，格式如YYYYMMDD的日期对应的数字 */
func (this *JdsTradesStatisticsDiffRequest) SetDate(value string) {
	this.Set("date", value)
}

/* 页数 */
func (this *JdsTradesStatisticsDiffRequest) SetPageNo(value string) {
	this.Set("page_no", value)
}

/* 需要比较的状态 */
func (this *JdsTradesStatisticsDiffRequest) SetPostStatus(value string) {
	this.Set("post_status", value)
}

/* 需要比较的状态，将会和post_status做比较 */
func (this *JdsTradesStatisticsDiffRequest) SetPreStatus(value string) {
	this.Set("pre_status", value)
}

type JdsTradesStatisticsDiffResponse struct {
	*ErrorResponse                        `json:"error_response,omitempty"`
	JdsTradesStatisticsDiffResponseResult `json:"jds_trades_statistics_diff_response"`
}
type JdsTradesStatisticsDiffResponseResult struct {
	/* pre_status比post_status多的tid列表 */
	Tids []int64 `json:"tids"`
	/* 总记录数 */
	TotalResults int64 `json:"total_results"`
}

/*taobao.jds.trades.statistics.get*/
type JdsTradesStatisticsGetRequest struct {
	values url.Values
}

func (this *JdsTradesStatisticsGetRequest) GetApiMethodName() string {
	return "taobao.jds.trades.statistics.get"
}
func (this *JdsTradesStatisticsGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *JdsTradesStatisticsGetRequest) GetValues() url.Values {
	return this.values
}

/* 查询的日期，格式如YYYYMMDD的日期对应的数字 */
func (this *JdsTradesStatisticsGetRequest) SetDate(value string) {
	this.Set("date", value)
}

type JdsTradesStatisticsGetResponse struct {
	*ErrorResponse                       `json:"error_response,omitempty"`
	JdsTradesStatisticsGetResponseResult `json:"jds_trades_statistics_get_response"`
}
type JdsTradesStatisticsGetResponseResult struct {
	/* 订单状态计数值 */
	StatusInfos struct {
		TradeStat []*TradeStat `json:"trade_stat"`
	} `json:"StatusInfos"`
}

/*taobao.jushita.jdp.rdsinfo.get*/
type JushitaJdpRdsinfoGetRequest struct {
	values url.Values
}

func (this *JushitaJdpRdsinfoGetRequest) GetApiMethodName() string {
	return "taobao.jushita.jdp.rdsinfo.get"
}
func (this *JushitaJdpRdsinfoGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *JushitaJdpRdsinfoGetRequest) GetValues() url.Values {
	return this.values
}

/* 查询的appkey */
func (this *JushitaJdpRdsinfoGetRequest) SetTargetAppkey(value string) {
	this.Set("target_appkey", value)
}

type JushitaJdpRdsinfoGetResponse struct {
	*ErrorResponse                     `json:"error_response,omitempty"`
	JushitaJdpRdsinfoGetResponseResult `json:"jushita_jdp_rdsinfo_get_response"`
}
type JushitaJdpRdsinfoGetResponseResult struct {
	/* app的rds列表 */
	RdsList struct {
		RdsConfig []*RdsConfig `json:"rds_config"`
	} `json:"RdsList"`
}

/*taobao.jushita.jdp.task.update*/
type JushitaJdpTaskUpdateRequest struct {
	values url.Values
}

func (this *JushitaJdpTaskUpdateRequest) GetApiMethodName() string {
	return "taobao.jushita.jdp.task.update"
}
func (this *JushitaJdpTaskUpdateRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *JushitaJdpTaskUpdateRequest) GetValues() url.Values {
	return this.values
}

/* 任务执行失败的错误信息，可以方便排查问题。<br /> 支持最大长度为：100<br /> 支持的最大列表长度为：100 */
func (this *JushitaJdpTaskUpdateRequest) SetErrorMessage(value string) {
	this.Set("error_message", value)
}

/* 任务执行的机器主机名,用于方便排查问题。<br /> 支持最大长度为：30<br /> 支持的最大列表长度为：30 */
func (this *JushitaJdpTaskUpdateRequest) SetExecuteHost(value string) {
	this.Set("execute_host", value)
}

/* 任务的编号 */
func (this *JushitaJdpTaskUpdateRequest) SetId(value string) {
	this.Set("id", value)
}

/* 任务的下次开始执行时间。 */
func (this *JushitaJdpTaskUpdateRequest) SetNextExecuteTime(value string) {
	this.Set("next_execute_time", value)
}

/* 任务当前同步到的时间点。 */
func (this *JushitaJdpTaskUpdateRequest) SetNowSyncTime(value string) {
	this.Set("now_sync_time", value)
}

/* 任务的参数，用json格式表示<br /> 支持最大长度为：512<br /> 支持的最大列表长度为：512 */
func (this *JushitaJdpTaskUpdateRequest) SetParams(value string) {
	this.Set("params", value)
}

/* 任务的状态，0表示任务新建或未执行完成，1表示任务停止，2表示任务已经完成，-1表示任务执行失败。<br /> 支持最大值为：2<br /> 支持最小值为：-1 */
func (this *JushitaJdpTaskUpdateRequest) SetStatus(value string) {
	this.Set("status", value)
}

/* 任务的类型,<br /> 支持最小值为：0 */
func (this *JushitaJdpTaskUpdateRequest) SetType(value string) {
	this.Set("type", value)
}

/* 任务更新时的版本号，此值在选择出的任务中获取到。 */
func (this *JushitaJdpTaskUpdateRequest) SetVersion(value string) {
	this.Set("version", value)
}

type JushitaJdpTaskUpdateResponse struct {
	*ErrorResponse                     `json:"error_response,omitempty"`
	JushitaJdpTaskUpdateResponseResult `json:"jushita_jdp_task_update_response"`
}
type JushitaJdpTaskUpdateResponseResult struct {
	/* 更新任务是否成功 */
	IsSuccess bool `json:"is_success"`
}

/*taobao.jushita.jdp.tasks.get*/
type JushitaJdpTasksGetRequest struct {
	values url.Values
}

func (this *JushitaJdpTasksGetRequest) GetApiMethodName() string {
	return "taobao.jushita.jdp.tasks.get"
}
func (this *JushitaJdpTasksGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *JushitaJdpTasksGetRequest) GetValues() url.Values {
	return this.values
}

/* 最大可返回的任务数量<br /> 支持最大值为：200<br /> 支持最小值为：1 */
func (this *JushitaJdpTasksGetRequest) SetFetchNum(value string) {
	this.Set("fetch_num", value)
}

/* 任务分组数量，表示把所有任务平均分成x组，在线订购应用此参数必传；非在线订购应用此参数忽略。<br /> 支持最小值为：1 */
func (this *JushitaJdpTasksGetRequest) SetTaskItemNum(value string) {
	this.Set("task_item_num", value)
}

/* 指定获取哪些分组的任务，取多个分组有半角逗号分隔，在线订购应用此参数必传；非在线订购应用此参数忽略。其中task_items >= 0且task_items < task_item_num */
func (this *JushitaJdpTasksGetRequest) SetTaskItems(value string) {
	this.Set("task_items", value)
}

/* 任务类型<br /> 支持最小值为：0 */
func (this *JushitaJdpTasksGetRequest) SetType(value string) {
	this.Set("type", value)
}

/* 需要查询哪些用户的任务，非在线订购应用此参数必传；在线订购应用此参数忽略。 */
func (this *JushitaJdpTasksGetRequest) SetUserIds(value string) {
	this.Set("user_ids", value)
}

type JushitaJdpTasksGetResponse struct {
	*ErrorResponse                   `json:"error_response,omitempty"`
	JushitaJdpTasksGetResponseResult `json:"jushita_jdp_tasks_get_response"`
}
type JushitaJdpTasksGetResponseResult struct {
	/* 查询到的未执行任务 */
	DataPushTasks struct {
		DataPushTask []*DataPushTask `json:"data_push_task"`
	} `json:"DataPushTasks"`
}

/*taobao.jushita.jdp.user.add*/
type JushitaJdpUserAddRequest struct {
	values url.Values
}

func (this *JushitaJdpUserAddRequest) GetApiMethodName() string {
	return "taobao.jushita.jdp.user.add"
}
func (this *JushitaJdpUserAddRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *JushitaJdpUserAddRequest) GetValues() url.Values {
	return this.values
}

/* RDS实例名称，如果有多个RDS必须设置。 */
func (this *JushitaJdpUserAddRequest) SetRdsName(value string) {
	this.Set("rds_name", value)
}

type JushitaJdpUserAddResponse struct {
	*ErrorResponse                  `json:"error_response,omitempty"`
	JushitaJdpUserAddResponseResult `json:"jushita_jdp_user_add_response"`
}
type JushitaJdpUserAddResponseResult struct {
	/* 是否添加成功 */
	IsSuccess bool `json:"is_success"`
}

/*taobao.jushita.jdp.user.delete*/
type JushitaJdpUserDeleteRequest struct {
	values url.Values
}

func (this *JushitaJdpUserDeleteRequest) GetApiMethodName() string {
	return "taobao.jushita.jdp.user.delete"
}
func (this *JushitaJdpUserDeleteRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *JushitaJdpUserDeleteRequest) GetValues() url.Values {
	return this.values
}

/* 要删除用户的昵称 */
func (this *JushitaJdpUserDeleteRequest) SetNick(value string) {
	this.Set("nick", value)
}

/* 需要删除的用户编号<br /> 支持最小值为：1 */
func (this *JushitaJdpUserDeleteRequest) SetUserId(value string) {
	this.Set("user_id", value)
}

type JushitaJdpUserDeleteResponse struct {
	*ErrorResponse                     `json:"error_response,omitempty"`
	JushitaJdpUserDeleteResponseResult `json:"jushita_jdp_user_delete_response"`
}
type JushitaJdpUserDeleteResponseResult struct {
	/* 是否删除成功 */
	IsSuccess bool `json:"is_success"`
}

/*taobao.jushita.jdp.users.get*/
type JushitaJdpUsersGetRequest struct {
	values url.Values
}

func (this *JushitaJdpUsersGetRequest) GetApiMethodName() string {
	return "taobao.jushita.jdp.users.get"
}
func (this *JushitaJdpUsersGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *JushitaJdpUsersGetRequest) GetValues() url.Values {
	return this.values
}

/* 此参数一般不用传，用于查询最后更改时间在某个时间段内的用户 */
func (this *JushitaJdpUsersGetRequest) SetEndModified(value string) {
	this.Set("end_modified", value)
}

/* 当前页数 */
func (this *JushitaJdpUsersGetRequest) SetPageNo(value string) {
	this.Set("page_no", value)
}

/* 每页记录数，默认200 */
func (this *JushitaJdpUsersGetRequest) SetPageSize(value string) {
	this.Set("page_size", value)
}

/* 此参数一般不用传，用于查询最后更改时间在某个时间段内的用户 */
func (this *JushitaJdpUsersGetRequest) SetStartModified(value string) {
	this.Set("start_modified", value)
}

/* 如果传了user_id表示单条查询 */
func (this *JushitaJdpUsersGetRequest) SetUserId(value string) {
	this.Set("user_id", value)
}

type JushitaJdpUsersGetResponse struct {
	*ErrorResponse                   `json:"error_response,omitempty"`
	JushitaJdpUsersGetResponseResult `json:"jushita_jdp_users_get_response"`
}
type JushitaJdpUsersGetResponseResult struct {
	/* 总记录数 */
	TotalResults int64 `json:"total_results"`
	/* 用户列表 */
	Users struct {
		JdpUser []*JdpUser `json:"jdp_user"`
	} `json:"Users"`
}

/*taobao.topats.jushita.jdp.datadelete*/
type TopatsJushitaJdpDatadeleteRequest struct {
	values url.Values
}

func (this *TopatsJushitaJdpDatadeleteRequest) GetApiMethodName() string {
	return "taobao.topats.jushita.jdp.datadelete"
}
func (this *TopatsJushitaJdpDatadeleteRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *TopatsJushitaJdpDatadeleteRequest) GetValues() url.Values {
	return this.values
}

/* 删除数据时间段的结束修改时间，格式为：yyyy-MM-dd HH:mm:ss，结束时间必须为前天的23:59:59秒以前，根据是业务的modified时间 */
func (this *TopatsJushitaJdpDatadeleteRequest) SetEndModified(value string) {
	this.Set("end_modified", value)
}

/* 删除数据时间段的起始修改时间，格式为：yyyy-MM-dd HH:mm:ss,根据是业务的modified时间 */
func (this *TopatsJushitaJdpDatadeleteRequest) SetStartModified(value string) {
	this.Set("start_modified", value)
}

/* 推送的数据类型，可选值为：tb_trade(淘宝交易)，tb_item(淘宝商品)，tb_refund(淘宝退款)，fx_trade(分销订单)，tm_refund(天猫退款)，同时删除多种类型以分号分隔，如："tb_trade;tb_item;tb_refund;fx_trade;tm_refund" */
func (this *TopatsJushitaJdpDatadeleteRequest) SetSyncType(value string) {
	this.Set("sync_type", value)
}

/* 用户昵称，不填表示删除所有用户的数据。 */
func (this *TopatsJushitaJdpDatadeleteRequest) SetUserNick(value string) {
	this.Set("user_nick", value)
}

type TopatsJushitaJdpDatadeleteResponse struct {
	*ErrorResponse                           `json:"error_response,omitempty"`
	TopatsJushitaJdpDatadeleteResponseResult `json:"topats_jushita_jdp_datadelete_response"`
}
type TopatsJushitaJdpDatadeleteResponseResult struct {
	/* 创建任务信息。里面只包含task_id和created */
	Task *Task `json:"task"`
}

/*alipay.xiaodai.user.permit*/
type AlipayXiaodaiUserPermitRequest struct {
	values url.Values
}

func (this *AlipayXiaodaiUserPermitRequest) GetApiMethodName() string {
	return "alipay.xiaodai.user.permit"
}
func (this *AlipayXiaodaiUserPermitRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *AlipayXiaodaiUserPermitRequest) GetValues() url.Values {
	return this.values
}

/* 用户数字ID */
func (this *AlipayXiaodaiUserPermitRequest) SetUserId(value string) {
	this.Set("user_id", value)
}

type AlipayXiaodaiUserPermitResponse struct {
	*ErrorResponse                        `json:"error_response,omitempty"`
	AlipayXiaodaiUserPermitResponseResult `json:"alipay_xiaodai_user_permit_response"`
}
type AlipayXiaodaiUserPermitResponseResult struct {
	/* 是否成功 */
	IsSuccess bool `json:"is_success"`
}

/*taobao.tmc.group.add*/
type TmcGroupAddRequest struct {
	values url.Values
}

func (this *TmcGroupAddRequest) GetApiMethodName() string {
	return "taobao.tmc.group.add"
}
func (this *TmcGroupAddRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *TmcGroupAddRequest) GetValues() url.Values {
	return this.values
}

/* 分组名称，同一个应用下需要保证唯一性，最长32个字符。添加分组后，消息通道会为用户的消息分配独立分组，但之前的消息还是存储于默认分组中。不能以default开头，default开头为系统默认组。<br /> 支持最大长度为：32<br /> 支持的最大列表长度为：32 */
func (this *TmcGroupAddRequest) SetGroupName(value string) {
	this.Set("group_name", value)
}

/* 用户昵称列表，以半角逗号分隔，支持子账号，支持增量添加用户 */
func (this *TmcGroupAddRequest) SetNicks(value string) {
	this.Set("nicks", value)
}

type TmcGroupAddResponse struct {
	*ErrorResponse            `json:"error_response,omitempty"`
	TmcGroupAddResponseResult `json:"tmc_group_add_response"`
}
type TmcGroupAddResponseResult struct {
	/* 创建时间 */
	Created string `json:"created"`
	/* 分组名称 */
	GroupName string `json:"group_name"`
}

/*taobao.tmc.group.delete*/
type TmcGroupDeleteRequest struct {
	values url.Values
}

func (this *TmcGroupDeleteRequest) GetApiMethodName() string {
	return "taobao.tmc.group.delete"
}
func (this *TmcGroupDeleteRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *TmcGroupDeleteRequest) GetValues() url.Values {
	return this.values
}

/* 分组名称，分组删除后，用户的消息将会存储于默认分组中。警告：由于分组已经删除，用户之前未消费的消息将无法再获取。不能以default开头，default开头为系统默认组。 */
func (this *TmcGroupDeleteRequest) SetGroupName(value string) {
	this.Set("group_name", value)
}

/* 用户列表，不传表示删除整个分组，如果用户全部删除后，也会自动删除整个分组 */
func (this *TmcGroupDeleteRequest) SetNicks(value string) {
	this.Set("nicks", value)
}

type TmcGroupDeleteResponse struct {
	*ErrorResponse               `json:"error_response,omitempty"`
	TmcGroupDeleteResponseResult `json:"tmc_group_delete_response"`
}
type TmcGroupDeleteResponseResult struct {
	/* 是否成功 */
	IsSuccess bool `json:"is_success"`
}

/*taobao.tmc.groups.get*/
type TmcGroupsGetRequest struct {
	values url.Values
}

func (this *TmcGroupsGetRequest) GetApiMethodName() string {
	return "taobao.tmc.groups.get"
}
func (this *TmcGroupsGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *TmcGroupsGetRequest) GetValues() url.Values {
	return this.values
}

/* 要查询分组的名称，多个分组用半角逗号分隔，不传代表查询所有分组信息，但不会返回组下面的用户信息。如果应用没有设置分组则返回空。组名不能以default开头，default开头是系统默认的组。 */
func (this *TmcGroupsGetRequest) SetGroupNames(value string) {
	this.Set("group_names", value)
}

/* 页码<br /> 支持最小值为：1 */
func (this *TmcGroupsGetRequest) SetPageNo(value string) {
	this.Set("page_no", value)
}

/* 每页返回多少个分组<br /> 支持最大值为：100<br /> 支持最小值为：1 */
func (this *TmcGroupsGetRequest) SetPageSize(value string) {
	this.Set("page_size", value)
}

type TmcGroupsGetResponse struct {
	*ErrorResponse             `json:"error_response,omitempty"`
	TmcGroupsGetResponseResult `json:"tmc_groups_get_response"`
}
type TmcGroupsGetResponseResult struct {
	/* 分组列表 */
	Groups struct {
		TmcGroup []*TmcGroup `json:"tmc_group"`
	} `json:"Groups"`
	/* 分组总数 */
	TotalResults int64 `json:"total_results"`
}

/*taobao.tmc.message.produce*/
type TmcMessageProduceRequest struct {
	values url.Values
}

func (this *TmcMessageProduceRequest) GetApiMethodName() string {
	return "taobao.tmc.message.produce"
}
func (this *TmcMessageProduceRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *TmcMessageProduceRequest) GetValues() url.Values {
	return this.values
}

/* 消息内容的JSON表述，必须按照topic的定义来填充<br /> 支持最大长度为：2000<br /> 支持的最大列表长度为：2000 */
func (this *TmcMessageProduceRequest) SetContent(value string) {
	this.Set("content", value)
}

/* 发布消息关联的主题<br /> 支持最大长度为：256<br /> 支持的最大列表长度为：256 */
func (this *TmcMessageProduceRequest) SetTopic(value string) {
	this.Set("topic", value)
}

type TmcMessageProduceResponse struct {
	*ErrorResponse                  `json:"error_response,omitempty"`
	TmcMessageProduceResponseResult `json:"tmc_message_produce_response"`
}
type TmcMessageProduceResponseResult struct {
	/* 是否成功 */
	IsSuccess bool `json:"is_success"`
}

/*taobao.tmc.messages.confirm*/
type TmcMessagesConfirmRequest struct {
	values url.Values
}

func (this *TmcMessagesConfirmRequest) GetApiMethodName() string {
	return "taobao.tmc.messages.confirm"
}
func (this *TmcMessagesConfirmRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *TmcMessagesConfirmRequest) GetValues() url.Values {
	return this.values
}

/* 处理失败的消息ID列表--已废弃，无需传此字段 */
func (this *TmcMessagesConfirmRequest) SetFMessageIds(value string) {
	this.Set("f_message_ids", value)
}

/* 分组名称，不传代表默认分组 */
func (this *TmcMessagesConfirmRequest) SetGroupName(value string) {
	this.Set("group_name", value)
}

/* 处理成功的消息ID列表
最大 200个ID */
func (this *TmcMessagesConfirmRequest) SetSMessageIds(value string) {
	this.Set("s_message_ids", value)
}

type TmcMessagesConfirmResponse struct {
	*ErrorResponse                   `json:"error_response,omitempty"`
	TmcMessagesConfirmResponseResult `json:"tmc_messages_confirm_response"`
}
type TmcMessagesConfirmResponseResult struct {
	/* 是否成功 */
	IsSuccess bool `json:"is_success"`
}

/*taobao.tmc.messages.consume*/
type TmcMessagesConsumeRequest struct {
	values url.Values
}

func (this *TmcMessagesConsumeRequest) GetApiMethodName() string {
	return "taobao.tmc.messages.consume"
}
func (this *TmcMessagesConsumeRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *TmcMessagesConsumeRequest) GetValues() url.Values {
	return this.values
}

/* 用户分组名称，不传表示消费默认分组，如果应用没有设置用户分组，传入分组名称将会返回错误 */
func (this *TmcMessagesConsumeRequest) SetGroupName(value string) {
	this.Set("group_name", value)
}

/* 每次批量消费消息的条数<br /> 支持最大值为：200<br /> 支持最小值为：10 */
func (this *TmcMessagesConsumeRequest) SetQuantity(value string) {
	this.Set("quantity", value)
}

type TmcMessagesConsumeResponse struct {
	*ErrorResponse                   `json:"error_response,omitempty"`
	TmcMessagesConsumeResponseResult `json:"tmc_messages_consume_response"`
}
type TmcMessagesConsumeResponseResult struct {
	/* 消息列表 */
	Messages struct {
		TmcMessage []*TmcMessage `json:"tmc_message"`
	} `json:"Messages"`
}

/*taobao.tmc.user.cancel*/
type TmcUserCancelRequest struct {
	values url.Values
}

func (this *TmcUserCancelRequest) GetApiMethodName() string {
	return "taobao.tmc.user.cancel"
}
func (this *TmcUserCancelRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *TmcUserCancelRequest) GetValues() url.Values {
	return this.values
}

/* 用户昵称，支持子账号 */
func (this *TmcUserCancelRequest) SetNick(value string) {
	this.Set("nick", value)
}

type TmcUserCancelResponse struct {
	*ErrorResponse              `json:"error_response,omitempty"`
	TmcUserCancelResponseResult `json:"tmc_user_cancel_response"`
}
type TmcUserCancelResponseResult struct {
	/* 是否成功,如果为false并且没有错误码，表示删除的用户不存在。 */
	IsSuccess bool `json:"is_success"`
}

/*taobao.tmc.user.get*/
type TmcUserGetRequest struct {
	values url.Values
}

func (this *TmcUserGetRequest) GetApiMethodName() string {
	return "taobao.tmc.user.get"
}
func (this *TmcUserGetRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *TmcUserGetRequest) GetValues() url.Values {
	return this.values
}

/* 需返回的字段列表，多个字段以半角逗号分隔。可选值：TmcUser结构体中的所有字段，一定要返回topic。 */
func (this *TmcUserGetRequest) SetFields(value string) {
	this.Set("fields", value)
}

/* 用户昵称<br /> 支持最大长度为：100<br /> 支持的最大列表长度为：100 */
func (this *TmcUserGetRequest) SetNick(value string) {
	this.Set("nick", value)
}

type TmcUserGetResponse struct {
	*ErrorResponse           `json:"error_response,omitempty"`
	TmcUserGetResponseResult `json:"tmc_user_get_response"`
}
type TmcUserGetResponseResult struct {
	/* 开通的用户数据 */
	TmcUser *TmcUser `json:"tmc_user"`
}

/*taobao.tmc.user.permit*/
type TmcUserPermitRequest struct {
	values url.Values
}

func (this *TmcUserPermitRequest) GetApiMethodName() string {
	return "taobao.tmc.user.permit"
}
func (this *TmcUserPermitRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *TmcUserPermitRequest) GetValues() url.Values {
	return this.values
}

/* 消息主题列表，用半角逗号分隔。当用户订阅的topic是应用订阅的子集时才需要设置，不设置表示继承应用所订阅的所有topic，一般情况建议不要设置。 */
func (this *TmcUserPermitRequest) SetTopics(value string) {
	this.Set("topics", value)
}

type TmcUserPermitResponse struct {
	*ErrorResponse              `json:"error_response,omitempty"`
	TmcUserPermitResponseResult `json:"tmc_user_permit_response"`
}
type TmcUserPermitResponseResult struct {
	/* 是否成功 */
	IsSuccess bool `json:"is_success"`
}

/*taobao.wireless.bunting.item.shorturl.create*/
type WirelessBuntingItemShorturlCreateRequest struct {
	values url.Values
}

func (this *WirelessBuntingItemShorturlCreateRequest) GetApiMethodName() string {
	return "taobao.wireless.bunting.item.shorturl.create"
}
func (this *WirelessBuntingItemShorturlCreateRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *WirelessBuntingItemShorturlCreateRequest) GetValues() url.Values {
	return this.values
}

/* 商品ID */
func (this *WirelessBuntingItemShorturlCreateRequest) SetItemId(value string) {
	this.Set("item_id", value)
}

type WirelessBuntingItemShorturlCreateResponse struct {
	*ErrorResponse                                  `json:"error_response,omitempty"`
	WirelessBuntingItemShorturlCreateResponseResult `json:"wireless_bunting_item_shorturl_create_response"`
}
type WirelessBuntingItemShorturlCreateResponseResult struct {
	/* 短地址 */
	Shorturl string `json:"shorturl"`
}

/*taobao.wireless.bunting.shop.shorturl.create*/
type WirelessBuntingShopShorturlCreateRequest struct {
	values url.Values
}

func (this *WirelessBuntingShopShorturlCreateRequest) GetApiMethodName() string {
	return "taobao.wireless.bunting.shop.shorturl.create"
}
func (this *WirelessBuntingShopShorturlCreateRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *WirelessBuntingShopShorturlCreateRequest) GetValues() url.Values {
	return this.values
}

/* 商店id */
func (this *WirelessBuntingShopShorturlCreateRequest) SetShopId(value string) {
	this.Set("shop_id", value)
}

type WirelessBuntingShopShorturlCreateResponse struct {
	*ErrorResponse                                  `json:"error_response,omitempty"`
	WirelessBuntingShopShorturlCreateResponseResult `json:"wireless_bunting_shop_shorturl_create_response"`
}
type WirelessBuntingShopShorturlCreateResponseResult struct {
	/* 短链 */
	Shorturl string `json:"shorturl"`
}

/*taobao.ma.packcode.create*/
type MaPackcodeCreateRequest struct {
	values url.Values
}

func (this *MaPackcodeCreateRequest) GetApiMethodName() string {
	return "taobao.ma.packcode.create"
}
func (this *MaPackcodeCreateRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *MaPackcodeCreateRequest) GetValues() url.Values {
	return this.values
}

type MaPackcodeCreateResponse struct {
	*ErrorResponse                 `json:"error_response,omitempty"`
	MaPackcodeCreateResponseResult `json:"ma_packcode_create_response"`
}
type MaPackcodeCreateResponseResult struct {
	/* 生成包裹码二维码图片链接 */
	Qrcodeurl string `json:"qrcodeurl"`
}

/*taobao.ma.qrcode.common.create*/
type MaQrcodeCommonCreateRequest struct {
	values url.Values
}

func (this *MaQrcodeCommonCreateRequest) GetApiMethodName() string {
	return "taobao.ma.qrcode.common.create"
}
func (this *MaQrcodeCommonCreateRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *MaQrcodeCommonCreateRequest) GetValues() url.Values {
	return this.values
}

/* 二维码需要布点的位置，方便用户在码平台上可以区分看到不同布点的扫码数据情况；列表值，用半角','号分割，单个渠道名不能超过16字符。<br /> 支持最大长度为：160<br /> 支持的最大列表长度为：160 */
func (this *MaQrcodeCommonCreateRequest) SetChannelName(value string) {
	this.Set("channel_name", value)
}

/* 二维码的内容之一，由type决定：type=page时，content传入无线页面的URL连接内容；type=item时，content传入宝贝数字ID；type=url时，content传入普通的URL连接内容；type=shop时，content传入店铺ID；<br /> 支持最大长度为：256<br /> 支持的最大列表长度为：256 */
func (this *MaQrcodeCommonCreateRequest) SetContent(value string) {
	this.Set("content", value)
}

/* 二维码的logo地址，只允许淘宝官方图片空间的图片地址，其他非图片空间图片不支持。官方淘logo图片地址：http://img01.taobaocdn.com/imgextra/T1Od8YFT8eXXXXXXXX。<br /> 支持最大长度为：256<br /> 支持的最大列表长度为：256 */
func (this *MaQrcodeCommonCreateRequest) SetLogo(value string) {
	this.Set("logo", value)
}

/* 二维码名字，即创建的二维码，在码平台上显示记录的名字。<br /> 支持最大长度为：32<br /> 支持的最大列表长度为：32 */
func (this *MaQrcodeCommonCreateRequest) SetName(value string) {
	this.Set("name", value)
}

/* 是否需要矢量图，如果需要矢量图，设置为true；只支持普通二维码，官方模板不支持矢量图 */
func (this *MaQrcodeCommonCreateRequest) SetNeedEps(value string) {
	this.Set("need_eps", value)
}

/* 二维码尺寸，只支持普通二维码，不支持官方模板，单位为像素，最小为60×60，最大为300×300，建议175×175。官方模板大小尺寸见style说明。<br /> 支持最大值为：300<br /> 支持最小值为：60 */
func (this *MaQrcodeCommonCreateRequest) SetSize(value string) {
	this.Set("size", value)
}

/* 二维码的样式名，支持普通码的颜色或官方模板的模板名；普通码的颜色可选输入：“000000”(黑色)、“EF4F2B”(橙色);官方模板的可选输入（实际尺寸比样例大）：“ww_color.png“ 尺寸290x320，样例：http://gtms03.alicdn.com/tps/i3/T1YLPRFRXXXXbsbYwb-100-102.png；“tb_scan.png“ 尺寸290x320，样例：http://gtms01.alicdn.com/tps/i1/T14vsEFThdXXbsbYwb-100-102.png；“ww_hide_color.png“  尺寸200x263，样例：http://gtms04.alicdn.com/tps/i4/TB1URvlFVXXXXbRXFXXwxcf6pXX-76-100.png；“tmall_hide_color.png“ 尺寸200x263，样例：http://gtms01.alicdn.com/tps/i1/TB1S5PiFVXXXXacXVXXwxcf6pXX-76-100.png。<br /> 支持最大长度为：50<br /> 支持的最大列表长度为：50 */
func (this *MaQrcodeCommonCreateRequest) SetStyle(value string) {
	this.Set("style", value)
}

/* 要制作的二维码业务类型：page:无线页面类型item:宝贝ID类型url:普通的URL连接类型shop:店铺ID类型<br /> 支持最大长度为：10<br /> 支持的最大列表长度为：10 */
func (this *MaQrcodeCommonCreateRequest) SetType(value string) {
	this.Set("type", value)
}

type MaQrcodeCommonCreateResponse struct {
	*ErrorResponse                     `json:"error_response,omitempty"`
	MaQrcodeCommonCreateResponseResult `json:"ma_qrcode_common_create_response"`
}
type MaQrcodeCommonCreateResponseResult struct {
	/* 二维码对像 */
	Modules struct {
		QrcodeDO []*QrcodeDO `json:"qrcode_d_o"`
	} `json:"Modules"`
	/* 执行是否成功 */
	Suc bool `json:"suc"`
}

/*taobao.mbc.promotion.use*/
type MbcPromotionUseRequest struct {
	values url.Values
}

func (this *MbcPromotionUseRequest) GetApiMethodName() string {
	return "taobao.mbc.promotion.use"
}
func (this *MbcPromotionUseRequest) Set(key, value string) {
	if this.values == nil {
		this.values = url.Values{}
	}
	this.values.Set(key, value)
}
func (this *MbcPromotionUseRequest) GetValues() url.Values {
	return this.values
}

/* 实际金额单位(分) */
func (this *MbcPromotionUseRequest) SetActualFee(value string) {
	this.Set("actual_fee", value)
}

/* 优惠金额单位(分) */
func (this *MbcPromotionUseRequest) SetDiscountFee(value string) {
	this.Set("discount_fee", value)
}

/* 有效结束时间 */
func (this *MbcPromotionUseRequest) SetEndTime(value string) {
	this.Set("end_time", value)
}

/* 外部流水号，
promotion_type+outer_no 唯一 */
func (this *MbcPromotionUseRequest) SetOuterNo(value string) {
	this.Set("outer_no", value)
}

/* 权益ID */
func (this *MbcPromotionUseRequest) SetPromotionId(value string) {
	this.Set("promotion_id", value)
}

/* 权限类型 */
func (this *MbcPromotionUseRequest) SetPromotionType(value string) {
	this.Set("promotion_type", value)
}

/* 有效开始时间 */
func (this *MbcPromotionUseRequest) SetStartTime(value string) {
	this.Set("start_time", value)
}

/* 交易总金额单位(分) */
func (this *MbcPromotionUseRequest) SetTotalFee(value string) {
	this.Set("total_fee", value)
}

/* 使用时间 */
func (this *MbcPromotionUseRequest) SetUseTime(value string) {
	this.Set("use_time", value)
}

/* 买家淘宝ID */
func (this *MbcPromotionUseRequest) SetUserId(value string) {
	this.Set("user_id", value)
}

type MbcPromotionUseResponse struct {
	*ErrorResponse                `json:"error_response,omitempty"`
	MbcPromotionUseResponseResult `json:"mbc_promotion_use_response"`
}
type MbcPromotionUseResponseResult struct {
	/* 淘宝这边生成的唯一ID */
	Id int64 `json:"id"`
	/* 0：失败；1：成功 */
	RetCode int64 `json:"ret_code"`
}
