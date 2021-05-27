package config

import (
	"database/sql"
	"fmt"
	"os"
)

const (
	connectionPattern = "host=%s user=%s password=%s dbname=%s sslmode=disable"
)

// OpenDbConnection for open connection
func OpenDbConnection(desc string) *sql.DB {
	db, err := sql.Open("postgres", desc)
	if err != nil {
		defer db.Close()
		return db
	}
	return db
}

// CloseDbConnection for closing connection
func CloseDbConnection(db *sql.DB) {
	if db != nil {
		db.Close()
		db = nil
	}
}

func PostgresDB() *sql.DB {
	return OpenDbConnection(
		fmt.Sprintf(connectionPattern, os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME")))
}
