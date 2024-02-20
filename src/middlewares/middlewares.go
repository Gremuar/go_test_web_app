package middlewares

import "github.com/gofiber/fiber/v2"

// This middleware authentication is minial
// It cehcks if the `x-api-key` heeader is empty.
// You can update the middleware and make it work with your own
// authentication philosophy

func AuthMiddleware() fiber.Handler {
 return func(c *fiber.Ctx) error {
  apiKey := c.GetRespHeader("x-api-key")
  if apiKey == "" {
	return c.Status(401).JSON(fiber.Map{"error": "Unauthorized to perform request. Please get a valid API key"})
  }
  return c.Next()
 }
}

// Register middleware on the base router
func RegisterMiddlewares(app *fiber.App) {
	//app.Use(AuthMiddleware())
}