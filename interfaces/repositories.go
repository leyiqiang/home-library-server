package interfaces

import "github.com/leyiqiang/home-library-server/models"

type MoviesRepository interface {
	GetMovieByID(id int) (*models.Movie, error)
	GetAllMovies() ([]*models.Movie, error)
}

type Repository interface {
	MoviesRepository
}
