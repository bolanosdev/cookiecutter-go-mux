package service

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"{{ cookiecutter.group_name }}/{{ cookiecutter.service_name }}/internal/cache"
	"{{ cookiecutter.group_name }}/{{ cookiecutter.service_name }}/internal/config"
	"{{ cookiecutter.group_name }}/{{ cookiecutter.service_name }}/internal/db"
	"{{ cookiecutter.group_name }}/{{ cookiecutter.service_name }}/internal/services"
	"{{ cookiecutter.group_name }}/{{ cookiecutter.service_name }}/internal/utils"
	"{{ cookiecutter.group_name }}/{{ cookiecutter.service_name }}/service/setup/authorization"
	"{{ cookiecutter.group_name }}/{{ cookiecutter.service_name }}/service/setup/middleware"
	"{{ cookiecutter.group_name }}/{{ cookiecutter.service_name }}/service/setup/telemetry"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
	"go.opentelemetry.io/otel"
)

type MainApp struct {
	ctx        context.Context
	router     *mux.Router
	pool       *pgxpool.Pool
	store      db.Store
	sf         services.ServiceFactory
	cfg        config.AppConfig
	paseto     authorization.Maker
	middleware middleware.Middleware
	logger     zerolog.Logger
}

func New() *MainApp {
	cfg := config.NewConfigMgr(".").Load()
	ctx := context.Background()
	logger := zerolog.New(os.Stderr)

	paseto, err := authorization.NewPasetoMaker(cfg.PASETO)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	pgx_config, err := utils.CreatePGXConfig(cfg.DATABASE)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	conn, pool, err := utils.OpenConnectionPool(ctx, pgx_config)
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
	tracer := utils.NewTracer(tp, cfg)
	store := db.NewStore(tracer, conn, cache)
	sf := services.NewServiceFactory(cfg, conn, store, tracer, cache)
	r := mux.NewRouter()

	middleware := middleware.NewMiddleware(cfg, paseto)

	app := MainApp{
		ctx:        context.Background(),
		router:     r,
		pool:       pool,
		store:      store,
		sf:         sf,
		cfg:        cfg,
		paseto:     paseto,
		middleware: middleware,
		logger:     logger,
	}

	return &app
}

func (app *MainApp) SetMiddleware() *MainApp {
	app.router.Use(app.middleware.Prometheus(app.pool))
	app.router.Use(app.middleware.Logging(app.logger))

	return app
}

func (app *MainApp) Start() {
	log.Println("AppStart")
	cors := middleware.CORS(app.router)
	port := fmt.Sprintf(":%s", app.cfg.SERVICE.PORT)
	log.Fatal(http.ListenAndServe(port, cors))
}
