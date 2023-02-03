package repository

import (
	"context"
	"database/sql"
	"sbank/internal/controller/dto"
	"sbank/internal/models"
)

type Account interface {
	CreateAccount(ctx context.Context, arg dto.CreateAccountDTO) (models.Account, error)
	GetAccount(ctx context.Context, reqID int64) (models.Account, error)
	UpdateAccount(ctx context.Context, arg dto.UpdateAccountDTO) (models.Account, error)
	DeleteAccount(ctx context.Context, id int64) error
}

type AccountStorage struct {
	db *sql.DB
}

func NewAccountStorage(db *sql.DB) *AccountStorage {
	return &AccountStorage{db: db}
}

func (as *AccountStorage) CreateAccount(ctx context.Context, arg dto.CreateAccountDTO) (models.Account, error) {
	var i models.Account

	query := `INSERT INTO accounts (owner, balance, currency) VALUES ($1, $2, $3) RETURNING id, owner, balance, currency, created_at`

	row := as.db.QueryRowContext(ctx, query, arg.Owner, arg.Balance, arg.Currency)

	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.Balance,
		&i.Currency,
		&i.CreatedAt,
	)

	return i, err
}

func (as *AccountStorage) GetAccount(ctx context.Context, reqID int64) (models.Account, error) {
	var account models.Account

	query := `SELECT id, owner, balance, currency, created_at FROM accounts Where id = $1 LIMIT 1`

	row := as.db.QueryRowContext(ctx, query, reqID)

	err := row.Scan(
		&account.ID,
		&account.Owner,
		&account.Balance,
		&account.Currency,
		&account.CreatedAt,
	)

	return account, err
}

func (as *AccountStorage) GetAccountForUpdate(ctx context.Context, id int64) (models.Account, error) {
	var account models.Account

	query := `SELECT id, owner, balance, currency, created_at FROM accounts Where id = $1 LIMIT 1 FOR UPDATE`

	row := as.db.QueryRowContext(ctx, query, id)

	err := row.Scan(
		&account.ID,
		&account.Owner,
		&account.Balance,
		&account.Currency,
		&account.CreatedAt,
	)

	return account, err
}

func (as *AccountStorage) UpdateAccount(ctx context.Context, arg dto.UpdateAccountDTO) (models.Account, error) {
	var account models.Account

	query := `UPDATE accounts SET balance = $2 WHERE id = $1 RETURNING id, owner, balance, currency, created_at`

	row := as.db.QueryRowContext(ctx, query, arg.ID, arg.Balance)

	err := row.Scan(
		&account.ID,
		&account.Owner,
		&account.Balance,
		&account.Currency,
		&account.CreatedAt,
	)

	return account, err
}

func (as *AccountStorage) DeleteAccount(ctx context.Context, id int64) error {
	query := `DELETE FROM accounts WHERE id = $1`

	_, err := as.db.ExecContext(ctx, query, id)

	return err
}
