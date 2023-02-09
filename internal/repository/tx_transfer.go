package repository

import (
	"context"
	"database/sql"
	"sbank/internal/controller/dto"
)

type TransferTx interface {
	CreateTransferTx(ctx context.Context, arg dto.CreateTransferDTO) (dto.TransferTxResult, error)
}

type TransferTxStorage struct {
	db *sql.DB
	Transfer
	Entry
	Account
}

func NewTransferTxStorage(db *sql.DB) *TransferTxStorage {
	return &TransferTxStorage{
		db:       db,
		Transfer: NewTransferStorage(db),
		Entry:    NewEntryStorage(db),
		Account:  NewAccountStorage(db),
	}
}

var txKey = struct{}{}

func (txs *TransferTxStorage) CreateTransferTx(ctx context.Context, arg dto.CreateTransferDTO) (dto.TransferTxResult, error) {
	var result dto.TransferTxResult

	tx, err := txs.db.BeginTx(ctx, nil)
	if err != nil {
		return result, err
	}

	result.Transfer, err = txs.Transfer.CraeteTransfer(ctx, dto.CreateTransferDTO{
		FromAccountID: arg.FromAccountID,
		ToAccountID:   arg.ToAccountID,
		Amount:        arg.Amount,
	})
	if err != nil {
		tx.Rollback()
		return dto.TransferTxResult{}, err
	}

	result.FromEntry, err = txs.Entry.CreateEntry(ctx, dto.CreateEntryDTO{
		AccountID: arg.FromAccountID,
		Amount:    -arg.Amount,
	})
	if err != nil {
		tx.Rollback()
		return dto.TransferTxResult{}, err
	}

	result.ToEntry, err = txs.Entry.CreateEntry(ctx, dto.CreateEntryDTO{
		AccountID: arg.ToAccountID,
		Amount:    arg.Amount,
	})
	if err != nil {
		tx.Rollback()
		return dto.TransferTxResult{}, err
	}

	// TODO: update accounts balans

	if arg.FromAccountID < arg.ToAccountID {
		result.FromAccount, err = txs.AddAccountBalance(ctx, dto.AddAccountBalanceDTO{
			ID:     arg.FromAccountID,
			Amount: -arg.Amount,
		})
		if err != nil {
			tx.Rollback()
			return dto.TransferTxResult{}, err
		}

		result.ToAccount, err = txs.AddAccountBalance(ctx, dto.AddAccountBalanceDTO{
			ID:     arg.ToAccountID,
			Amount: arg.Amount,
		})
		if err != nil {
			tx.Rollback()
			return dto.TransferTxResult{}, err
		}
	} else {
		result.ToAccount, err = txs.AddAccountBalance(ctx, dto.AddAccountBalanceDTO{
			ID:     arg.ToAccountID,
			Amount: arg.Amount,
		})
		if err != nil {
			tx.Rollback()
			return dto.TransferTxResult{}, err
		}

		result.FromAccount, err = txs.AddAccountBalance(ctx, dto.AddAccountBalanceDTO{
			ID:     arg.FromAccountID,
			Amount: -arg.Amount,
		})
		if err != nil {
			tx.Rollback()
			return dto.TransferTxResult{}, err
		}
	}

	tx.Commit()
	return result, err
}
