package sql

import (
	"context"

	"{{ cookiecutter.group_name }}/{{ cookiecutter.service_name }}/internal/db/models"

	"github.com/bolanosdev/go-snacks/observability/logging"
	qb "github.com/bolanosdev/query-builder"
	"github.com/jackc/pgx/v5"
	"go.opentelemetry.io/otel/attribute"
)

func (q *Queries) GetPermissions(ctx context.Context) ([]models.Permission, error) {
	query := `select * from permissions`
	ctx, span := q.tracer.Trace(ctx, "sql.GetPermissions")
	logger := ctx.Value("logger").(*logging.ContextLogger)
	defer span.End()

	span.SetAttributes(attribute.String("query", query))

	d := []models.Permission{}
	rows, err := q.db.Query(ctx, query)
	if err != nil {
		logger.Error().Err(err).Msg("sql.GetPermissions")
		return d, err
	}

	accounts, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Permission])
	if err != nil {
		logger.Error().Err(err).Msg("sql.GetPermissions")
		return d, err
	}

	return accounts, nil
}

func (q *Queries) GetPermission(ctx context.Context, conditions ...qb.QueryCondition) (models.Permission, error) {
	ctx, span := q.tracer.Trace(ctx, "sql.GetPermission")
	logger := ctx.Value("logger").(*logging.ContextLogger)
	defer span.End()

	builder := qb.NewQueryBuilder("select * from permissions").Where(conditions...)
	query, values := builder.Commit()

	span.SetAttributes(attribute.String("query", query))

	d := models.Permission{}
	rows, err := q.db.Query(ctx, query, values...)
	if err != nil {
		logger.Error().Err(err).Msg("sql.GetPermission")
		return d, err
	}

	permission, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.Permission])
	if err != nil {
		logger.Error().Err(err).Msg("sql.GetPermission")
		return d, err
	}

	return permission, nil
}
