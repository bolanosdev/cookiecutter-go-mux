package services

import (
	"context"

	"{{ cookiecutter.group_name }}/{{ cookiecutter.service_name }}/internal/cache"
	"{{ cookiecutter.group_name }}/{{ cookiecutter.service_name }}/internal/config"
	"{{ cookiecutter.group_name }}/{{ cookiecutter.service_name }}/internal/db"
	"{{ cookiecutter.group_name }}/{{ cookiecutter.service_name }}/internal/db/sql"
	"{{ cookiecutter.group_name }}/{{ cookiecutter.service_name }}/internal/utils"

	"go.opentelemetry.io/otel"
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
	tracer utils.TracerInterface,
	cache *cache.InMemoryCacheStore,
) ServiceFactory {
	return ServiceFactory{
		Cfg:         cfg,
		Accounts:    NewAccountService(db, store, tracer, cache),
		Roles:       NewRoleService(db, store, tracer, cache),
		Permissions: NewPermissionService(db, store, tracer, cache),
	}
}

func GetRealServiceFactory(ctx context.Context) ServiceFactory {
	cache := cache.NewCacheStore()
	cfg := config.NewConfigMgr("../../").Load()
	tp := otel.GetTracerProvider()
	tracer := utils.NewTracer(tp, cfg)

	pgx_config, _ := utils.CreatePGXConfig(cfg.DATABASE)
	conn, _, _ := utils.OpenConnectionPool(ctx, pgx_config)
	store := db.NewStore(tracer, conn, cache)
	sf := NewServiceFactory(cfg, conn, store, tracer, cache)

	return sf
}

func GetTestServiceFactory(mocker sql.PGXMocker) ServiceFactory {
	cache := cache.NewCacheStore()
	cfg := config.NewConfigMgr("../../").Load()
	tp := otel.GetTracerProvider()
	tracer := utils.NewTracer(tp, cfg)

	store := db.NewStore(tracer, mocker.Mock, cache)
	sf := NewServiceFactory(cfg, mocker.Mock, store, tracer, cache)

	return sf
}
