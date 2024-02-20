package controllers

import (
	//"net/http"

	"main/src/models"

	"github.com/gofiber/fiber/v2"
)

func GetAllUsers(c *fiber.Ctx) error {
	users, err := models.FetchAllUsers()
	if err != nil {
		return c.JSON(fiber.Map{"message": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Users fetched successfully", "status": "success", "data": users})
}

func GetUserByID(c *fiber.Ctx) error {
	userID := c.AllParams()["id"]
	if userID == "" {
		return c.JSON(fiber.Map{"message": "User ID is required"})
	}
	user, err := models.FetchUser(userID)
	if err != nil {
		return c.JSON(fiber.Map{"message": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "User fetched successfully", "status": "success", "data": user})
}

func CreateUser(c *fiber.Ctx) error {
	userModel := new(models.User)

	if err := c.BodyParser(userModel); err != nil {
		return c.JSON(fiber.Map{"status": "failed parse model", "message": err.Error(), "data": nil})
	}

	savedUser, err := userModel.Save()

	if err != nil {
		return c.JSON(fiber.Map{"status": "failed save in db", "message": err.Error(), "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "User saved successfully", "data": savedUser})
}

func UpdateUser(c *fiber.Ctx) error {
	UserID := c.AllParams()["id"]

	var updatedUser *models.User
	if err := c.JSON(&updatedUser); err != nil {
		return c.JSON(fiber.Map{"status": "failed", "message": err.Error(), "data": nil})
	}

	updatedUser, err := updatedUser.UpdateUser(UserID)
	if err != nil {
		return c.JSON(fiber.Map{"status": "failed", "message": err.Error(), "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "User updated successfully", "data": updatedUser})
}

func DeleteUser(c *fiber.Ctx) error {
	UserID := c.AllParams()["id"]
	if UserID == "" {
		return c.JSON(fiber.Map{"message": "User ID is required"})
	}

	err := models.DeleteUser(UserID)
	if err != nil {
		return c.JSON(fiber.Map{"message": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "User deleted successfully", "status": "success", "data": nil})
}
