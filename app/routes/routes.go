package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jeffthorne/tasky/app/auth"
	_ "github.com/jeffthorne/tasky/app/auth"
	_ "os"
)

func SetupRoutes(app *fiber.App) {

	app.Get("/auth", auth.Auth)
	app.Post("/login", auth.Login)
	app.Get("/register", auth.Register)
	app.Post("/signup", auth.Signup)

	//app.Static("/docs", "./docs")

}
