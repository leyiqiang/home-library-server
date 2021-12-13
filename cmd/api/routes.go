package main

import "github.com/go-chi/chi/v5"

func (app *application) routes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/status", app.statusHandler)
	return router
}
