package repository

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/wagaru/Recodar/server/internal/config"
	"github.com/wagaru/Recodar/server/internal/domain"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	MONGO_DATABASE        = "recodar"
	MONGO_USER_COLLECTION = "users"
)

type Repository interface {
	// Disconnect
	Disconnect()

	// User
	GetUser(ctx context.Context, key string, value interface{}) (*domain.User, error)
	StoreUser(ctx context.Context, u *domain.User) (interface{}, error)
	UpdateUser(ctx context.Context, id string, u *domain.User) error
}
type mongoRepo struct {
	client *mongo.Client
	db     *mongo.Database
}

func NewMongoRepo(config *config.Config) (Repository, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.MongoURI))
	if err != nil {
		return nil, fmt.Errorf("Connect mongo db failed: %w", err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, fmt.Errorf("Ping mongo db failed: %w", err)
	}

	// collection := client.Database("sample_airbnb").Collection("listingsAndReviews")
	// result := bson.M{}
	// err = collection.FindOne(context.Background(), bson.M{"_id": "10006546"}).Decode(&result)
	// if errors.Is(err, mongo.ErrNoDocuments) {
	// 	fmt.Println("yaaa")
	// }
	return &mongoRepo{
		client: client,
		db:     client.Database(MONGO_DATABASE),
	}, nil
}

func (repo *mongoRepo) Disconnect() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	if err := repo.client.Disconnect(ctx); err != nil {
		log.Fatal(err)
	}
}
