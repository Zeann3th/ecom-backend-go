package order

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/zeann3th/ecom/internal/api/models"
)

type OrderHandler struct {
	DB *sql.DB
}

func (o *OrderHandler) HandleOrdersAcquisition(c echo.Context) error {
	userId := c.Get("user").(models.User).Id

	orders, total, err := GetOrdersByUserId(o.DB, userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err,
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"orders": orders,
		"total":  total,
	})
}

func (o *OrderHandler) HandleOrderCreation(c echo.Context) error {
	userId := c.Get("user").(models.User).Id
	req := new(models.OrderPayload)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid request payload",
		})
	}

	if req.ProductId == 0 || req.Quantity == 0 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Missing required fields",
		})
	}
	order := &models.Order{
		UserId:    userId,
		ProductId: req.ProductId,
		Quantity:  req.Quantity,
	}
	err := CreateOrder(o.DB, order)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err,
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"msg": "Order created successfully",
	})
}

func (o *OrderHandler) HandleOrderUpdate(c echo.Context) error {
	userId := c.Get("user").(models.User).Id
	req := new(models.OrderPayload)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid request payload",
		})
	}

	if req.ProductId == 0 || req.Quantity == 0 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Missing required fields",
		})
	}
	ok, err := CheckOrderExist(o.DB, userId, req.ProductId)
	if !ok {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"error": err,
		})
	}

	order := &models.Order{
		UserId:    userId,
		ProductId: req.ProductId,
		Quantity:  req.Quantity,
	}

	err = UpdateOrder(o.DB, order)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err,
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"msg": "Order updated successfully",
	})
}

func (o *OrderHandler) HandleOrderDeletion(c echo.Context) error {
	userId := c.Get("user").(models.User).Id
	productIdStr := c.Param("id")

	productId, err := strconv.Atoi(productIdStr)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err,
		})
	}
	err = DeleteOrder(o.DB, userId, productId)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"msg": "Order deleted successfully",
	})
}

func (o *OrderHandler) HandleCheckout(c echo.Context) error {
	return nil
}
