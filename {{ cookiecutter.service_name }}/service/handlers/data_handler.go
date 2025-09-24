package handlers

import (
	"net/http"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

type DataApi struct{}

func NewDataApi() DataApi {
	return DataApi{}
}

func (h DataApi) Get(w http.ResponseWriter, r *http.Request) {
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
