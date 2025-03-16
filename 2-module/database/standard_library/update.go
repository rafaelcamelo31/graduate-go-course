package main

import "database/sql"

func updateProduct(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare("update Products set name = ?, price = ? where id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(product.Name, product.Price, product.ID)
	if err != nil {
		return err
	}
	return nil
}
