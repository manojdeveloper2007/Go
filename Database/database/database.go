package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Opendb(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		log.Fatal("database failed to connect", err)
		return nil, err
	}

	if err = db.Ping(); err != nil {
		log.Fatal("Something wrong to communicate between database", err)
		return nil, err
	}

	return db, nil
}
