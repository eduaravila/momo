package config

import (
	"database/sql"
	"fmt"
	"os"
)

func InitPostgresDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME")))

	if err != nil {
		return nil, err
	}
	return db, db.Ping()

}
