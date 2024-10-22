package routers

import (
	"Project_PBO/controllers"

	"github.com/gofiber/fiber/v2"
)

func Route(c *fiber.App) {
	c.Get("/api/showAll", controllers.SiswaControllerShow)
	c.Post("/api/create", controllers.SiswaControllerCreate)
	c.Get("/api/showAllAccount", controllers.AkunControllerShow)
	c.Post("/api/createAccount", controllers.AkunControllerCreate)
}
