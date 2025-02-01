package db

import (
	"context"
	"fmt"
	"log"

	"{{ cookiecutter.group_name }}/{{ cookiecutter.project_name }}/cmd/cfg"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreatePGXConfig(db_cfg cfg.DBConfig) (*pgxpool.Config, error) {
	conn_string := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", db_cfg.USERNAME, db_cfg.PASSWORD, db_cfg.HOSTNAME, db_cfg.PORT, db_cfg.DATABASE, db_cfg.SSL)

	cfg, err := pgxpool.ParseConfig(conn_string)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func OpenConnectionPool(ctx context.Context, db_config cfg.DBConfig) (*pgxpool.Conn, error) {
	cfg, err := CreatePGXConfig(db_config)
	if err != nil {
		log.Fatalf("Error while creating db config %v", err)
		return nil, err
	}

	pool, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		log.Fatalf("Error while creating pool to the database!! %v", err)
		return nil, err
	}

	conn, err := pool.Acquire(context.Background())
	if err != nil {
		log.Fatalf("Error while acquiring connection from the database pool!! %v", err)
		return nil, err
	}

	return conn, nil
}
