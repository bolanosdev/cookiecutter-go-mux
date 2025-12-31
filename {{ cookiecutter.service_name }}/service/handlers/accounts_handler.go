package handlers

import (
	"net/http"
	"strconv"

	"{{cookiecutter.group_name}}/{{cookiecutter.service_name}}/internal/services"

	"github.com/gorilla/mux"
)

type AccountApi struct {
	BaseHandler
	svc services.AccountService
}

func NewAccountApi(sf services.ServiceFactory) AccountApi {
	return AccountApi{
		BaseHandler: NewBaseHandler(),
		svc:         sf.Accounts,
	}
}

func (h AccountApi) GetAll(w http.ResponseWriter, r *http.Request) {

	accounts, err := h.svc.GetAll(r.Context())
	if err != nil {
		h.Error(w, r, err)
		return
	}

	h.Success(w, r, accounts)
}

func (h AccountApi) GetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	account, err := h.svc.GetByID(r.Context(), id)
	if err != nil {
		h.Error(w, r, err)
		return
	}

	h.Success(w, r, account)
}
