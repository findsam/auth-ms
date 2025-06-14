package cmd

import (
	"log"

	"github.com/findsam/auth-micro/internal/handler"
	"github.com/findsam/auth-micro/internal/repo"
	"github.com/findsam/auth-micro/internal/router"
	"github.com/findsam/auth-micro/internal/service"
	"github.com/findsam/auth-micro/pkg/mongo"
)

func Execute() {
	db, err := mongo.New()
	if err != nil {
		log.Fatal(err)
	}

	userRepo := repo.NewUserRepositoryImpl(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	
	storeHandler := handler.NewStoreHandler()

	if err := router.New("8080", &router.Handlers{
		User: userHandler,
		Store: storeHandler, 
	}).Start(); err != nil {
		log.Fatal(err)
	}
}
