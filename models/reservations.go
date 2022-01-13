package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Reservation struct {
	ID       primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Schedule primitive.ObjectID `json:"schedule" bson:"schedule" validate:"required"`
	Name     string             `json:"name" bson:"name" validate:"required"`
}
