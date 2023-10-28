package router

import (
	"github.com/arioprima/Point-Of-Sale/controller"
	"github.com/arioprima/Point-Of-Sale/repository"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter(repository repository.UserRepository, userController *controller.UserController) *gin.Engine {
	router := gin.Default()

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
	router.GET("/api/users", userController.FindAll)
	router.GET("/api/users/:id", userController.FindById)
	router.GET("/api/users/username/:username", userController.FindByUserName)
	router.GET("/api/users/email/:email", userController.FindByEmail)
	router.PUT("/api/users/:id", userController.Update)
	router.POST("/api/users/:id", userController.Delete)

	return router
}
