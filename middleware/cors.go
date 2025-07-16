package middleware

import "github.com/gofiber/fiber/v3/middleware/cors"

var Cors = cors.New(cors.Config{
    AllowOrigins: []string{"http://localhost:4000"},
    AllowHeaders: []string{"Origin", "Content-Type", "Accept"},
})