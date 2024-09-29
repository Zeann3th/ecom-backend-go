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
	"github.com/zeann3th/ecom/internal/api/upload"
	"github.com/zeann3th/ecom/internal/config"
	"github.com/zeann3th/ecom/internal/db"
)

func main() {
	err := config.LoadEnv()
	if err != nil {
		log.Printf("Cannot detect .env file, switching to os env variables")
	}

	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))

	v1 := e.Group("/api/v1")

	// Database
	instdb, err := db.ConnectStorage(os.Getenv("DB_DRIVER"), os.Getenv("DB_CONN"))
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

	v1.GET("/products/search", p.HandleProductsSearch)

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

	// Cloudinary
	up := &upload.UploadHandler{}

	v1.GET("/upload", up.GenerateSignature)

	// Render port
	port := os.Getenv("PORT")
	if port == "" {
		port = "6969"
	}

	log.Fatal(e.Start(fmt.Sprintf("0.0.0.0:%v", port)))
}
