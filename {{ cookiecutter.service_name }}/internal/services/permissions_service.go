package services

import (
	"context"

	"{{ cookiecutter.group_name }}/{{ cookiecutter.service_name }}/internal/cache"
	"{{ cookiecutter.group_name }}/{{ cookiecutter.service_name }}/internal/db"
	"{{ cookiecutter.group_name }}/{{ cookiecutter.service_name }}/internal/db/models"
	"{{ cookiecutter.group_name }}/{{ cookiecutter.service_name }}/internal/db/sql"
	"{{ cookiecutter.group_name }}/{{ cookiecutter.service_name }}/internal/utils/obs"

	"github.com/bolanosdev/go-snacks/observability/logging"
	qb "github.com/bolanosdev/query-builder"
)

type PermissionService struct {
	db     sql.PgxPoolConn
	store  db.Store
	cache  *cache.InMemoryCacheStore
	tracer obs.TracerInterface
}

func NewPermissionService(
	db sql.PgxPoolConn,
	store db.Store,
	tracer obs.TracerInterface,
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
	ctx, span := svc.tracer.Trace(c, "svc.permission.GetAll")
	logger := c.Value("logger").(*logging.ContextLogger)
	defer span.End()

	permissions, err := svc.store.GetPermissions(ctx)
	if err != nil {
		logger.Error().Err(err).Msg("svc.GetAll")
		return nil, err
	}

	return permissions, nil
}

func (svc PermissionService) GetByID(c context.Context, id int) (*models.Permission, error) {
	ctx, span := svc.tracer.Trace(c, "svc.permission.GetByID")
	logger := c.Value("logger").(*logging.ContextLogger)
	defer span.End()

	permission, err := svc.store.GetPermission(ctx, qb.ByIntColumn("p.id", []int{id}))
	if err != nil {
		logger.Error().Err(err).Msg("svc.GetByID")
		return nil, err
	}

	return &permission, nil
}
