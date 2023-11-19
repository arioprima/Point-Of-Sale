package controller

import (
	"github.com/arioprima/Point-Of-Sale/tree/main/backend/go-postgres/models"
	"github.com/arioprima/Point-Of-Sale/tree/main/backend/go-postgres/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthController struct {
	AuthService service.AuthService
}

func NewAuthController(authService service.AuthService) *AuthController {
	return &AuthController{AuthService: authService}
}

func (controller *AuthController) Login(ctx *gin.Context) {
	loginRequest := models.LoginRequest{}
	err := ctx.ShouldBindJSON(&loginRequest)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	loginResponse, err := controller.AuthService.Login(ctx, loginRequest)

	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{
			"Status":  http.StatusInternalServerError,
			"Message": "Internal Server Error",
		})
	} else {
		ctx.IndentedJSON(http.StatusOK, gin.H{
			"Status":  http.StatusOK,
			"Message": "OK",
			"Data":    loginResponse,
		})
	}
}

func (controller *AuthController) Register(ctx *gin.Context) {
	registerRequest := models.RegisterRequest{}
	err := ctx.ShouldBindJSON(&registerRequest)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	registerResponse, err := controller.AuthService.Register(ctx, registerRequest)

	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{
			"Status":  http.StatusInternalServerError,
			"Message": "Internal Server Error",
			"Error":   err.Error(),
		})
	} else {
		ctx.IndentedJSON(http.StatusCreated, gin.H{
			"Status":  http.StatusCreated,
			"Message": "OK",
			"Data":    registerResponse,
		})
	}
}
