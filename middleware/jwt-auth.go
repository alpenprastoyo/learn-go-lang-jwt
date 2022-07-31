package middleware

import (
	"log"
	"net/http"

	"github.com/alpenprastoyo/learn-go-lang-jwt/helper"
	"github.com/alpenprastoyo/learn-go-lang-jwt/service"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

//AuthorizeJWT validate the token user given, return 401 if not valid
func AutorizeJWT(jwtService service.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response := helper.BuildErrorResponse("Failed to process requrest", "No Token Found", nil)
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
		token, err := jwtService.ValidateToken(authHeader)
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Println("Claim[user_id]", claims["user_id"])
			log.Println("Claim[issuer]", claims["issuer"])
		} else {
			log.Println(err)
			response := helper.BuildErrorResponse("Token is not valid", err.Error(), nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		}

	}
}