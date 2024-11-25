package routers

import (
	"Project_PBO/controllers"
	"Project_PBO/handlers"
	"Project_PBO/middleware"

	"github.com/gofiber/fiber/v2"
)

func Route(app *fiber.App) {

	app.Get("db/admin/sign_in", handlers.SignIn)
	app.Get("db/admin/sign_up", handlers.SignUp)
	app.Post("/signup", controllers.SignUp)
	app.Post("/login", controllers.Login)

	siswaGroup := app.Group("/api/siswa")
	siswaGroup.Get("/showAll", controllers.SiswaControllerShow)
	siswaGroup.Post("/create", controllers.SiswaControllerCreate)
	siswaGroup.Delete("/delete/:id", controllers.SiswaControllerDelete)
	siswaGroup.Put("/update/:id", controllers.SiswaControllerUpdate)
	siswaGroup.Get("/db/admin/siswa/:id", controllers.SiswaControllerShowByID)

	beritaGroup := app.Group("/api/berita")
	beritaGroup.Get("/showAll", controllers.BeritaControllerShow)
	beritaGroup.Post("/create", controllers.BeritaControllerCreate)
	beritaGroup.Delete("/delete/:id", controllers.BeritaControllerDelete)
	beritaGroup.Put("/update/:id", controllers.BeritaControllerUpdate)
	beritaGroup.Get("/show/:id", controllers.BeritaControllerShowByID)

	prestasiGroup := app.Group("/api/prestasi")
	prestasiGroup.Get("/showAll", controllers.PrestasiControllerShow)
	prestasiGroup.Post("/create", controllers.PrestasiControllerCreate)
	prestasiGroup.Put("/update/:id", controllers.PrestasiControllerUpdate)
	prestasiGroup.Delete("/delete/:id", controllers.PrestasiControllerDelete)

	EkskulGroup := app.Group("/api/ekskul")
	EkskulGroup.Get("/showAll", controllers.EkskulControllerShow)
	EkskulGroup.Post("/create", controllers.EkskulControllerCreate)
	EkskulGroup.Put("/update/:id", controllers.EkskulControllerUpdate)
	EkskulGroup.Delete("/delete/:id", controllers.EkskulControllerDelete)
	EkskulGroup.Get("/show/:id", controllers.EkskulControllerShowByID)

	app.Get("/static/*", handlers.StaticHandler)
	app.Get("/pkg/*", handlers.PkgHandler)
	app.Get("/uploads/*", handlers.UploadsHandler)

	app.Get("/", handlers.Home)

	admin := app.Group("/db/admin", middleware.AuthMiddleware())
	admin.Get("/dashboard", handlers.Dashboard)
	admin.Get("/prestasi", handlers.FormPrestasi)
	admin.Get("/infoSiswa", handlers.InfoSiswa)
	admin.Get("/berita", handlers.FormBerita)
	admin.Get("/ekskul", handlers.FormEkskul)
	app.Post("/Sendmail", handlers.HandleContactForm)

	//Buat user gabisa ke mana mana selain yang ada di route
	app.Use(func(c *fiber.Ctx) error {
		return c.Redirect("/", fiber.StatusFound)
	})
}
