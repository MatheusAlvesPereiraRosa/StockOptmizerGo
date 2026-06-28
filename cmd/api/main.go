package main

import (
	"fmt"

	"gear-priority-api/internal/config"
	"gear-priority-api/internal/repository/mongodb"

	"log"
)

func main() {
	cfg, err := config.LoadConfig()

	if err != nil {
		log.Fatal(err)
	}

	client, err := mongodb.New(cfg.MongoURI, cfg.MongoDatabase)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Gear Priority API")
	fmt.Printf("Port: %s\n", cfg.Port)
	fmt.Printf("Mongo URI: %s\n", cfg.MongoURI)

	log.Printf("Successfully connected to MongoDB!")
	log.Printf("Database: %s\n", client.DB.Name())
}
