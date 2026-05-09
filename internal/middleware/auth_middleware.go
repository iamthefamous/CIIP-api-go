package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

func AdminOnly(appJWTSecret, supabaseJWTSecret string, getProfileRoleByID func(userID string) (string, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")

		if auth == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
			c.Abort()
			return
		}

		parts := strings.Split(auth, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid authorization header"})
			c.Abort()
			return
		}

		token, err := jwt.ParseWithClaims(
			parts[1],
			&Claims{},
			func(token *jwt.Token) (interface{}, error) {
				return []byte(appJWTSecret), nil
			},
			jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}),
		)

		if err == nil && token.Valid {
			claims, ok := token.Claims.(*Claims)
			if ok && claims.Role == "admin" {
				c.Next()
				return
			}
			c.JSON(http.StatusForbidden, gin.H{"error": "admin only"})
			c.Abort()
			return
		}

		if supabaseJWTSecret == "" || getProfileRoleByID == nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}

		supabaseToken, err := jwt.Parse(
			parts[1],
			func(token *jwt.Token) (interface{}, error) {
				return []byte(supabaseJWTSecret), nil
			},
			jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}),
		)
		if err != nil || !supabaseToken.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}

		claims, ok := supabaseToken.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token claims"})
			c.Abort()
			return
		}

		sub, ok := claims["sub"].(string)
		if !ok || sub == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token claims"})
			c.Abort()
			return
		}

		role, err := getProfileRoleByID(sub)
		if err != nil || role != "admin" {
			c.JSON(http.StatusForbidden, gin.H{"error": "admin only"})
			c.Abort()
			return
		}

		c.Next()
	}
}
