package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Falcer/elearning/server/auth"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

var (
	port        string
	databaseURL string
	authRepo    auth.Repository
	authService auth.Service
)

func init() {
	// Load ENV
	err := godotenv.Load("./.env")
	if err != nil {
		log.Println("Error load .env")
	}
	// Get PORT
	port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	// Get DATABASE_URL
	databaseURL = os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "" // TODO: Set default url
	}
	// Set UserRePository
	authRepo = auth.NewRepo(databaseURL)
	// Set UserService
	authService = auth.NewService(authRepo)
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(&fiber.Map{
			"message": "app running 🔥",
		})
	})

	// auth
	app.Post("/login", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(&fiber.Map{
			"message": "Login is running 🔥",
		})
	})
	app.Post("/register", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(&fiber.Map{
			"message": "Register is running 🔥",
		})
	})
	app.Get("/user", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(&fiber.Map{
			"message": "Users is running 🔥",
		})
	})
	// UserRole Add
	app.Post("/user/role", func(c *fiber.Ctx) error {
		userrole := new(auth.UserRoleInput)
		if err := c.BodyParser(userrole); err != nil {
			return err
		}
		return c.Status(200).JSON(&fiber.Map{
			"message": "User role added is running 🔥",
			"data":    &userrole,
		})
	})
	app.Delete("/user/role", func(c *fiber.Ctx) error {
		userrole := new(auth.UserRoleInput)
		if err := c.BodyParser(userrole); err != nil {
			return err
		}
		return c.Status(200).JSON(&fiber.Map{
			"message": "User role deleted is running 🔥",
			"data":    &userrole,
		})
	})

	// Role
	app.Get("/role", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(&fiber.Map{
			"message": "Roles is running 🔥",
		})
	})
	app.Get("/role/:id", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(&fiber.Map{
			"message": fmt.Sprintf("Role is running 🔥, id : %s", c.Params("id")),
		})
	})
	app.Post("/role", func(c *fiber.Ctx) error {
		role := new(auth.RoleInput)
		if err := c.BodyParser(role); err != nil {
			return err
		}
		return c.Status(200).JSON(&fiber.Map{
			"message": "Role added is running 🔥",
			"data":    &role,
		})
	})
	app.Delete("/role/:id", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(&fiber.Map{
			"message": fmt.Sprintf("Role is running 🔥, deleted id : %s", c.Params("id")),
		})
	})

	app.Listen(":8080")
}
