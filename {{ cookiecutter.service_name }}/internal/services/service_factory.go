package services

import (
	"{{ cookiecutter.group_name }}/{{ cookiecutter.service_name }}/internal/config"
	"{{ cookiecutter.group_name }}/{{ cookiecutter.service_name }}/internal/db"
	"{{ cookiecutter.group_name }}/{{ cookiecutter.service_name }}/internal/db/sql"

	"github.com/bolanosdev/go-snacks/observability/jaeger"
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
	tracer jaeger.JaegerInterface,
	cache *storage.InMemoryCacheStore,
) ServiceFactory {
	return ServiceFactory{
		Cfg:         cfg,
		Accounts:    NewAccountService(db, store, tracer, cache),
		Roles:       NewRoleService(db, store, tracer, cache),
		Permissions: NewPermissionService(db, store, tracer, cache),
	}
}

func GetTestServiceFactory(mocker sql.PGXMocker) ServiceFactory {
	cache := storage.NewCacheStore()
	cfg := config.NewConfigMgr("../../").Load()
	tracer := jaeger.NewMockTracer()

	store := db.NewStore(tracer, mocker.Mock, cache)
	sf := NewServiceFactory(cfg, mocker.Mock, store, tracer, cache)

	return sf
}
