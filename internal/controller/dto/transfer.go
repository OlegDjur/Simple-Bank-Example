package dto

import "sbank/internal/models"

type CreateTransferDTO struct {
	FromAccountID int64  `json:"from_account_id"`
	ToAccountID   int64  `json:"to_account_id"`
	Amount        int64  `json:"amount"`
	Currency      string `json:"currency" binding:"required,oneof=USD EUR CAD"`
}

type TransferTxParamsDTO struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

type TransferTxResult struct {
	Transfer    models.Transfer `json:"transfer"`
	FromAccount models.Account  `json:"from_account"`
	ToAccount   models.Account  `json:"to_account"`
	FromEntry   models.Entry    `json:"from_entry"`
	ToEntry     models.Entry    `json:"to_entry"`
}

type ListTransfersDTO struct {
	FromAccountID int64 `json:"from_acount_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Limit         int32 `json:"limit"`
	Offset        int32 `json:"offset"`
}
