package middleware

import (
	"context"
	"net/http"
	"strings"

	"{{cookiecutter.group_name}}/{{cookiecutter.service_name}}/internal/consts/keys"
	"{{cookiecutter.group_name}}/{{cookiecutter.service_name}}/service/setup/authorization"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
)

func (m *Middleware) Authorize(h func(w http.ResponseWriter, r *http.Request), operation_name string) http.Handler {
	tp := otel.GetTracerProvider()

	handler := otelhttp.NewHandler(
		http.HandlerFunc(h),
		operation_name,
		otelhttp.WithTracerProvider(tp),
	)

	return authorize(m.Paseto, handler)
}

func authorize(maker authorization.Maker, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorization_header := r.Header.Get(keys.AuthorizationHeaderKey)

		if len(authorization_header) == 0 {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		fields := strings.Fields(authorization_header)

		if len(fields) != 2 {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		access_token := fields[1]
		session, err := maker.VerifyToken(access_token)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), keys.SessionKey, session)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
