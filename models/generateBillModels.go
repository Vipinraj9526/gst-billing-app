package models

type GenerateBillRequest struct {
	Items []BillItemRequest `json:"items" binding:"required"`
}

type BillItemRequest struct {
	ProductCode string `json:"productCode" validate:"required" example:"P12"`
	Quantity    uint64 `json:"quantity" validate:"required" example:"20"`
}
