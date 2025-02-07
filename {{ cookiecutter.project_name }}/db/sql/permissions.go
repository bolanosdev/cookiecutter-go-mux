package sql

import (
	"context"

	"{{ cookiecutter.group_name }}/{{ cookiecutter.project_name }}/db/models"
	"{{ cookiecutter.group_name }}/{{ cookiecutter.project_name }}/internal/utils"

	"github.com/jackc/pgx/v5"
)

func (q *Queries) GetPermissions(ctx context.Context) ([]models.Permission, error) {
	ctx, span := utils.TracerWithContext(ctx, "db.GetPermissions")
	d := []models.Permission{}
	query := `select * from permissions;`
	rows, err := q.db.Query(ctx, query)
	if err != nil {
		return d, err
	}

	accounts, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Permission])
	if err != nil {
		return d, err
	}

	span.End()
	return accounts, nil
}

func (q *Queries) GetPermissionById(ctx context.Context) (models.Permission, error) {
	ctx, span := utils.TracerWithContext(ctx, "db.GetPermissionById")

	d := models.Permission{}
	query := `select * from permissions where id = 1;`
	rows, err := q.db.Query(ctx, query)
	if err != nil {
		return d, err
	}

	account, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.Permission])
	if err != nil {
		return d, err
	}

	span.End()
	return account, nil
}
