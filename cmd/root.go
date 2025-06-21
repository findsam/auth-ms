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
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)


	paymentRepo := repo.NewPaymentRepositoryImpl(db)
	paymentService := service.NewPaymentService(paymentRepo)
	paymentHandler := handler.NewPaymentHandler(paymentService)

	router := router.New("8080", &router.Handlers{
		User:  userHandler,
		Store: storeHandler,
		Payment: paymentHandler,
	})

	err = router.Start()
	if err != nil {
		log.Fatal(err)
	}
}
