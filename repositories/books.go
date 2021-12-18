package repositories

import (
	"context"
	"github.com/leyiqiang/home-library-server/models"
	"time"
)

func (r *postgresRepo) GetBookByID(id int) (*models.Book, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select id, title, description, year, release_date, rating,
			created_at, updated_at from movies where id = $1`

	row := r.DB.QueryRowContext(ctx, query, id)

	var book models.Book

	err := row.Scan(
		&book.ID,
		&book.Title,
		&book.Description,
		&book.Year,
		&book.ReleaseDate,
		&book.Rating,
		&book.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &book, nil
}

func (r *postgresRepo) GetAllBooks() ([]*models.Book, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select id, title, description, year, release_date, rating,
			created_at, updated_at from books order by title`
	rows, err := r.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	var books []*models.Book

	for rows.Next() {
		var book models.Book
		err := rows.Scan(
			&book.ID,
			&book.Title,
			&book.Description,
			&book.Year,
			&book.ReleaseDate,
			&book.Rating,
			&book.CreatedAt,
		)

		if err != nil {
			return nil, err
		}
		books = append(books, &book)
	}
	return books, nil
}
