package db

import (
	"{{ cookiecutter.group_name }}/{{ cookiecutter.service_name }}/internal/cache"
	"{{ cookiecutter.group_name }}/{{ cookiecutter.service_name }}/internal/db/sql"
	"{{ cookiecutter.group_name }}/{{ cookiecutter.service_name }}/internal/utils/obs"
)

type (
	Store struct {
		tracer obs.TracerInterface
		db     sql.PgxPoolConn
		*sql.Queries
	}
)

func NewStore(
	tracer obs.TracerInterface,
	conn sql.PgxPoolConn,
	cache *cache.InMemoryCacheStore,
) Store {
	return Store{
		db:      conn,
		Queries: sql.NewQueries(tracer, conn, cache),
	}
}
