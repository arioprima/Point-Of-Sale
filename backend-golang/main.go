package main

import (
	"github.com/arioprima/Point-Of-Sale/config"
	"github.com/arioprima/Point-Of-Sale/controller"
	_ "github.com/arioprima/Point-Of-Sale/docs"
	"github.com/arioprima/Point-Of-Sale/repository"
	"github.com/arioprima/Point-Of-Sale/router"
	"github.com/arioprima/Point-Of-Sale/service"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"time"
)

// @title Point Of Sale API Documentation
// @version 1.0
// @description Tag a service for point of sale using golang and gin framework

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type "Bearer"
// @host localhost:8080
// basePath: /api
func main() {
	loadConfig, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}
	db := config.ConnectionDB(&loadConfig)
	validate := validator.New()
	userRepository := repository.NewUserRepositoryImpl(db)
	userService := service.NewUserServiceImpl(userRepository, db, validate)
	userController := controller.NewUserController(userService)

	routes := router.NewRouter(userRepository, userController, db)

	server := &http.Server{
		Addr:           ":" + loadConfig.ServerPort,
		Handler:        routes,
		ReadTimeout:    15 * time.Second,
		WriteTimeout:   15 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err_server := server.ListenAndServe()

	if err_server != nil {
		log.Fatal("🚀 Could not start server", err_server)
	}
}
