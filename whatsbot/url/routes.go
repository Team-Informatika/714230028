package url

import (
	"github.com/gofiber/fiber/v2"
	"whatsbot/controller"
)

func SetupRoutes(app *fiber.App) {
	// Grouping endpoint jika ingin diatur per kategori
	api := app.Group("/api")

	// Route untuk menerima pesan masuk
	api.Post("/incoming", controller.HandleIncomingMessage)
}
