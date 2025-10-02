package sql

import (
	"context"
	"errors"

	"{{ cookiecutter.group_name }}/{{ cookiecutter.service_name }}/internal/db/models"
	"{{ cookiecutter.group_name }}/{{ cookiecutter.service_name }}/internal/db/queries"

	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog/log"
)

func (q *Queries) GetAccounts(ctx context.Context) ([]models.Account, error) {
	ctx, span := q.tracer.Trace(ctx, "sql.GetAccounts")
	defer span.End()

	d := []models.Account{}
	rows, err := q.db.Query(ctx, queries.GET_ACCOUNTS_QUERY)
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

func (q *Queries) GetAccountByEmail(ctx context.Context, email string) (models.Account, error) {
	ctx, span := q.tracer.Trace(ctx, "sql.GetAccountByEmail")
	defer span.End()

	d := models.Account{}
	rows, err := q.db.Query(ctx, queries.GET_ACCOUNTS_BY_EMAIL_QUERY, email)
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

func (q *Queries) GetAccountById(ctx context.Context, id int) (models.Account, error) {
	ctx, span := q.tracer.Trace(ctx, "sql.GetAccountById")
	defer span.End()

	d := models.Account{}
	rows, err := q.db.Query(ctx, queries.GET_ACCOUNTS_BY_ID_QUERY, id)
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
	ctx, span := q.tracer.Trace(ctx, "sql.CreateAccount")
	defer span.End()
	d := models.Account{}

	_, err := q.db.Exec(ctx, `
      insert into accounts (email, password) 
      values (@email, @password) 
      returning *;
    `, pgx.StrictNamedArgs{
		"email":    email,
		"password": password,
	})
	if err != nil {
		return nil, err
	}

	return &d, nil
}
