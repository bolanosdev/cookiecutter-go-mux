package pgx

import (
	"context"
	"fmt"

	"{{ cookiecutter.group_name }}/{{ cookiecutter.service_name }}/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateConnectionString(db_cfg config.DBConfig) string {
	conn_string := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", db_cfg.USERNAME, db_cfg.PASSWORD, db_cfg.HOSTNAME, db_cfg.PORT, db_cfg.DATABASE, db_cfg.SSL)
	return conn_string
}

func CreatePGXConfig(db_cfg config.DBConfig) (*pgxpool.Config, error) {
	conn_string := CreateConnectionString(db_cfg)
	cfg, err := pgxpool.ParseConfig(conn_string)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func OpenConnectionPool(ctx context.Context, cfg *pgxpool.Config) (*pgxpool.Conn, *pgxpool.Pool, error) {
	pool, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		// log.Fatalf("Error while creating pool to the database!! %v", err)
		return nil, nil, err
	}

	conn, err := pool.Acquire(context.Background())
	if err != nil {
		// log.Fatalf("Error while acquiring connection from the database pool!! %v", err)
		return nil, nil, err
	}

	return conn, pool, nil
}
