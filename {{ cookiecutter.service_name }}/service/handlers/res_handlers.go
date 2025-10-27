package handlers

import (
	"net/http"

	"{{cookiecutter.group_name}}/{{cookiecutter.service_name}}/internal/utils/encoder"
	"{{cookiecutter.group_name}}/{{cookiecutter.service_name}}/service/entities/response"
)

func Error(w http.ResponseWriter, r *http.Request, status int, external_err error) {
	ext_err_message := ""

	if external_err != nil {
		ext_err_message = external_err.Error()
	}

	response := response.V1ErrorResponse{
		Code:    status,
		Message: ext_err_message,
	}

	encoder.Encode(w, status, response)
}

func Success(w http.ResponseWriter, r *http.Request, status int, result any) {
	response := response.V1SuccessResponse{
		Success: true,
		Result:  result,
	}

	encoder.Encode(w, status, response)
}
