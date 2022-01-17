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
	SchedulesController
	ReservationsController
}

type BooksController interface {
	GetOneBook(w http.ResponseWriter, r *http.Request)
	GetAllBooks(w http.ResponseWriter, r *http.Request)
	AddBook(w http.ResponseWriter, r *http.Request)
	DeleteBook(w http.ResponseWriter, r *http.Request)
	UpdateBook(w http.ResponseWriter, r *http.Request)
}

type SchedulesController interface {
	GetOneSchedule(w http.ResponseWriter, r *http.Request)
	GetAllSchedules(w http.ResponseWriter, r *http.Request)
	AddSchedule(w http.ResponseWriter, r *http.Request)
	DeleteSchedule(w http.ResponseWriter, r *http.Request)
	UpdateSchedule(w http.ResponseWriter, r *http.Request)
}

type ReservationsController interface {
}

type UsersController interface {
	Register(w http.ResponseWriter, r *http.Request)
	SignIn(w http.ResponseWriter, r *http.Request)
}
