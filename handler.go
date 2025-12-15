package main

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreatePostHandler(c *fiber.Ctx) error {
	postRequest := new(Request)

	err := c.BodyParser(postRequest)
	if err != nil {
		badRequestResponse := NewInternalErrorResponse()
		return c.Status(fiber.StatusBadRequest).JSON(badRequestResponse)
	}

	var tags []Tag
	for _, tag := range postRequest.Tags {
		tempTag := Tag{
			Name: tag,
		}
		tags = append(tags, tempTag)
	}

	post := Post{
		Title:    postRequest.Title,
		Content:  postRequest.Content,
		Category: postRequest.Category,
		Tags:     tags,
	}

	createdPost, err := InsertNewPost(post)
	if err != nil {
		internalErrorResponse := NewInternalErrorResponse()
		return c.Status(fiber.StatusInternalServerError).JSON(internalErrorResponse)
	}

	successRespones := NewSuccessResponse(NewPostResponseFormat(createdPost), "new post created")
	return c.Status(fiber.StatusCreated).JSON(successRespones)
}

func UpdatePostHandler(c *fiber.Ctx) error {
	postRequest := new(Request)

	err := c.BodyParser(postRequest)
	if err != nil {
		badRequestResponse := NewInternalErrorResponse()
		return c.Status(fiber.StatusBadRequest).JSON(badRequestResponse)
	}

	postID, err := c.ParamsInt("id")
	if err != nil {
		badRequestResponse := NewBadRequestResponse()
		return c.Status(fiber.StatusBadRequest).JSON(badRequestResponse)
	}

	var tags []Tag
	for _, tag := range postRequest.Tags {
		tempTag := Tag{
			Name: tag,
		}
		tags = append(tags, tempTag)
	}

	post := Post{
		Title:    postRequest.Title,
		Content:  postRequest.Content,
		Category: postRequest.Category,
		Tags:     tags,
	}

	updatedPost, err := UpdatePost(postID, post)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			notFoundResponse := NewNotFoundResponse()
			return c.Status(fiber.StatusNotFound).JSON(notFoundResponse)
		}

		internalErrorResponse := NewInternalErrorResponse()
		return c.Status(fiber.StatusInternalServerError).JSON(internalErrorResponse)
	}

	successRespones := NewSuccessResponse(NewPostResponseFormat(updatedPost), "new post updated")
	return c.Status(fiber.StatusOK).JSON(successRespones)
}

func DeletePostHandler(c *fiber.Ctx) error {
	postID, err := c.ParamsInt("id")
	if err != nil {
		badRequestResponse := NewBadRequestResponse()
		return c.Status(fiber.StatusBadRequest).JSON(badRequestResponse)
	}

	err = DeletePost(postID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			notFoundResponse := NewNotFoundResponse()
			return c.Status(fiber.StatusNotFound).JSON(notFoundResponse)
		}

		internalErrorResponse := NewInternalErrorResponse()
		return c.Status(fiber.StatusInternalServerError).JSON(internalErrorResponse)
	}

	successRespones := NewSuccessResponse(nil, "post deleted")
	return c.Status(fiber.StatusNoContent).JSON(successRespones)
}

func GetPostHandler(c *fiber.Ctx) error {
	postID, err := c.ParamsInt("id")
	if err != nil {
		badRequestResponse := NewBadRequestResponse()
		return c.Status(fiber.StatusBadRequest).JSON(badRequestResponse)
	}

	post, err := GetPost(postID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			notFoundResponse := NewNotFoundResponse()
			return c.Status(fiber.StatusNotFound).JSON(notFoundResponse)
		}

		internalErrorResponse := NewInternalErrorResponse()
		return c.Status(fiber.StatusInternalServerError).JSON(internalErrorResponse)
	}

	successRespones := NewSuccessResponse(NewPostResponseFormat(post), "get post success")
	return c.Status(fiber.StatusOK).JSON(successRespones)
}

func GetPostsHandler(c *fiber.Ctx) error {
	term := c.Query("term", "")

	posts, err := GetPosts(term)
	if err != nil {
		internalErrorResponse := NewInternalErrorResponse()
		return c.Status(fiber.StatusInternalServerError).JSON(internalErrorResponse)
	}

	if len(posts) == 0 {
		notFoundResponse := NewNotFoundResponse()
		return c.Status(fiber.StatusNotFound).JSON(notFoundResponse)
	}

	successRespones := NewSuccessResponse(NewPostsResponseFormat(posts), "get all post success")
	return c.Status(fiber.StatusOK).JSON(successRespones)
}
