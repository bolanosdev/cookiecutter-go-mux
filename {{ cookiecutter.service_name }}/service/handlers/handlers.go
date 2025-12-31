package handlers

import (
	"errors"
	"net/http"

	"{{cookiecutter.group_name}}/{{cookiecutter.service_name}}/internal/kerr"
	"{{cookiecutter.group_name}}/{{cookiecutter.service_name}}/internal/utils/encoder"
	"{{cookiecutter.group_name}}/{{cookiecutter.service_name}}/service/entities/output"

	"github.com/bolanosdev/go-snacks/observability/sentry"
)

type errorInfoSetter interface {
	SetErrorInfo(errorMsg string, errorCode int, eventID string)
}

type BaseHandler struct {
	sentry *sentry.SentryObs
}

func NewBaseHandler() BaseHandler {
	return BaseHandler{}
}

func (h *BaseHandler) Error(w http.ResponseWriter, r *http.Request, err error) {
	status_code := http.StatusInternalServerError
	var error_code int
	var message string
	var event_id string

	type statusCoder interface {
		StatusCode() int
	}

	var sc statusCoder
	if errors.As(err, &sc) {
		status_code = sc.StatusCode()
	}

	var knownError *kerr.KnownError
	if errors.As(err, &knownError) {
		error_code = knownError.Code
		message = knownError.Message
	} else {
		error_code = 1000
		message = "Something went wrong"
		if h.sentry != nil {
			sentryEventID := h.sentry.CaptureError(err, error_code)
			if sentryEventID != nil {
				event_id = string(*sentryEventID)
			}
		}
	}

	if setter, ok := w.(errorInfoSetter); ok {
		setter.SetErrorInfo(err.Error(), error_code, event_id)
	}

	response := output.V1ErrorResponse{
		Success: false,
		Error: output.Error{
			Code:    error_code,
			Message: message,
		},
	}

	encoder.Encode(w, status_code, response)
}

func (h *BaseHandler) Success(w http.ResponseWriter, r *http.Request, result any) {
	response := output.V1SuccessResponse{
		Success: true,
		Result:  result,
	}

	encoder.Encode(w, http.StatusOK, response)
}
