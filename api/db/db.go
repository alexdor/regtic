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
	DB, err = sql.Open("postgres", os.Getenv("REGTIC_DATABASE_URL"))
	if err != nil {
		panic(err)
	}
}
