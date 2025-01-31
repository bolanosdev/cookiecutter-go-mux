package test

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestE2E(t *testing.T) {
	assert := assert.New(t)

	endpoint := os.Getenv("API_ADDRESS")

	t.Run("GET /pong", func(t *testing.T) {
		log.Printf("endpoint %v", endpoint)
		res, err := http.Get(fmt.Sprintf("%s/ping", endpoint))
		assert.NoError(err)
		assert.Equal(http.StatusOK, res.StatusCode)
	})
}
