package service

import (
	"sbank/internal/controller/dto"
	"sbank/internal/models"
	"sbank/internal/repository"

	"github.com/gin-gonic/gin"
)

type User interface {
	CreateUser(ctx *gin.Context, arg dto.CreateUserDTO) (models.User, error)
}

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(ctx *gin.Context, arg dto.CreateUserDTO) (models.User, error) {
	return s.repo.CreateUser(ctx, arg)
}
