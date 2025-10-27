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

	dataApi := handlers.NewDataApi()

	app.router.Handle("/metrics", promhttp.Handler())
	app.router.HandleFunc("/health", HealthHandler)
	app.router.HandleFunc("/data", dataApi.Get)
	app.router.Handle("/restricted", app.middleware.Authorize(RestrictedHandler, "GET /restricted")).Methods("GET")
	http.Handle("/", app.router)

	return app
}

func (app *MainApp) SetAuthRouter() {
	api := handlers.NewAuthApi(app.sf, app.paseto)

	app.router.HandleFunc("/login", api.Login).Methods("POST")
	app.router.HandleFunc("/signup", api.SignUp).Methods("POST")
}

func (app *MainApp) SetAccountsRouter() {
	api := handlers.NewAccountApi(app.sf)

	app.router.HandleFunc("/accounts", api.GetAll).Methods("GET")
	app.router.HandleFunc("/accounts/{id}", api.GetByID).Methods("GET")
}

func (app *MainApp) SetRolesRouter() {
	api := handlers.NewRoleApi(app.sf)

	app.router.HandleFunc("/roles", api.GetAll).Methods("GET")
	app.router.HandleFunc("/roles/{id}", api.GetByID).Methods("GET")
}

func (app *MainApp) SetPermissionsRouter() {
	api := handlers.NewPermissionApi(app.sf)

	app.router.HandleFunc("/permissions", api.GetAll).Methods("GET")
	app.router.HandleFunc("/permissions/{id}", api.GetByID).Methods("GET")
}
