package handlers

import (
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

func StaticHandler(c *fiber.Ctx) error {
	path := c.Params("*")
	filePath := filepath.Join("static", path)

	// Kirim file ke client
	return c.SendFile(filePath, true)
}
