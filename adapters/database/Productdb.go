package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/nobruin/hexagonal-arch-example/app"
)

func NewProductDB(db *sql.DB) *ProductDB {
	return &ProductDB{
		db: db,
	}
}

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

func (p *ProductDB) Save(product app.ProductInterface) (app.ProductInterface, error) {

	if p.exists(product.GetID()) {
		_, err := p.update(product)
		if err != nil {
			return nil, err
		}
	} else {
		_, err := p.create(product)
		if err != nil {
			return nil, err
		}
	}
	return product, nil
}

func (p *ProductDB) create(product app.ProductInterface) (app.ProductInterface, error) {
	stmt, err := p.db.Prepare(`insert into products( id, name, price, status) values (?,?,?,?);`)
	if err != nil {
		return nil, err
	}
	_, err = stmt.Exec(product.GetID(), product.GetName(), product.GetPrice(), product.GetStatus())
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (p *ProductDB) update(product app.ProductInterface) (app.ProductInterface, error) {
	_, err := p.db.Exec(`UPDATE products SET name = ?, price = ?, status = ? WHERE id = ?;`, product.GetName(), product.GetPrice(), product.GetStatus(), product.GetID())
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (p *ProductDB) exists(id string) bool {
	var rows int
	err := p.db.QueryRow("select count(id) from products where id=?;", id).Scan(&rows)
	if err != nil {
		return false
	}
	return rows > 0
}
