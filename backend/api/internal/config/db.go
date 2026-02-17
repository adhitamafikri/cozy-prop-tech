package config

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // PostgreSQL driver
)

func NewDBConnection(cfg *DBConfig) *sqlx.DB {
	db, err := sqlx.Connect("postgres", cfg.DSN())
	fmt.Println("What is the conn", cfg.DSN())
	fmt.Println("What is the err", err)
	if err != nil {
		log.Fatal("Something went wrong when opening new connection")
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully connected to postgres")
	}

	return db
}

type DBConfig struct {
	Host     string
	Port     int
	Name     string
	User     string
	Password string
}

func (c *DBConfig) DSN() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		c.Host, c.Port, c.User, c.Password, c.Name)
}
