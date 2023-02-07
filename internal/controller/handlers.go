package controller

import (
	"sbank/config"
	"sbank/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
	config  *config.Config
}

func NewHandler(service *service.Service, config *config.Config) *Handler {
	return &Handler{
		service: service,
		config:  config,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	router.POST("/user", h.createUser)
	router.POST("/user/login", h.loginUser)

	router.POST("/accounts", h.CreateAccount)

	return router
}
