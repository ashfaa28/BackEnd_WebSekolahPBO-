// File: handlers/html_handler.go
package handlers

import (
	"github.com/gofiber/fiber/v2"
)

// RenderHome menampilkan halaman beranda (index.html)
func RenderHome(c *fiber.Ctx) error {
	return c.Render("formBerita", fiber.Map{
		"Title": "Form Berita",
	})
}
