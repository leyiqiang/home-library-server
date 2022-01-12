package repositories

import (
	"context"
	"github.com/leyiqiang/home-library-server/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

const BooksCollectionString = "Books"
const UsersCollectionString = "Users"
const SchedulesCollectionString = "Schedules"
const reservationsCollectionString = "Reservations"

type mongoRepo struct {
	booksCollection        *mongo.Collection
	schedulesCollection    *mongo.Collection
	reservationsCollection *mongo.Collection
	usersCollection        *mongo.Collection
}

var cfg config.Config

func NewMongoRepo() Repository {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cfg.Read()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.Database.URI).SetAuth(
		options.Credential{
			Username: cfg.Database.Username,
			Password: cfg.Database.Password,
		}))
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	database := client.Database(cfg.Database.Name)
	booksCollection := database.Collection(BooksCollectionString)
	usersCollection := database.Collection(UsersCollectionString)
	schedulesCollection := database.Collection(SchedulesCollectionString)
	reservationsCollection := database.Collection(reservationsCollectionString)

	//defer func() {
	//	if err = client.Disconnect(ctx); err != nil {
	//		panic(err)
	//	}
	//}()
	return &mongoRepo{
		booksCollection:        booksCollection,
		schedulesCollection:    schedulesCollection,
		reservationsCollection: reservationsCollection,
		usersCollection:        usersCollection,
	}
}
