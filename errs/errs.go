package errs

import (
	"errors"
)

var (
	JWTValidationError = errors.New("JWT token is not validated")
	EmailNotInJWT = errors.New("Could not extract email from JWT token")
	UnvalidatedEmail = errors.New("Email not registered")
)