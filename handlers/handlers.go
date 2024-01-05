package handlers

import (
	"github.com/3n3a/server-control-api/lib/docker"
	"github.com/gofiber/fiber/v2"
)

type HandlersList []func()

var app *fiber.App
var handlersList HandlersList
var innerConfig InnerConfig

type InnerConfig struct {
	CONTAINER_TYPE docker.ContainerTool
}

func Setup(currentApp *fiber.App, innerConfig_ InnerConfig) {
	app = currentApp
	innerConfig = innerConfig_

	// Register all Handlers
	for _, handlerFunc := range handlersList {
		handlerFunc()
	}
}

func AddHandler(f func()) bool {
	handlersList = append(handlersList, f)
	return true
}