package services

import (
	"context"

	"{{ cookiecutter.group_name }}/{{ cookiecutter.service_name }}/internal/db"
	"{{ cookiecutter.group_name }}/{{ cookiecutter.service_name }}/internal/db/models"
	"{{ cookiecutter.group_name }}/{{ cookiecutter.service_name }}/internal/db/sql"
	pw "{{ cookiecutter.group_name }}/{{ cookiecutter.service_name }}/internal/utils/password"

	"github.com/bolanosdev/go-snacks/observability/jaeger"
	"github.com/bolanosdev/go-snacks/storage"
	qb "github.com/bolanosdev/query-builder"
	"github.com/pkg/errors"
)

type AccountService struct {
	db     sql.PgxPoolConn
	store  db.Store
	cache  *storage.InMemoryCacheStore
	tracer jaeger.JaegerInterface
}

func NewAccountService(
	db sql.PgxPoolConn,
	store db.Store,
	tracer jaeger.JaegerInterface,
	cache *storage.InMemoryCacheStore,
) AccountService {
	return AccountService{
		db:     db,
		store:  store,
		tracer: tracer,
		cache:  cache,
	}
}

func (svc AccountService) GetAll(c context.Context) ([]*models.Account, error) {
	ctx := svc.tracer.TraceFunc(c)

	accounts, err := svc.store.GetAccounts(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "fail to retrieve accounts")
	}
	return accounts, nil
}

func (svc AccountService) GetByEmail(c context.Context, email string) (*models.Account, error) {
	ctx := svc.tracer.TraceFunc(c)

	account, err := svc.store.GetAccount(ctx, qb.ByStringColumn("a.email", []string{email}))
	if err != nil {
		return nil, errors.Wrap(err, "fail to retrieve account by email")
	}
	return account, nil
}

func (svc AccountService) GetByID(c context.Context, id int) (*models.Account, error) {
	ctx := svc.tracer.TraceFunc(c)
	account, err := svc.store.GetAccount(ctx, qb.ByIntColumn("a.id", []int{id}))
	if err != nil {
		return nil, errors.Wrap(err, "fail to retrieve account by id")
	}
	return account, nil
}

func (svc AccountService) Login(c context.Context, email string, password string) (*models.Account, error) {
	ctx := svc.tracer.TraceFunc(c)
	account, err := svc.store.GetAccount(ctx, qb.ByStringColumn("a.email", []string{email}))
	if err != nil {
		return nil, errors.Wrap(err, "fail to retrieve account by email")
	}

	err = pw.CheckPassword(password, account.Password)
	if err != nil {
		return nil, errors.Wrap(err, "wrong credentials")
	}

	return account, nil
}

func (svc AccountService) SignUp(c context.Context, email string, password string) (*models.Account, error) {
	ctx := svc.tracer.TraceFunc(c)
	hash_password, err := pw.HashPassword(password)
	if err != nil {
		return nil, errors.Wrap(err, "fail to hash password")
	}

	account, err := svc.store.CreateAccount(ctx, email, hash_password)
	if err != nil {
		return nil, errors.Wrap(err, "fail to create account")
	}

	return account, nil
}
