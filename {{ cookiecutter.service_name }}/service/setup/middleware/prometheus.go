package middleware

import (
	"net/http"

	"{{ cookiecutter.group_name }}/{{ cookiecutter.service_name }}/service/prometheus"

	"github.com/jackc/pgx/v5/pgxpool"
)

func (m *Middleware) Prometheus(pool *pgxpool.Pool) func(http.Handler) http.Handler {
	monitor := prometheus.GetMonitor(pool)
	monitor.Use(nil)
	return monitor.Interceptor
}
