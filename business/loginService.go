package business

import (
	"fmt"
	"gst-billing/commons/constants"
	"gst-billing/models"
	"gst-billing/repositories"
	"gst-billing/utils/authorization"

	"gst-billing/utils/postgres"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type LoginService struct {
	loginServiceRegistry repositories.UserRepository
}

func NewLoginService(loginServiceRegistry repositories.UserRepository) *LoginService {
	return &LoginService{
		loginServiceRegistry: loginServiceRegistry,
	}
}

func (service *LoginService) Login(ctx *gin.Context, request models.LoginRequest) (*models.LoginResponse, error) {

	db := postgres.GetPostGresClient().GormDb
	condition := map[string]interface{}{
		"username": request.Username,
		"password": request.Password,
	}
	userDetails, err := service.loginServiceRegistry.ReadRecordWithCondition(ctx, db, condition)
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return nil, fmt.Errorf(constants.InvalidUsernameOrPasswordError)
		}
		return nil, err
	}

	token, err := authorization.GenerateJWTToken(userDetails.Username)
	if err != nil {
		return nil, err
	}

	loginResponse := &models.LoginResponse{
		Token: token,
	}
	return loginResponse, nil
}
