package response

import (
	"github.com/one-han/taobao/api"
	. "github.com/one-han/taobao/api/domain"
)

type TbkDgMaterialOptionalResponse struct {
	*api.ErrorResponse                      `json:"error_response,omitempty"`
	TbkDgMaterialOptionalResponseResultList `json:"tbk_dg_material_optional_response"`
}
type TbkDgMaterialOptionalResponseResultList struct {
	/* resultList */
	ResultList TbkDgMaterialOptionalResponseResult `json:"result_list"`
	/* 搜索到符合条件的结果总数 */
	TotalResults int64 `json:"total_results"`
}

type TbkDgMaterialOptionalResponseResult struct {
	/* resultList */
	MapData []*DgSearchMapData `json:"map_data"`
}
