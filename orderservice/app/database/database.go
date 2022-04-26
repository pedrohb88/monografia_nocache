package database

import (
	"database/sql"
	"os"

	"github.com/go-gorp/gorp"
)

func New() (gorp.DbMap, error) {

	// Database
	db, err := sql.Open("mysql", os.Getenv("DBCONN"))
	if err != nil {
		return gorp.DbMap{}, err
	}

	dbmap := gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{}}
	return dbmap, nil
}
