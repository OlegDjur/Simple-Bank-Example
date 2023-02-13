package controller

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"sbank/internal/controller/dto"
	"sbank/internal/token"
	"sbank/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

func (h *Handler) CreateAccount(ctx *gin.Context) {
	var req dto.CreateAccountRequestDTO

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	arg := dto.CreateAccountParamsDTO{
		Owner:    authPayload.Username,
		Currency: req.Currency,
		// Balance:  1000,
	}

	account, err := h.service.CreateAccount(ctx, arg)
	if err != nil {
		fmt.Println("error", err)
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "foreign_key_violation", "unique_violation":
				ctx.JSON(http.StatusForbidden, utils.ErrorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}

func (h *Handler) GetAccount(ctx *gin.Context) {
	var req dto.GetAccountDTO

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	account, err := h.service.GetAccount(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, utils.ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusNotFound, utils.ErrorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	if account.Owner != authPayload.Username {
		err := errors.New("account doesn't belong to authenticated user")
		ctx.JSON(http.StatusUnauthorized, utils.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}

func (h *Handler) GetListAccounts(ctx *gin.Context) {
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	arg := dto.ListAccountsDTO{
		Owner: authPayload.Username,
	}

	accounts, err := h.service.GetListAccounts(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, accounts)
}
