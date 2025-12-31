package handlers

import (
	"net/http"
	"strconv"

	"{{cookiecutter.group_name}}/{{cookiecutter.service_name}}/internal/services"

	"github.com/gorilla/mux"
)

type PermissionApi struct {
	BaseHandler
	svc services.PermissionService
}

func NewPermissionApi(sf services.ServiceFactory) PermissionApi {
	return PermissionApi{
		BaseHandler: NewBaseHandler(),
		svc:         sf.Permissions,
	}
}

func (h PermissionApi) GetAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	permissions, err := h.svc.GetAll(ctx)

	if err != nil {
		h.Error(w, r, err)
		return
	}

	h.Success(w, r, permissions)
}

func (h PermissionApi) GetByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	role, err := h.svc.GetByID(ctx, id)
	if err != nil {
		h.Error(w, r, err)
		return
	}

	h.Success(w, r, role)
}
