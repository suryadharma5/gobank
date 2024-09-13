package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/suryadharma5/gobank/api"
	db "github.com/suryadharma5/gobank/db/sqlc"
)

const (
	dbDriver   = "postgres"
	dbSource   = "postgresql://postgres:postgres@localhost:5500/gobank?sslmode=disable"
	serverAddr = "127.0.0.1:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("Can't connect to db: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddr)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
