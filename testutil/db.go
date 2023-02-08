package testutil

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func OpenDBForTest(t *testing.T) *sqlx.DB {
	t.Helper()

	port := 33306
	if _, defined := os.LookupEnv("CI"); defined {
		port = 3306
	}
	db, err := sql.Open(
		"mysql",
		fmt.Sprintf("todo:todo@tcp(127.0.0.1:%d)/todo?parseTime=true", port),
	)

	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(
		func() { _ = db.Close },
	)

	return sqlx.NewDb(db, "mysql")

}
