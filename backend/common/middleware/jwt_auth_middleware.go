package middleware

import (
	"fmt"
	"github.com/RustamSafiulin/mesh_cloud_computation/backend/common/errors_helper"
	"github.com/pkg/errors"
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("test_secret_key")

// JwtClaims for validation
type JwtClaims struct {
	AccountID string
	jwt.StandardClaims
}

// JwtTokenValidation middleware
func JwtTokenValidation(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {

		tokenString, err := extractTokenFromRequest(req)

		if errors.Cause(err) == errors_helper.ErrParseAuthorizationHeader {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		claims := &JwtClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Error during check access token")
			}

			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		ctx := NewAccountIDContext(req.Context(), claims.AccountID)
		req = req.WithContext(ctx)

		next(w, req)
	})
}

func CreateToken(accountID string) (string, error) {

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &JwtClaims{
		AccountID: accountID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	return tokenString, err
}

func extractTokenFromRequest(r *http.Request) (string, error) {

	authorizationHeader := r.Header.Get("Authorization")
	if authorizationHeader != "" && strings.HasPrefix(authorizationHeader, "Bearer_") {
		return strings.TrimPrefix(authorizationHeader, "Bearer_"), nil
	} else {
		return "", errors_helper.ErrParseAuthorizationHeader
	}
}
