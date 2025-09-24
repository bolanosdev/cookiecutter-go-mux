package services

import (
	"context"

	"{{ cookiecutter.group_name }}/{{ cookiecutter.service_name }}/internal/cache"
	"{{ cookiecutter.group_name }}/{{ cookiecutter.service_name }}/internal/db"
	"{{ cookiecutter.group_name }}/{{ cookiecutter.service_name }}/internal/db/models"
	"{{ cookiecutter.group_name }}/{{ cookiecutter.service_name }}/internal/db/sql"
	"{{ cookiecutter.group_name }}/{{ cookiecutter.service_name }}/internal/utils"
)

type PermissionService struct {
	db     sql.PgxPoolConn
	store  db.Store
	cache  *cache.InMemoryCacheStore
	tracer utils.TracerInterface
}

func NewPermissionService(
	db sql.PgxPoolConn,
	store db.Store,
	tracer utils.TracerInterface,
	cache *cache.InMemoryCacheStore,
) PermissionService {
	return PermissionService{
		db:     db,
		store:  store,
		cache:  cache,
		tracer: tracer,
	}
}

func (svc PermissionService) GetAll(c context.Context) ([]models.Permission, error) {
	ctx, span := svc.tracer.Trace(c, "svc.GetAll")
	defer span.End()

	permissions, err := svc.store.GetPermissions(ctx)
	if err != nil {
		return nil, err
	}

	return permissions, nil
}

func (svc PermissionService) GetByID(c context.Context, id int) (*models.Permission, error) {
	ctx, span := svc.tracer.Trace(c, "svc.GetByID")
	defer span.End()

	permission, err := svc.store.GetPermissionById(ctx, id)
	if err != nil {
		return nil, err
	}

	return &permission, nil
}
