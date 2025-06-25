package storage

import (
	"context"
	"fmt"

	"github.com/M-kos/crm-user/internal/config"
	"github.com/jackc/pgx/v5"
)

type Storage struct {
	Conn *pgx.Conn
}

func NewStorage(config *config.Config) *Storage {
	url := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", config.DbUser, config.DbPassword, config.DbHost, config.DbPort, config.DbName)
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, url)
	if err != nil {
		panic(err)
	}

	defer conn.Close(ctx)

	if err = conn.Ping(ctx); err != nil {
		panic(err)
	}

	fmt.Println("Connected to database")

	return &Storage{
		Conn: conn,
	}
}
