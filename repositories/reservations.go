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

// TODO remember to populate!
func (r *mongoRepo) GetReservationsByID(id string) (*models.Reservation, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var reservation models.Reservation
	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{
		"_id", objID,
	}}
	err := r.reservationsCollection.FindOne(ctx, filter).Decode(&reservation)
	if err == mongo.ErrNoDocuments {
		// Do something when no record was found
		return nil, errors.New("record does not exist")
	} else if err != nil {
		log.Fatal(err)
	}
	return &reservation, nil
}

// TODO : be careful! need to delete reservations as well
func (r *mongoRepo) DeleteReservationByID(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{
		"_id", objID,
	}}
	_, err := r.reservationsCollection.DeleteOne(ctx, filter)
	if err == mongo.ErrNoDocuments {
		// Do something when no record was found
		return err
	} else if err != nil {
		log.Fatal(err)
	}
	return nil
}

// TODO remember to populate!
func (r *mongoRepo) GetReservationsByScheduleID(id string) ([]*models.Reservation, error) {
	var reservations []*models.Reservation

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{
		"schedule", objID,
	}}
	cur, err := r.reservationsCollection.Find(ctx, filter)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var reservation *models.Reservation
		err := cur.Decode(&reservation)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		reservations = append(reservations, reservation)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}
	return reservations, nil
}

func (r *mongoRepo) AddReservation(reservation models.Reservation) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	res, err := r.reservationsCollection.InsertOne(ctx, &reservation)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	oid := res.InsertedID.(primitive.ObjectID)
	return oid.Hex(), nil
}

func (r *mongoRepo) UpdateReservationByID(id string, newData models.Reservation) (*models.Reservation, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	objID, _ := primitive.ObjectIDFromHex(id)

	filter := bson.D{{
		"_id", objID,
	}}

	update := bson.M{
		"$set": newData,
	}
	res := r.reservationsCollection.FindOneAndUpdate(ctx, filter, update)

	if res.Err() != nil {
		return nil, res.Err()
	}
	var reservation models.Reservation
	res.Decode(&reservation)
	return &reservation, nil
}
