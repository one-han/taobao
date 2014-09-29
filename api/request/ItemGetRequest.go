
	package request
	import (
		"net/url"
	)
	
	type ItemGetRequest struct{
		values url.Values
	}
	
	func (r *ItemGetRequest)GetApiMethodName() string{
		return "taobao.item.get"
	}
	func (r *ItemGetRequest)SetValue(key,value string){
		if r.values == nil {
			r.values = url.Values{}
		}
		r.values.Set(key, value)
	}
	func (r *ItemGetRequest)GetValues() url.Values{
		return r.values
	}
	



	
	/* 需要返回的商品对象字段，如title,price,desc_modules等。可选值：Item商品结构体中所有字段均可返回；多个字段用“,”分隔。<br>新增返回字段：item_weight(商品的重量，格式为数字，包含小数)、item_size(商品的体积，格式为数字，包含小数)、change_prop（商品基础色数据） */
	func (r *ItemGetRequest) SetFields(value string) {
		r.SetValue("fields", value)
	}
	
	/* 商品数字ID<br /> 支持最小值为：1 */
	func (r *ItemGetRequest) SetNumIid(value string) {
		r.SetValue("num_iid", value)
	}
	
	/* 商品数字ID(带有跟踪效果) */
	func (r *ItemGetRequest) SetTrackIid(value string) {
		r.SetValue("track_iid", value)
	}
	
	