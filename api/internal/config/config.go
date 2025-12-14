package config

import (
	"log"
	"os"
)

type Config struct {
	DatabaseURL string
	RedisAddr   string
	RedisPass   string
	WorkDir     string
}

func Load() Config {
	cfg := Config{
		DatabaseURL: os.Getenv("DATABASE_URL"),
		RedisAddr:   os.Getenv("REDIS_ADDR"),
		RedisPass:   os.Getenv("REDIS_PASSWORD"),
		WorkDir:     os.Getenv("OJ_WORKDIR"),
	}

	if cfg.DatabaseURL == "" {
		log.Fatal("DATABASE_URL is required")
	}
	if cfg.RedisAddr == "" {
		cfg.RedisAddr = "localhost:6379"
	}
	if cfg.WorkDir == "" {
		cfg.WorkDir = "jobs"
	}

	return cfg
}
