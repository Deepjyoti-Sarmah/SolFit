package storage

import (
	"context"
	"fmt"
	"time"

	"github.com/Deepjyoti-Sarmah/sol-kit-backend/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Store struct {
	db *pgxpool.Pool
}

func NewStore(pool *pgxpool.Pool) *Store {
	return &Store{
		db: pool,
	}
}

func NewPostgresDb(conf *config.Config) (*pgxpool.Pool, error) {
	connString := conf.DatabaseUrl()

	poolConfig, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, fmt.Errorf("failed to parse pgx config: %w", err)
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create connection pool: %w", err)
	}

	poolConfig.MaxConns = conf.DatabaseMaxConn
	poolConfig.MinConns = conf.DatabaseMinConn
	poolConfig.MaxConnLifetime = conf.DatabaseMaxConnLifetime * time.Minute
	poolConfig.MaxConnIdleTime = conf.DatabaseMaxConnIdletime * time.Minute
	poolConfig.HealthCheckPeriod = conf.DatabaseHealthCheckPeriod * time.Minute

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	return pool, nil
}

func CreateDatabase(conf *config.Config) error {
	baseConnString := conf.DatabaseUrl()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	basePool, err := pgxpool.New(ctx, baseConnString)
	if err != nil {
		return fmt.Errorf("failed to connect to postgres database: %w", err)
	}
	defer basePool.Close()

	var exists bool
	err = basePool.QueryRow(ctx, "SELECT EXISTS(SELECT 1 FROM pg_database WHERE datname = $1)", conf.DatabaseName).Scan(&exists)
	if err != nil {
		return fmt.Errorf("failed to check if database exists: %w", err)
	}

	if !exists {
		_, err = basePool.Exec(ctx, fmt.Sprintf("CREATE DATABASE %s", conf.DatabaseName))
		if err != nil {
			return fmt.Errorf("failed to createe database: %w", err)
		}
	}

	return nil
}
