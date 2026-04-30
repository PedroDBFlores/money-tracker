package handler

import (
	"money-tracker/api/internal/user/service"

	"github.com/gofiber/fiber/v3"
)

func UserRouter(app *fiber.App, service *service.UserService) {
	app.Route("/users", func(router fiber.Router) {
		router.Post("/", service.CreateUser)      // Needs handler
		router.Get("/:email", service.GetByEmail) // Needs handler
	})
}
