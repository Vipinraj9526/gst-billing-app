package business

import (
	"fmt"
	"gst-billing/commons/constants"
	"gst-billing/models"
	"gst-billing/repositories"
	"strings"

	"gst-billing/utils/postgres"

	"github.com/gin-gonic/gin"
)

type GenerateBillService struct {
	productServiceRegistry      repositories.ProductRepository
	generateBillServiceRegistry repositories.BillRepository
}

func NewGenerateBillService(productServiceRegistry repositories.ProductRepository, generateBillServiceRegistry repositories.BillRepository) *GenerateBillService {
	return &GenerateBillService{
		productServiceRegistry:      productServiceRegistry,
		generateBillServiceRegistry: generateBillServiceRegistry,
	}
}

func (service *GenerateBillService) GenerateBill(ctx *gin.Context, request models.GenerateBillRequest) (*models.Bill, error) {

	db := postgres.GetPostGresClient().GormDb

	// Create a slice to hold product codes
	productCodes := make([]string, 0)
	for _, item := range request.Items {
		productCodes = append(productCodes, item.ProductCode)
	}

	// Query for products using the list of product codes
	condition := map[string]interface{}{
		"product_code": strings.Join(productCodes, "','"),
	}
	products, err := service.productServiceRegistry.ReadRecordsWithCondition(ctx, db, condition)
	if err != nil {
		return nil, err
	}

	// Create a map for easy product lookup
	productMap := make(map[string]models.Product)
	for _, product := range *products {
		productMap[product.ProductCode] = product
	}

	var total, totalTax, subtotal float64
	var billDetails models.Bill

	for _, item := range request.Items {
		product, exists := productMap[item.ProductCode]
		if !exists {
			return nil, fmt.Errorf("product with code %s not found", item.ProductCode)
		}

		price := float64(item.Quantity) * product.ProductPrice
		gst := price * (product.ProductGST / 100)

		subtotal += price
		totalTax += gst
		total += price + gst

		billItem := models.BillItem{
			ProductCode: item.ProductCode,
			Quantity:    item.Quantity,
			Price:       price,
			ProductGST:  product.ProductGST,
			ProductName: product.ProductName,
		}
		billDetails.Items = append(billDetails.Items, billItem)
	}
	billDetails.BillerUserName = ctx.GetString(constants.UserName)
	billDetails.Total = total
	billDetails.TotalTax = totalTax
	billDetails.Subtotal = subtotal
	billDetails.Subtotal = subtotal

	err = service.generateBillServiceRegistry.CreateRecord(ctx, db, &billDetails)
	if err != nil {
		return nil, err
	}

	return &billDetails, nil
}
