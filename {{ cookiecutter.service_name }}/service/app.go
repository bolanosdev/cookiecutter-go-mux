package service

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"{{cookiecutter.group_name}}/{{cookiecutter.service_name}}/internal/config"
	"{{cookiecutter.group_name}}/{{cookiecutter.service_name}}/internal/db"
	"{{cookiecutter.group_name}}/{{cookiecutter.service_name}}/internal/services"
	"{{cookiecutter.group_name}}/{{cookiecutter.service_name}}/internal/utils/jwt"
	"{{cookiecutter.group_name}}/{{cookiecutter.service_name}}/internal/utils/obs"
	"{{cookiecutter.group_name}}/{{cookiecutter.service_name}}/internal/utils/pgx"
	"{{cookiecutter.group_name}}/{{cookiecutter.service_name}}/service/middleware"

	"github.com/bolanosdev/go-snacks/automapper"
	"github.com/bolanosdev/go-snacks/storage"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
)

type MainApp struct {
	router     *mux.Router
	pool       *pgxpool.Pool
	store      db.Store
	sf         services.ServiceFactory
	cfg        config.AppConfig
	paseto     jwt.Maker
	sentry     *obs.SentryObs
	middleware middleware.Middleware
	mapper     *automapper.AutoMapper
}

func New() *MainApp {
	cfg := config.NewConfigMgr(".").Load()
	ctx := context.Background()

	autoMapper, err := automapper.New().Configure(ApiMappers)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	paseto, err := jwt.NewPasetoMaker(cfg.PASETO)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	sentry, err := obs.NewSentryObs(cfg.OBSERVABILITY)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	tracer, err := obs.NewTracer(ctx, cfg.SERVICE.NAME, cfg.OBSERVABILITY).Initialize()
	if err != nil {
		log.Fatal(err)
		return nil
	}

	cache := storage.GetCacheStore()
	conn, pool, err := pgx.OpenConnectionPool(ctx, cfg.DATABASE)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	store := db.NewStore(tracer, conn, cache)

	sf := services.NewServiceFactory(cfg, conn, store, tracer, cache)
	r := mux.NewRouter()

	app := MainApp{
		router: r,
		pool:   pool,
		store:  store,
		sf:     sf,
		cfg:    cfg,
		paseto: paseto,
		sentry: sentry,
		mapper: autoMapper,
	}

	return &app
}

func (app *MainApp) SetMiddleware() *MainApp {
	app.middleware = middleware.NewMiddleware(app.cfg, app.paseto)
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
