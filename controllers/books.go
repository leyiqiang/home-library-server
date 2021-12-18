package controllers

import (
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/leyiqiang/home-library-server/utils"
	"log"
	"net/http"
	"strconv"
)

var logger *log.Logger

func (c *Controller) GetOneMovie(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "movieID"))

	if err != nil {
		logger.Print(errors.New("invalid id parameter"))
		utils.ErrorJSON(w, err)
		return
	}

	movie, err := c.Service.GetOneBook(id)

	err = utils.WriteJSON(w, http.StatusOK, movie, "movie")
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

}

func (c *Controller) GetAllMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := c.Service.GetAllBooks()
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	err = utils.WriteJSON(w, http.StatusOK, movies, "movies")
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {

}

func InsertMovie(w http.ResponseWriter, r *http.Request) {

}

func UpdateMovie(w http.ResponseWriter, r *http.Request) {

}

func SearchMovie(w http.ResponseWriter, r *http.Request) {

}
