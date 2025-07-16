package handler

import (
	"file-organizer/database"
	"fmt"

	"github.com/gofiber/fiber/v3"
)

func GetFilesByType(c fiber.Ctx) error {
	typeName := c.Params("type")
	fmt.Println("Type:", typeName)
	files, err := database.GetFilesByExtension()
	if err != nil {

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Dosyalar alınamadı"})
	}
	if len(files) == 0 {
		fmt.Println(err)
		fmt.Println(files)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Dosya bulunamadı"})
	}

	return c.JSON(files)
}
