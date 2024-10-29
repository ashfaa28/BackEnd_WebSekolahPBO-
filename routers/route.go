package routers

import (
	"Project_PBO/controllers"
	"Project_PBO/handlers"

	"github.com/gofiber/fiber/v2"
)

func Route(app *fiber.App) {

	siswaGroup := app.Group("/api/siswa")
	siswaGroup.Get("/showAll", controllers.SiswaControllerShow)
	siswaGroup.Post("/create", controllers.SiswaControllerCreate)
	siswaGroup.Delete("/delete/:id", controllers.SiswaControllerDelete)
	siswaGroup.Put("/update/:id", controllers.SiswaControllerUpdate)

	beritaGroup := app.Group("/api/berita")
	beritaGroup.Get("/showAll", controllers.BeritaControllerShow)
	beritaGroup.Post("/create", controllers.BeritaControllerCreate)
	beritaGroup.Delete("/delete/:id", controllers.BeritaControllerDelete)
	beritaGroup.Put("/update/:id", controllers.BeritaControllerUpdate)

	prestasiGroup := app.Group("api/prestasi")
	prestasiGroup.Get("/showAll", controllers.PrestasiControllerShow)
	prestasiGroup.Post("/create", controllers.PrestasiControllerCreate)
	prestasiGroup.Put("/update/:id", controllers.PrestasiControllerUpdate)
	prestasiGroup.Delete("/delete/:id", controllers.PrestasiControllerDelete)

	app.Get("/static/*", handlers.StaticHandler)

	app.Get("/", handlers.FormBerita)
	app.Get("/upload/prestasi", handlers.FormPrestasi)

	//Buat user gabisa ke mana mana selain yang ada di route
	app.Use(func(c *fiber.Ctx) error {
		return c.Redirect("/", fiber.StatusFound)
	})
}
