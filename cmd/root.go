package cmd

import (
	"log"

	"github.com/findsam/auth-micro/internal/handler"
	"github.com/findsam/auth-micro/internal/repo"
	"github.com/findsam/auth-micro/internal/router"
	"github.com/findsam/auth-micro/internal/service"
	"github.com/findsam/auth-micro/pkg/mongo"
	util "github.com/findsam/auth-micro/pkg/util"
)	


func Execute(){
	client, err := mongo.New();
	if err != nil {
		log.Fatal(err)
	}

	userRepo := repo.NewUserRepositoryImpl(client) 
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	if err := router.New("8080", &util.Handlers{
		User: userHandler,
	}).Start(); err != nil {
		log.Fatal(err)
	}
}