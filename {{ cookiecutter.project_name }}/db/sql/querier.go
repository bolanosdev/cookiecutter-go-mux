package sql

import (
	"context"

	"{{ cookiecutter.group_name }}/{{ cookiecutter.project_name }}/db/models"
)

type Querier interface {
	GetAccounts(ctx context.Context) ([]models.Account, error)
	GetAccountById(ctx context.Context) (models.Account, error)
	GetRoles(ctx context.Context) ([]models.Role, error)
	GetRoleById(ctx context.Context) (models.Role, error)
	GetPermissions(ctx context.Context) ([]models.Permission, error)
	GetPermissionById(ctx context.Context) (models.Permission, error)
}
