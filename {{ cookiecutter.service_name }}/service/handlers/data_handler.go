package handlers

import (
	"net/http"

	"{{ cookiecutter.group_name }}/{{ cookiecutter.service_name }}/internal/consts/errors"

	"github.com/bolanosdev/go-snacks/observability/logging"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

type DataApi struct{}

func NewDataApi() DataApi {
	return DataApi{}
}

func (h DataApi) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := ctx.Value("logger").(*logging.ContextLogger)

	res, err := otelhttp.Get(r.Context(), "http://localhost:8081/accounts")

	if err != nil {
		logger.Error().Err(err).Msg("failed to get accounts")
		Error(w, r, http.StatusNotFound, errors.ErrorUnexpected)
		return
	}

	res2, err := otelhttp.Get(r.Context(), "http://localhost:8082/roles")
	if err != nil {
		logger.Error().Err(err).Msg("failed to get roles")
		Error(w, r, http.StatusNotFound, errors.ErrorUnexpected)
		return
	}

	res3, err := otelhttp.Get(r.Context(), "http://localhost:8083/permissions")
	if err != nil {
		logger.Error().Err(err).Msg("failed to get permissions")
		Error(w, r, http.StatusNotFound, errors.ErrorUnexpected)
		return
	}

	res.Body.Close()
	res2.Body.Close()
	res3.Body.Close()
}
