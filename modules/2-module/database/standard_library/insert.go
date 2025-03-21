package main

import "database/sql"

func insertProduct(db *sql.DB, product *Product) (*Product, error) {
	stmt, err := db.Prepare("insert into Products(id, name, price) values(?,?,?)")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(product.ID, product.Name, product.Price)
	if err != nil {
		return nil, err
	}

	return product, nil
}
