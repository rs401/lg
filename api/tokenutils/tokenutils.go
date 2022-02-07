// Package tokenutils provides utilities for JWTs
package tokenutils

import (
	"errors"
	"fmt"
	"log"
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

func RefreshAccessToken(refreshToken *jwt.Token) (string, error) {
	if claims, ok := refreshToken.Claims.(*Claims); ok && refreshToken.Valid {
		userid := claims.UserId
		atClaims := Claims{}
		atClaims.UserId = userid
		atClaims.Id = uuid.NewString()
		atClaims.IssuedAt = time.Now().Unix()
		atClaims.ExpiresAt = time.Now().Add(time.Minute * 15).Unix()
		at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
		return at.SignedString([]byte(os.Getenv("JWT_SECRET")))
	}
	// Something is wrong, how did they get this far?
	log.Printf("Bad Refresh Token: %v\n", refreshToken)
	return "", errors.New("bad refresh token")
}

// ExtractToken extracts an access token from request header
func ExtractToken(r *http.Request) string {
	bearerToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearerToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func ExtractRefreshToken(r *http.Request) string {
	cookie, err := r.Cookie("refresh_token")
	if err != nil {
		return ""
	}
	return cookie.Value
}

// VerifyToken verifies a token is legit
func VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := ExtractToken(r)
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
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
	tokenString := ExtractRefreshToken(r)
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
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

func StillValid(token *jwt.Token) bool {
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims.StandardClaims.ExpiresAt >= time.Now().Unix()
	}
	// Not ok
	return false
}
