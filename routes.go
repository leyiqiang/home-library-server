package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/leyiqiang/home-library-server/controllers"
	"github.com/leyiqiang/home-library-server/middleware"
	"net/http"
)

func Routers(c *controllers.Controller) *chi.Mux {
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowCredentials: true,
	}))
	r.Route("/book", func(r chi.Router) {
		r.Get("/all", c.GetAllBooks)
		r.Get("/{bookID}", c.GetOneBook)

	})

	r.Route("/user", func(r chi.Router) {
		r.Post("/register", c.Register)
	})

	r.Mount("/admin", adminRouter(c))

	return r
}

func adminRouter(c *controllers.Controller) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.AdminOnly)
	r.Route("/book", func(r chi.Router) {
		r.Post("/", c.AddBook)
		r.Put("/{bookID}", c.UpdateBook)
		r.Delete("/{bookID}", c.DeleteBook)
	})
	return r
}
