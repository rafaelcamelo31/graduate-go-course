package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type Product struct {
	ID    string
	Name  string
	Price float64
}

func NewProduct(name string, price float64) *Product {
	return &Product{
		ID:    uuid.New().String(),
		Name:  name,
		Price: price,
	}
}

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	product := NewProduct("basketball", 259.99)
	result, err := insertProduct(db, product)
	if err != nil {
		panic(err)
	}

	result.Price = 369.99
	err = updateProduct(db, result)
	if err != nil {
		panic(err)
	}

	p, err := selectProduct(db, result.ID)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Product: %v, Price: %.2f\n", p.Name, p.Price)

	ps, err := selectAllProducts(db)
	if err != nil {
		panic(err)
	}
	for _, p := range ps {
		fmt.Printf("Product: %v, Price: %.2f\n", p.Name, p.Price)
	}

	err = deleteProduct(db, result.ID)
	if err != nil {
		panic(err)
	}
}
