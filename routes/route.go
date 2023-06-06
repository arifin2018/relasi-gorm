package routes

import "github.com/gofiber/fiber/v2"


func Router() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "i'm not okay",
			"name":    "arifin",
		})
	})

	app.Listen(":8000")
}