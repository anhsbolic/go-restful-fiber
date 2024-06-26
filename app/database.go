package app

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"go-restful-fiber/config"
	"go-restful-fiber/pkg"
	"time"
)

func NewDB() *sql.DB {
	// Get Config
	env := config.GetEnvConfig()

	// Connect to database
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		env.Get("DB_HOST"),
		env.Get("DB_PORT"),
		env.Get("DB_USER"),
		env.Get("DB_PASSWORD"),
		env.Get("DB_NAME"),
		env.Get("DB_SSL_MODE"),
	)
	db, err := sql.Open("postgres", psqlInfo)
	pkg.PanicIfError(err)

	// Set up database connection
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
