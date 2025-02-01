package sql

import (
	_ "{{ cookiecutter.group_name }}/{{ cookiecutter.project_name }}/db/models"
)

type Querier interface{}
