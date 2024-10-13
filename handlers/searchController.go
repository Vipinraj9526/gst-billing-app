package handlers

import (
	"gst-billing/business"
	"gst-billing/models"
	"net/http"

	"gst-billing/commons/constants"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type SearchProductController struct {
	service *business.SearchProductService
}

func NewSearchProductController(service *business.SearchProductService) *SearchProductController {
	return &SearchProductController{
		service: service,
	}
}

// @Summary Search product
// @Description Search product
// @Tags Products
// @Accept json
// @security ApiKeyAuth
// @Accept json
// @Produce json
// @Param searchProductRequest body models.SearchProductRequest true "Search product request"
// @Success 200 {object} models.SearchProductResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /v1/login/products/search [post]
func (controller *SearchProductController) SearchProductHandler(ctx *gin.Context) {
	var request models.SearchProductRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}

	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}

	products, err := controller.service.SearchProduct(ctx, request)
	if err != nil {
		if err.Error() == constants.ProductNotFoundError {
			ctx.JSON(http.StatusNotFound, models.ErrorResponse{Error: err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, products)
}
