package services

import (
	"context"

	"{{ cookiecutter.group_name }}/{{ cookiecutter.service_name }}/internal/cache"
	"{{ cookiecutter.group_name }}/{{ cookiecutter.service_name }}/internal/db"
	"{{ cookiecutter.group_name }}/{{ cookiecutter.service_name }}/internal/db/models"
	"{{ cookiecutter.group_name }}/{{ cookiecutter.service_name }}/internal/db/sql"
	"{{ cookiecutter.group_name }}/{{ cookiecutter.service_name }}/internal/utils"

	qb "github.com/bolanosdev/query-builder"
)

type AccountService struct {
	db     sql.PgxPoolConn
	store  db.Store
	cache  *cache.InMemoryCacheStore
	tracer utils.TracerInterface
}

func NewAccountService(
	db sql.PgxPoolConn,
	store db.Store,
	tracer utils.TracerInterface,
	cache *cache.InMemoryCacheStore,
) AccountService {
	return AccountService{
		db:     db,
		store:  store,
		tracer: tracer,
		cache:  cache,
	}
}

func (svc AccountService) GetAll(c context.Context) ([]models.Account, error) {
	ctx, span := svc.tracer.Trace(c, "svc.Login")
	defer span.End()

	accounts, err := svc.store.GetAccounts(ctx)
	if err != nil {
		return nil, err
	}

	return accounts, nil
}

func (svc AccountService) GetByEmail(c context.Context, email string) (*models.Account, error) {
	ctx, span := svc.tracer.Trace(c, "svc.GetByEmail")
	defer span.End()
	account, err := svc.store.GetAccount(ctx, qb.ByStringColumn("a.email", email))
	if err != nil {
		return nil, err
	}

	return &account, nil
}

func (svc AccountService) GetByID(c context.Context, id int) (*models.Account, error) {
	ctx, span := svc.tracer.Trace(c, "svc.GetByID")
	defer span.End()

	account, err := svc.store.GetAccount(ctx, qb.ByIntColumn("a.id", id))
	if err != nil {
		return nil, err
	}

	return &account, nil
}

func (svc AccountService) Login(c context.Context, email string, password string) (*models.Account, error) {
	ctx, span := svc.tracer.Trace(c, "svc.Login")
	defer span.End()

	account, err := svc.store.GetAccount(ctx, qb.ByStringColumn("a.email", email))
	if err != nil {
		return nil, err
	}

	_, hash_span := svc.tracer.Trace(c, "svc.login.check_password")
	err = utils.CheckPassword(password, account.Password)
	if err != nil {
		return nil, err
	}
	hash_span.End()

	return &account, nil
}

func (svc AccountService) SignUp(c context.Context, email string, password string) (*models.Account, error) {
	ctx, span := svc.tracer.Trace(c, "svc.SignUp")
	defer span.End()

	hash_password, err := utils.HashPassword(password)
	if err != nil {
		return nil, err
	}

	account, err := svc.store.CreateAccount(ctx, email, hash_password)
	if err != nil {
		return nil, err
	}

	return account, nil
}
