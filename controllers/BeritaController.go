// File: controllers/berita_controller.go
package controllers

import (
	"Project_PBO/database"
	"Project_PBO/models/entity"
	"Project_PBO/models/req"
	"log"
	"os"
	"path/filepath"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// Menampilkan semua data Berita
func BeritaControllerShow(c *fiber.Ctx) error {
	var beritas []entity.Berita
	err := database.DB.Find(&beritas).Error
	if err != nil {
		log.Println(err)
	}
	return c.JSON(beritas)
}

// Membuat Berita baru dengan gambar
func BeritaControllerCreate(c *fiber.Ctx) error {
	isiBerita := c.FormValue("isi_berita")
	if isiBerita == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Isi berita harus diisi",
		})
	}

	validation := validator.New()
	reqData := req.BeritaReq{IsiBerita: isiBerita}
	if err := validation.Struct(&reqData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Validasi data berita gagal",
			"error":   err.Error(),
		})
	}

	// Menyimpan file gambar ke folder uploads/berita
	file, err := c.FormFile("gambar")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Gambar harus diupload",
		})
	}

	filePath := filepath.Join("uploads", "berita", file.Filename)
	if err := c.SaveFile(file, filePath); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal menyimpan gambar",
			"error":   err.Error(),
		})
	}

	// Membuat entitas Berita baru
	beritaBaru := entity.Berita{
		Gambar:    filePath,
		IsiBerita: isiBerita,
	}

	if err := database.DB.Create(&beritaBaru).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal membuat berita baru",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Berita berhasil dibuat",
		"data":    beritaBaru,
	})
}

// Memperbarui Berita berdasarkan ID
func BeritaControllerUpdate(c *fiber.Ctx) error {
	id := c.Params("id")
	var berita entity.Berita

	if err := database.DB.First(&berita, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Berita tidak ditemukan",
		})
	}

	isiBerita := c.FormValue("isi_berita")
	if isiBerita == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Isi berita harus diisi",
		})
	}

	validation := validator.New()
	reqData := req.BeritaReq{IsiBerita: isiBerita}
	if err := validation.Struct(&reqData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Validasi data berita gagal",
			"error":   err.Error(),
		})
	}

	file, err := c.FormFile("gambar")
	if err == nil {
		filePath := filepath.Join("uploads", "berita", file.Filename)
		if err := c.SaveFile(file, filePath); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Gagal menyimpan gambar",
				"error":   err.Error(),
			})
		}

		if berita.Gambar != "" && berita.Gambar != filePath {
			os.Remove(berita.Gambar)
		}
		berita.Gambar = filePath
	}

	berita.IsiBerita = isiBerita

	if err := database.DB.Save(&berita).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal memperbarui data berita",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Data berita berhasil diperbarui",
		"data":    berita,
	})
}

// Menghapus Berita berdasarkan ID
func BeritaControllerDelete(c *fiber.Ctx) error {
	id := c.Params("id")
	var berita entity.Berita

	if err := database.DB.First(&berita, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Berita tidak ditemukan",
		})
	}

	if berita.Gambar != "" {
		os.Remove(berita.Gambar)
	}

	if err := database.DB.Delete(&berita).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal menghapus data berita",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Data berita berhasil dihapus",
	})
}
