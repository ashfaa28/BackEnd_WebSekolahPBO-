package controllers

import (
	"Project_PBO/database"
	"Project_PBO/models/entity"
	"Project_PBO/models/req"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// AkunControllerShow untuk menampilkan semua akun
func AkunControllerShow(c *fiber.Ctx) error {
	var akuns []entity.Akun
	err := database.DB.Find(&akuns).Error
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal mendapatkan data akun",
		})
	}
	return c.JSON(akuns)
}

// AkunControllerCreate untuk membuat akun baru
func AkunControllerCreate(c *fiber.Ctx) error {
	akun := new(req.AkunReq) // Pastikan Anda memiliki req.AkunReq yang sesuai
	if err := c.BodyParser(akun); err != nil {
		return err
	}

	validation := validator.New()
	if err := validation.Struct(akun); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Validasi data akun gagal",
			"error":   err.Error(),
		})
	}

	akunBaru := entity.Akun{
		UserName:    akun.UserName,
		Password:    akun.Password,
		AsalSekolah: akun.AsalSekolah,
	}

	if err := database.DB.Create(&akunBaru).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Gagal membuat akun baru",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Akun berhasil dibuat",
		"data":    akunBaru,
	})
}
