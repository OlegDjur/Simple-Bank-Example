package controller

import (
	"net/http"
	"sbank/internal/controller/dto"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

func (h *Handler) CreateAccount(ctx *gin.Context) {
	var req dto.CreateAccountDTO

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	account, err := h.service.CreateAccount(ctx, req)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "foreign_key_violation", "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse)
		return
	}

	ctx.JSON(http.StatusOK, account)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
