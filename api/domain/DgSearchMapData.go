package domain

/* resultList */
type DgSearchMapData struct {
	/* 叶子类目id */
	CategoryId int64 `json:"category_id"`

	/* 叶子类目名称 */
	CategoryName string `json:"category_name"`

	/* 佣金比率 */
	CommissionRate string `json:"commission_rate"`

	/* 佣金类型 */
	CommissionType string `json:"commission_type"`

	/* 优惠券结束时间 */
	CouponEndTime string `json:"coupon_end_time"`

	/* 优惠券id */
	CouponId string `json:"coupon_id"`

	/* 优惠券面额 */
	CouponInfo string `json:"coupon_info"`

	/* 优惠券剩余量 */
	CouponRemainCount int64 `json:"coupon_remain_count"`

	/* 券二合一页面链接 */
	CouponShareUrl string `json:"coupon_share_url"`

	/* 优惠券开始时间 */
	CouponStartTime string `json:"coupon_start_time"`

	/* 优惠券总量 */
	CouponTotalCount int64 `json:"coupon_total_count"`

	/* 是否包含定向计划 */
	IncludeDxjh string `json:"include_dxjh"`

	/* 券面额 */
	CouponAmount string `json:"coupon_amount"`

	/* 是否包含营销计划 */
	IncludeMkt string `json:"include_mkt"`

	/* 定向计划信息 */
	InfoDxjh string `json:"info_dxjh"`

	/* 宝贝id */
	ItemId int64 `json:"item_id"`

	/* 商品地址 */
	ItemUrl string `json:"item_url"`

	/* 拼团：几人团 */
	JddNum int64 `json:"jdd_num"`

	/* 拼团：拼成价，单位元 */
	JddPrice string `json:"jdd_price"`

	/* 一级类目ID */
	LevelOneCategoryId int64 `json:"level_one_category_id"`

	/* 一级类目名称 */
	LevelOneCategoryName string `json:"level_one_category_name"`

	/* 宝贝id */
	NumIid int64 `json:"num_iid"`

	/* 拼团：结束时间 */
	Oetime string `json:"oetime"`

	/* 拼团：开始时间 */
	Ostime string `json:"ostime"`

	/* 商品主图 */
	PictUrl string `json:"pict_url"`

	/* 宝贝所在地 */
	Provcity string `json:"provcity"`

	/* 商品一口价格 */
	ReservePrice string `json:"reserve_price"`

	/* 卖家id */
	SellerId int64 `json:"seller_id"`

	/* 店铺dsr评分 */
	ShopDsr int64 `json:"shop_dsr"`

	/* 店铺名称 */
	ShopTitle string `json:"shop_title"`

	/* 商品短标题 */
	ShortTitle string `json:"short_title"`

	/* 商品小图列表 */
	SmallImages []string `json:"small_images"`

	/* 商品标题 */
	Title string `json:"title"`

	/* 月支出佣金 */
	TkTotalCommi string `json:"tk_total_commi"`

	/* 淘客30天月推广量 */
	TkTotalSales string `json:"tk_total_sales"`

	/* 商品淘客链接 */
	Url string `json:"url"`

	/* 卖家类型，0表示集市，1表示商城 */
	UserType int64 `json:"user_type"`

	/* 预售数量 */
	UvSumPreSale int64 `json:"uv_sum_pre_sale"`

	/* 30天销量 */
	Volume int64 `json:"volume"`

	/* 商品白底图 */
	WhiteImage string `json:"white_image"`

	/* 商品折扣价格 */
	ZkFinalPrice string `json:"zk_final_price"`
}
