package main

import (
	"context"
	"log"
	"project8/api"
	"project8/cookies"
	db "project8/db/sqlc"
	"project8/firebaseadmin"
	"project8/utils"

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
	cookies, err := cookies.NewSecureCookie(config.CookieSecret)
	if err != nil {
		log.Fatalf("cannot create secure cookie")
	}

	firebaseAdmin, err := firebaseadmin.NewFirebaseAdmin(config.FirebaseServiceAccount)
	if err != nil {
		log.Fatalf("cannot create firebase admin %v", err)
	}

	runGinServer(config, store, cookies, firebaseAdmin)
}

func runGinServer(
	config utils.Config,
	store db.Store,
	cookies cookies.Cookies,
	firebaseAdmin *firebaseadmin.FirebaseAdmin,
) {
	server, err := api.NewServer(config, store, cookies, firebaseAdmin)

	if err != nil {
		log.Fatalf("cannot create server: %v", err)
	}

	err = server.Start(config.HttpServerAddress)

	if err != nil {
		log.Fatalf("cannot start server: %v", err)
	}
}
