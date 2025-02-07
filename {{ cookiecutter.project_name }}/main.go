package main

import (
	"{{ cookiecutter.group_name }}/{{ cookiecutter.project_name }}/cmd"
)

func main() {
	cmd.New().
		SetMiddleware().
		SetRouter().
		Start()
}
