package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product
}

type Product struct {
	ID           int `gorm:"primaryKey"`
	Name         string
	Price        float64
	CategoryID   int
	Category     Category
	SerialNumber SerialNumber
	gorm.Model
}

type SerialNumber struct {
	ID        int `gorm:"primaryKey"`
	Number    int
	ProductID int
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// Should be used only in Dev environment
	db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})

	// Create if not exists Category
	category := new(Category)
	category.Name = "Electronics"
	db.FirstOrCreate(category, Category{Name: "Electronics"})

	// Create if not exists Product
	product := new(Product)
	product.Name = "Calculator"
	product.Price = 850
	product.CategoryID = category.ID
	db.FirstOrCreate(product, &Product{Name: "Calculator"})
	products := []Product{
		{Name: "Mouse", Price: 18, CategoryID: category.ID},
		{Name: "Keyboard", Price: 32, CategoryID: category.ID},
		{Name: "Monitor", Price: 200, CategoryID: category.ID},
	}

	// Create if not exists SerialNumber
	serialNumber := new(SerialNumber)
	serialNumber.Number = 123
	serialNumber.ProductID = product.ID
	db.FirstOrCreate(serialNumber, SerialNumber{Number: 123})

	// Find or create
	var found []Product
	result := db.Find(&found)
	if result.RowsAffected < int64(len(products)) {
		// Create in batches
		db.CreateInBatches(products, 50)
	}

	err = db.Model(&category).Preload("Products.SerialNumber").Find(&category).Error
	if err != nil {
		panic(err)
	}

	for _, product := range category.Products {
		fmt.Printf("%s %s %.2f %d\n", category.Name, product.Name, product.Price, product.SerialNumber.Number)
	}

	// Locking
	tx := db.Begin()
	newProduct := new(Product)
	err = tx.Debug().Clauses(clause.Locking{Strength: "UPDATE"}).First(newProduct, "id = ?", product.ID).Error
	if err != nil {
		tx.Rollback()
		panic(err)
	}
	newProduct.Price = 2000
	newProduct.Name = "New Product"
	newProduct.CategoryID = category.ID
	tx.Debug().Save(newProduct)
	tx.Commit()

}
