package business

import (
	"fmt"
	"gst-billing/models"
	"gst-billing/repositories"

	"gst-billing/commons/constants"
	"gst-billing/utils/postgres"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type SearchProductService struct {
	productServiceRegistry repositories.ProductRepository
}

func NewSearchProductService(productServiceRegistry repositories.ProductRepository) *SearchProductService {
	return &SearchProductService{
		productServiceRegistry: productServiceRegistry,
	}
}

func (service *SearchProductService) SearchProduct(ctx *gin.Context, request models.SearchProductRequest) (*models.SearchProductResponse, error) {

	db := postgres.GetPostGresClient().GormDb
	condition := map[string]interface{}{
		"product_name": request.ProductName,
		"product_code": request.ProductCode,
	}

	Products, err := service.productServiceRegistry.SearchProduct(ctx, db, condition)
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return nil, fmt.Errorf(constants.ProductNotFoundError)
		}
		return nil, err
	}

	searchProductResponse := &models.SearchProductResponse{
		Products: Products,
	}
	return searchProductResponse, nil
}
