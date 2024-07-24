package middleware

import (
	"api-gateway/api/rbac"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

const signingkey = "ssecca"

func Check(c *gin.Context) {
	accessToken := c.GetHeader("Authorization")

	if accessToken == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Authorization header is required",
		})
		return
	}

	token, err := jwt.Parse(accessToken, func(t *jwt.Token) (interface{}, error) {
		return []byte(signingkey), nil
	})

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Token could not be parsed",
		})
		return
	}

	if !token.Valid {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid token provided",
		})
		return
	}

	claims := token.Claims.(jwt.MapClaims)
	userRole := claims["role"].(string)

	e, err := rbac.Policy()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "Policy could not be loaded",
		})
		return
	}

	if ok, err := e.Enforce(userRole, c.Request.URL.Path, c.Request.Method); !ok || err != nil {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"error": "Access denied",
		})
		return
	}

	c.Next()
}
