package service

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"{{cookiecutter.group_name}}/{{cookiecutter.service_name}}/service/handlers"
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
	app.router.Handle("/health", app.middleware.Tracing(HealthHandler, "GET /health"))
	app.router.Handle("/data", app.middleware.Tracing(handlers.NewDataApi().Get, "GET /data"))
	app.router.Handle("/restricted", app.middleware.Authorize(RestrictedHandler, "GET /restricted"))

	http.Handle("/", app.router)

	return app
}

func (app *MainApp) SetAuthRouter() {
	api := handlers.NewAuthApi(app.sf, app.paseto)

	app.router.Handle("/login", app.middleware.Tracing(api.Login, "POST /login")).Methods("POST")
	app.router.Handle("/signup", app.middleware.Tracing(api.SignUp, "POST /signup")).Methods("POST")
}

func (app *MainApp) SetAccountsRouter() {
	api := handlers.NewAccountApi(app.sf)

	app.router.Handle("/accounts", app.middleware.Tracing(api.GetAll, "GET /accounts"))
	app.router.Handle("/accounts/{id}", app.middleware.Tracing(api.GetByID, "GET /accounts/:id"))
}

func (app *MainApp) SetRolesRouter() {
	api := handlers.NewRoleApi(app.sf)

	app.router.Handle("/roles", app.middleware.Tracing(api.GetAll, "GET /roles"))
	app.router.Handle("/roles/{id}", app.middleware.Tracing(api.GetByID, "GET /roles/:id"))
}

func (app *MainApp) SetPermissionsRouter() {
	api := handlers.NewPermissionApi(app.sf)

	app.router.Handle("/permissions", app.middleware.Tracing(api.GetAll, "GET /permissions"))
	app.router.Handle("/permissions/{id}", app.middleware.Tracing(api.GetByID, "GET /permissions/:id"))
}
