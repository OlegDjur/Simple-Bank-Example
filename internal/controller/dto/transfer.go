package dto

type TransferDTO struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

type TransferTxDTO struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

type ListTransfersDTO struct {
	FromAccountID int64 `json:"from_acount_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Limit         int32 `json:"limit"`
	Offset        int32 `json:"offset"`
}
