package services

import (
	"context"

	"{{ cookiecutter.group_name }}/{{ cookiecutter.project_name }}/db"
	"{{ cookiecutter.group_name }}/{{ cookiecutter.project_name }}/db/models"
	"{{ cookiecutter.group_name }}/{{ cookiecutter.project_name }}/internal/utils"
)

type PermissionService struct {
	store db.Store
}

func NewPermissionService(store db.Store) PermissionService {
	return PermissionService{
		store: store,
	}
}

func (svc PermissionService) GetAll(c context.Context) ([]models.Permission, error) {
	ctx, span := utils.TracerWithContext(c, "svc.GetAll")

	permissions, err := svc.store.GetPermissions(ctx)
	if err != nil {
		return nil, err
	}

	span.End()
	return permissions, nil
}

func (svc PermissionService) GetByID(c context.Context, id int) (*models.Permission, error) {
	ctx, span := utils.TracerWithContext(c, "svc.GetByID")

	permission, err := svc.store.GetPermissionById(ctx, id)
	if err != nil {
		return nil, err
	}

	span.End()
	return &permission, nil
}
