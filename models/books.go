package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Book struct {
	ID           primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Title        string             `json:"title" bson:"title" validate:"required""`
	Author       string             `json:"author" bson:"author"`
	Publisher    string             `json:"publisher" bson:"publisher"`
	ImportedDate time.Time          `json:"importedDate" json:"importedDate"`
	Location     string             `json:"location" bson:"location"`
	Category     string             `json:"category" bson:"category"`
	ISBN         string             `json:"isbn" bson:"isbn"`
	Description  string             `json:"-" bson:"-"`
}
