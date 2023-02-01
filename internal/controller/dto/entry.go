package dto

type CreateEntryDTO struct {
	AccountID int64 `json:"account_id"`
	Amount    int64 `json:"amount"`
}
