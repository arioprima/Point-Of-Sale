package main

import (
	"github.com/arioprima/Point-Of-Sale/tree/main/backend/go-postgres/controller"
	"github.com/arioprima/Point-Of-Sale/tree/main/backend/go-postgres/initializers"
	"github.com/arioprima/Point-Of-Sale/tree/main/backend/go-postgres/repository"
	"github.com/arioprima/Point-Of-Sale/tree/main/backend/go-postgres/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	db, err := initializers.ConnectDB(&config)
	if err != nil {
		log.Fatal("🚀 Could not connect to database", err)
	}

	// Check if db is connected
	err = db.Ping()
	if err != nil {
		log.Fatal("🚀 Could not ping database", err)
	}

	log.Println("🚀 Database connection successful")

	// Close the database connection when done
	defer func() {
		err := db.Close()
		if err != nil {
			log.Println("🚀 Could not close database connection", err)
		}
	}()

	router := gin.Default()

	validate := validator.New()
	authRepository := repository.NewAuthRepositoryImpl(db)
	authService := service.NewAuthServiceImpl(authRepository, db, validate)
	authController := controller.NewAuthController(authService)

	router.POST("/api/auth/login", authController.Login)

	router.POST("/api/auth/register", authController.Register)

	if err := router.Run(":8080"); err != nil {
		log.Fatal("🚀 Failed to start server:", err)
	}
	log.Println("🚀 Server started")
}
