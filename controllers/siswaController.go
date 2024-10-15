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
			"message": "Validasi data siswa gagal",
			"error":   err.Error(),
		})
	}

	siswaBaru := entity.Siswa{
		Nama:           siswa.Nama,
		NISN:           siswa.NISN,
		NoTelp:         siswa.NoTelp,
		Email:          siswa.Email,
		Usia:           siswa.Usia,
		Alamat:         siswa.Alamat,
		NamaAyah:       siswa.NamaAyah,
		NamaIbu:        siswa.NamaIbu,
		NoTelpOrangTua: siswa.NoTelpOrangTua,
		AsalSekolah:    siswa.AsalSekolah,
	}

	if err := database.DB.Create(&siswaBaru).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Gagal membuat siswa baru",
		})

	}
	return c.JSON(fiber.Map{
		"message": "Siswa berhasil dibuat",
		"data":    siswaBaru,
	})
}
