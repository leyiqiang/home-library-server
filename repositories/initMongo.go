package repositories

import (
	"context"
	"github.com/leyiqiang/home-library-server/config"
	"github.com/leyiqiang/home-library-server/interfaces"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type mongoRepo struct {
	Client *mongo.Client
}

var cfg config.Config

func NewMongoRepo() interfaces.Repository {
	cfg.Read()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.Database.URI))
	if err != nil {
		log.Fatal(err)
	}

	return &mongoRepo{
		Client: client,
	}
}
