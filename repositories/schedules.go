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

func (r *mongoRepo) GetScheduleByID(id string) (*models.Schedule, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var schedule models.Schedule
	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{
		"_id", objID,
	}}
	err := r.schedulesCollection.FindOne(ctx, filter).Decode(&schedule)
	if err == mongo.ErrNoDocuments {
		// Do something when no record was found
		return nil, errors.New("record does not exist")
	} else if err != nil {
		log.Fatal(err)
	}
	return &schedule, nil
}

func (r *mongoRepo) DeleteScheduleByID(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{
		"_id", objID,
	}}
	_, err := r.schedulesCollection.DeleteOne(ctx, filter)
	if err == mongo.ErrNoDocuments {
		// Do something when no record was found
		return err
	} else if err != nil {
		log.Fatal(err)
	}
	return nil
}

func (r *mongoRepo) GetAllSchedules() ([]*models.Schedule, error) {
	var schedules []*models.Schedule

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	cur, err := r.schedulesCollection.Find(ctx, bson.D{})

	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var schedule *models.Schedule
		err := cur.Decode(&schedule)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		schedules = append(schedules, schedule)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}
	return schedules, nil
}

func (r *mongoRepo) AddSchedule(schedule models.Schedule) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	res, err := r.schedulesCollection.InsertOne(ctx, &schedule)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	oid := res.InsertedID.(primitive.ObjectID)
	return oid.Hex(), nil
}

func (r *mongoRepo) UpdateScheduleByID(id string, newData models.Schedule) (*models.Schedule, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	objID, _ := primitive.ObjectIDFromHex(id)

	filter := bson.D{{
		"_id", objID,
	}}

	update := bson.M{
		"$set": newData,
	}
	res := r.schedulesCollection.FindOneAndUpdate(ctx, filter, update)

	if res.Err() != nil {
		return nil, res.Err()
	}
	var schedule models.Schedule
	res.Decode(&schedule)
	return &schedule, nil
}
