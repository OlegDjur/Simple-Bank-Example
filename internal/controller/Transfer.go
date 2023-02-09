package controller

import (
	"net/http"
	"sbank/internal/controller/dto"
	"sbank/internal/utils"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createTransfer(ctx *gin.Context) {
	var req dto.CreateTransferDTO

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	arg := dto.CreateTransferDTO{
		FromAccountID: req.FromAccountID,
		ToAccountID:   req.ToAccountID,
		Amount:        req.Amount,
		Currency:      req.Currency,
	}

	result, err := h.service.CreateTransfer(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, result)
}
