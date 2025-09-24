package middleware

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"
	"{{cookiecutter.group_name}}/{{cookiecutter.service_name}}/api/setup/authorization"
	"{{cookiecutter.group_name}}/{{cookiecutter.service_name}}/internal/config"
	"{{cookiecutter.group_name}}/{{cookiecutter.service_name}}/internal/consts/enums"
	"{{cookiecutter.group_name}}/{{cookiecutter.service_name}}/internal/mocks"
	"{{cookiecutter.group_name}}/{{cookiecutter.service_name}}/internal/models/entities"

	"github.com/stretchr/testify/require"
)

var (
	duration = 24 * time.Hour
	expired  = -24 * time.Hour * 365
	password = "12345678"
	account  = entities.Account{
		ID:        1,
		Tag:       mocks.RandomString(32),
		Firstname: mocks.RandomFirstname(),
		Lastname:  mocks.RandomLastname(),
		Password:  password,
		Email:     mocks.RandomEmail(),
		Role:      enums.User,
	}
)

var (
	cfg       = config.NewConfigMgr("../../../").Load()
	paseto, _ = authorization.NewPasetoMaker(cfg.PASETO)

	m           = NewMiddleware(cfg, paseto)
	mockHandler = func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("authorized"))
	}
	mockSuccessHandlerFunc = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// w.Header().Set("x-error-message", err.Error())
		w.Header().Set("x-status-code", strconv.Itoa(200))
		w.Header().Set("x-trace-id", "2ec68d14-d2e7-4e1a-a21d-8615282bbac7")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("authorized"))
	})
)

func TestAuthorizationMiddlewareNoHeader(t *testing.T) {
	handler := m.Authorize(mockHandler, "test-op")
	rr := httptest.NewRecorder()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	handler.ServeHTTP(rr, req)
	require.Equal(t, rr.Result().StatusCode, 401)
}

func TestAuthorizationMiddlewareWrongHeader(t *testing.T) {
	handler := m.Authorize(mockHandler, "test-op")
	rr := httptest.NewRecorder()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Add("authorizer", "bearer foo123")
	handler.ServeHTTP(rr, req)
	require.Equal(t, rr.Result().StatusCode, 401)
}

func TestAuthorizationMiddlewareWrongHeader2(t *testing.T) {
	handler := m.Authorize(mockHandler, "test-op")
	rr := httptest.NewRecorder()

	token, _ := paseto.CreateToken(&account, time.Now())
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Add("authorization", fmt.Sprintf("bearer%v", token))
	handler.ServeHTTP(rr, req)
	require.Equal(t, rr.Result().StatusCode, 401)
}

func TestAuthorizationMiddlewareExpired(t *testing.T) {
	handler := m.Authorize(mockHandler, "test-op")
	rr := httptest.NewRecorder()

	issued := time.Now().AddDate(0, 0, -366)
	token, _ := paseto.CreateToken(&account, issued)
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Add("authorization", fmt.Sprintf("bearer %v", token))
	handler.ServeHTTP(rr, req)
	require.Equal(t, rr.Result().StatusCode, 401)
}

func TestAuthorizationMiddleware(t *testing.T) {
	handler := m.Authorize(mockHandler, "test-op")
	rr := httptest.NewRecorder()

	token, _ := paseto.CreateToken(&account, time.Now())
	session, _ := paseto.DecryptToken(token)
	log.Printf("session %v", session)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Add("authorization", fmt.Sprintf("bearer %v", token))
	handler.ServeHTTP(rr, req)
	require.Equal(t, rr.Result().StatusCode, 200)
}
