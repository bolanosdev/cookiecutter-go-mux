package service

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"{{ cookiecutter.group_name }}/{{ cookiecutter.service_name }}/internal/cache"
	"{{ cookiecutter.group_name }}/{{ cookiecutter.service_name }}/internal/config"
	"{{ cookiecutter.group_name }}/{{ cookiecutter.service_name }}/internal/db"
	"{{ cookiecutter.group_name }}/{{ cookiecutter.service_name }}/internal/services"
	"{{ cookiecutter.group_name }}/{{ cookiecutter.service_name }}/internal/utils/obs"
	"{{ cookiecutter.group_name }}/{{ cookiecutter.service_name }}/internal/utils/pgx"
	"{{ cookiecutter.group_name }}/{{ cookiecutter.service_name }}/service/setup/authorization"
	"{{ cookiecutter.group_name }}/{{ cookiecutter.service_name }}/service/setup/middleware"
	"{{ cookiecutter.group_name }}/{{ cookiecutter.service_name }}/service/setup/telemetry"

	"github.com/bolanosdev/go-snacks/automapper"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.opentelemetry.io/otel"
)

type MainApp struct {
	router     *mux.Router
	pool       *pgxpool.Pool
	store      db.Store
	sf         services.ServiceFactory
	cfg        config.AppConfig
	paseto     authorization.Maker
	middleware middleware.Middleware
	mapper     *automapper.AutoMapper
}

func New() *MainApp {
	cfg := config.NewConfigMgr(".").Load()
	ctx := context.Background()

	paseto, err := authorization.NewPasetoMaker(cfg.PASETO)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	pgx_config, err := pgx.CreatePGXConfig(cfg.DATABASE)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	conn, pool, err := pgx.OpenConnectionPool(ctx, pgx_config)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	err = telemetry.Initialize(ctx, cfg)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	cache := cache.GetCacheStore()
	tp := otel.GetTracerProvider()
	tracer := obs.NewTracer(tp, cfg)
	store := db.NewStore(tracer, conn, cache)

	autoMapper, err := automapper.New().Configure(ApiMappers)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	sf := services.NewServiceFactory(cfg, conn, store, tracer, cache)
	r := mux.NewRouter()

	middleware := middleware.NewMiddleware(cfg, paseto)

	app := MainApp{
		router:     r,
		pool:       pool,
		store:      store,
		sf:         sf,
		cfg:        cfg,
		paseto:     paseto,
		middleware: middleware,
		mapper:     autoMapper,
	}

	return &app
}

func (app *MainApp) SetMiddleware() *MainApp {
	app.router.Use(app.middleware.Prometheus(app.pool))
	app.router.Use(app.middleware.Tracing())
	app.router.Use(app.middleware.Logging())

	return app
}

func (app *MainApp) Start() {
	cors := middleware.CORS(app.router)
	port := fmt.Sprintf(":%s", app.cfg.SERVICE.PORT)

	log.Printf("App.Start%s", port)
	log.Fatal(http.ListenAndServe(port, cors))
}
