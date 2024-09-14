package middleware

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/zeann3th/ecom/internal/api/models"
	"github.com/zeann3th/ecom/internal/api/services/product"
	"github.com/zeann3th/ecom/internal/config"
	"github.com/zeann3th/ecom/internal/db"
)

func IsSeller(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sellerId := c.Get("user").(models.User).Id
		productIdStr := c.Param("id")

		productId, err := strconv.Atoi(productIdStr)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"error": "Invalid Product Id",
			})
		}

		instdb, err := db.ConnectStorage(config.Env["DB_DRIVER"], config.Env["DB_CONN"])
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error": "Database connection failed",
			})
		}
		ok, err := product.CheckSellerPrivilege(instdb, sellerId, productId)
		if ok {
			return next(c)
		}
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"error": err,
		})
	}
}
