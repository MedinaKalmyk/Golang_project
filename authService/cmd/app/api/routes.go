package api

import (
	"Go/cmd/controllers"
	"github.com/gofiber/fiber/v2"
)

var authHandler = controllers.AuthHandler{}

func Routes(app *fiber.App) {
	publicGroup := app.Group("/api")
	publicGroup.Post("/registration", authHandler.Registration)
	publicGroup.Post("/login", authHandler.Login)
}
