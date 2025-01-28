package auth

import (
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
)

var secretKey = []byte("dajkvnklsmeiWJE329IREJ3JSEKLDFNKMDSKP'asd]dpsa")

// validate Request and return the username if success and the error if unable to validate
func ValidateRequest(r *http.Request) (string, error) {
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		return "", fmt.Errorf("Missing authorization header")
	}
	tokenString = tokenString[len("Bearer "):]

	username, err := verifyToken(tokenString)
	if err != nil {
		return "", fmt.Errorf("Invalid token")
	}
	return username, nil
}

func CreateToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"exp":      time.Now().Add(time.Hour * 1).Unix(),
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func verifyToken(tokenString string) (string, error) {
	claims := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	// ... error handling

	// do something with decoded claims
	username, ok := claims["username"]

	if !ok {
		return "", fmt.Errorf("cant decode username")
	}

	if err != nil {
		return "", err
	}

	if !token.Valid {
		return "", fmt.Errorf("invalid token")
	}

	return username.(string), nil
}
