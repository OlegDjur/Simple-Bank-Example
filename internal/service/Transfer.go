package service

import (
	"database/sql"
	"sbank/internal/controller/dto"
	"sbank/internal/repository"

	"github.com/gin-gonic/gin"
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
	valid1 := ts.validAccount(ctx, arg.ToAccountID, arg.Currency)
	if !valid1 {
		// return
	}

	valid2 := ts.validAccount(ctx, arg.ToAccountID, arg.Currency)
	if !valid2 {
		// return
	}

	return ts.repo.CreateTransferTx(ctx, arg)
}

func (ts *TransferService) validAccount(ctx *gin.Context, accountID int64, currency string) bool {
	account, err := ts.repo.GetAccount(ctx, accountID)
	if err != nil {
		if err == sql.ErrNoRows {
			// ctx.JSON(http.StatusNotFound, utils.ErrorResponse(err))
			return false
		}
		// ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return false
	}

	if account.Currency != currency {
		// err := fmt.Errorf("account [%d] currency mismatch: %s vs %s", account.ID, account.Currency, currency)
		// ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return false
	}

	return true
}
