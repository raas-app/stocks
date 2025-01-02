package database

import (
	"context"
	"database/sql"
)

// Queryer is an interface used for selection queries.
type Queryer interface {
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
}

// Execer is an interface used for executing queries.
type Execer interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error)
}

type DBPinger interface {
	PingContext(ctx context.Context) error
}

// Common is a union interface which can query, and exec, with Context
// It can be pure db handle or tx handle.
type Common interface {
	Queryer
	Execer
}

type HookMysql interface {
	beforeOperation(ctx context.Context) context.Context
	afterOperation(ctx context.Context, query string)
}
