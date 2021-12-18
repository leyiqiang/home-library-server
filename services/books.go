package services

import (
	"github.com/leyiqiang/home-library-server/models"
	"log"
)

var logger *log.Logger

func (s *service) GetOneBook(id int) (*models.Book, error) {
	movie, err := s.repo.GetBookByID(id)
	return movie, err
}

func (s *service) GetAllBooks() ([]*models.Book, error) {
	movies, err := s.repo.GetAllBooks()
	return movies, err
}
