package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

var Store = session.New()

func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		sess, err := Store.Get(c)
		if err != nil {
			return c.Redirect("/db/admin/sign_in")
		}

		auth := sess.Get("authenticated")
		if auth == nil {
			return c.Redirect("/db/admin/sign_in")
		}

		return c.Next()
	}
}
