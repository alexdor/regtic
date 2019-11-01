package db

import (
	"database/sql"
	"os"

	"github.com/volatiletech/sqlboiler/boil"

	_ "github.com/lib/pq"
)

var DB boil.ContextExecutor

func init() {
	var err error
	dbUrl := os.Getenv("REGTIC_DATABASE_URL")
	// FIXME: Remove hardcoded creds from here
	if len(dbUrl) == 0 {
		dbUrl = "postgres://admin:admin@docker.for.mac.localhost:5432/regtic?sslmode=disable"
	}
	DB, err = sql.Open("postgres", dbUrl)
	if err != nil {
		panic(err)
	}
}
