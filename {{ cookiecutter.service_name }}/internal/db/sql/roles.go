package sql

import (
	"context"

	"{{ cookiecutter.group_name }}/{{ cookiecutter.service_name }}/internal/db/models"
	"{{ cookiecutter.group_name }}/{{ cookiecutter.service_name }}/internal/db/queries"

	"github.com/bolanosdev/go-snacks/observability/logging"
	qb "github.com/bolanosdev/query-builder"
	"github.com/jackc/pgx/v5"
	"go.opentelemetry.io/otel/attribute"
)

func (q *Queries) GetRoles(ctx context.Context) ([]models.Role, error) {
	query := queries.GET_ALL_ROLES_QUERY
	ctx, span := q.tracer.Trace(ctx, "sql.GetRoles")
	logger := ctx.Value("logger").(*logging.ContextLogger)
	defer span.End()

	span.SetAttributes(attribute.String("query", query))

	d := []models.Role{}
	rows, err := q.db.Query(ctx, query)
	if err != nil {
		logger.Error().Err(err).Msg("sql.GetRoles")
		return d, err
	}

	roles, err := pgx.CollectRows(rows, pgx.RowToStructByPos[models.Role])
	if err != nil {
		logger.Error().Err(err).Msg("sql.GetRoles")
		return d, err
	}

	return roles, nil
}

func (q *Queries) GetRole(ctx context.Context, conditions ...qb.QueryCondition) (models.Role, error) {
	ctx, span := q.tracer.Trace(ctx, "sql.GetRole")
	logger := ctx.Value("logger").(*logging.ContextLogger)
	defer span.End()

	builder := qb.NewQueryBuilder(queries.GET_ALL_ROLES_QUERY).Where(conditions...)
	query, values := builder.Commit()

	span.SetAttributes(attribute.String("query", query))

	d := models.Role{}
	rows, err := q.db.Query(ctx, query, values...)
	if err != nil {
		logger.Error().Err(err).Msg("sql.GetRole")
		return d, err
	}

	role, err := pgx.CollectOneRow(rows, pgx.RowToStructByPos[models.Role])
	if err != nil {
		logger.Error().Err(err).Msg("sql.GetRole")
		return d, err
	}

	return role, nil
}
