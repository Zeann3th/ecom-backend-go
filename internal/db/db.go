package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
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

	log.Println(fmt.Sprintf("[%v]: Successfully connected!!!", os.Getenv("DB_DRIVER")))

	return db, nil
}

func CloseStorage(db *sql.DB) {
	db.Close()
}
