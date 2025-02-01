package main

import (
	"context"
	"log"

	"{{ cookiecutter.group_name }}/{{ cookiecutter.project_name }}/cmd"
)

func main() {
	ctx := context.Background()
	err := cmd.New(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
