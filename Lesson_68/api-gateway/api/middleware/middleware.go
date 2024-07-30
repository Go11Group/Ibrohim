package middleware

import (
	"api-gateway/api/rbac"
	"api-gateway/config"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

func Check(c *gin.Context) {
	accessToken := c.GetHeader("Authorization")

	if accessToken == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Authorization header is required",
		})
		return
	}

	token, err := jwt.Parse(accessToken, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.Load().ACCESS_TOKEN), nil
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

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid claims' type",
		})
		return
	}

	userRole, roleOk := claims["role"].(string)
	userID, idOk := claims["user_id"].(string)
	if !roleOk || !idOk {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid token claims",
		})
		return
	}

	_, err = uuid.Parse(userID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid user id",
		})
		return
	}

	c.Set("user_id", userID)

	e, err := rbac.Policy()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "Policy could not be loaded",
		})
		return
	}

	if ok, err := e.Enforce(userRole, c.Request.URL.Path, c.Request.Method); !ok || err != nil {
		msg := fmt.Sprintf("Access denied: %s cannot %s %s",
			userRole, c.Request.Method, c.Request.URL.Path,
		)
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"error": msg,
		})
		return
	}

	c.Next()
}
