package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
	"github.com/mirhijinam/backend-bootcamp-assignment-2024/internal/config"
)

func MustOpenDB(ctx context.Context, dbCfg config.DBConfig) (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig("")
	if err != nil {
		return nil, fmt.Errorf("MustOpenDB config parse: %w", err)
	}

	config.ConnConfig.Host = dbCfg.PgHost
	config.ConnConfig.Port = dbCfg.PgPort
	config.ConnConfig.Database = dbCfg.PgDatabase
	config.ConnConfig.User = dbCfg.PgUser
	config.ConnConfig.Password = dbCfg.PgPassword

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("failed to open db: %w", err)
	}

	if err = pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("failed to ping db: %w", err)
	}

	return pool, nil
}
