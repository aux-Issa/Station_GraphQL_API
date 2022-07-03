package utils

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

// db接続: https://qiita.com/taka23kz/items/cb7ae9ac6cf343b3dec2

func openDB(dbDriver string, dsn string) (*sql.DB, error) {
	db, err := sql.Open(dbDriver, dsn)
	if err != nil {
		return nil, err
	}
	return db, err
}

func SetupDB() (*sql.DB, error) {
	dbDriver := "postgres"
	dsn := "host=192.168.11.11 port=5432 user=postgres password=postgres dbname=postgres sslmode=disable"
	db, err := openDB(dbDriver, dsn)

	// defer db.Close()

	if err != nil {
		log.Fatal(err)
	}
	return db, err
}
