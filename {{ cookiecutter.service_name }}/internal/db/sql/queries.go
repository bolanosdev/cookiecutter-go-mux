package sql

import (
	"context"

	"github.com/bolanosdev/go-snacks/observability/jaeger"
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
	tracer jaeger.JaegerInterface
	db     PgxPoolConn
	cache  *storage.InMemoryCacheStore
}

func NewQueries(
	tracer jaeger.JaegerInterface,
	conn PgxPoolConn,
	cache *storage.InMemoryCacheStore,
) *Queries {
	return &Queries{tracer: tracer, db: conn, cache: cache}
}
