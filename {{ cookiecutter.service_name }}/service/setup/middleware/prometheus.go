package middleware

import (
	"net/http"

	prometheus "github.com/bolanosdev/prometheus-mux-monitor"
)

func (m *Middleware) Prometheus() func(http.Handler) http.Handler {
	monitor := prometheus.GetMonitor()
	return monitor.Interceptor
}
