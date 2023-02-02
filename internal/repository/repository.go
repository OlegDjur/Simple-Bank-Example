package repository

import (
	"database/sql"
)

type Repository struct {
	Account
	User
	Transfer
	Entry
	TransferTx
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Account:    NewAccountStorage(db),
		User:       NewUserStorage(db),
		Transfer:   NewTransferStorage(db),
		Entry:      NewEntryStorage(db),
		TransferTx: NewTransferTxStorage(db),
	}
}
