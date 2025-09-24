package sql

import (
	"context"

	"{{ cookiecutter.group_name }}/{{ cookiecutter.service_name }}/internal/db/models"

	"github.com/jackc/pgx/v5"
)

func (q *Queries) GetPermissions(ctx context.Context) ([]models.Permission, error) {
	ctx, span := q.tracer.Trace(ctx, "sql.GetPermissions")
	defer span.End()

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

	return accounts, nil
}

func (q *Queries) GetPermissionById(ctx context.Context, id int) (models.Permission, error) {
	ctx, span := q.tracer.Trace(ctx, "sql.GetPermissionById")
	defer span.End()

	d := models.Permission{}
	query := `select * from permissions where id = $1;`
	rows, err := q.db.Query(ctx, query, id)
	if err != nil {
		return d, err
	}

	account, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.Permission])
	if err != nil {
		return d, err
	}

	return account, nil
}
