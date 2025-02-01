package sql

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type Queries struct {
	db *pgxpool.Conn
}

func NewQueries(conn *pgxpool.Conn) *Queries {
	return &Queries{db: conn}
}
