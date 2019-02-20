package response

import (
	"github.com/one-han/taobao/api"
	. "github.com/one-han/taobao/api/domain"
)

type TbkDgOptimusMaterialResponse struct {
	*api.ErrorResponse                     `json:"error_response,omitempty"`
	TbkDgOptimusMaterialResponseResultList `json:"tbk_dg_optimus_material_response"`
}

type TbkDgOptimusMaterialResponseResultList struct {
	/* resultList */
	ResultList TbkDgOptimusMaterialResponseResult `json:"result_list"`
}

type TbkDgOptimusMaterialResponseResult struct {
	/* resultList */
	MapData []*DgMapData `json:"map_data"`
}
