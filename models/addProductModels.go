package models

type AddProductRequest struct {
	ProductCode  string  `json:"product_code" validate:"required" example:"P001"`
	ProductName  string  `json:"name" validate:"required" example:"Book"`
	Quantity     uint64  `json:"quantity" validate:"required" exaple:"10"`
	ProductPrice float64 `json:"product_price" validate:"required" example:"100"`
	ProductGST   float64 `json:"product_gst" validate:"required" example:"10"`
	Description  string  `json:"description" validate:"required" example:"This is a book"`
}

type AddProductResponse struct {
	Message string `json:"message" example:"Product added successfully"`
}
