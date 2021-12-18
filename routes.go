package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/leyiqiang/home-library-server/controllers"
)

func Routers(c *controllers.Controller) *chi.Mux {
	router := chi.NewRouter()
	router.Get("/book/{bookID}", c.GetOneMovie)
	router.Get("/books", c.GetAllMovies)
	return router
}
