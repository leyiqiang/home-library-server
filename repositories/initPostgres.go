package repositories

import (
	"database/sql"
	"github.com/leyiqiang/home-library-server/interfaces"
)

type postgresRepo struct {
	DB *sql.DB
}

func NewPostgresRepo(db *sql.DB) interfaces.Repository {
	return &postgresRepo{
		DB: db,
	}
}
