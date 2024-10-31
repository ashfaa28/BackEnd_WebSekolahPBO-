package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func Home(c *fiber.Ctx) error {
	return c.Render("Home", fiber.Map{
		"Title": "SMK Citra Negara",
	})
}

func Form(c *fiber.Ctx) error {
	return c.Render("Form", fiber.Map{
		"Title": "Form Pendaftaran - Sekolah Citra Negara",
	})
}

func Register(c *fiber.Ctx) error {
	return c.Render("Register", fiber.Map{
		"Title": "Form Register - Sekolah Citra Negara",
	})
}

func Dashboard(c *fiber.Ctx) error {
	return c.Render("dashboard", fiber.Map{
		"Title": "Dashboard Admin SMK CN",
	})
}

func FormBerita(c *fiber.Ctx) error {
	return c.Render("berita", fiber.Map{
		"Title": "Form Berita SMK CN",
	})
}

func FormPrestasi(c *fiber.Ctx) error {
	return c.Render("prestasi", fiber.Map{
		"Title": "Form Prestasi SMK CN",
	})
}

func InfoSiswa(c *fiber.Ctx) error {
	return c.Render("infoSiswa", fiber.Map{
		"Title": "Info Siswa SMK CN",
	})
}

func SignIn(c *fiber.Ctx) error {
	return c.Render("signIn", fiber.Map{
		"Title": "Sign In SMK CN",
	})
}

func SignUp(c *fiber.Ctx) error {
	return c.Render("signUp", fiber.Map{
		"Title": "Sign Up SMK CN",
	})
}

func Profile(c *fiber.Ctx) error {
	return c.Render("profile", fiber.Map{
		"Title": "Profile admin SMK CN",
	})
}
