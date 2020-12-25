package main

import (
	"context"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/dsouzadyn/micron-go.go/handlers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gopkg.in/yaml.v2"
)

// Config structure for yml
type Config struct {
	Server struct {
		Host string
		Port string
	}
	Database struct {
		Host     string
		Port     string
		Username string
		Password string
	}
}

func loadConfig() Config {
	content, err := ioutil.ReadFile("config.yml")
	if err != nil {
		log.Fatal(err)
	}
	c := Config{}
	yaml.Unmarshal(content, &c)
	return c
}

func main() {
	config := loadConfig()

	// Create an echo instance
	e := echo.New()

	// Add the middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Setup the health endpoint
	e.GET("/", handlers.HealthHandler)

	// Setup the server address
	address := config.Server.Host + ":" + config.Server.Port

	// Start the server
	go func() {
		if err := e.Start(address); err != nil {
			e.Logger.Info("shutting down health service.")
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
