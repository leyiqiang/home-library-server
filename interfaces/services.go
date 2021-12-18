package interfaces

import "github.com/leyiqiang/home-library-server/models"

type BooksService interface {
	GetOneBook(id int) (*models.Book, error)
	GetAllBooks() ([]*models.Book, error)
}

type Service interface {
	BooksService
}
