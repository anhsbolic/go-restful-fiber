package app

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"go-restful-fiber/helper"
	"time"
)

const (
	host     = "localhost"
	port     = 5435
	user     = "postgres"
	password = "goApiPostgres1234"
	dbname   = "go_api_db"
)

func NewDB() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
