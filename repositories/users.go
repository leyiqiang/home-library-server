package repositories

import (
	"context"
	"errors"
	"github.com/leyiqiang/home-library-server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

func (r *mongoRepo) AddUser(user models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	_, err := r.usersCollection.InsertOne(ctx, &user)

	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (r *mongoRepo) FindUserByUsername(username string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var user models.User
	filter := bson.D{{
		"username", username,
	}}

	err := r.usersCollection.FindOne(ctx, filter).Decode(&user)

	if err == mongo.ErrNoDocuments {
		// Do something when no record was found
		return nil, errors.New("record does not exist")
	} else if err != nil {
		log.Fatal(err)
	}
	return &user, nil
}
