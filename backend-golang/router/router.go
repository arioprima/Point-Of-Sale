// router.go
package router

import (
	"database/sql"
	"github.com/arioprima/Point-Of-Sale/config"
	"github.com/arioprima/Point-Of-Sale/controller"
	"github.com/arioprima/Point-Of-Sale/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter(userController *controller.UserController, productController *controller.ProductController, db *sql.DB, config *config.Config) *gin.Engine {
	router := gin.Default()
	middleware.SetupCorsMiddleware(router)

	// add swagger
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(404, gin.H{
			"code":    "PAGE_NOT_FOUND",
			"message": "Page not found",
		})
	})

	// User routes
	router.POST("/api/auth/login", userController.Login)
	router.POST("/api/auth/register", userController.Create)
	router.GET("/api/users", middleware.AuthMiddleware("staff"), func(ctx *gin.Context) {
		userController.FindAll(ctx)
	})
	router.GET("/api/users/:id", middleware.AuthMiddleware("staff"), func(ctx *gin.Context) {
		userController.FindById(ctx)
	})

	router.PUT("/api/users/:id", middleware.AuthMiddleware("staff"), func(ctx *gin.Context) {
		userController.Update(ctx)
	})
	router.DELETE("/api/users/:id", middleware.AuthMiddleware("staff"), func(ctx *gin.Context) {
		userController.Delete(ctx)
	})

	// Product routes
	router.POST("/api/products", middleware.AuthMiddleware("staff"), func(ctx *gin.Context) {
		productController.Create(ctx)
	})
	router.GET("/api/products", middleware.AuthMiddleware("staff"), func(ctx *gin.Context) {
		productController.FindAll(ctx)
	})
	router.GET("/api/products/:id", middleware.AuthMiddleware("staff"), func(ctx *gin.Context) {
		productController.FindById(ctx)
	})
	router.GET("/api/products/name/:name", middleware.AuthMiddleware("staff"), func(ctx *gin.Context) {
		productController.FindByName(ctx)
	})
	router.PUT("/api/products/:id", middleware.AuthMiddleware("staff"), func(ctx *gin.Context) {
		productController.Update(ctx)
	})
	router.POST("/api/products/:id", middleware.AuthMiddleware("staff"), func(ctx *gin.Context) {
		productController.Delete(ctx)
	})

	return router
}
