package main

import (
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/zeann3th/ecom/internal/api/middleware"
	"github.com/zeann3th/ecom/internal/api/services/order"
	"github.com/zeann3th/ecom/internal/api/services/product"
	"github.com/zeann3th/ecom/internal/api/services/user"
	"github.com/zeann3th/ecom/internal/config"
)

func main() {
	e := echo.New()

	v1 := e.Group("/api/v1")

	// User

	v1.POST("/register", user.UserRegisterHandler)

	v1.POST("/login", user.UserLoginHandler)

	// Product

	v1.GET("/products", product.AllProductsAcquisitionHandler)

	v1.POST("/products", product.ProductCreationHandler, middleware.JWTMiddleware)

	v1.GET("/products/:id", product.ProductAcquisitionHandler)

	v1.PUT("/products/:id", product.ProductUpdateHandler, middleware.JWTMiddleware)

	// Order

	v1.POST("/cart", order.OrderHandler, middleware.JWTMiddleware)

	log.Fatal(e.Start(fmt.Sprintf(":%v", config.Env["PORT"])))
}
