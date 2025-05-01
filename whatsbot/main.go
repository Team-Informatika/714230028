package main

import (
	"log"
	"os"

	"whatsbot/config"
	"whatsbot/url"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Koneksi ke database
	config.CreateDBConnection()

	// Fiber app
	app := fiber.New()

	// Setup routes
	url.SetupRoutes(app)

	// Port dari .env
	port := os.Getenv("PORT")

	log.Printf("🚀 Server running at http://localhost:%s", port)
	log.Fatal(app.Listen(":" + port))
}
