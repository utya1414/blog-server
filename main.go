package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/utya1414/blog-server/api"
	"github.com/utya1414/blog-server/infrastructure/db"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:password@localhost:5434/blog?sslmode=disable"
	serveraddress = "0.0.0.0:8080"
)

func main() {
	// config, err := util.LoadConfig(".")
	// if err != nil {
	// 	log.Fatal("cannot load config:", err)
	// }
	
	conn, err := sql.Open(dbDriver, dbSource);
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	router, err := api.NewServer(store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	err = router.Run(serveraddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}