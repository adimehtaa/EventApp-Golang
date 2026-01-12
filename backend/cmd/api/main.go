package main

import (
	"app-event/internal/database"
	"app-event/internal/env"
	"database/sql"
	"log"
)

type application struct {
	port      int
	jwtSecret string
	model     database.Models
}

func main() {

	db, err := sql.Open("sqlite3", "./data.db")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	models := database.NewModels(db)

	app := &application{
		port:      env.GetEnvInteger("PORT", 8080),
		jwtSecret: env.GetEnvString("JWT_SECRET", "Some-Secret-808080"),
		model:     models,
	}

	if err := app.serve(); err != nil {
		log.Fatal(err)
	}
}
