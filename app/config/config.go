package config

import (
	"github.com/spf13/viper"
	"os"
	"time"
)

var C Config

type (
	Config struct {
		App      AppConfig
		Postgres PostgresConfig
		Redis    RedisConfig
		HTTP     HTTPConfig
		Auth     AuthConfig
		Stats    StatsConfig
	}

	AppConfig struct {
		BaseDir  string
		MediaDir string
	}

	HTTPConfig struct {
		Host string
		Port string
	}

	PostgresConfig struct {
		Host     string
		Port     string
		Username string
		Password string
		DBName   string
	}

	RedisConfig struct {
		Host string
		Port string
	}

	AuthConfig struct {
		PasswordSalt string
		TTL          time.Duration
		SigningKey   string
	}

	StatsConfig struct {
		TTL time.Duration
	}
)

func InitConfig() *Config {
	SetConfig(&C)
	return &C
}

func SetConfig(cfg *Config) {
	// App
	cfg.App.BaseDir = viper.GetString("app.dirs.base")
	cfg.App.MediaDir = viper.GetString("app.dirs.media")

	// HTTP
	cfg.HTTP.Host = viper.GetString("app.server.http.host")
	cfg.HTTP.Port = viper.GetString("app.server.http.port")

	// Postgres
	cfg.Postgres.Host = viper.GetString("postgresql.host")
	cfg.Postgres.Port = viper.GetString("postgresql.port")
	cfg.Postgres.Username = os.Getenv("POSTGRES_USER")
	cfg.Postgres.Password = os.Getenv("POSTGRES_PASSWORD")
	cfg.Postgres.DBName = os.Getenv("POSTGRES_DB")

	// Redis
	cfg.Redis.Host = viper.GetString("redis.host")
	cfg.Redis.Port = viper.GetString("redis.port")

	// Auth
	cfg.Auth.PasswordSalt = os.Getenv("PASSWORD_SALT")
	cfg.Auth.TTL = viper.GetDuration("app.token.ttl")
	cfg.Auth.SigningKey = os.Getenv("SIGNING_KEY")

	// Stats
	cfg.Stats.TTL = viper.GetDuration("app.stats.ttl")
}
