package service

import (
	"context"
	"sbank/internal/controller/dto"
	"sbank/internal/models"
	"sbank/internal/repository"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

const (
	signingKey = "sgkj#laksjd#LKJSL89dfg"
	tokenTTl   = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	Username string `json:"username"`
}

type User interface {
	CreateUser(ctx *gin.Context, arg dto.CreateUserRequestDTO) (models.User, error)
	GetUser(ctx *gin.Context, req dto.LoginUserRequestDTO) (models.User, error)
}

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(ctx *gin.Context, arg dto.CreateUserRequestDTO) (models.User, error) {
	return s.repo.CreateUser(ctx, arg)
}

// func (s *UserService) GetUser(ctx *gin.Context, req dto.LoginUserRequestDTO) (models.User, error) {
// 	user, err := s.repo.GetUser(ctx, req.Username)
// 	if err != nil {
// 		return models.User{}, err
// 	}

// 	if err = checkPassword(req.Password, user.HashedPassword); err != nil {
// 		ctx.JSON(http.StatusUnauthorized, utils.ErrorResponse(err))
// 		return models.User{}, err
// 	}

// 	accessToken, err := CreateToken(user.Username)

// 	return user, nil
// }

func (s *UserService) GenerateToken(ctx context.Context, username, password string) (string, error) {
	user, err := s.repo.GetUser(ctx, username)
	if err != nil {
		return "", err
	}

	// err = util.CheckPassword(req.Password, user.HashedPassword)
	// if err != nil {
	// 	ctx.JSON(http.StatusUnauthorized, errorResponse(err))
	// 	return
	// }

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTl).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Username,
	})

	return token.SignedString([]byte(signingKey))
}

// func generatePassword
