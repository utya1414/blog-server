package sqlc

import (
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/utya1414/blog-server/util"
)



var testQueries *Queries;

func TestMain(m *testing.M) {
	conn := util.SetUpConnection("../../../../");
	testQueries = New(conn);

	os.Exit(m.Run());
}