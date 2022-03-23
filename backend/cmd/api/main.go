package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":2022"
		fmt.Println("🚀 Server run port:" + port)
	}

	// Create a new Fiber instance
	app := fiber.New()

	// Response with a hello message for calling root path
	app.Get("/", welcome)

	err := app.Listen(port)
	if err != nil {
		log.Fatal(err)
	}
}

func welcome(c *fiber.Ctx) error {
	return c.SendString("💃🏻 Welcome to my System 👋!")
}
