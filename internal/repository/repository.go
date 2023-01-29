package repository

import (
	"database/sql"
)

type Repository struct {
	Account
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Account: NewAccountStorage(db),
	}
}
