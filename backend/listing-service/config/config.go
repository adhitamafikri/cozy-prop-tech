package config

import (
	"fmt"
	"os"
)

type Config struct {
	Port    string
	GinMode string
	DB      DatabaseConfig
	Redis   RedisConfig
}

type DatabaseConfig struct {
	Host     string
	Port     string
	Name     string
	User     string
	Password string
}

type RedisConfig struct {
	Host   string
	Port   string
	Prefix string
}

func LoadConfig() Config {
	return Config{
		Port:    getEnv("PORT", "8082"),
		GinMode: getEnv("GIN_MODE", "debug"),
		DB: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			Name:     getEnv("DB_NAME", "cozy_listings"),
			User:     getEnv("DB_USER", "cozy"),
			Password: getEnv("DB_PASSWORD", "cozy123"),
		},
		Redis: RedisConfig{
			Host:   getEnv("REDIS_HOST", "localhost"),
			Port:   getEnv("REDIS_PORT", "6379"),
			Prefix: getEnv("REDIS_CACHE_PREFIX", "sv_listing:"),
		},
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func (c Config) DSN() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.DB.Host, c.DB.Port, c.DB.User, c.DB.Password, c.DB.Name)
}

func (c Config) RedisAddr() string {
	return fmt.Sprintf("%s:%s", c.Redis.Host, c.Redis.Port)
}
