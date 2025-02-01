package api

import (
	"{{ cookiecutter.group_name }}/{{ cookiecutter.project_name }}/db"
	"{{ cookiecutter.group_name }}/{{ cookiecutter.project_name }}/internal/services"
)

type ApiFactory struct {
	Accounts    AccountApi
	Roles       RoleApi
	Permissions PermissionApi
	Health      HealthApi
}

func NewApiFactory(store db.Store) ApiFactory {
	sf := services.NewServiceFactory(store)
	return ApiFactory{
		Accounts:    NewAccountApi(sf),
		Roles:       NewRoleApi(sf),
		Permissions: NewPermissionApi(sf),
		Health:      HealthApi{},
	}
}
