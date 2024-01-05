package handlers

import (
	"github.com/gofiber/fiber/v2"
)

// Will trick go into adding to HandlersList Array without executing a function 
var setupSystemd = AddHandler(func() {
	api := app.Group("/systemd")
	api.Get("/ola", Ola)
})

func Ola(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "ola"})
}
