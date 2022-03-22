package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":2022"
		fmt.Println("Server run port: " + port)
	}

	// Create a new Fiber instance
	app := fiber.New()

	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Errorf("error running server")
		return
	}
}
