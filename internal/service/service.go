package service

import "sbank/internal/repository"

type Service struct {
	Account
	User
	Maker
}

func NewService(repo *repository.Repository, secretKey string) *Service {
	return &Service{
		Account: NewAccountService(repo.Account),
		User:    NewUserService(repo.User),
		Maker:   NewJWTMaker(secretKey),
	}
}
