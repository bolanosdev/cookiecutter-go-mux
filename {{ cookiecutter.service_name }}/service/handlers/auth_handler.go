package handlers

import (
	"net/http"
	"strings"
	"time"

	"{{cookiecutter.group_name}}/{{cookiecutter.service_name}}/internal/consts/errors"
	"{{cookiecutter.group_name}}/{{cookiecutter.service_name}}/internal/services"
	"{{cookiecutter.group_name}}/{{cookiecutter.service_name}}/internal/utils"
	"{{cookiecutter.group_name}}/{{cookiecutter.service_name}}/service/entities/request"
	"{{cookiecutter.group_name}}/{{cookiecutter.service_name}}/service/setup/authorization"
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
	req, err := utils.Decode[request.LoginRequest](r)

	if err != nil || (req.Email == "" || req.Password == "") {
		Error(w, r, http.StatusBadRequest, errors.ErrorBadRequest, "")
		return
	}

	account, err := h.svc.Login(r.Context(), req.Email, req.Password)
	if err != nil {
		Error(w, r, http.StatusNoContent, err, "")
		return
	}

	access_token, err := h.paseto.CreateToken(account, time.Now())
	if err != nil {
		Error(w, r, http.StatusInternalServerError, err, "")
		return
	}

	w.Header().Set("access-token", access_token)

	Success(w, r, http.StatusOK, account)
}

func (h AuthApi) SignUp(w http.ResponseWriter, r *http.Request) {
	req, err := utils.Decode[request.SignUpRequest](r)

	if err != nil {
		Error(w, r, http.StatusBadRequest, err, "bad request")
	}

	account, err := h.svc.SignUp(r.Context(), req.Email, req.Password)
	if err != nil {
		if strings.Contains(err.Error(), "23505") {
			Error(w, r, http.StatusBadRequest, err, errors.ErrorDuplicatedAccount.Error())
		} else {
			Error(w, r, http.StatusBadRequest, err, errors.ErrorUnexpected.Error())
		}
		return
	}

	access_token, err := h.paseto.CreateToken(account, time.Now())
	if err != nil {
		Error(w, r, http.StatusBadRequest, err, errors.ErrorUnexpected.Error())
		return
	}

	w.Header().Set("access-token", access_token)

	Success(w, r, http.StatusOK, "")
}
