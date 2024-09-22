package config

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	migrate "github.com/rubenv/sql-migrate"
	"log"
	"os"
)

func initMigration(db *sql.DB) (err error) {
	migrations := &migrate.FileMigrationSource{
		Dir: os.Getenv("MIGRATION_PATH"),
	}

	n, err := migrate.Exec(db, "postgres", migrations, migrate.Up)
	if err != nil {
		log.Fatalf(`%v`, err)
		os.Exit(2)
	}
	fmt.Printf("Applied %d migrations!\n", n)

	return
}
