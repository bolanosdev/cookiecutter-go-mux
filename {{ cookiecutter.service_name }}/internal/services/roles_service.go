package services

import (
	"context"

	"{{ cookiecutter.group_name }}/{{ cookiecutter.service_name }}/internal/cache"
	"{{ cookiecutter.group_name }}/{{ cookiecutter.service_name }}/internal/db"
	"{{ cookiecutter.group_name }}/{{ cookiecutter.service_name }}/internal/db/models"
	"{{ cookiecutter.group_name }}/{{ cookiecutter.service_name }}/internal/db/sql"
	"{{ cookiecutter.group_name }}/{{ cookiecutter.service_name }}/internal/utils"
)

type RoleService struct {
	db     sql.PgxPoolConn
	store  db.Store
	cache  *cache.InMemoryCacheStore
	tracer utils.TracerInterface
}

func NewRoleService(
	db sql.PgxPoolConn,
	store db.Store,
	tracer utils.TracerInterface,
	cache *cache.InMemoryCacheStore,
) RoleService {
	return RoleService{
		db:     db,
		store:  store,
		cache:  cache,
		tracer: tracer,
	}
}

func (svc RoleService) GetAll(c context.Context) ([]models.Role, error) {
	ctx, span := svc.tracer.Trace(c, "svc.GetAll")
	defer span.End()

	roles, err := svc.store.GetRoles(ctx)
	if err != nil {
		return nil, err
	}

	return roles, nil
}

func (svc RoleService) GetByID(c context.Context, id int) (*models.Role, error) {
	ctx, span := svc.tracer.Trace(c, "svc.GetByID")
	defer span.End()

	role, err := svc.store.GetRoleById(ctx, id)
	if err != nil {
		return nil, err
	}

	return &role, nil
}
