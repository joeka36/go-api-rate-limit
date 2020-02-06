package transport

import (
	"go-rate-limit/errs"
	"github.com/dgrijalva/jwt-go"
)

// ValidateUserEmail extracts the JWT and check if the user email is valid
func ValidateUserEmail(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("SECRET"), nil
	})

	if err != nil {
		return "", errs.JWTValidationError
	}

	c := token.Claims.(jwt.MapClaims)

	if c["email"] == 0 {
		return "", errs.EmailNotInJWT
	}

	_, ok := APIUsers[c["email"].(string)]
	if !ok {
		return "", errs.UnvalidatedEmail
	}

	return c["email"].(string), nil
}

// ValidateAPIKey validates the given API key against the user data
func ValidateAPIKey(email string, apiKey string) error {
	if apiKey == "" {
		return errs.NoAPIKey
	}

	if APIUsers[email].Key != apiKey {
		return errs.IncorrectAPIKey
	}

	return nil
} 
