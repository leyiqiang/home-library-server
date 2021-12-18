package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/leyiqiang/home-library-server/controllers"
)

func Routers(c *controllers.Controller) *chi.Mux {
	router := chi.NewRouter()
	router.Get("/v1/movie/{movieID}", c.GetOneMovie)
	router.Get("/v1/movies", c.GetAllMovies)
	return router
}
