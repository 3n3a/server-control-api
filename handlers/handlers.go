package handlers

import (
	"github.com/gofiber/fiber/v2"
)

type HandlersList []func()

var app *fiber.App
var handlersList HandlersList

func Setup(currentApp *fiber.App) {
	app = currentApp

	// Register all Handlers
	for _, handlerFunc := range handlersList {
		handlerFunc()
	}
}

func AddHandler(f func()) bool {
	handlersList = append(handlersList, f)
	return true
}