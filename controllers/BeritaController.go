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

// Fungsi untuk menampilkan semua data Berita
func BeritaControllerShow(c *fiber.Ctx) error {
	var beritas []entity.Berita
	err := database.DB.Find(&beritas).Error
	if err != nil {
		log.Println(err)
	}
	return c.JSON(beritas)
}

// Fungsi untuk membuat Berita baru dengan gambar dan isi berita
func BeritaControllerCreate(c *fiber.Ctx) error {
	// Mengambil isi berita dari form-data
	isiBerita := c.FormValue("isi_berita")
	if isiBerita == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Isi berita harus diisi",
		})
	}

	// Validasi data input
	validation := validator.New()
	reqData := req.BeritaReq{IsiBerita: isiBerita}
	if err := validation.Struct(&reqData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Validasi data berita gagal",
			"error":   err.Error(),
		})
	}

	// Mengambil file gambar dari form-data
	file, err := c.FormFile("gambar")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Gambar harus diupload",
		})
	}

	// Menyimpan file gambar ke direktori server
	filePath := filepath.Join("uploads", file.Filename)
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

// Fungsi untuk memperbarui Berita berdasarkan ID
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

	// Validasi data input
	validation := validator.New()
	reqData := req.BeritaReq{IsiBerita: isiBerita}
	if err := validation.Struct(&reqData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Validasi data berita gagal",
			"error":   err.Error(),
		})
	}

	// Memperbarui isi berita dan gambar jika ada file baru yang diupload
	file, err := c.FormFile("gambar")
	if err == nil {
		filePath := filepath.Join("uploads", file.Filename)
		if err := c.SaveFile(file, filePath); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Gagal menyimpan gambar",
				"error":   err.Error(),
			})
		}

		// Hapus file gambar lama jika berbeda
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

// Fungsi untuk menghapus Berita berdasarkan ID
func BeritaControllerDelete(c *fiber.Ctx) error {
	id := c.Params("id")
	var berita entity.Berita

	if err := database.DB.First(&berita, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Berita tidak ditemukan",
		})
	}

	// Menghapus file gambar dari server
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
