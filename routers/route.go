package routers

import (
	"Project_PBO/controllers"
	"Project_PBO/handlers"

	"github.com/gofiber/fiber/v2"
)

func Route(c *fiber.App) {
	c.Get("/api/showAll", controllers.SiswaControllerShow)
	c.Post("/api/create", controllers.SiswaControllerCreate)
	c.Delete("/api/delete/:id", controllers.SiswaControllerDelete)
	c.Put("/api/update/:id", controllers.SiswaControllerUpdate)

	c.Get("/api/showAllBerita", controllers.SiswaControllerShow)
	c.Post("/api/createBerita", controllers.SiswaControllerCreate)
	c.Delete("/api/deleteBerita/:id", controllers.SiswaControllerDelete)
	c.Put("/api/updateBerita/:id", controllers.SiswaControllerUpdate)

	c.Get("/static/*", handlers.StaticHandler)

	c.Get("/", handlers.RenderHome)
}
