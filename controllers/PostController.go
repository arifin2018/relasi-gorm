package controllers

import (
	"relasi-gorm/databases"
	"relasi-gorm/models"

	"github.com/gofiber/fiber/v2"
)

func PostGetAll(c *fiber.Ctx) error {
	var Posts []models.Post
	databases.DB.Preload("User").Find(&Posts)

	return c.Status(200).JSON(fiber.Map{
		"Post": Posts,
	})
}

func CreatePost(c *fiber.Ctx) error {
	post := new(models.Post)

	if err := c.BodyParser(post); err != nil {
		return c.Status(503).JSON(fiber.Map{
			"err":"can't handle request",
		})
	}

	if post.Title == "" || post.Body == "" {
		return c.Status(400).JSON(fiber.Map{
			"err":"title and body is required",
		})
	}

	databases.DB.Create(&post)

	return c.JSON(fiber.Map{
		"message":"create data successfully",
		"post":post,
	})
}