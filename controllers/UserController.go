package controllers

import (
	"relasi-gorm/databases"
	"relasi-gorm/models"

	"github.com/gofiber/fiber/v2"
)

func UserGetAll(c *fiber.Ctx) error {
	var users []models.User
	databases.DB.Preload("Locker").Find(&users)

	return c.Status(200).JSON(fiber.Map{
		"users": users,
	})
}

func CreateUser(c *fiber.Ctx) error {
	user := new(models.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(503).JSON(fiber.Map{
			"err":"can't handle request",
		})
	}

	if user.Name == "" {
		return c.Status(400).JSON(fiber.Map{
			"err":"name is required",
		})
	}

	databases.DB.Create(&user)

	return c.JSON(fiber.Map{
		"message":"create data successfully",
		"user":user,
	})
}