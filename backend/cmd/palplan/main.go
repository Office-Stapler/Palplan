package main

import (
	"context"
	"log"

	"github.com/Office-Stapler/Palplan/backend/srv/config"
	"github.com/Office-Stapler/Palplan/backend/srv/db"
)

func main() {
	log.Println("Starting PalPlan...")

	log.Println("Loading env config...")
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load env config:", err)
	}

	db, err := db.NewDB(context.Background(), config)
	if err != nil {
		log.Fatal("Error connecting to DB.")
	}
	defer db.Close()

	log.Println("Database: ", db)
	log.Println("Config:", config)
}
