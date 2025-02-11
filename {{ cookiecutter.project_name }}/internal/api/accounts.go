package api

import (
	"net/http"
	"strconv"

	"{{ cookiecutter.group_name }}/{{ cookiecutter.project_name }}/internal/api/response"
	"{{ cookiecutter.group_name }}/{{ cookiecutter.project_name }}/internal/services"

	"github.com/gorilla/mux"
)

type AccountApi struct {
	svc services.AccountService
}

func NewAccountApi(sf services.ServiceFactory) AccountApi {
	return AccountApi{
		svc: sf.Accounts,
	}
}

func (h AccountApi) GetAll(w http.ResponseWriter, r *http.Request) {
	accounts, err := h.svc.GetAll(r.Context())
	if err != nil {
		response.Error(w, r, http.StatusBadRequest, err, "bad request")
		return
	}

	response.Success(w, r, http.StatusOK, accounts)
}

func (h AccountApi) GetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	account, err := h.svc.GetByID(r.Context(), id)
	if err != nil {
		if err.Error() == "no records found" {
			response.Error(w, r, http.StatusNoContent, err, "")
		} else {
			response.Error(w, r, http.StatusBadRequest, err, "")
		}
		return
	}

	response.Success(w, r, http.StatusOK, account)
}
