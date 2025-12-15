package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func NewRoute(app *fiber.App) {
	app.Use(logger.New())

	app.Post("/posts", CreatePostHandler)
	app.Put("/posts/:id", UpdatePostHandler)
	app.Delete("/posts/:id", DeletePostHandler)
	app.Get("/posts/:id", GetPostHandler)
	app.Get("/posts", GetPostsHandler)
}
