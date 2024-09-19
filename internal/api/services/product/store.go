package product

import (
	"database/sql"
	"fmt"
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
		err := rows.Scan(&p.Id, &p.Name, &p.Description, &p.Image, &p.Price, &p.Stock, &p.SellerId, &p.CreatedAt)
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

	err := rows.Scan(&p.Id, &p.Name, &p.Description, &p.Image, &p.Price, &p.Stock, &p.SellerId, &p.CreatedAt)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func GetProductBySellerId(db *sql.DB, sellerId int) ([]models.Product, error) {
	var products []models.Product
	rows, err := db.Query("SELECT * FROM products WHERE sellerId = $1", sellerId)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		p := &models.Product{}
		err := rows.Scan(&p.Id, &p.Name, &p.Description, &p.Image, &p.Price, &p.Stock, &p.SellerId, &p.CreatedAt)
		if err != nil {
			return nil, err
		}
	}
	return products, nil
}

func CheckSellerPrivilege(db *sql.DB, sellerId, productId int) (bool, error) {
	p := &models.Product{}
	rows := db.QueryRow("SELECT id FROM products WHERE sellerId = $1 AND id = $2", sellerId, productId)
	err := rows.Scan(&p.Id)
	if err != nil {
		return false, err
	}
	if p.Id == 0 {
		return false, fmt.Errorf("Unauthorized product or product does not exist")
	}
	return true, nil
}

func UpdateProduct(db *sql.DB, product *models.Product) error {
	_, err := db.Exec("UPDATE products SET name = $1, description = $2, image = $3, price = $4, stock = $5, createdAt = $6 WHERE id = $7", product.Name, product.Description, product.Image, product.Price, product.Stock, time.Now(), product.Id)
	if err != nil {
		return err
	}
	return nil
}

func CreateProduct(db *sql.DB, product *models.Product) error {
	_, err := db.Exec("INSERT INTO products(name, description, image, price, stock, sellerId) VALUES ($1, $2, $3, $4, $5, $6)", product.Name, product.Description, product.Image, product.Price, product.Stock, product.SellerId)
	if err != nil {
		return err
	}
	return nil
}

func DeleteProduct(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM products WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func SearchProducts(db *sql.DB, searchTerm string) ([]models.Product, error) {
	var products []models.Product

	rows, err := db.Query(
		`SELECT id, name, description, image, price, stock, sellerId
    FROM products
    WHERE search_vector @@ to_tsquery('english', $1)
    ORDER BY ts_rank(search_vector, to_tsquery('english', $1)) DESC`, searchTerm)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		p := &models.Product{}
		err := rows.Scan(&p.Id, &p.Name, &p.Description, &p.Image, &p.Price, &p.Stock, &p.SellerId)
		if err != nil {
			return nil, err
		}
		products = append(products, *p)
	}
	return products, nil
}
