package domain

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
