package user

import (
	"database/sql"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/zeann3th/ecom/internal/api/auth"
	"github.com/zeann3th/ecom/internal/api/models"
)

type UserHandler struct {
	DB *sql.DB
}

func (u *UserHandler) HandleUserRegister(c echo.Context) error {
	req := new(models.RegisterUserPayload)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid request payload",
		})
	}

	if req.FirstName == "" || req.LastName == "" || req.Email == "" || req.Password == "" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Missing required fields",
		})
	}
	// Check if database have user
	_, err := GetUserByEmail(u.DB, req.Email)
	if err == nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "User already exists",
		})
	}
	// Hash password
	hashedPassword, err := auth.HashPassword(req.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "Password hasing failed",
		})
	}
	// Add new user to database
	user := &models.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  hashedPassword,
	}

	err = CreateUser(u.DB, user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"email":     req.Email,
		"createdAt": time.Now().Format(time.RFC3339),
	})
}

func (u *UserHandler) HandleUserLogin(c echo.Context) error {
	req := new(models.LoginUserPayload)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid request payload",
		})
	}

	if req.Email == "" || req.Password == "" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Missing required fields",
		})
	}
	// Check if database have user
	user, err := GetUserByEmail(u.DB, req.Email)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err,
		})
	}
	// Compare password
	if !auth.ComparePassword(user.Password, req.Password) {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Incorrect password",
		})
	}

	secret := []byte(os.Getenv("JWTSecret"))

	token, err := auth.CreateJWT(secret, user.Id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"msg":   "Login successfully",
		"token": token,
	})
}
