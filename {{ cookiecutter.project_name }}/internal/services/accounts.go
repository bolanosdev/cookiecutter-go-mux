package services

import (
	"context"

	"{{ cookiecutter.group_name }}/{{ cookiecutter.project_name }}/db"
	"{{ cookiecutter.group_name }}/{{ cookiecutter.project_name }}/db/models"
	"{{ cookiecutter.group_name }}/{{ cookiecutter.project_name }}/internal/utils"
)

type AccountService struct {
	store db.Store
}

func NewAccountService(store db.Store) AccountService {
	return AccountService{
		store: store,
	}
}

func (svc AccountService) GetAll(c context.Context) ([]models.Account, models.Account, error) {
	ctx, span := utils.TracerWithContext(c, "svc.GetAll")
	accounts, err := svc.store.GetAccounts(ctx)
	if err != nil {
		return nil, models.Account{}, err
	}

	account, err := svc.store.GetAccountById(ctx)
	if err != nil {
		return nil, models.Account{}, err
	}

	span.End()
	return accounts, account, nil
}
