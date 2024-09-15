package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/zeann3th/ecom/internal/api/services/user"
	"github.com/zeann3th/ecom/internal/db"
)

func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenStr := getTokenFromRequest(c)
		token, err := validateToken(tokenStr)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"error": "Invalid token",
			})
		}

		if !token.Valid {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"error": "Invalid token",
			})
		}

		claims := token.Claims.(jwt.MapClaims)
		str := claims["userId"].(string)

		id, _ := strconv.Atoi(str)

		instdb, err := db.ConnectStorage(os.Getenv("DB_DRIVER"), os.Getenv("DB_CONN"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error": "Database connection failed",
			})
		}

		u, err := user.GetUserById(instdb, id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error": "User does not exist",
			})
		}

		c.Set("user", u)

		return next(c)
	}
}

func getTokenFromRequest(c echo.Context) string {
	tokenAuth := c.Request().Header.Get("Authorization")

	if tokenAuth != "" {
		return tokenAuth
	}

	return ""
}

func validateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(os.Getenv("JWTSecret")), nil
	})
}
