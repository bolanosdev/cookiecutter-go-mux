package sql

import (
	"context"
	"errors"

	"{{ cookiecutter.group_name }}/{{ cookiecutter.service_name }}/internal/db/models"
	"{{ cookiecutter.group_name }}/{{ cookiecutter.service_name }}/internal/db/queries"

	"github.com/bolanosdev/go-snacks/observability/logging"
	qb "github.com/bolanosdev/query-builder"
	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/otel/attribute"
)

func (q *Queries) GetAccounts(ctx context.Context) ([]models.Account, error) {
	ctx, span := q.tracer.Trace(ctx, "sql.GetAccounts")
	logger := ctx.Value("logger").(*logging.ContextLogger)
	defer span.End()

	query, _ := qb.NewQueryBuilder(queries.GET_ACCOUNTS_QUERY).
		SortBy(qb.Sort("a.id", qb.SortDesc)).
		Commit()

	span.SetAttributes(attribute.String("query", query))

	d := []models.Account{}
	rows, err := q.db.Query(ctx, query)
	if err != nil {
		logger.Error().Err(err).Msg("sql.GetAccounts")
		return d, err
	}

	for rows.Next() {
		a := models.Account{}
		r := models.Role{
			Permissions: []models.Permission{},
		}

		if err := rows.Scan(&a.ID, &a.Email, &a.Password, &r.ID, &r.Name, &a.IsActive, &a.CreatedAt, &a.UpdatedAt); err != nil {
			logger.Error().Err(err).Msg("sql.GetAccounts")
			return d, err
		}

		a.Role = r

		d = append(d, a)
	}

	return d, nil
}

func (q *Queries) GetAccount(ctx context.Context, conditions ...qb.QueryCondition) (models.Account, error) {
	ctx, span := q.tracer.Trace(ctx, "sql.GetAccount")
	logger := ctx.Value("logger").(*logging.ContextLogger)
	defer span.End()

	builder := qb.NewQueryBuilder(queries.GET_ACCOUNTS_QUERY).Where(conditions...)
	query, values := builder.Commit()

	span.SetAttributes(attribute.String("query", query))

	d := models.Account{}
	rows, err := q.db.Query(ctx, query, values...)
	if err != nil {
		logger.Error().Err(err).Msg("sql.GetAccount")
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
		logger.Error().Msg("sql.GetAccount")
		return d, errors.New("no records found")
	}

	return d, nil
}

func (q *Queries) CreateAccount(ctx context.Context, email string, password string) (*models.Account, error) {
	query := queries.CREATE_ACCOUNT_QUERY
	ctx, span := q.tracer.Trace(ctx, "sql.CreateAccount")
	logger := ctx.Value("logger").(*logging.ContextLogger)
	defer span.End()

	span.SetAttributes(attribute.String("query", query))

	args := pgx.StrictNamedArgs{
		"email":    email,
		"password": password,
	}

	d := models.Account{}

	_, err := q.db.Exec(ctx, query, args)

	if err != nil {
		logger.Error().Err(err).Msg("sql.CreateAccount")
		return nil, err
	}

	return &d, nil
}
