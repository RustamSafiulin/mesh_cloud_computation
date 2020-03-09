package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
)

var JwtKey = []byte("test_secret_key")

// ErrParseAuthorizationHeader describes error when parse auth header
var ErrParseAuthorizationHeader = errors.New("Error during parse authorization header")

// JwtClaims for validation
type JwtClaims struct {
	Username string
	jwt.StandardClaims
}

// JwtTokenValidation middleware
func JwtTokenValidation(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {

		tokenString, err := extractTokenFromRequest(req)
		if err == ErrParseAuthorizationHeader {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		claims := &JwtClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Error during check access token")
			}

			return JwtKey, nil
		})

		if err != nil || !token.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
	})
}

func extractTokenFromRequest(r *http.Request) (string, error) {

	authorizationHeader := r.Header.Get("Authorization")
	if authorizationHeader != "" {
		bearerToken := strings.Split(authorizationHeader, " ")

		if len(bearerToken) == 2 {
			return bearerToken[1], nil
		} else {
			return "", ErrParseAuthorizationHeader
		}
	} else {
		return "", ErrParseAuthorizationHeader
	}
}
