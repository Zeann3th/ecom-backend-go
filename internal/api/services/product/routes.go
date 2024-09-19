package product

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/zeann3th/ecom/internal/api/models"
)

type ProductHandler struct {
	DB *sql.DB
}

func (p *ProductHandler) HandleAllProducts(c echo.Context) error {
	products, err := GetAllProducts(p.DB)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"error": "Product does not exist",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"products": products,
	})
}

func (p *ProductHandler) HandleProductById(c echo.Context) error {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid Id",
		})
	}

	product, err := GetProductById(p.DB, id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"error": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"product": product,
	})
}

func (p *ProductHandler) HandleProductUpdate(c echo.Context) error {
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

	product, err := GetProductById(p.DB, id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"error": err,
		})
	}

	if req.Name != "" {
		product.Name = req.Name
	}
	if req.Image != "" {
		product.Image = req.Name
	}
	if req.Price != 0 {
		product.Price = req.Price
	}
	if req.Stock != 0 {
		product.Stock = req.Stock
	}
	if req.Description != "" {
		product.Description = req.Description
	}

	err = UpdateProduct(p.DB, product)

	return c.JSON(http.StatusAccepted, map[string]interface{}{
		"msg": "Product updated",
	})
}

func (p *ProductHandler) HandleProductCreation(c echo.Context) error {
	sellerId := c.Get("user").(*models.User).Id
	req := new(models.ProductPayload)

	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid request payload",
		})
	}

	if req.Name == "" || req.Image == "" || req.Price == 0 || req.Stock == 0 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Missing required fields",
		})
	}

	product := &models.Product{
		Name:        req.Name,
		Description: req.Description,
		Image:       req.Image,
		Price:       req.Price,
		Stock:       req.Stock,
		SellerId:    sellerId,
	}
	err := CreateProduct(p.DB, product)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err,
		})
	}

	return c.JSON(http.StatusAccepted, map[string]interface{}{
		"msg": "Product created",
	})
}

func (p *ProductHandler) HandleProductDeletion(c echo.Context) error {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "Invalid Id",
		})
	}

	_, err = GetProductById(p.DB, id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"error": err,
		})
	}

	err = DeleteProduct(p.DB, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"msg": "Product deleted successfully",
	})
}

func (p *ProductHandler) HandleProductsSearch(c echo.Context) error {
	searchTerm := c.QueryParam("term")

	products, err := SearchProducts(p.DB, searchTerm)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err,
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"results": products,
	})
}
