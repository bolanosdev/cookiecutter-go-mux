# About

This project is a [Cookiecutter](https://github.com/cookiecutter/cookiecutter) template for quickly spinning up a
Go microservice. It was designed to be used by the Application Services team to quickly spin up new microservices
to prevent a lot of startup overhead.

## Getting Started

### Prequisites

- Cookiecutter
- Go
- Docker
- Golangci-lint

### Creating a new project

1.) From the root of your project workspaces, run:

```
cookiecutter git@github.com:bolanosdev/cookiecutter-go-mux.git
```

2.) Follow the prompts - if you're just trying it out, just use the defaults. For more info, see Project Options below.

```
$ groupname [company.com]
$ projectname [example_service]:
$ projectdescription [A brief overview of your service.]:
$ go_module [company.com/example-service]:
$ go_version [1.23]:
$ docker_base_image [gcr.io/distroless/base]:
$ namespace [default]:
```

3.) Change directores to the project that was created and run the following:

```
cd ~/path/to/example-service
go mod tidy
go mod vendor
make build
docker-compose up --force-recreate --build --abort-on-container-exit
```

4.) The application and e2e test should exit successfully.

### Project Options

| Option              | Details                                                                                                      |
| ------------------- | ------------------------------------------------------------------------------------------------------------ |
| project_name        | This is the name of your project. If you use multiple words, make it spinal-case. (e.g. example-service)     |
| project_description | This is a description of your project - short and sweet works here.                                          |
| project_slug        | This is the Bitbucket project slug where you plan to host the code - (e.g. as, plat, etc.). Keep lower case. |
| go_module           | This is the go module. This will be auto-generated from your project name and project slug.                  |
| go_version          | This is the version of Go we want to use. Defaults to 1.15.                                                  |
| docker_base_image   | This is the base docker image to use when creating the project (excluding the hostname). Defaults to buster. |
| namespace           | The Kubernetes namespace, for example u-&lt;CLOUDFLARE-USERNAME&gt; for your personal namespace.             |

## Misc

Remember to set `REMOTE_CHECK_MASTER_BRANCH: main` in the `env` field of the teamcity branch & production docker-build stages.
