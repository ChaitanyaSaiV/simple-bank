package main

import (
	"context"
	"log"

	"github.com/ChaitanyaSaiV/simple-bank/api"
	db "github.com/ChaitanyaSaiV/simple-bank/internal/db/methods"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	conn, err := pgxpool.New(context.Background(), dbSource)
	if err != nil {
		log.Fatal("Unable to establish the connection")
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("Error Starting the GIN server")
	}
}
