package services

import (
	"context"

	"{{ cookiecutter.group_name }}/{{ cookiecutter.service_name }}/internal/db"
	"{{ cookiecutter.group_name }}/{{ cookiecutter.service_name }}/internal/db/models"
	"{{ cookiecutter.group_name }}/{{ cookiecutter.service_name }}/internal/db/sql"

	"github.com/bolanosdev/go-snacks/observability/jaeger"
	"github.com/bolanosdev/go-snacks/storage"
	qb "github.com/bolanosdev/query-builder"
	"github.com/pkg/errors"
)

type RoleService struct {
	db     sql.PgxPoolConn
	store  db.Store
	cache  *storage.InMemoryCacheStore
	tracer jaeger.JaegerInterface
}

func NewRoleService(
	db sql.PgxPoolConn,
	store db.Store,
	tracer jaeger.JaegerInterface,
	cache *storage.InMemoryCacheStore,
) RoleService {
	return RoleService{
		db:     db,
		store:  store,
		cache:  cache,
		tracer: tracer,
	}
}

func (svc RoleService) GetAll(c context.Context) ([]*models.Role, error) {
	ctx := svc.tracer.TraceFunc(c)

	roles, err := svc.store.GetRoles(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "fail to retrieve roles")
	}

	return roles, nil
}

func (svc RoleService) GetByID(c context.Context, id int) (*models.Role, error) {
	ctx := svc.tracer.TraceFunc(c)
	permission, err := svc.store.GetRole(ctx, qb.ByIntColumn("id", []int{id}))

	if err != nil {
		return nil, errors.Wrap(err, "fail to retrieve role by id")
	}

	return permission, nil
}
