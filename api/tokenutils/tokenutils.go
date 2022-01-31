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

type Tokens struct {
	AccessToken  string
	RefreshToken string
}

type Claims struct {
	UserId uint `json:"userid"` // User ID
	jwt.StandardClaims
}

// Take ID and generate access token
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

func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

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
