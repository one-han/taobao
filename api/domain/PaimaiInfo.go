
	package domain
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
	