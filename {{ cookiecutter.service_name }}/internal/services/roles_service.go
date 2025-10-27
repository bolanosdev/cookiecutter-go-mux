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

type RoleService struct {
	db     sql.PgxPoolConn
	store  db.Store
	cache  *cache.InMemoryCacheStore
	tracer obs.TracerInterface
}

func NewRoleService(
	db sql.PgxPoolConn,
	store db.Store,
	tracer obs.TracerInterface,
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
	ctx, span := svc.tracer.Trace(c, "svc.role.GetAll")
	logger := c.Value("logger").(*logging.ContextLogger)
	defer span.End()

	roles, err := svc.store.GetRoles(ctx)
	if err != nil {
		logger.Error().Err(err).Msg("svc.GetAll")
		return nil, err
	}

	return roles, nil
}

func (svc RoleService) GetByID(c context.Context, id int) (*models.Role, error) {
	ctx, span := svc.tracer.Trace(c, "svc.role.GetByID")
	logger := c.Value("logger").(*logging.ContextLogger)
	defer span.End()

	role, err := svc.store.GetRole(ctx, qb.ByIntColumn("r.id", []int{id}))
	if err != nil {
		logger.Error().Err(err).Msg("svc.GetByID")
		return nil, err
	}

	return &role, nil
}
