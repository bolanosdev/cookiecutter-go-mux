package handlers

import (
	"net/http"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

type DataApi struct {
	BaseHandler
}

func NewDataApi() DataApi {
	return DataApi{
		BaseHandler: NewBaseHandler(),
	}
}

func (h DataApi) Get(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	res, err := otelhttp.Get(ctx, "http://localhost:8081/accounts")

	if err != nil {
		h.Error(w, r, err)
		return
	}

	res2, err := otelhttp.Get(ctx, "http://localhost:8082/roles")
	if err != nil {
		h.Error(w, r, err)
		return
	}

	res3, err := otelhttp.Get(ctx, "http://localhost:8083/permissions")
	if err != nil {
		h.Error(w, r, err)
		return
	}

	res.Body.Close()
	res2.Body.Close()
	res3.Body.Close()
}
