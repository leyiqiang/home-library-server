package repositories

import (
	"context"
	"fmt"
	"github.com/leyiqiang/home-library-server/config"
	"github.com/leyiqiang/home-library-server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

const BooksCollection = "Books"

var collection = new(mongo.Collection)

func (r *mongoRepo) init() {
	var cfg config.Config
	cfg.Read()
	collection = r.Client.Database(cfg.Database.Name).Collection(BooksCollection)
}

func (r *mongoRepo) GetBookByID(id int) (*models.Book, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var book models.Book
	filter := bson.D{{
		"_id", id,
	}}
	err := collection.FindOne(ctx, filter).Decode(&book)
	if err == mongo.ErrNoDocuments {
		// Do something when no record was found
		fmt.Println("record does not exist")
		return nil, err
	} else if err != nil {
		log.Fatal(err)
	}

	return &book, nil
}

func (r *mongoRepo) GetAllBooks() ([]*models.Book, error) {
	var books []*models.Book

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	cur, err := collection.Find(ctx, bson.D{})

	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var book *models.Book
		err := cur.Decode(&book)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		books = append(books, book)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return books, nil
}

func (r *mongoRepo) AddBook(book models.Book) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	_, err := collection.InsertOne(ctx, &book)

	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
