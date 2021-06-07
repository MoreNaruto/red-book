package database

import (
	"database/sql"
	"log"
	"redbook/util"

	_ "github.com/go-sql-driver/mysql"
)

func StartDB() (*sql.DB, error) {
	config, err := util.LoadConfig("../../config/app.env")
	if err != nil {
		log.Printf("Error %s when trying to access config\n", err)
		return nil, err
	}
	db, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Printf("Error %s when opening DB\n", err)
		return nil, err
	}

	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(20)

	err = db.Ping()
	if err != nil {
		log.Printf("Error %s when pinging DB\n", err)
		return nil, err
	}

	return db, nil
}
