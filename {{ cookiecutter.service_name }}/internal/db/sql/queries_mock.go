package sql

import (
	"context"

	"github.com/bolanosdev/go-snacks/observability/jaeger"
	"github.com/bolanosdev/go-snacks/storage"
	"github.com/jackc/pgx/v5"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/pashagolub/pgxmock/v4"
)

type MockPgxPoolConn struct {
	Mock pgxmock.PgxConnIface
}

func (m *MockPgxPoolConn) QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row {
	return m.Mock.QueryRow(ctx, sql, args...)
}

func (m *MockPgxPoolConn) Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error) {
	return m.Mock.Query(ctx, sql, args...)
}

func (m *MockPgxPoolConn) Begin(ctx context.Context) (pgx.Tx, error) {
	return m.Mock.Begin(ctx)
}

func (m *MockPgxPoolConn) Exec(ctx context.Context, sql string, args ...interface{}) (pgconn.CommandTag, error) {
	return m.Mock.Exec(ctx, sql, args...)
}

func (m *MockPgxPoolConn) Release() {
	// No-op, since Release() in pgxpool.Conn is just returning the connection to the pool
}

type (
	MockQueryReturnType int
)

const (
	QRT_NA MockQueryReturnType = iota
	QRT_TX
	QRT_Error
	QRT_Rows
)

type PGXMocker struct {
	Ctx     context.Context
	Conn    pgxmock.PgxConnIface
	Mock    *MockPgxPoolConn
	Querier *Queries
}

func GetPGXMocks() (pgxmock.PgxConnIface, *MockPgxPoolConn, *Queries) {
	conn, _ := pgxmock.NewConn()
	mock := &MockPgxPoolConn{conn}
	store := storage.NewCacheStore()
	tracer := jaeger.NewMockTracer()
	querier := NewQueries(tracer, mock, store)

	return conn, mock, querier
}

func NewPGXMocker(ctx context.Context) PGXMocker {
	conn, mock, querier := GetPGXMocks()
	return PGXMocker{
		Ctx:     ctx,
		Conn:    conn,
		Mock:    mock,
		Querier: querier,
	}
}

func (m *PGXMocker) MockQuerySuccess(query string, args_count int, rows *pgxmock.Rows) {
	args := GetArgs(args_count)
	m.Conn.ExpectQuery(query).
		WithArgs(args...).
		WillReturnRows(rows)
}

func (m *PGXMocker) MockQueryFailure(query string, args_count int, return_type MockQueryReturnType, err error) {
	args := GetArgs(args_count)

	if return_type == QRT_Error {
		m.Conn.ExpectQuery(query).WithArgs(args...).WillReturnError(err)
	} else {
		m.Conn.ExpectQuery(query).
			WithArgs(args...).
			WillReturnRows(
				pgxmock.NewRows([]string{"id"}).AddRow(1),
			)
	}
}

func (m *PGXMocker) MockQueryPGFailure(query string, args_count int, return_type MockQueryReturnType, err pgconn.PgError) {
	args := GetArgs(args_count)

	if return_type == QRT_Error {
		m.Conn.ExpectQuery(query).WithArgs(args...).WillReturnError(&err)
	} else {
		m.Conn.ExpectQuery(query).
			WithArgs(args...).
			WillReturnRows(
				pgxmock.NewRows([]string{"id"}).AddRow(1),
			)
	}
}

func (m *PGXMocker) MockExecSuccess(query string, args_count int) {
	args := GetArgs(args_count)

	m.Conn.ExpectExec(query).
		WithArgs(args...).
		WillReturnResult(pgxmock.NewResult("inserted", 1))
}

func (m *PGXMocker) MockExecFailure(query string, args_count int, return_type MockQueryReturnType, err error) {
	args := GetArgs(args_count)
	if return_type == QRT_Error {
		m.Conn.ExpectExec(query).WithArgs(args...).WillReturnError(err)
	}
}

func GetArgs(arg_count int) []any {
	args := []any{}

	for arg_count > 0 {
		args = append(args, pgxmock.AnyArg())
		arg_count--
	}

	return args
}
