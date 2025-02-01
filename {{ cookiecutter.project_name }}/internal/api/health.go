package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthApi struct{}

func (h HealthApi) Get(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
