package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func healthCheck(c *fiber.Ctx) error {
	return c.SendString("ok")
}

func main() {
	app := fiber.New()

	app.Get("/", healthCheck)

	fmt.Println("Server is running on port 3000")

	log.Fatal(app.Listen(":3000"))

}
