package services

import (
	"github.com/leyiqiang/home-library-server/models"
	"log"
	"net/http"
)

var logger *log.Logger

func (s *service) GetOneMovie(id int) (*models.Movie, error) {
	movie, err := s.repo.GetMovieByID(id)
	return movie, err
}

func (s *service) GetAllMovies() ([]*models.Movie, error) {
	movies, err := s.repo.GetAllMovies()
	return movies, err
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {

}

func InsertMovie(w http.ResponseWriter, r *http.Request) {

}

func UpdateMovie(w http.ResponseWriter, r *http.Request) {

}

func SearchMovie(w http.ResponseWriter, r *http.Request) {

}
