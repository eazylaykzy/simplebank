package main

import (
	"database/sql"
	"github.com/eazylaykzy/simplebank/api"
	db "github.com/eazylaykzy/simplebank/db/sqlc"
	"github.com/eazylaykzy/simplebank/util"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	config, err := util.LoadConfig(".")
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
