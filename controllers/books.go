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

func (c *Controller) GetOneBook(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "bookID"))

	if err != nil {
		logger.Print(errors.New("invalid id parameter"))
		utils.ErrorJSON(w, err)
		return
	}

	book, err := c.Service.GetOneBook(id)

	err = utils.WriteJSON(w, http.StatusOK, book, "book")
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

}

func (c *Controller) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	books, err := c.Service.GetAllBooks()
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	err = utils.WriteJSON(w, http.StatusOK, books, "books")
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}
}
