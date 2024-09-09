package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/zeann3th/ecom/internal/config"
)

func ConnectStorage(driver string, conn string) (*sql.DB, error) {
	db, err := sql.Open(driver, conn)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println(fmt.Sprintf("[%v]: Successfully connected!!!", config.Env["DB_DRIVER"]))

	return db, nil
}

func CloseStorage(db *sql.DB) {
	db.Close()
}
