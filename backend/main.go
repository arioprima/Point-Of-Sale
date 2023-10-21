package main

import (
	"github.com/arioprima/blog_web/config"
	"github.com/arioprima/blog_web/controller"
	"github.com/arioprima/blog_web/middleware"
	"github.com/arioprima/blog_web/repository"
	"github.com/arioprima/blog_web/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	r := gin.Default()

	// connect to database
	db := config.ConnectionDB()

	// validate
	validate := validator.New()

	//repository
	userRepository := repository.NewUserRepositoryImpl(db)

	//service
	userService := service.NewUserServiceImpl(userRepository, db, validate)

	//controller
	userController := controller.NewUserController(userService)

	userRouter := r.Group("/api")

	userRouter.GET("/users", middleware.UserHandler(userRepository, db, "user"), userController.FindAll)

	r.POST("/auth/login", userController.Login)

	r.GET("/users/:id", userController.FindById)
	r.POST("/auth/register", userController.Create)
	r.PUT("/users", userController.Update)
	r.DELETE("/users/:id", userController.Delete)

	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
