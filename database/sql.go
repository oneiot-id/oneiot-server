package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"time"
)

func NewSqlConnection() *sql.DB {
	connectionString := os.Getenv("SQL_USER") + ":" + os.Getenv("SQL_PASSWORD") + "@tcp(127.0.0.1:3306)/" + os.Getenv("SQL_DATABASE")

	fmt.Println(connectionString)

	db, err := sql.Open("mysql", connectionString)

	if err != nil {
		return nil
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return db
}
