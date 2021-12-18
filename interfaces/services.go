package interfaces

import "github.com/leyiqiang/home-library-server/models"

type MoviesService interface {
	GetOneMovie(id int) (*models.Movie, error)
	GetAllMovies() ([]*models.Movie, error)
}

type Service interface {
	MoviesService
}
