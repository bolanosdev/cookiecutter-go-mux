package middleware

import (
	"net/http"

	prometheus "github.com/bolanosdev/prometheus-mux-monitor"

	"github.com/jackc/pgx/v5/pgxpool"
)

func (m *Middleware) Prometheus(pool *pgxpool.Pool) func(http.Handler) http.Handler {
	monitor := prometheus.GetMonitor(pool)
	monitor.Use(nil)
	return monitor.Interceptor
}
