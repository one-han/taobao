package domain

/* resultList */
type DgMapData struct {
	/* 叶子类目id */
	CategoryId int64 `json:"category_id"`

	/* 叶子类目名称 */
	CategoryName string `json:"category_name"`

	/* 单品淘客链接 */
	ClickUrl string `json:"click_url"`

	/* 佣金比率(%) */
	CommissionRate string `json:"commission_rate"`

	/* 券面额 */
	CouponAmount int64 `json:"coupon_amount"`

	/* 优惠券链接 */
	CouponClickUrl string `json:"coupon_click_url"`

	/* 优惠券结束时间 */
	CouponEndTime string `json:"coupon_end_time"`

	/* 优惠券剩余量 */
	CouponRemainCount int64 `json:"coupon_remain_count"`

	/* 优惠券起用门槛，满X元可用 */
	CouponStartFee string `json:"coupon_start_fee"`

	/* 优惠券开始时间 */
	CouponStartTime string `json:"coupon_start_time"`

	/* 券总量 */
	CouponTotalCount int64 `json:"coupon_total_count"`

	/* 宝贝描述（推荐理由,不一定有） */
	ItemDescription string `json:"item_description"`

	/* 宝贝id */
	ItemId int64 `json:"item_id"`

	/* 拼团：几人团 */
	JddNum int64 `json:"jdd_num"`

	/* 拼团：拼成价，单位元 */
	JddPrice string `json:"jdd_price"`

	/* 一级类目ID */
	LevelOneCategoryId int64 `json:"level_one_category_id"`

	/* 一级类目名称 */
	LevelOneCategoryName string `json:"level_one_category_name"`

	/* 拼团：结束时间 */
	Oetime string `json:"oetime"`

	/* 拼团：一人价（原价)，单位元 */
	OrigPrice string `json:"orig_price"`

	/* 拼团：开始时间 */
	Ostime string `json:"ostime"`

	/* 商品主图 */
	PictUrl string `json:"pict_url"`

	/* 拼团：已售数量 */
	SellNum int64 `json:"sell_num"`

	/* 卖家id */
	SellerId int64 `json:"seller_id"`

	/* 店铺名称 */
	ShopTitle string `json:"shop_title"`

	/* 商品短标题 */
	ShortTitle string `json:"short_title"`

	/* 商品小图列表 */
	SmallImages []string `json:"small_images"`

	/* 拼团：剩余库存 */
	Stock int64 `json:"stock"`

	/* 商品标题 */
	Title string `json:"title"`

	/* 天猫营销玩法 */
	TmallPlayActivityInfo string `json:"tmall_play_activity_info"`

	/* 拼团：库存数量 */
	TotalStock int64 `json:"total_stock"`

	/* 卖家类型，0表示集市，1表示商城 */
	UserType int64 `json:"user_type"`

	/* 预售数量 */
	UvSumPreSale int64 `json:"uv_sum_pre_sale"`

	/* 30天销量 */
	Volume int64 `json:"volume"`

	/* 商品白底图 */
	WhiteImage string `json:"white_image"`

	/* 商品关联词 */
	WordList []*WordMapData `json:"word_list>word_map_data"`

	/* 折扣价 */
	ZkFinalPrice string `json:"zk_final_price"`
}
