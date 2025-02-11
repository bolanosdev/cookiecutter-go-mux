package services

import (
	"context"

	"{{ cookiecutter.group_name }}/{{ cookiecutter.project_name }}/db"
	"{{ cookiecutter.group_name }}/{{ cookiecutter.project_name }}/db/models"
	"{{ cookiecutter.group_name }}/{{ cookiecutter.project_name }}/internal/utils"
)

type RoleService struct {
	store db.Store
}

func NewRoleService(store db.Store) RoleService {
	return RoleService{
		store: store,
	}
}

func (svc RoleService) GetAll(c context.Context) ([]models.Role, error) {
	ctx, span := utils.TracerWithContext(c, "svc.GetAll")

	roles, err := svc.store.GetRoles(ctx)
	if err != nil {
		return nil, err
	}

	span.End()
	return roles, nil
}

func (svc RoleService) GetByID(c context.Context, id int) (*models.Role, error) {
	ctx, span := utils.TracerWithContext(c, "svc.GetByID")

	role, err := svc.store.GetRoleById(ctx, id)
	if err != nil {
		return nil, err
	}

	span.End()
	return &role, nil
}
