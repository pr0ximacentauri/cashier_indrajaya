package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"cashier_indrajaya/config"
	"cashier_indrajaya/routes"
	"log"
)

func main() {
	// Init Firebase
	config.InitFirebase()
	defer config.Firestore.Close()

	app := fiber.New()
	
	app.Use(cors.New(cors.Config{
        AllowOrigins: "http://localhost:3000",
        AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
    }))

	routes.SetupRoutes(app)

	log.Println("Server running on http://localhost:8080")
	app.Listen(":8080")
	app.Static("/", "./web/build")
}
