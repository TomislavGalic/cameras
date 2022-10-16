package routes

import (
	"github.com/TomislavGalic/cameras/controllers"
	"github.com/gofiber/fiber/v2"
)

func Create(app *fiber.App) {

	app.Get("/", controllers.Home)
	app.Get("/list", controllers.List)
	app.Get("/*", controllers.NotFound)
}
