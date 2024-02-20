package routes

import (
	"fmt"
	"main/src/controllers"

	"github.com/gofiber/fiber/v2"
)

func middleware(c *fiber.Ctx) error {
	fmt.Println("Don't mind me!")
	//Эта ф-ция вызывается при обработке каждого запроса пути /startups/*
	return c.Next()
}

func SetupRoutes(app *fiber.App) fiber.Router {
	users := app.Group("/users", middleware)

	apiv1 := users.Group("/api/v1", middleware)
	apiv1.Get("/all", controllers.GetAllUsers)
	apiv1.Get("/get/:id", controllers.GetUserByID)
	apiv1.Post("/create", controllers.CreateUser)
	apiv1.Patch("/update", controllers.UpdateUser)
	apiv1.Delete("/delete/:id", controllers.DeleteUser)

	return apiv1
}
