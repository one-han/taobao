package response

import (
	"github.com/ohohco/taobao/api"
	. "github.com/ohohco/taobao/api/domain"
)

type ItemUpdateResponse struct {
	*api.ErrorResponse       `json:"error_response,omitempty"`
	ItemUpdateResponseResult `json:"item_update_response"`
}
type ItemUpdateResponseResult struct {
	/* 商品结构里的num_iid，modified */
	Item *Item `json:"item"`
}
