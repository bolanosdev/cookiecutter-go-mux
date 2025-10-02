package handlers

import (
	"net/http"
	"strconv"

	"{{ cookiecutter.group_name }}/{{ cookiecutter.service_name }}/internal/services"

	"github.com/gorilla/mux"
)

type RoleApi struct {
	svc services.RoleService
}

func NewRoleApi(sf services.ServiceFactory) RoleApi {
	return RoleApi{
		svc: sf.Roles,
	}
}

func (h RoleApi) GetAll(w http.ResponseWriter, r *http.Request) {
	roles, err := h.svc.GetAll(r.Context())
	if err != nil {
		Error(w, r, http.StatusBadRequest, err, "bad request")
		return
	}

	Success(w, r, http.StatusOK, roles)
}

func (h RoleApi) GetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	role, err := h.svc.GetByID(r.Context(), id)
	if err != nil {
		Error(w, r, http.StatusBadRequest, err, "bad request")
		return
	}

	Success(w, r, http.StatusOK, role)
}
