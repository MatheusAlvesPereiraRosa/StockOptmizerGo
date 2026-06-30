package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"

	"gear-priority-api/internal/config"
	"gear-priority-api/internal/repository/mongodb"
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

	repository := mongodb.NewGearRepository(client)
	ctx := context.Background()

	command := "insert"

	if len(os.Args) > 1 {
		command = os.Args[1]
	}

	switch command {
	case "insert":
		runInsert(ctx, repository)

	case "clear":
		runClear(ctx, repository)

	default:
		log.Fatalf("Unknown command: %s. Use: insert or clear", command)
	}
}

func runInsert(ctx context.Context, repository *mongodb.MongoGearRepository) {
	total := 300

	if len(os.Args) > 2 {
		value, err := strconv.Atoi(os.Args[2])
		if err == nil && value > 0 {
			total = value
		}
	}

	criticalLimit := int(float64(total) * 0.20)
	mediumLimit := criticalLimit + int(float64(total)*0.40)

	random := rand.New(rand.NewSource(42))

	fmt.Printf("Generating %d gears...\n\n", total)

	inserted := 0

	for i := 0; i < total; i++ {
		gear := generateGear(
			random,
			i,
			criticalLimit,
			mediumLimit,
		)

		if err := repository.Create(ctx, &gear); err != nil {
			log.Printf("Error creating %s: %v\n", gear.Name, err)
			continue
		}

		inserted++
	}

	for _, gear := range edgeCases {
		if err := repository.Create(ctx, &gear); err != nil {
			log.Printf("Error creating %s: %v\n", gear.Name, err)
			continue
		}

		inserted++
	}

	fmt.Println("Seed completed successfully!")
	fmt.Printf("Inserted: %d\n", inserted)
}

func runClear(ctx context.Context, repository *mongodb.MongoGearRepository) {
	if err := repository.DeleteAll(ctx); err != nil {
		log.Fatal(err)
	}

	fmt.Println("All gears deleted successfully.")
}
