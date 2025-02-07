package api

import (
	"encoding/json"
	"net/http"

	"{{ cookiecutter.group_name }}/{{ cookiecutter.project_name }}/internal/services"
)

type PermissionApi struct {
	ac_svc services.AccountService
	r_svc  services.RoleService
	p_svc  services.PermissionService
}

func NewPermissionApi(sf services.ServiceFactory) PermissionApi {
	return PermissionApi{
		ac_svc: sf.Accounts,
		r_svc:  sf.Roles,
		p_svc:  sf.Permissions,
	}
}

func (h PermissionApi) GetAll(w http.ResponseWriter, r *http.Request) {
	permissions, _, err := h.p_svc.GetAll(r.Context())
	if err != nil {
		// c.JSON(http.StatusInternalServerError, errors.ErrorInternalServer)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(permissions)
}
