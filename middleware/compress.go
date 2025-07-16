package middleware

import(
	"github.com/gofiber/fiber/v3/middleware/compress"
	"github.com/gofiber/fiber/v3"
)


func Compress(c fiber.Ctx) error {
	compress.New(compress.Config{
		Level: compress.LevelBestCompression,
	})
	return c.Next()
}