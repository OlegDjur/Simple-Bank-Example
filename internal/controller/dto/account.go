package dto

type CreateAccountDTO struct {
	Owner    string `json:"owner"`
	Balance  int64  `json:"balance"`
	Currency string `json:"currency"`
}

type GetAccountDTO struct {
	ID int64
}
