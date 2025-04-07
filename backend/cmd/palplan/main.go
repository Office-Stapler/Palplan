package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Office-Stapler/Palplan/backend/srv/config"
	"github.com/Office-Stapler/Palplan/backend/srv/db"
	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("Starting PalPlan...")

	log.Println("Loading env config...")
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load env config:", err)
	}
	log.Println("Config:", config)

	db, err := db.NewDB(context.Background(), config)
	if err != nil {
		log.Fatal("Error connecting to DB.")
	}
	log.Println("Database: ", db)
	defer db.Close()

	r := gin.Default()
	err = r.Run(fmt.Sprintf(":%d", config.Port))
	if err != nil {
		panic("[Error] Gin server failed to start")
	}

}
