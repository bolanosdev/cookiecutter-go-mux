package sql

import (
	"context"

	"{{ cookiecutter.group_name }}/{{ cookiecutter.project_name }}/db/models"
	"{{ cookiecutter.group_name }}/{{ cookiecutter.project_name }}/internal/utils"

	"github.com/jackc/pgx/v5"
)

func (q *Queries) GetRoles(ctx context.Context) ([]models.Role, error) {
	ctx, span := utils.TracerWithContext(ctx, "db.GetRoles")
	d := []models.Role{}
	query := `select * from roles;`
	rows, err := q.db.Query(ctx, query)
	if err != nil {
		return d, err
	}

	accounts, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Role])
	if err != nil {
		return d, err
	}

	span.End()
	return accounts, nil
}

func (q *Queries) GetRoleById(ctx context.Context) (models.Role, error) {
	ctx, span := utils.TracerWithContext(ctx, "db.GetRoleById")

	d := models.Role{}
	query := `select * from roles where id = 1;`
	rows, err := q.db.Query(ctx, query)
	if err != nil {
		return d, err
	}

	account, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.Role])
	if err != nil {
		return d, err
	}

	span.End()
	return account, nil
}
