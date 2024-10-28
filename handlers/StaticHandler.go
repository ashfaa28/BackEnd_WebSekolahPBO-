// File: handlers/static_handler.go
package handlers

import (
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

// StaticHandler menyajikan file statis dari folder public
func StaticHandler(c *fiber.Ctx) error {
	// Dapatkan path dari request dan gabungkan dengan folder public
	path := c.Params("*")
	filePath := filepath.Join("static", path)

	// Kirim file ke client
	return c.SendFile(filePath, true)
}
