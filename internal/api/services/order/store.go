package order

import (
	"database/sql"

	"github.com/zeann3th/ecom/internal/api/models"
)

func GetOrdersByUserId(db *sql.DB, userId int) ([]models.Order, float64, error) {
	var orders []models.Order
	var total float64

	rows, err := db.Query("SELECT * FROM orders WHERE user_id = $1", userId)
	if err != nil {
		return nil, 0, err
	}

	for rows.Next() {
		o := &models.Order{}
		err := rows.Scan(&o.UserId, &o.ProductId, o.Quantity, o.CreatedAt)
		if err != nil {
			return nil, 0, err
		}

		orders = append(orders, *o)
	}

	return orders, total, nil
}

func GetOrdersByProductId(db *sql.DB, productId int) ([]models.Order, error) {
	var orders []models.Order

	rows, err := db.Query("SELECT * FROM orders WHERE product_id = $1", productId)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		o := &models.Order{}
		err := rows.Scan(&o.UserId, &o.ProductId, o.Quantity, o.CreatedAt)
		if err != nil {
			return nil, err
		}

		orders = append(orders, *o)
	}

	return orders, nil
}

func CreateOrder(db *sql.DB, order *models.Order) error {
	_, err := db.Exec("INSERT INTO orders(user_id, product_id, quantity) VALUES ($1, $2, $3)", order.UserId, order.ProductId, order.Quantity)
	if err != nil {
		return err
	}
	return nil
}

func UpdateOrder(db *sql.DB, order *models.Order) error {
	_, err := db.Exec("UPDATE orders SET quantity = $1 WHERE user_id = $2 AND product_id = $3", order.Quantity, order.UserId, order.ProductId)
	if err != nil {
		return err
	}
	return nil
}

func UpdateOrders(db *sql.DB, orders []models.Order) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	for _, order := range orders {
		_, err = tx.Exec("UPDATE orders SET quantity = $1 WHERE user_id = $2 AND product_id = $3", order.Quantity, order.UserId, order.ProductId)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func DeleteOrder(db *sql.DB, userId, productId int) error {
	_, err := db.Exec("DELETE FROM orders WHERE user_id = $1 AND product_id = $2", userId, productId)
	if err != nil {
		return err
	}
	return nil
}
