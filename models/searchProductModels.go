package models

type SearchProductRequest struct {
	ProductName string `json:"productName" validate:"required" example:"Book"`
	ProductCode string `json:"productCode" validate:"required" example:"P001"`
}

type SearchProductResponse struct {
	Products []Product `json:"products"`
}
