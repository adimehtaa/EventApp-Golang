package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/sqlite3"
	"github.com/golang-migrate/migrate/source/file"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Plese provie a migration direction 'up' or 'down' ")
	}

	direction := os.Args[1]

	db, err := sql.Open("sqlite3", "./data.db")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	instance, err := sqlite3.WithInstance(db, &sqlite3.Config{})

	if err != nil {
		log.Fatal(err)
	}

	fSrc, err := (&file.File{}).Open("cmd/migrate/migrations")

	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithInstance("file", fSrc, "sqlite3", instance)

	if err != nil {
		log.Fatal(err)
	}

	switch direction {
	case "up":
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}

	case "down":
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}

	default:
		log.Fatal("Invalid migration direction. Use 'up' or 'down'")
	}

}

/**
# run this command to generate migration file
	(Golang migrate need to be install starting migration)
	- migrate create -ext sql -dir ./cmd/migrate/migrations -seq table_name
	for example
	: migrate create -ext sql -dir ./cmd/migrate/migrations -seq create_users_table

*/
