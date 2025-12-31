package db

import (
	"github.com/bolanosdev/go-snacks/observability/jaeger"
	"github.com/bolanosdev/go-snacks/storage"
	"{{ cookiecutter.group_name }}/{{ cookiecutter.service_name }}/internal/db/sql"
)

type (
	Store struct {
		tracer jaeger.JaegerInterface
		db     sql.PgxPoolConn
		*sql.Queries
	}
)

func NewStore(
	tracer jaeger.JaegerInterface,
	conn sql.PgxPoolConn,
	cache *storage.InMemoryCacheStore,
) Store {
	return Store{
		tracer:  tracer,
		db:      conn,
		Queries: sql.NewQueries(tracer, conn, cache),
	}
}
