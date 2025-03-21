package main

import (
	"context"
	"database/sql"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func bootstrapSQLite() *sql.DB {
	db := getConnection()

	createTable(db)

	return db
}

func getConnection() *sql.DB {
	db, err := sql.Open("sqlite3", "exchange_rate.db")
	if err != nil {
		log.Println(err)
		return nil
	}

	log.Println("Connected to SQLite")

	return db
}

func createTable(db *sql.DB) {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS exchange_rate (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		code TEXT,
		codein TEXT,
		name TEXT,
		high REAL,
		low REAL,
		varBid REAL,
		pctChange REAL,
		bid REAL,
		ask REAL,
		timestamp TEXT,
		createdAt TEXT
	)`)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("Table exchange_rate created")
}

func insertExchangeRate(ctx context.Context, db *sql.DB, ex *ExchangeRate) error {
	log.Printf("Inserting ExchangeRate: %+v\n", ex.USDBRL)

	stmt, err := db.Prepare(`insert into exchange_rate(
	code, 
	codein, 
	name, 
	high, 
	low, 
	varBid, 
	pctChange, 
	bid, 
	ask, 
	timestamp,
	createdAt
	) values(?,?,?,?,?,?,?,?,?,?,?)`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	ctx, cancel := context.WithTimeout(ctx, 10*time.Millisecond)
	defer cancel()

	_, err = stmt.ExecContext(ctx, ex.USDBRL.Code, ex.USDBRL.Codein, ex.USDBRL.Name, ex.USDBRL.High, ex.USDBRL.Low, ex.USDBRL.VarBid, ex.USDBRL.PctChange, ex.USDBRL.Bid, ex.USDBRL.Ask, ex.USDBRL.Timestamp, ex.USDBRL.CreatedAt)
	if err != nil {
		log.Println(err)
		return err
	}

	log.Println("ExchangeRate inserted")

	return nil
}
