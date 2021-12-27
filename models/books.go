package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Book struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Title       string             `json:"title" bson:"title" validate:"required""`
	Description string             `json:"description" bson:"description"`
	Year        int                `json:"year" bson:"year"`
	ReleaseDate time.Time          `json:"releaseDate" bson:"releaseDate"`
	Rating      int                `json:"rating bson:"rating"`
	CreatedAt   time.Time          `json:"-" json:"createdAt"`
	Genre       map[int]string     `json:"genres" bson:"genres"`
	ISBN        string             `json:"isbn" bson:"isbn"`
}
