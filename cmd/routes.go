package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ahsan-sabri/gorest/handlers"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.Home)

	app.Get("/facts", handlers.ListFact)
	app.Post("/fact", handlers.CreateFact)
}