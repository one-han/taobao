package response

import (
	"github.com/one-han/taobao/api"
	. "github.com/one-han/taobao/api/domain"
)

type TbkTpwdCreateResponse struct {
	*api.ErrorResponse          `json:"error_response,omitempty"`
	TbkTpwdCreateResponseResult `json:"tbk_tpwd_create_response"`
}
type TbkTpwdCreateResponseResult struct {
	/* data */
	Data *MapData `json:"data"`
}
