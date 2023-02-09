package service

import "sbank/internal/repository"

type Service struct {
	Account
	User
	Transfer
}

func NewService(repo *repository.Repository, secretKey string) *Service {
	return &Service{
		Account:  NewAccountService(repo.Account),
		User:     NewUserService(repo.User, secretKey),
		Transfer: NewTransferService(repo),
	}
}
