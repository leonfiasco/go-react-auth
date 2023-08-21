package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/leonfiasco/go-auth/database"
	"github.com/leonfiasco/go-auth/routes"
)




func main() {
	database.ConnectDb()
    app := fiber.New()

	app.Use(cors.New(cors.Config{


		// allows the cookie to be received on the FE
	AllowCredentials: true,
	}))

	routes.SetUpRoutes(app)

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World 👋!")
    })

   
	log.Fatal(app.Listen(":2402"))
}