package migration

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/hossein1376/BehKhan/catalogue/internal/domain/repository"
	"github.com/hossein1376/BehKhan/catalogue/internal/infrastructure/database/maria/pool"
)

func Migrate(db *pool.Pool, path string) error {
	dir, err := os.ReadDir(path)
	if err != nil {
		return fmt.Errorf("read directory: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	return db.Query(ctx, func(r *repository.Repo) error {
		for _, file := range dir {
			if file.IsDir() {
				continue
			}
			m, err := os.ReadFile(fmt.Sprintf("%s/%s", path, file.Name()))
			if err != nil {
				return fmt.Errorf("read file: %w", err)
			}
			if _, err := r.ExecContext(ctx, string(m)); err != nil {
				return fmt.Errorf("execute migration: %w", err)
			}
		}
		return nil
	})
}
