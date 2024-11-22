package main

import (
	"app/config"
	"app/router"
	"log"

	"github.com/gofiber/fiber/v2"
	// "github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	config.SetupEnvFile()
	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:       "Redpay",
	})
	// app := fiber.New()
	// app.Use(cors.New())

	// database.ConnectDB()
	// database.SetupMongoDB()

	router.SetupRoutes(app)
	log.Fatal(app.Listen(":4000"))
}
