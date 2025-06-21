package cmd

import (
	"log"

	"github.com/findsam/auth-micro/internal/handler"
	"github.com/findsam/auth-micro/internal/repo"
	"github.com/findsam/auth-micro/internal/router"
	"github.com/findsam/auth-micro/internal/service"
	m "github.com/findsam/auth-micro/pkg/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func Execute() error {
	db, err := m.New()
	if err != nil {
		log.Fatal(err)
	}

	deps := &router.Handlers{
		Store: buildStoreHandler(db),
		Payment: buildPaymentHandler(db),
		User: buildUserHandler(db),
	}

	router := router.New("8080", deps)
	return router.Start()

}


func buildStoreHandler(db *mongo.Database) *handler.StoreHandler {
	storeRepo := repo.NewStoreRepositoryImpl(db)
	storeService := service.NewStoreService(storeRepo)
	storeHandler := handler.NewStoreHandler(storeService)
	return storeHandler
}

func buildUserHandler(db *mongo.Database) *handler.UserHandler {
	userRepo := repo.NewUserRepositoryImpl(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)
	return userHandler
}

func buildPaymentHandler(db *mongo.Database) *handler.PaymentHandler {
	paymentRepo := repo.NewPaymentRepositoryImpl(db)
	paymentService := service.NewPaymentService(paymentRepo)
	paymentHandler := handler.NewPaymentHandler(paymentService)
	return paymentHandler
}