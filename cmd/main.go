package main

import (
	a "sbank/internal/adapters/api/account"
	"sbank/internal/repository"
	"sbank/internal/service/account"
	"sbank/pkg/store/postgres"
)

func main() {
	db := postgres.InitDB()

	repo := repository.NewRepository(db)
	service := account.NewAccountService(repo)
	handler := a.NewHandler(service)
}
