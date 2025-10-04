package sql

import (
	"context"

	"{{ cookiecutter.group_name }}/{{ cookiecutter.service_name }}/internal/db/models"
	"{{ cookiecutter.group_name }}/{{ cookiecutter.service_name }}/internal/db/queries"

	qb "github.com/bolanosdev/query-builder"
	"github.com/jackc/pgx/v5"
	"go.opentelemetry.io/otel/attribute"
)

func (q *Queries) GetRoles(ctx context.Context) ([]models.Role, error) {
	query := queries.GET_ALL_ROLES_QUERY
	ctx, span := q.tracer.Trace(ctx, "sql.GetRoles")
	defer span.End()
	span.SetAttributes(
		attribute.String("query", query),
	)

	d := []models.Role{}
	rows, err := q.db.Query(ctx, query)
	if err != nil {
		return d, err
	}

	roles, err := pgx.CollectRows(rows, pgx.RowToStructByPos[models.Role])
	if err != nil {
		return d, err
	}

	return roles, nil
}

func (q *Queries) GetRole(ctx context.Context, conditions ...qb.QueryCondition) (models.Role, error) {
	ctx, span := q.tracer.Trace(ctx, "sql.GetRole")
	defer span.End()

	builder := qb.NewQueryBuilder(queries.GET_ALL_ROLES_QUERY).Where(conditions...)
	query := builder.Apply()
	values := builder.GetValues()

	span.SetAttributes(
		attribute.String("query", query),
	)

	d := models.Role{}
	rows, err := q.db.Query(ctx, query, values...)
	if err != nil {
		return d, err
	}

	role, err := pgx.CollectOneRow(rows, pgx.RowToStructByPos[models.Role])
	if err != nil {
		return d, err
	}

	return role, nil
}
