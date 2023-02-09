package service

import (
	"database/sql"
	"errors"
	"sbank/internal/controller/dto"
	"sbank/internal/models"
	"sbank/internal/repository"

	"github.com/gin-gonic/gin"
)

var (
	ErrCurrency = errors.New("account currency mismatch")
	ErrAuthUser = errors.New("from account doesn't belog to the authenticated user")
)

type Transfer interface {
	CreateTransfer(ctx *gin.Context, arg dto.CreateTransferDTO) (dto.TransferTxResult, error)
}

type TransferService struct {
	repo *repository.Repository
}

func NewTransferService(repo *repository.Repository) *TransferService {
	return &TransferService{repo: repo}
}

func (ts *TransferService) CreateTransfer(ctx *gin.Context, arg dto.CreateTransferDTO) (dto.TransferTxResult, error) {
	fromAccount, err := ts.validAccount(ctx, arg.FromAccountID, arg.Currency)
	if err != nil {
		return dto.TransferTxResult{}, err
	}

	if fromAccount.Owner != arg.AuthUsername {
		return dto.TransferTxResult{}, ErrAuthUser
	}

	_, err = ts.validAccount(ctx, arg.ToAccountID, arg.Currency)
	if err != nil {
		return dto.TransferTxResult{}, err
	}

	return ts.repo.CreateTransferTx(ctx, arg)
}

func (ts *TransferService) validAccount(ctx *gin.Context, accountID int64, currency string) (models.Account, error) {
	account, err := ts.repo.GetAccount(ctx, accountID)
	if err != nil {
		if err == sql.ErrNoRows {
			return account, err
		}

		return account, err
	}

	if account.Currency != currency {
		// err := fmt.Errorf("account [%d] currency mismatch: %s vs %s", account.ID, account.Currency, currency)

		return account, ErrCurrency
	}

	return account, nil
}
