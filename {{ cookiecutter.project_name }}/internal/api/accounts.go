package api

import (
	"encoding/json"
	"net/http"

	"{{ cookiecutter.group_name }}/{{ cookiecutter.project_name }}/internal/services"
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
	accounts, _, err := h.svc.GetAll(r.Context())
	if err != nil {
		// c.JSON(http.StatusInternalServerError, errors.ErrorInternalServer)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(accounts)
}
