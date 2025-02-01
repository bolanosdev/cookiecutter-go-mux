package db

import (
	"{{ cookiecutter.group_name }}/{{ cookiecutter.project_name }}/db/sql"

	"github.com/jackc/pgx/v5/pgxpool"
)

type (
	Store struct {
		db *pgxpool.Conn
		*sql.Queries
	}
)

func NewStore(conn *pgxpool.Conn) Store {
	return Store{
		db:      conn,
		Queries: sql.NewQueries(conn),
	}
}
