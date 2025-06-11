package cmd

import (
	"log"

	"github.com/findsam/auth-micro/internal/handler"
	"github.com/findsam/auth-micro/internal/repo"
	"github.com/findsam/auth-micro/internal/router"
	"github.com/findsam/auth-micro/internal/service"
	"github.com/findsam/auth-micro/pkg/mongo"
	"github.com/findsam/auth-micro/pkg/util"
)

func Execute() {
	db, err := mongo.New()
	if err != nil {
		log.Fatal(err)
	}

	validator := util.NewValidator()
	baseHandler := handler.NewBaseHanlder(validator)

	userRepo := repo.NewUserRepositoryImpl(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(baseHandler, userService)

	if err := router.New("8080", &router.Handlers{
		User: userHandler,
	}).Start(); err != nil {
		log.Fatal(err)
	}
}
