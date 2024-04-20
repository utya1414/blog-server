package repository

import (
	"database/sql"
	"log"
	"os"
	"testing"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:password@localhost:5434/blog?sslmode=disable"
)

var testQueries *Queries;

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource);
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(conn);

	os.Exit(m.Run());
}