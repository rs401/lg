// Package tokenutils provides utilities for JWTs
package tokenutils

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"

	"github.com/google/uuid"
)

// Tokens holds an access token and a refresh token
type Tokens struct {
	AccessToken  string
	RefreshToken string
}

// Claims holds a UserId and standard claims
type Claims struct {
	UserId uint `json:"userid"` // User ID
	jwt.StandardClaims
}

// CreateToken takes an ID and generates a Tokens
func CreateToken(userid uint) (*Tokens, error) {
	ts := &Tokens{}

	var err error
	//Create Access Token
	atClaims := Claims{}
	atClaims.UserId = userid
	atClaims.Id = uuid.NewString()
	atClaims.IssuedAt = time.Now().Unix()
	atClaims.ExpiresAt = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	ts.AccessToken, err = at.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return nil, err
	}
	//Create Refresh Token
	rtClaims := Claims{}
	rtClaims.UserId = userid
	rtClaims.Id = uuid.NewString()
	rtClaims.IssuedAt = time.Now().Unix()
	rtClaims.ExpiresAt = time.Now().Add(time.Hour * 24 * 7).Unix()
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	ts.RefreshToken, err = rt.SignedString([]byte(os.Getenv("REFRESH_SECRET")))
	if err != nil {
		return nil, err
	}
	return ts, nil
}

// ExtractToken extracts an access token from request header
func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

// VerifyToken verifies a token is legit
func VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

// VerifyRefreshToken verifies a refresh token is legit
func VerifyRefreshToken(r *http.Request) (*jwt.Token, error) {
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("REFRESH_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}
