package main

import (
	"log"
	"net/http"
	"sbank/config"
	"sbank/internal/controller"
	"sbank/internal/repository"
	"sbank/internal/service"
	"sbank/pkg/db/postgres"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func main() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load")
	}

	db := postgres.InitDB(config)

	repo := repository.NewRepository(db)
	service := service.NewService(repo, config.TokenSymmetricKey)
	handler := controller.NewHandler(service, config)

	router := handler.InitRoutes()

	srv := new(Server)

	log.Println("Starting the server")
	if err := srv.Start("8000", router); err != nil {
		log.Fatalf("error server: %v", err)
	}
}

func (s *Server) Start(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:         ":" + port,
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	return s.httpServer.ListenAndServe()
}
