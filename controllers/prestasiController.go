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

// Fungsi untuk menampilkan semua data Prestasi
func PrestasiControllerShow(c *fiber.Ctx) error {
	var prestasis []entity.Prestasi
	if err := database.DB.Find(&prestasis).Error; err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal mengambil data prestasi",
		})
	}
	return c.JSON(prestasis)
}

// Fungsi untuk membuat prestasi baru
func PrestasiControllerCreate(c *fiber.Ctx) error {
	file, err := c.FormFile("gambar")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Gambar harus diupload",
			"error":   err.Error(),
		})
	}

	// Validasi input
	validation := validator.New()
	reqData := req.PrestasiReq{Gambar: file.Filename}
	if err := validation.Struct(&reqData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Validasi data prestasi gagal",
			"error":   err.Error(),
		})
	}

	filePath := filepath.Join("uploads", "prestasi", file.Filename)
	if err := c.SaveFile(file, filePath); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal menyimpan gambar",
			"error":   err.Error(),
		})
	}

	prestasiBaru := entity.Prestasi{
		Gambar: filePath,
	}

	if err := database.DB.Create(&prestasiBaru).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal membuat data prestasi",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Prestasi berhasil diupload",
		"data":    prestasiBaru,
	})
}

// Fungsi untuk memperbarui data prestasi
func PrestasiControllerUpdate(c *fiber.Ctx) error {
	id := c.Params("id")
	var prestasi entity.Prestasi

	if err := database.DB.First(&prestasi, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data prestasi tidak ditemukan",
		})
	}

	file, err := c.FormFile("gambar")
	if err == nil {
		// Validasi input
		validation := validator.New()
		reqData := req.PrestasiReq{Gambar: file.Filename}
		if err := validation.Struct(&reqData); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Validasi data prestasi gagal",
				"error":   err.Error(),
			})
		}

		filePath := filepath.Join("uploads", "prestasi", file.Filename)
		if err := c.SaveFile(file, filePath); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Gagal menyimpan gambar baru",
				"error":   err.Error(),
			})
		}

		// Hapus file gambar lama jika berbeda
		if prestasi.Gambar != "" && prestasi.Gambar != filePath {
			os.Remove(prestasi.Gambar)
		}
		prestasi.Gambar = filePath
	}

	if err := database.DB.Save(&prestasi).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal memperbarui data prestasi",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Data prestasi berhasil diperbarui",
		"data":    prestasi,
	})
}

func PrestasiControllerDelete(c *fiber.Ctx) error {
	id := c.Params("id")
	var prestasi entity.Prestasi

	if err := database.DB.First(&prestasi, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data prestasi tidak ditemukan",
		})
	}

	if prestasi.Gambar != "" {
		os.Remove(prestasi.Gambar)
	}

	if err := database.DB.Delete(&prestasi).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal menghapus data prestasi",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Data prestasi berhasil dihapus",
	})
}
