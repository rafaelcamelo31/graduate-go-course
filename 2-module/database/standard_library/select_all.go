package main

import "database/sql"

func selectAllProducts(db *sql.DB) ([]Product, error) {
	rows, err := db.Query("select * from Products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := []Product{}
	for rows.Next() {
		product := Product{}
		err = rows.Scan(&product.ID, &product.Name, &product.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}
