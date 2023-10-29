package controller

import (
	"github.com/arioprima/Point-Of-Sale/models/web/request"
	"github.com/arioprima/Point-Of-Sale/models/web/response"
	"github.com/arioprima/Point-Of-Sale/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) *UserController {
	return &UserController{UserService: userService}
}

// Login handles user login.
// @Summary Login
// @Description Logs in a user.
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body string true "Username and Password"
// @Success 200 {string} string "OK"
// @Router /api/auth/login [post]
func (controller *UserController) Login(ctx *gin.Context) {
	loginRequest := request.UserLoginRequest{}
	err := ctx.ShouldBindJSON(&loginRequest)
	if err != nil {
		panic(err)
	}

	responses, err := controller.UserService.Login(ctx, loginRequest)

	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, response.Response{
			Status:  http.StatusNotFound,
			Message: err.Error(),
			Data:    nil,
		})
	} else {
		ctx.IndentedJSON(http.StatusOK, response.Response{
			Status:  http.StatusOK,
			Message: "Success",
			Data:    responses,
		})
	}
}

// Create registers a new user.
// @Summary Register
// @Description Registers a new user.
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body string true
// @Success 201 {string} string "Created"
// @Router /api/auth/register [post]
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
	} else {
		ctx.IndentedJSON(http.StatusCreated, response.Response{
			Status:  http.StatusCreated,
			Message: "Success",
			Data:    responses,
		})
	}

}

// Update modifies user information.
// @Summary Update
// @Description Updates user information.
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param username body string true "Username"
// @Param email body string true "Email"
// @Param password body string true "Password"
// @Param role body string true "Role"
// @Success 200 {string} string "OK"
// @Router /api/users/{id} [put]
// @Security Bearer
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
	} else {
		ctx.IndentedJSON(http.StatusCreated, response.Response{
			Status:  http.StatusCreated,
			Message: "Success",
			Data:    responses,
		})
	}
}

// Update modifies user information.
// @Summary Update
// @Description Updates user information.
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param username body string true "Username"
// @Param email body string true "Email"
// @Param password body string true "Password"
// @Param role body string true "Role"
// @Success 200 {string} string "OK"
// @Router /api/users/{id} [put]
// @Security Bearer
func (controller *UserController) Delete(ctx *gin.Context) {
	userId := ctx.Param("id")

	err := controller.UserService.Delete(ctx, userId)
	if err != nil {
		return
	}

	ctx.IndentedJSON(http.StatusOK, response.Response{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    nil,
	})
}

// FindById retrieves user by ID.
// @Summary Find By Id
// @Description Finds a user by ID.
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {string} string "OK"
// @Router /api/users/{id} [get]
// @Security Bearer
func (controller *UserController) FindById(ctx *gin.Context) {
	userId := ctx.Param("id")

	responses, err := controller.UserService.FindById(ctx, userId)
	if err != nil {
		// Menangani kesalahan dari service
		ctx.IndentedJSON(http.StatusInternalServerError, response.Response{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	if responses.ID == "" || responses.UserName == "" {
		// Jika ID kosong, berarti pengguna tidak ditemukan
		ctx.IndentedJSON(http.StatusInternalServerError, response.Response{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	ctx.IndentedJSON(http.StatusOK, response.Response{
		Status:  http.StatusOK,
		Message: "Get user by userId success",
		Data:    responses,
	})
	return
}

// FindByUserName retrieves user by username.
// @Summary Find By Username
// @Description Finds a user by username.
// @Tags User
// @Accept json
// @Produce json
// @Param username path string true "Username"
// @Success 200 {string} string "OK"
// @Router /api/users/username/{username} [get]
// @Security Bearer
func (controller *UserController) FindByUserName(ctx *gin.Context) {
	userName := ctx.Param("username")

	responses, err := controller.UserService.FindByUserName(ctx, userName)
	if err != nil {
		// Menangani kesalahan dari service
		ctx.IndentedJSON(http.StatusInternalServerError, response.Response{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	if responses.ID == "" || responses.UserName == "" {
		// Jika ID kosong, berarti pengguna tidak ditemukan
		ctx.IndentedJSON(http.StatusInternalServerError, response.Response{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	ctx.IndentedJSON(http.StatusOK, response.Response{
		Status:  http.StatusOK,
		Message: "Get user by username success",
		Data:    responses,
	})
	return
}

// FindByEmail retrieves user by email.
// @Summary Find By Email
// @Description Finds a user by email.
// @Tags User
// @Accept json
// @Produce json
// @Param email path string true "Email"
// @Success 200 {string} string "OK"
// @Router /api/users/email/{email} [get]
// @Security Bearer
func (controller *UserController) FindByEmail(ctx *gin.Context) {
	email := ctx.Param("email")

	responses, err := controller.UserService.FindByEmail(ctx, email)
	if err != nil {
		// Menangani kesalahan dari service
		ctx.IndentedJSON(http.StatusInternalServerError, response.Response{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	if responses.ID == "" || responses.UserName == "" {
		// Jika ID kosong, berarti pengguna tidak ditemukan
		ctx.IndentedJSON(http.StatusInternalServerError, response.Response{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	ctx.IndentedJSON(http.StatusOK, response.Response{
		Status:  http.StatusOK,
		Message: "Get user by email success",
		Data:    responses,
	})
	return
}

// FindAll retrieves all users.
// @Summary Find All
// @Description Retrieves all users.
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {string} string "OK"
// @Router /api/users [get]
// @Security Bearer
func (controller *UserController) FindAll(ctx *gin.Context) {
	responses, err := controller.UserService.FindAll(ctx)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, response.Response{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	ctx.IndentedJSON(http.StatusOK, response.Response{
		Status:  http.StatusOK,
		Message: "Get all user success",
		Data:    responses,
	})
}
