package response

type ResultInfo struct {
	Page       int `json:"page"`
	PageSize   int `json:"page_size"`
	Count      int `json:"count"`
	TotalCount int `json:"total_count"`
}

type V1ErrorResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type V1SuccessResponse struct {
	AccessToken interface{} `json:"access_token,omitempty"`
	Success     bool        `json:"success"`
	Error       string      `json:"error"`
	Messages    []string    `json:"messages"`
	Result      any         `json:"result"`
	ResultInfo  interface{} `json:"result_info,omitempty"`
}
