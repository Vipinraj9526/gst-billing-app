// your_project/authorization/authorization.go
package authorization

import (
	// Adjust the import path as necessary

	"fmt"
	"gst-billing/commons/constants"
	"gst-billing/utils/configs"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Claims structure for JWT
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// GenerateJWTToken generates a JWT token for a user
func GenerateJWTToken(username string) (string, error) {
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // Token expiration time
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	applicationConfig, err := configs.LoadConfig("configs/application.yml")
	if err != nil {
		return "", fmt.Errorf(constants.GetApplicationConfigError, err)
	}
	var jwtKey = []byte(applicationConfig.Application.JwtSecretKey)
	signedToken, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

// JWTAuthMiddleware is a Gin middleware to protect routes
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the token from the Authorization header
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing or malformed"})
			c.Abort()
			return
		}

		// Trim the "Bearer " prefix
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		// Parse the token
		claims := &Claims{}
		applicationConfig, err := configs.LoadConfig("configs/application.yml")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		jwtKey := []byte(applicationConfig.Application.JwtSecretKey)
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Set(constants.UserName, claims.Username)
		c.Next()
	}
}
