package controllers

import (
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/leyiqiang/home-library-server/models"
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

	book, err := c.Repo.GetBookByID(id)

	err = utils.WriteJSON(w, http.StatusOK, book, "book")
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

}

func (c *Controller) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	books, err := c.Repo.GetAllBooks()
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

func (c *Controller) AddBook(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var book models.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		utils.ErrorJSON(w, err)
		return
	}
	err := c.Repo.AddBook(book)

	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	err = utils.WriteJSON(w, http.StatusOK, "", "books")

}
