package main

import (
	"context"
	"log"
	"project9/api"
	db "project9/db/sqlc"
	"project9/utils"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	config, err := utils.LoadConfig(".")

	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	conn, err := pgxpool.New(context.Background(), config.DBSource)

	if err != nil {
		log.Fatalf("cannot connect to db")
	}

	store := db.NewStore(conn)

	runGinServer(config, store)
}

func runGinServer(config utils.Config, store db.Store) {
	server, err := api.NewServer(config, store)

	if err != nil {
		log.Fatalf("cannot create server: %v", err)
	}

	err = server.Start(config.HttpServerAddress)

	if err != nil {
		log.Fatalf("cannot start server: %v", err)
	}
}
