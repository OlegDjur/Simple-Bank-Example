package controller

import (
	"database/sql"
	"errors"
	"net/http"
	"sbank/internal/controller/dto"
	"sbank/internal/service"
	"sbank/internal/token"
	"sbank/internal/utils"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createTransfer(ctx *gin.Context) {
	var req dto.CreateTransferDTO

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	tempAuthPayload := ctx.MustGet(authorizationPayloadKey)
	authPayload := tempAuthPayload.(*token.Payload)

	arg := dto.CreateTransferDTO{
		FromAccountID: req.FromAccountID,
		ToAccountID:   req.ToAccountID,
		Amount:        req.Amount,
		Currency:      req.Currency,
		AuthUsername:  authPayload.Username,
	}

	result, err := h.service.CreateTransfer(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, utils.ErrorResponse(err))
			return
		}

		if errors.Is(err, service.ErrCurrency) {
			ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		}

		if errors.Is(err, service.ErrAuthUser) {
			ctx.JSON(http.StatusUnauthorized, utils.ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, result)
}
