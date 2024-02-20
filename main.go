package main

import (
	"github.com/gofiber/fiber/v2"

	"main/src/middlewares"
	"main/src/models"
	"main/src/routes"
	"main/src/utils"
)

func main() {
	app := fiber.New()
	utils.LoadEnv()
	models.OpenDatabaseConnection()
	models.AutoMigrateModels()
	routes.SetupRoutes(app)
	middlewares.RegisterMiddlewares(app)

	app.Listen(":3000")
}
