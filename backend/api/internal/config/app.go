package config

import (
	"log/slog"
	"os"
	"strconv"

	"log"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

type Config struct {
	AppEnv  string
	Host    string
	Port    int
	GinMode string
	DB      *DBConfig
	Redis   *RedisConfig
}

func LoadConfig() (*Config, error) {
	// Try .env.local first (for local development), fallback to .env
	err := godotenv.Load(".env.local")
	if err != nil {
		// .env.local doesn't exist, try .env
		err = godotenv.Load()
		if err != nil {
			log.Fatal("failed to load .env or .env.local file")
			return nil, err
		}
	}

	cfg := &Config{
		AppEnv:  GetValue("APP_ENV", "local"),
		Host:    GetValue("APP_HOST", "localhost"),
		Port:    GetInt("APP_PORT", 8082),
		GinMode: GetValue("GIN_MODE", "debug"),
		DB: &DBConfig{
			Host:     GetValue("DB_HOST", "localhost"),
			Port:     GetInt("DB_PORT", 5432),
			Name:     GetValue("DB_NAME", "cozy_prop_db"),
			User:     GetValue("DB_USER", "cozy"),
			Password: GetValue("DB_PASSWORD", "cozy123"),
		},
		Redis: &RedisConfig{
			Host:   GetValue("REDIS_HOST", "localhost"),
			Port:   GetInt("REDIS_PORT", 6379),
			Prefix: GetValue("REDIS_CACHE_PREFIX", "cory_prop:"),
		},
	}

	return cfg, err
}

func GetValue(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func GetInt(key string, fallback int) int {
	if value, exists := os.LookupEnv(key); exists {
		val, err := strconv.Atoi(value)
		if err != nil {
			panic("Invalid env value for: " + key)
		}
		return val
	}
	return fallback
}

type BoostrapConfig struct {
	Logger *slog.Logger
	DB     *sqlx.DB
	Redis  *redis.Client
	App    *gin.Engine
}

// Setup the Gin Engine for routing, middlewares, etc
func Bootstrap(cfg *BoostrapConfig) {
	cfg.App.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "index"})
	})
	cfg.App.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "hello"})
	})
	cfg.App.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "test"})
	})
}
