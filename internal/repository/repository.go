package repository

import (
	"database/sql"
)

type Repository struct {
	Account
	User
	Transfer
	Entry
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Account: NewAccountStorage(db),
		User:    NewUserStorage(db),
	}
}
