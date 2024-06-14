package jwt

import (
	"go_playground/internal/core/entity"
	"time"

	"github.com/golang-jwt/jwt"
)

type UserClaims struct {
	jwt.StandardClaims
	Username string `json:"username"`
}

func GenerateAccessToken(customer entity.Customers) (string, error) {
	claims := UserClaims{Username: customer.Username, StandardClaims: jwt.StandardClaims{
		Issuer:    "Synapsis",
		ExpiresAt: time.Now().Add(2 * time.Hour).Unix(),
	}}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return accessToken.SignedString([]byte("SECRET_KEY"))
}

func GenerateRefreshToken() (string, error) {
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
	})
	return refreshToken.SignedString([]byte("SECRET_KEY"))
}

func ParseAccessToken(accessToken string) *UserClaims {
	parsedAccessToken, _ := jwt.ParseWithClaims(accessToken, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("SECRET_KEY"), nil
	})

	return parsedAccessToken.Claims.(*UserClaims)
}
