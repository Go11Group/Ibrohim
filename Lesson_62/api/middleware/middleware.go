package middleware

import (
	"fmt"
	"log"
	"net/http"
	"redis-crud/api/rbac"

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
		log.Print(err)
		return
	}

	if !token.Valid {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid token provided",
		})
		return
	}

	claims := token.Claims.(jwt.MapClaims)
	userRole, ok := claims["role"].(string)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid token claims",
		})
		return
	}

	e, err := rbac.Policy()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "Policy could not be loaded",
		})
		log.Print(err)
		return
	}

	if ok, err := e.Enforce(userRole, c.Request.URL.Path, c.Request.Method); !ok || err != nil {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"error": fmt.Sprintf("Access denied: %s cannot %s %s", userRole, c.Request.Method, c.Request.URL.Path),
		})
		log.Print(ok, err)
		return
	}
	log.Println(userRole, c.Request.URL.Path, c.Request.Method)

	c.Next()
}
