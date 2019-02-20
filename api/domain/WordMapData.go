package domain

/* 商品关联词 */
type WordMapData struct {
	/* 商品相关关联词落地页地址 */
	Url string `json:"url"`

	/* 商品相关的关联词 */
	Word string `json:"word"`
}
