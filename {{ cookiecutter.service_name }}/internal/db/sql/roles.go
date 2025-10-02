package sql

import (
	"context"

	"{{ cookiecutter.group_name }}/{{ cookiecutter.service_name }}/internal/db/models"
	"{{ cookiecutter.group_name }}/{{ cookiecutter.service_name }}/internal/db/queries"

	"github.com/jackc/pgx/v5"
)

func (q *Queries) GetRoles(ctx context.Context) ([]models.Role, error) {
	ctx, span := q.tracer.Trace(ctx, "sql.GetRoles")
	defer span.End()

	d := []models.Role{}
	rows, err := q.db.Query(ctx, queries.GET_ALL_ROLES_QUERY)
	if err != nil {
		return d, err
	}

	roles, err := pgx.CollectRows(rows, pgx.RowToStructByPos[models.Role])
	if err != nil {
		return d, err
	}

	return roles, nil
}

func (q *Queries) GetRoleById(ctx context.Context, id int) (models.Role, error) {
	ctx, span := q.tracer.Trace(ctx, "sql.GetRoleById")
	defer span.End()

	d := models.Role{}
	rows, err := q.db.Query(ctx, queries.GET_ROLES_BY_ID_QUERY, id)
	if err != nil {
		return d, err
	}

	role, err := pgx.CollectOneRow(rows, pgx.RowToStructByPos[models.Role])
	if err != nil {
		return d, err
	}

	return role, nil
}
