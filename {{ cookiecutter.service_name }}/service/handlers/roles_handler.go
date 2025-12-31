package handlers

import (
	"net/http"
	"strconv"

	"{{cookiecutter.group_name}}/{{cookiecutter.service_name}}/internal/services"

	"github.com/gorilla/mux"
)

type RoleApi struct {
	BaseHandler
	svc services.RoleService
}

func NewRoleApi(sf services.ServiceFactory) RoleApi {
	return RoleApi{
		BaseHandler: NewBaseHandler(),
		svc:         sf.Roles,
	}
}

func (h RoleApi) GetAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	roles, err := h.svc.GetAll(ctx)
	if err != nil {
		h.Error(w, r, err)
		return
	}

	h.Success(w, r, roles)
}

func (h RoleApi) GetByID(w http.ResponseWriter, r *http.Request) {
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
