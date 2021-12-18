package models

import (
	"time"
)

type Book struct {
	ID          int            `json:"id"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Year        int            `json:"year"`
	ReleaseDate time.Time      `json:"releaseDate"`
	Rating      int            `json:"rating"`
	CreatedAt   time.Time      `json:"-"`
	Genre       map[int]string `json:"genres"`
}
