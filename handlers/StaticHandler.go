package handlers

import (
	"fmt"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

func StaticHandler(c *fiber.Ctx) error {
	path := c.Params("*")
	filePath := filepath.Join("static", path)

	// Kirim file ke client
	return c.SendFile(filePath, true)
}

func PkgHandler(c *fiber.Ctx) error {
	path := c.Params("*")
	filePath := filepath.Join("pkg", path)

	// Kirim file ke client
	return c.SendFile(filePath, true)
}

func UploadsHandler(c *fiber.Ctx) error {
	path := c.Params("*")
	filePath := filepath.Join("uploads", path)

	// Log filePath untuk debugging
	fmt.Println("Mengakses file:", filePath)

	// Kirim file ke client
	return c.SendFile(filePath, true)
}
