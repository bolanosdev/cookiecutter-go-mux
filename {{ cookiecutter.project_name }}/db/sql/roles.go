package sql

import (
	"context"

	"{{ cookiecutter.group_name }}/{{ cookiecutter.project_name }}/db/models"
	"{{ cookiecutter.group_name }}/{{ cookiecutter.project_name }}/db/queries"
	"{{ cookiecutter.group_name }}/{{ cookiecutter.project_name }}/internal/utils"

	"github.com/jackc/pgx/v5"
)

func (q *Queries) GetRoles(ctx context.Context) ([]models.Role, error) {
	ctx, span := utils.TracerWithContext(ctx, "GetRoles")
	d := []models.Role{}
	rows, err := q.db.Query(ctx, queries.GET_ALL_ROLES_QUERY)
	if err != nil {
		return d, err
	}

	roles, err := pgx.CollectRows(rows, pgx.RowToStructByPos[models.Role])
	if err != nil {
		return d, err
	}

	span.End()
	return roles, nil
}

func (q *Queries) GetRoleById(ctx context.Context, id int) (models.Role, error) {
	ctx, span := utils.TracerWithContext(ctx, "GetRoleById")

	d := models.Role{}
	rows, err := q.db.Query(ctx, queries.GET_ROLES_BY_ID_QUERY, id)
	if err != nil {
		return d, err
	}

	role, err := pgx.CollectOneRow(rows, pgx.RowToStructByPos[models.Role])
	if err != nil {
		return d, err
	}

	span.End()
	return role, nil
}
