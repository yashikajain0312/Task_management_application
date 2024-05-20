package controllers

import (
    "context"
    "net/http"
    "time"
	"fmt"

    "github.com/gin-gonic/gin"
    "github.com/dgrijalva/jwt-go"
	"github.com/asaskevich/govalidator"
    "Tma_backend/ent"
	"Tma_backend/constants"
)

type SignupRequest struct {
    Name      string `json:"username" binding:"required"`
    Email     string `json:"email" binding:"required"`
    Password  string `json:"password" binding:"required"`
}

type SignupResponse struct {
	Token     string  `json:"token"` 
	Message   string   `json:"message"`
}

func SignupHandler(c *gin.Context) {
    var req SignupRequest

    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	// Validating email format.
    if !govalidator.IsEmail(req.Email) {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email format"})
        return
    }

    // Validating password length.
    if len(req.Password) < 8 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Password must be at least 8 characters long"})
        return
    }

    // Creating user record in the database.
    client := c.MustGet("ent").(*ent.Client)
	fmt.Println("client", client)
	fmt.Println("password", req.Password)

    u, err := client.User.
        Create().
        SetUsername(req.Name).
        SetEmail(req.Email).
        SetPassword(req.Password).
        Save(context.Background())
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "Failed to create user"})
        return
    }

    // Generating JWT token.
    token, err := generateJWTToken(int(u.ID))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "Failed to generate token"})
        return
    }

    // Preparing response.
    resp := SignupResponse{
        Token:   token,
        Message: "Signup successful!",
    }

    c.JSON(http.StatusOK, resp)
}

// generateJWTToken generates a JWT token for the given user ID.
func generateJWTToken(userID int) (string, error) {
    token := jwt.New(jwt.SigningMethodHS256)
    claims := token.Claims.(jwt.MapClaims)
    claims["user_id"] = userID
    claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token expiration time (24 hours)

    // Sign the token with the secret key
    tokenString, err := token.SignedString([]byte((constants.JWTSecretKey)))
    if err != nil {
        return "", err
    }

    return tokenString, nil
}
