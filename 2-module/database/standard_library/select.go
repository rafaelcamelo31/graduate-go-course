package main

import "database/sql"

func selectProduct(db *sql.DB, id string) (*Product, error) {
	stmt, err := db.Prepare("select id, name, price from Products where id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	p := Product{}
	err = stmt.QueryRow(id).Scan(&p.ID, &p.Name, &p.Price)
	if err != nil {
		return nil, err
	}

	return &p, nil
}
