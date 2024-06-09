package main

import (
	"goapi/app/config"
	"goapi/app/routes"
	"log"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	e := echo.New()

	// Initialize database
	config.InitDB()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Routes
	routes.Router(e)

	// server
	e.Logger.Fatal(e.Start(":3541"))
}
