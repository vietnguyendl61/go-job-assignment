package main

import (
	"github.com/joho/godotenv"
	"testing"
)

const tableCreationQuery = `CREATE TABLE IF NOT EXISTS products
(
    id SERIAL,
    name TEXT NOT NULL,
    price NUMERIC(10,2) NOT NULL DEFAULT 0.00,
    CONSTRAINT products_pkey PRIMARY KEY (id)
)`

func TestInitDB(t *testing.T) {
	_ = godotenv.Load()
	db := InitDB()

	t.Run("connect to database", func(t *testing.T) {
		if err := db.Exec(tableCreationQuery).Error; err != nil {
			t.Errorf("Error when connect to db: " + err.Error())
		}
		db.Exec("DELETE FROM products")
		db.Exec("ALTER SEQUENCE products_id_seq RESTART WITH 1")
	})
}
