package business

import (
	"fmt"
	"gst-billing/commons/constants"
	"gst-billing/models"
	"gst-billing/repositories"

	"gst-billing/utils/postgres"

	"github.com/gin-gonic/gin"
)

type AddProductService struct {
	productServiceRegistry repositories.ProductRepository
}

func NewAddProductService(productServiceRegistry repositories.ProductRepository) *AddProductService {
	return &AddProductService{
		productServiceRegistry: productServiceRegistry,
	}
}

func (service *AddProductService) AddProduct(ctx *gin.Context, request models.AddProductRequest) (*models.AddProductResponse, error) {

	db := postgres.GetPostGresClient().GormDb

	err := service.productServiceRegistry.CreateRecord(ctx, db, &models.Product{
		ProductCode:        request.ProductCode,
		ProductName:        request.ProductName,
		Quantity:           request.Quantity,
		ProductPrice:       request.ProductPrice,
		ProductGST:         request.ProductGST,
		ProductDescription: request.Description,
	})
	if err != nil {
		return nil, fmt.Errorf("error creating product: %v", err)
	}

	addProductResponse := &models.AddProductResponse{
		Message: constants.ProductAddedMessage,
	}
	return addProductResponse, nil
}
