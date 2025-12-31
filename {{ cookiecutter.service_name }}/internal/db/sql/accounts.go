package sql

import (
	"context"

	"{{ cookiecutter.group_name }}/{{ cookiecutter.service_name }}/internal/db/models"
	"{{ cookiecutter.group_name }}/{{ cookiecutter.service_name }}/internal/db/queries"

	qb "github.com/bolanosdev/query-builder"
	"github.com/jackc/pgx/v5"
)

func (q *Queries) GetAccounts(ctx context.Context) ([]*models.Account, error) {
	query, _ := qb.NewQueryBuilder(queries.GET_ACCOUNTS_QUERY).
		SortBy(qb.Sort("id", qb.SortDesc)).
		Commit()
	ctx = q.tracer.TraceDB(ctx, query, nil)

	rows, err := q.db.Query(ctx, query)
	if err != nil {
		return nil, NewQueryError(err, "sql.query.GetAccounts", nil)
	}

	accounts, err := pgx.CollectRows(rows, pgx.RowToAddrOfStructByName[models.Account])

	if err != nil {
		return nil, NewQueryError(err, "sql.map.GetAccounts", nil)
	}

	return accounts, nil
}

func (q *Queries) GetAccount(ctx context.Context, conditions ...qb.QueryCondition) (*models.Account, error) {
	builder := qb.NewQueryBuilder(queries.GET_ACCOUNTS_QUERY).Where(conditions...)
	query, values := builder.Commit()
	ctx = q.tracer.TraceDB(ctx, query, values)

	rows, err := q.db.Query(ctx, query, values...)
	if err != nil {
		return nil, NewQueryError(err, "sql.query.GetAccount", nil)
	}

	account, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.Account])

	if err != nil {
		return nil, NewQueryError(err, "sql.map.GetAccount", nil)
	}

	return &account, nil
}

func (q *Queries) CreateAccount(ctx context.Context, email string, password string) (*models.Account, error) {
	query := queries.CREATE_ACCOUNT_QUERY
	args := pgx.StrictNamedArgs{
		"email":    email,
		"password": password,
	}

	ctx = q.tracer.TraceDB(ctx, query, args)

	_, err := q.db.Exec(ctx, query, args)

	if err != nil {
		return nil, NewQueryError(err, "sql.CreateAccount", nil)
	}

	return nil, nil
}
