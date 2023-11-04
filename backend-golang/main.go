// main.go
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

func main() {
	loadConfig, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("Could not load environment variables: %v", err)
	}

	db := config.ConnectionDB(&loadConfig)
	defer db.Close()

	validate := validator.New()
	userRepository := repository.NewUserRepositoryImpl(db)
	userService := service.NewUserServiceImpl(userRepository, db, validate)
	userController := controller.NewUserController(userService)

	productRepository := repository.NewProductRepositoryImpl(db)
	productService := service.NewProductServiceImpl(productRepository, db, validate)
	productController := controller.NewProductController(productService)

	routes := router.SetupRouter(userController, productController, db, &loadConfig)

	server := &http.Server{
		Addr:           ":" + loadConfig.ServerPort,
		Handler:        routes,
		ReadTimeout:    15 * time.Second,
		WriteTimeout:   15 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Printf("Server is running on :%s\n", loadConfig.ServerPort)
	if err_server := server.ListenAndServe(); err_server != nil {
		log.Fatalf("Could not start server: %v", err_server)
	}
}
