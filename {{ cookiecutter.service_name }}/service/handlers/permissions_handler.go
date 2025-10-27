package handlers

import (
	"net/http"
	"strconv"

	"{{ cookiecutter.group_name }}/{{ cookiecutter.service_name }}/internal/services"

	"github.com/bolanosdev/go-snacks/observability/logging"
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
	ctx := r.Context()
	logger := ctx.Value("logger").(*logging.ContextLogger)

	permissions, err := h.svc.GetAll(ctx)

	if err != nil {
		logger.Error().Err(err).Msg("failed to get all permissions")
		Error(w, r, http.StatusBadRequest, err)
		return
	}

	Success(w, r, http.StatusOK, permissions)
}

func (h PermissionApi) GetByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := ctx.Value("logger").(*logging.ContextLogger)

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	role, err := h.svc.GetByID(ctx, id)
	if err != nil {
		logger.Error().Err(err).Msg("failed to get permission by id")
		Error(w, r, http.StatusBadRequest, err)
		return
	}

	Success(w, r, http.StatusOK, role)
}
