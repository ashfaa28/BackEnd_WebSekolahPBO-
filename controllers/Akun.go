package controllers

import (
	"Project_PBO/database"
	"Project_PBO/middleware"
	"Project_PBO/models/entity"
	"Project_PBO/models/req"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *fiber.Ctx) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	if isValidCredentials(email, password) {
		sess, err := middleware.Store.Get(c)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Error getting session",
			})
		}

		sess.Set("authenticated", true)
		sess.Save()

		return c.Redirect("/db/admin/dashboard")
	}

	return c.Redirect("/db/admin/sign_in")
}

func SignUp(c *fiber.Ctx) error {
	akun := new(req.AkunReq)
	if err := c.BodyParser(akun); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request data",
		})
	}

	validate := validator.New()
	if err := validate.Struct(akun); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Validation failed",
			"error":   err.Error(),
		})
	}

	var existingAkun entity.Akun
	if err := database.DB.Where("email = ?", akun.Email).First(&existingAkun).Error; err == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Email already registered",
		})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(akun.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to process password",
		})
	}

	newAkun := entity.Akun{
		Email:    akun.Email,
		Password: string(hashedPassword),
	}

	if err := database.DB.Create(&newAkun).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create account",
		})
	}

	return c.Redirect("/db/admin/sign_in")
}

func isValidCredentials(email, password string) bool {
	var akun entity.Akun
	if err := database.DB.Where("email = ?", email).First(&akun).Error; err != nil {
		return false
	}

	err := bcrypt.CompareHashAndPassword([]byte(akun.Password), []byte(password))
	return err == nil
}
