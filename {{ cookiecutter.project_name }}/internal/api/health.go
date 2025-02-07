package api

import (
	"net/http"
)

type HealthApi struct{}

func (h HealthApi) Get(w http.ResponseWriter, r *http.Request) {
}
