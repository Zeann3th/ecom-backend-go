package user

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/zeann3th/ecom/internal/api/models"
)

func GetUserByEmail(db *sql.DB, email string) (*models.User, error) {
	user := &models.User{}
	err := db.QueryRow("SELECT * FROM users WHERE email = $1", email).Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		return nil, err
	}

	if user.Id == 0 {
		return nil, fmt.Errorf("User %v not found", email)
	}

	return user, nil
}

func GetUserById(db *sql.DB, id int) (*models.User, error) {
	user := &models.User{}
	err := db.QueryRow("SELECT * FROM users WHERE id = $1", id).Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		return nil, err
	}

	if user.Id == 0 {
		return nil, fmt.Errorf("User ID %v not found", id)
	}

	return user, nil
}

func CreateUser(db *sql.DB, user *models.User) error {
	_, err := db.Exec("INSERT INTO users (firstName, lastName, email, password) VALUES ($1, $2, $3, $4)", user.FirstName, user.LastName, user.Email, user.Password)
	if err != nil {
		return err
	}
	return nil
}
