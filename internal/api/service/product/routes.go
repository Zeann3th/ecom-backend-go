package product

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/zeann3th/ecom/internal/config"
	"github.com/zeann3th/ecom/internal/db"
)

func GetProductsHandler(c echo.Context) error {
	instdb, err := db.ConnectStorage(config.Env["DB_DRIVER"], config.Env["DB_CONN"])
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "Database connection failed",
		})
	}

	products, err := GetProducts(instdb)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "Product does not exist",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"products": products,
	})
}
