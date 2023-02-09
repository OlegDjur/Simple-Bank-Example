package dto

type CreateAccountRequestDTO struct {
	Currency string `json:"currency"`
}

type CreateAccountParamsDTO struct {
	Owner    string `json:"owner"`
	Balance  int64  `json:"balance"`
	Currency string `json:"currency"`
}

type GetAccountDTO struct {
	ID int64 `json:"id"`
}

type UpdateAccountDTO struct {
	ID      int64 `json:"id"`
	Balance int64 `json:"balance"`
}

type AddAccountBalanceDTO struct {
	ID     int64 `json:"id"`
	Amount int64 `json:"balance"`
}
