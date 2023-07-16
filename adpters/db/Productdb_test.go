package db

import (
	"database/sql"
	"log"
)

var DB *sql.DB

func setup() {
	DB, _ = sql.Open("sqlite3", ":memory")
	createTable(DB)
	createProduct(DB)
}

func createTable(db *sql.DB) {
	queryCreateTable := `CREATE TABLE products( "id" string, "name" string, "price" float,"status" bolean);`
	stmt, err := db.Prepare(queryCreateTable)
	if err != nil {
		log.Fatal(err.Error())
	}
	stmt.Exec()
}

func createProduct(db *sql.DB) {
	queryCreateTable := `INSERT INTO products values ( "id1", "P1", "price", 0, false);`
	stmt, err := db.Prepare(queryCreateTable)
	if err != nil {
		log.Fatal(err.Error())
	}
	stmt.Exec()
}
