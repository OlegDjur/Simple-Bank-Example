package repository

import (
	"context"
	"database/sql"
	"sbank/internal/controller/dto"
)

type Account interface {
	CreateAccount(ctx context.Context, arg dto.CreateAccountParams) (Account, error)
}

type AccountStorage struct {
	db *sql.DB
}

func NewAccountStorage(db *sql.DB) *AccountStorage {
	return &AccountStorage{db: db}
}

func (db *AccountStorage) CreateAccount(ctx context.Context, arg dto.CreateAccountParams) (Account, error) {
	var i Account

	return i, nil
}
