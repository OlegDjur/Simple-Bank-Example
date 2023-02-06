package service

import (
	"context"
	"errors"
	"sbank/internal/controller/dto"
	"sbank/internal/models"
	"sbank/internal/repository"

	"github.com/gin-gonic/gin"
)

var (
	ErrInvalidcurrencyType = errors.New("invalid currency type")
	ErrInvalidOwner        = errors.New("invalid user")
	ErrInvalidBalance      = errors.New("invalid balance")
)

type Account interface {
	CreateAccount(ctx context.Context, arg dto.CreateAccountDTO) (models.Account, error)
	GetAccount(ctx *gin.Context, reqID int64) (models.Account, error)
}

type AccountService struct {
	repo repository.Account
}

func NewAccountService(repo repository.Account) *AccountService {
	return &AccountService{repo: repo}
}

func (as *AccountService) CreateAccount(ctx context.Context, arg dto.CreateAccountDTO) (models.Account, error) {
	if err := validAccount(arg); err != nil {
		return models.Account{}, err
	}

	return as.repo.CreateAccount(ctx, arg)
}

func (as *AccountService) GetAccount(ctx *gin.Context, reqID int64) (models.Account, error) {
	return as.repo.GetAccount(ctx, reqID)
}

func validAccount(arg dto.CreateAccountDTO) error {
	if err := validateAccountOwner(arg.Owner); err != nil {
		return ErrInvalidOwner
	}

	if arg.Balance > 0 {
		return ErrInvalidBalance
	}

	if err := validateAccountCurrency(arg.Currency); err != nil {
		return ErrInvalidcurrencyType
	}

	return nil
}

func validateAccountOwner(owner string) error {
	if owner == "" {
		return ErrInvalidOwner
	}

	if len(owner) > 15 || len(owner) < 3 {
		return ErrInvalidOwner
	}

	for _, v := range owner {
		if v < 33 || v > 126 {
			return ErrInvalidOwner
		}
	}

	return nil
}

func validateAccountCurrency(currency string) error {
	if currency == "" {
		return ErrInvalidcurrencyType
	}

	currencyType := map[string]bool{
		"usd": true,
		"eur": true,
		"kzt": true,
	}

	if _, ok := currencyType[currency]; !ok {
		return ErrInvalidcurrencyType
	}

	return nil
}
