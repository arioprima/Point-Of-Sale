package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/arioprima/Point-Of-Sale/models/entity"
)

type ProductRepository interface {
	Create(ctx context.Context, tx *sql.Tx, product *entity.Product) (*entity.Product, error)
	Update(ctx context.Context, tx *sql.Tx, product *entity.Product) (*entity.Product, error)
	Delete(ctx context.Context, tx *sql.Tx, productId string) error
	FindById(ctx context.Context, tx *sql.Tx, productId string) (*entity.Product, error)
	FindBYName(ctx context.Context, tx *sql.Tx, productName string) (*entity.Product, error)
	FindAll(ctx context.Context, tx *sql.Tx) ([]*entity.Product, error)
}

type ProductRepositoryImpl struct {
	DB *sql.DB
}

func NewProductRepositoryImpl(db *sql.DB) ProductRepository {
	return &ProductRepositoryImpl{DB: db}
}

func (p *ProductRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, product *entity.Product) (*entity.Product, error) {
	//TODO implement me
	var count int
	checkQuery := "SELECT COUNT(*) FROM products WHERE product_name = ?"
	err := tx.QueryRowContext(ctx, checkQuery, product.ProductName).Scan(&count)
	if err != nil {
		return nil, err
	}

	if count > 0 {
		return nil, errors.New("product already exist")
	}

	SQL := "INSERT INTO products (product_id, product_name, category_id, price, description, " +
		"quantity, `product_condition`, image, supplier_id, date_of_arrival, expiry_date, created_at, updated_at) " +
		"VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

	_, err = tx.ExecContext(
		ctx,
		SQL,
		product.ProductId,
		product.ProductName,
		product.CategoryId,
		product.Price,
		product.Description,
		product.Quantity,
		product.ProductCondition,
		product.Image,
		product.SupplierId,
		product.DateOfArrival,
		product.ExpiryDate,
		product.CreatedAt,
		product.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *ProductRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, product *entity.Product) (*entity.Product, error) {
	//TODO implement me
	checkQuery := "SELECT product_id FROM products WHERE product_id = ? LIMIT 1"
	var existingProductId string
	err := tx.QueryRowContext(ctx, checkQuery, product.ProductId).Scan(&existingProductId)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("product not found")
		}
		return nil, err
	}

	SQL := "UPDATE products SET product_name = ?, category_id = ?, price = ?, description = ?, " +
		"quantity = ?, product_condition = ?, image = ?, supplier_id = ?, date_of_arrival = ?, expiry_date = ?, updated_at = ? " +
		"WHERE product_id = ?"

	_, err = tx.ExecContext(
		ctx,
		SQL,
		product.ProductName,
		product.CategoryId,
		product.Price,
		product.Description,
		product.Quantity,
		product.ProductCondition,
		product.Image,
		product.SupplierId,
		product.DateOfArrival,
		product.ExpiryDate,
		product.UpdatedAt,
		product.ProductId,
	)

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return product, nil
}

func (p *ProductRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, productId string) error {
	// TODO: Implementasi Anda di sini

	// Cek apakah produk sudah dihapus sebelumnya (is_deleted = 1)
	checkIsDeletedQuery := "SELECT is_deleted FROM products WHERE product_id = ?"
	var isDeleted int

	err := tx.QueryRowContext(ctx, checkIsDeletedQuery, productId).Scan(&isDeleted)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errors.New("product not found")
		}
		return err
	}

	// Jika is_deleted sudah diatur ke 1, maka return pesan kesalahan
	if isDeleted == 1 {
		return errors.New("product is already deleted")
	}

	// Jika belum dihapus, jalankan perintah UPDATE
	updateQuery := "UPDATE products SET is_deleted = 1 WHERE product_id = ?"
	_, err = tx.ExecContext(ctx, updateQuery, productId)
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (p *ProductRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, productId string) (*entity.Product, error) {
	//TODO implement me
	checkQuery := "SELECT product_id FROM products WHERE product_id = ? LIMIT 1 AND is_deleted = 0"
	var existingProductId string

	err := tx.QueryRowContext(ctx, checkQuery, productId).Scan(&existingProductId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("product not found")
		}
		return nil, err
	}

	SQL := "SELECT products.product_id, products.product_name, products.category_id, " +
		"products.price, products.description, products.quantity, products.product_condition, products.image, " +
		"products.supplier_id, products.date_of_arrival, products.expiry_date, products.created_at, products.updated_at " +
		"category.category_name, supplier.supplier_name FROM products where product_id = ? and is_deleted = 0"

	var product entity.Product
	err = tx.QueryRowContext(ctx, SQL, productId).Scan(
		&product.ProductId,
		&product.ProductName,
		&product.CategoryId,
		&product.Price,
		&product.Description,
		&product.Quantity,
		&product.ProductCondition,
		&product.Image,
		&product.SupplierId,
		&product.DateOfArrival,
		&product.ExpiryDate,
		&product.CreatedAt,
		&product.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (p *ProductRepositoryImpl) FindBYName(ctx context.Context, tx *sql.Tx, productName string) (*entity.Product, error) {
	//TODO implement me
	checkQuery := "SELECT product_name FROM products WHERE product_name = ? LIMIT 1 AND is_deleted = 0"
	var existingProductName string

	err := tx.QueryRowContext(ctx, checkQuery, productName).Scan(&existingProductName)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("product not found")
		}
		return nil, err
	}

	SQL := "SELECT products.product_id, products.product_name, products.category_id, " +
		"products.price, products.description, products.quantity, products.product_condition, products.image, " +
		"products.supplier_id, products.date_of_arrival, products.expiry_date, products.created_at, products.updated_at " +
		"category.category_name, supplier.supplier_name FROM products where product_name = ? and is_deleted = 0"

	var product entity.Product
	err = tx.QueryRowContext(ctx, SQL, productName).Scan(
		&product.ProductId,
		&product.ProductName,
		&product.CategoryId,
		&product.Price,
		&product.Description,
		&product.Quantity,
		&product.ProductCondition,
		&product.Image,
		&product.SupplierId,
		&product.DateOfArrival,
		&product.ExpiryDate,
		&product.CreatedAt,
		&product.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (p *ProductRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]*entity.Product, error) {
	//TODO implement me

	checkQuery := "SELECT product_id FROM products WHERE is_deleted = 0"
	var existingProductId string

	err := tx.QueryRowContext(ctx, checkQuery).Scan(&existingProductId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("product not found")
		}
		return nil, err
	}

	SQL := "SELECT products.product_id, products.product_name, products.category_id, " +
		"products.price, products.description, products.quantity, products.product_condition, products.image, " +
		"products.supplier_id, products.date_of_arrival, products.expiry_date, products.created_at, products.updated_at, " +
		"category.category_name, supplier.supplier_name FROM products " +
		"LEFT JOIN category ON products.category_id = category.category_id " +
		"LEFT JOIN supplier ON products.supplier_id = supplier.supplier_id " +
		"WHERE products.is_deleted = 0"

	rows, err := tx.QueryContext(ctx, SQL)
	if err != nil {
		return nil, err
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			panic(err)
		}
	}(rows)

	var products []*entity.Product
	for rows.Next() {
		var product entity.Product
		err := rows.Scan(
			&product.ProductId,
			&product.ProductName,
			&product.CategoryId,
			&product.Price,
			&product.Description,
			&product.Quantity,
			&product.ProductCondition,
			&product.Image,
			&product.SupplierId,
			&product.DateOfArrival,
			&product.ExpiryDate,
			&product.CreatedAt,
			&product.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}
		products = append(products, &product)
	}

	return products, nil
}
