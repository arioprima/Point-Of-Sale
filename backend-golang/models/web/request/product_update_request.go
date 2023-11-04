package request

import "time"

type ProductUpdateRequest struct {
	ProductId        string    `json:"product_id"`
	ProductName      string    `json:"product_name" validate:"required"`
	CategoryId       string    `json:"category_id"`
	Price            int       `json:"price"`
	Description      *string   `json:"description"`
	Quantity         int       `json:"quantity"`
	ProductCondition string    `json:"product_condition"`
	Image            *string   `json:"image"`
	SupplierId       string    `json:"supplier_id"`
	DateOfArrival    time.Time `json:"date_of_arrival"`
	ExpiryDate       time.Time `json:"expiry_date"`
}
