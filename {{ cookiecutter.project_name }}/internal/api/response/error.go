package response

import (
	"net/http"
	"strconv"

	"go.opentelemetry.io/otel/trace"
)

type V1ErrorResponse struct {
	Success   bool   `json:"success"`
	Error     string `json:"error"`
	ErrorCode int    `json:"error_code"`
}

func Error(w http.ResponseWriter, r *http.Request, status int, err error, message string) {
	trace_id := trace.SpanContextFromContext(r.Context()).TraceID()

	w.Header().Set("x-status-code", strconv.Itoa(status))
	w.Header().Set("x-trace-id", trace_id.String())
	w.Header().Set("x-error-message", err.Error())

	http.Error(w, message, http.StatusBadRequest)
}
