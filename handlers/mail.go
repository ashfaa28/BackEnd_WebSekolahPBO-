package handlers

import (
	"crypto/tls"
	"log"
	"os"

	"github.com/go-gomail/gomail"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
}

func sendEmail(to, subject, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("EMAIL"))
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	d := gomail.NewDialer("smtp.gmail.com", 587, os.Getenv("EMAIL"), os.Getenv("PASSWORD"))
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}

func HandleContactForm(c *fiber.Ctx) error {
	type ContactForm struct {
		Name    string `form:"name" json:"name"`
		Email   string `form:"email" json:"email"`
		Message string `form:"message" json:"message"`
	}

	var form ContactForm
	if err := c.BodyParser(&form); err != nil {
		return c.Status(400).SendString("Invalid form data")
	}

	subject := "New Message from " + form.Name
	body := `
        <h2>Contact Form Message</h2>
        <p><strong>Name:</strong> ` + form.Name + `</p>
        <p><strong>Email:</strong> ` + form.Email + `</p>
        <p><strong>Message:</strong></p>
        <p>` + form.Message + `</p>
    `

	err := sendEmail(os.Getenv("RECIPIENT"), subject, body)
	if err != nil {
		log.Println("Failed to send email:", err)
		return c.Status(500).JSON(fiber.Map{"success": false, "message": "Gagal Mengirim email"})
	}

	return c.JSON(fiber.Map{"success": true, "message": "Email sent successfully"})
}
