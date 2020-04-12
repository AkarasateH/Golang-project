package database

import (
	"fmt"

	"database/sql"

	_ "github.com/lib/pq"
)

type Config struct { // Config structure to use to connect database
	User string
	Host string
	Port int
	Name string
}

func Open(cfg Config) (*sql.DB, error) {
	connStr := fmt.Sprintf("user=%s dbname=%s host=%s port=%d sslmode=verify-full", cfg.User, cfg.Name, cfg.Host, cfg.Port)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("Opening Database Error : >> ", err)
	}

	return db, err
}
