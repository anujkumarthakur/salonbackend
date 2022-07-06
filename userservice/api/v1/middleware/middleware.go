package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"userservice/userservice/errors"
	"userservice/userservice/respond"

	"github.com/golang-jwt/jwt"
)

// SPACE space string literal
const SPACE = " "

func AuthRequired(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		tokenStr := r.Header.Get("Authorization")
		if strings.TrimSpace(tokenStr) == "" {
			respond.Fail(w, errors.Unauthorized("No credentials sent"))
			return
		}
		validateErr := ValidateToken(tokenStr)
		if validateErr != nil {
			respond.Fail(w, errors.Unauthorized("Invalid token"))
			return
		}
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

// parseJWTToken parse the token and verify the token with signing secret key
func parseJWTToken(authToken string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(authToken, &JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			// Validate expected algorithm.
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			if token.Header["alg"] != "HS256" || token.Header["typ"] != "JWT" {
				return nil, fmt.Errorf("Unexpected signing algorithm: %s or type: %s",
					token.Header["alg"], token.Header["typ"])
			}
			return []byte(jwtKey), nil
		},
	)
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	if err := token.Claims.Valid(); err != nil {
		return nil, err
	}

	return token, err
}
