package repository

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/utya1414/blog-server/infrastructure/db"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:password@localhost:5434/blog?sslmode=disable"
)

// var testQueries *sqlc.Queries;
var r UserRepository;

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource);
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	// testQueries = sqlc.New(conn);
	store := db.NewStore(conn);
	r = NewUserRepository(store);
	os.Exit(m.Run());
}