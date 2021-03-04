package main

import (
	"github.com/Falcer/elearning/server/module/auth"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	auth.New(app)
	app.Listen(":8080")
}
