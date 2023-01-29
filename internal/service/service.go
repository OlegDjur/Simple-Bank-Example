package service

import "sbank/internal/repository"

type Service struct {
	Account
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Account: NewAccountService(repo.Account),
	}
}
