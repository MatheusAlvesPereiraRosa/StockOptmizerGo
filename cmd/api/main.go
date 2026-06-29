package main

import (
	"fmt"
	"net/http"

	"gear-priority-api/internal/config"
	"gear-priority-api/internal/controller"
	"gear-priority-api/internal/handler"
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

	handler := handler.NewGearHandler(service)

	controller := controller.NewRouter(handler)

	fmt.Println("Gear Priority API")
	fmt.Printf("Port: %s\n", cfg.Port)
	fmt.Printf("Mongo URI: %s\n", cfg.MongoURI)

	log.Printf("Successfully connected to MongoDB!")
	log.Printf("Database: %s\n", client.DB.Name())

	log.Fatal(http.ListenAndServe(":"+cfg.Port, controller))
}
