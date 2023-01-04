package app

import (
	"cart/helper"
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func NewBD() *sql.DB {
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/shopping-cart")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxIdleTime(60 * time.Minute)
	db.SetConnMaxLifetime(10 * time.Minute)

	return db
}
