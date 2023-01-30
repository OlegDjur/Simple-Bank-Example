package service

import (
	"context"
	"sbank/internal/controller/dto"
	"sbank/internal/models"
	"sbank/internal/repository"
)

type Account interface {
	CreateAccount(ctx context.Context, arg dto.CreateAccountDTO) (models.Account, error)
}

type AccountService struct {
	repo repository.Account
}

func NewAccountService(repo repository.Account) *AccountService {
	return &AccountService{repo: repo}
}

func (as *AccountService) CreateAccount(ctx context.Context, arg dto.CreateAccountDTO) (models.Account, error) {
	return as.repo.CreateAccount(ctx, arg)
}

func validateAccount(arg dto.CreateAccountDTO) {
}
