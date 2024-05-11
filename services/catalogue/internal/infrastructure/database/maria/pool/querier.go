package pool

import (
	"context"
	"database/sql"
)

type Querier struct {
	tx *sql.Tx
}

func (q Querier) ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error) {
	return q.tx.ExecContext(ctx, query, args...)
}

func (q Querier) QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error) {
	return q.tx.QueryContext(ctx, query, args...)
}
