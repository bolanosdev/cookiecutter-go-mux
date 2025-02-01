package sql

import (
	"context"

	"{{ cookiecutter.group_name }}/{{ cookiecutter.project_name }}/db/models"

	"github.com/jackc/pgx/v5"
)

func (q *Queries) GetAccount(ctx context.Context) ([]models.Account, error) {
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

	return accounts, nil
}
