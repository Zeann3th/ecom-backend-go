package main

import (
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	m "github.com/zeann3th/ecom/internal/api/middleware"
	"github.com/zeann3th/ecom/internal/api/services/order"
	"github.com/zeann3th/ecom/internal/api/services/product"
	"github.com/zeann3th/ecom/internal/api/services/user"
	"github.com/zeann3th/ecom/internal/config"
	"github.com/zeann3th/ecom/internal/db"
)

func main() {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))

	v1 := e.Group("/api/v1")

	// Database
	instdb, err := db.ConnectStorage(config.Env["DB_DRIVER"], config.Env["DB_CONN"])
	if err != nil {
		log.Fatal(err)
	}

	// User
	u := &user.UserHandler{DB: instdb}

	v1.POST("/register", u.HandleUserRegister)

	v1.POST("/login", u.HandleUserLogin)

	// Product
	p := &product.ProductHandler{DB: instdb}

	v1.GET("/products", p.HandleAllProducts)

	v1.POST("/products", p.HandleProductCreation, m.JWTMiddleware)

	v1.GET("/products/:id", p.HandleProductById)

	v1.PUT("/products/:id", p.HandleProductUpdate, m.JWTMiddleware, m.IsSeller)

	v1.DELETE("/products/:id", p.HandleProductDeletion, m.JWTMiddleware, m.IsSeller)

	// Order
	o := &order.OrderHandler{DB: instdb}

	v1.GET("/cart", o.HandleOrdersAcquisition, m.JWTMiddleware)

	v1.POST("/cart", o.HandleOrderCreation, m.JWTMiddleware)

	v1.PUT("/cart/:id", o.HandleOrderUpdate, m.JWTMiddleware)

	v1.DELETE("/cart/:id", o.HandleOrderDeletion, m.JWTMiddleware)

	// Render port
	port := config.Env["PORT"]
	if port == "" {
		port = "6969"
	}

	log.Fatal(e.Start(fmt.Sprintf(":%v", port)))
}
