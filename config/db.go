package config

import (
	"database/sql"
	"fmt"
	// "github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func initDB() *sql.DB {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	connStr := fmt.Sprintf(`user=%s password=%s dbname=%s host=%s port=%s sslmode=disable search_path=%s`,
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DEFAULT_SEARCH_PATH"))
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec(fmt.Sprintf(`
		create schema if not exists sample_schema;
	`))

	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(fmt.Sprintf(`
		set search_path to sample_schema;
	`))

	if err != nil {
		log.Fatal(err)
	}

	return db
}
