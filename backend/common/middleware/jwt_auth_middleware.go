package middleware

import (
	"errors"
	"fmt"
	"github.com/RustamSafiulin/mesh_cloud_computation/backend/common/errors_helper"
	"net/http"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
)

var JwtKey = []byte("test_secret_key")

// JwtClaims for validation
type JwtClaims struct {
	Username string
	jwt.StandardClaims
}

// JwtTokenValidation middleware
func JwtTokenValidation(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {

		tokenString, err := extractTokenFromRequest(req)

		var appErr *errors_helper.ApplicationError
		if err != nil && errors.As(err, &appErr) && appErr.Code() == errors_helper.ErrParseAuthorizationHeader {
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

		next(w, req)
	})
}

func extractTokenFromRequest(r *http.Request) (string, error) {

	authorizationHeader := r.Header.Get("Authorization")
	if authorizationHeader != "" {
		bearerToken := strings.Split(authorizationHeader, " ")

		if len(bearerToken) == 2 {
			return bearerToken[1], nil
		} else {
			return "", errors_helper.NewApplicationError(errors_helper.ErrParseAuthorizationHeader)
		}
	} else {
		return "", errors_helper.NewApplicationError(errors_helper.ErrParseAuthorizationHeader)
	}
}
