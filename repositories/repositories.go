package repositories

import "github.com/leyiqiang/home-library-server/models"

type BooksRepository interface {
	GetBookByID(id string) (*models.Book, error)
	GetAllBooks() ([]*models.Book, error)
	AddBook(book models.Book) (string, error)
	DeleteBookByID(id string) error
	UpdateBookByID(id string, newBookInfo models.Book) (*models.Book, error)
}

type SchedulesRepository interface {
	GetScheduleByID(id string) (*models.Schedule, error)
	DeleteScheduleByID(id string) error
	GetAllSchedules() ([]*models.Schedule, error)
	AddSchedule(schedule models.Schedule) (string, error)
	UpdateScheduleByID(id string, newScheduleInfo models.Schedule) (*models.Schedule, error)
}

type ReservationsRepository interface {
}

type UsersRepository interface {
	Register(user models.User) error
}

type Repository interface {
	BooksRepository
	SchedulesRepository
	ReservationsRepository
	UsersRepository
}
