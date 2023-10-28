package router

import (
	"database/sql"
	"github.com/arioprima/Point-Of-Sale/controller"
	"github.com/arioprima/Point-Of-Sale/middleware"
	"github.com/arioprima/Point-Of-Sale/repository"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter(userRepository repository.UserRepository, userController *controller.UserController, db *sql.DB) *gin.Engine {
	router := gin.Default()
	middleware.SetupCorsMiddleware(router)

	// add swagger
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.NoRoute(func(ctx *gin.Context) {
		ctx.IndentedJSON(404, gin.H{
			"code":    "PAGE_NOT_FOUND",
			"message": "Page not found",
		})
	})

	router.POST("/api/auth/login", userController.Login)
	router.POST("/api/auth/register", userController.Create)
	router.GET("/api/users", middleware.UserHandler(userRepository, db, "staff"), userController.FindAll)
	router.GET("/api/users/:id", middleware.UserHandler(userRepository, db, "staff"), userController.FindById)
	router.GET("/api/users/username/:username", middleware.UserHandler(userRepository, db, "staff"), userController.FindByUserName)
	router.GET("/api/users/email/:email", middleware.UserHandler(userRepository, db, "staff"), userController.FindByEmail)
	router.PUT("/api/users/:id", middleware.UserHandler(userRepository, db, "employee"), userController.Update)
	router.POST("/api/users/:id", middleware.UserHandler(userRepository, db, "admin"), userController.Update)

	return router
}
