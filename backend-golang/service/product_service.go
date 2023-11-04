package service

import (
	"context"
	"database/sql"
	"github.com/arioprima/Point-Of-Sale/models/entity"
	"github.com/arioprima/Point-Of-Sale/models/web/request"
	"github.com/arioprima/Point-Of-Sale/models/web/response"
	"github.com/arioprima/Point-Of-Sale/repository"
	"github.com/arioprima/Point-Of-Sale/utils"
	"github.com/go-playground/validator/v10"
	"log"
	"time"
)

type ProductService interface {
	Create(ctx context.Context, request request.ProductCreateRequest) (response.ProductResponse, error)
	Update(ctx context.Context, request request.ProductUpdateRequest) (response.ProductResponse, error)
	Delete(ctx context.Context, productId string) error
	FindById(ctx context.Context, productId string) (response.ProductResponse, error)
	FindByName(ctx context.Context, productName string) (response.ProductResponse, error)
	FindAll(ctx context.Context) ([]response.ProductResponse, error)
}

type ProductServiceImpl struct {
	ProductRepository repository.ProductRepository
	DB                *sql.DB
	Validate          *validator.Validate
}

func NewProductServiceImpl(productRepository repository.ProductRepository, db *sql.DB, validate *validator.Validate) ProductService {
	return &ProductServiceImpl{
		ProductRepository: productRepository,
		DB:                db,
		Validate:          validate,
	}
}

func (product *ProductServiceImpl) Create(ctx context.Context, request request.ProductCreateRequest) (response.ProductResponse, error) {
	//TODO implement me
	err := product.Validate.Struct(request)
	if err != nil {
		return response.ProductResponse{}, err
	}

	tx, err := product.DB.Begin()
	if err != nil {
		return response.ProductResponse{}, err
	}

	defer func() {
		if r := recover(); r != nil {
			// An error occurred, rollback the transaction
			if rollbackErr := tx.Rollback(); rollbackErr != nil {
				log.Fatalf("Error rolling back transaction: %v", rollbackErr)
			}
		} else {
			// No error, commit the transaction
			if commitErr := tx.Commit(); commitErr != nil {
				log.Fatalf("Error committing transaction: %v", commitErr)
			}
		}
	}()

	productEntity := entity.Product{
		ProductId:        utils.GenerateUUID(),
		ProductName:      request.ProductName,
		CategoryId:       request.CategoryId,
		Price:            request.Price,
		Description:      request.Description,
		Quantity:         request.Quantity,
		ProductCondition: request.ProductCondition,
		Image:            request.Image,
		SupplierId:       request.SupplierId,
		DateOfArrival:    request.DateOfArrival,
		ExpiryDate:       request.ExpiryDate,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}

	productResponse, err := product.ProductRepository.Create(ctx, tx, &productEntity)
	if err != nil {
		return response.ProductResponse{}, err
	}

	return response.ProductResponse{
		ProductId:        productResponse.ProductId,
		ProductName:      productResponse.ProductName,
		CategoryId:       productResponse.CategoryId,
		Price:            productResponse.Price,
		Description:      productResponse.Description,
		Quantity:         productResponse.Quantity,
		ProductCondition: productResponse.ProductCondition,
		Image:            productResponse.Image,
		SupplierId:       productResponse.SupplierId,
		DateOfArrival:    productResponse.DateOfArrival,
		ExpiryDate:       productResponse.ExpiryDate,
		CreatedAt:        productResponse.CreatedAt,
	}, nil
}

func (product *ProductServiceImpl) Update(ctx context.Context, request request.ProductUpdateRequest) (response.ProductResponse, error) {
	//TODO implement me
	err := product.Validate.Struct(request)
	if err != nil {
		return response.ProductResponse{}, err
	}

	tx, err := product.DB.Begin()
	if err != nil {
		return response.ProductResponse{}, err
	}

	defer func() {
		if r := recover(); r != nil {
			// An error occurred, rollback the transaction
			if rollbackErr := tx.Rollback(); rollbackErr != nil {
				log.Fatalf("Error rolling back transaction: %v", rollbackErr)
			}
		} else {
			// No error, commit the transaction
			if commitErr := tx.Commit(); commitErr != nil {
				log.Fatalf("Error committing transaction: %v", commitErr)
			}
		}
	}()

	productEntity := entity.Product{
		ProductId:        request.ProductId,
		ProductName:      request.ProductName,
		CategoryId:       request.CategoryId,
		Price:            request.Price,
		Description:      request.Description,
		Quantity:         request.Quantity,
		ProductCondition: request.ProductCondition,
		Image:            request.Image,
		SupplierId:       request.SupplierId,
		DateOfArrival:    request.DateOfArrival,
		ExpiryDate:       request.ExpiryDate,
		UpdatedAt:        time.Now(),
	}

	productResponse, err := product.ProductRepository.Update(ctx, tx, &productEntity)
	if err != nil {
		return response.ProductResponse{}, err
	}

	return response.ProductResponse{
		ProductId:        productResponse.ProductId,
		ProductName:      productResponse.ProductName,
		CategoryId:       productResponse.CategoryId,
		Price:            productResponse.Price,
		Description:      productResponse.Description,
		Quantity:         productResponse.Quantity,
		ProductCondition: productResponse.ProductCondition,
		Image:            productResponse.Image,
		SupplierId:       productResponse.SupplierId,
		DateOfArrival:    productResponse.DateOfArrival,
		ExpiryDate:       productResponse.ExpiryDate,
		CreatedAt:        productResponse.CreatedAt,
		UpdatedAt:        productResponse.UpdatedAt,
	}, nil

}

func (product *ProductServiceImpl) Delete(ctx context.Context, productId string) error {
	//TODO implement me
	err := product.Validate.Var(productId, "required")
	if err != nil {
		return err
	}

	tx, err := product.DB.Begin()
	if err != nil {
		panic(err)
	}

	defer func() {
		if r := recover(); r != nil {
			// An error occurred, rollback the transaction
			if rollbackErr := tx.Rollback(); rollbackErr != nil {
				log.Fatalf("Error rolling back transaction: %v", rollbackErr)
			}
		} else {
			// No error, commit the transaction
			if commitErr := tx.Commit(); commitErr != nil {
				log.Fatalf("Error committing transaction: %v", commitErr)
			}
		}
	}()

	err = product.ProductRepository.Delete(ctx, tx, productId)
	if err != nil {
		return err

	}

	return nil

}

func (product *ProductServiceImpl) FindById(ctx context.Context, productId string) (response.ProductResponse, error) {
	//TODO implement me
	err := product.Validate.Var(productId, "required")
	if err != nil {
		return response.ProductResponse{}, err
	}

	tx, err := product.DB.Begin()
	if err != nil {
		panic(err)
	}

	defer func() {
		if r := recover(); r != nil {
			// An error occurred, rollback the transaction
			if rollbackErr := tx.Rollback(); rollbackErr != nil {
				log.Fatalf("Error rolling back transaction: %v", rollbackErr)
			}
		} else {
			// No error, commit the transaction
			if commitErr := tx.Commit(); commitErr != nil {
				log.Fatalf("Error committing transaction: %v", commitErr)
			}
		}
	}()

	productEntity, err := product.ProductRepository.FindById(ctx, tx, productId)
	if err != nil {
		return response.ProductResponse{}, err
	}

	return response.ProductResponse{
		ProductId:        productEntity.ProductId,
		ProductName:      productEntity.ProductName,
		CategoryId:       productEntity.CategoryId,
		Price:            productEntity.Price,
		Description:      productEntity.Description,
		Quantity:         productEntity.Quantity,
		ProductCondition: productEntity.ProductCondition,
		Image:            productEntity.Image,
		SupplierId:       productEntity.SupplierId,
		DateOfArrival:    productEntity.DateOfArrival,
		ExpiryDate:       productEntity.ExpiryDate,
		CreatedAt:        productEntity.CreatedAt,
		UpdatedAt:        productEntity.UpdatedAt,
	}, nil
}

func (product *ProductServiceImpl) FindByName(ctx context.Context, productName string) (response.ProductResponse, error) {
	//TODO implement me
	err := product.Validate.Var(productName, "required")
	if err != nil {
		return response.ProductResponse{}, err
	}

	tx, err := product.DB.Begin()
	if err != nil {
		panic(err)
	}

	defer func() {
		if r := recover(); r != nil {
			// An error occurred, rollback the transaction
			if rollbackErr := tx.Rollback(); rollbackErr != nil {
				log.Fatalf("Error rolling back transaction: %v", rollbackErr)
			}
		} else {
			// No error, commit the transaction
			if commitErr := tx.Commit(); commitErr != nil {
				log.Fatalf("Error committing transaction: %v", commitErr)
			}
		}
	}()

	productEntity, err := product.ProductRepository.FindBYName(ctx, tx, productName)
	if err != nil {
		return response.ProductResponse{}, err
	}

	return response.ProductResponse{
		ProductId:        productEntity.ProductId,
		ProductName:      productEntity.ProductName,
		CategoryId:       productEntity.CategoryId,
		Price:            productEntity.Price,
		Description:      productEntity.Description,
		Quantity:         productEntity.Quantity,
		ProductCondition: productEntity.ProductCondition,
		Image:            productEntity.Image,
		SupplierId:       productEntity.SupplierId,
		DateOfArrival:    productEntity.DateOfArrival,
		ExpiryDate:       productEntity.ExpiryDate,
		CreatedAt:        productEntity.CreatedAt,
		UpdatedAt:        productEntity.UpdatedAt,
	}, nil
}

func (product *ProductServiceImpl) FindAll(ctx context.Context) ([]response.ProductResponse, error) {
	//TODO implement me
	tx, err := product.DB.Begin()
	if err != nil {
		panic(err)
	}

	defer func() {
		if r := recover(); r != nil {
			// An error occurred, rollback the transaction
			if rollbackErr := tx.Rollback(); rollbackErr != nil {
				log.Fatalf("Error rolling back transaction: %v", rollbackErr)
			}
		} else {
			// No error, commit the transaction
			if commitErr := tx.Commit(); commitErr != nil {
				log.Fatalf("Error committing transaction: %v", commitErr)
			}
		}
	}()

	productEntity, err := product.ProductRepository.FindAll(ctx, tx)
	if err != nil {
		return []response.ProductResponse{}, err
	}

	var productResponse []response.ProductResponse

	for _, product := range productEntity {
		productResponse = append(productResponse, response.ProductResponse{
			ProductId:        product.ProductId,
			ProductName:      product.ProductName,
			CategoryId:       product.CategoryId,
			Price:            product.Price,
			Description:      product.Description,
			Quantity:         product.Quantity,
			ProductCondition: product.ProductCondition,
			Image:            product.Image,
			SupplierId:       product.SupplierId,
			DateOfArrival:    product.DateOfArrival,
			ExpiryDate:       product.ExpiryDate,
			CreatedAt:        product.CreatedAt,
			UpdatedAt:        product.UpdatedAt,
		})
	}
	return productResponse, nil
}
