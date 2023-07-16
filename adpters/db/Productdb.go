package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/nobruin/hexagonal-arch-example/app"
)

type ProductDB struct {
	db *sql.DB
}

func (p *ProductDB) Get(id string) (app.ProductInterface, error) {
	var product app.Product
	stmt, err := p.db.Prepare("select id, name, price, status from products where id=?")
	if err != nil {
		return nil, err
	}
	err = stmt.QueryRow(id).Scan(&product.ID, &product.Name, &product.Price, &product.Status)
	if err != nil {
		return nil, err
	}
	return &product, nil
}
