package controller

import (
	"strings"
	"whatsbot/model"
	"whatsbot/repository"

	"github.com/gofiber/fiber/v2"
)

func HandleIncomingMessage(c *fiber.Ctx) error {
	var incoming model.IncomingMessage

	// Parse request body
	if err := c.BodyParser(&incoming); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request",
		})
	}

	// Ambil semua keyword dari database
	keywords, err := repository.GetKeywordsFromDB()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve keywords",
		})
	}

	// Normalisasi pesan menjadi lowercase
	message := strings.ToLower(incoming.Message)

	// Cek apakah ada keyword yang cocok
	for _, keyword := range keywords {
		if strings.Contains(message, strings.ToLower(keyword.Keyword)) {
			// Return reply jika keyword cocok
			return c.JSON(fiber.Map{
				"number": incoming.Number,
				"reply":  keyword.Reply,
			})
		}
	}

	// Jika tidak ada keyword yang cocok
	return c.JSON(fiber.Map{
		"number": incoming.Number,
		"reply":  "Maaf, saya tidak mengerti pesan Anda.",
	})
}
