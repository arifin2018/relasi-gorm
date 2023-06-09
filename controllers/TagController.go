package controllers

import (
	"relasi-gorm/databases"
	"relasi-gorm/models"

	"github.com/gofiber/fiber/v2"
)

func TagsGetAll(c *fiber.Ctx) error {
	var Tag []models.Tag
	databases.DB.Find(&Tag)

	return c.Status(200).JSON(fiber.Map{
		"Tag": Tag,
	})
}

func CreateTags(c *fiber.Ctx) error {
	Tag := new(models.Tag)

	if err := c.BodyParser(Tag); err != nil {
		return c.Status(503).JSON(fiber.Map{
			"err":"can't handle request",
		})
	}

	if Tag.Name == "" {
		return c.Status(400).JSON(fiber.Map{
			"err":"title and body is required",
		})
	}

	databases.DB.Create(&Tag)

	return c.JSON(fiber.Map{
		"message":"create data successfully",
		"tag":Tag,
	})
}