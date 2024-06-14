package jwt

import (
	"errors"
	"go_playground/internal/core/entity"
	"time"

	"github.com/golang-jwt/jwt"
)

type UserClaims struct {
	jwt.StandardClaims
	Id       int    `json:"id"`
	Username string `json:"username"`
}

func GenerateAccessToken(customer entity.Customers) (string, error) {
	claims := UserClaims{Id: int(customer.Id), Username: customer.Username, StandardClaims: jwt.StandardClaims{
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

func ParseAccessToken(accessToken string) (*UserClaims, error) {
	to, err := jwt.Parse(accessToken, func(t *jwt.Token) (interface{}, error) {

		if t.Method != jwt.SigningMethodHS256 {
			return nil, errors.New("invalid signing method")
		}
		return []byte("SECRET_KEY"), nil
	})
	if err != nil {
		return nil, err
	}
	if !to.Valid {
		return nil, errors.New("invalid token")
	}
	claims, ok := to.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid claims")
	}

	exp, _ := claims["exp"].(int64)
	id, _ := claims["id"].(int)
	if !ok {
		return nil, errors.New("invalid id claim")
	}
	username, _ := claims["username"].(string)

	user := &UserClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: exp,
		},
		Username: username,
		Id:       id,
	}
	return user, nil

}
