package services

import (
	"context"

	"{{ cookiecutter.group_name }}/{{ cookiecutter.service_name }}/internal/config"
	"{{ cookiecutter.group_name }}/{{ cookiecutter.service_name }}/internal/db"
	"{{ cookiecutter.group_name }}/{{ cookiecutter.service_name }}/internal/db/sql"
	"{{ cookiecutter.group_name }}/{{ cookiecutter.service_name }}/internal/utils/obs"
	"{{ cookiecutter.group_name }}/{{ cookiecutter.service_name }}/internal/utils/pgx"

	"github.com/bolanosdev/go-snacks/storage"
)

type ServiceFactory struct {
	Cfg         config.AppConfig
	Accounts    AccountService
	Roles       RoleService
	Permissions PermissionService
}

func NewServiceFactory(
	cfg config.AppConfig,
	db sql.PgxPoolConn,
	store db.Store,
	tracer obs.TracerInterface,
	cache *storage.InMemoryCacheStore,
) ServiceFactory {
	return ServiceFactory{
		Cfg:         cfg,
		Accounts:    NewAccountService(db, store, tracer, cache),
		Roles:       NewRoleService(db, store, tracer, cache),
		Permissions: NewPermissionService(db, store, tracer, cache),
	}
}

func GetRealServiceFactory(ctx context.Context) ServiceFactory {
	cache := storage.NewCacheStore()
	cfg := config.NewConfigMgr("../../").Load()
	tracer := obs.NewTracer(ctx, cfg.SERVICE.NAME, cfg.OBSERVABILITY)

	conn, _, _ := pgx.OpenConnectionPool(ctx, cfg.DATABASE)
	store := db.NewStore(tracer, conn, cache)
	sf := NewServiceFactory(cfg, conn, store, tracer, cache)

	return sf
}

func GetTestServiceFactory(mocker sql.PGXMocker) ServiceFactory {
	cache := storage.NewCacheStore()
	cfg := config.NewConfigMgr("../../").Load()
	tracer := obs.NewTracer(mocker.Ctx, cfg.SERVICE.NAME, cfg.OBSERVABILITY)

	store := db.NewStore(tracer, mocker.Mock, cache)
	sf := NewServiceFactory(cfg, mocker.Mock, store, tracer, cache)

	return sf
}
