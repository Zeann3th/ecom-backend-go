package order

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/zeann3th/ecom/internal/api/models"
	"github.com/zeann3th/ecom/internal/api/services/product"
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
	userId := c.Get("user").(models.User).Id
	orders, _, err := GetOrdersByUserId(o.DB, userId)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"error": err,
		})
	}
	tx, err := o.DB.Begin()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err,
		})
	}
	for _, order := range orders {
		product, err := product.GetProductById(o.DB, order.ProductId)
		if err != nil {
			tx.Rollback()
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"error": err,
			})
		}
		if order.Quantity > product.Stock {
			tx.Rollback()
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"error": fmt.Errorf("Product %v having only %v in stock, can't fullfil the request of %v", product.Name, product.Stock, order.Quantity),
			})
		} else {
			res, err := tx.Exec("UPDATE products SET stock = $1 WHERE id = $2", product.Stock-order.Quantity, product.Id)
			if err != nil {
				tx.Rollback()
				return err
			}
			res, err = tx.Exec("DELETE FROM orders WHERE userId = $1 AND productId = $2", userId, product.Id)
			if err != nil {
				tx.Rollback()
				return err
			}
			_ = res
		}
	}
	err = tx.Commit()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"msg": "Hand over your Credit card!!!",
	})
}
