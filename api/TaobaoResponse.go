package api

type ErrorResponse struct {
	ErrorCode          int    `json:"code" xml:"code"`
	Msg                string `json:"msg" xml:"msg"`
	SubCode            string `json:"sub_code" xml:"sub_code"`
	SubMsg             string `json:"sub_msg" xml:"sub_msg"`
	TopForbiddenFields string `json:"top_forbidden_fields" xml:"top_forbidden_fields"`
	Body               string `json:"body" xml:"body"`
	// Params             map[string]string `json:"params" xml:"params"`
}
