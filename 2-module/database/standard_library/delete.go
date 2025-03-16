package main

import "database/sql"

func deleteProduct(db *sql.DB, id string) error {
	stmnt, err := db.Prepare("delete from Products where id = ?")
	if err != nil {
		return err
	}
	defer stmnt.Close()

	_, err = stmnt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}
