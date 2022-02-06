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
	// ErrEmptyName
	ErrEmptyName = errors.New("name cannot be empty")
	// ErrEmptyEmail
	ErrEmptyEmail = errors.New("email cannot be empty")
	// ErrEmptyPassword
	ErrEmptyPassword = errors.New("password cannot be empty")
	// ErrInvalidEmail
	ErrInvalidEmail = errors.New("email not valid")
	// ErrEmailExists
	ErrEmailExists = errors.New("email already exists")
	// ErrNameExists
	ErrNameExists = errors.New("name already exists")
	// ErrNotFound
	ErrNotFound = errors.New("user not found")
	// ErrInvalidPassword
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
