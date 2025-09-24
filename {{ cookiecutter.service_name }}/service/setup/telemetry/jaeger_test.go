package telemetry

import (
	"context"
	"testing"
	"{{cookiecutter.group_name}}/{{cookiecutter.service_name}}/internal/config"

	"github.com/stretchr/testify/require"
)

func TestTelemetryProvider(t *testing.T) {
	ctx := context.Background()
	cfg := config.NewConfigMgr("../../../").Load()

	err := Initialize(ctx, cfg)
	require.NoError(t, err)
}
