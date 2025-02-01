package api

import "{{ cookiecutter.group_name }}/{{ cookiecutter.project_name }}/db"

type ApiFactory struct {
	Accounts AccountApi
	Health   HealthApi
}

func NewApiFactory(store db.Store) ApiFactory {
	return ApiFactory{
		Accounts: NewAccountApi(store),
		Health:   HealthApi{},
	}
}
