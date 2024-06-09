package middlewares

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/ricardoalcantara/go-template/internal/domain"
	"github.com/ricardoalcantara/go-template/internal/token"
	"github.com/ricardoalcantara/go-template/internal/utils"
)

func AuthMiddleware() gin.HandlerFunc {
	identity_endpoint := os.Getenv("IDENTITY_ENDPOINT") + "/api/auth/claims"
	return func(c *gin.Context) {

		authToken := getBearerToken(c)

		if len(authToken) == 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, domain.ErrorResponse{Error: "Not authorized"})
			return
		}

		claims, err := token.ValidateTokenWithIdp(authToken, identity_endpoint)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, domain.ErrorResponse{Error: utils.PrintError(err)})
			return
		}

		c.Set("x-user-id", claims.RegisteredClaims.Subject)
		c.Set("x-claims", claims)
		c.Next()

	}
}

func AuthMiddlewareWithRole(roles []string) gin.HandlerFunc {
	identity_endpoint := os.Getenv("IDENTITY_ENDPOINT") + "/api/auth/claims"
	return func(c *gin.Context) {

		authToken := getBearerToken(c)

		if len(authToken) == 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, domain.ErrorResponse{Error: "Not authorized"})
			return
		}

		claims, err := token.ValidateTokenWithIdp(authToken, identity_endpoint)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, domain.ErrorResponse{Error: utils.PrintError(err)})
			return
		}

		if !containsAny(roles, claims.Roles) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, domain.ErrorResponse{Error: "Not in role"})
			return
		}

		c.Set("x-user-id", claims.RegisteredClaims.Subject)
		c.Set("x-claims", claims)
		c.Next()
	}
}

func containsAny(arr1, arr2 []string) bool {
	for _, a := range arr1 {
		for _, b := range arr2 {
			if a == b {
				return true
			}
		}
	}
	return false
}

func getBearerToken(c *gin.Context) string {
	authHeader := c.Request.Header.Get("Authorization")

	t := strings.Split(authHeader, " ")
	if len(t) == 2 {
		return t[1]
	}

	return ""
}
