package dto

type CreateEntryDTO struct {
	AccountID int64 `json:"account_id"`
	Amount    int64 `json:"amount"`
}

type ListEntriesDTO struct {
	AccountID int64 `json:"amount"`
	Limit     int32 `json:"limit"`
	Offset    int32 `json:"offset"`
}
