package sql

import (
	"context"

	"{{ cookiecutter.group_name }}/{{ cookiecutter.service_name }}/internal/db/models"
	"{{ cookiecutter.group_name }}/{{ cookiecutter.service_name }}/internal/db/queries"

	qb "github.com/bolanosdev/query-builder"
	"github.com/jackc/pgx/v5"
)

func (q *Queries) GetRoles(ctx context.Context) ([]*models.Role, error) {
	query := queries.GET_ROLES_QUERY
	ctx = q.tracer.TraceDB(ctx, query, nil)
	rows, err := q.db.Query(ctx, query)
	if err != nil {
		return nil, NewQueryError(err, "sql.query.GetRoles", nil)
	}

	roles, err := pgx.CollectRows(rows, pgx.RowToAddrOfStructByName[models.Role])
	if err != nil {
		return nil, NewQueryError(err, "sql.map.GetRoles", nil)
	}

	return roles, nil
}

func (q *Queries) GetRole(ctx context.Context, conditions ...qb.QueryCondition) (*models.Role, error) {
	query, values := qb.NewQueryBuilder(queries.GET_ROLES_QUERY).Where(conditions...).Commit()
	ctx = q.tracer.TraceDB(ctx, query, values)

	rows, err := q.db.Query(ctx, query, values...)
	if err != nil {
		return nil, NewQueryError(err, "sql.query.GetRole", nil)
	}

	role, err := pgx.CollectOneRow(rows, pgx.RowToStructByPos[models.Role])
	if err != nil {
		return nil, NewQueryError(err, "sql.map.GetRole", nil)
	}

	return &role, nil
}
