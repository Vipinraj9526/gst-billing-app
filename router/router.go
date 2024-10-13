package router

import (
	"gst-billing/business"
	"gst-billing/commons/constants"
	"gst-billing/docs"
	"gst-billing/handlers"
	"gst-billing/repositories"
	"gst-billing/utils/authorization"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func init() {
	gin.SetMode(gin.ReleaseMode)
}

// GetRouter returns a new Gin engine configured with middleware and routes
func GetRouter(middlewares ...gin.HandlerFunc) *gin.Engine {
	router := gin.New()
	router.Use(middlewares...)
	router.Use(gin.Recovery())

	// Swagger documentation
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	// Swagger endpoint
	router.GET(constants.SwaggerRoute, ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Repositories
	userRepository := repositories.GetUserRepository()
	productRepository := repositories.GetProductRepository()
	billRepository := repositories.GetBillRepository()

	// authrnticatiom
	jwtMiddeleWare := authorization.JWTAuthMiddleware()

	loginService := business.NewLoginService(userRepository)
	loginController := handlers.NewLoginController(loginService)

	addProductService := business.NewAddProductService(productRepository)
	addProductController := handlers.NewAddProductController(addProductService)

	searchProductService := business.NewSearchProductService(productRepository)
	searchProductController := handlers.NewSearchProductController(searchProductService)

	generateBillService := business.NewGenerateBillService(productRepository, billRepository)
	generateBillController := handlers.NewGenerateBillController(generateBillService)

	v1Routes := router.Group(constants.V1Routes)
	{
		// Health Check
		v1Routes.GET(constants.HealthCheck, func(c *gin.Context) {
			response := map[string]string{
				"message": "API is up and running",
			}
			c.JSON(http.StatusOK, response)
		})

		// Define routes
		v1Routes.POST(constants.Login, loginController.LoginHandler)
		v1Routes.POST(constants.AddProduct, jwtMiddeleWare, addProductController.AddProductHandler)
		v1Routes.POST(constants.SeachProduct, jwtMiddeleWare, searchProductController.SearchProductHandler)
		v1Routes.POST(constants.GenerateBill, jwtMiddeleWare, generateBillController.GenerateBillHandler)
	}

	return router
}
