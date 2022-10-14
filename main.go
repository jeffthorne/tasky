package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/jeffthorne/tasky/app/database"
	"github.com/jeffthorne/tasky/app/routes"
)

func main() {
	engine := html.New("./app/views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	routes.SetupRoutes(app)
	database.InitMongo()

	if err := app.Listen(":3000"); err != nil {
		fmt.Println("ERROR: ", err)
	}
}
