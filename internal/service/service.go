package service

import "sbank/internal/repository"

type Service struct {
	Account
	User
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Account: NewAccountService(repo.Account),
		User:    NewUserService(repo.User),
	}
}
