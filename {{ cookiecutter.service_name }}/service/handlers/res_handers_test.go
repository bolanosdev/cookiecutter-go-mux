package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"{{cookiecutter.group_name}}/{{cookiecutter.service_name}}/internal/consts/errors"
	"{{cookiecutter.group_name}}/{{cookiecutter.service_name}}/service/entities/response"

	"github.com/stretchr/testify/require"
)

func TestErrorHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	err := errors.ErrorBadRequest
	msg := "something went wrong"
	rr := httptest.NewRecorder()

	Error(rr, req, 400, err, msg)

	var res response.V1ErrorResponse
	err = json.NewDecoder(rr.Body).Decode(&res)

	require.NoError(t, err)
	require.Equal(t, rr.Code, 400)
	require.Equal(t, res.Code, 400)
	require.Equal(t, response.Message, msg)
}

func TestSuccessHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rr := httptest.NewRecorder()

	result := struct{ foo string }{foo: "bar"}
	Success(rr, req, 200, result)

	var res response.V1SuccessResponse
	err := json.NewDecoder(rr.Body).Decode(&res)

	require.NoError(t, err)
	require.Equal(t, rr.Code, 200)
	require.Equal(t, res.Success, true)
}
