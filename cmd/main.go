package main

import (
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/zeann3th/ecom/internal/api/service/general"
	"github.com/zeann3th/ecom/internal/api/service/user"
	"github.com/zeann3th/ecom/internal/config"
)

func main() {
	e := echo.New()

	e.GET("/api/v1", general.HelloHandler)

	e.POST("/api/v1/register", user.UserRegisterHandler)

	e.POST("/api/v1/login", user.UserLoginHandler)

	log.Fatal(e.Start(fmt.Sprintf(":%v", config.Env["PORT"])))
}
