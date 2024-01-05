package handlers

import (
	"fmt"

	"github.com/3n3a/server-control-api/lib/docker"
	"github.com/3n3a/server-control-api/lib/utils"
	"github.com/gofiber/fiber/v2"
)

var dockerConn docker.Client

// // Will trick go into adding to HandlersList Array without executing a function 
var setupDocker = AddHandler(func() {
	dockerConn = docker.New(innerConfig.CONTAINER_TYPE)

	api := app.Group("/docker")

	imagesApi := api.Group("/images")
	imagesApi.Post("/pull", PullImage)
})

// Pull Image
// Request
// image: Param in form of "docker.io/library/busybox" with optional tag (default is latest). With tag "docker.io/library/busybox:latest" 
// Responses:
// 200: Successfully Pulled Image
// 500: Some Error with Docker Tool occurred
func PullImage(c *fiber.Ctx) error {
	imageUrl := c.FormValue("image")
	err := dockerConn.PullImage(imageUrl)
	if err != nil {
		if utils.IsDev() {
			fmt.Println(err)
		}
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
			"status": "error",
		})
	}
	return c.JSON(fiber.Map{"message": "image successfully pulled", "status": "success"})
}
