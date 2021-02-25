package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/wagaru/Recodar/server/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoRepo struct {
	client *mongo.Client
}

func NewMongoRepo(config *config.Config) (Repository, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.MongoURI))
	if err != nil {
		return nil, fmt.Errorf("connect mongo db failed: %w", err)
	}
	return &mongoRepo{
		client: client,
	}, nil
}
