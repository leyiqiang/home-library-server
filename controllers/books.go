package controllers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/leyiqiang/home-library-server/models"
	"github.com/leyiqiang/home-library-server/utils"
	"net/http"
)

func (c *Controller) GetOneBook(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "bookID")

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
}

func (c *Controller) AddBook(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var book models.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	validate = validator.New()
	err = validate.Struct(book)
	if err != nil {
		utils.ErrorJSON(w, err)
		return

	}
	err = c.Repo.AddBook(book)

	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	err = utils.WriteJSON(w, http.StatusCreated, "", "book")

}
