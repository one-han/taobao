
	package domain
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
	