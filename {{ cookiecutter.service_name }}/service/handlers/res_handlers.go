package handlers

import (
	"net/http"
	"strconv"

	"{{cookiecutter.group_name}}/{{cookiecutter.service_name}}/internal/utils"
	"{{cookiecutter.group_name}}/{{cookiecutter.service_name}}/service/entities/response"

	"go.opentelemetry.io/otel/trace"
)

func Error(w http.ResponseWriter, r *http.Request, status int, err error, message string) {
	trace_id := trace.SpanContextFromContext(r.Context()).TraceID()
	res := response.V1ErrorResponse{
		Code:    status,
		Message: message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("x-server-readonly", "true")
	w.Header().Set("x-status-code", strconv.Itoa(status))
	w.Header().Set("x-trace-id", trace_id.String())
	w.Header().Set("x-error-message", err.Error())

	utils.Encode(w, status, res)
}

func Success(w http.ResponseWriter, r *http.Request, status int, result any) {
	res := response.V1SuccessResponse{
		Success: true,
		Result:  result,
	}
	trace_id := trace.SpanContextFromContext(r.Context()).TraceID()

	w.Header().Set("x-server-readonly", "true")
	w.Header().Set("x-status-code", strconv.Itoa(status))
	w.Header().Set("x-trace-id", trace_id.String())
	utils.Encode(w, status, res)
}
