package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitPostgres(dbUrl string) error {
	var err error
	DB, err = sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal("Unable to connect to DB:", err)
		return err
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Cannot reach database:", err)
	}
	log.Println("✅ Connected to PostgreSQL!")

	return nil

}
