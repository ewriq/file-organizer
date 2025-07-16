package middleware

import (
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/limiter"
)

var RateLimit = limiter.New(limiter.Config{
	Max:        100,
	Expiration: 1 * time.Minute,
	KeyGenerator: func(c fiber.Ctx) string {
		return c.IP()
	},
	LimitReached: func(c fiber.Ctx) error {
		return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
			"msg": "sonra dene amk.",
		})
	},
})
