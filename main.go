// main.go
package main

import (
	"Project_PBO/database"
	"Project_PBO/database/migration"
	"Project_PBO/routers"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {

	database.ConnectDB()

	migration.StartMigration()

	if _, err := os.Stat("uploads/berita"); os.IsNotExist(err) {
		os.MkdirAll("uploads/berita", os.ModePerm)
	}

	if _, err := os.Stat("uploads/prestasi"); os.IsNotExist(err) {
		os.MkdirAll("uploads/prestasi", os.ModePerm)
	}

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	routers.Route(app)

	if err := app.Listen(":3000"); err != nil {
		log.Fatalf("Gagal menjalankan server: %v", err)
	}
}
