package repositories

import "github.com/leyiqiang/home-library-server/models"

type BooksRepository interface {
	GetBookByID(id int) (*models.Book, error)
	GetAllBooks() ([]*models.Book, error)
	AddBook(book models.Book) error
}

type Repository interface {
	BooksRepository
}
