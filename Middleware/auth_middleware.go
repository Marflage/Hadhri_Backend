package middleware

import (
	"errors"
	"fmt"
	constants "hadhri/Constants"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.GetHeader(constants.AuthHeader)

		if tokenStr == "" {
			// TODO: Log.
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing authorization header."})
			return
		}

		tokenStr = strings.TrimPrefix(tokenStr, constants.BearerTokenPrefix)

		claims := jwt.RegisteredClaims{}

		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (any, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			return constants.Jwtkey, nil
		})

		if err != nil || !token.Valid {
			if errors.Is(err, jwt.ErrTokenExpired) {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token expired."})
			}
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token."})
			return
		}

		c.Set("userEmail", claims.Subject)

		c.Next()
	}
}
