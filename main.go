package main

import (
	"log"
	"relasi-gorm/databases"
	"relasi-gorm/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	databases.DatabaseInit()
	
	app := fiber.New()
	routes.Router(app)
	
	log.Fatal(app.Listen(":8000"))
}