package cmd

import (
	"github.com/findsam/auth-micro/internal/handler"
	"github.com/findsam/auth-micro/internal/repo"
	"github.com/findsam/auth-micro/internal/router"
	"github.com/findsam/auth-micro/internal/service"
	m "github.com/findsam/auth-micro/pkg/mongo"
)

func Execute() error {
	db, err := m.New()
	if err != nil {
		return err
	}

	userRepo := repo.NewUserRepositoryImpl(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)
	
	storeRepo := repo.NewStoreRepositoryImpl(db)
	storeService := service.NewStoreService(storeRepo, userRepo)
	storeHandler := handler.NewStoreHandler(storeService)
	
	paymentRepo := repo.NewPaymentRepositoryImpl(db)
	paymentService := service.NewPaymentService(paymentRepo, storeRepo, userRepo)
	paymentHandler := handler.NewPaymentHandler(paymentService)


	deps := &router.Handlers{
		Store:   storeHandler,
		Payment: paymentHandler,
		User:    userHandler,
	}

	router := router.New("8080", deps)
	return router.Start()
}
