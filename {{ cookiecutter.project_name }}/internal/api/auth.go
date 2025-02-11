package api

import (
	"net/http"
	"time"

	"{{ cookiecutter.group_name }}/{{ cookiecutter.project_name }}/cmd/authorization"
	"{{ cookiecutter.group_name }}/{{ cookiecutter.project_name }}/internal/api/request"
	"{{ cookiecutter.group_name }}/{{ cookiecutter.project_name }}/internal/api/response"
	"{{ cookiecutter.group_name }}/{{ cookiecutter.project_name }}/internal/services"
	"{{ cookiecutter.group_name }}/{{ cookiecutter.project_name }}/internal/utils"

	"github.com/jackc/pgx/v5/pgconn"
)

type AuthApi struct {
	svc    services.AccountService
	paseto authorization.Maker
}

func NewAuthApi(sf services.ServiceFactory, paseto authorization.Maker) AuthApi {
	return AuthApi{
		svc:    sf.Accounts,
		paseto: paseto,
	}
}

func (h AuthApi) Login(w http.ResponseWriter, r *http.Request) {
	req, err := utils.Decode[request.LoginRequest](r)
	if err != nil {
		response.Error(w, r, http.StatusBadRequest, err, "bad request")
	}

	account, err := h.svc.Login(r.Context(), req.Email, req.Password)
	if err != nil {
		if err.Error() == "no records found" || err.Error() == "hashedPassword is not the hash of the given password" {
			response.Error(w, r, http.StatusNoContent, err, "Invalid credentials")
		} else {
			response.Error(w, r, http.StatusBadRequest, err, "bad request")
		}
		return
	}

	access_token, err := h.paseto.CreateToken(account.ID, account.Email, time.Hour*24)
	if err != nil {
		response.Error(w, r, http.StatusBadRequest, err, "bad request")

		return
	}

	w.Header().Set("access-token", access_token)

	response.Success(w, r, http.StatusOK, "")
}

func (h AuthApi) SignUp(w http.ResponseWriter, r *http.Request) {
	req, err := utils.Decode[request.SignUpRequest](r)
	if err != nil {
		response.Error(w, r, http.StatusBadRequest, err, "bad request")
	}

	hash_password, err := utils.HashPassword(req.Password)
	if err != nil {
		response.Error(w, r, http.StatusBadRequest, err, "bad request")
	}

	account, err := h.svc.Create(r.Context(), req.Email, hash_password)
	if err != nil {
		if pqErr, ok := err.(*pgconn.PgError); ok {
			switch pqErr.Code {
			case "23505":
				response.Error(w, r, http.StatusBadRequest, err, "Duplicated account")
			}
		} else {
			response.Error(w, r, http.StatusBadRequest, err, "bad request")
		}
		return
	}

	access_token, err := h.paseto.CreateToken(account.ID, account.Email, time.Hour*24)
	if err != nil {
		response.Error(w, r, http.StatusBadRequest, err, "bad request")

		return
	}

	w.Header().Set("access-token", access_token)

	response.Success(w, r, http.StatusOK, "")
}
