package cmd

import (
	"context"
	"log"
	"net/http"

	"{{ cookiecutter.group_name }}/{{ cookiecutter.project_name }}/cmd/authorization"
	"{{ cookiecutter.group_name }}/{{ cookiecutter.project_name }}/cmd/cfg"
	"{{ cookiecutter.group_name }}/{{ cookiecutter.project_name }}/cmd/middleware"
	"{{ cookiecutter.group_name }}/{{ cookiecutter.project_name }}/cmd/telemetry"
	"{{ cookiecutter.group_name }}/{{ cookiecutter.project_name }}/db"
	"{{ cookiecutter.group_name }}/{{ cookiecutter.project_name }}/internal/services"

	"github.com/gorilla/mux"
	muxMonitor "github.com/labbsr0x/mux-monitor"
)

type MainApp struct {
	ctx     context.Context
	router  *mux.Router
	store   db.Store
	sf      services.ServiceFactory
	cfg     cfg.AppConfig
	monitor *muxMonitor.Monitor
	paseto  authorization.Maker
}

func New() *MainApp {
	cfg := cfg.Load(".")
	ctx := context.Background()
	paseto, err := authorization.NewPasetoMaker(cfg.PASETO.TOKEN_SYMETRIC_KEY)
	if err != nil {
		log.Fatal(err)
		return nil
	}

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
	sf := services.NewServiceFactory(store)
	r := mux.NewRouter()

	app := MainApp{
		ctx:     context.Background(),
		router:  r,
		store:   store,
		sf:      sf,
		cfg:     cfg,
		monitor: monitor,
		paseto:  paseto,
	}

	return &app
}

func (app *MainApp) SetMiddleware() *MainApp {
	app.router.Use(middleware.Logging)
	app.router.Use(app.monitor.Prometheus)

	return app
}

func (app *MainApp) Start() {
	log.Fatal(http.ListenAndServe(":9000", nil))
}
