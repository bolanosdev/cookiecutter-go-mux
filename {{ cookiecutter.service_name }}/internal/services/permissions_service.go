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

type PermissionService struct {
	db     sql.PgxPoolConn
	store  db.Store
	cache  *storage.InMemoryCacheStore
	tracer jaeger.JaegerInterface
}

func NewPermissionService(
	db sql.PgxPoolConn,
	store db.Store,
	tracer jaeger.JaegerInterface,
	cache *storage.InMemoryCacheStore,
) PermissionService {
	return PermissionService{
		db:     db,
		store:  store,
		cache:  cache,
		tracer: tracer,
	}
}

func (svc PermissionService) GetAll(c context.Context) ([]*models.Permission, error) {
	ctx := svc.tracer.TraceFunc(c)

	permissions, err := svc.store.GetPermissions(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "fail to retrieve permissions")
	}

	return permissions, nil
}

func (svc PermissionService) GetByID(c context.Context, id int) (*models.Permission, error) {
	ctx := svc.tracer.TraceFunc(c)

	permission, err := svc.store.GetPermission(ctx, qb.ByIntColumn("id", []int{id}))
	if err != nil {
		return nil, errors.Wrap(err, "fail to retrieve permission by id")
	}

	return permission, nil
}
