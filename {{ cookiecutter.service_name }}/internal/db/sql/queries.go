package sql

import (
	"context"

	"{{ cookiecutter.group_name }}/{{ cookiecutter.service_name }}/internal/utils/obs"

	"github.com/bolanosdev/go-snacks/storage"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type PgxPoolConn interface {
	QueryRow(context.Context, string, ...interface{}) pgx.Row
	Query(context.Context, string, ...interface{}) (pgx.Rows, error)
	Exec(context.Context, string, ...interface{}) (pgconn.CommandTag, error)
	Begin(ctx context.Context) (pgx.Tx, error)
	Release()
}

type Queries struct {
	tracer obs.TracerInterface
	db     PgxPoolConn
	cache  *storage.InMemoryCacheStore
}

func NewQueries(
	tracer obs.TracerInterface,
	conn PgxPoolConn,
	cache *storage.InMemoryCacheStore,
) *Queries {
	return &Queries{tracer: tracer, db: conn, cache: cache}
}
