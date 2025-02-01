package integration

import (
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestE2E(t *testing.T) {
	assert := assert.New(t)
	api_address := os.Getenv("API_ADDRESS")

	t.Run("GET /health", func(t *testing.T) {
		res, err := http.Get(fmt.Sprintf("%s/health", api_address))
		assert.NoError(err)
		assert.Equal(http.StatusOK, res.StatusCode)
	})
}
