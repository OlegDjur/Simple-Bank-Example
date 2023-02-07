package controller

import (
	"database/sql"
	"net/http"
	"sbank/internal/controller/dto"
	"sbank/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

func (h *Handler) createUser(ctx *gin.Context) {
	var req dto.CreateUserRequestDTO

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	user, err := h.service.CreateUser(ctx, req)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				ctx.JSON(http.StatusForbidden, utils.ErrorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	rsp := dto.NewUserResponse(user)

	ctx.JSON(http.StatusOK, rsp)
}

func (h *Handler) loginUser(ctx *gin.Context) {
	var req dto.LoginUserRequestDTO

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
	}

	user, accessToken, err := h.service.GenerateToken(ctx, req, h.config.AccessTokenDuration)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, utils.ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	rsp := dto.LoginUserResponseDTO{
		AccessToken: accessToken,
		User:        *dto.NewUserResponse(user),
	}
	ctx.JSON(http.StatusOK, rsp)
}
