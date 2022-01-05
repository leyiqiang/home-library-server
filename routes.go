package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/leyiqiang/home-library-server/controllers"
)

func Routers(c *controllers.Controller) *chi.Mux {
	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowCredentials: true,
	}))
	router.Route("/book", func(r chi.Router) {
		r.Get("/{bookID}", c.GetOneBook)
		r.Get("/all", c.GetAllBooks)
		r.Post("/", c.AddBook)
	})

	router.Route("/user", func(r chi.Router) {
		r.Post("/register", c.Register)
	})

	return router
}
