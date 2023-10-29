package controller

import (
	"github.com/arioprima/Point-Of-Sale/models/web/request"
	"github.com/arioprima/Point-Of-Sale/models/web/response"
	"github.com/arioprima/Point-Of-Sale/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ProductController struct {
	ProductService service.ProductService
}

func NewProductController(productService service.ProductService) *ProductController {
	return &ProductController{ProductService: productService}
}

// Create registers a new product.
// @Summary Register
// @Description Registers a new product.
// @Tags Product
// @Accept json
// @Produce json
// @Param request body string true
// @Success 201 {string} string "Created"
// @Router /api/product [post]
func (controller *ProductController) Create(ctx *gin.Context) {
	//TODO implement me
	createProductRequest := request.ProductCreateRequest{}
	err := ctx.ShouldBindJSON(&createProductRequest)
	if err != nil {
		panic(err)
	}

	responses, err := controller.ProductService.Create(ctx, createProductRequest)

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

// Update registers a new product.
// @Summary Register
// @Description Registers a new product.
// @Tags Product
// @Accept json
// @Produce json
// @Param request body string true
// @Success 201 {string} string "Created"
// @Router /api/product [put]
func (controller *ProductController) Update(ctx *gin.Context) {
	//TODO implement me
	updateProductRequest := request.ProductUpdateRequest{}
	err := ctx.ShouldBindJSON(&updateProductRequest)
	if err != nil {
		panic(err)
	}

	responses, err := controller.ProductService.Update(ctx, updateProductRequest)

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

// Delete registers a new product.
// @Summary Register
// @Description Registers a new product.
// @Tags Product
// @Accept json
// @Produce json
// @Param request body string true
// @Success 201 {string} string "Created"
// @Router /api/product [delete]
func (controller *ProductController) Delete(ctx *gin.Context) {
	//TODO implement me
	productId := ctx.Param("product_id")

	err := controller.ProductService.Delete(ctx, productId)
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
			Data:    nil,
		})
	}
}

// FindById registers a new product.
// @Summary Register
// @Description Registers a new product.
// @Tags Product
// @Accept json
// @Produce json
// @Param request body string true
// @Success 201 {string} string "Created"
// @Router /api/product [get]
func (controller *ProductController) FindById(ctx *gin.Context) {
	//TODO implement me
	productId := ctx.Param("product_id")

	responses, err := controller.ProductService.FindById(ctx, productId)

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

// FindByName registers a new product.
// @Summary Register
// @Description Registers a new product.
// @Tags Product
// @Accept json
// @Produce json
// @Param request body string true
// @Success 201 {string} string "Created"
// @Router /api/product [get]
func (controller *ProductController) FindByName(ctx *gin.Context) {
	//TODO implement me
	productName := ctx.Param("product_name")

	responses, err := controller.ProductService.FindByName(ctx, productName)

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

// FindAll registers a new product.
// @Summary Register
// @Description Registers a new product.
// @Tags Product
// @Accept json
// @Produce json
// @Param request body string true
// @Success 201 {string} string "Created"
// @Router /api/product [get]
func (controller *ProductController) FindAll(ctx *gin.Context) {
	//TODO implement me
	responses, err := controller.ProductService.FindAll(ctx)

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
