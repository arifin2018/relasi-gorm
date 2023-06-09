package routes

import (
	"relasi-gorm/controllers"

	"github.com/gofiber/fiber/v2"
)


func Router(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("I'm a GET request!")
	})
	app.Get("/users", controllers.UserGetAll)
	app.Post("/users", controllers.CreateUser)
	
	app.Get("/loker", controllers.LokerGetAll)
	app.Post("/loker", controllers.CreateLoker)

	app.Get("/posts", controllers.PostGetAll)
	app.Post("/posts", controllers.CreatePost)

	app.Get("/tags", controllers.TagsGetAll)
	app.Post("/tags", controllers.CreateTags)
}