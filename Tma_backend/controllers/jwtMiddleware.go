package controllers

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/dgrijalva/jwt-go"
    "Tma_backend/constants"
)

// JWTMiddleware is a middleware function to validate JWT tokens.
func JWTMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        // Extracting token from request header.
        tokenString := c.GetHeader("Authorization")
        if tokenString == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Token not provided"})
            return
        }

        // Parsing and validating token.
        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            return []byte(constants.JWTSecretKey), nil
        })
        if err != nil || !token.Valid {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            return
        }

        // Setting the user ID extracted from the token in the context.
        claims := token.Claims.(jwt.MapClaims)
        userID := claims["user_id"].(float64) 
        c.Set("user_id", int(userID))

        // Continue processing.
        c.Next()
    }
}
