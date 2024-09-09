package product

import (
	"database/sql"

	"github.com/zeann3th/ecom/internal/api/models"
)

func GetProducts(db *sql.DB) ([]models.Product, error) {
	var products []models.Product
	rows, err := db.Query("SELECT * FROM products")
	if err != nil {
		return products, err
	}

	for rows.Next() {
		p := &models.Product{}
		err := rows.Scan(&p.Id, &p.Name, &p.Description, &p.Image, &p.Price, &p.CreatedAt)
		if err != nil {
			return products, err
		}

		products = append(products, *p)
	}

	return products, nil
}
