package controller

import (
	"github.com/arioprima/blog_web/models/web/request"
	"github.com/arioprima/blog_web/models/web/response"
	"github.com/arioprima/blog_web/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) *UserController {
	return &UserController{UserService: userService}
}

func (controller *UserController) Create(ctx *gin.Context) {
	createUserRequest := request.UserCreateRequest{}
	err := ctx.ShouldBindJSON(&createUserRequest)

	if err != nil {
		panic(err)
	}

	responses, err := controller.UserService.Create(ctx, createUserRequest)

	if err != nil {
		ctx.IndentedJSON(http.StatusUnauthorized, response.Response{
			Status:  http.StatusUnauthorized,
			Message: err.Error(),
			Data:    nil,
		})
	}

	ctx.IndentedJSON(http.StatusCreated, response.Response{
		Status:  http.StatusCreated,
		Message: "Success",
		Data:    responses,
	})
}

func (controller *UserController) Update(ctx *gin.Context) {
	updateUserRequest := request.UserUpdateRequest{}
	err := ctx.ShouldBindJSON(&updateUserRequest)
	if err != nil {
		panic(err)
	}

	responses, err := controller.UserService.Update(ctx, updateUserRequest)
	if err != nil {
		ctx.IndentedJSON(http.StatusUnauthorized, response.Response{
			Status:  http.StatusUnauthorized,
			Message: err.Error(),
			Data:    nil,
		})
	}

	ctx.IndentedJSON(http.StatusCreated, response.Response{
		Status:  http.StatusCreated,
		Message: "Success",
		Data:    responses,
	})
}

func (controller *UserController) Delete(ctx *gin.Context) {
	userId := ctx.Param("id")

	controller.UserService.Delete(ctx, userId)

	ctx.IndentedJSON(http.StatusOK, response.Response{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    nil,
	})
}

func (controller *UserController) FindById(ctx *gin.Context) {
	userId := ctx.Param("id")

	responses, err := controller.UserService.FindById(ctx, userId)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, response.Response{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})
	}

	if responses.UserName == "" || responses.Email == "" {
		ctx.IndentedJSON(http.StatusNotFound, response.Response{
			Status:  http.StatusNotFound,
			Message: "User not found",
			Data:    nil,
		})
	}

	ctx.IndentedJSON(http.StatusOK, response.Response{
		Status:  http.StatusOK,
		Message: "Get user by id success",
		Data:    responses,
	})
}

func (controller *UserController) FindByUserName(ctx *gin.Context) {
	userName := ctx.Param("username")

	responses, err := controller.UserService.FindByUserName(ctx, userName)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, response.Response{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})
	}

	if responses.UserName == "" || responses.Email == "" {
		ctx.IndentedJSON(http.StatusNotFound, response.Response{
			Status:  http.StatusNotFound,
			Message: "User not found",
			Data:    nil,
		})
	}

	ctx.IndentedJSON(http.StatusOK, response.Response{
		Status:  http.StatusOK,
		Message: "Get user by username success",
		Data:    responses,
	})
}

func (controller *UserController) FindByEmail(ctx *gin.Context) {
	email := ctx.Param("email")

	responses, err := controller.UserService.FindByEmail(ctx, email)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, response.Response{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})
	}

	if responses.UserName == "" || responses.Email == "" {
		ctx.IndentedJSON(http.StatusNotFound, response.Response{
			Status:  http.StatusNotFound,
			Message: "User not found",
			Data:    nil,
		})
	}

	ctx.IndentedJSON(http.StatusOK, response.Response{
		Status:  http.StatusOK,
		Message: "Get user by email success",
		Data:    responses,
	})
}

func (controller *UserController) FindAll(ctx *gin.Context) {
	responses, err := controller.UserService.FindAll(ctx)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, response.Response{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})
		return // Tambahkan pernyataan return di sini
	}

	ctx.IndentedJSON(http.StatusOK, response.Response{
		Status:  http.StatusOK,
		Message: "Get all user success",
		Data:    responses,
	})
}
