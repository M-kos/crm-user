package db

import (
	"context"
	"fmt"
	"net"
	"net/url"

	"github.com/M-kos/crm-user/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

const DBScheme = "postgres"

type PostgressDB struct {
	Pool *pgxpool.Pool
}

func NewDB(ctx context.Context, config *config.Config) (*PostgressDB, error) {
	pool, err := MakePostgresPool(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("postgres connection: %w", err)
	}

	defer pool.Close()

	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("ping postgres: %w", err)
	}

	fmt.Println("Connected to database")

	return &PostgressDB{
		Pool: pool,
	}, nil
}

func MakePostgresPool(ctx context.Context, config *config.Config) (*pgxpool.Pool, error) {
	connString := ConnectionString(config)

	pool, err := pgxpool.New(ctx, connString)
	if err != nil {
		return nil, err
	}

	return pool, nil
}

func ConnectionString(config *config.Config) string {
	query := url.Values{}
	query.Set("sslmode", "disable")

	value := url.URL{
		Scheme:   DBScheme,
		User:     url.UserPassword(config.Postgres.User, config.Postgres.Password),
		Host:     net.JoinHostPort(config.Postgres.Host, config.Postgres.Port),
		Path:     config.Postgres.Name,
		RawQuery: query.Encode(),
	}

	return value.String()
}
