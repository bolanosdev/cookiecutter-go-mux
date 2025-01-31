package main

import (
	"flag"
	"log"
	"os"

	"{{ cookiecutter.group_name }}/{{ cookiecutter.project_name }}/cmd"
)

func main() {
	flag.Parse()
	app := cmd.MainApp{}
	err := app.New().Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
