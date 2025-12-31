package kerr

import (
	"fmt"
)

type KnownError struct {
	Code        int
	Message     string
	Range, Name string
}

func (k *KnownError) Error() string {
	return k.Message
}

type HttpStatusCoder struct {
	error
	statusCode int
}

func (h *HttpStatusCoder) Unwrap() error {
	return h.error
}

func (h *HttpStatusCoder) StatusCode() int {
	return h.statusCode
}

const (
	GeneralInvalidRequestParamErrorCode = 1001
	GeneralForbiddenErrorCode           = 1002
	GeneralInvalidRequestJsonErrorCode  = 1003
)

func GeneralInvalidRequestParamError(original_err string, param_kind string) error {
	err := &KnownError{
		Code:    1003,
		Message: fmt.Sprintf("invalid {param_kind}: {original_err}", original_err, param_kind),
		Range:   "general",
		Name:    "invalid_request_param",
	}
	var out error
	out = err
	out = &HttpStatusCoder{
		error:      out,
		statusCode: 400,
	}
	return out
}

func GeneralForbiddenError() error {
	err := &KnownError{
		Code:    1005,
		Message: "unauthorized",
		Range:   "general",
		Name:    "forbidden",
	}
	var out error
	out = err
	out = &HttpStatusCoder{
		error:      out,
		statusCode: 403,
	}
	return out
}

func GeneralInvalidRequestJsonError(problem string) error {
	err := &KnownError{
		Code:    1008,
		Message: fmt.Sprintf("Invalid JSON: {problem}", problem),
		Range:   "general",
		Name:    "invalid_request_json",
	}
	var out error
	out = err
	out = &HttpStatusCoder{
		error:      out,
		statusCode: 400,
	}
	return out
}
