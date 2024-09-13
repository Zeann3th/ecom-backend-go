package main

import (
	"fmt"
	"log"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	m "github.com/zeann3th/ecom/internal/api/middleware"
	"github.com/zeann3th/ecom/internal/api/services/order"
	"github.com/zeann3th/ecom/internal/api/services/product"
	"github.com/zeann3th/ecom/internal/api/services/user"
	"github.com/zeann3th/ecom/internal/config"
)

func main() {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))

	v1 := e.Group("/api/v1")

	// User
	v1.POST("/register", user.UserRegisterHandler)

	v1.POST("/login", user.UserLoginHandler)

	// Product
	v1.GET("/products", product.AllProductsAcquisitionHandler)

	v1.POST("/products", product.ProductCreationHandler, m.JWTMiddleware)

	v1.GET("/products/:id", product.ProductAcquisitionHandler)

	v1.PUT("/products/:id", product.ProductUpdateHandler, m.JWTMiddleware)

	// Order
	v1.POST("/cart", order.OrderHandler, m.JWTMiddleware)

	// Render port
	port := config.Env["PORT"]
	if port == "" {
		port = "6969"
	}

	log.Fatal(e.Start(fmt.Sprintf(":%v", port)))
}
