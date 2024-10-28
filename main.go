// main.go
package main

import (
	"Project_PBO/database"
	"Project_PBO/database/migration"
	"Project_PBO/routers"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {

	database.ConnectDB()

	migration.StartMigration()

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	routers.Route(app)

	if err := app.Listen(":3000"); err != nil {
		log.Fatalf("Gagal menjalankan server: %v", err)
	}
}
