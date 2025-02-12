package main

import (
	"log"

	"github.com/alexesp/Go_FibePostgreSQL_Docker.git/database"
	"github.com/gofiber/fiber/v2"
)
func welcome(c *fiber.Ctx) error{
	return c.SendString("Welcome to fiber")
}

func setupRoutes(app *fiber.App){
	
	app.Get("/", welcome)
}

func main(){
	database.ConnectDb()
	app := fiber.New()

	//app.Get("/", welcome)
	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}