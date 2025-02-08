package cmd

import (
	"context"
	"log"
	"net/http"

	"{{ cookiecutter.group_name }}/{{ cookiecutter.project_name }}/cmd/cfg"
	"{{ cookiecutter.group_name }}/{{ cookiecutter.project_name }}/cmd/telemetry"
	"{{ cookiecutter.group_name }}/{{ cookiecutter.project_name }}/cmd/middleware"

	"{{ cookiecutter.group_name }}/{{ cookiecutter.project_name }}/db"
	"{{ cookiecutter.group_name }}/{{ cookiecutter.project_name }}/internal/api"
	"{{ cookiecutter.group_name }}/{{ cookiecutter.project_name }}/internal/utils"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5"
  muxMonitor "github.com/labbsr0x/mux-monitor"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type MainApp struct {
	ctx   context.Context
	conn  *pgx.Conn
	r     *mux.Router
	store db.Store
	cfg   cfg.AppConfig
	monitor *muxMonitor.Monitor
}

func New() *MainApp {
	cfg := cfg.Load(".")
	ctx := context.Background()
	conn, err := db.OpenConnectionPool(ctx, cfg.DATABASE)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	err = telemetry.Initialize(ctx, cfg)
	if err != nil {
		log.Fatal(err)
		return nil
	}

  monitor, err := muxMonitor.New("v1.0.0", muxMonitor.DefaultErrorMessageKey, muxMonitor.DefaultBuckets)
	if err != nil {
		panic(err)
	}

	store := db.NewStore(conn)
	r := mux.NewRouter()

	app := MainApp{
		ctx:   context.Background(),
		r:     r,
		store: store,
		cfg:   cfg,
		monitor: monitor,
	}

	return &app
}

func (app *MainApp) SetRouter() *MainApp {
	api := api.NewApiFactory(app.store)

	app.r.Handle("/metrics", promhttp.Handler())
	app.r.Handle("/data", utils.Instrument(api.Data.Get, "GET /data"))
	app.r.Handle("/accounts", utils.Instrument(api.Accounts.GetAll, "GET /accounts"))
	app.r.Handle("/roles", utils.Instrument(api.Roles.GetAll, "GET /roles"))
	app.r.Handle("/permissions", utils.Instrument(api.Permissions.GetAll, "GET /permissions"))

	http.Handle("/", app.r)
	return app
}

func (app *MainApp) SetMiddleware() *MainApp {
	app.r.Use(middleware.Logging)
	app.r.Use(app.monitor.Prometheus)
	return app
}

func (app *MainApp) Start() {
	log.Fatal(http.ListenAndServe("localhost:9000", nil))
}
