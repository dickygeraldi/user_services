package authentication

import (
	"creart_new/models"
	"net/http"
	"os"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {
	tokenString := c.Request.Header.Get("Authorization")
	tk := &models.Token{}

	tkn, err := jwt.ParseWithClaims(tokenString, tk, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("KEY_JWT")), nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			result := gin.H{
				"Message": "Invalid Signature",
			}
			c.JSON(http.StatusUnauthorized, result)
		}
		result := gin.H{
			"Message": "Bad Request",
		}
		c.JSON(http.StatusBadRequest, result)
	}

	if !tkn.Valid {
		result := gin.H{
			"Message": "Invalid Token",
		}
		c.JSON(http.StatusUnauthorized, result)
	}
}
