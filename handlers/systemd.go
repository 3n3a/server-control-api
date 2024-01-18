package handlers

import (
	"fmt"

	"github.com/3n3a/server-control-api/lib/systemd"
	"github.com/3n3a/server-control-api/lib/utils"
	"github.com/gofiber/fiber/v2"
)

var systemdConn systemd.SystemDConn

// // Will trick go into adding to HandlersList Array without executing a function 
var setupSystemd = AddHandler(func() {
	systemdConn = systemd.New()

	api := app.Group("/systemd")

	api.Post("/restart", RestartService)
})

// Restart Service
// Request
// name: Name of service
// Responses:
// 200: Successfully restarted service
// 500: service could not be restarted
func RestartService(c *fiber.Ctx) error {
	serviceName := c.FormValue("name")
	err := systemdConn.RestartService(serviceName)
	if err != nil {
		if utils.IsDev() {
			fmt.Println(err)
		}
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
			"status": "error",
		})
	}
	return c.JSON(fiber.Map{"message": "service successfully restarted", "status": "success"})
}
