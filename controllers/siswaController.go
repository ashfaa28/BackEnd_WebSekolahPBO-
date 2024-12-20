package controllers

import (
	"Project_PBO/database"
	"Project_PBO/models/entity"
	"Project_PBO/models/req"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func SiswaControllerShow(c *fiber.Ctx) error {
	var siswas []entity.Siswa
	err := database.DB.Find(&siswas).Error
	if err != nil {
		log.Println(err)
	}
	return c.JSON(siswas)
}

func SiswaControllerCreate(c *fiber.Ctx) error {
	siswa := new(req.SiswaReq)
	if err := c.BodyParser(siswa); err != nil {
		return err
	}

	validation := validator.New()
	if err := validation.Struct(siswa); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Validasi data siswa gagal",
			"details": err.Error(),
		})
	}

	// Check for duplicate NISN
	var existingSiswa entity.Siswa
	if err := database.DB.Where("nisn = ?", siswa.NISN).First(&existingSiswa).Error; err == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "NISN sudah terdaftar.",
		})
	}

	siswaBaru := entity.Siswa{
		Nama:           siswa.Nama,
		NISN:           siswa.NISN,
		JenisKelamin:   siswa.JenisKelamin,
		Email:          siswa.Email,
		Usia:           siswa.Usia,
		Alamat:         siswa.Alamat,
		NamaAyah:       siswa.NamaAyah,
		NamaIbu:        siswa.NamaIbu,
		NoTelpOrangTua: siswa.NoTelpOrangTua,
		AsalSekolah:    siswa.AsalSekolah,
		Jurusan:        siswa.Jurusan,
	}

	if err := database.DB.Create(&siswaBaru).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Gagal membuat siswa baru.",
		})
	}
	return c.JSON(fiber.Map{
		"error":   false,
		"message": "Siswa berhasil dibuat.",
		"data":    siswaBaru,
	})
}

func SiswaControllerUpdate(c *fiber.Ctx) error {
	id := c.Params("id")
	var siswa entity.Siswa

	if err := database.DB.First(&siswa, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "Siswa tidak ditemukan",
		})
	}

	reqData := new(req.SiswaReq)
	if err := c.BodyParser(reqData); err != nil {
		return err
	}

	validation := validator.New()
	if err := validation.Struct(reqData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Validasi data siswa gagal",
			"details": err.Error(),
		})
	}

	// Check if the new NISN is taken by another student
	var existingSiswa entity.Siswa
	if err := database.DB.Where("nisn = ? AND id != ?", reqData.NISN, id).First(&existingSiswa).Error; err == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "NISN sudah terdaftar.",
		})
	}

	siswa.Nama = reqData.Nama
	siswa.NISN = reqData.NISN
	siswa.JenisKelamin = reqData.JenisKelamin
	siswa.Email = reqData.Email
	siswa.Usia = reqData.Usia
	siswa.Alamat = reqData.Alamat
	siswa.NamaAyah = reqData.NamaAyah
	siswa.NamaIbu = reqData.NamaIbu
	siswa.NoTelpOrangTua = reqData.NoTelpOrangTua
	siswa.AsalSekolah = reqData.AsalSekolah
	siswa.Jurusan = reqData.Jurusan

	if err := database.DB.Save(&siswa).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Gagal memperbarui data siswa.",
		})
	}

	return c.JSON(fiber.Map{
		"error":   false,
		"message": "Data siswa berhasil diperbarui.",
		"data":    siswa,
	})
}

func SiswaControllerDelete(c *fiber.Ctx) error {
	id := c.Params("id")
	var siswa entity.Siswa

	if err := database.DB.First(&siswa, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "Siswa tidak ditemukan.",
		})
	}

	if err := database.DB.Delete(&siswa).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Gagal menghapus data siswa.",
		})
	}

	return c.JSON(fiber.Map{
		"error":   false,
		"message": "Data siswa berhasil dihapus.",
	})
}

func SiswaControllerShowByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var siswa entity.Siswa

	if err := database.DB.First(&siswa, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "Siswa tidak ditemukan",
		})
	}

	return c.JSON(siswa)
}
