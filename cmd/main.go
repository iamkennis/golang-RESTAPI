package main

import (
	"fmt"
	"log"

	"github.com/GoNodeapp/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func healthcheck(c *fiber.Ctx) error {
	return c.SendString("OK")
}

func main() {
	app := fiber.New()

	app.Get("/healthcheck", healthcheck)
	app.Post("/api/products", handlers.CreateProduct)
	app.Get("/api/products", handlers.GetAllProducts)

	fmt.Println("Hello world")

	log.Fatal(app.Listen(":3000"))
}
