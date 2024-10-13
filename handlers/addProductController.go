package handlers

import (
	"gst-billing/business"
	"gst-billing/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AddProductController struct {
	service *business.AddProductService
}

func NewAddProductController(service *business.AddProductService) *AddProductController {
	return &AddProductController{
		service: service,
	}
}

// @Summary Add product
// @Description Add product
// @Tags Products
// @Accept json
// @Produce json
// @security ApiKeyAuth
// @Param addProductRequest body models.AddProductRequest true "Add product request"
// @Success 200 {object} models.AddProductResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /v1/login/products/add [post]
func (controller *AddProductController) AddProductHandler(ctx *gin.Context) {
	var request models.AddProductRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}

	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}

	addProductResponse, err := controller.service.AddProduct(ctx, request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, addProductResponse)
}
