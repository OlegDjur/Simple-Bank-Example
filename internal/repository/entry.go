package repository

import (
	"context"
	"database/sql"
	"sbank/internal/controller/dto"
	"sbank/internal/models"
)

type Entry interface {
	CreateEntry(ctx context.Context, arg dto.CreateEntryDTO) (models.Entry, error)
	GetEntry(ctx context.Context, id int64) (models.Entry, error)
	ListEntries(ctx context.Context, arg dto.ListEntriesDTO) ([]models.Entry, error)
}

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

func (es *EntryStorage) ListEntries(ctx context.Context, arg dto.ListEntriesDTO) ([]models.Entry, error) {
	var listEntries []models.Entry

	query := `SELECT id, account_id, amount, created_at FROM entries WHERE account_id = $1 ORDER BY id LIMIT $2, OFFSET $3`

	rows, err := es.db.QueryContext(ctx, query, arg.AccountID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var i models.Entry

		if err := rows.Scan(
			&i.ID,
			&i.AccountID,
			&i.Amount,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		listEntries = append(listEntries, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return listEntries, nil
}
