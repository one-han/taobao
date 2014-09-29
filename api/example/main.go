package main

import (
	"encoding/json"
	"github.com/ohohco/taobao/api"
	"github.com/ohohco/taobao/api/request"
	"github.com/ohohco/taobao/api/response"
	"log"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}
func main() {

	client := api.DefaultTaobaoClient{
		"http://gw.api.taobao.com/router/rest",
		"your appKey",
		"your appSecret",
	}
	var resp response.ItemGetResponse
	var req = new(request.ItemGetRequest)
	req.SetFields("detail_url,num_iid,title,nick,type,cid,seller_cids,props,input_pids,input_str,desc,pic_url,num,valid_thru,list_time,delist_time,stuff_status,location,price,post_fee,express_fee,ems_fee,has_discount,freight_payer,has_invoice,has_warranty,has_showcase,modified,increment,approve_status,postage_id,product_id,auction_point,property_alias,item_img,prop_img,sku,video,outer_id,is_virtual")
	req.SetNumIid("41195705534")
	_, err := client.Excute(req, &resp, "your sessionKey")
	// log.Println("data:", string(_bytes))
	if err != nil {
		log.Fatalln(err)
	}
	if resp.ErrorResponse != nil {
		log.Fatalln(resp.Msg)
	}
	log.Println("error:", resp.Response.Item.Desc)
}

func t3() {
	data := `{"item_get_response":{"item":{"nick":"歌莉娅官方旗舰店","num_iid":41195705534,"title":"[纪念款]歌莉娅2014冬季连衣裙印花欧根纱长袖连衣裙14DE4E790"}}}`
	var resp response.ItemGetResponse
	// resp.Body = "歌莉娅官方旗舰店"
	// resp.SubCode = "41195705534"
	// resp.Response.Item.Nick = "歌莉娅官方旗舰店"
	// if _bytes, err := json.Marshal(&resp); err != nil {
	// 	log.Fatalln(err)
	// } else {
	// 	log.Println(string(_bytes))
	// }
	// return
	if err := json.Unmarshal([]byte(data), &resp); err != nil {
		log.Fatalln(err)
	}
	log.Println(resp.Response.Item.Nick)
}
