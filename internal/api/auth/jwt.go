package auth

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateJWT(secret []byte, userId int) (string, error) {
	jwtExpirationInSeconds, err := strconv.Atoi(os.Getenv("JWTExpirationInSeconds"))
	if err != nil {
		return "", fmt.Errorf("Invalid time period, failed to parse into int")
	}

	expiration := time.Second * time.Duration(jwtExpirationInSeconds)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":    strconv.Itoa(userId),
		"expiredAt": time.Now().Add(expiration).Unix(),
	})

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", fmt.Errorf("Failed to sign JWT token")
	}
	return tokenString, nil
}
