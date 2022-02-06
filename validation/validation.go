// Package validation provides methods for validating various inputs
package validation

import (
	"errors"
	"net/mail"
	"strings"
	"unicode"

	"github.com/rs401/lg/auth/models"
)

var (
	// ErrEmptyName error for empty name
	ErrEmptyName = errors.New("name cannot be empty")
	// ErrEmptyEmail error for empty email
	ErrEmptyEmail = errors.New("email cannot be empty")
	// ErrEmptyPassword error for empty password
	ErrEmptyPassword = errors.New("password cannot be empty")
	// ErrInvalidEmail error for invalid email
	ErrInvalidEmail = errors.New("email not valid")
	// ErrEmailExists error for email already exists
	ErrEmailExists = errors.New("email already exists")
	// ErrNameExists error for name already exists
	ErrNameExists = errors.New("name already exists")
	// ErrNotFound error for not found
	ErrNotFound = errors.New("user not found")
	// ErrInvalidPassword error for invalid password
	ErrInvalidPassword = errors.New("invalid password, 8-50 characters, one upper, lower, number and special character")

	maxPwLen int = 50
	minPwLen int = 8
)

// IsValidSignUp takes a *SignUpRequest and verifies if the request is valid
func IsValidSignUp(user *models.SignUpRequest) error {
	if IsEmptyString(user.Name) {
		return ErrEmptyName
	}
	if IsEmptyString(user.Email) {
		return ErrEmptyEmail
	}
	if IsEmptyString(string(user.Password)) {
		return ErrEmptyPassword
	}
	if !IsValidEmail(user.Email) {
		return ErrInvalidEmail
	}
	if !IsValidPassword(string(user.Password)) {
		return ErrInvalidPassword
	}

	return nil
}

// IsEmptyString verifies if a string is empty
func IsEmptyString(in string) bool {
	return strings.TrimSpace(in) == ""
}

// IsValidEmail verifies if an email is valid
func IsValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

// IsValidPassword verifies if a password is valid
func IsValidPassword(s string) bool {

	special := false
	number := false
	upper := false
	lower := false

	// Check length
	if len(s) < minPwLen || len(s) > maxPwLen {
		return false
	}

	// Check other requirements
	for _, c := range s {
		if special && number && upper && lower {
			break
		}

		switch {
		case unicode.IsUpper(c):
			upper = true
		case unicode.IsLower(c):
			lower = true
		case unicode.IsNumber(c):
			number = true
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			special = true
		}
	}

	for _, v := range []bool{special, number, upper, lower} {
		if !v {
			return false
		}
	}

	// No errors
	return true
}

// NormalizeEmail normalizes email string
func NormalizeEmail(email string) string {
	return strings.TrimSpace(strings.ToLower(email))
}
