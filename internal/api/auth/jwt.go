package auth

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/zeann3th/ecom/internal/config"
)

func CreateJWT(secret []byte, userId int) (string, error) {
	jwtExpirationInSeconds, err := strconv.Atoi(config.Env["JWTExpirationInSeconds"])
	if err != nil {
		return "", err
	}

	expiration := time.Second * time.Duration(jwtExpirationInSeconds)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":    strconv.Itoa(userId),
		"expiredAt": time.Now().Add(expiration).Unix(),
	})

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", nil
	}
	return tokenString, nil
}
