package middleware

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CORS(r *gin.Engine) {
	cors_config := cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"OPTIONS", "GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization", "Access-Token", "Baggage", "Sentry-Trace"},
		ExposeHeaders:    []string{"Access-Token"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}
	r.Use(cors.New(cors_config))
}
