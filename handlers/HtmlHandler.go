package handlers

import (
	"Project_PBO/database"
	"time"

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
	var schoolCounts []struct {
		AsalSekolah string
		Count       int64
	}

	var latestDate time.Time
	database.DB.Table("siswas").Select("created_at").Order("created_at desc").Limit(1).Scan(&latestDate)

	database.DB.Table("siswas").
		Select("asal_sekolah, COUNT(*) as count").
		Group("asal_sekolah").
		Order("count ASC").
		Find(&schoolCounts)

	schools := make([]string, len(schoolCounts))
	counts := make([]int64, len(schoolCounts))

	for i, sc := range schoolCounts {
		schools[i] = sc.AsalSekolah
		counts[i] = sc.Count
	}

	return c.Render("dashboard", fiber.Map{
		"Title":      "Dashboard Admin SMK CN",
		"Schools":    schools,
		"Counts":     counts,
		"LatestDate": latestDate.Format("02 January 2006"),
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
