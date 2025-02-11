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

func (svc AccountService) GetAll(c context.Context) ([]models.Account, error) {
	ctx, span := utils.TracerWithContext(c, "svc.GetAll")
	accounts, err := svc.store.GetAccounts(ctx)
	if err != nil {
		return nil, err
	}

	span.End()
	return accounts, nil
}

func (svc AccountService) GetByEmail(c context.Context, email string) (*models.Account, error) {
	ctx, span := utils.TracerWithContext(c, "svc.GetByEmail")
	account, err := svc.store.GetAccountByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	span.End()
	return &account, nil
}

func (svc AccountService) GetByID(c context.Context, id int) (*models.Account, error) {
	ctx, span := utils.TracerWithContext(c, "svc.GetByID")
	account, err := svc.store.GetAccountById(ctx, id)
	if err != nil {
		return nil, err
	}

	span.End()
	return &account, nil
}

func (svc AccountService) Login(c context.Context, email string, password string) (*models.Account, error) {
	ctx, span := utils.TracerWithContext(c, "svc.Login")
	account, err := svc.store.GetAccountByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	_, hash_span := utils.TracerWithContext(c, "svc.login.check_password")
	err = utils.CheckPassword(password, account.Password)
	if err != nil {
		return nil, err
	}
	hash_span.End()

	span.End()
	return &account, nil
}

func (svc AccountService) Create(c context.Context, email string, password string) (*models.Account, error) {
	ctx, span := utils.TracerWithContext(c, "svc.Create")
	account, err := svc.store.CreateAccount(ctx, email, password)
	if err != nil {
		return nil, err
	}

	span.End()
	return account, nil
}
