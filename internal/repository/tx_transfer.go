package repository

import (
	"context"
	"database/sql"
	"fmt"
	"sbank/internal/controller/dto"
	"sbank/internal/models"
)

type TransferTx interface {
	TtransferTx(ctx context.Context, arg dto.TransferTxDTO) (TransferTxResult, error)
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

type TransferTxResult struct {
	Transfer    models.Transfer `json:"transfer"`
	FromAccount models.Account  `json:"from_account"`
	ToAccount   models.Account  `json:"to_account"`
	FromEntry   models.Entry    `json:"from_entry"`
	ToEntry     models.Entry    `json:"to_entry"`
}

var txKey = struct{}{}

func (txs *TransferTxStorage) TtransferTx(ctx context.Context, arg dto.TransferTxDTO) (TransferTxResult, error) {
	var result TransferTxResult

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
		return TransferTxResult{}, err
	}

	result.FromEntry, err = txs.Entry.CreateEntry(ctx, dto.CreateEntryDTO{
		AccountID: arg.FromAccountID,
		Amount:    -arg.Amount,
	})
	if err != nil {
		tx.Rollback()
		return TransferTxResult{}, err
	}

	result.ToEntry, err = txs.Entry.CreateEntry(ctx, dto.CreateEntryDTO{
		AccountID: arg.ToAccountID,
		Amount:    arg.Amount,
	})
	if err != nil {
		tx.Rollback()
		return TransferTxResult{}, err
	}

	// TODO: update accounts balans

	// переводим деньги с account1
	result.FromAccount, err = txs.Account.AddAccountBalance(ctx, dto.AddAccountBalanceDTO{
		ID:     arg.FromAccountID,
		Amount: -arg.Amount,
	})
	if err != nil {
		return TransferTxResult{}, err
	}

	// переводим деньги на account2
	result.ToAccount, err = txs.Account.AddAccountBalance(ctx, dto.AddAccountBalanceDTO{
		ID:     arg.ToAccountID,
		Amount: arg.Amount,
	})
	fmt.Println()
	if err != nil {
		return TransferTxResult{}, err
	}

	tx.Commit()
	return result, err
}
