package main

import (
	"EcomerceProject/internal/env"
	"context"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5"
)

func main() {
	ctx := context.Background()
	cfg := config{
		addr: ":8080",
		db: dbConfig{
			dsn: env.GetString(
				"GOOSE_DBSTRING",
				"host=localhost port=5432 user=postgres password=postgres dbname=ecom sslmode=disable",
			),
		},
	}
	//loggin
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	conn, err := pgx.Connect(ctx, cfg.db.dsn)
	if err != nil {
		panic(err)
		// os.Exit(1)
	}
	defer conn.Close(ctx)

	logger.Info("connected to database", "dsn", cfg.db.dsn)
	api := aplicaction{
		config: cfg,
		db:     conn,
	}

	if err := api.run(api.mount()); err != nil {
		slog.Error("server failed to start", "error", err)
		os.Exit(1)
	}
}
