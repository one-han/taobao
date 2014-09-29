package domain

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
