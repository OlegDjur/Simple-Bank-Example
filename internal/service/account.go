package service

import (
	"context"
	"sbank/internal/controller/dto"
	"sbank/internal/models"
	"sbank/internal/repository"

	"github.com/gin-gonic/gin"
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
	return as.repo.CreateAccount(ctx, arg)
}

func (as *AccountService) GetAccount(ctx *gin.Context, reqID int64) (models.Account, error) {
	return as.repo.GetAccount(ctx, reqID)
}

func (as *AccountService) validAccount(arg dto.CreateAccountDTO) bool {
	// account, err := as.repo.GetAccount(arg.)

	return true
}

func validateAccountCurrency(currency string) bool {
	return true
}
