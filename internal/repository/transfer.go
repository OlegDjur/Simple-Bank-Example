package repository

import (
	"context"
	"database/sql"
	"sbank/internal/controller/dto"
	"sbank/internal/models"
)

type Transfer interface {
	CraeteTransfer(ctx context.Context, arg dto.TransferDTO) (models.Transfer, error)
	GetTransfer(ctx context.Context, id int64) (models.Transfer, error)
}

type TransferStorage struct {
	db *sql.DB
}

func NewTransferStorage(db *sql.DB) *TransferStorage {
	return &TransferStorage{db: db}
}

func (ts *TransferStorage) CraeteTransfer(ctx context.Context, arg dto.TransferDTO) (models.Transfer, error) {
	var transfer models.Transfer

	query := `INSERT INTO transfers (
		from_account_id, 
		to_account_id, 
		amount, 
		created_at
		) VALUES (
		$1, $2, $3
		) RETURNING id, from_account_id, to_account_id, amount, created_at
		`

	row := ts.db.QueryRowContext(ctx, query, arg.FromAccountID, arg.ToAccountID, arg.Amount)

	err := row.Scan(
		&transfer.ID,
		&transfer.FromAccountID,
		&transfer.ToAccountID,
		&transfer.Amount,
		&transfer.CreatedAt,
	)

	return transfer, err
}

func (ts *TransferStorage) GetTransfer(ctx context.Context, id int64) (models.Transfer, error) {
	var transfer models.Transfer

	query := `SELECT id, from_account_id, to_account_id, amount, created_at FROM transfers`

	row := ts.db.QueryRowContext(ctx, query, id)

	err := row.Scan(
		&transfer.ID,
		&transfer.FromAccountID,
		&transfer.ToAccountID,
		&transfer.Amount,
		&transfer.CreatedAt,
	)

	return transfer, err
}

func (ts *TransferStorage) ListTransfers(ctx context.Context, arg dto.ListTransfersDTO) ([]models.Transfer, error) {
	query := `SELECT id, from_account_id, to_account_id, created_at FROM transfers 
		WHERE 
			from_account_id = $1 OR 
			to_account_id = 2,
		ORDER BY id
		LIMIT $3
		OFFSET $4
	`

	rows, err := ts.db.QueryContext(ctx, query, arg.FromAccountID, arg.ToAccountID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	items := []models.Transfer{}

	for rows.Next() {
		var i models.Transfer
		if err := rows.Scan(
			&i.ID,
			&i.FromAccountID,
			&i.ToAccountID,
			&i.Amount,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}
