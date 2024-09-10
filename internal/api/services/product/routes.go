package product

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/zeann3th/ecom/internal/api/models"
	"github.com/zeann3th/ecom/internal/config"
	"github.com/zeann3th/ecom/internal/db"
)

func AllProductsAcquisitionHandler(c echo.Context) error {
	instdb, err := db.ConnectStorage(config.Env["DB_DRIVER"], config.Env["DB_CONN"])
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "Database connection failed",
		})
	}

	products, err := GetAllProducts(instdb)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"error": "Product does not exist",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"products": products,
	})
}

func ProductAcquisitionHandler(c echo.Context) error {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid Id",
		})
	}

	instdb, err := db.ConnectStorage(config.Env["DB_DRIVER"], config.Env["DB_CONN"])
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "Database connection failed",
		})
	}

	product, err := GetProductById(instdb, id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"error": "Product does not exist",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"product": product,
	})
}

func ProductUpdateHandler(c echo.Context) error {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid Id",
		})
	}

	req := new(models.ProductPayload)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid request payload",
		})
	}

	instdb, err := db.ConnectStorage(config.Env["DB_DRIVER"], config.Env["DB_CONN"])
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "Database connection failed",
		})
	}

	p, err := GetProductById(instdb, id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"error": "Product does not exist",
		})
	}

	if req.Name != "" {
		p.Name = req.Name
	}
	if req.Image != "" {
		p.Image = req.Name
	}
	if req.Price != 0 {
		p.Price = req.Price
	}
	if req.Description != "" {
		p.Description = req.Description
	}

	err = UpdateProduct(instdb, p)

	return c.JSON(http.StatusAccepted, map[string]interface{}{
		"msg": "Product updated",
	})
}

func ProductCreationHandler(c echo.Context) error {
	req := new(models.ProductPayload)

	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid request payload",
		})
	}

	if req.Name == "" || req.Image == "" || req.Price == 0 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Missing required fields",
		})
	}
	// DB instance
	instdb, err := db.ConnectStorage(config.Env["DB_DRIVER"], config.Env["DB_CONN"])
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "Database connection failed",
		})
	}

	p := &models.Product{
		Name:        req.Name,
		Description: req.Description,
		Image:       req.Image,
		Price:       req.Price,
	}
	err = CreateProduct(instdb, p)

	return c.JSON(http.StatusAccepted, map[string]interface{}{
		"msg": "Product created",
	})
}
