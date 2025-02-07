package sql

import (
	"context"

	"{{ cookiecutter.group_name }}/{{ cookiecutter.project_name }}/db/models"
	"{{ cookiecutter.group_name }}/{{ cookiecutter.project_name }}/internal/utils"

	"github.com/jackc/pgx/v5"
)

func (q *Queries) GetAccounts(ctx context.Context) ([]models.Account, error) {
	ctx, span := utils.TracerWithContext(ctx, "db.GetAccounts")
	d := []models.Account{}
	query := `select * from accounts;`
	rows, err := q.db.Query(ctx, query)
	if err != nil {
		return d, err
	}

	accounts, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Account])
	if err != nil {
		return d, err
	}

	span.End()
	return accounts, nil
}

func (q *Queries) GetAccountById(ctx context.Context) (models.Account, error) {
	ctx, span := utils.TracerWithContext(ctx, "db.GetAccountById")

	d := models.Account{}
	query := `select * from accounts where id = 1;`
	rows, err := q.db.Query(ctx, query)
	if err != nil {
		return d, err
	}

	account, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.Account])
	if err != nil {
		return d, err
	}

	span.End()
	return account, nil
}
