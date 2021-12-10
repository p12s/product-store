package repository

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Config - db
type Config struct {
	Username, Password, Url string
}

// NewMongoDB - open connect and ping trying
func NewMongoDB(ctx context.Context, cfg Config) (*mongo.Client, error) {
	clientOptions := options.Client()
	clientOptions.SetAuth(options.Credential{
		Username: cfg.Username,
		Password: cfg.Password,
	})
	clientOptions.ApplyURI(cfg.Url)

	dbClient, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("mongodb connect: %w", err)
	}

	if err := dbClient.Ping(context.Background(), nil); err != nil {
		return nil, fmt.Errorf("mongodb ping: %w", err)
	}

	return dbClient, nil
}
