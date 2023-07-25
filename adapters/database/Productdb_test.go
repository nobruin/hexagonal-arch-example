package database_test

import (
	"database/sql"
	"log"
	"testing"

	"github.com/nobruin/hexagonal-arch-example/adapters/database"
	"github.com/nobruin/hexagonal-arch-example/app"
	"github.com/stretchr/testify/require"
)

var Db *sql.DB

func setup() {
	Db, _ = sql.Open("sqlite3", ":memory:")
	createTable(Db)
	createProduct(Db)
}

func createTable(db *sql.DB) {
	table := `CREATE TABLE products ("id" string,"name" string,"price" float,"status" string);`
	stmt, err := db.Prepare(table)
	if err != nil {
		log.Fatal(err.Error())
	}
	stmt.Exec()
}

func createProduct(db *sql.DB) {
	insert := `insert into products values("abc","Product Test",0, "disabled");`
	stmt, err := db.Prepare(insert)
	if err != nil {
		log.Fatal(err.Error())
	}
	stmt.Exec()
}

func TestProductDB_Get(t *testing.T) {
	setup()
	defer Db.Close()
	productDB := database.NewProductDB(Db)
	product, err := productDB.Get("abc")
	require.Nil(t, err)

	require.Equal(t, "Product Test", product.GetName())
	require.Equal(t, 0.0, product.GetPrice())
	require.Equal(t, app.DISABLED, product.GetStatus())
}

func TestProductDB_Create(t *testing.T) {
	setup()
	defer Db.Close()
	product := app.NewProduct()
	productDB := database.NewProductDB(Db)
	product.ID = "a1"
	product.Name = "p1"
	product.Price = 10.0
	product.Status = app.ENABLED

	productSaved, err := productDB.Save(product)
	require.Nil(t, err)
	require.Equal(t, product, productSaved)
	product.Name = "p2"
	productUpdated, errUpdate := productDB.Save(product)
	require.Nil(t, errUpdate)
	require.Equal(t, product, productUpdated)
}
