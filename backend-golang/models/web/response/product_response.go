package response

import "time"

type ProductResponse struct {
	ProductId     string    `json:"product_id"`
	ProductName   string    `json:"product_name"`
	CategoryId    string    `json:"category_id"`
	Price         int       `json:"price"`
	Description   *string   `json:"description"`
	Quantity      int       `json:"quantity"`
	Condition     string    `json:"condition"`
	Image         *string   `json:"image"`
	SupplierId    string    `json:"supplier_id"`
	DateOfArrival time.Time `json:"date_of_arrival"`
	ExpiryDate    time.Time `json:"expiry_date"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
