package controllers

import (
	"encoding/json"
	"fmt"
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

func (c *Controller) DeleteBook(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "bookID")

	err := c.Repo.DeleteBookByID(id)
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}
	err = utils.WriteJSON(w, http.StatusOK, "success", "message")
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
	if book.ImportedDate.IsZero() {
		fmt.Println("nil!")
	}
	err = c.Repo.AddBook(book)

	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	err = utils.WriteJSON(w, http.StatusCreated, "created", "message")

}

func (c *Controller) UpdateBook(w http.ResponseWriter, r *http.Request) {
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
	var updatedBook *models.Book
	updatedBook, err = c.Repo.UpdateBookByID(book.ID.String(), book)

	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	err = utils.WriteJSON(w, http.StatusCreated, updatedBook, "book")
}
