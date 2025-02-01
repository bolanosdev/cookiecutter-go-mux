# About

This project was created using a [Cookiecutter](https://github.com/cookiecutter/cookiecutter) [template](https://google.com) .

## Pending

- GH Actions
- Telemetry
- Metrics
- Logging

## Getting Started

### Prequisites

- [Cookiecutter](https://github.com/cookiecutter/cookiecutter)

```
pipx install cookiecutter
```

- [Go](https://go.dev/doc/install)
- Docker
- Golangci-lint

```
brew install golangci-lint
```

- Go Migrate

```
brew install migrate
```

### Running the project

1.) Install Go dependenciies
Change directores to the project that was created and run the following do download go dependencies:

```
cd ~/path/to/example-service
go mod tidy
go mod vendor
```

2.) Update your database info
if you wish to create a local postgres database, see Makefile to update credentials,
else update the credentials in the app.yaml

```
make postgres
```

Note: the ssl property is used at the end of the connectin string, "disable" for most versions of postgres in development mode, "require" for most of the postgres productions databases.

3.) Apply your first database migration.

The project is already comes with 2 migrations, you can add all your table, keys, index, relationships tear up/tear down scripts on the db/migrations/000001_initial_schema.

3.1) if you wish to create more migrations in the future you can use the go migrate tool.
`migrate create -ext sql -dir db/migrations -seq initial_schema`.
For more information see [golang-migrate github](https://github.com/golang-migrate/migrate).

4.) Run the App
`make run` or `docker-compose up`
