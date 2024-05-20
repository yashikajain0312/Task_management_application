package controllers

import (
    "context"
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
    "Tma_backend/ent"
    "Tma_backend/ent/user"
    "github.com/dgrijalva/jwt-go"
	"Tma_backend/constants"
)

type LoginRequest struct {
    Email    string `json:"email" binding:"required"`
    Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
    Token   string `json:"token"`
    Message string `json:"message"`
}

func LoginHandler(c *gin.Context) {
    var loginReq LoginRequest

    if err := c.BindJSON(&loginReq); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Querying the user by email.
	client := c.MustGet("ent").(*ent.Client)
    u, err := client.User.Query().
        Where(user.EmailEQ(loginReq.Email)).
        Only(context.Background())
    if err != nil {
        if ent.IsNotFound(err) {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
            return
        }
    
        c.JSON(http.StatusInternalServerError, gin.H{"message": "Database error", "error": err.Error()})
        return
    }

    // Validate password
    if u.Password != loginReq.Password {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    // Generate JWT token
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "email": loginReq.Email,
        "exp":   time.Now().Add(time.Hour * 24).Unix(), // Token expiration time (1 day)
    })

    tokenString, err := token.SignedString([]byte((constants.JWTSecretKey)))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
        return
    }

    // Return token in response
    loginResp := LoginResponse{
        Token:   tokenString,
        Message: "Login Successfully!",
    }

    c.JSON(http.StatusOK, loginResp)
}
