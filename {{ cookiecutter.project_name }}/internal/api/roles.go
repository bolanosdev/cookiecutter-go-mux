package api

import (
	"encoding/json"
	"net/http"

	"{{ cookiecutter.group_name }}/{{ cookiecutter.project_name }}/internal/services"
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
	roles, _, err := h.svc.GetAll(r.Context())
	if err != nil {
		// c.JSON(http.StatusInternalServerError, errors.ErrorInternalServer)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(roles)
}
