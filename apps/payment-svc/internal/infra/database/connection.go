package database

import (
	"context"
	"fmt"
	"os"

	"github.com/buemura/event-driven-commerce/payment-svc/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	Conn *pgxpool.Pool
)

func Connect() {
	dbConfig, err := pgxpool.ParseConfig(config.DATABASE_URL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create pool config: %v\n", err)
		os.Exit(1)
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), dbConfig)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	Conn = pool
}
