package product

import (
	"database/sql"
	"time"

	"github.com/zeann3th/ecom/internal/api/models"
)

func GetAllProducts(db *sql.DB) ([]models.Product, error) {
	var products []models.Product
	rows, err := db.Query("SELECT * FROM products")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		p := &models.Product{}
		err := rows.Scan(&p.Id, &p.Name, &p.Description, &p.Image, &p.Price, &p.Stock, &p.CreatedAt)
		if err != nil {
			return nil, err
		}

		products = append(products, *p)
	}

	return products, nil
}

func GetProductById(db *sql.DB, id int) (*models.Product, error) {
	p := &models.Product{}
	rows := db.QueryRow("SELECT * FROM products WHERE id = $1", id)

	err := rows.Scan(&p.Id, &p.Name, &p.Description, &p.Image, &p.Price, &p.Stock, &p.CreatedAt)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func UpdateProduct(db *sql.DB, product *models.Product) error {
	_, err := db.Exec("UPDATE products SET name = $1, description = $2, image = $3, price = $4, stock = $5, createdAt = $6 WHERE id = $7", product.Name, product.Description, product.Image, product.Price, product.Stock, time.Now(), product.Id)
	if err != nil {
		return err
	}
	return nil
}

func CreateProduct(db *sql.DB, product *models.Product) error {
	_, err := db.Exec("INSERT INTO products(name, description, image, price) VALUES ($1, $2, $3, $4, $5)", product.Name, product.Description, product.Image, product.Price, product.Stock)
	if err != nil {
		return err
	}
	return nil
}
