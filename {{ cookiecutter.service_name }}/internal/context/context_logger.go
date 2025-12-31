package context

import (
	"context"

	"{{ cookiecutter.group_name }}/{{ cookiecutter.service_name }}/internal/consts"

	"github.com/bolanosdev/go-snacks/observability/logging"
)

func GetLogger(ctx context.Context) logging.ContextLogger {
	if logger, ok := ctx.Value(consts.LoggerKey).(*logging.ContextLogger); ok {
		return *logger
	}
	return logging.ContextLogger{}
}
