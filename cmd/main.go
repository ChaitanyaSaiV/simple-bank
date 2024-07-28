package main

import (
	"context"
	"log"

	"github.com/ChaitanyaSaiV/simple-bank/api"
	db "github.com/ChaitanyaSaiV/simple-bank/internal/db/methods"
	"github.com/ChaitanyaSaiV/simple-bank/util"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Unable to load env variables")
	}

	conn, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal("Unable to establish the connection")
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("Cannot Start the Server: ", err)
	}
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Error Starting the GIN server")
	}
}
