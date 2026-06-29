package main

import (
	"fmt"
	"log"
	"net/http"

	"gear-priority-api/internal/config"
	"gear-priority-api/internal/controller"
	"gear-priority-api/internal/handler"
	"gear-priority-api/internal/repository/mongodb"
	"gear-priority-api/internal/service"
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

	// Repository
	repository := mongodb.NewGearRepository(client)

	// Services
	gearService := service.NewGearService(repository)
	restockService := service.NewRestockService(repository)

	// Handlers
	gearHandler := handler.NewGearHandler(gearService)
	restockHandler := handler.NewRestockHandler(restockService)

	// Router
	router := controller.NewRouter(
		gearHandler,
		restockHandler,
	)

	fmt.Println("Gear Priority API")
	fmt.Printf("Port: %s\n", cfg.Port)
	fmt.Printf("Mongo URI: %s\n", cfg.MongoURI)

	log.Println("Successfully connected to MongoDB!")
	log.Printf("Database: %s\n", client.DB.Name())

	log.Fatal(http.ListenAndServe(":"+cfg.Port, router))
}
