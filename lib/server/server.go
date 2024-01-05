package server

// Should only be imported by app.go

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"strings"
	"time"

	"github.com/3n3a/server-control-api/handlers"
	"github.com/3n3a/server-control-api/lib/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

const (
	ENV_NAME = "ENVIRONMENT"
)

type AppConfig struct {
	VERSION           string
	CACHE_INCLUDE_RAW string
	CACHE_INCLUDE     []string
	CACHE_LENGTH      time.Duration
	APP_PORT          int
	ENVIRONMENT       string
}

func (a *AppConfig) Setup() {
	// Setup Cache Includes
	a.CACHE_INCLUDE = slices.DeleteFunc(
		strings.Split(a.CACHE_INCLUDE_RAW, ";"),
		func(e string) bool {
			return e == ""
		},
	)

	// ENv
	a.ENVIRONMENT = os.Getenv("ENVIRONMENT")

	// Print config
	fmt.Printf("=== Server Configuration ===\n")
	configJson, _ := json.MarshalIndent(a, "", "  ")
	fmt.Printf("%s\n", configJson)

	// Set Version to DEV
	if utils.IsDev() {
		a.VERSION = "devel"
	}

	// start Gofiber server
	a.setupServer()
}

func (a *AppConfig) setupServer() {
	

	// Create fiber app
	app := fiber.New(fiber.Config{})

	// Middleware
	app.Use(recover.New())
	app.Use(logger.New())
	if !utils.IsDev() {
		app.Use(compress.New())
		app.Use(cache.New(cache.Config{
			Next: func(c *fiber.Ctx) bool {
				for _, pathMatch := range a.CACHE_INCLUDE {
					match, _ := regexp.MatchString(pathMatch, c.Path())
					if match {
						return false // cached
					}
				}
				return true // not cached
			},
			Expiration:   a.CACHE_LENGTH,
			CacheControl: true,
			KeyGenerator: func(c *fiber.Ctx) string {
				return c.OriginalURL()
			},
		}))
	}

	// Setup routes & configure handlers
	handlers.Setup(app)

	// Handle not founds
	app.Use(handlers.NotFound)

	// Listen on port 3000
	log.Fatal(app.Listen(fmt.Sprintf(":%d", a.APP_PORT)))
}
