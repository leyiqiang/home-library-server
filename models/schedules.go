package models

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strings"
	"time"
)

type Schedule struct {
	ID    primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Title string             `json:"title" bson:"title" validate:"required""`
	Start DateTime           `json:"start" json:"start" validate:"required"`
	End   DateTime           `json:"end" json:"end" validate:"required"`
}

type DateTime struct {
	time.Time
}

var dateTimeLayout = "2006-01-02|15:04:05"

func (d *DateTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	if s == "null" || s == "" {
		d.Time = time.Now()
		return nil
	}
	var err error
	d.Time, err = time.Parse(dateTimeLayout, s)
	return err
}

func (d *DateTime) MarshalJSON() ([]byte, error) {
	if d.Time.UnixNano() == (time.Time{}).UnixNano() {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%s\"", d.Format(dateTimeLayout))), nil
}
