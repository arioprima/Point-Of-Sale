package main

import (
	"github.com/arioprima/Point-Of-Sale/config"
	"github.com/arioprima/Point-Of-Sale/controller"
	"github.com/arioprima/Point-Of-Sale/repository"
	"github.com/arioprima/Point-Of-Sale/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	router := gin.Default()

	loadConfig, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}
	db := config.ConnectionDB(&loadConfig)
	validate := validator.New()
	userRepository := repository.NewUserRepositoryImpl(db)
	userService := service.NewUserServiceImpl(userRepository, db, validate)
	userController := controller.NewUserController(userService)

	router.POST("/login", userController.Login)
	router.POST("/register", userController.Create)
	router.GET("/users", userController.FindAll)
	router.GET("/users/:id", userController.FindById)
	router.GET("/users/username/:username", userController.FindByUserName)
	router.GET("/users/email/:email", userController.FindByEmail)
	router.PUT("/users/:id", userController.Update)
	router.DELETE("/users/:id", userController.Delete)

	err = router.Run(":8080")
	if err != nil {
		return
	}
}
