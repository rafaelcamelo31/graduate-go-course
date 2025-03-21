package main

import (
	"database/sql"
	"log"
	"net/http"
)

type BaseHandler struct {
	DB *sql.DB
}

func NewBaseHandler(db *sql.DB) *BaseHandler {
	return &BaseHandler{
		DB: db,
	}
}

func main() {
	log.Println("Server starting at :8080")

	db := bootstrapSQLite()
	defer db.Close()

	bh := NewBaseHandler(db)

	http.HandleFunc("/exchange-rate", bh.exchangeRateHandler)
	http.ListenAndServe(":8080", nil)
}
