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

	storeRepo := repo.NewStoreRepositoryImpl(db)
	storeService := service.NewStoreService(storeRepo)
	storeHandler := handler.NewStoreHandler(storeService)

	userRepo := repo.NewUserRepositoryImpl(db)
	userService := service.NewUserService(userRepo, storeRepo)
	userHandler := handler.NewUserHandler(userService)

	if err := router.New("8080", &router.Handlers{
		User:  userHandler,
		Store: storeHandler,
	}).Start(); err != nil {
		log.Fatal(err)
	}
}
