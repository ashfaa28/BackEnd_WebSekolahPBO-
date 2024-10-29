// File: handlers/html_handler.go
package handlers

import (
	"github.com/gofiber/fiber/v2"
)

// RenderHome menampilkan halaman beranda (index.html)
func FormBerita(c *fiber.Ctx) error {
	return c.Render("formBerita", fiber.Map{
		"Title": "Form Berita",
	})
}

func FormPrestasi(c *fiber.Ctx) error {
	return c.Render("formPrestasi", fiber.Map{
		"Title": "Form Prestasi",
	})
}
