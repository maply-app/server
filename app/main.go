package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"maply/api"
	"maply/cache"
	"maply/config"
	prometheusMiddleware "maply/prometheus"
	"maply/repository"
	"maply/ws"
	"os"
)

func main() {
	// Initialise Viper
	if err := InitViper(); err != nil {
		log.Fatalf("Failed to open config file: %s", err.Error())
	}

	// Initialise Logrus
	InitLogrus()

	// Config initialise
	var cfg = config.InitConfig()

	// Initialise a PostgreSQL connection
	repository.InitPostgres(cfg.Postgres)

	// Initialise a Redis connection
	cache.InitRedis(cfg.Redis)

	// Initialise Fiber web server
	app := fiber.New(fiber.Config{ServerHeader: "Maply-Server", Prefork: false}) // true
	app.Use(cors.New())

	// Setup prometheus
	prometheus := prometheusMiddleware.New("maply-server")
	prometheus.RegisterAt(app, "/metrics")
	app.Use(prometheus.Middleware)

	// Setup server
	api.SetupRoutes(app)
	ws.SetupRoutes(app)
	addr := fmt.Sprintf("%s:%s", cfg.HTTP.Host, cfg.HTTP.Port)
	if err := app.Listen(addr); err != nil {
		log.Fatalf("Failed to start web server: %s", err.Error())
	}
}

func InitViper() error {
	// Viper config
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")
	return viper.ReadInConfig()
}

func InitLogrus() {
	// Initialise a logger
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above
	log.SetLevel(log.TraceLevel) // WarnLevel
}
