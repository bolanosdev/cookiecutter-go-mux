package api

import (
	"net/http"
	"strconv"

	"{{ cookiecutter.group_name }}/{{ cookiecutter.project_name }}/internal/api/response"
	"{{ cookiecutter.group_name }}/{{ cookiecutter.project_name }}/internal/services"

	"github.com/gorilla/mux"
)

type PermissionApi struct {
	svc services.PermissionService
}

func NewPermissionApi(sf services.ServiceFactory) PermissionApi {
	return PermissionApi{
		svc: sf.Permissions,
	}
}

func (h PermissionApi) GetAll(w http.ResponseWriter, r *http.Request) {
	permissions, err := h.svc.GetAll(r.Context())
	if err != nil {
		response.Error(w, r, http.StatusBadRequest, err, "bad request")
		return
	}

	response.Success(w, r, http.StatusOK, permissions)
}

func (h PermissionApi) GetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	role, err := h.svc.GetByID(r.Context(), id)
	if err != nil {
		response.Error(w, r, http.StatusBadRequest, err, "bad request")
		return
	}

	response.Success(w, r, http.StatusOK, role)
}
