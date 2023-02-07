package service

import (
	"context"
	"fmt"
	"sbank/internal/controller/dto"
	"sbank/internal/models"
	"sbank/internal/repository"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type User interface {
	CreateUser(ctx *gin.Context, arg dto.CreateUserRequestDTO) (models.User, error)
	GenerateToken(ctx context.Context, req dto.LoginUserRequestDTO, tokenDuration time.Duration) (models.User, string, error)
	// GetUser(ctx *gin.Context, req dto.LoginUserRequestDTO) (models.User, error)
}

type UserService struct {
	repo  repository.User
	maker Maker
}

func NewUserService(repo repository.User, secretKey string) *UserService {
	return &UserService{
		repo:  repo,
		maker: NewJWTMaker(secretKey),
	}
}

func (s *UserService) CreateUser(ctx *gin.Context, arg dto.CreateUserRequestDTO) (models.User, error) {
	var err error

	arg.Password, err = HashPassword(arg.Password)
	if err != nil {
		// ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return models.User{}, err
	}

	return s.repo.CreateUser(ctx, arg)
}

func (s *UserService) GenerateToken(ctx context.Context, req dto.LoginUserRequestDTO, tokenDuration time.Duration) (models.User, string, error) {
	user, err := s.repo.GetUser(ctx, req.Username)
	if err != nil {
		return models.User{}, "", err
	}

	err = CheckPassword(req.Password, user.HashedPassword)
	if err != nil {
		// ctx.JSON(http.StatusUnauthorized, utils.ErrorResponse(err))
		return models.User{}, "", err
	}

	accessToken, err := s.maker.CreateToken(user.Username, tokenDuration)
	if err != nil {
		// ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return models.User{}, "", err
	}

	return user, accessToken, nil
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}

	return string(hashedPassword), nil
}

func CheckPassword(password string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
