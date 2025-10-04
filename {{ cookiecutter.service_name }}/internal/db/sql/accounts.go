package sql

import (
	"context"
	"errors"

	"{{ cookiecutter.group_name }}/{{ cookiecutter.service_name }}/internal/db/models"
	"{{ cookiecutter.group_name }}/{{ cookiecutter.service_name }}/internal/db/queries"

	qb "github.com/bolanosdev/query-builder"
	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/otel/attribute"
)

func (q *Queries) GetAccounts(ctx context.Context) ([]models.Account, error) {
	query := qb.NewQueryBuilder(queries.GET_ACCOUNTS_QUERY).SortBy(qb.Sort("a.id", qb.SortDesc)).Apply()
	ctx, span := q.tracer.Trace(ctx, "sql.GetAccounts")
	span.SetAttributes(attribute.String("query", query))
	defer span.End()

	d := []models.Account{}
	rows, err := q.db.Query(ctx, query)
	if err != nil {
		return d, err
	}

	for rows.Next() {
		a := models.Account{}
		r := models.Role{
			Permissions: []models.Permission{},
		}

		if err := rows.Scan(&a.ID, &a.Email, &a.Password, &r.ID, &r.Name, &a.IsActive, &a.CreatedAt, &a.UpdatedAt); err != nil {
			return d, err
		}

		a.Role = r

		d = append(d, a)
	}

	return d, nil
}

func (q *Queries) GetAccount(ctx context.Context, conditions ...qb.QueryCondition) (models.Account, error) {
	ctx, span := q.tracer.Trace(ctx, "sql.GetAccount")
	defer span.End()

	builder := qb.NewQueryBuilder(queries.GET_ACCOUNTS_QUERY).Where(conditions...)
	query := builder.Apply()
	values := builder.GetValues()

	span.SetAttributes(
		attribute.String("query", query),
	)

	d := models.Account{}
	rows, err := q.db.Query(ctx, query, values...)
	if err != nil {
		return d, err
	}

	for rows.Next() {
		a := models.Account{}
		r := models.Role{
			Permissions: []models.Permission{},
		}

		if err := rows.Scan(&a.ID, &a.Email, &a.Password, &r.ID, &r.Name, &a.IsActive, &a.CreatedAt, &a.UpdatedAt); err != nil {
			log.Error().Err(err).Send()
		}
		a.Role = r
		d = a
	}

	if d.ID == 0 {
		return d, errors.New("no records found")
	}

	return d, nil
}

func (q *Queries) CreateAccount(ctx context.Context, email string, password string) (*models.Account, error) {
	query := queries.CREATE_ACCOUNT_QUERY
	ctx, span := q.tracer.Trace(ctx, "sql.CreateAccount")
	defer span.End()

	span.SetAttributes(
		attribute.String("query", query),
		attribute.String("email", email),
		attribute.String("password", password),
	)

	args := pgx.StrictNamedArgs{
		"email":    email,
		"password": password,
	}

	d := models.Account{}

	_, err := q.db.Exec(ctx, query, args)

	if err != nil {
		return nil, err
	}

	return &d, nil
}
