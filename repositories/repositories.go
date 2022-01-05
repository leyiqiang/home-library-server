package repositories

import "github.com/leyiqiang/home-library-server/models"

type BooksRepository interface {
	GetBookByID(id string) (*models.Book, error)
	GetAllBooks() ([]*models.Book, error)
	AddBook(book models.Book) error
	DeleteBookByID(id string) error
	UpdateBookByID(id string, newBookInfo models.Book) (*models.Book, error)
}

type UsersRepository interface {
	Register(user models.User) error
}

type Repository interface {
	BooksRepository
	UsersRepository
}
