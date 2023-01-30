package repository

import (
	"context"
	"database/sql"
	"fmt"
	"sbank/internal/controller/dto"
	"sbank/internal/models"
)

type Account interface {
	CreateAccount(ctx context.Context, arg dto.CreateAccountDTO) (models.Account, error)
}

type AccountStorage struct {
	db *sql.DB
}

func NewAccountStorage(db *sql.DB) *AccountStorage {
	return &AccountStorage{db: db}
}

func (as *AccountStorage) CreateAccount(ctx context.Context, arg dto.CreateAccountDTO) (models.Account, error) {
	var i models.Account

	query := `INSERT INTO accounts (
	  owner,
	  currency
	) VALUES (
	  $1, $2
	) RETURNING id, owner, balance, currency, created_at
	`

	row := as.db.QueryRowContext(ctx, query, arg.Owner, arg.Currency)

	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.Balance,
		&i.Currency,
		&i.CreatedAt,
	)
	fmt.Println(err)
	return i, err
}
