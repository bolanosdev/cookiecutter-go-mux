package cmd

import (
	"fmt"
	"net/http"

	"{{ cookiecutter.group_name }}/{{ cookiecutter.project_name }}/cmd/middleware"
	"{{ cookiecutter.group_name }}/{{ cookiecutter.project_name }}/internal/api"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func RestrictedHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello I am  restricted")
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "OK")
}

func (app *MainApp) SetRouter() *MainApp {
	app.SetAccountsRouter()
	app.SetAuthRouter()
	app.SetRolesRouter()
	app.SetPermissionsRouter()

	app.router.Handle("/metrics", promhttp.Handler())
	app.router.Handle("/health", middleware.Tracing(HealthHandler, "GET /health"))
	app.router.Handle("/data", middleware.Tracing(api.NewDataApi().Get, "GET /data"))
	app.router.Handle("/restricted", middleware.Authorization(RestrictedHandler, "GET /restricted"))

	http.Handle("/", app.router)

	return app
}

func (app *MainApp) SetAuthRouter() {
	api := api.NewAuthApi(app.sf, app.paseto)

	app.router.Handle("/login", middleware.Tracing(api.Login, "POST /login")).Methods("POST")
	app.router.Handle("/signup", middleware.Tracing(api.SignUp, "POST /signup")).Methods("POST")
}

func (app *MainApp) SetAccountsRouter() {
	api := api.NewAccountApi(app.sf)

	app.router.Handle("/accounts", middleware.Tracing(api.GetAll, "GET /accounts"))
	app.router.Handle("/accounts/{id}", middleware.Tracing(api.GetByID, "GET /accounts/:id"))
}

func (app *MainApp) SetRolesRouter() {
	api := api.NewRoleApi(app.sf)

	app.router.Handle("/roles", middleware.Tracing(api.GetAll, "GET /roles"))
	app.router.Handle("/roles/{id}", middleware.Tracing(api.GetByID, "GET /roles/:id"))
}

func (app *MainApp) SetPermissionsRouter() {
	api := api.NewPermissionApi(app.sf)

	app.router.Handle("/permissions", middleware.Tracing(api.GetAll, "GET /permissions"))
	app.router.Handle("/permissions/{id}", middleware.Tracing(api.GetByID, "GET /permissions/:id"))
}
