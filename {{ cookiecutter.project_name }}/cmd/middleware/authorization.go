package middleware

import (
	"log"
	"net/http"
	"strings"

	"{{ cookiecutter.group_name }}/{{ cookiecutter.project_name }}/cmd/authorization"
	"{{ cookiecutter.group_name }}/{{ cookiecutter.project_name }}/cmd/cfg"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
)

const (
	AuthorizationHeaderKey  = "authorization"
	AuthorizationBearerKey  = "bearer"
	AuthorizationPayloadKey = "authorization_payload"
)

type Middleware = func(next http.Handler) http.Handler

func Authorization(h func(w http.ResponseWriter, r *http.Request), operation_name string) http.Handler {
	tp := otel.GetTracerProvider()

	handler := otelhttp.NewHandler(
		http.HandlerFunc(h),
		operation_name,
		otelhttp.WithTracerProvider(tp),
	)

	return authorize(handler)
}

func authorize(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		config := cfg.Load(".")
		token_maker, err := authorization.NewPasetoMaker(config.PASETO.TOKEN_SYMETRIC_KEY)
		if err != nil {
			log.Fatal(err)
			w.WriteHeader(http.StatusUnauthorized)
		}

		authorization_header := r.Header.Get(AuthorizationHeaderKey)

		if len(authorization_header) == 0 {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		fields := strings.Fields(authorization_header)
		if len(fields) != 2 {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		authorization_type := strings.ToLower(fields[0])
		if authorization_type != AuthorizationBearerKey {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		access_token := fields[1]
		_, err = token_maker.VerifyToken(access_token)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
