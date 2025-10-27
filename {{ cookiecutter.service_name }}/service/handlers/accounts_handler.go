package handlers

import (
	"net/http"
	"strconv"

	"{{ cookiecutter.group_name }}/{{ cookiecutter.service_name }}/internal/services"

	"github.com/bolanosdev/go-snacks/observability/logging"
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
	ctx := r.Context()
	logger := ctx.Value("logger").(*logging.ContextLogger)

	accounts, err := h.svc.GetAll(ctx)
	if err != nil {
		logger.Error().Err(err).Msg("failed to get all accounts")
		Error(w, r, http.StatusBadRequest, err)
		return
	}

	Success(w, r, http.StatusOK, accounts)
}

func (h AccountApi) GetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	ctx := r.Context()
	logger := ctx.Value("logger").(*logging.ContextLogger)

	account, err := h.svc.GetByID(ctx, id)
	if err != nil {
		logger.Error().Err(err).Msg("failed to get account by id")
		Error(w, r, http.StatusBadRequest, err)

		return
	}

	Success(w, r, http.StatusOK, account)
}
