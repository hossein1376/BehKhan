package pool

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/hossein1376/BehKhan/catalogue/internal/infrastructure/database/maria"
)

type Pool struct {
	*maria.Tables
	*sql.Tx
}

type Query = func(ctx context.Context, r *Pool) error

type DB struct {
	db *sql.DB
}

func New(dsn string) (*DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("sql.Open(%s): %w", dsn, err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("pinging database: %w", err)
	}

	return &DB{db: db}, nil
}

func (p *DB) Query(ctx context.Context, f Query) error {
	tx, err := p.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("starting transaction: %w", err)
	}

	r := &Pool{
		Tx:     tx,
		Tables: maria.New(tx),
	}
	err = f(ctx, r)
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			return fmt.Errorf("rollback error: %w %w", rollbackErr, err)
		}
		return err
	}
	return nil
}

func (p *DB) GetDB() *sql.DB {
	return p.db
}
