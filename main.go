package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/suryadharma5/gobank/api"
	db "github.com/suryadharma5/gobank/db/sqlc"
	"github.com/suryadharma5/gobank/util"
)

func main() {
	conf, err := util.LoadConfig(".") // pas location dari .env
	if err != nil {
		log.Fatal("Cannot load config", err)
	}

	conn, err := sql.Open(conf.DBDriver, conf.DBSource)

	if err != nil {
		log.Fatal("Can't connect to db: ", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(conf, store)
	if err != nil {
		log.Fatal("Failed to create server")
	}

	err = server.Start(conf.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
