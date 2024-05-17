package pool

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/hossein1376/BehKhan/catalogue/internal/domain/repository"
	"github.com/hossein1376/BehKhan/catalogue/internal/infrastructure/database/maria"
)

// Pool embeds an instance of *repository.Tables, offering available repository methods. Also, due to
// embedding an instance of Querier, it's possible to directly run raw sql statements against the database.
type Pool struct {
	db *sql.DB
}

// New opens a database connection to a MariaDB (or MySQL) database with the given dsn address.
func New(dsn string) (*Pool, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("sql.Open(%s): %w", dsn, err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("pinging database: %w", err)
	}

	return &Pool{db: db}, nil
}

// Query starts a transaction and executes the provided function. It creates an instance of *Repo with the concrete
// types and passes it the function. In case of an error, it will attempt to rollback.
func (p *Pool) Query(ctx context.Context, f repository.QueryFunc) error {
	tx, err := p.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("starting transaction: %w", err)
	}

	g, err := gorm.Open(mysql.New(mysql.Config{
		Conn: tx,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return fmt.Errorf("opening gorm connection: %w", err)
	}

	r := &repository.Repo{
		Querier: tx,
		Tables:  maria.New(g),
	}
	err = f(r)
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			return fmt.Errorf("rollback: %w query: %w", rollbackErr, err)
		}
		return err
	}
	return nil
}

// Close closes the database connection.
func (p *Pool) Close() error {
	return p.db.Close()
}
