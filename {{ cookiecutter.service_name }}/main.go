package main

import (
	"{{ cookiecutter.group_name }}/{{ cookiecutter.service_name }}/service"
)

func main() {
	service.New().
		SetMiddleware().
		SetRouter().
		Start()
}
