package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    int `gorm:"primaryKey"`
	Name  string
	Price float64
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// Should be used only in Dev environment
	db.AutoMigrate(&Product{})

	// Create if not exists
	db.FirstOrCreate(&Product{Name: "Laptop", Price: 1000}, &Product{Name: "Laptop"})

	products := []Product{
		{Name: "Mouse", Price: 18},
		{Name: "Keyboard", Price: 32},
		{Name: "Monitor", Price: 200},
	}

	// Find or create
	result := db.Find(&products)
	if result.RowsAffected == 0 {
		// Create in batches
		db.CreateInBatches(products, 50)
	}

}
