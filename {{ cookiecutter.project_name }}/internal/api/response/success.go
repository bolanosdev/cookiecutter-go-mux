package response

import (
	"net/http"
	"strconv"

	"{{ cookiecutter.group_name }}/{{ cookiecutter.project_name }}/internal/utils"

	"go.opentelemetry.io/otel/trace"
)

type ResultInfo struct {
	Page       int `json:"page"`
	PageSize   int `json:"page_size"`
	Count      int `json:"count"`
	TotalCount int `json:"total_count"`
}

type V1SuccessResponse struct {
	AccessToken interface{} `json:"access_token,omitempty"`
	Success     bool        `json:"success"`
	Error       string      `json:"error"`
	Messages    []string    `json:"messages"`
	Result      any         `json:"result"`
	ResultInfo  interface{} `json:"result_info,omitempty"`
}

func Success(w http.ResponseWriter, r *http.Request, status int, result any) {
	response := V1SuccessResponse{
		Success: true,
		Result:  result,
	}
	trace_id := trace.SpanContextFromContext(r.Context()).TraceID()

	w.Header().Set("x-status-code", strconv.Itoa(status))
	w.Header().Set("x-trace-id", trace_id.String())
	utils.Encode(w, status, response)
}
