package repository

import (
	"os"
	"testing"

	"github.com/utya1414/blog-server/infrastructure/db"

	_ "github.com/lib/pq"
	userDomain "github.com/utya1414/blog-server/domain/user"
	"github.com/utya1414/blog-server/util"
)

// var testQueries *sqlc.Queries;
var r userDomain.UserRepository;

func TestMain(m *testing.M) {
	conn := util.SetUpConnection("../../../");

	// testQueries = sqlc.New(conn);
	store := db.NewStore(conn);
	r = NewUserRepository(store);
	os.Exit(m.Run());
}