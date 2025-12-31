package handlers

import (
	"net/http"
	"time"

	"{{cookiecutter.group_name}}/{{cookiecutter.service_name}}/internal/services"
	"{{cookiecutter.group_name}}/{{cookiecutter.service_name}}/internal/utils/encoder"
	"{{cookiecutter.group_name}}/{{cookiecutter.service_name}}/internal/utils/jwt"
	"{{cookiecutter.group_name}}/{{cookiecutter.service_name}}/service/entities/input"
)

type AuthApi struct {
	BaseHandler
	sf     services.ServiceFactory
	svc    services.AccountService
	paseto jwt.Maker
}

func NewAuthApi(sf services.ServiceFactory, paseto jwt.Maker) AuthApi {
	return AuthApi{
		BaseHandler: NewBaseHandler(),
		sf:          sf,
		svc:         sf.Accounts,
		paseto:      paseto,
	}
}

func (h AuthApi) Login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	req, err := encoder.Decode[input.LoginRequest](r)

	if err != nil || (req.Email == "" || req.Password == "") {
		h.Error(w, r, err)
		return
	}

	account, err := h.svc.Login(ctx, req.Email, req.Password)
	if err != nil {
		h.Error(w, r, err)
		return
	}

	access_token, err := h.paseto.CreateToken(account, time.Now())
	if err != nil {
		h.Error(w, r, err)
		return
	}

	w.Header().Set("access-token", access_token)

	h.Success(w, r, account)
}

func (h AuthApi) SignUp(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	req, err := encoder.Decode[input.SignUpRequest](r)

	if err != nil {
		h.Error(w, r, err)
		return
	}

	account, err := h.svc.SignUp(ctx, req.Email, req.Password)
	if err != nil {
		h.Error(w, r, err)
		return
	}

	access_token, err := h.paseto.CreateToken(account, time.Now())
	if err != nil {
		h.Error(w, r, err)
		return
	}

	w.Header().Set("access-token", access_token)
	h.Success(w, r, "")
}
