package service

import (
	"context"
	"sbank/internal/controller/dto"
	"sbank/internal/repository"
)

type Account interface {
	CreateAccount(ctx context.Context, arg dto.CreateAccountParams) (Account, error)
}

type AccountService struct {
	repo repository.Account
}

func NewAccountService(repo repository.Account) *AccountService {
	return &AccountService{repo: repo}
}

func (as *AccountService) CreateAccount(ctx context.Context, arg dto.CreateAccountParams) (Account, error) {
	return nil, nil
}
