package database

import (
	"cmd/main/util"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func StartDB() (db *sql.DB) {
	config, err := util.LoadConfig("../../config/app.env")
	if err != nil {
		panic(err.Error())
	}
	db, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	return db
}
