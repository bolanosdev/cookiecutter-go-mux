package handlers

import (
	"net/http"
	"strconv"

	"{{ cookiecutter.group_name }}/{{ cookiecutter.service_name }}/internal/services"

	"github.com/bolanosdev/go-snacks/observability/logging"
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
	ctx := r.Context()
	logger := ctx.Value("logger").(*logging.ContextLogger)

	roles, err := h.svc.GetAll(ctx)
	if err != nil {
		logger.Error().Err(err).Msg("failed to get all roles")
		Error(w, r, http.StatusBadRequest, err)
		return
	}

	Success(w, r, http.StatusOK, roles)
}

func (h RoleApi) GetByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := ctx.Value("logger").(*logging.ContextLogger)

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	role, err := h.svc.GetByID(ctx, id)

	if err != nil {
		logger.Error().Err(err).Msg("failed to get role by id")
		Error(w, r, http.StatusBadRequest, err)
		return
	}

	Success(w, r, http.StatusOK, role)
}
