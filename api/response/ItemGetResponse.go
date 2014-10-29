package response

import (
	"github.com/ohohco/taobao/api"
	. "github.com/ohohco/taobao/api/domain"
)

type ItemGetResponse struct {
	*api.ErrorResponse    `json:"error_response,omitempty"`
	ItemGetResponseResult `json:"item_get_response"`
}
type ItemGetResponseResult struct {
	/* 获取的商品 具体字段根据权限和设定的fields决定 */
	Item *Item `json:"item"`
}
