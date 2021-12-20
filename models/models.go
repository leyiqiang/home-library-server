package models

import (
	"time"
)

type Book struct {
	Title       string         `json:"title" validate:"required""`
	Description string         `json:"description"`
	Year        int            `json:"year"`
	ReleaseDate time.Time      `json:"releaseDate"`
	Rating      int            `json:"rating"`
	CreatedAt   time.Time      `json:"-"`
	Genre       map[int]string `json:"genres"`
	ISBN        string         `json:"isbn"`
}
