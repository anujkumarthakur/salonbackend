package utils

import (
	"time"
	"userservice/userservice/schema"

	"github.com/golang-jwt/jwt"
)

func GenerateJWTWebToken(user *schema.User) (
	string, *JWTClaims, error) {
	claims := getClaimWithExpiry(user, time.Now().Add(183*24*time.Hour))
	jwtStr, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).
		SignedString([]byte(""))

	return jwtStr, claims, err
}

func GenerateJWTWebTokenWithExpiry(user *schema.User, expiryTime time.Time) (
	string, *JWTClaims, error) {
	claims := getClaimWithExpiry(user, expiryTime)
	jwtStr, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).
		SignedString([]byte("config.SecretKey"))

	return jwtStr, claims, err
}

func getClaimWithExpiry(user *schema.User, expiryTime time.Time) *JWTClaims {
	claims := &JWTClaims{
		StandardClaims: jwt.StandardClaims{
			//Id:        //UniqStr(),
			IssuedAt:  time.Now().Unix(),
			NotBefore: time.Now().Unix(),
			ExpiresAt: expiryTime.Unix(),
		},
		User:      nil,
		UserID:    user.ID,
		FirstName: user.FirstName,
		Email:     user.Email,
	}
	return claims
}

// JWTClaims Custom Claim Struct
type JWTClaims struct {
	jwt.StandardClaims
	// custom claims
	UserID    int          `json:"id"`
	FirstName string       `json:"first_name"`
	Email     string       `json:"email"`
	User      *schema.User `json:"user"`
}
