package services

import (
	"context"

	"{{ cookiecutter.group_name }}/{{ cookiecutter.service_name }}/internal/cache"
	"{{ cookiecutter.group_name }}/{{ cookiecutter.service_name }}/internal/db"
	"{{ cookiecutter.group_name }}/{{ cookiecutter.service_name }}/internal/db/models"
	"{{ cookiecutter.group_name }}/{{ cookiecutter.service_name }}/internal/db/sql"
	"{{ cookiecutter.group_name }}/{{ cookiecutter.service_name }}/internal/utils/obs"
	pw "{{ cookiecutter.group_name }}/{{ cookiecutter.service_name }}/internal/utils/password"

	"github.com/bolanosdev/go-snacks/observability/logging"
	qb "github.com/bolanosdev/query-builder"
)

type AccountService struct {
	db     sql.PgxPoolConn
	store  db.Store
	cache  *cache.InMemoryCacheStore
	tracer obs.TracerInterface
}

func NewAccountService(
	db sql.PgxPoolConn,
	store db.Store,
	tracer obs.TracerInterface,
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
	ctx, span := svc.tracer.Trace(c, "svc.account.GetAll")
	logger := c.Value("logger").(*logging.ContextLogger)
	defer span.End()

	accounts, err := svc.store.GetAccounts(ctx)
	if err != nil {
		logger.Error().Err(err).Msg("svc.GetAll")
		return nil, err
	}

	return accounts, nil
}

func (svc AccountService) GetByEmail(c context.Context, email string) (*models.Account, error) {
	ctx, span := svc.tracer.Trace(c, "svc.account.GetByEmail")
	logger := c.Value("logger").(*logging.ContextLogger)
	defer span.End()

	account, err := svc.store.GetAccount(ctx, qb.ByStringColumn("a.email", []string{email}))
	if err != nil {
		logger.Error().Err(err).Msg("svc.GetByEmail")
		return nil, err
	}

	return &account, nil
}

func (svc AccountService) GetByID(c context.Context, id int) (*models.Account, error) {
	ctx, span := svc.tracer.Trace(c, "svc.account.GetByID")
	logger := c.Value("logger").(*logging.ContextLogger)
	defer span.End()

	account, err := svc.store.GetAccount(ctx, qb.ByIntColumn("a.id", []int{id}))
	if err != nil {
		logger.Error().Err(err).Msg("svc.GetByID")
		return nil, err
	}

	return &account, nil
}

func (svc AccountService) Login(c context.Context, email string, password string) (*models.Account, error) {
	ctx, span := svc.tracer.Trace(c, "svc.account.Login")
	logger := c.Value("logger").(*logging.ContextLogger)
	defer span.End()

	account, err := svc.store.GetAccount(ctx, qb.ByStringColumn("a.email", []string{email}))
	if err != nil {
		logger.Error().Err(err).Msg("svc.Login")
		return nil, err
	}

	_, hash_span := svc.tracer.Trace(c, "svc.login.check_password")
	err = pw.CheckPassword(password, account.Password)
	if err != nil {
		return nil, err
	}
	hash_span.End()

	return &account, nil
}

func (svc AccountService) SignUp(c context.Context, email string, password string) (*models.Account, error) {
	ctx, span := svc.tracer.Trace(c, "svc.SignUp")
	logger := c.Value("logger").(*logging.ContextLogger)
	defer span.End()

	hash_password, err := pw.HashPassword(password)
	if err != nil {
		logger.Error().Err(err).Msg("svc.SignUp")
		return nil, err
	}

	account, err := svc.store.CreateAccount(ctx, email, hash_password)
	if err != nil {
		logger.Error().Err(err).Msg("svc.SignUp")
		return nil, err
	}

	return account, nil
}
