package middleware

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func CORS(router *mux.Router) http.Handler {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			http.MethodHead,
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
		},
		AllowedHeaders: []string{"Origin", "Content-Length", "Content-Type", "Authorization", "Access-Token", "Baggage", "Sentry-Trace"},
		ExposedHeaders: []string{"Access-Token"},
	})
	cors := c.Handler(router)
	return cors
}
