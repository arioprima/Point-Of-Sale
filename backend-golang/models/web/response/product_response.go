package response

type ProductResponse struct {
	ProductId     string  `json:"product_id"`
	ProductName   string  `json:"product_name"`
	CategoryId    string  `json:"category_id"`
	Price         int     `json:"price"`
	Description   *string `json:"description"`
	Quantity      int     `json:"quantity"`
	Condition     string  `json:"condition"`
	Image         *string `json:"image"`
	SupplierId    string  `json:"supplier_id"`
	DateOfArrival string  `json:"date_of_arrival"`
	ExpiryDate    string  `json:"expiry_date"`
	CreatedAt     string  `json:"created_at"`
	UpdatedAt     string  `json:"updated_at"`
}
