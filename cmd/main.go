package main

import (
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/zeann3th/ecom/internal/api/middleware"
	"github.com/zeann3th/ecom/internal/api/service/product"
	"github.com/zeann3th/ecom/internal/api/service/user"
	"github.com/zeann3th/ecom/internal/config"
)

func main() {
	e := echo.New()

	v1 := e.Group("/api/v1")

	// User

	v1.POST("/register", user.UserRegisterHandler)

	v1.POST("/login", user.UserLoginHandler)

	// Product

	v1.GET("/products", product.GetProductsHandler, middleware.JWTMiddleware)

	// Order

	log.Fatal(e.Start(fmt.Sprintf(":%v", config.Env["PORT"])))
}
