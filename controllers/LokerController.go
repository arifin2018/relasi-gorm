package controllers

import (
	"relasi-gorm/databases"
	"relasi-gorm/models"

	"github.com/gofiber/fiber/v2"
)

func LokerGetAll(c *fiber.Ctx) error {
	
	var locker []models.Locker

	databases.DB.Preload("User").Find(&locker)

	return c.JSON(fiber.Map{
		"locker": locker,
	})
}

func CreateLoker(c *fiber.Ctx) error {
	locker := new(models.Locker)

	if err := c.BodyParser(locker); err != nil {
		return c.Status(503).JSON(fiber.Map{
			"err":"can't handle request",
		})
	}

	if locker.Code == "" {
		return c.Status(400).JSON(fiber.Map{
			"err":"name is required",
		})
	}

	databases.DB.Preload("User").Create(&locker)

	return c.JSON(fiber.Map{
		"message":"create data successfully",
		"locker":locker,
	})
}