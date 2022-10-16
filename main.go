package main

import (
	"log"
	"os"

	"github.com/TomislavGalic/cameras/database"
	"github.com/TomislavGalic/cameras/routes"
	"github.com/gofiber/template/html"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file.")
	}

	engine := html.New("./templates", ".html")

	database.Connect()
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	routes.Create(app)
	app.Listen(os.Getenv("LOCALHOST"))
}
