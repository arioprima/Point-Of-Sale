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
// @Summary Register a new product
// @Description Registers a new product.
// @Tags Product
// @Accept json
// @Produce json
// @Param request body ProductCreateRequest true "Product creation request"
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
		ctx.JSON(http.StatusNotFound, response.Response{
			Status:  http.StatusNotFound,
			Message: err.Error(),
			Data:    nil,
		})
	} else {
		ctx.JSON(http.StatusCreated, response.Response{
			Status:  http.StatusCreated,
			Message: "Created",
			Data:    responses,
		})
	}
}

// Update modifies an existing product.
// @Summary Update an existing product
// @Description Modifies an existing product.
// @Tags Product
// @Accept json
// @Produce json
// @Param request body ProductUpdateRequest true "Product update request"
// @Success 200 {string} string "OK"
// @Router /api/product [put]
func (controller *ProductController) Update(ctx *gin.Context) {
	// TODO: Implement the update logic here
	updateProductRequest := request.ProductUpdateRequest{}
	err := ctx.ShouldBindJSON(&updateProductRequest)
	if err != nil {
		panic(err)
	}

	responses, err := controller.ProductService.Update(ctx, updateProductRequest)

	if err != nil {
		ctx.JSON(http.StatusNotFound, response.Response{
			Status:  http.StatusNotFound,
			Message: err.Error(),
			Data:    nil,
		})
	} else {
		ctx.JSON(http.StatusOK, response.Response{
			Status:  http.StatusOK,
			Message: "OK",
			Data:    responses,
		})
	}
}

// Delete deletes an existing product.
// @Summary Delete an existing product
// @Description Deletes an existing product by its ID.
// @Tags Product
// @Accept json
// @Produce json
// @Param product_id path string true "Product ID to delete"
// @Success 200 {string} string "OK"
// @Router /api/product/{product_id} [delete]
func (controller *ProductController) Delete(ctx *gin.Context) {
	// TODO: Implement the delete logic here
	productId := ctx.Param("product_id")

	err := controller.ProductService.Delete(ctx, productId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, response.Response{
			Status:  http.StatusNotFound,
			Message: err.Error(),
			Data:    nil,
		})
	} else {
		ctx.JSON(http.StatusOK, response.Response{
			Status:  http.StatusOK,
			Message: "OK",
			Data:    nil,
		})
	}
}

// FindById Retrieves a product by ID.
// @Summary Find a product by ID
// @Description Retrieves a product by its ID.
// @Tags Product
// @Accept json
// @Produce json
// @Param product_id path string true "Product ID to retrieve"
// @Success 200 {string} string "OK"
// @Router /api/product/{product_id} [get]
func (controller *ProductController) FindById(ctx *gin.Context) {
	// TODO: Implement the retrieval logic here
	productId := ctx.Param("product_id")

	responses, err := controller.ProductService.FindById(ctx, productId)

	if err != nil {
		ctx.JSON(http.StatusNotFound, response.Response{
			Status:  http.StatusNotFound,
			Message: err.Error(),
			Data:    nil,
		})
	} else {
		ctx.JSON(http.StatusOK, response.Response{
			Status:  http.StatusOK,
			Message: "OK",
			Data:    responses,
		})
	}
}

// FindByName retrieves a product by name.
// @Summary Find a product by name
// @Description Retrieves a product by its name.
// @Tags Product
// @Accept json
// @Produce json
// @Param product_name path string true "Product name to retrieve"
// @Success 200 {string} string "OK"
// @Router /api/product/{product_name} [get]
func (controller *ProductController) FindByName(ctx *gin.Context) {
	// TODO: Implement the retrieval logic here
	productName := ctx.Param("product_name")

	responses, err := controller.ProductService.FindByName(ctx, productName)

	if err != nil {
		ctx.JSON(http.StatusNotFound, response.Response{
			Status:  http.StatusNotFound,
			Message: err.Error(),
			Data:    nil,
		})
	} else {
		ctx.JSON(http.StatusOK, response.Response{
			Status:  http.StatusOK,
			Message: "OK",
			Data:    responses,
		})
	}
}

// FindAll retrieves all products.
// @Summary Retrieve all products
// @Description Retrieves all products.
// @Tags Product
// @Accept json
// @Produce json
// @Success 200 {string} string "OK"
// @Router /api/product [get]
func (controller *ProductController) FindAll(ctx *gin.Context) {
	// TODO: Implement the retrieval logic here
	responses, err := controller.ProductService.FindAll(ctx)

	if err != nil {
		ctx.JSON(http.StatusNotFound, response.Response{
			Status:  http.StatusNotFound,
			Message: err.Error(),
			Data:    nil,
		})
	} else {
		ctx.JSON(http.StatusOK, response.Response{
			Status:  http.StatusOK,
			Message: "OK",
			Data:    responses,
		})
	}
}
