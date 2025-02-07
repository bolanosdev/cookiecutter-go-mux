package api

import (
	"net/http"

	"{{ cookiecutter.group_name }}/{{ cookiecutter.project_name }}/db/models"
	"{{ cookiecutter.group_name }}/{{ cookiecutter.project_name }}/internal/services"

	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

type Response struct {
	Account  models.Account   `mapstructure:"account"`
	Accounts []models.Account `mapstructure:"accounts"`
}

type DataApi struct{}

func NewDataApi(sf services.ServiceFactory) DataApi {
	return DataApi{}
}

func (h DataApi) Get(w http.ResponseWriter, r *http.Request) {
	log.Info().Msg("handler")

	res, err := otelhttp.Get(r.Context(), "http://localhost:8081/accounts")
	if err != nil {
		return
	}

	res2, err := otelhttp.Get(r.Context(), "http://localhost:8083/roles")
	if err != nil {
		return
	}

	res3, err := otelhttp.Get(r.Context(), "http://localhost:8084/permissions")
	if err != nil {
		return
	}

	res.Body.Close()
	res2.Body.Close()
	res3.Body.Close()
}
