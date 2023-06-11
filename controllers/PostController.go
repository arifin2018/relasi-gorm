package controllers

import (
	"relasi-gorm/databases"
	"relasi-gorm/models"

	"github.com/gofiber/fiber/v2"
)

func PostGetAll(c *fiber.Ctx) error {
	var Posts []models.ResponsePostWithTag
	databases.DB.Preload("User").Preload("Tag").Find(&Posts)

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

	if post.Title == "" || post.Body == "" || post.TagID == nil || len(post.TagID) <= 0 {
		return c.Status(400).JSON(fiber.Map{
			"err":"title,body or tagId is required",
		})
	}

	// postt := models.Post{ID: 8}

	// err := databases.DB.Model(&postt).Preload("Tag").Find(&postt)
	// log.Println(err.Error)
	// return nil

	databases.DB.Debug().Create(&post)

	if len(post.TagID) > 0 {
		for _, tagId := range post.TagID {
			postTag := new(models.PostTag)
			postTag.PostId = post.ID
			postTag.TagId = tagId 
			databases.DB.Create(&postTag)
		}
	}


	return c.JSON(fiber.Map{
		"message":"create data successfully",
		"post":post,
	})
}