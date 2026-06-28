package mongodb

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Client struct {
	DB *mongo.Database
}

func New(uri string, database string) (*Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(options.Client().ApplyURI(uri))

	if err != nil {
		return nil, err
	}

	/* Verify the connection */
	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	return &Client{
		DB: client.Database(database),
	}, nil
}
