package repositories

import (
	"context"
	"errors"
	"github.com/leyiqiang/home-library-server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

func (r *mongoRepo) GetBookByID(id string) (*models.Book, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var book models.Book
	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{
		"_id", objID,
	}}
	err := r.booksCollection.FindOne(ctx, filter).Decode(&book)
	if err == mongo.ErrNoDocuments {
		// Do something when no record was found
		return nil, errors.New("record does not exist")
	} else if err != nil {
		log.Fatal(err)
	}
	return &book, nil
}

func (r *mongoRepo) DeleteBookByID(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{
		"_id", objID,
	}}
	_, err := r.booksCollection.DeleteOne(ctx, filter)
	if err == mongo.ErrNoDocuments {
		// Do something when no record was found
		return err
	} else if err != nil {
		log.Fatal(err)
	}
	return nil
}

func (r *mongoRepo) GetAllBooks() ([]*models.Book, error) {
	var books []*models.Book

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	cur, err := r.booksCollection.Find(ctx, bson.D{})

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

func (r *mongoRepo) AddBook(book models.Book) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	res, err := r.booksCollection.InsertOne(ctx, &book)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	oid := res.InsertedID.(primitive.ObjectID)
	return oid.Hex(), nil
}

func (r *mongoRepo) UpdateBookByID(id string, newBookInfo models.Book) (*models.Book, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	objID, _ := primitive.ObjectIDFromHex(id)

	filter := bson.D{{
		"_id", objID,
	}}
	res := r.booksCollection.FindOneAndUpdate(ctx, filter, &newBookInfo)

	if res.Err() != nil {
		return nil, res.Err()
	}
	var book models.Book
	res.Decode(&book)
	return &book, nil
}
