package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/orhanince/go-wikiquote/src/model"
)

func Home(c *fiber.Ctx) error {
	err := c.Status(fiber.StatusOK).JSON(model.Response{
		Status:  true,
		Message: "Success",
		Data: map[string]interface{}{
			"name": "Home",
		}})

	return err
}

func Check(c *fiber.Ctx) error {
	err := c.Status(fiber.StatusOK).JSON(model.Response{
		Status:  true,
		Message: "Success",
		Data: map[string]interface{}{
			"name": "User",
		}})

	return err
}

func Test(c *fiber.Ctx) error {
	err := c.Status(fiber.StatusOK).JSON(model.Response{
		Status:  true,
		Message: "Success",
		Data: map[string]interface{}{
			"name": "User Test!",
		}})

	return err
}
