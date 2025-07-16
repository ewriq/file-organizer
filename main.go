package main

import (
	database "file-organizer/database"
	"file-organizer/handler"
	"file-organizer/middleware"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v3/middleware/helmet"
	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	app.Use(middleware.Cors)
	app.Use(middleware.RateLimit)
	app.Use(helmet.New())
	app.Use(middleware.Compress)
	app.Use(middleware.Logger)
	app.Use(middleware.Security)

	app.Get("/", handler.Home)

	//app.Get("/api/v1/files", handler.GetFiles)
	app.Get("/api/v1/files/:id", handler.GetFilesByType)
	//app.Post("/api/v1/files/delete", handler.DeleteFile)

	go database.Connect()
	go Cronjob()
	app.Use(middleware.NotFound)

	app.Listen(":3000")
}
