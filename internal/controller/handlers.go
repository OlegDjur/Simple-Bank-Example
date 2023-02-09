package controller

import (
	"sbank/config"
	"sbank/internal/service"
	"sbank/internal/token"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service    *service.Service
	config     *config.Config
	tokenMaker token.Maker
}

func NewHandler(service *service.Service, config *config.Config) *Handler {
	tokenMaker := token.NewJWTMaker(config.TokenSymmetricKey)

	return &Handler{
		service:    service,
		config:     config,
		tokenMaker: tokenMaker,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	router.POST("/user", h.createUser)
	router.POST("/user/login", h.loginUser)

	// router.POST("/accounts", h.CreateAccount)

	authRoutes := router.Group("/").Use(authMiddleware(h.tokenMaker))
	authRoutes.POST("/accounts", h.CreateAccount)
	authRoutes.POST("/transfers", h.createTransfer)

	return router
}
