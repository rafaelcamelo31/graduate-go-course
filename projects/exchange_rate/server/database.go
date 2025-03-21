package main

import (
	"database/sql"
	"log"
)

func getConnection() *sql.DB {
	db, err := sql.Open("sqlite3", "exchange_rate.db")
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer db.Close()

	return db
}
