package main

import (
	"fmt"

	"gear-priority-api/internal/config"
	"gear-priority-api/internal/repository/mongodb"
	"gear-priority-api/internal/service"

	"log"
)

func main() {
	cfg, err := config.LoadConfig()

	client, _ := mongodb.New(cfg.MongoURI, cfg.MongoDatabase)

	repository := mongodb.NewGearRepository(client)

	service := service.NewGearService(repository)

	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Gear Priority API")
	fmt.Printf("Port: %s\n", cfg.Port)
	fmt.Printf("Mongo URI: %s\n", cfg.MongoURI)

	log.Printf("Successfully connected to MongoDB!")
	log.Printf("Database: %s\n", client.DB.Name())
}
