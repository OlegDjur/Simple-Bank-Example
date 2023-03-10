package controller

import (
	"os"
	"sbank/config"
	"sbank/internal/repository"
	"sbank/internal/service"
	"sbank/internal/token"
	"sbank/internal/utils"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
)

func newTestServer(t *testing.T) (*Handler, *token.JWTMaker) {
	config := &config.Config{
		TokenSymmetricKey:   utils.RandomString(32),
		AccessTokenDuration: time.Minute,
	}

	repo := repository.NewRepository(nil)
	services := service.NewService(repo, config.TokenSymmetricKey)
	tokenMaker := token.NewJWTMaker(config.TokenSymmetricKey)
	handler := NewHandler(services, config)

	return handler, tokenMaker
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)

	os.Exit(m.Run())
}
