package sql

import (
	"context"

	"{{ cookiecutter.group_name }}/{{ cookiecutter.service_name }}/internal/db/models"

	qb "github.com/bolanosdev/query-builder"
	"github.com/jackc/pgx/v5"
	"go.opentelemetry.io/otel/attribute"
)

func (q *Queries) GetPermissions(ctx context.Context) ([]models.Permission, error) {
	query := `select * from permissions`
	ctx, span := q.tracer.Trace(ctx, "sql.GetPermissions")
	defer span.End()

	span.SetAttributes(
		attribute.String("query", query),
	)

	d := []models.Permission{}
	rows, err := q.db.Query(ctx, query)
	if err != nil {
		return d, err
	}

	accounts, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Permission])
	if err != nil {
		return d, err
	}

	return accounts, nil
}

func (q *Queries) GetPermission(ctx context.Context, conditions ...qb.QueryCondition) (models.Permission, error) {
	ctx, span := q.tracer.Trace(ctx, "sql.GetPermission")
	defer span.End()

	builder := qb.NewQueryBuilder("select * from permissions").Where(conditions...)
	query := builder.Apply()
	values := builder.GetValues()

	span.SetAttributes(
		attribute.String("query", query),
	)

	d := models.Permission{}
	rows, err := q.db.Query(ctx, query, values...)
	if err != nil {
		return d, err
	}

	permission, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.Permission])
	if err != nil {
		return d, err
	}

	return permission, nil
}
