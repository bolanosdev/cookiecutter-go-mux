package output

type ResultInfo struct {
	Page       int `json:"page"`
	PageSize   int `json:"page_size"`
	Count      int `json:"count"`
	TotalCount int `json:"total_count"`
}

type Error struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type V1ErrorResponse struct {
	Success bool  `json:"success"`
	Error   Error `json:"error"`
}

type V1SuccessResponse struct {
	AccessToken interface{} `json:"access_token,omitempty"`
	Success     bool        `json:"success"`
	Error       *Error      `json:"error,omitempty"`
	Messages    []string    `json:"messages"`
	Result      any         `json:"result"`
	ResultInfo  interface{} `json:"result_info,omitempty"`
}
