package cmd

import (
	"context"
	"log"

	"{{ cookiecutter.group_name }}/{{ cookiecutter.project_name }}/cmd/cfg"
	"{{ cookiecutter.group_name }}/{{ cookiecutter.project_name }}/cmd/middleware"
	"{{ cookiecutter.group_name }}/{{ cookiecutter.project_name }}/cmd/telemetry"
	"{{ cookiecutter.group_name }}/{{ cookiecutter.project_name }}/db"
	"{{ cookiecutter.group_name }}/{{ cookiecutter.project_name }}/internal/api"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

type MainApp struct {
	r    *gin.Engine
	conn *pgx.Conn
	db   db.Store
	cfg  cfg.AppConfig
}

func New(ctx context.Context) *MainApp {
	cfg := cfg.Load(".")
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
	r := gin.Default()
	db := db.NewStore(conn)

	app := MainApp{
		r:   r,
		db:  db,
		cfg: cfg,
	}
	return &app
}

func (app *MainApp) SetRouter() *MainApp {
	api := api.NewApiFactory(app.db)
	app.r.GET("/health", api.Health.Get)
	app.r.GET("/accounts", api.Accounts.GetAll)
	app.r.GET("/roles", api.Roles.GetAll)
	app.r.GET("/permissions", api.Permissions.GetAll)

	return app
}

func (app *MainApp) SetMiddleware() *MainApp {
	middleware.CORS(app.r)
	middleware.Metrics(app.r)

	app.r.Use(middleware.Logger(app.cfg))
	return app
}

func (app *MainApp) Start() error {
	return app.r.Run("0.0.0.0:9000")
}
