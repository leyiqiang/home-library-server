package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Reservation struct {
	ID       primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Calendar primitive.ObjectID `json:"calendar" bson:"calendar,required"`
	Name     string             `json:"name" bson:"name" validate:"required"`
}
