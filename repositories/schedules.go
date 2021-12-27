package repositories

import "go.mongodb.org/mongo-driver/bson/primitive"

type Schedule struct {
	ID       primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Username string             `json:"username" bson:"username" validate:"required,alphanum""`
	Time     string             `json:"name" bson:"time"`
}
