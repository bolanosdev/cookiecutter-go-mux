package cmd

import (
	"fmt"
	"net/http"

	"{{ cookiecutter.group_name }}/{{ cookiecutter.project_name }}/cmd/cfg"

	"github.com/gin-gonic/gin"
	"github.com/urfave/cli/v2"
)

type MainApp struct {
	config cfg.AppConfig
}

func (main_app *MainApp) New() *cli.App {
	app := cli.NewApp()
	app.Name = "{{ cookiecutter.project_name }}"
	app.Usage = "{{ cookiecutter.project_description }}"
	main_app.config = cfg.Load(".")
	app.Action = main_app.Run()

	return app
}

func (main_app *MainApp) Run() func(c *cli.Context) error {
	return func(c *cli.Context) error {
		fmt.Println(main_app.config.SERVICE_NAME)

		r := gin.Default()
		r.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
		return r.Run("0.0.0.0:9000")
	}
}
