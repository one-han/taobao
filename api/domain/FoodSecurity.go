
	package domain
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
	