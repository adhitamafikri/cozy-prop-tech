package main

import (
	"fmt"
	"log"

	"github.com/adhitamafikri/cozy-prop-tech/backend/api/internal/config"
)

func main() {
	cfg, err := config.LoadConfig()

	if err != nil {
		log.Fatal("Failed to load configuration")
	}

	logger := config.NewLogger()
	db := config.NewDBConnection(&config.DBConfig{
		Host:     cfg.DB.Host,
		Port:     cfg.DB.Port,
		Name:     cfg.DB.Name,
		User:     cfg.DB.User,
		Password: cfg.DB.Password,
	})
	cache := config.NewRedis(&config.RedisConfig{
		Host:   cfg.Redis.Host,
		Port:   cfg.Redis.Port,
		Prefix: cfg.Redis.Prefix,
	})
	app := config.NewGinEngine()

	config.Bootstrap(&config.BoostrapConfig{
		Logger: logger,
		DB:     db,
		Redis:  cache,
		App:    app,
	})

	address := fmt.Sprintf("%s:%d", config.GetValue("APP_HOST", "0.0.0.0"), config.GetInt("APP_PORT", 8082))
	app.Run(address)
}
