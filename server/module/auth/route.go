package auth

import "github.com/gofiber/fiber/v2"

type app struct {
	service Service
}

func (a *app) login(c *fiber.Ctx) error {
	return c.Status(200).JSON(&fiber.Map{
		"message": "login running",
	})
}

func (a *app) register(c *fiber.Ctx) error {
	return c.Status(200).JSON(&fiber.Map{
		"message": "register running",
	})
}

func (a *app) refresh(c *fiber.Ctx) error {
	return c.Status(200).JSON(&fiber.Map{
		"message": "refresh token running",
	})
}

func (a *app) health(c *fiber.Ctx) error {
	return c.Status(200).JSON(&fiber.Map{
		"health": "UP",
	})
}

// New Router
func New(route *fiber.App) {
	service := &app{}
	group := route.Group("/auth")
	group.Post("/login", service.login)
	group.Post("/register", service.register)
	group.Post("/refresh", service.refresh)
	group.Get("/health", service.health)
}
