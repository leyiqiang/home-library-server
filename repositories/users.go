package repositories

import (
	"context"
	"github.com/leyiqiang/home-library-server/models"
	"log"
	"time"
)

func (r *mongoRepo) Register(user models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	_, err := r.usersCollection.InsertOne(ctx, &user)

	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
