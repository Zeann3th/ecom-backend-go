package user

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/zeann3th/ecom/internal/api/auth"
	"github.com/zeann3th/ecom/internal/api/models"
	"github.com/zeann3th/ecom/internal/config"
	"github.com/zeann3th/ecom/internal/db"
)

func UserRegisterHandler(c echo.Context) error {
	req := new(models.RegisterUserPayload)
	if err := c.Bind(req); err != nil {
		return c.JSONBlob(http.StatusBadRequest, []byte(`{"error": "Invalid request payload"}`))
	}

	if req.FirstName == "" || req.LastName == "" || req.Email == "" || req.Password == "" {
		return c.JSONBlob(http.StatusBadRequest, []byte(`{"error": "Missing required fields"}`))
	}
	// DB instance
	instdb, err := db.ConnectStorage(config.Env["DB_DRIVER"], config.Env["DB_CONN"])
	if err != nil {
		return c.JSONBlob(http.StatusInternalServerError, []byte(`{"error": "Internal server error"}`))
	}

	// Check if database have user
	_, err = GetUserByEmail(instdb, req.Email)
	if err == nil {
		return c.JSONBlob(http.StatusBadRequest, []byte(`{"error": "User already exists"}`))
	}

	// Hash password
	hashedPassword, err := auth.HashPassword(req.Password)
	if err != nil {
		return err
	}

	// Add new user to database
	user := &models.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  hashedPassword,
	}
	err = CreateUser(instdb, user)

	return c.JSONBlob(http.StatusOK, []byte(
		fmt.Sprintf(`{
      "email": %q,
      "createdAt": %q
    }`, req.Email, time.Now().Format(time.RFC3339))))
}

func UserLoginHandler(c echo.Context) error {
	req := new(models.LoginUserPayload)
	if err := c.Bind(req); err != nil {
		return c.JSONBlob(http.StatusBadRequest, []byte(`{"error": "Invalid request payload"}`))
	}

	if req.Email == "" || req.Password == "" {
		return c.JSONBlob(http.StatusBadRequest, []byte(`{"error": "Missing required fields"}`))
	}
	// DB instance
	instdb, err := db.ConnectStorage(config.Env["DB_DRIVER"], config.Env["DB_CONN"])
	if err != nil {
		return c.JSONBlob(http.StatusInternalServerError, []byte(`{"error": "Internal server error"}`))
	}

	// Check if database have user
	user, err := GetUserByEmail(instdb, req.Email)
	if err != nil {
		return c.JSONBlob(http.StatusBadRequest, []byte(fmt.Sprintf(`{"error": "%v"}`, err)))
	}

	// Compare password
	if !auth.ComparePassword(user.Password, req.Password) {
		return c.JSONBlob(http.StatusBadRequest, []byte(`{"error": "Incorrect password"}`))
	}

	secret := []byte(config.Env["JWTSecret"])
	token, err := auth.CreateJWT(secret, user.Id)

	return c.JSONBlob(http.StatusOK, []byte(fmt.Sprintf(`{"msg": "Login successfully", "token": "%v"}`, token)))
}
