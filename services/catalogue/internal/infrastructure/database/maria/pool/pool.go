package pool

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/hossein1376/BehKhan/catalogue/internal/infrastructure/database/maria"
)

// Pool embeds an instance of *maria.Tables, offering available repository methods. Also, due to
// embedding *sql.Tx, it's possible to directly run raw sql statements against the database.
type Pool struct {
	*maria.Tables
	*sql.Tx
}

// Query is an alias for closure function type to run against DB instances to query database using
// repository methods or direct sql.
type Query = func(ctx context.Context, r *Pool) error

// DB contains an instance of *sql.DB. It is safe for concurrent use by multiple goroutines and
// maintains its own pool of idle connections.
type DB struct {
	db *sql.DB
}

// New opens a database connection to a MariaDB (or MySQL) database with the given dsn address.
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

// Query starts a transaction and executes the given function. It creates and passes an instance of
// Pool to the function. In case of an error, it will attempt to roll back.
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
			return fmt.Errorf("rollback: %w query: %w", rollbackErr, err)
		}
		return fmt.Errorf("query: %w", err)
	}
	return nil
}

// DB returns the underlying *sql.DB struct.
func (p *DB) DB() *sql.DB {
	return p.db
}
