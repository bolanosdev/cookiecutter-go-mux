package services

import "{{ cookiecutter.group_name }}/{{ cookiecutter.project_name }}/db"

type ServiceFactory struct {
	Accounts    AccountService
	Roles       RoleService
	Permissions PermissionService
}

func NewServiceFactory(store db.Store) ServiceFactory {
	return ServiceFactory{
		Accounts:    NewAccountService(store),
		Roles:       NewRoleService(store),
		Permissions: NewPermissionService(store),
	}
}
