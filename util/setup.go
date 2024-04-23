package util

import (
	"database/sql"
	"log"
)

func SetUpConnection(path string) (*sql.DB) {
	config, err := LoadConfig(path)
	if err != nil {
		log.Fatal("cannot load config:", err)
		return nil
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
		return nil
	}

	return conn
}