// main.go
package main

import (
	"Project_PBO/database"
	"Project_PBO/database/migration"
	"Project_PBO/routers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDB()

	migration.StartMigration()

	app := fiber.New()

	// Setup routing dan file statis
	routers.Route(app)

	// Menjalankan server di port 3000
	app.Listen(":3000")
}
