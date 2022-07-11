package main

import (
	"io"
	"log"
	"net/http"
	"os"

	// api "github.com/capungkoneng/gomcommerce/server"
	// "github.com/capungkoneng/gomcommerce/util"
	_ "github.com/lib/pq"
)

func main() {

	port := os.Getenv("PORT")
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	}
	http.HandleFunc("/", helloHandler)
	log.Println("Listing for" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

	// config, err := util.LoadConfig(".")
	// if err != nil {
	// 	log.Fatal("cannot load config", err)
	// }

	// // conn, err := sql.Open(config.DBDriver, config.DBSource)
	// // if err != nil {
	// // 	log.Fatal("cannot connect to database", err)
	// // }

	// // store := db.NewStore(conn)
	// server, err := api.NewServer(config)
	// if err != nil {
	// 	log.Fatal("cannot create server", err)
	// }

	// err = server.Start(os.Getenv("PORT"))
	// if err != nil {
	// 	log.Fatal("cannot start server", err)
	// }
}
