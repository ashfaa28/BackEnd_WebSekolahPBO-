package controllers

import (
	"Project_PBO/database"
	"Project_PBO/models/entity"
	"log"
	"os"
	"path/filepath"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// EkskulControllerShow - Menampilkan semua data ekskul
func EkskulControllerShow(c *fiber.Ctx) error {
	var ekskuls []entity.Ekskul
	err := database.DB.Find(&ekskuls).Error
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal mengambil data ekskul",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Data ekskul berhasil diambil",
		"data":    ekskuls,
	})
}

// EkskulControllerCreate - Membuat data ekskul baru
func EkskulControllerCreate(c *fiber.Ctx) error {
	namaEkskul := c.FormValue("nama_ekskul")
	if namaEkskul == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Nama ekskul harus diisi",
		})
	}

	validation := validator.New()
	if err := validation.Var(namaEkskul, "required"); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Validasi data ekskul gagal",
			"error":   err.Error(),
		})
	}

	file, err := c.FormFile("gambar")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Gambar harus diupload",
		})
	}

	filePath := filepath.Join("uploads", "ekskul", file.Filename)
	if err := c.SaveFile(file, filePath); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal menyimpan gambar",
			"error":   err.Error(),
		})
	}

	ekskulBaru := entity.Ekskul{
		NamaEkskul: namaEkskul,
		Gambar:     filePath,
	}

	if err := database.DB.Create(&ekskulBaru).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal membuat ekskul baru",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Ekskul berhasil dibuat",
		"data":    ekskulBaru,
	})
}

// EkskulControllerUpdate - Memperbarui data ekskul
func EkskulControllerUpdate(c *fiber.Ctx) error {
	id := c.Params("id")
	var ekskul entity.Ekskul

	if err := database.DB.First(&ekskul, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Ekskul tidak ditemukan",
		})
	}

	namaEkskul := c.FormValue("nama_ekskul")
	if namaEkskul == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Nama ekskul harus diisi",
		})
	}

	validation := validator.New()
	if err := validation.Var(namaEkskul, "required"); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Validasi data ekskul gagal",
			"error":   err.Error(),
		})
	}

	file, err := c.FormFile("gambar")
	if err == nil {
		filePath := filepath.Join("uploads", "ekskul", file.Filename)
		if err := c.SaveFile(file, filePath); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Gagal menyimpan gambar",
				"error":   err.Error(),
			})
		}

		if ekskul.Gambar != "" && ekskul.Gambar != filePath {
			os.Remove(ekskul.Gambar)
		}
		ekskul.Gambar = filePath
	}

	ekskul.NamaEkskul = namaEkskul

	if err := database.DB.Save(&ekskul).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal memperbarui data ekskul",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Data ekskul berhasil diperbarui",
		"data":    ekskul,
	})
}

// EkskulControllerDelete - Menghapus data ekskul
func EkskulControllerDelete(c *fiber.Ctx) error {
	id := c.Params("id")
	var ekskul entity.Ekskul

	if err := database.DB.First(&ekskul, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Ekskul tidak ditemukan",
		})
	}

	if ekskul.Gambar != "" {
		os.Remove(ekskul.Gambar)
	}

	if err := database.DB.Delete(&ekskul).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal menghapus data ekskul",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Data ekskul berhasil dihapus",
	})
}

// EkskulControllerShowByID - Menampilkan data ekskul berdasarkan ID
func EkskulControllerShowByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var ekskul entity.Ekskul

	if err := database.DB.First(&ekskul, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Ekskul tidak ditemukan",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Data ekskul berhasil ditemukan",
		"data":    ekskul,
	})
}
