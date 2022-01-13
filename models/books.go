package models

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strings"
	"time"
)

type Book struct {
	ID           primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Title        string             `json:"title" bson:"title" validate:"required""`
	Author       string             `json:"author" bson:"author"`
	Publisher    string             `json:"publisher" bson:"publisher"`
	ImportedDate ImportedDate       `json:"importedDate" bson:"importedDate"`
	Location     string             `json:"location" bson:"location"`
	Category     string             `json:"category" bson:"category"`
	ISBN         string             `json:"isbn" bson:"isbn"`
	Description  string             `json:"-" bson:"-"`
}

type ImportedDate struct {
	time.Time
}

const dateLayout = "2006-01-02"

func (i *ImportedDate) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	if s == "null" || s == "" {
		i.Time = time.Now()
		return nil
	}
	var err error
	i.Time, err = time.Parse(dateLayout, s)
	return err
}

func (i *ImportedDate) MarshalJSON() ([]byte, error) {
	if i.Time.UnixNano() == (time.Time{}).UnixNano() {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%s\"", i.Format(dateLayout))), nil
}
