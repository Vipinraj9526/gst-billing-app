package handlers

import (
	"gst-billing/business"
	"gst-billing/commons/constants"
	"gst-billing/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type LoginController struct {
	service *business.LoginService
}

func NewLoginController(service *business.LoginService) *LoginController {
	return &LoginController{
		service: service,
	}
}

// @Summary Login user
// @Description Login user
// @Tags Login
// @Accept json
// @Produce json
// @Param login body models.LoginRequest true "Login request"
// @Success 200 {object} models.LoginResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /v1/login [post]
func (controller *LoginController) LoginHandler(ctx *gin.Context) {
	var request models.LoginRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}

	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}

	token, err := controller.service.Login(ctx, request)
	if err != nil {
		if strings.Contains(err.Error(), constants.InvalidUsernameOrPasswordError) {
			ctx.JSON(http.StatusUnauthorized, models.ErrorResponse{Error: err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, token)
}
