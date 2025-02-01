package cmd

import (
	"context"

	"{{ cookiecutter.group_name }}/{{ cookiecutter.project_name }}/cmd/cfg"
	"{{ cookiecutter.group_name }}/{{ cookiecutter.project_name }}/db"
	"{{ cookiecutter.group_name }}/{{ cookiecutter.project_name }}/internal/api"

	"github.com/gin-gonic/gin"
)

func New(ctx context.Context) error {
	config := cfg.Load(".")
	conn, err := db.OpenConnectionPool(ctx, config.DATABASE)
	if err != nil {
		return err
	}

	store := db.NewStore(conn)
	api := api.NewApiFactory(store)

	r := gin.Default()

	r.GET("/health", api.Health.Get)
	r.GET("/accounts", api.Accounts.GetAll)

	return r.Run("0.0.0.0:9000")
}
