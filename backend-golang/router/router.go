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

// RouterConfig adalah struktur data untuk mengelompokkan parameter-parameter yang dibutuhkan oleh SetupRouter
type RouterConfig struct {
	UserController    *controller.UserController
	ProductController *controller.ProductController
	DB                *sql.DB
	Config            *config.Config
}

func SetupRouter(config RouterConfig) *gin.Engine {
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
	router.POST("/api/auth/login", config.UserController.Login)
	router.POST("/api/auth/register", config.UserController.Create)
	router.GET("/api/users", middleware.AuthMiddleware("staff"), func(ctx *gin.Context) {
		config.UserController.FindAll(ctx)
	})
	router.GET("/api/users/:id", middleware.AuthMiddleware("staff"), func(ctx *gin.Context) {
		config.UserController.FindById(ctx)
	})

	router.PUT("/api/users/:id", middleware.AuthMiddleware("staff"), func(ctx *gin.Context) {
		config.UserController.Update(ctx)
	})
	router.DELETE("/api/users/:id", middleware.AuthMiddleware("staff"), func(ctx *gin.Context) {
		config.UserController.Delete(ctx)
	})

	// Product routes
	router.POST("/api/products", middleware.AuthMiddleware("staff"), func(ctx *gin.Context) {
		config.ProductController.Create(ctx)
	})
	router.GET("/api/products", middleware.AuthMiddleware("staff"), func(ctx *gin.Context) {
		config.ProductController.FindAll(ctx)
	})
	router.GET("/api/products/:id", middleware.AuthMiddleware("staff"), func(ctx *gin.Context) {
		config.ProductController.FindById(ctx)
	})
	router.GET("/api/products/name/:name", middleware.AuthMiddleware("staff"), func(ctx *gin.Context) {
		config.ProductController.FindByName(ctx)
	})
	router.PUT("/api/products/:id", middleware.AuthMiddleware("staff"), func(ctx *gin.Context) {
		config.ProductController.Update(ctx)
	})
	router.POST("/api/products/delete/:product_id", middleware.AuthMiddleware("staff"), func(ctx *gin.Context) {
		config.ProductController.Delete(ctx)
	})

	return router
}
