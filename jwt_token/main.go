package main

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Secret key for signing tokens
var secretKey = []byte("your_secret_key")

// CustomClaims structure to include additional data in JWT
type CustomClaims struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

// Function to create a token (JWT or Refresh token)
func createToken(userID string, duration time.Duration) (string, error) {
	// Define token expiration time
	expirationTime := time.Now().Add(duration)

	// Create the JWT claims, which includes the user ID and expiry time
	claims := &CustomClaims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Set custom header
	token.Header["typ"] = "JWT"
	token.Header["alg"] = "HS256"
	token.Header["kid"] = "your_key_id"

	// Sign the token with the secret key
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// Function to parse and validate token
func parseToken(tokenString string) (*CustomClaims, error) {
	claims := &CustomClaims{}

	// Parse the token
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		// Verify the signing algorithm
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// Verify the key ID
		if kid, ok := token.Header["kid"].(string); !ok || kid != "your_key_id" {
			return nil, fmt.Errorf("unexpected key ID: %v", kid)
		}

		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	// Check if token is valid
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}

func main() {
	userID := "123456"

	// Create JWT token
	jwtToken, err := createToken(userID, 15*time.Minute)
	if err != nil {
		fmt.Println("Error creating JWT:", err)
		return
	}

	// Create Refresh token
	refreshToken, err := createToken(userID, 7*24*time.Hour)
	if err != nil {
		fmt.Println("Error creating Refresh Token:", err)
		return
	}

	fmt.Println("JWT Token:", jwtToken)
	fmt.Println("Refresh Token:", refreshToken)

	// Parse and validate JWT token
	parsedClaims, err := parseToken(jwtToken)
	if err != nil {
		fmt.Println("Error parsing JWT:", err)
		return
	}

	fmt.Printf("Parsed JWT Claims: %+v\n", parsedClaims)

	// Parse and validate Refresh token
	parsedRefreshClaims, err := parseToken(refreshToken)
	if err != nil {
		fmt.Println("Error parsing Refresh Token:", err)
		return
	}

	fmt.Printf("Parsed Refresh Token Claims: %+v\n", parsedRefreshClaims)
}
