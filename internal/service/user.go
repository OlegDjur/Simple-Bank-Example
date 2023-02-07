package service

import (
	"context"
	"sbank/internal/controller/dto"
	"sbank/internal/models"
	"sbank/internal/repository"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type User interface {
	CreateUser(ctx *gin.Context, arg dto.CreateUserRequestDTO) (models.User, error)
	GenerateToken(ctx context.Context, req dto.LoginUserRequestDTO) (string, error)
	// GetUser(ctx *gin.Context, req dto.LoginUserRequestDTO) (models.User, error)
}

type UserService struct {
	repo  repository.User
	maker Maker
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(ctx *gin.Context, arg dto.CreateUserRequestDTO) (models.User, error) {
	return s.repo.CreateUser(ctx, arg)
}

func (s *UserService) GenerateToken(ctx context.Context, req dto.LoginUserRequestDTO) (string, error) {
	user, err := s.repo.GetUser(ctx, req.Username)
	if err != nil {
		return "", err
	}

	err = CheckPassword(req.Password, user.HashedPassword)
	if err != nil {
		// ctx.JSON(http.StatusUnauthorized, utils.ErrorResponse(err))
		return "", err
	}

	return
}

func CheckPassword(password string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
