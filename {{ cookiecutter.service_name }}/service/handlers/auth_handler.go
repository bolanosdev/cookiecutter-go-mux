package handlers

import (
	"net/http"
	"strings"
	"time"

	"{{cookiecutter.group_name}}/{{cookiecutter.service_name}}/internal/consts/errors"
	"{{cookiecutter.group_name}}/{{cookiecutter.service_name}}/internal/services"
	"{{cookiecutter.group_name}}/{{cookiecutter.service_name}}/internal/utils/encoder"
	"{{cookiecutter.group_name}}/{{cookiecutter.service_name}}/service/entities/request"
	"{{cookiecutter.group_name}}/{{cookiecutter.service_name}}/service/setup/authorization"

	"github.com/bolanosdev/go-snacks/observability/logging"
)

type AuthApi struct {
	sf     services.ServiceFactory
	svc    services.AccountService
	paseto authorization.Maker
}

func NewAuthApi(sf services.ServiceFactory, paseto authorization.Maker) AuthApi {
	return AuthApi{
		sf:     sf,
		svc:    sf.Accounts,
		paseto: paseto,
	}
}

func (h AuthApi) Login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := ctx.Value("logger").(*logging.ContextLogger)
	req, err := encoder.Decode[request.LoginRequest](r)

	if err != nil || (req.Email == "" || req.Password == "") {
		logger.Error().Err(err).Msg("failed to decode login request")
		Error(w, r, http.StatusBadRequest, errors.ErrorBadRequest)
		return
	}

	account, err := h.svc.Login(ctx, req.Email, req.Password)
	if err != nil {
		logger.Error().Err(err).Msg("failed to login")
		Error(w, r, http.StatusNoContent, err)
		return
	}

	access_token, err := h.paseto.CreateToken(account, time.Now())
	if err != nil {
		logger.Error().Err(err).Msg("failed to create access token")
		Error(w, r, http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("access-token", access_token)

	Success(w, r, http.StatusOK, account)
}

func (h AuthApi) SignUp(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := ctx.Value("logger").(*logging.ContextLogger)
	req, err := encoder.Decode[request.SignUpRequest](r)

	if err != nil {
		logger.Error().Err(err).Msg("failed to decode sign up request")
		Error(w, r, http.StatusBadRequest, errors.ErrorBadRequest)
		return
	}

	account, err := h.svc.SignUp(ctx, req.Email, req.Password)
	if err != nil {
		if strings.Contains(err.Error(), "23505") {
			logger.Error().Err(err).Msg("duplicated account")
			Error(w, r, http.StatusBadRequest, errors.ErrorDuplicatedAccount)
		} else {
			logger.Error().Err(err).Msg("failed to sign up")
			Error(w, r, http.StatusBadRequest, errors.ErrorUnexpected)
		}
		return
	}

	access_token, err := h.paseto.CreateToken(account, time.Now())
	if err != nil {
		logger.Error().Err(err).Msg("failed to create access token")
		Error(w, r, http.StatusInternalServerError, errors.ErrorUnexpected)
		return
	}

	w.Header().Set("access-token", access_token)
	Success(w, r, http.StatusOK, "")
}
