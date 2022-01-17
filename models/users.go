package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Username string             `json:"username" bson:"username" validate:"required,alphanum""`
	Name     string             `json:"name" bson:"name" validate:"alphanumunicode"`
	Password string             `json:"password" bson:"password" validate:"required"`
	IsAdmin  bool               `json:"isAdmin" bson:"isAdmin" validate:"required"`
}
