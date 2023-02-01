package repository

import (
	"context"
	"database/sql"
	"sbank/internal/controller/dto"
	"sbank/internal/models"
)

type Entry interface{}

type EntryStorage struct {
	db *sql.DB
}

func NewEntryStorage(db *sql.DB) *EntryStorage {
	return &EntryStorage{db: db}
}

func (es *EntryStorage) CreateEntry(ctx context.Context, arg dto.CreateEntryDTO) (models.Entry, error) {
	var entry models.Entry

	query := `INSERT INTO entries (accoint_id, amount) VALUES ($1, $2) RETURNING id, account_id, amount, created_at`

	row := es.db.QueryRowContext(ctx, query, arg.AccountID, arg.Amount)

	err := row.Scan(
		&entry.ID,
		&entry.AccountID,
		&entry.Amount,
		&entry.CreatedAt,
	)

	return entry, err
}

func (es *EntryStorage) GetEntry(ctx context.Context, id int64) (models.Entry, error) {
	var entry models.Entry

	query := `SELECT id, account_id, amount, created_at FROM entries WHERE id = $1 LIMIT 1`

	row := es.db.QueryRowContext(ctx, query, id)

	err := row.Scan(
		&entry.ID,
		&entry.AccountID,
		&entry.Amount,
		&entry.CreatedAt,
	)

	return entry, err
}