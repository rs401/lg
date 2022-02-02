package tokenutils

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error loading environment variables: %v\n", err)
	}
}

func isValidToken(tokenString string) bool {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return false
	}
	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return true
	}
	return false
}

func isValidRefreshToken(tokenString string) bool {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("REFRESH_SECRET")), nil
	})
	if err != nil {
		return false
	}
	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return true
	}
	return false
}

func TestCreateToken(t *testing.T) {
	type args struct {
		userid uint
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "Test 1",
			args:    args{userid: 1},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateToken(tt.args.userid)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !isValidToken(got.AccessToken) {
				t.Errorf("CreateToken() AccessToken not valid: %v", got.AccessToken)
			}
			if !isValidRefreshToken(got.RefreshToken) {
				t.Errorf("CreateToken() RefreshToken not valid: %v", got.RefreshToken)
			}
		})
	}
}

func TestExtractToken(t *testing.T) {
	type args struct {
		r *http.Request
	}
	tokenString := "authorization test"
	// req := &http.Request{}
	req := httptest.NewRequest(http.MethodGet, "/api/", nil)
	req.Header.Set("Authorization", tokenString)
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Extract test token",
			args: args{r: req},
			want: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExtractToken(tt.args.r); got != tt.want {
				t.Errorf("ExtractToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVerifyToken(t *testing.T) {
	type args struct {
		r *http.Request
	}
	tokens, err := CreateToken(2)
	assert.NoError(t, err)

	tokenString := tokens.AccessToken
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	assert.NoError(t, err)
	req := httptest.NewRequest(http.MethodGet, "/api/", nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", tokenString))
	tests := []struct {
		name    string
		args    args
		want    *jwt.Token
		wantErr bool
	}{
		{
			name:    "TestVerifyToken",
			args:    args{r: req},
			want:    token,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := VerifyToken(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("VerifyToken() error = %v, wantErr %v\n\tGOT: %v\n", err, tt.wantErr, got)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("VerifyToken() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestVerifyRefreshToken(t *testing.T) {
	type args struct {
		r *http.Request
	}
	tokens, err := CreateToken(2)
	assert.NoError(t, err)

	tokenString := tokens.RefreshToken
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("REFRESH_SECRET")), nil
	})
	assert.NoError(t, err)
	req := httptest.NewRequest(http.MethodGet, "/api/", nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", tokenString))
	tests := []struct {
		name    string
		args    args
		want    *jwt.Token
		wantErr bool
	}{
		{
			name:    "TestVerifyRefreshToken",
			args:    args{r: req},
			want:    token,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := VerifyRefreshToken(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("VerifyToken() error = %v, wantErr %v\n\tGOT: %v\n", err, tt.wantErr, got)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("VerifyToken() = %v, want %v", got, tt.want)
			}
		})
	}
}
