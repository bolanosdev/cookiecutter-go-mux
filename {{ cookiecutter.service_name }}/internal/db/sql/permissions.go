package sql

import (
	"context"

	"{{ cookiecutter.group_name }}/{{ cookiecutter.service_name }}/internal/db/models"
	"{{ cookiecutter.group_name }}/{{ cookiecutter.service_name }}/internal/db/queries"

	qb "github.com/bolanosdev/query-builder"
	"github.com/jackc/pgx/v5"
)

func (q *Queries) GetPermissions(ctx context.Context) ([]*models.Permission, error) {
	query := queries.GET_PERMISSIONS_QUERY
	ctx = q.tracer.TraceDB(ctx, query, nil)

	rows, err := q.db.Query(ctx, query)
	if err != nil {
		return nil, NewQueryError(err, "sql.query.Permissions", nil)
	}

	permissions, err := pgx.CollectRows(rows, pgx.RowToAddrOfStructByName[models.Permission])

	if err != nil {
		return nil, NewQueryError(err, "sql.map.Permissions", nil)
	}

	return permissions, nil
}

func (q *Queries) GetPermission(ctx context.Context, conditions ...qb.QueryCondition) (*models.Permission, error) {
	query, values := qb.NewQueryBuilder(queries.GET_PERMISSIONS_QUERY).
		Where(conditions...).
		Commit()

	ctx = q.tracer.TraceDB(ctx, query, values)

	rows, err := q.db.Query(ctx, query, values...)
	if err != nil {
		return nil, NewQueryError(err, "sql.query.GetPermission", nil)
	}

	permission, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.Permission])
	if err != nil {
		return nil, NewQueryError(err, "sql.map.GetPermission", nil)
	}

	return &permission, nil
}
