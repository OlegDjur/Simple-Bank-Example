package controller

import (
	"database/sql"
	"fmt"
	"net/http"
	"sbank/internal/controller/dto"
	"sbank/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

func (h *Handler) CreateAccount(ctx *gin.Context) {
	var req dto.CreateAccountDTO

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
	}

	account, err := h.service.CreateAccount(ctx, req)
	if err != nil {
		fmt.Println("error", err)
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "foreign_key_violation", "unique_violation":
				ctx.JSON(http.StatusForbidden, utils.ErrorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse)
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

	ctx.JSON(http.StatusOK, account)
}
