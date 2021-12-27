package controllers

import (
	"github.com/go-playground/validator/v10"
	"github.com/leyiqiang/home-library-server/repositories"
	"net/http"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

type Controller struct {
	Repo repositories.Repository
	BooksController
	UsersController
}

type BooksController interface {
	GetOneBook(w http.ResponseWriter, r *http.Request)
	GetAllBooks(w http.ResponseWriter, r *http.Request)
	AddBook(w http.ResponseWriter, r *http.Request)
}

type UsersController interface {
	Register(w http.ResponseWriter, r *http.Request)
}
