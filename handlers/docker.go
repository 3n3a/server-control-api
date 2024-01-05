package handlers

import (
	"github.com/3n3a/server-control-api/lib/docker"
	"github.com/gofiber/fiber/v2"
)

var dockerConn docker.Client

// // Will trick go into adding to HandlersList Array without executing a function 
var setupDocker = AddHandler(func() {
	dockerConn = docker.New(docker.DockerTool)

	api := app.Group("/podman")
	api.Post("/pull", PullImage)
})

func PullImage(c *fiber.Ctx) error {
	imageUrl := c.FormValue("image")
	_, err := dockerConn.PullImage(imageUrl)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "image could not be pulled",
			"status": "error",
		})
	}
	return c.JSON(fiber.Map{"message": "image successfully pulled", "status": "success"})
}
