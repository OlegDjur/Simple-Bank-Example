package repository

import (
	"context"
	"database/sql"
	"sbank/internal/controller/dto"
	"sbank/internal/models"
)

type transfer interface {
	CraeteTransfer(ctx context.Context, arg dto.TransferDTO) (models.Transfer, error)
}

type TransferStorage struct {
	db *sql.DB
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
