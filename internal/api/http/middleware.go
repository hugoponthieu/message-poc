package http

import (
	"log"
	"message/infrastructure"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	oidcClient *infrastructure.OidcClient
}

func NewAuthMiddleware(oidcClient *infrastructure.OidcClient) *AuthMiddleware {
	return &AuthMiddleware{
		oidcClient: oidcClient,
	}
}

func (am *AuthMiddleware) Verify() gin.HandlerFunc {
	return func(c *gin.Context) {	
		log.Println(c.Request)

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			log.Println("no header")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			log.Println("wrong format")

			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Bearer token missing"})
			return
		}

		userClaims, err := am.oidcClient.VerifyToken(tokenString)
		if err != nil {
			log.Println("invalid token",err)

			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		// Store typed claims
		c.Set("user", userClaims)

		c.Next()
	}

}
