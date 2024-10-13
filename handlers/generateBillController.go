package handlers

import (
	"gst-billing/business"
	"gst-billing/models"
	"net/http"

	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"
)

type GenerateBillController struct {
	service *business.GenerateBillService
}

func NewGenerateBillController(service *business.GenerateBillService) *GenerateBillController {
	return &GenerateBillController{
		service: service,
	}
}

// @Summary Generate bill
// @Description Generate bill
// @Tags GenerateBill
// @Accept json
// @Produce json
// @security ApiKeyAuth
// @Param billRequest body models.GenerateBillRequest true "GenerateBill request"
// @Success 200 {object} models.Bill
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /v1/login/bill/generate [post]
func (controller *GenerateBillController) GenerateBillHandler(ctx *gin.Context) {
	var request models.GenerateBillRequest
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

	billDetails, err := controller.service.GenerateBill(ctx, request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, billDetails)
}
