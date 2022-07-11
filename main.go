package main

import (
	"log"
	"os"

	api "github.com/capungkoneng/gomcommerce/server"
	"github.com/capungkoneng/gomcommerce/util"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config", err)
	}

	// conn, err := sql.Open(config.DBDriver, config.DBSource)
	// if err != nil {
	// 	log.Fatal("cannot connect to database", err)
	// }

	// store := db.NewStore(conn)
	server, err := api.NewServer(config)
	if err != nil {
		log.Fatal("cannot create server", err)
	}

	err = server.Start(os.Getenv("PORT"))
	if err != nil {
		log.Fatal("cannot start server", err)
	}
}
